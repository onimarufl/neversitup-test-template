package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/onimarufl/neversitup-test-template/models"
	"github.com/onimarufl/neversitup-test-template/service"
	"github.com/onimarufl/neversitup-test-template/validate"
)

type Handler struct {
	service service.Servicer
}

func NewHandler(service service.Servicer) Handler {
	return Handler{service: service}
}

func (h Handler) HandlerGetUser(c fiber.Ctx) error {
	req := models.Request{}

	err := c.Bind().Body(&req)
	if err != nil {
		return fiber.NewError(http.StatusBadRequest, err.Error())
	}

	validate := validate.Validate(req)
	if validate != nil {
		byte, err := json.Marshal(validate)
		if err != nil {
			return fiber.NewError(http.StatusInternalServerError, err.Error())
		}
		return fiber.NewError(http.StatusBadRequest, string(byte))
	}

	resp, err := h.service.GetUserService(req)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(resp)
}
