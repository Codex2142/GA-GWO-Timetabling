package services

import (
	"errors"

	"timetabling/models"
	"timetabling/repositories"
)

type MapelService interface {
	Create(mapel *models.Mapel) (*models.Mapel, error)
	GetAll() ([]models.Mapel, error)
	GetByID(id int64) (*models.Mapel, error)
	Update(mapel *models.Mapel) (*models.Mapel, error)
	Delete(id int64) error
}

type mapelService struct {
	repo repositories.MapelRepository
}

func NewMapelService(repo repositories.MapelRepository) MapelService {
	return &mapelService{repo: repo}
}

func (s *mapelService) Create(mapel *models.Mapel) (*models.Mapel, error) {
	if mapel.KodeMapel == "" || mapel.NamaMapel == "" || mapel.Mgmp == "" {
		return nil, errors.New("kode_mapel, nama_mapel, dan mgmp harus diisi")
	}
	if mapel.JamPerMinggu <= 0 {
		return nil, errors.New("jam_per_minggu harus lebih besar dari 0")
	}
	id, err := s.repo.Create(mapel)
	if err != nil {
		return nil, err
	}
	mapel.MapelID = id
	return mapel, nil
}

func (s *mapelService) GetAll() ([]models.Mapel, error) {
	return s.repo.GetAll()
}

func (s *mapelService) GetByID(id int64) (*models.Mapel, error) {
	return s.repo.GetByID(id)
}

func (s *mapelService) Update(mapel *models.Mapel) (*models.Mapel, error) {
	if mapel.MapelID == 0 {
		return nil, errors.New("mapel_id tidak boleh 0")
	}
	if mapel.KodeMapel == "" || mapel.NamaMapel == "" || mapel.Mgmp == "" {
		return nil, errors.New("kode_mapel, nama_mapel, dan mgmp harus diisi")
	}
	if mapel.JamPerMinggu <= 0 {
		return nil, errors.New("jam_per_minggu harus lebih besar dari 0")
	}
	if err := s.repo.Update(mapel); err != nil {
		return nil, err
	}
	return mapel, nil
}

func (s *mapelService) Delete(id int64) error {
	return s.repo.Delete(id)
}
