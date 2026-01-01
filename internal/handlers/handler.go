package handlers

import (
	"gamecheck-backend/db"

	"github.com/dimuska139/rawg-sdk-go/v3"
)

type Handler struct {
	q *db.Queries
	r *rawg.Client
}

func NewHander(q *db.Queries, r *rawg.Client) *Handler {
	return &Handler{q: q, r: r}
}
