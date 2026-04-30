package repositories

import (
	"database/sql"

	"timetabling/models"
)

type GuruRepository interface {
	Create(guru *models.Guru) (int64, error)
	GetAll() ([]models.Guru, error)
	GetByID(id int64) (*models.Guru, error)
	Update(guru *models.Guru) error
	Delete(id int64) error
}

type guruRepository struct {
	db *sql.DB
}

func NewGuruRepository(db *sql.DB) GuruRepository {
	return &guruRepository{db: db}
}

func (r *guruRepository) Create(guru *models.Guru) (int64, error) {
	result, err := r.db.Exec(`INSERT INTO guru (nama_guru) VALUES (?)`, guru.NamaGuru)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *guruRepository) GetAll() ([]models.Guru, error) {
	rows, err := r.db.Query(`SELECT guru_id, nama_guru FROM guru`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var gurus []models.Guru
	for rows.Next() {
		var g models.Guru
		if err := rows.Scan(&g.GuruID, &g.NamaGuru); err != nil {
			return nil, err
		}
		gurus = append(gurus, g)
	}
	return gurus, nil
}

func (r *guruRepository) GetByID(id int64) (*models.Guru, error) {
	var g models.Guru
	row := r.db.QueryRow(`SELECT guru_id, nama_guru FROM guru WHERE guru_id = ?`, id)
	if err := row.Scan(&g.GuruID, &g.NamaGuru); err != nil {
		return nil, err
	}
	return &g, nil
}

func (r *guruRepository) Update(guru *models.Guru) error {
	_, err := r.db.Exec(`UPDATE guru SET nama_guru = ? WHERE guru_id = ?`, guru.NamaGuru, guru.GuruID)
	return err
}

func (r *guruRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM guru WHERE guru_id = ?`, id)
	return err
}
