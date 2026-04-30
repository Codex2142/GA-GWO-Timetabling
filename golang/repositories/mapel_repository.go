package repositories

import (
	"database/sql"

	"timetabling/models"
)

type MapelRepository interface {
	Create(mapel *models.Mapel) (int64, error)
	GetAll() ([]models.Mapel, error)
	GetByID(id int64) (*models.Mapel, error)
	Update(mapel *models.Mapel) error
	Delete(id int64) error
}

type mapelRepository struct {
	db *sql.DB
}

func NewMapelRepository(db *sql.DB) MapelRepository {
	return &mapelRepository{db: db}
}

func (r *mapelRepository) Create(mapel *models.Mapel) (int64, error) {
	result, err := r.db.Exec(`INSERT INTO mapel (kode_mapel, nama_mapel, mgmp, jam_per_minggu) VALUES (?, ?, ?, ?)`, mapel.KodeMapel, mapel.NamaMapel, mapel.Mgmp, mapel.JamPerMinggu)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *mapelRepository) GetAll() ([]models.Mapel, error) {
	rows, err := r.db.Query(`SELECT mapel_id, kode_mapel, nama_mapel, mgmp, jam_per_minggu FROM mapel`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Mapel
	for rows.Next() {
		var m models.Mapel
		if err := rows.Scan(&m.MapelID, &m.KodeMapel, &m.NamaMapel, &m.Mgmp, &m.JamPerMinggu); err != nil {
			return nil, err
		}
		items = append(items, m)
	}
	return items, nil
}

func (r *mapelRepository) GetByID(id int64) (*models.Mapel, error) {
	var m models.Mapel
	row := r.db.QueryRow(`SELECT mapel_id, kode_mapel, nama_mapel, mgmp, jam_per_minggu FROM mapel WHERE mapel_id = ?`, id)
	if err := row.Scan(&m.MapelID, &m.KodeMapel, &m.NamaMapel, &m.Mgmp, &m.JamPerMinggu); err != nil {
		return nil, err
	}
	return &m, nil
}

func (r *mapelRepository) Update(mapel *models.Mapel) error {
	_, err := r.db.Exec(`UPDATE mapel SET kode_mapel = ?, nama_mapel = ?, mgmp = ?, jam_per_minggu = ? WHERE mapel_id = ?`, mapel.KodeMapel, mapel.NamaMapel, mapel.Mgmp, mapel.JamPerMinggu, mapel.MapelID)
	return err
}

func (r *mapelRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM mapel WHERE mapel_id = ?`, id)
	return err
}
