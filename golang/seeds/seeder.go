package seeds

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const csvPath = "./data/csv"

func SeedDatabase(db *sql.DB) error {
	if err := seedGuru(db); err != nil {
		return err
	}
	if err := seedKelas(db); err != nil {
		return err
	}
	if err := seedMapel(db); err != nil {
		return err
	}
	if err := seedSlot(db); err != nil {
		return err
	}
	if err := seedRelasiGuruMapel(db); err != nil {
		return err
	}
	if err := seedWaliKelas(db); err != nil {
		return err
	}
	return nil
}

func seedGuru(db *sql.DB) error {
	file, err := os.Open(csvPath + "/guru.csv")
	if err != nil {
		return fmt.Errorf("failed open guru.csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed read guru.csv: %w", err)
	}

	stmt, err := db.Prepare(`INSERT INTO guru (nama_guru) VALUES (?)`)
	if err != nil {
		return fmt.Errorf("failed prepare guru insert: %w", err)
	}
	defer stmt.Close()

	for i, record := range records {
		if i == 0 {
			continue
		}
		// record[0] adalah guru_id, record[1] adalah nama_guru
		if _, err := stmt.Exec(record[1]); err != nil {
			return fmt.Errorf("failed insert guru row %d: %w", i, err)
		}
	}
	return nil
}

func seedKelas(db *sql.DB) error {
	file, err := os.Open(csvPath + "/kelas.csv")
	if err != nil {
		return fmt.Errorf("failed open kelas.csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed read kelas.csv: %w", err)
	}

	stmt, err := db.Prepare(`INSERT INTO kelas (nama_kelas, tingkatan) VALUES (?, ?)`)
	if err != nil {
		return fmt.Errorf("failed prepare kelas insert: %w", err)
	}
	defer stmt.Close()

	for i, record := range records {
		if i == 0 {
			continue
		}
		tingkatan, err := strconv.Atoi(record[2])
		if err != nil {
			return fmt.Errorf("failed parse tingkatan at row %d: %w", i, err)
		}
		if _, err := stmt.Exec(record[1], tingkatan); err != nil {
			return fmt.Errorf("failed insert kelas row %d: %w", i, err)
		}
	}
	return nil
}

func seedMapel(db *sql.DB) error {
	file, err := os.Open(csvPath + "/mapel.csv")
	if err != nil {
		return fmt.Errorf("failed open mapel.csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed read mapel.csv: %w", err)
	}

	stmt, err := db.Prepare(`INSERT INTO mapel (kode_mapel, nama_mapel, mgmp, jam_per_minggu) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed prepare mapel insert: %w", err)
	}
	defer stmt.Close()

	for i, record := range records {
		if i == 0 {
			continue
		}
		jamPerMinggu, err := strconv.Atoi(record[4])
		if err != nil {
			return fmt.Errorf("failed parse jam_per_minggu at row %d: %w", i, err)
		}
		if _, err := stmt.Exec(record[1], record[2], record[3], jamPerMinggu); err != nil {
			return fmt.Errorf("failed insert mapel row %d: %w", i, err)
		}
	}
	return nil
}

func seedSlot(db *sql.DB) error {
	file, err := os.Open(csvPath + "/slot.csv")
	if err != nil {
		return fmt.Errorf("failed open slot.csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed read slot.csv: %w", err)
	}

	stmt, err := db.Prepare(`INSERT INTO slot (hari, jam_mulai, jam_selesai, jenis_slot) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed prepare slot insert: %w", err)
	}
	defer stmt.Close()

	for i, record := range records {
		if i == 0 {
			continue
		}
		// record[0] adalah slot_id
		// record[1]: hari, record[2]: jam_mulai, record[3]: jam_selesai, record[4]: jenis_slot
		if _, err := stmt.Exec(record[1], record[2], record[3], record[4]); err != nil {
			return fmt.Errorf("failed insert slot row %d: %w", i, err)
		}
	}
	return nil
}

func seedRelasiGuruMapel(db *sql.DB) error {
	file, err := os.Open(csvPath + "/relasi_guru_mapel.csv")
	if err != nil {
		return fmt.Errorf("failed open relasi_guru_mapel.csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed read relasi_guru_mapel.csv: %w", err)
	}

	stmt, err := db.Prepare(`INSERT INTO relasi_guru_mapel (guru_id, mapel_id, tingkatan, durasi) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return fmt.Errorf("failed prepare relasi_guru_mapel insert: %w", err)
	}
	defer stmt.Close()

	for i, record := range records {
		if i == 0 {
			continue
		}
		guruID, _ := strconv.ParseInt(strings.TrimSpace(record[0]), 10, 64)
		mapelID, _ := strconv.ParseInt(strings.TrimSpace(record[1]), 10, 64)
		tingkatan, _ := strconv.Atoi(strings.TrimSpace(record[2]))
		durasi, _ := strconv.Atoi(strings.TrimSpace(record[3]))

		if _, err := stmt.Exec(guruID, mapelID, tingkatan, durasi); err != nil {
			return fmt.Errorf("failed insert relasi_guru_mapel row %d: %w", i, err)
		}
	}
	return nil
}

func seedWaliKelas(db *sql.DB) error {
	file, err := os.Open(csvPath + "/wali_kelas.csv")
	if err != nil {
		return fmt.Errorf("failed open wali_kelas.csv: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed read wali_kelas.csv: %w", err)
	}

	stmt, err := db.Prepare(`INSERT INTO wali_kelas (guru_id, kelas_id) VALUES (?, ?)`)
	if err != nil {
		return fmt.Errorf("failed prepare wali_kelas insert: %w", err)
	}
	defer stmt.Close()

	for i, record := range records {
		if i == 0 {
			continue
		}
		guruID, _ := strconv.ParseInt(strings.TrimSpace(record[0]), 10, 64)
		kelasID, _ := strconv.ParseInt(strings.TrimSpace(record[1]), 10, 64)

		if _, err := stmt.Exec(guruID, kelasID); err != nil {
			return fmt.Errorf("failed insert wali_kelas row %d: %w", i, err)
		}
	}
	return nil
}
