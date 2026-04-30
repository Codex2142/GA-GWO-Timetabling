package handlers

import (
	"strconv"

	"timetabling/models"
	"timetabling/services"

	"github.com/gofiber/fiber/v2"
)

type MapelHandler struct {
	service services.MapelService
}

func NewMapelHandler(service services.MapelService) *MapelHandler {
	return &MapelHandler{service: service}
}

// CreateMapel creates a new mapel
// @Summary Create mapel
// @Description Create a new mapel
// @Tags Mapel
// @Accept json
// @Produce json
// @Param mapel body models.Mapel true "Mapel data"
// @Success 201 {object} models.Mapel
// @Failure 400 {object} map[string]interface{}
// @Router /mapel [post]
func (h *MapelHandler) Create(c *fiber.Ctx) error {
	var payload models.Mapel
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	result, err := h.service.Create(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

// GetAllMapel retrieves all mapel
// @Summary Get all mapel
// @Description Get a list of all mapel
// @Tags Mapel
// @Produce json
// @Success 200 {array} models.Mapel
// @Router /mapel [get]
func (h *MapelHandler) GetAll(c *fiber.Ctx) error {
	items, err := h.service.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(items)
}

// GetByIDMapel retrieves a mapel by ID
// @Summary Get mapel by ID
// @Description Get a specific mapel by ID
// @Tags Mapel
// @Produce json
// @Param id path int true "Mapel ID"
// @Success 200 {object} models.Mapel
// @Failure 404 {object} map[string]interface{}
// @Router /mapel/{id} [get]
func (h *MapelHandler) GetByID(c *fiber.Ctx) error {
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

// UpdateMapel updates a mapel
// @Summary Update mapel
// @Description Update an existing mapel
// @Tags Mapel
// @Accept json
// @Produce json
// @Param id path int true "Mapel ID"
// @Param mapel body models.Mapel true "Updated mapel data"
// @Success 200 {object} models.Mapel
// @Failure 400 {object} map[string]interface{}
// @Router /mapel/{id} [put]
func (h *MapelHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	var payload models.Mapel
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	payload.MapelID = id
	result, err := h.service.Update(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(result)
}

// DeleteMapel deletes a mapel
// @Summary Delete mapel
// @Description Delete a mapel by ID
// @Tags Mapel
// @Param id path int true "Mapel ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Router /mapel/{id} [delete]
func (h *MapelHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	if err := h.service.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
