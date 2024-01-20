package controller

import (
	"bufio"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/nxadm/tail"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
)

func LogStream(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/event-stream")
	c.Set("Cache-Control", "no-cache")
	c.Set("Connection", "keep-alive")
	c.Set("Transfer-Encoding", "chunked")

	path := viper.GetString("app.logfile")
	if path == "" {
		path = "logs/dnic.log"
	}

	t, err := tail.TailFile(path, tail.Config{Follow: true, ReOpen: true})
	if err != nil {
		return err
	}

	c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
		for line := range t.Lines {
			fmt.Fprintf(w, "data: %s\n\n", line.Text)

			err := w.Flush()
			if err != nil {
				// Refreshing page in web browser will establish a new
				// SSE connection, but only (the last) one is alive, so
				// dead connections must be closed here.
				log.Error().Err(err).Msg("error while flushing.")
				break
			}
		}
	}))

	return nil
}
