package controllers

import (
	"github.com/gofiber/fiber/v2"
	requestvalidator "github.com/henriquerocha2004/sistema-escolar/internal/infra/http/request_validator"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/parsers"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary"
)

type ClassRoomController struct {
	classRoomActions secretary.ClassRoomActionsInterface
}

func NewClassRoomController(ca secretary.ClassRoomActionsInterface) *ClassRoomController {
	return &ClassRoomController{
		classRoomActions: ca,
	}
}

func (c *ClassRoomController) Create(ctx *fiber.Ctx) error {
	var dtoRequest dto.ClassRoomRequestDto

	err := ctx.BodyParser(&dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"invalid data provided",
			nil,
		))
	}

	validationMessages := requestvalidator.ValidateRequest(&dtoRequest)
	if validationMessages != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Failed to validate data",
			validationMessages,
		))
	}

	err = c.classRoomActions.Create(dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"class room created with success",
		nil,
	))
}

func (c *ClassRoomController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"class room id is not provided",
			nil,
		))
	}

	var dtoRequest dto.ClassRoomRequestDto
	err := ctx.BodyParser(&dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"invalid data provided",
			nil,
		))
	}

	validationMessages := requestvalidator.ValidateRequest(&dtoRequest)
	if validationMessages != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Failed to validate data",
			validationMessages,
		))
	}

	err = c.classRoomActions.Update(id, dtoRequest)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"class room updated with success",
		nil,
	))
}

func (c *ClassRoomController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"class room id is not provided",
			nil,
		))
	}

	err := c.classRoomActions.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"class room deleted with success",
		nil,
	))
}

func (c *ClassRoomController) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"class room id is not provided",
			nil,
		))
	}

	classRoom, err := c.classRoomActions.Find(id)
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
		classRoom,
	))
}

func (c *ClassRoomController) FindAll(ctx *fiber.Ctx) error {
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

	classRooms, err := c.classRoomActions.FindAll(*paginatorRequestDto)
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
		classRooms,
	))
}
