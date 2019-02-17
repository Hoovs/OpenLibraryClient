package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

// DB wraps a sqlite DB with specific calls for wish list.
type DB struct {
	db *sql.DB
}

// WishListRow is a json annotated representation of the db schema.
type WishListRow struct {
	UserId    int    `json:"userId"`
	BookTitle string `json:"bookTitle"`
	Deleted   bool   `json:"deleted"`
}

// InitDB opens a filepath or :memory: and returns bacl the pointer to the db.
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

// DeleteWishList removes a record for the given id.
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
	} else if rows == 0 {
		return errors.New("wish list record not found")
	}
	return nil
}

// GetWishList returns back the search endpoints values plus the marshaled db row.
// ie) if a row is "1984" -> you get back /search?q=1984 + {.., "bookTitle":"1984".. }
func (d *DB) GetWishList(id int) (*WishListRow, error) {
	sqlSelect := `SELECT userId, booktitle, deleted FROM wishlist WHERE id = ?`

	stmt, err := d.db.Prepare(sqlSelect)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rs := stmt.QueryRow(id)
	if rs == nil {
		return nil, errors.New("wish list record not found")
	}
	w := WishListRow{}
	err = rs.Scan(&w.UserId, &w.BookTitle, &w.Deleted)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

// InsertRow unmarshales a request body and checks for the required fields.
// If the fields are valid it inserts it into the db.
func (d *DB) InsertRow(row WishListRow) error {
	if row.BookTitle == "" {
		return errors.New("bookTitle cannot be empty")
	} else if row.UserId == 0 {
		return errors.New("userId cannot be empty")
	}

	sqlInsert := `INSERT INTO wishlist (userid, booktitle) VALUES (?, ?)`
	stmt, err := d.db.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(row.UserId, row.BookTitle)
	return err
}
