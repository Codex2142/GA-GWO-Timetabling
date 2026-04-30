package services

import (
	"errors"

	"timetabling/models"
	"timetabling/repositories"
)

type SlotService interface {
	Create(slot *models.Slot) (*models.Slot, error)
	GetAll() ([]models.Slot, error)
	GetByID(id int64) (*models.Slot, error)
	Update(slot *models.Slot) (*models.Slot, error)
	Delete(id int64) error
}

type slotService struct {
	repo repositories.SlotRepository
}

func NewSlotService(repo repositories.SlotRepository) SlotService {
	return &slotService{repo: repo}
}

func (s *slotService) Create(slot *models.Slot) (*models.Slot, error) {
	if slot.Hari == "" || slot.JamMulai == "" || slot.JamSelesai == "" || slot.JenisSlot == "" {
		return nil, errors.New("hari, jam_mulai, jam_selesai, dan jenis_slot harus diisi")
	}
	id, err := s.repo.Create(slot)
	if err != nil {
		return nil, err
	}
	slot.SlotID = id
	return slot, nil
}

func (s *slotService) GetAll() ([]models.Slot, error) {
	return s.repo.GetAll()
}

func (s *slotService) GetByID(id int64) (*models.Slot, error) {
	return s.repo.GetByID(id)
}

func (s *slotService) Update(slot *models.Slot) (*models.Slot, error) {
	if slot.SlotID == 0 {
		return nil, errors.New("slot_id tidak boleh 0")
	}
	if slot.Hari == "" || slot.JamMulai == "" || slot.JamSelesai == "" || slot.JenisSlot == "" {
		return nil, errors.New("hari, jam_mulai, jam_selesai, dan jenis_slot harus diisi")
	}
	if err := s.repo.Update(slot); err != nil {
		return nil, err
	}
	return slot, nil
}

func (s *slotService) Delete(id int64) error {
	return s.repo.Delete(id)
}
