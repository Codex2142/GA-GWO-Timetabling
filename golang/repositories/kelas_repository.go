package repositories

import (
	"database/sql"

	"timetabling/models"
)

type KelasRepository interface {
	Create(kelas *models.Kelas) (int64, error)
	GetAll() ([]models.Kelas, error)
	GetByID(id int64) (*models.Kelas, error)
	Update(kelas *models.Kelas) error
	Delete(id int64) error
}

type kelasRepository struct {
	db *sql.DB
}

func NewKelasRepository(db *sql.DB) KelasRepository {
	return &kelasRepository{db: db}
}

func (r *kelasRepository) Create(kelas *models.Kelas) (int64, error) {
	result, err := r.db.Exec(`INSERT INTO kelas (nama_kelas, tingkatan) VALUES (?, ?)`, kelas.NamaKelas, kelas.Tingkatan)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *kelasRepository) GetAll() ([]models.Kelas, error) {
	rows, err := r.db.Query(`SELECT kelas_id, nama_kelas, tingkatan FROM kelas`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Kelas
	for rows.Next() {
		var k models.Kelas
		if err := rows.Scan(&k.KelasID, &k.NamaKelas, &k.Tingkatan); err != nil {
			return nil, err
		}
		items = append(items, k)
	}
	return items, nil
}

func (r *kelasRepository) GetByID(id int64) (*models.Kelas, error) {
	var k models.Kelas
	row := r.db.QueryRow(`SELECT kelas_id, nama_kelas, tingkatan FROM kelas WHERE kelas_id = ?`, id)
	if err := row.Scan(&k.KelasID, &k.NamaKelas, &k.Tingkatan); err != nil {
		return nil, err
	}
	return &k, nil
}

func (r *kelasRepository) Update(kelas *models.Kelas) error {
	_, err := r.db.Exec(`UPDATE kelas SET nama_kelas = ?, tingkatan = ? WHERE kelas_id = ?`, kelas.NamaKelas, kelas.Tingkatan, kelas.KelasID)
	return err
}

func (r *kelasRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM kelas WHERE kelas_id = ?`, id)
	return err
}
