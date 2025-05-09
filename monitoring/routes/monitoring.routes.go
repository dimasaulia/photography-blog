package articles_routes

import (
	cont "blog/monitoring/controllers"

	"github.com/gofiber/fiber/v2"
)

func SetupMonitoringRoutes(r fiber.Router, cont cont.MonitoringControllers) {
	r.Get("/stat", cont.GetServerStat)
	r.Get("/metrics", cont.GetServerMetrics)
}
