package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var (
	uniformDomain = flag.Float64(
		"uniform.domain", 0.0002, "The domain for the uniform distribution.",
	)
	normDomain    = flag.Float64("normal.domain", 0.0002, "The domain for the normal distribution.")
	normMean      = flag.Float64("normal.mean", 0.00001, "The mean for the normal distribution.")
	histogram     = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "default",
			Subsystem: "go_exercise",
			Name:      "err_rate",
			Help:      "The error code stat",
		},
		[]string{"id", "method"},
	)
	h2 = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: "default",
			Subsystem: "go_exercise2",
			Name:      "err_rate2",
			Help:      "The error code stat",
		},
	)
)

func init() {
	prometheus.MustRegister(histogram)
	prometheus.MustRegister(h2)
}

func main() {
	rand.Seed(int64(time.Now().Nanosecond()))

	for i := 0; i < 100; i++ {
		id := strconv.Itoa(i)
		for j := 0; j < 1000; j++ {
			v := (rand.NormFloat64() * *normDomain) + *normMean
			fmt.Println(v)
			histogram.WithLabelValues(id, "test").Observe(v)
			h2.(prometheus.ExemplarObserver).ObserveWithExemplar(
				v, prometheus.Labels{
					"id": id,
				},
			)
		}
	}

	// Expose the registered metrics via HTTP.
	http.Handle(
		"/metrics", promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{
				// Opt into OpenMetrics to support exemplars.
				EnableOpenMetrics: true,
			},
		),
	)
	log.Fatal(http.ListenAndServe(":9001", nil))
}
