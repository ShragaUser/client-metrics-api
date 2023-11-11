package clientmetrics

import (
	"fmt"

	"github.com/ShragaUser/gin-metrics/ginmetrics"
)

type ClientMetricsRequestBody struct {
	MetricName   string   `json:"metricName"`
	MetricType   string   `json:"metricType"`
	MetricValue  *float64 `json:"metricValue,omitempty"`
	MetricLabels []string `json:"metricLabels,omitempty"`
}

func (c *ClientMetricsRequestBody) Validate() error {
	if c.MetricName == "" || c.MetricType == "" {
		return fmt.Errorf("metric name and type are required")
	}

	if c.GetMetricType() == ginmetrics.Summary || c.GetMetricType() == ginmetrics.Histogram {
		if c.MetricValue == nil {
			return fmt.Errorf("metric value is required for %s metric type", c.MetricType)
		}

		if GetMonitor().GetMetric(c.MetricName).Name != c.MetricName {
			return fmt.Errorf("metric of type %s must pre-exist", c.MetricType)
		}
	}

	if ok := isAllowedMetricType(c.MetricType); !ok {
		return fmt.Errorf("metric type %s is not allowed", c.MetricType)
	}

	return nil
}

func (c *ClientMetricsRequestBody) GetMetricType() ginmetrics.MetricType {
	return getMetricType(c.MetricType)
}

func getMetricType(metricType string) ginmetrics.MetricType {
	switch metricType {
	case "counter":
		return ginmetrics.Counter
	case "gauge":
		return ginmetrics.Gauge
	case "histogram":
		return ginmetrics.Histogram
	case "summary":
		return ginmetrics.Summary
	default:
		return -1
	}
}
