package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/amjadjibon/utils/metrics"
)

// Histogram implements Histogram via a Prometheus HistogramVec. The difference
// between a Histogram and a Summary is that Histograms require predefined
// quantile buckets, and can be statistically aggregated.
type Histogram struct {
	hv  *prometheus.HistogramVec
	lvs labelValues
}

// With implements Histogram.
func (h *Histogram) With(labelValues ...string) metrics.Histogram {
	return &Histogram{
		hv:  h.hv,
		lvs: h.lvs.With(labelValues...),
	}
}

// Observe implements Histogram.
func (h *Histogram) Observe(value float64) {
	h.hv.With(makeLabels(h.lvs...)).Observe(value)
}

// NewHistogram wraps the HistogramVec and returns a usable Histogram object.
func NewHistogram(hv *prometheus.HistogramVec) *Histogram {
	return &Histogram{
		hv: hv,
	}
}

// NewHistogramFrom constructs and registers a Prometheus HistogramVec,
// and returns a usable Histogram object.
func NewHistogramFrom(opts prometheus.HistogramOpts, labelNames []string) *Histogram {
	var hv = prometheus.NewHistogramVec(opts, labelNames)
	prometheus.MustRegister(hv)
	return NewHistogram(hv)
}
