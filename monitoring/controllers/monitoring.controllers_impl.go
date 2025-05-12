package controllers

import (
	h "blog/monitoring/helpers"
	"blog/provider/config"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type MonitoringControllersImpl struct {
}

func NewMonitoringController() MonitoringControllers {
	return &MonitoringControllersImpl{}
}

func (m *MonitoringControllersImpl) GetServerStat(c *fiber.Ctx) error {
	serverStat := h.GetServerStatistic()
	return c.JSON(serverStat)
}

func (m *MonitoringControllersImpl) GetServerMetrics(c *fiber.Ctx) error {
	env := config.ReadConfigigurationFile()
	appEnv := env.GetString("ENV")

	serverStat := h.GetServerStatistic()
	serverMetrics := fmt.Sprintf(`# HELP program_memory Show how much program use system memory
    # TYPE program_memory gauge
    resident_set_size_bytes{program="be-blog-life", env="%v"} %v
    heap_limit_bytes{program="be-blog-life", env="%v"} %v
    heap_used_bytes{program="be-blog-life", env="%v"} %v
    heap_total_bytes{program="be-blog-life", env="%v"} %v
    external_bytes{program="be-blog-life", env="%v"} %v
    array_buffers_bytes{program="be-blog-life", env="%v"} %v
	`,
		appEnv, serverStat.RSSBytes,
		appEnv, serverStat.HeapSysBytes,
		appEnv, serverStat.HeapUsedBytes,
		appEnv, serverStat.HeapTotalBytes,
		appEnv, 0,
		appEnv, 0,
	)

	// Trim spaces from each line
	// Split the metrics string into lines, trim each line, and join them back
	lines := strings.Split(serverMetrics, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}
	trimmedMetrics := strings.Join(lines, "\n")

	return c.Send([]byte(trimmedMetrics))
}
