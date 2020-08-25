package infra

import (
	"database/sql"

	"github.com/thiagotbl/hackatons-api/interfaces"
)

type PostgresDb struct {
	Conn *sql.DB
}

func (db *PostgresDb) Query(statement string) (interfaces.IRows, error) {
	rows, err := db.Conn.Query(statement)
	if err != nil {
		return &emptyRows{}, err
	}

	return rows, nil
}

type emptyRows struct{}

func (*emptyRows) Scan(dest ...interface{}) error {
	return nil
}

func (*emptyRows) Next() bool {
	return false
}
