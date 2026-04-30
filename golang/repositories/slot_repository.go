package repositories

import (
	"database/sql"

	"timetabling/models"
)

type SlotRepository interface {
	Create(slot *models.Slot) (int64, error)
	GetAll() ([]models.Slot, error)
	GetByID(id int64) (*models.Slot, error)
	Update(slot *models.Slot) error
	Delete(id int64) error
}

type slotRepository struct {
	db *sql.DB
}

func NewSlotRepository(db *sql.DB) SlotRepository {
	return &slotRepository{db: db}
}

func (r *slotRepository) Create(slot *models.Slot) (int64, error) {
	result, err := r.db.Exec(`INSERT INTO slot (hari, jam_mulai, jam_selesai, jenis_slot) VALUES (?, ?, ?, ?)`, slot.Hari, slot.JamMulai, slot.JamSelesai, slot.JenisSlot)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *slotRepository) GetAll() ([]models.Slot, error) {
	rows, err := r.db.Query(`SELECT slot_id, hari, jam_mulai, jam_selesai, jenis_slot FROM slot`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Slot
	for rows.Next() {
		var s models.Slot
		if err := rows.Scan(&s.SlotID, &s.Hari, &s.JamMulai, &s.JamSelesai, &s.JenisSlot); err != nil {
			return nil, err
		}
		items = append(items, s)
	}
	return items, nil
}

func (r *slotRepository) GetByID(id int64) (*models.Slot, error) {
	var s models.Slot
	row := r.db.QueryRow(`SELECT slot_id, hari, jam_mulai, jam_selesai, jenis_slot FROM slot WHERE slot_id = ?`, id)
	if err := row.Scan(&s.SlotID, &s.Hari, &s.JamMulai, &s.JamSelesai, &s.JenisSlot); err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *slotRepository) Update(slot *models.Slot) error {
	_, err := r.db.Exec(`UPDATE slot SET hari = ?, jam_mulai = ?, jam_selesai = ?, jenis_slot = ? WHERE slot_id = ?`, slot.Hari, slot.JamMulai, slot.JamSelesai, slot.JenisSlot, slot.SlotID)
	return err
}

func (r *slotRepository) Delete(id int64) error {
	_, err := r.db.Exec(`DELETE FROM slot WHERE slot_id = ?`, id)
	return err
}
