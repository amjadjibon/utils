package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
)

type labelValues []string

// With validates the input, and returns a new aggregate labelValues.
func (lvs labelValues) With(labelValues ...string) labelValues {
	if len(labelValues)%2 != 0 {
		labelValues = append(labelValues, "unknown")
	}
	return append(lvs, labelValues...)
}

func makeLabels(labelValues ...string) prometheus.Labels {
	var labels = prometheus.Labels{}
	for i := 0; i < len(labelValues); i += 2 {
		labels[labelValues[i]] = labelValues[i+1]
	}
	return labels
}
