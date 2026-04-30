package repositories

import (
	"database/sql"

	"timetabling/models"
)

type RelasiGuruMapelRepository interface {
	Create(relasi *models.RelasiGuruMapel) (int64, error)
	GetAll() ([]models.RelasiGuruMapel, error)
	GetByID(id int64) (*models.RelasiGuruMapel, error)
	Update(relasi *models.RelasiGuruMapel) error
	Delete(id int64) error
}

type relasiGuruMapelRepository struct {
	db *sql.DB
}

func NewRelasiGuruMapelRepository(db *sql.DB) RelasiGuruMapelRepository {
	return &relasiGuruMapelRepository{db: db}
}

func (r *relasiGuruMapelRepository) Create(relasi *models.RelasiGuruMapel) (int64, error) {
	result, err := r.db.Exec(`INSERT INTO relasi_guru_mapel (guru_id, mapel_id, tingkatan, durasi) VALUES (?, ?, ?, ?)`, relasi.GuruID, relasi.MapelID, relasi.Tingkatan, relasi.Durasi)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *relasiGuruMapelRepository) GetAll() ([]models.RelasiGuruMapel, error) {
	rows, err := r.db.Query(`SELECT id, guru_id, mapel_id, tingkatan, durasi FROM relasi_guru_mapel`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.RelasiGuruMapel
	for rows.Next() {
		var rel models.RelasiGuruMapel
		if err := rows.Scan(&rel.ID, &rel.GuruID, &rel.MapelID, &rel.Tingkatan, &rel.Durasi); err != nil {
			return nil, err
		}
		items = append(items, rel)
	}
	return items, nil
}

func (r *relasiGuruMapelRepository) GetByID(id int64) (*models.RelasiGuruMapel, error) {
	var rel models.RelasiGuruMapel
	row := r.db.QueryRow(`SELECT id, guru_id, mapel_id, tingkatan, durasi FROM relasi_guru_mapel WHERE id = ?`, id)
	if err := row.Scan(&rel.ID, &rel.GuruID, &rel.MapelID, &rel.Tingkatan, &rel.Durasi); err != nil {
		return nil, err
	}
	return &rel, nil
}

func (r *relasiGuruMapelRepository) Update(relasi *models.RelasiGuruMapel) error {
	_, err := r.db.Exec(`UPDATE relasi_guru_mapel SET guru_id = ?, mapel_id = ?, tingkatan = ?, durasi = ? WHERE id = ?`, relasi.GuruID, relasi.MapelID, relasi.Tingkatan, relasi.Durasi, relasi.ID)
	return err
}

func (r *relasiGuruMapelRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM relasi_guru_mapel WHERE id = ?`, id)
	return err
}
