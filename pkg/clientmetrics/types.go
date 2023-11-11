package clientmetrics

import "github.com/penglongli/gin-metrics/ginmetrics"

type ClientMetricsRequestBody struct {
	MetricName   string   `json:"metricName"`
	MetricType   string   `json:"metricType"`
	MetricValue  *float64 `json:"metricValue,omitempty"`
	MetricLabels []string `json:"metricLabels,omitempty"`
}

func (c *ClientMetricsRequestBody) Validate() bool {
	if c.MetricName == "" || c.MetricType == "" {
		return false
	}

	return isAllowedMetricType(c.MetricType)
}

func (c *ClientMetricsRequestBody) GetMetricType() ginmetrics.MetricType {
	switch c.MetricType {
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
