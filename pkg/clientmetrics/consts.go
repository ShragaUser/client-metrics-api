package clientmetrics

import (
	"strings"
)

const (
	allowedMetricTypes = "counter|gauge|summary|histogram"
)

func isAllowedMetricType(metricType string) bool {
	if metricType == "" {
		return false
	}

	return strings.Contains(allowedMetricTypes, metricType)
}
