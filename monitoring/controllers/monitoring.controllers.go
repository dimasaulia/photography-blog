package controllers

import "github.com/gofiber/fiber/v2"

type MonitoringControllers interface {
	GetServerStat(c *fiber.Ctx) error
	GetServerMetrics(c *fiber.Ctx) error
}
