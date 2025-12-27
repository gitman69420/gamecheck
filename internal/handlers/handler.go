package handlers

import "gamecheck-backend/db"

type Handler struct {
	q *db.Queries
}

func NewHander(q *db.Queries) *Handler {
	return &Handler{q: q}
}
