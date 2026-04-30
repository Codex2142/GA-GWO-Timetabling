package models

type Mapel struct {
	MapelID      int64  `json:"mapel_id" db:"mapel_id"`
	KodeMapel    string `json:"kode_mapel" db:"kode_mapel"`
	NamaMapel    string `json:"nama_mapel" db:"nama_mapel"`
	Mgmp         string `json:"mgmp" db:"mgmp"`
	JamPerMinggu int    `json:"jam_per_minggu" db:"jam_per_minggu"`
}
