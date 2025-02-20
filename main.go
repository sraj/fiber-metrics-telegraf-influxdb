package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"time"
)

func isPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

var httpDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "http_response_time_seconds",
	Help:    "Duration of HTTP requests.",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{"method", "path"})

func NewMetrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start).Seconds()
		httpDuration.WithLabelValues(r.Method, r.URL.Path).Observe(duration)
	})
}

func main() {
	app := fiber.New()
	app.Use(adaptor.HTTPMiddleware(NewMetrics))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Hot Reload!")
	})

	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	log.Fatal(app.Listen(":3000"))

	//str := []string{"RACEACAR", "A", "ABCDEFGFEDCBA", "ABC", "ABCBA", "ABBA", "RACEACAR"}
	//for i, s := range str {
	//	fmt.Printf("Checking on Index: %d value: %s\n", i, s)
	//	fmt.Printf("Is it a palindrome?.....%v\n\n", isPalindrome(s))
	//}

}
