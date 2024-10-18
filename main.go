package main

import (
	"log"
	"metricsExtractor/helpers"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"go.uber.org/zap"
)

var logger *zap.Logger

func init() {
	// Initialize zap logger
	var err error
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
}

func collectAndSendMetrics(logger *zap.Logger) {
	metrics := make(map[string]interface{})

	// Collect CPU metrics
	cpuUsage, _ := cpu.Percent(time.Second, false)
	metrics["metric_name:cpu.percent"] = cpuUsage[0]

	// Collect memory metrics
	vmStat, _ := mem.VirtualMemory()
	metrics["metric_name:memory.percent"] = vmStat.UsedPercent

	// Collect network metrics
	netStat, _ := net.IOCounters(false)
	if len(netStat) > 0 {
		metrics["metric_name:network.bytes_sent"] = netStat[0].BytesSent
		metrics["metric_name:network.bytes_recv"] = netStat[0].BytesRecv
	}

	// Send metrics to Splunk
	helpers.SendMetricsToSplunk(logger, metrics)

	time.Sleep(9 * time.Second) // Adjust the sleep duration as needed
}
func main() {
	log.Println("Starting metrics collection")

	// Send metrics at intervals
	for {
		collectAndSendMetrics(logger)
		time.Sleep(20 * time.Second)
	}

}
