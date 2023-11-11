package clientmetrics

import (
	"strings"
)

const (
	allowedMetricTypes = "counter|gauge|" // add histogram and summary in the future if needed
)

func isAllowedMetricType(metricType string) bool {
	if metricType == "" {
		return false
	}

	return strings.Contains(allowedMetricTypes, metricType)
}
