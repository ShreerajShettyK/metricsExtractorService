package helpers

import (
	"bytes"
	"encoding/json"
	"log"
	"metricsExtractor/configs"
	"metricsExtractor/models"
	"net/http"
	"time"

	"go.uber.org/zap"
)

// SendMetricsToSplunk sends metrics data to Splunk using HTTP Event Collector
func SendMetricsToSplunk(logger *zap.Logger, metrics map[string]interface{}) {
	// Prepare metric event with "_metrics" sourcetype
	metricEvent := models.SplunkMetric{
		Time:       time.Now().Unix(),
		Event:      "system_metrics",
		Host:       configs.Envs.SplunkHost,
		Sourcetype: configs.Envs.SplunkSType,
		Fields:     metrics,
		Index:      configs.Envs.SplunkIndex,
	}

	// Convert to JSON
	jsonData, err := json.Marshal(metricEvent)
	if err != nil {
		logger.Error("Error encoding JSON", zap.Error(err))
		return
	}

	// Send the request
	req, err := http.NewRequest("POST", configs.Envs.SplunkURL, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Error("Error creating HTTP request", zap.Error(err))
		return
	}

	req.Header.Set("Authorization", "Splunk "+configs.Envs.SplunkToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	// Retry logic for sending metrics
	maxRetries := 3
	for attempt := 0; attempt < maxRetries; attempt++ {
		resp, err := client.Do(req)
		if err != nil {
			logger.Error("Error sending metric to Splunk", zap.Error(err))
			time.Sleep(2 * time.Second) // Wait before retrying
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			log.Println("Metrics sent successfully")
			return
		} else {
			logger.Error("Splunk HEC returned non-200 status code", zap.Int("status_code", resp.StatusCode))
			time.Sleep(2 * time.Second) // Wait before retrying
		}
	}

	log.Println("Failed to send metrics after maximum retries")
}
