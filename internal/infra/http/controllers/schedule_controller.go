package controllers

import (
	"github.com/gofiber/fiber/v2"
	requestvalidator "github.com/henriquerocha2004/sistema-escolar/internal/infra/http/request_validator"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/parsers"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary"
)

type ScheduleController struct {
	scheduleActions secretary.ScheduleActionsInterface
}

func NewScheduleController(scheduleActions secretary.ScheduleActionsInterface) *ScheduleController {
	return &ScheduleController{
		scheduleActions: scheduleActions,
	}
}

func (s *ScheduleController) Create(ctx *fiber.Ctx) error {
	inputRequest := dto.ScheduleRequestDto{}
	err := ctx.BodyParser(&inputRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Data provided is invalid",
			nil,
		))
	}

	validateMessages := requestvalidator.ValidateRequest(&inputRequest)
	if validateMessages != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Failed to validate data",
			nil,
		))
	}

	err = s.scheduleActions.Create(inputRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusCreated).JSON(NewResponseDto(
		"success",
		"schedule created with success",
		nil,
	))
}

func (s *ScheduleController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"schedule id is not provided",
			nil,
		))
	}

	inputRequest := dto.ScheduleRequestDto{}
	err := ctx.BodyParser(&inputRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Data provided is invalid",
			nil,
		))
	}

	validationMessages := requestvalidator.ValidateRequest(&inputRequest)
	if validationMessages != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Failed to validate data",
			validationMessages,
		))
	}

	err = s.scheduleActions.Update(id, inputRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"schedule updated with success",
		nil,
	))
}

func (s *ScheduleController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"schedule id is not provided",
			nil,
		))
	}

	err := s.scheduleActions.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"schedule deleted with success",
		nil,
	))
}

func (s *ScheduleController) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"schedule id is not provided",
			nil,
		))
	}

	schedule, err := s.scheduleActions.FindOne(id)
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
		schedule,
	))
}

func (s *ScheduleController) FindAll(ctx *fiber.Ctx) error {
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

	schedules, err := s.scheduleActions.FindAll(*paginatorRequestDto)
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
		schedules,
	))
}
