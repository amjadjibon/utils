package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/amjadjibon/utils/metrics"
)

// Counter implements Counter, via a Prometheus CounterVec.
type Counter struct {
	cv  *prometheus.CounterVec
	lvs labelValues
}

// With implements Counter.
func (c *Counter) With(labelValues ...string) metrics.Counter {
	return &Counter{
		cv:  c.cv,
		lvs: c.lvs.With(labelValues...),
	}
}

// Add implements Counter.
func (c *Counter) Add(delta float64) {
	c.cv.With(makeLabels(c.lvs...)).Add(delta)
}

// Inc increments the counter by 1.
func (c *Counter) Inc() {
	c.cv.With(makeLabels(c.lvs...)).Inc()
}

// NewCounter wraps the CounterVec and returns a usable Counter object.
func NewCounter(cv *prometheus.CounterVec) *Counter {
	return &Counter{
		cv: cv,
	}
}

// NewCounterFrom constructs and registers a Prometheus CounterVec,
// and returns a usable Counter object.
func NewCounterFrom(opts prometheus.CounterOpts, labelNames []string) *Counter {
	var cv = prometheus.NewCounterVec(opts, labelNames)
	prometheus.MustRegister(cv)
	return NewCounter(cv)
}
