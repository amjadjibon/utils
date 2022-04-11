package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/amjadjibon/utils/metrics"
)

// Gauge implements Gauge, via a Prometheus GaugeVec.
type Gauge struct {
	gv  *prometheus.GaugeVec
	lvs labelValues
}

// With implements Gauge.
func (g *Gauge) With(labelValues ...string) metrics.Gauge {
	return &Gauge{
		gv:  g.gv,
		lvs: g.lvs.With(labelValues...),
	}
}

// Set implements Gauge.
func (g *Gauge) Set(value float64) {
	g.gv.With(makeLabels(g.lvs...)).Set(value)
}

// Add is supported by Prometheus GaugeVecs.
func (g *Gauge) Add(delta float64) {
	g.gv.With(makeLabels(g.lvs...)).Add(delta)
}

func (g *Gauge) Inc() {
	g.gv.With(makeLabels(g.lvs...)).Inc()
}

func (g *Gauge) Dec() {
	g.gv.With(makeLabels(g.lvs...)).Dec()
}

func (g *Gauge) Sub(f float64) {
	g.gv.With(makeLabels(g.lvs...)).Sub(f)
}

func (g *Gauge) SetToCurrentTime() {
	g.gv.With(makeLabels(g.lvs...)).SetToCurrentTime()
}

// NewGauge wraps the GaugeVec and returns a usable Gauge object.
func NewGauge(gv *prometheus.GaugeVec) *Gauge {
	return &Gauge{
		gv: gv,
	}
}

// NewGaugeFrom constructs and registers a Prometheus GaugeVec,
// and returns a usable Gauge object.
func NewGaugeFrom(opts prometheus.GaugeOpts, labelNames []string) *Gauge {
	var gv = prometheus.NewGaugeVec(opts, labelNames)
	prometheus.MustRegister(gv)
	return NewGauge(gv)
}
