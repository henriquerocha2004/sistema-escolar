package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestShouldReturnErrorIfNotPassAllNecessaryParamsInCreateService(t *testing.T) {
	actionsService := new(mocks.ServiceActionsMock)
	serviceController := NewServiceController(actionsService)

	scenaries := []ErrorScenarios{
		{
			Endpoint:             "/service",
			Method:               "POST",
			Description:          "when not send data request",
			Data:                 ``,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Data provided is invalid",
		},
		{
			Endpoint:             "/service",
			Method:               "POST",
			Description:          "when description is not provided",
			Data:                 `{"description": ""}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "failed to validate data",
		},
	}

	app := fiber.New()

	for _, scenario := range scenaries {
		t.Run(scenario.Description, func(t *testing.T) {
			app.Post(scenario.Endpoint, serviceController.Create)
			request := httptest.NewRequest(scenario.Method, scenario.Endpoint, bytes.NewReader([]byte(scenario.Data)))
			request.Header.Set("Content-Type", "application/json")
			response, _ := app.Test(request)
			var m map[string]interface{}
			_ = json.NewDecoder(response.Body).Decode(&m)
			_ = response.Body.Close()
			assert.Equal(t, scenario.ExpectedCodeResponse, response.StatusCode)
			assert.Equal(t, scenario.ExpectedErrorMessage, m["message"])
		})
	}
}

func TestShouldCreateServiceWithSuccess(t *testing.T) {

	actionsService := new(mocks.ServiceActionsMock)
	actionsService.On("Create", mock.AnythingOfType("dto.ServiceRequestDto")).Return(nil)
	serviceController := NewServiceController(actionsService)

	app := fiber.New()
	app.Post("/service", serviceController.Create)
	data := `{"description": "any description"}`
	request := httptest.NewRequest("POST", "/service", bytes.NewReader([]byte(data)))
	request.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(request)
	var m map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&m)
	_ = response.Body.Close()
	assert.Equal(t, 201, response.StatusCode)
}

func TestShouldReturnErrorIfNotPassAllNecessaryDataInUpdateService(t *testing.T) {
	actionsService := new(mocks.ServiceActionsMock)
	serviceController := NewServiceController(actionsService)

	scenaries := []ErrorScenarios{
		{
			Endpoint:             "/service/:id",
			Method:               "PUT",
			Param:                "/service/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when not send data request",
			Data:                 ``,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Data provided is invalid",
		},
		{
			Endpoint:             "/service/1da90050-e182-4551-923d-2c60f72b545a",
			Method:               "PUT",
			Param:                "/service/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when description is not provided",
			Data:                 `{"description": ""}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
	}

	app := fiber.New()

	for _, scenario := range scenaries {
		t.Run(scenario.Description, func(t *testing.T) {
			app.Put(scenario.Endpoint, serviceController.Update)
			request := httptest.NewRequest(scenario.Method, scenario.Param, bytes.NewReader([]byte(scenario.Data)))
			request.Header.Set("Content-Type", "application/json")
			response, _ := app.Test(request)
			var m map[string]interface{}
			_ = json.NewDecoder(response.Body).Decode(&m)
			_ = response.Body.Close()
			assert.Equal(t, scenario.ExpectedCodeResponse, response.StatusCode)
			assert.Equal(t, scenario.ExpectedErrorMessage, m["message"])
		})
	}
}

func TestShouldUpdateServiceWithSuccess(t *testing.T) {
	actionsService := new(mocks.ServiceActionsMock)
	actionsService.On("Update", "1da90050-e182-4551-923d-2c60f72b545a", mock.AnythingOfType("dto.ServiceRequestDto")).Return(nil)
	serviceController := NewServiceController(actionsService)
	data := `{"description": "2020"}`
	app := fiber.New()
	app.Put("/service/:id", serviceController.Update)
	request := httptest.NewRequest("PUT", "/service/1da90050-e182-4551-923d-2c60f72b545a", bytes.NewReader([]byte(data)))
	request.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(request)
	var m map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&m)
	_ = response.Body.Close()
	fmt.Println(m)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "service updated with success", m["message"])
}
