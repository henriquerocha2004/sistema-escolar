package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/dto"
	"github.com/henriquerocha2004/sistema-escolar/internal/school/secretary"
)

type RegisterController struct {
	registerActions secretary.RegistrationActionsInterface
}

func NewRegisterController(regActions secretary.RegistrationActionsInterface) *RegisterController {
	return &RegisterController{
		registerActions: regActions,
	}
}

func (r *RegisterController) Create(ctx *fiber.Ctx) error {
	var registerDto dto.RegistrationDto

	err := ctx.BodyParser(&registerDto)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"Data provided is invalid",
			nil,
		))
	}

	registrationResponse, err := r.registerActions.Create(registerDto)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(NewResponseDto(
			"error",
			"error in create registration",
			nil,
		))
	}

	return ctx.Status(fiber.StatusOK).JSON(NewResponseDto(
		"success",
		"registration created with success",
		registrationResponse,
	))
}
