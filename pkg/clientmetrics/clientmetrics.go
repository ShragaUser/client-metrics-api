package clientmetrics

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/ShragaUser/gin-metrics/ginmetrics"
	"github.com/gin-gonic/gin"
)

var singletonOnce sync.Once
var metrics *ginmetrics.Monitor

func GetMonitor() *ginmetrics.Monitor {
	singletonOnce.Do(
		func() {
			metrics = ginmetrics.GetMonitor()
		})

	return metrics
}

var customMetricsMap = sync.Map{}

func PostMetricHandler(c *gin.Context) {
	body := ClientMetricsRequestBody{}
	if err := c.BindJSON(&body); err != nil {
		slog.Error("invalid request body", "err", err.Error())
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if err := body.Validate(); err != nil {
		c.String(400, err.Error())
		return
	}

	slog.Info("received metric", "metric", body)

	metricDef := getMetricFromRequestBody(body)
	metric, err := createNewMetricOnce(metricDef)
	if err != nil {
		slog.Error("could not create metric", "err", err.Error())
		c.String(500, "could not create metric")
		return
	}

	if err = handleMetricByType(metric, body); err != nil {
		slog.Error("could not handle metric", "err", err.Error())
		c.String(500, "could not handle metric")
		return
	}
}

func getMetricFromRequestBody(body ClientMetricsRequestBody) *ginmetrics.Metric {
	return &ginmetrics.Metric{
		Type:   body.GetMetricType(),
		Name:   body.MetricName,
		Labels: body.MetricLabels,
	}
}

func createNewMetricOnce(metricDef *ginmetrics.Metric) (newMetric *ginmetrics.Metric, err error) {
	once, _ := customMetricsMap.LoadOrStore(metricDef.Name, &sync.Once{})
	once.(*sync.Once).Do(
		func() {
			if err = GetMonitor().AddMetric(metricDef); err != nil {
				slog.Error("could not add metric", "err", err.Error())
				return
			}
		},
	)

	if err != nil {
		customMetricsMap.Delete(metricDef.Name)
		return nil, err
	}

	return GetMonitor().GetMetric(metricDef.Name), nil
}

func handleCounterAndGauge(metric *ginmetrics.Metric, body ClientMetricsRequestBody) error {
	incrementAmount := 1.0
	if body.MetricValue != nil {
		incrementAmount = *body.MetricValue
	}

	if err := metric.Add(body.MetricLabels, incrementAmount); err != nil {
		return err
	}

	return nil
}

func handleSummaryAndHistogram(metric *ginmetrics.Metric, body ClientMetricsRequestBody) error {
	if body.MetricValue == nil {
		return nil
	}

	if err := metric.Observe(body.MetricLabels, *body.MetricValue); err != nil {
		return err
	}

	return nil
}

func handleMetricByType(metric *ginmetrics.Metric, body ClientMetricsRequestBody) error {
	switch metric.Type {
	case ginmetrics.Counter:
		return handleCounterAndGauge(metric, body)
	case ginmetrics.Gauge:
		return handleCounterAndGauge(metric, body)
	case ginmetrics.Summary:
		return handleSummaryAndHistogram(metric, body)
	case ginmetrics.Histogram:
		return handleSummaryAndHistogram(metric, body)
	default:
		return nil
	}
}
