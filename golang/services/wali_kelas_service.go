package services

import (
	"errors"

	"timetabling/models"
	"timetabling/repositories"
)

type WaliKelasService interface {
	Create(wali *models.WaliKelas) (*models.WaliKelas, error)
	GetAll() ([]models.WaliKelas, error)
	GetByID(id int64) (*models.WaliKelas, error)
	Update(wali *models.WaliKelas) (*models.WaliKelas, error)
	Delete(id int64) error
}

type waliKelasService struct {
	repo      repositories.WaliKelasRepository
	guruRepo  repositories.GuruRepository
	kelasRepo repositories.KelasRepository
}

func NewWaliKelasService(repo repositories.WaliKelasRepository, guruRepo repositories.GuruRepository, kelasRepo repositories.KelasRepository) WaliKelasService {
	return &waliKelasService{repo: repo, guruRepo: guruRepo, kelasRepo: kelasRepo}
}

func (s *waliKelasService) Create(wali *models.WaliKelas) (*models.WaliKelas, error) {
	if wali.GuruID == 0 || wali.KelasID == 0 {
		return nil, errors.New("guru_id dan kelas_id harus diisi")
	}
	if _, err := s.guruRepo.GetByID(wali.GuruID); err != nil {
		return nil, errors.New("guru tidak ditemukan")
	}
	if _, err := s.kelasRepo.GetByID(wali.KelasID); err != nil {
		return nil, errors.New("kelas tidak ditemukan")
	}
	id, err := s.repo.Create(wali)
	if err != nil {
		return nil, err
	}
	wali.ID = id
	return wali, nil
}

func (s *waliKelasService) GetAll() ([]models.WaliKelas, error) {
	return s.repo.GetAll()
}

func (s *waliKelasService) GetByID(id int64) (*models.WaliKelas, error) {
	return s.repo.GetByID(id)
}

func (s *waliKelasService) Update(wali *models.WaliKelas) (*models.WaliKelas, error) {
	if wali.ID == 0 {
		return nil, errors.New("id tidak boleh 0")
	}
	if wali.GuruID == 0 || wali.KelasID == 0 {
		return nil, errors.New("guru_id dan kelas_id harus diisi")
	}
	if _, err := s.guruRepo.GetByID(wali.GuruID); err != nil {
		return nil, errors.New("guru tidak ditemukan")
	}
	if _, err := s.kelasRepo.GetByID(wali.KelasID); err != nil {
		return nil, errors.New("kelas tidak ditemukan")
	}
	if err := s.repo.Update(wali); err != nil {
		return nil, err
	}
	return wali, nil
}

func (s *waliKelasService) Delete(id int64) error {
	return s.repo.Delete(id)
}
