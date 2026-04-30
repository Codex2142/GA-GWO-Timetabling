package services

import (
	"errors"

	"timetabling/models"
	"timetabling/repositories"
)

type GuruService interface {
	Create(guru *models.Guru) (*models.Guru, error)
	GetAll() ([]models.Guru, error)
	GetByID(id int64) (*models.Guru, error)
	Update(guru *models.Guru) (*models.Guru, error)
	Delete(id int64) error
}

type guruService struct {
	repo repositories.GuruRepository
}

func NewGuruService(repo repositories.GuruRepository) GuruService {
	return &guruService{repo: repo}
}

func (s *guruService) Create(guru *models.Guru) (*models.Guru, error) {
	if guru.NamaGuru == "" {
		return nil, errors.New("nama_guru harus diisi")
	}

	id, err := s.repo.Create(guru)
	if err != nil {
		return nil, err
	}
	guru.GuruID = id
	return guru, nil
}

func (s *guruService) GetAll() ([]models.Guru, error) {
	return s.repo.GetAll()
}

func (s *guruService) GetByID(id int64) (*models.Guru, error) {
	return s.repo.GetByID(id)
}

func (s *guruService) Update(guru *models.Guru) (*models.Guru, error) {
	if guru.GuruID == 0 {
		return nil, errors.New("guru_id tidak boleh 0")
	}
	if guru.NamaGuru == "" {
		return nil, errors.New("nama_guru harus diisi")
	}
	if err := s.repo.Update(guru); err != nil {
		return nil, err
	}
	return guru, nil
}

func (s *guruService) Delete(id int64) error {
	return s.repo.Delete(id)
}
