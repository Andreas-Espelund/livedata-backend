package models

import "github.com/go-openapi/strfmt"

type IndividualPatch struct {
	BirthDate *strfmt.Date `json:"birth_date,omitempty"`
	Status    *string      `json:"status,omitempty"`
	Gender    *string      `json:"gender,omitempty"`
	Mother    *int64       `json:"mother,omitempty"`
	Father    *int64       `json:"father,omitempty"`
}
