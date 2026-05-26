package main

import (
	"io"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog"
)

var log zerolog.Logger

func init() {
	logDir := "/var/log/app"
	os.MkdirAll(logDir, 0755)
	f, _ := os.OpenFile(logDir+"/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	writers := io.MultiWriter(os.Stderr, f)
	log = zerolog.New(writers).With().Timestamp().Logger()
}

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_response_time_seconds",
	Help:    "Duration of HTTP requests.",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{"method", "path"})

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		if c.Path() == "/metrics" {
			return c.Next()
		}
		start := time.Now()
		err := c.Next()
		httpDuration.WithLabelValues(c.Method(), c.Path()).Observe(time.Since(start).Seconds())
		log.Info().
			Str("method", c.Method()).
			Str("path", c.Path()).
			Float64("duration", time.Since(start).Seconds()).
			Msg("request completed")
		return err
	})

	app.Get("/", func(c *fiber.Ctx) error {
		log.Info().Str("path", "/").Msg("handling root request")
		return c.SendString("Hello, Hot Reload!")
	})

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	log.Info().Msg("starting server on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatal().Err(err).Msg("server stopped")
	}
}
