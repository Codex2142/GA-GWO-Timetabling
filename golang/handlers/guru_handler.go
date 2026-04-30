package handlers

import (
	"strconv"

	"timetabling/models"
	"timetabling/services"

	"github.com/gofiber/fiber/v2"
)

type GuruHandler struct {
	service services.GuruService
}

func NewGuruHandler(service services.GuruService) *GuruHandler {
	return &GuruHandler{service: service}
}

// CreateGuru creates a new guru record
// @Summary Create guru
// @Description Create a new guru
// @Tags Guru
// @Accept json
// @Produce json
// @Param guru body models.Guru true "Guru data"
// @Success 201 {object} models.Guru
// @Failure 400 {object} map[string]interface{}
// @Router /guru [post]
func (h *GuruHandler) Create(c *fiber.Ctx) error {
	var payload models.Guru
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	result, err := h.service.Create(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}

// GetAllGuru retrieves all guru records
// @Summary Get all guru
// @Description Get a list of all guru
// @Tags Guru
// @Produce json
// @Success 200 {array} models.Guru
// @Router /guru [get]
func (h *GuruHandler) GetAll(c *fiber.Ctx) error {
	items, err := h.service.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(items)
}

// GetByIDGuru retrieves a guru by ID
// @Summary Get guru by ID
// @Description Get a specific guru by ID
// @Tags Guru
// @Produce json
// @Param id path int true "Guru ID"
// @Success 200 {object} models.Guru
// @Failure 404 {object} map[string]interface{}
// @Router /guru/{id} [get]
func (h *GuruHandler) GetByID(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	item, err := h.service.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return c.JSON(item)
}

// UpdateGuru updates a guru
// @Summary Update guru
// @Description Update an existing guru
// @Tags Guru
// @Accept json
// @Produce json
// @Param id path int true "Guru ID"
// @Param guru body models.Guru true "Updated guru data"
// @Success 200 {object} models.Guru
// @Failure 400 {object} map[string]interface{}
// @Router /guru/{id} [put]
func (h *GuruHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}

	var payload models.Guru
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	payload.GuruID = id
	result, err := h.service.Update(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(result)
}

// DeleteGuru deletes a guru
// @Summary Delete guru
// @Description Delete a guru by ID
// @Tags Guru
// @Param id path int true "Guru ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Router /guru/{id} [delete]
func (h *GuruHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	if err := h.service.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
