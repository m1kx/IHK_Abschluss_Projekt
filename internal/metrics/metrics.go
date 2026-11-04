package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	MessagesProcessed = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "messages_processed_total",
		Help: "Number of processed messages.",
	}, []string{"team"})

	ProcessingDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "message_processing_duration_seconds",
		Help:    "Time spent processing a message.",
		Buckets: prometheus.DefBuckets,
	}, []string{"team"})

	Cost = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "processing_cost_total",
		Help: "Accumulated processing cost.",
	}, []string{"team"})
)

func InitTeam(team string) {
	MessagesProcessed.WithLabelValues(team)
	ProcessingDuration.WithLabelValues(team)
	Cost.WithLabelValues(team)
}
