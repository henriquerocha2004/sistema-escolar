package controllers

import (
	"github.com/gofiber/fiber/v2"
	requestvalidator "github.com/henriquerocha2004/sistema-escolar/internal/infra/http/request_validator"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/parsers"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/room"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/room/roomService"
	dto2 "github.com/henriquerocha2004/sistema-escolar/internal/school/secretary/schedule"
	"log"
)

type RoomController struct {
	roomActions roomService.RoomActionsInterface
}

func NewRoomController(roomActions roomService.RoomActionsInterface) *RoomController {
	return &RoomController{
		roomActions: roomActions,
	}
}

func (r *RoomController) Create(ctx *fiber.Ctx) error {
	var requestDto room.RoomRequestDto
	err := ctx.BodyParser(&requestDto)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			NewResponseDto(
				"error",
				"Data provided is invalid",
				nil,
			))
	}

	validateMessages := requestvalidator.ValidateRequest(&requestDto)
	if validateMessages != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			NewResponseDto(
				"error",
				"Failed to validate data",
				validateMessages,
			))
	}

	err = r.roomActions.Create(requestDto)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusCreated).JSON(NewResponseDto(
		"success",
		"Room created with success",
		nil,
	))
}

func (r *RoomController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"room id is not provided",
			nil,
		))
	}

	var inputDto room.RoomRequestDto
	err := ctx.BodyParser(&inputDto)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Data provided is invalid",
			nil,
		))
	}

	validationMessages := requestvalidator.ValidateRequest(&inputDto)
	if validationMessages != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Failed to validate data",
			validationMessages,
		))
	}

	err = r.roomActions.Update(id, inputDto)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"room updated with success",
		nil,
	))
}

func (r *RoomController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"room id is not provided",
			nil,
		))
	}

	err := r.roomActions.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"room deleted with success",
		nil,
	))
}

func (r *RoomController) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"room id is not provided",
			nil,
		))
	}

	room, err := r.roomActions.FindById(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"",
		room,
	))
}

func (r *RoomController) FindAll(ctx *fiber.Ctx) error {
	paginatorRequestDto, err := parsers.ParseRequestPaginator(ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	validateMessages := requestvalidator.ValidateRequest(paginatorRequestDto)
	if validateMessages != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"failed to validate data",
			validateMessages,
		))
	}

	rooms, err := r.roomActions.FindAll(*paginatorRequestDto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"",
		rooms,
	))
}

func (r *RoomController) SyncSchedule(ctx *fiber.Ctx) error {
	roomScheduleDto := dto2.RoomScheduleDto{}
	err := ctx.BodyParser(&roomScheduleDto)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(
			NewResponseDto(
				"error",
				"Data provided is invalid",
				nil,
			))
	}

	validateMessages := requestvalidator.ValidateRequest(&roomScheduleDto)
	if validateMessages != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"failed to validate data",
			validateMessages,
		))
	}

	err = r.roomActions.SyncSchedule(roomScheduleDto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"schedule room sync with success",
		nil,
	))
}
