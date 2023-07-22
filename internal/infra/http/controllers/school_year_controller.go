package controllers

import (
	"github.com/gofiber/fiber/v2"
	requestvalidator "github.com/henriquerocha2004/sistema-escolar/internal/infra/http/request_validator"
	"github.com/henriquerocha2004/sistema-escolar/internal/infra/parsers"
	"github.com/henriquerocha2004/sistema-escolar/internal/presenters"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary"
	"log"
)

type SchoolYearController struct {
	actions secretary.SchoolYearActionsInterface
}

func NewSchoolYearController(actions secretary.SchoolYearActionsInterface) *SchoolYearController {
	return &SchoolYearController{
		actions: actions,
	}
}

func (s *SchoolYearController) Create(ctx *fiber.Ctx) error {
	requestDto := dto.SchoolYearRequestDto{}
	err := ctx.BodyParser(&requestDto)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Data provided is invalid",
			nil,
		))
	}

	validateMessages := requestvalidator.ValidateRequest(&requestDto)
	if validateMessages != nil {
		log.Println(validateMessages)
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Failed to validate data",
			validateMessages,
		))
	}

	err = s.actions.Create(requestDto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusCreated).JSON(NewResponseDto(
		"success",
		"school year created with success",
		nil,
	))
}

func (s *SchoolYearController) Update(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"school year id is not provided",
			nil,
		))
	}

	var inputDto dto.SchoolYearRequestDto
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

	err = s.actions.Update(id, inputDto)
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
		"school year updated with success",
		nil,
	))
}

func (s *SchoolYearController) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"school year id is not provided",
			nil,
		))
	}
	log.Println(id)
	err := s.actions.Delete(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"error",
		"school year deleted with success",
		nil,
	))
}

func (s *SchoolYearController) Find(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"school year id is not provided",
			nil,
		))
	}

	schoolYear, err := s.actions.FindOne(id)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	presenter := presenters.WebSchoolYearPresenter{}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"",
		presenter.FillSingleSchoolYearPresenter(*schoolYear),
	))
}

func (s *SchoolYearController) FindAll(ctx *fiber.Ctx) error {
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

	schoolYears, err := s.actions.FindAll(*paginatorRequestDto)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(NewResponseDto(
			"error",
			err.Error(),
			nil,
		))
	}

	presenter := presenters.WebSchoolYearPresenter{}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"",
		presenter.FillMultipleSchoolYearPresenter(*schoolYears),
	))
}
