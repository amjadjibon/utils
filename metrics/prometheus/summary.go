package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/amjadjibon/utils/metrics"
)

// Summary implements Histogram, via a Prometheus SummaryVec. The difference
// between a Summary and a Histogram is that Summaries don't require predefined
// quantile buckets, but cannot be statistically aggregated.
type Summary struct {
	sv  *prometheus.SummaryVec
	lvs labelValues
}

// With implements Histogram.
func (s *Summary) With(labelValues ...string) metrics.Histogram {
	return &Summary{
		sv:  s.sv,
		lvs: s.lvs.With(labelValues...),
	}
}

// Observe implements Histogram.
func (s *Summary) Observe(value float64) {
	s.sv.With(makeLabels(s.lvs...)).Observe(value)
}

// NewSummary wraps the SummaryVec and returns a usable Summary object.
func NewSummary(sv *prometheus.SummaryVec) *Summary {
	return &Summary{
		sv: sv,
	}
}

// NewSummaryFrom constructs and registers a Prometheus SummaryVec,
// and returns a usable Summary object.
func NewSummaryFrom(opts prometheus.SummaryOpts, labelNames []string) *Summary {
	var sv = prometheus.NewSummaryVec(opts, labelNames)
	prometheus.MustRegister(sv)
	return NewSummary(sv)
}
