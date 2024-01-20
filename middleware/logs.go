package middleware

import (
	"api/model"
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type logFields struct {
	ID         string
	RemoteIP   string
	Host       string
	Method     string
	Path       string
	Protocol   string
	StatusCode int
	Latency    int64
	Error      error
}

func (lf *logFields) MarshalZerologObject(e *zerolog.Event) {
	e.
		Str("id", lf.ID).
		Str("remote_ip", lf.RemoteIP).
		Str("host", lf.Host).
		Str("method", lf.Method).
		Str("path", lf.Path).
		Str("protocol", lf.Protocol).
		Int("status_code", lf.StatusCode).
		Int64("latency", lf.Latency).
		Str("tag", "request")

	if lf.Error != nil {
		e.Err(lf.Error)
	}
}

func Logger(c *fiber.Ctx) error {
	start := time.Now()
	rid := c.GetRespHeader(fiber.HeaderXRequestID)

	fields := &logFields{
		ID:       rid,
		RemoteIP: c.IP(),
		Method:   c.Method(),
		Host:     c.Hostname(),
		Path:     c.Path(),
		Protocol: c.Protocol(),
	}

	defer func() {
		rvr := recover()

		if rvr != nil {
			err, ok := rvr.(error)
			if !ok {
				err = fmt.Errorf("%v", rvr)
			}

			fields.Error = err

			c.Status(http.StatusInternalServerError)
			c.JSON(model.Error{
				Error:            http.StatusText(fiber.StatusInternalServerError),
				ErrorCode:        fiber.StatusInternalServerError,
				ErrorDescription: err.Error(),
			})
		}

		fields.StatusCode = c.Response().StatusCode()
		fields.Latency = time.Since(start).Milliseconds()

		switch {
		case rvr != nil:
			log.WithLevel(zerolog.PanicLevel).EmbedObject(fields).Msg("panic recover")
		case fields.StatusCode >= 500:
			log.Error().EmbedObject(fields).Msg("server error")
		case fields.StatusCode >= 400:
			log.Error().EmbedObject(fields).Msg("client error")
		case fields.StatusCode >= 300:
			log.Warn().EmbedObject(fields).Msg("redirect")
		case fields.StatusCode >= 200:
			log.Info().EmbedObject(fields).Msg("success")
		case fields.StatusCode >= 100:
			log.Info().EmbedObject(fields).Msg("informative")
		default:
			log.Warn().EmbedObject(fields).Msg("unknown status")
		}
	}()

	return c.Next()
}
