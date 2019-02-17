package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

func InitDB(filepath string) (*DB, error) {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, err
	}
	if db == nil {
		return nil, errors.New("db couldn't be read")
	}
	return &DB{db}, nil
}

func (d *DB) DeleteWishList(id int) error {
	sqlDelete := `DELETE FROM wishlist WHERE id = ?`

	stmt, err := d.db.Prepare(sqlDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(id)
	if err != nil {
		return err
	} else if rs == nil {
		return errors.New("wish list record not found")
	}

	rows, err := rs.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return errors.New("wish list record not found")
	}
	return nil
}
