package handlers

import (
	"strconv"

	"timetabling/models"
	"timetabling/services"

	"github.com/gofiber/fiber/v2"
)

type WaliKelasHandler struct {
	service services.WaliKelasService
}

func NewWaliKelasHandler(service services.WaliKelasService) *WaliKelasHandler {
	return &WaliKelasHandler{service: service}
}

// CreateWaliKelas creates a new wali kelas
// @Summary Create wali kelas
// @Description Create a new wali kelas
// @Tags WaliKelas
// @Accept json
// @Produce json
// @Param wali body models.WaliKelas true "Wali kelas data"
// @Success 201 {object} models.WaliKelas
// @Failure 400 {object} map[string]interface{}
// @Router /wali-kelas [post]
func (h *WaliKelasHandler) Create(c *fiber.Ctx) error {
	var payload models.WaliKelas
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	result, err := h.service.Create(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

// GetAllWaliKelas retrieves all wali kelas
// @Summary Get all wali kelas
// @Description Get a list of all wali kelas
// @Tags WaliKelas
// @Produce json
// @Success 200 {array} models.WaliKelas
// @Router /wali-kelas [get]
func (h *WaliKelasHandler) GetAll(c *fiber.Ctx) error {
	items, err := h.service.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(items)
}

// GetByIDWaliKelas retrieves a wali kelas by ID
// @Summary Get wali kelas by ID
// @Description Get a specific wali kelas by ID
// @Tags WaliKelas
// @Produce json
// @Param id path int true "Wali ID"
// @Success 200 {object} models.WaliKelas
// @Failure 404 {object} map[string]interface{}
// @Router /wali-kelas/{id} [get]
func (h *WaliKelasHandler) GetByID(c *fiber.Ctx) error {
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

// UpdateWaliKelas updates a wali kelas
// @Summary Update wali kelas
// @Description Update an existing wali kelas
// @Tags WaliKelas
// @Accept json
// @Produce json
// @Param id path int true "Wali ID"
// @Param wali body models.WaliKelas true "Updated wali data"
// @Success 200 {object} models.WaliKelas
// @Failure 400 {object} map[string]interface{}
// @Router /wali-kelas/{id} [put]
func (h *WaliKelasHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	var payload models.WaliKelas
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	payload.ID = id
	result, err := h.service.Update(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(result)
}

// DeleteWaliKelas deletes a wali kelas
// @Summary Delete wali kelas
// @Description Delete a wali kelas by ID
// @Tags WaliKelas
// @Param id path int true "Wali ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Router /wali-kelas/{id} [delete]
func (h *WaliKelasHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	if err := h.service.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
