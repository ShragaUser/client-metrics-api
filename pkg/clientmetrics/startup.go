package clientmetrics

import (
	"clientmetrics/pkg/config"
	"log/slog"

	"github.com/ShragaUser/gin-metrics/ginmetrics"
)

func Init() error {
	cfg := config.GetConfig()
	if !cfg.ConfigFileSupported() {
		return nil
	}

	customMetricCfg, err := cfg.GetPreDefinedCustomMetricsConfig()
	if err != nil {
		return err
	}

	for _, inputMetric := range customMetricCfg.Metrics {
		metric := &ginmetrics.Metric{
			Type:        getMetricType(inputMetric.Type),
			Name:        inputMetric.Name,
			Description: inputMetric.Description,
			Labels:      inputMetric.Labels,
			Buckets:     inputMetric.Buckets,
			Objectives:  inputMetric.Objectives,
		}

		slog.Info("creating metric", "metric", metric.Name)

		if _, err := createNewMetricOnce(metric); err != nil {
			return err
		}
	}

	return nil
}
