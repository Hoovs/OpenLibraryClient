package db

import (
	"io/ioutil"
	"testing"
)

func setupDb() (*DB, error) {
	db, _ := InitDB(":memory:")
	f, err := ioutil.ReadFile("../../sql/wishlist.sql")
	if err != nil {
		return nil, err
	}
	_, err = db.db.Exec(string(f))
	if err != nil {
		return nil, err
	}

	_, err = db.db.Exec(`INSERT INTO wishlist (id, userId) VALUES (1, 1)`)
	return db, err
}

func TestDeleteWishList(t *testing.T) {
	db, err := setupDb()
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		name     string
		id       int
		expected func(error) bool
	}{
		{
			name: "Delete missing id fails",
			id:   2,
			expected: func(e error) bool {
				return e != nil
			},
		}, {
			name: "Delete for existing id removes it",
			id:   1,
			expected: func(e error) bool {
				return e == nil
			},
		}, {
			name: "Delete for just deleted id fails",
			id:   1,
			expected: func(e error) bool {
				return e != nil
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := db.DeleteWishList(c.id)
			if !c.expected(err) {
				t.Error(err)
			}
		})
	}
}

func TestGetWishList(t *testing.T) {
	db, err := setupDb()
	if err != nil {
		t.Error(err)
	}

	cases := []struct {
		name     string
		id       int
		expected func(error) bool
	}{
		{
			name: "Get for missing id fails",
			id:   2,
			expected: func(e error) bool {
				return e != nil
			},
		}, {
			name: "Get for existing id returns it",
			id:   1,
			expected: func(e error) bool {
				return e == nil
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := db.DeleteWishList(c.id)
			if !c.expected(err) {
				t.Error(err)
			}
		})
	}
}
