package handlers

import (
	"gamecheck-backend/db"
	"gamecheck-backend/utils"
)

type Handler struct {
	q *db.Queries
	c *utils.Config
}

func NewHander(q *db.Queries, c *utils.Config) *Handler {
	return &Handler{q: q, c: c}
}
