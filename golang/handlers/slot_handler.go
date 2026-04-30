package handlers

import (
	"strconv"

	"timetabling/models"
	"timetabling/services"

	"github.com/gofiber/fiber/v2"
)

type SlotHandler struct {
	service services.SlotService
}

func NewSlotHandler(service services.SlotService) *SlotHandler {
	return &SlotHandler{service: service}
}

// CreateSlot creates a new slot
// @Summary Create slot
// @Description Create a new slot
// @Tags Slot
// @Accept json
// @Produce json
// @Param slot body models.Slot true "Slot data"
// @Success 201 {object} models.Slot
// @Failure 400 {object} map[string]interface{}
// @Router /slot [post]
func (h *SlotHandler) Create(c *fiber.Ctx) error {
	var payload models.Slot
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	result, err := h.service.Create(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

// GetAllSlot retrieves all slot
// @Summary Get all slot
// @Description Get a list of all slot
// @Tags Slot
// @Produce json
// @Success 200 {array} models.Slot
// @Router /slot [get]
func (h *SlotHandler) GetAll(c *fiber.Ctx) error {
	items, err := h.service.GetAll()
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.JSON(items)
}

// GetByIDSlot retrieves a slot by ID
// @Summary Get slot by ID
// @Description Get a specific slot by ID
// @Tags Slot
// @Produce json
// @Param id path int true "Slot ID"
// @Success 200 {object} models.Slot
// @Failure 404 {object} map[string]interface{}
// @Router /slot/{id} [get]
func (h *SlotHandler) GetByID(c *fiber.Ctx) error {
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

// UpdateSlot updates a slot
// @Summary Update slot
// @Description Update an existing slot
// @Tags Slot
// @Accept json
// @Produce json
// @Param id path int true "Slot ID"
// @Param slot body models.Slot true "Updated slot data"
// @Success 200 {object} models.Slot
// @Failure 400 {object} map[string]interface{}
// @Router /slot/{id} [put]
func (h *SlotHandler) Update(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	var payload models.Slot
	if err := c.BodyParser(&payload); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	payload.SlotID = id
	result, err := h.service.Update(&payload)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	return c.JSON(result)
}

// DeleteSlot deletes a slot
// @Summary Delete slot
// @Description Delete a slot by ID
// @Tags Slot
// @Param id path int true "Slot ID"
// @Success 204
// @Failure 400 {object} map[string]interface{}
// @Router /slot/{id} [delete]
func (h *SlotHandler) Delete(c *fiber.Ctx) error {
	id, err := strconv.ParseInt(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid id")
	}
	if err := h.service.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}
