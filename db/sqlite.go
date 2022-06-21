package db

import "database/sql"

func open() (*sql.DB, error) {
	db, err := sql.Open("", "")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Query(sql string, args ...any) error {
	db, err := open()
	if err != nil {
		return err
	}
	rows, err := db.Query(sql, args)
	if err != nil {
		return err
	}
	defer db.Close()
	defer rows.Close()
	cls, err := rows.Columns()
	if err != nil {
		return err
	}
}
