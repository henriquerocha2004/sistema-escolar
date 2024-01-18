package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http/httptest"
	"testing"
)

func TestShouldReturnErrorIfNotPassAllNecessaryParamsInCreateSchedule(t *testing.T) {
	actionScheduleClass := new(mocks.ScheduleActionsMock)
	scheduleClassController := NewScheduleController(actionScheduleClass)

	scenaries := []ErrorScenarios{
		{
			Endpoint:             "/schedule",
			Method:               "POST",
			Description:          "when not send data request",
			Data:                 ``,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Data provided is invalid",
		},
		{
			Endpoint:             "/schedule",
			Method:               "POST",
			Description:          "when initial_time is not provided",
			Data:                 `{"description": "any description", "initial_time":"", "final_time" : "09:00", "school_year" : "2020"}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/schedule",
			Method:               "POST",
			Description:          "when initial_time is invalid",
			Data:                 `{"description": "any description", "initial_time":"asdsadsa", "final_time" : "09:00", "school_year" : "2020"}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/schedule",
			Method:               "POST",
			Description:          "when final_time is not provided",
			Data:                 `{"description": "any description", "initial_time":"09:00", "final_time" : "", "school_year" : "2020"}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/schedule",
			Method:               "POST",
			Description:          "when final_time is invalid",
			Data:                 `{"description": "any description", "initial_time":"09:00", "final_time" : "asdsadas", "school_year" : "2020"}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/schedule",
			Method:               "POST",
			Description:          "when school_year is not provided",
			Data:                 `{"description": "any description", "initial_time":"09:00", "final_time" : "asdsadas", "school_year" : ""}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
	}

	app := fiber.New()

	for _, scenario := range scenaries {
		t.Run(scenario.Description, func(t *testing.T) {
			app.Post(scenario.Endpoint, scheduleClassController.Create)
			request := httptest.NewRequest(scenario.Method, scenario.Endpoint, bytes.NewReader([]byte(scenario.Data)))
			request.Header.Set("Content-Type", "application/json")
			response, _ := app.Test(request)
			var m map[string]interface{}
			_ = json.NewDecoder(response.Body).Decode(&m)
			response.Body.Close()
			assert.Equal(t, scenario.ExpectedCodeResponse, response.StatusCode)
			assert.Equal(t, scenario.ExpectedErrorMessage, m["message"])
		})
	}
}

func TestShouldCreateScheduleWithSuccess(t *testing.T) {
	actionScheduleClass := new(mocks.ScheduleActionsMock)
	actionScheduleClass.On("Create", mock.AnythingOfType("paginator.ScheduleRequestDto")).Return(nil)
	scheduleClassController := NewScheduleController(actionScheduleClass)

	app := fiber.New()
	app.Post("/schedule", scheduleClassController.Create)
	data := `{"description": "any description", "initial_time":"09:00", "final_time" : "10:00", "school_year" : "2020"}`
	request := httptest.NewRequest("POST", "/schedule", bytes.NewReader([]byte(data)))
	request.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(request)
	var m map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&m)
	_ = response.Body.Close()
	assert.Equal(t, 201, response.StatusCode)
}

func TestShouldReturnErrorIfNotPassAllNecessaryDataInUpdateScheduleClass(t *testing.T) {
	actionScheduleClass := new(mocks.ScheduleActionsMock)
	scheduleClassController := NewScheduleController(actionScheduleClass)

	scenaries := []ErrorScenarios{
		{
			Endpoint:             "/schedule/:id",
			Method:               "PUT",
			Param:                "/schedule/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when not send data request",
			Data:                 ``,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Data provided is invalid",
		},
		{
			Endpoint:             "/schedule/1da90050-e182-4551-923d-2c60f72b545a",
			Method:               "PUT",
			Param:                "/schedule/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when initial_time is not provided",
			Data:                 `{"description": "any description", "initial_time":"", "final_time" : "10:00", "school_year" : "2020"}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/schedule/1da90050-e182-4551-923d-2c60f72b545a",
			Method:               "PUT",
			Param:                "/schedule/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when final_time is not provided",
			Data:                 `{"description": "any description", "initial_time":"09:00", "final_time" : "", "school_year" : "2020"}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/schedule/1da90050-e182-4551-923d-2c60f72b545a",
			Method:               "PUT",
			Param:                "/schedule/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when school_year is not provided",
			Data:                 `{"description": "any description", "initial_time":"09:00", "final_time" : "10:00", "school_year" : ""}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
	}

	app := fiber.New()

	for _, scenario := range scenaries {
		t.Run(scenario.Description, func(t *testing.T) {
			app.Put(scenario.Endpoint, scheduleClassController.Update)
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

func TestShouldUpdateScheduleWithSuccess(t *testing.T) {
	actionScheduleClass := new(mocks.ScheduleActionsMock)
	actionScheduleClass.On("Update", "1da90050-e182-4551-923d-2c60f72b545a", mock.AnythingOfType("paginator.ScheduleRequestDto")).Return(nil)
	scheduleClassController := NewScheduleController(actionScheduleClass)
	data := `{"description": "any description", "initial_time":"09:00", "final_time" : "10:00", "school_year" : "2002"}`
	app := fiber.New()
	app.Put("/schedule/:id", scheduleClassController.Update)
	request := httptest.NewRequest("PUT", "/schedule/1da90050-e182-4551-923d-2c60f72b545a", bytes.NewReader([]byte(data)))
	request.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(request)
	var m map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&m)
	_ = response.Body.Close()
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "schedule updated with success", m["message"])
}
