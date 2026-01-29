package config

import (
	"context"
	"fmt"
	"gamecheck-backend/db"

	"github.com/jackc/pgx/v5"
)

var q *db.Queries

func InitSqlcQueries(dbConfig *DbConfig) (*db.Queries, error) {
	connString := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s",
		dbConfig.DbHost,
		dbConfig.DbPort,
		dbConfig.DbName,
		dbConfig.DbUser,
		dbConfig.DbPassword,
	)
	conn, err := pgx.Connect(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to the db with error: %s", err)
	}
	q = db.New(conn)
	return q, nil
}
