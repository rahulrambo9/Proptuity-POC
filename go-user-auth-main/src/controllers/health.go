package controllers

import (
	"fmt"
	"go-user-auth/config"
	"go-user-auth/services"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type HealthController interface {
	Apphealth(c *fiber.Ctx) error
	Home(c *fiber.Ctx) error
}

type healthController struct {
	config    config.AccountsConfig
	healthSvc services.HealthService
}

func NewHealthController(config config.AccountsConfig, healthSvc services.HealthService) HealthController {
	return &healthController{
		config:    config,
		healthSvc: healthSvc,
	}
}

func (healthCtrl *healthController) Home(c *fiber.Ctx) error {
	version := fmt.Sprintf("API: %s, Version: %s", healthCtrl.config.AppCode, healthCtrl.config.AppVersion)
	log.Debug("app version: ", version)
	return c.JSON(version)
}

func (healthCtrl *healthController) Apphealth(c *fiber.Ctx) error {
	response, err := healthCtrl.healthSvc.GetHealth(c.Context())
	if err != nil {
		return err
	}
	return c.SendString(response)
}
