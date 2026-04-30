package models

type RelasiGuruMapel struct {
	ID        int64 `json:"id" db:"id"`
	GuruID    int64 `json:"guru_id" db:"guru_id"`
	MapelID   int64 `json:"mapel_id" db:"mapel_id"`
	Tingkatan int   `json:"tingkatan" db:"tingkatan"`
	Durasi    int   `json:"durasi" db:"durasi"`
}
