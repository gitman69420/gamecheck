package handlers

import (
	"gamecheck-backend/db"

	"log"

	"github.com/dimuska139/rawg-sdk-go/v3"
	"github.com/valkey-io/valkey-go"
)

type Handler struct {
	q      *db.Queries
	r      *rawg.Client
	c      valkey.Client
	logger *log.Logger
}

func NewHander(q *db.Queries, r *rawg.Client, c valkey.Client, l *log.Logger) *Handler {
	return &Handler{q: q, r: r, c: c, logger: l}
}
