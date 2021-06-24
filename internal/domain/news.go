package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type News struct {
	ID        uuid.UUID `json:"id" form:"id"`
	Title     string    `json:"title" form:"title" validate:"required"`
	Text      string    `json:"text" form:"text" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (n *News) IsValid() bool {
	title := strings.Trim(n.Title, " ")
	text := strings.Trim(n.Text, " ")

	return title != "" && text != ""
}
