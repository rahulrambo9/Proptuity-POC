package controllers

import (
	"encoding/json"
	"go-user-auth/config"
	model "go-user-auth/models"
	"go-user-auth/responses"
	"go-user-auth/services"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ApplicationsController interface {
	RegisterApplication(c *fiber.Ctx) error
	GetApplicationByClientID(c *fiber.Ctx) error
	GetApplicationByTag(c *fiber.Ctx) error
	GetAllApplications(c *fiber.Ctx) error
	DeleteApplication(c *fiber.Ctx) error
}

type applicationController struct {
	config config.AccountsConfig
	appLib services.ApplicationService
}

func NewApplicationsHandler(config config.AccountsConfig, appLib services.ApplicationService) ApplicationsController {
	return &applicationController{
		config: config,
		appLib: appLib,
	}
}

func (app *applicationController) RegisterApplication(c *fiber.Ctx) error {
	log.Println("Registered application in controller")
	// Unmarshal the request body into the User struct
	var applicationObj model.Application
	json.Unmarshal(c.Body(), &applicationObj)

	response, err := app.appLib.RegisterApplication(c.UserContext(), &applicationObj)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ApplicationResponse{
			Status:  http.StatusNotFound,
			Message: model.Error,
			Error:   err,
		})
	}
	// Respond with the updated list of users
	return c.Status(http.StatusOK).JSON(responses.ApplicationResponse{
		Status:  http.StatusOK,
		Message: model.Success,
		Data:    response,
	})
}

func (app *applicationController) GetApplicationByClientID(c *fiber.Ctx) error {
	clientID := c.Params("client_id")
	response, err := app.appLib.GetApplicationByClientID(c.UserContext(), clientID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ApplicationResponse{
			Status:  http.StatusNotFound,
			Message: model.Error,
			Error:   err,
		})
	}
	// Respond with the updated list of users
	return c.Status(http.StatusOK).JSON(responses.ApplicationResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

func (app *applicationController) GetApplicationByTag(c *fiber.Ctx) error {
	tag := c.Params("tag")
	response, err := app.appLib.GetApplicationByTag(c.UserContext(), tag)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ApplicationResponse{
			Status:  http.StatusNotFound,
			Message: model.Error,
			Error:   err,
		})
	}
	// Respond with the updated list of users
	return c.Status(http.StatusOK).JSON(responses.ApplicationResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

func (app *applicationController) GetAllApplications(c *fiber.Ctx) error {
	response, err := app.appLib.GetAllApplications(c.UserContext())
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ApplicationResponse{
			Status:  http.StatusNotFound,
			Message: model.Error,
			Error:   err,
		})
	}
	// Respond with the updated list of users
	return c.Status(http.StatusOK).JSON(responses.ApplicationsResponse{
		Status: http.StatusOK,
		Data:   response,
	})
}

func (app *applicationController) DeleteApplication(c *fiber.Ctx) error {
	clientID := c.Params("client_id")
	response, err := app.appLib.DeleteApplication(c.UserContext(), clientID)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(responses.ApplicationResponse{
			Status:  http.StatusNotFound,
			Message: model.Error,
			Error:   err,
		})
	}
	if !response {
		return c.Status(http.StatusNotFound).JSON(responses.ApplicationResponse{
			Status:  http.StatusNotFound,
			Message: model.Error,
			Error:   err,
		})
	}
	// Respond with the updated list of users
	return c.Status(http.StatusOK).JSON(responses.ApplicationResponse{
		Status: http.StatusOK,
		Data:   nil,
	})
}
