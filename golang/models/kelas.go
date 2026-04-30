package models

type Kelas struct {
	KelasID   int64  `json:"kelas_id" db:"kelas_id"`
	NamaKelas string `json:"nama_kelas" db:"nama_kelas"`
	Tingkatan int    `json:"tingkatan" db:"tingkatan"`
}
