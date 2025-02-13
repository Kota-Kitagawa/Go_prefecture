package database

import (
	"database/sql"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"sync"
	_ "github.com/mattn/go-sqlite3"
)

var(
	DB *sql.DB
	once sync.Once
)


func InitDB(filepath string) (*sql.DB, error) {
	var Initerr error
	once.Do(func() {
		var err error
		DB, err = sql.Open("sqlite3", filepath)
		if err != nil {
			Initerr=err
			return 
		}
		if DB == nil {
			Initerr = errors.New("db is nil")
			return
		}

		_, err = DB.Exec(`PRAGMA encoding = "UTF-8";`)
		if err != nil {
			Initerr = err
			return
		}

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
			Initerr = err
			return
		}
	})
	if Initerr != nil {
		return nil, Initerr
	}
	return DB, nil
}

func NormalizeTable() error {
	tx, err := DB.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`
    DROP TABLE IF EXISTS normalized_utf_ken_all;

    CREATE TABLE IF NOT EXISTS normalized_utf_ken_all (
        field1 INTEGER,
        field2 INTEGER,
        field3 INTEGER,
        field4 TEXT,
        field5 TEXT,
        field6 TEXT,
        field7 TEXT,
        field8 TEXT,
        Normalizedfield9 TEXT,
        OutsideParentheses TEXT,
        InsideParentheses TEXT,
        PRIMARY KEY (Normalizedfield9, field7, field8)
    );

    INSERT OR IGNORE INTO normalized_utf_ken_all (field1, field2, field3, field4, field5, field6, field7, field8, Normalizedfield9, OutsideParentheses, InsideParentheses)
    SELECT field1, field2, field3, field4, field5, field6, field7, field8,
        CASE 
            WHEN field9 = '以下に掲載がない場合' THEN '未掲載'
            ELSE field9
        END AS Normalizedfield9,
        CASE 
            WHEN field9 LIKE '%（%' THEN substr(field9, 1, instr(field9, '（') - 1)
            ELSE field9
        END AS OutsideParentheses,
        CASE 
            WHEN field9 LIKE '%（%' THEN substr(field9, instr(field9, '（') + 1, instr(field9, '）') - instr(field9, '（') - 1)
            ELSE NULL
        END AS InsideParentheses
    FROM addresses;
`)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
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

	// Shift_JISエンコーディングのCSVファイルをUTF-8として読み込む
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

	for i, record := range records {
		if i ==0 {
			continue
		}
		if i < 5 {
			fmt.Printf("Record %d: %v\n", i, record)
		}
		_, err = stmt.Exec(record[0], record[1], record[2], record[3], record[4], record[5], record[6], record[7], record[8], record[9], record[10], record[11], record[12], record[13], record[14])
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
