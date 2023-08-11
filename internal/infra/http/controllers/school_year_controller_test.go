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

func TestShouldReturnErrorIfNotPassAllNecessaryParamsInCreateSchoolYear(t *testing.T) {
	actionSchoolYear := new(mocks.SchoolYearActionsMock)
	schoolYearController := NewSchoolYearController(actionSchoolYear)

	scenaries := []ErrorScenarios{
		{
			Endpoint:             "/room",
			Method:               "POST",
			Description:          "when not send data request",
			Data:                 ``,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Data provided is invalid",
		},
		{
			Endpoint:             "/room",
			Method:               "POST",
			Description:          "when school year is not provided",
			Data:                 `{"year": "", "start_at":"2020-01-01", "end_at" : "2020-12-20"}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/room",
			Method:               "POST",
			Description:          "when started_at is not provided",
			Data:                 `{"year": "2020", "start_at":"", "end_at" : "2020-12-20"}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/room",
			Method:               "POST",
			Description:          "when end_at is not provided",
			Data:                 `{"year": "2020", "start_at":"2020-01-01", "end_at" : ""}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
	}

	app := fiber.New()

	for _, scenario := range scenaries {
		t.Run(scenario.Description, func(t *testing.T) {
			app.Post(scenario.Endpoint, schoolYearController.Create)
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

func TestShouldCreateSchoolYearWithSuccess(t *testing.T) {
	actionSchoolYear := new(mocks.SchoolYearActionsMock)
	actionSchoolYear.On("Create", mock.AnythingOfType("dto.SchoolYearRequestDto")).Return(nil)
	schoolYearController := NewSchoolYearController(actionSchoolYear)

	app := fiber.New()
	app.Post("/school-year", schoolYearController.Create)
	data := `{"year": "2020","start_at":"2020-01-01","end_at":"2020-12-20"}`
	request := httptest.NewRequest("POST", "/school-year", bytes.NewReader([]byte(data)))
	request.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(request)
	var m map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&m)
	response.Body.Close()
	assert.Equal(t, 201, response.StatusCode)
}

func TestShouldReturnErrorIfNotPassAllNecessaryDataInUpdateSchoolYear(t *testing.T) {
	actionSchoolYear := new(mocks.SchoolYearActionsMock)
	schoolYearController := NewSchoolYearController(actionSchoolYear)

	scenaries := []ErrorScenarios{
		{
			Endpoint:             "/school-year/:id",
			Method:               "PUT",
			Param:                "/school-year/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when not send data request",
			Data:                 ``,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Data provided is invalid",
		},
		{
			Endpoint:             "/school-year/1da90050-e182-4551-923d-2c60f72b545a",
			Method:               "PUT",
			Param:                "/school-year/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when year is not provided",
			Data:                 `{"year": "", "start_at":"2020-01-01", "end_at" : "2020-12-20"}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/school-year/1da90050-e182-4551-923d-2c60f72b545a",
			Method:               "PUT",
			Param:                "/school-year/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when started_at is not provided",
			Data:                 `{"year": "2020", "start_at":"", "end_at" : "2020-12-20"}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/school-year/1da90050-e182-4551-923d-2c60f72b545a",
			Method:               "PUT",
			Param:                "/school-year/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when end_at is not provided",
			Data:                 `{"year": "2020", "start_at":"2020-01-01", "end_at" : ""}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
	}

	app := fiber.New()

	for _, scenario := range scenaries {
		t.Run(scenario.Description, func(t *testing.T) {
			app.Put(scenario.Endpoint, schoolYearController.Update)
			request := httptest.NewRequest(scenario.Method, scenario.Param, bytes.NewReader([]byte(scenario.Data)))
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

func TestShouldUpdateSchoolYearWithSuccess(t *testing.T) {
	actionSchoolYear := new(mocks.SchoolYearActionsMock)
	actionSchoolYear.On("Update", "1da90050-e182-4551-923d-2c60f72b545a", mock.AnythingOfType("dto.SchoolYearRequestDto")).Return(nil)
	schoolYearroomController := NewSchoolYearController(actionSchoolYear)
	data := `{"year": "2020", "start_at":"2020-01-01", "end_at" : "2020-12-20"}`
	app := fiber.New()
	app.Put("/school-year/:id", schoolYearroomController.Update)
	request := httptest.NewRequest("PUT", "/school-year/1da90050-e182-4551-923d-2c60f72b545a", bytes.NewReader([]byte(data)))
	request.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(request)
	var m map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&m)
	response.Body.Close()
	fmt.Println(m)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "school year updated with success", m["message"])
}
