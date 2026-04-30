package models

type Guru struct {
	GuruID   int64  `json:"guru_id" db:"guru_id"`
	NamaGuru string `json:"nama_guru" db:"nama_guru"`
}
