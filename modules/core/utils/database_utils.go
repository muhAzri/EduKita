package utils

import (
	"database/sql"
)

func WithTransaction(db *sql.DB, operation func(tx *sql.Tx) error) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			if err := tx.Rollback(); err != nil {
				panic(err)
			}
			panic(r)
		} else {
			if err := tx.Commit(); err != nil {
				panic(err)
			}
		}
	}()
	if err := operation(tx); err != nil {
		return err
	}
	return nil
}
