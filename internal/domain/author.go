package domain

import (
	"github.com/google/uuid"
	"strings"
)

type Author struct {
	ID         uuid.UUID `json:"id" form:"id"`
	FirstName  string    `json:"firstName" form:"firstName" validate:"required"`
	LastName   string    `json:"lastName" form:"lastName" validate:"required"`
	FatherName string    `json:"fatherName" form:"fatherName"`
}

func (a *Author) IsValid() bool {
	firstName := strings.Trim(a.FirstName, " ")
	lastName := strings.Trim(a.LastName, " ")

	return firstName != "" && lastName != ""
}
