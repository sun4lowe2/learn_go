package main

import (
	"database/sql"
	"github.com/pkg/errors"
)

type user struct {
	id int
	name string
}
func mock() (user, error) {
	return user{}, sql.ErrNoRows
}

func query() (user, error) {
	user, err := mock()

	if err != nil {
		return user, errors.Wrap(err, "query no result")
	}
	return user, nil
}

func main() {
	_, err := query()
	if errors.Cause(err) == sql.ErrNoRows {
		println("no result")
	}
}
