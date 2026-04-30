package services

import (
	"errors"

	"timetabling/models"
	"timetabling/repositories"
)

type RelasiGuruMapelService interface {
	Create(relasi *models.RelasiGuruMapel) (*models.RelasiGuruMapel, error)
	GetAll() ([]models.RelasiGuruMapel, error)
	GetByID(id int64) (*models.RelasiGuruMapel, error)
	Update(relasi *models.RelasiGuruMapel) (*models.RelasiGuruMapel, error)
	Delete(id int64) error
}

type relasiGuruMapelService struct {
	repo      repositories.RelasiGuruMapelRepository
	guruRepo  repositories.GuruRepository
	mapelRepo repositories.MapelRepository
}

func NewRelasiGuruMapelService(repo repositories.RelasiGuruMapelRepository, guruRepo repositories.GuruRepository, mapelRepo repositories.MapelRepository) RelasiGuruMapelService {
	return &relasiGuruMapelService{repo: repo, guruRepo: guruRepo, mapelRepo: mapelRepo}
}

func (s *relasiGuruMapelService) Create(relasi *models.RelasiGuruMapel) (*models.RelasiGuruMapel, error) {
	if relasi.GuruID == 0 || relasi.MapelID == 0 || relasi.Tingkatan <= 0 || relasi.Durasi <= 0 {
		return nil, errors.New("guru_id, mapel_id, tingkatan, dan durasi harus diisi")
	}
	if _, err := s.guruRepo.GetByID(relasi.GuruID); err != nil {
		return nil, errors.New("guru tidak ditemukan")
	}
	if _, err := s.mapelRepo.GetByID(relasi.MapelID); err != nil {
		return nil, errors.New("mapel tidak ditemukan")
	}
	id, err := s.repo.Create(relasi)
	if err != nil {
		return nil, err
	}
	relasi.ID = id
	return relasi, nil
}

func (s *relasiGuruMapelService) GetAll() ([]models.RelasiGuruMapel, error) {
	return s.repo.GetAll()
}

func (s *relasiGuruMapelService) GetByID(id int64) (*models.RelasiGuruMapel, error) {
	return s.repo.GetByID(id)
}

func (s *relasiGuruMapelService) Update(relasi *models.RelasiGuruMapel) (*models.RelasiGuruMapel, error) {
	if relasi.ID == 0 {
		return nil, errors.New("id tidak boleh 0")
	}
	if relasi.GuruID == 0 || relasi.MapelID == 0 || relasi.Tingkatan <= 0 || relasi.Durasi <= 0 {
		return nil, errors.New("guru_id, mapel_id, tingkatan, dan durasi harus diisi")
	}
	if _, err := s.guruRepo.GetByID(relasi.GuruID); err != nil {
		return nil, errors.New("guru tidak ditemukan")
	}
	if _, err := s.mapelRepo.GetByID(relasi.MapelID); err != nil {
		return nil, errors.New("mapel tidak ditemukan")
	}
	if err := s.repo.Update(relasi); err != nil {
		return nil, err
	}
	return relasi, nil
}

func (s *relasiGuruMapelService) Delete(id int64) error {
	return s.repo.Delete(id)
}
