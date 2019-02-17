package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	db *sql.DB
}

type WishListRow struct {
	UserId  int  `json:"userId"`
	BookId  int  `json:"bookId"`
	Deleted bool `json:"deleted"`
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

func (d *DB) GetWishList(id int) (*WishListRow, error) {
	sqlSelect := `SELECT userId, deleted FROM wishlist WHERE id = ?`

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
	err = rs.Scan(&w.UserId, &w.Deleted)
	if err != nil {
		return nil, err
	}
	return &w, nil
}

func (d *DB) InsertRow(row WishListRow) error {
	if row.BookId == 0 {
		return errors.New("bookId cannot be empty")
	}
	if row.UserId == 0 {
		return errors.New("userId cannot be empty")
	}

	sqlInsert := `INSERT INTO wishlist (userid, bookid) VALUES (?, ?)`

	stmt, err := d.db.Prepare(sqlInsert)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(row.UserId, row.BookId)
	return err
}
