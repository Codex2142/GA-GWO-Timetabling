package services

import (
	"errors"

	"timetabling/models"
	"timetabling/repositories"
)

type KelasService interface {
	Create(kelas *models.Kelas) (*models.Kelas, error)
	GetAll() ([]models.Kelas, error)
	GetByID(id int64) (*models.Kelas, error)
	Update(kelas *models.Kelas) (*models.Kelas, error)
	Delete(id int64) error
}

type kelasService struct {
	repo repositories.KelasRepository
}

func NewKelasService(repo repositories.KelasRepository) KelasService {
	return &kelasService{repo: repo}
}

func (s *kelasService) Create(kelas *models.Kelas) (*models.Kelas, error) {
	if kelas.NamaKelas == "" {
		return nil, errors.New("nama_kelas harus diisi")
	}
	if kelas.Tingkatan <= 0 {
		return nil, errors.New("tingkatan harus lebih besar dari 0")
	}
	id, err := s.repo.Create(kelas)
	if err != nil {
		return nil, err
	}
	kelas.KelasID = id
	return kelas, nil
}

func (s *kelasService) GetAll() ([]models.Kelas, error) {
	return s.repo.GetAll()
}

func (s *kelasService) GetByID(id int64) (*models.Kelas, error) {
	return s.repo.GetByID(id)
}

func (s *kelasService) Update(kelas *models.Kelas) (*models.Kelas, error) {
	if kelas.KelasID == 0 {
		return nil, errors.New("kelas_id tidak boleh 0")
	}
	if kelas.NamaKelas == "" {
		return nil, errors.New("nama_kelas harus diisi")
	}
	if kelas.Tingkatan <= 0 {
		return nil, errors.New("tingkatan harus lebih besar dari 0")
	}
	if err := s.repo.Update(kelas); err != nil {
		return nil, err
	}
	return kelas, nil
}

func (s *kelasService) Delete(id int64) error {
	return s.repo.Delete(id)
}
