package repositories

import (
	"database/sql"

	"timetabling/models"
)

type WaliKelasRepository interface {
	Create(wali *models.WaliKelas) (int64, error)
	GetAll() ([]models.WaliKelas, error)
	GetByID(id int64) (*models.WaliKelas, error)
	Update(wali *models.WaliKelas) error
	Delete(id int64) error
}

type waliKelasRepository struct {
	db *sql.DB
}

func NewWaliKelasRepository(db *sql.DB) WaliKelasRepository {
	return &waliKelasRepository{db: db}
}

func (r *waliKelasRepository) Create(wali *models.WaliKelas) (int64, error) {
	result, err := r.db.Exec(`INSERT INTO wali_kelas (guru_id, kelas_id) VALUES (?, ?)`, wali.GuruID, wali.KelasID)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *waliKelasRepository) GetAll() ([]models.WaliKelas, error) {
	rows, err := r.db.Query(`SELECT id, guru_id, kelas_id FROM wali_kelas`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.WaliKelas
	for rows.Next() {
		var w models.WaliKelas
		if err := rows.Scan(&w.ID, &w.GuruID, &w.KelasID); err != nil {
			return nil, err
		}
		items = append(items, w)
	}
	return items, nil
}

func (r *waliKelasRepository) GetByID(id int64) (*models.WaliKelas, error) {
	var w models.WaliKelas
	row := r.db.QueryRow(`SELECT id, guru_id, kelas_id FROM wali_kelas WHERE id = ?`, id)
	if err := row.Scan(&w.ID, &w.GuruID, &w.KelasID); err != nil {
		return nil, err
	}
	return &w, nil
}

func (r *waliKelasRepository) Update(wali *models.WaliKelas) error {
	_, err := r.db.Exec(`UPDATE wali_kelas SET guru_id = ?, kelas_id = ? WHERE id = ?`, wali.GuruID, wali.KelasID, wali.ID)
	return err
}

func (r *waliKelasRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM wali_kelas WHERE id = ?`, id)
	return err
}
