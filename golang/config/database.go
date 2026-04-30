package config

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "./database/app.db"

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		db.Close()
		return nil, err
	}

	if err := createTables(db); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func createTables(db *sql.DB) error {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS guru (
			guru_id INTEGER PRIMARY KEY AUTOINCREMENT,
			nama_guru TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS kelas (
			kelas_id INTEGER PRIMARY KEY AUTOINCREMENT,
			nama_kelas TEXT NOT NULL,
			tingkatan INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS mapel (
			mapel_id INTEGER PRIMARY KEY AUTOINCREMENT,
			kode_mapel TEXT NOT NULL,
			nama_mapel TEXT NOT NULL,
			mgmp TEXT NOT NULL,
			jam_per_minggu INTEGER NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS slot (
			slot_id INTEGER PRIMARY KEY AUTOINCREMENT,
			hari TEXT NOT NULL,
			jam_mulai TEXT NOT NULL,
			jam_selesai TEXT NOT NULL,
			jenis_slot TEXT NOT NULL
		);`,
		`CREATE TABLE IF NOT EXISTS relasi_guru_mapel (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			guru_id INTEGER NOT NULL,
			mapel_id INTEGER NOT NULL,
			tingkatan INTEGER NOT NULL,
			durasi INTEGER NOT NULL,
			FOREIGN KEY (guru_id) REFERENCES guru(guru_id) ON DELETE CASCADE,
			FOREIGN KEY (mapel_id) REFERENCES mapel(mapel_id) ON DELETE CASCADE
		);`,
		`CREATE TABLE IF NOT EXISTS wali_kelas (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			guru_id INTEGER NOT NULL,
			kelas_id INTEGER NOT NULL,
			FOREIGN KEY (guru_id) REFERENCES guru(guru_id) ON DELETE CASCADE,
			FOREIGN KEY (kelas_id) REFERENCES kelas(kelas_id) ON DELETE CASCADE
		);`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			return fmt.Errorf("failed create table: %w", err)
		}
	}

	return nil
}

// ResetDatabase drops all tables and recreates them (fresh start).
func ResetDatabase(db *sql.DB) error {
	// Disable FK temporarily to allow dropping in any order
	if _, err := db.Exec("PRAGMA foreign_keys = OFF"); err != nil {
		return fmt.Errorf("failed to disable foreign keys: %w", err)
	}

	drops := []string{
		`DROP TABLE IF EXISTS wali_kelas;`,
		`DROP TABLE IF EXISTS relasi_guru_mapel;`,
		`DROP TABLE IF EXISTS slot;`,
		`DROP TABLE IF EXISTS mapel;`,
		`DROP TABLE IF EXISTS kelas;`,
		`DROP TABLE IF EXISTS guru;`,
	}

	for _, q := range drops {
		if _, err := db.Exec(q); err != nil {
			return fmt.Errorf("failed drop table: %w", err)
		}
	}

	// Re-enable FK
	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return fmt.Errorf("failed to re-enable foreign keys: %w", err)
	}

	return createTables(db)
}
