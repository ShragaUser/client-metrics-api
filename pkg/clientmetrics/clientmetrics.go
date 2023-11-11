package clientmetrics

import (
	"log/slog"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
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

	if !body.Validate() {
		c.String(400, "invalid metric")
		return
	}

	slog.Info("received metric", "metric", body)

	metric, err := createNewMetricOnce(body)
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

func createNewMetricOnce(body ClientMetricsRequestBody) (metric *ginmetrics.Metric, err error) {
	once, _ := customMetricsMap.LoadOrStore(body.MetricName, &sync.Once{})
	once.(*sync.Once).Do(
		func() {
			metric := &ginmetrics.Metric{
				Type:   body.GetMetricType(),
				Name:   body.MetricName,
				Labels: body.MetricLabels,
			}

			if err = GetMonitor().AddMetric(metric); err != nil {
				slog.Error("could not add metric", "err", err.Error())
				return
			}
		},
	)

	if err != nil {
		return nil, err
	}

	return GetMonitor().GetMetric(body.MetricName), nil
}

func handleMetricByType(metric *ginmetrics.Metric, body ClientMetricsRequestBody) error {
	incrementAmount := 1.0
	if body.MetricValue != nil {
		incrementAmount = *body.MetricValue
	}

	if err := metric.Add(body.MetricLabels, incrementAmount); err != nil {
		return err
	}

	return nil
}
