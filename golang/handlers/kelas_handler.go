package handlers

import (
	"strconv"

	"timetabling/models"
	"timetabling/services"

	"github.com/gofiber/fiber/v2"
)

type KelasHandler struct {
	service services.KelasService
}

func NewKelasHandler(service services.KelasService) *KelasHandler {
	return &KelasHandler{service: service}
}

// CreateKelas creates a new kelas
// @Summary Create kelas
// @Description Create a new kelas
// @Tags Kelas
// @Accept json
// @Produce json
// @Param kelas body models.Kelas true "Kelas data"
// @Success 201 {object} models.Kelas
// @Failure 400 {object} map[string]interface{}
// @Router /kelas [post]
func (h *KelasHandler) Create(c *fiber.Ctx) error {
	var payload models.Kelas
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	result, err := h.service.Create(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

// GetAllKelas retrieves all kelas
// @Summary Get all kelas
// @Description Get a list of all kelas
// @Tags Kelas
// @Produce json
// @Success 200 {array} models.Kelas
// @Router /kelas [get]
func (h *KelasHandler) GetAll(c *fiber.Ctx) error {
	items, err := h.service.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(items)
}

// GetByIDKelas retrieves a kelas by ID
// @Summary Get kelas by ID
// @Description Get a specific kelas by ID
// @Tags Kelas
// @Produce json
// @Param id path int true "Kelas ID"
// @Success 200 {object} models.Kelas
// @Failure 404 {object} map[string]interface{}
// @Router /kelas/{id} [get]
func (h *KelasHandler) GetByID(c *fiber.Ctx) error {
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

// UpdateKelas updates a kelas
// @Summary Update kelas
// @Description Update an existing kelas
// @Tags Kelas
// @Accept json
// @Produce json
// @Param id path int true "Kelas ID"
// @Param kelas body models.Kelas true "Updated kelas data"
// @Success 200 {object} models.Kelas
// @Failure 400 {object} map[string]interface{}
// @Router /kelas/{id} [put]
func (h *KelasHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	var payload models.Kelas
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	payload.KelasID = id
	result, err := h.service.Update(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(result)
}

// DeleteKelas deletes a kelas
// @Summary Delete kelas
// @Description Delete a kelas by ID
// @Tags Kelas
// @Param id path int true "Kelas ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Router /kelas/{id} [delete]
func (h *KelasHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	if err := h.service.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
