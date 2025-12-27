package utils

import (
	"context"
	"fmt"
	"gamecheck-backend/db"

	"github.com/jackc/pgx/v5"
)

var q *db.Queries

func InitSqlcQueries() (*db.Queries, error) {
	conn, err := pgx.Connect(context.Background(), "host=localhost port=5432 dbname=app user=app_user password=app_password")
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the db with error: %s", err)
	}
	q = db.New(conn)
	return q, nil
}
