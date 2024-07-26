package database

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"os"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB(filepath string) (*sql.DB, error) {
	var err error
	DB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	if DB == nil {
		return nil, errors.New("db is nil")
	}

	// Create table if not exists
	_, err = DB.Exec(`CREATE TABLE IF NOT EXISTS addresses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		field1 INTEGER,
		field2 INTEGER,
		field3 INTEGER,
		field4 TEXT,
		field5 TEXT,
		field6 TEXT,
		field7 TEXT,
		field8 TEXT,
		field9 TEXT,
		field10 INTEGER,
		field11 INTEGER,
		field12 INTEGER,
		field13 INTEGER,
		field14 INTEGER,
		field15 INTEGER
	)`)
	if err != nil && err.Error() != "table addresses already exists" {
		return nil, err
	}

	return DB, nil
}

func ImportCSV(filepath string) error {
	if DB == nil {
		return errors.New("database not initialized")
	}

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO addresses (field1, field2, field3, field4, field5, field6, field7, field8, field9, field10, field11, field12, field13, field14, field15) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, record := range records {
		_, err = stmt.Exec(record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9], record[10], record[11], record[12], record[13], record[14])
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
