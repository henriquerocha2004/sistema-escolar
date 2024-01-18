package controllers

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ErrorScenarios struct {
	Endpoint             string
	Param                string
	Method               string
	Description          string
	Data                 string
	ExpectedCodeResponse int
	ExpectedErrorMessage string
}

func TestShouldReturnErrorIfNotPassAllNecessaryParamsInCreateRoom(t *testing.T) {
	actionRoom := new(mocks.RoomActionsMock)
	roomController := NewRoomController(actionRoom)

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
			Description:          "when code room is not provided",
			Data:                 `{"code": "", "description":"desc", "capacity" : 20}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/room",
			Method:               "POST",
			Description:          "when description room is not provided",
			Data:                 `{"code": "SL-03", "description":"", "capacity" : 20}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/room",
			Method:               "POST",
			Description:          "when capacity room is not provided",
			Data:                 `{"code": "SL-03", "description":"desc", "capacity" : 0}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
	}

	app := fiber.New()

	for _, scenario := range scenaries {
		t.Run(scenario.Description, func(t *testing.T) {
			app.Post(scenario.Endpoint, roomController.Create)
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

func TestShouldCreateRoomWithSuccess(t *testing.T) {
	actionRoom := new(mocks.RoomActionsMock)
	actionRoom.On("Create", mock.AnythingOfType("paginator.RoomRequestDto")).Return(nil)
	roomController := NewRoomController(actionRoom)

	app := fiber.New()
	app.Post("/room", roomController.Create)
	data := `{"code" : "SL-01", "description" : "desc", "capacity" : 20}`
	request := httptest.NewRequest("POST", "/room", bytes.NewReader([]byte(data)))
	request.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(request)
	var m map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&m)
	response.Body.Close()
	assert.Equal(t, 201, response.StatusCode)
}

func TestShouldReturnErrorIfNotPassAllNecessaryDataInUpdate(t *testing.T) {
	actionRoom := new(mocks.RoomActionsMock)
	roomController := NewRoomController(actionRoom)

	scenaries := []ErrorScenarios{
		{
			Endpoint:             "/room/:id",
			Method:               "PUT",
			Param:                "/room/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when not send data request",
			Data:                 ``,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Data provided is invalid",
		},
		{
			Endpoint:             "/room/1da90050-e182-4551-923d-2c60f72b545a",
			Method:               "PUT",
			Param:                "/room/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when code room is not provided",
			Data:                 `{"code": "", "description":"desc", "capacity" : 20}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/room/1da90050-e182-4551-923d-2c60f72b545a",
			Method:               "PUT",
			Param:                "/room/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when description room is not provided",
			Data:                 `{"code": "SL-03", "description":"", "capacity" : 20}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
		{
			Endpoint:             "/room/1da90050-e182-4551-923d-2c60f72b545a",
			Method:               "PUT",
			Param:                "/room/1da90050-e182-4551-923d-2c60f72b545a",
			Description:          "when capacity room is not provided",
			Data:                 `{"code": "SL-03", "description":"desc", "capacity" : 0}`,
			ExpectedCodeResponse: 400,
			ExpectedErrorMessage: "Failed to validate data",
		},
	}

	app := fiber.New()

	for _, scenario := range scenaries {
		t.Run(scenario.Description, func(t *testing.T) {
			app.Put(scenario.Endpoint, roomController.Update)
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

func TestShouldUpdateRoomWithSuccess(t *testing.T) {
	actionRoom := new(mocks.RoomActionsMock)
	actionRoom.On("Update", "1da90050-e182-4551-923d-2c60f72b545a", mock.AnythingOfType("paginator.RoomRequestDto")).Return(nil)
	roomController := NewRoomController(actionRoom)
	data := `{"code" : "SL-01", "description" : "desc", "capacity" : 20}`
	app := fiber.New()
	app.Put("/room/:id", roomController.Update)
	request := httptest.NewRequest("PUT", "/room/1da90050-e182-4551-923d-2c60f72b545a", bytes.NewReader([]byte(data)))
	request.Header.Set("Content-Type", "application/json")
	response, _ := app.Test(request)
	var m map[string]interface{}
	_ = json.NewDecoder(response.Body).Decode(&m)
	response.Body.Close()
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "room updated with success", m["message"])
}
