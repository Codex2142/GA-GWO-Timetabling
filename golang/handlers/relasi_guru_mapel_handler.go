package handlers

import (
	"strconv"

	"timetabling/models"
	"timetabling/services"

	"github.com/gofiber/fiber/v2"
)

type RelasiGuruMapelHandler struct {
	service services.RelasiGuruMapelService
}

func NewRelasiGuruMapelHandler(service services.RelasiGuruMapelService) *RelasiGuruMapelHandler {
	return &RelasiGuruMapelHandler{service: service}
}

// CreateRelasiGuruMapel creates a new relasi guru mapel
// @Summary Create relasi guru mapel
// @Description Create a new relasi guru mapel
// @Tags RelasiGuruMapel
// @Accept json
// @Produce json
// @Param relasi body models.RelasiGuruMapel true "Relasi guru mapel data"
// @Success 201 {object} models.RelasiGuruMapel
// @Failure 400 {object} map[string]interface{}
// @Router /relasi-guru-mapel [post]
func (h *RelasiGuruMapelHandler) Create(c *fiber.Ctx) error {
	var payload models.RelasiGuruMapel
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	result, err := h.service.Create(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

// GetAllRelasiGuruMapel retrieves all relasi guru mapel
// @Summary Get all relasi guru mapel
// @Description Get a list of all relasi guru mapel
// @Tags RelasiGuruMapel
// @Produce json
// @Success 200 {array} models.RelasiGuruMapel
// @Router /relasi-guru-mapel [get]
func (h *RelasiGuruMapelHandler) GetAll(c *fiber.Ctx) error {
	items, err := h.service.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(items)
}

// GetByIDRelasiGuruMapel retrieves a relasi guru mapel by ID
// @Summary Get relasi guru mapel by ID
// @Description Get a specific relasi guru mapel by ID
// @Tags RelasiGuruMapel
// @Produce json
// @Param id path int true "Relasi ID"
// @Success 200 {object} models.RelasiGuruMapel
// @Failure 404 {object} map[string]interface{}
// @Router /relasi-guru-mapel/{id} [get]
func (h *RelasiGuruMapelHandler) GetByID(c *fiber.Ctx) error {
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

// UpdateRelasiGuruMapel updates a relasi guru mapel
// @Summary Update relasi guru mapel
// @Description Update an existing relasi guru mapel
// @Tags RelasiGuruMapel
// @Accept json
// @Produce json
// @Param id path int true "Relasi ID"
// @Param relasi body models.RelasiGuruMapel true "Updated relasi data"
// @Success 200 {object} models.RelasiGuruMapel
// @Failure 400 {object} map[string]interface{}
// @Router /relasi-guru-mapel/{id} [put]
func (h *RelasiGuruMapelHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	var payload models.RelasiGuruMapel
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

// DeleteRelasiGuruMapel deletes a relasi guru mapel
// @Summary Delete relasi guru mapel
// @Description Delete a relasi guru mapel by ID
// @Tags RelasiGuruMapel
// @Param id path int true "Relasi ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Router /relasi-guru-mapel/{id} [delete]
func (h *RelasiGuruMapelHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	if err := h.service.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
