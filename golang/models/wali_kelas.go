package models

type WaliKelas struct {
	ID      int64 `json:"id" db:"id"`
	GuruID  int64 `json:"guru_id" db:"guru_id"`
	KelasID int64 `json:"kelas_id" db:"kelas_id"`
}
