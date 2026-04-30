package models

type Slot struct {
	SlotID     int64  `json:"slot_id" db:"slot_id"`
	Hari       string `json:"hari" db:"hari"`
	JamMulai   string `json:"jam_mulai" db:"jam_mulai"`
	JamSelesai string `json:"jam_selesai" db:"jam_selesai"`
	JenisSlot  string `json:"jenis_slot" db:"jenis_slot"`
}
