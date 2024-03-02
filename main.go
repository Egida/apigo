package main

import (
	"api/active"
	"api/cloudns"
	"api/combahton"
	"api/controller"
	"api/database"
	"api/db1"
	"api/dynadot"
	apierror "api/error"
	"api/middleware"
	"api/model"
	"api/pdns"
	"api/synlinq"
	"api/vpn"
	"fmt"
	"os"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"

	//"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2/middleware/earlydata"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/gofiber/template/html"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/spf13/viper"
)

var Commit string = "dev"
var json = jsoniter.ConfigCompatibleWithStandardLibrary

func main() {
	loadEnv()
	setupLogging()
	loadDatabase()
	loadIntegrations()
	model.InitValidator()

	serveApplication()
}

func loadEnv() {
	viper.SetConfigName("app") // config file name without extension
	viper.SetConfigType("yaml")

	viper.AddConfigPath("./config") // config file path
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv() // read value ENV variable

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		os.Exit(1)
	}
}

func setupLogging() {
	path := viper.GetString("app.logfile")
	if path == "" {
		path = "logs/dnic.log"
	}

	file := &lumberjack.Logger{
		Filename:   path,
		MaxBackups: 10, // files
		MaxAge:     1,  // days
	}

	w := zerolog.MultiLevelWriter(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.TimeOnly},
		zerolog.ConsoleWriter{Out: file, TimeFormat: time.TimeOnly, NoColor: true},
	)

	log.Logger = zerolog.New(w).With().Timestamp().Caller().Logger()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{}, &model.APIKey{}, &model.IP{})
	model.SetupInitialAdmin()
	db1.Connect()
}

func loadIntegrations() {
	combahton.Init()
	dynadot.Init()
	cloudns.Init()
	active.Init()
	vpn.Init()
	pdns.Init()
	synlinq.Init()
}

func serveApplication() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Prefork:                      false,
		ServerHeader:                 "DNIC",
		StrictRouting:                false,
		CaseSensitive:                false,
		Immutable:                    false,
		UnescapePath:                 false,
		ETag:                         false,
		BodyLimit:                    0,
		Concurrency:                  0,
		Views:                        engine,
		ViewsLayout:                  "layouts/main",
		PassLocalsToViews:            false,
		ReadTimeout:                  0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		ReadBufferSize:               0,
		WriteBufferSize:              0,
		CompressedFileSuffix:         "",
		JSONEncoder:                  json.Marshal,
		JSONDecoder:                  json.Unmarshal,
		ProxyHeader:                  "", //fiber.HeaderXForwardedFor,
		GETOnly:                      false,
		ErrorHandler:                 apierror.Handler,
		DisableKeepalive:             false,
		DisableDefaultDate:           false,
		DisableDefaultContentType:    false,
		DisableHeaderNormalizing:     false,
		DisableStartupMessage:        false,
		AppName:                      "DNIC-API",
		StreamRequestBody:            true,
		DisablePreParseMultipartForm: true,
		ReduceMemoryUsage:            false,

		Network:                 "",
		EnableTrustedProxyCheck: false,
		TrustedProxies:          []string{},
		EnableIPValidation:      false,
		EnablePrintRoutes:       false,
		ColorScheme:             fiber.Colors{},
		RequestMethods:          []string{},
	})

	app.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        240,
		Expiration: 60 * time.Second,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.SendFile("./toofast.html")
		},
	}))
	app.Use(earlydata.New())
	app.Use(requestid.New())
	app.Use(middleware.Logger)
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "PUT,PATCH,POST,GET,OPTIONS,DELETE,HEAD",
		AllowHeaders:     "bearer,X-Apikey, Content-Type, X-XSRF-TOKEN, Accept, Origin, X-Requested-With, Authorization",
		ExposeHeaders:    "Content-Length",
		AllowCredentials: true,
		MaxAge:           int(12 * time.Hour),
	}))

	app.Static("/", "./public")
	app.Get("/metrics", monitor.New())
	app.Get("/version", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"commit": Commit})
	})

	logstream := app.Group("logstream")
	logstream.Use(middleware.JWTAuthMiddleware)
	logstream.Use(middleware.BindCurrentUser)
	logstream.Use(middleware.RequireAdmin)
	logstream.Get("/", controller.LogStream)

	admin := app.Group("/admin")
	admin.Use(middleware.BindCurrentUser)
	admin.Get("/", controller.IndexView)
	admin.Get("/login", controller.LoginView)
	admin.Post("/login", controller.LoginForm)
	admin.Get("/logout", controller.LogoutForm)
	admin.Use(middleware.RequireAdmin)
	admin.Get("/accounts", controller.AccountsIndex)
	admin.Post("/accounts/new", controller.AccountCreateForm)
	admin.Post("/accounts/:id/edit", controller.AccountEditForm)
	admin.Get("/accounts/:id/revoke", controller.AccountRevoke)
	admin.Get("/accounts/:id/token/new", controller.CreateToken)
	admin.Get("/accounts/:id/delete", controller.AccountDeleteForm)
	admin.Get("/logs", controller.LogsView)

	auth := app.Group("/auth")
	auth.Post("/token", controller.Token)

	account := app.Group("/account")
	account.Use(middleware.APIKeyAuthMiddleware)
	account.Get("/info", controller.ShowUser)
	account.Post("/password", controller.ChangePassword)

	ddos := app.Group("/ddos")
	ddos.Use(middleware.APIKeyAuthMiddleware)
	//ddos.Use(middleware.RequireIP)
	ddos.Get("/status/:ip", controller.GetRouting)
	ddos.Get("/incidents/:ip", controller.GetIncidents)
	ddos.Put("/routing/:ip/:mask", controller.AddRouting)
	ddos.Post("/vhost/:ip/:mask", controller.AddVhost)

	path := app.Group("/path")
	path.Use(middleware.APIKeyAuthMiddleware)
	path.Use(middleware.RequireRole(model.RoleAdmin))
	//path.Use(middleware.RequireIP)
	path.Get("/incidents/:ip", controller.GetPathIncidents)
	path.Get("/rules/:ip", controller.GetPathRules)
	path.Post("/addrules/:ip", controller.AddPathRules)
	path.Delete("/rules/:ip/:id", controller.DeleteRule)
	path.Get("/ratelimits/:ip", controller.GetRateLimits)
	path.Post("/ratelimits/:ip", controller.AddRateLimit)
	path.Delete("/ratelimits/:ip/:id", controller.DeleterRateLimit)
	path.Get("/filters/:ip", controller.GetFilters)
	path.Get("/filters/available/:ip", controller.AvailableFilter)
	path.Post("/filters/:ip/:filter_type", controller.AddFilter)
	path.Delete("/filters/:ip/:id/:filter_type", controller.DeleteFilter)

	domain := app.Group("/domain")
	domain.Use(middleware.APIKeyAuthMiddleware)
	domain.Get("/whois/:domain", controller.Whois)
	domain.Post("/addContact", controller.AddContact)

	//vpn := app.Group("/vpn")
	//vpn.Use(middleware.APIKeyAuthMiddleware)
	//vpn.Get("/accounts", controller.GetAccounts)

	dns := app.Group("/dns")
	dns.Use(middleware.APIKeyAuthMiddleware)
	//dns.Post("/zone/:domain", controller.AddZone)
	dns.Post("/zone", controller.CZone)
	dns.Delete("/zone/:domain", controller.RemoveZone)
	dns.Post("/record/:domain", controller.AddRecord)
	dns.Delete("/record/:domain", controller.RemoveRecord)
	dns.Post("/ptr/:id", controller.ChangePtr)

	health := app.Group("/")
	health.Use(middleware.APIKeyAuthMiddleware)
	health.Get("/health", controller.Health)

	//ping := app.Group("/ping")
	//ping.Use(middleware.JWTAuthMiddleware)
	//ping.Get("/host/:host/:port", controller.Ping)
	//app.Listen(":8000")
	app.ListenTLS(":443", "./fullchain.crt", "./fullchain.key")
}
