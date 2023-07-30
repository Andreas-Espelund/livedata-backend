package models

import (
	"github.com/go-openapi/strfmt"
)

type Individual struct {
	ID        int64       `json:"id" db:"ID"`
	BirthDate strfmt.Date `json:"birth_date" db:"BirthDate"`
	Father    int64       `json:"father" db:"Father"`
	Gender    string      `json:"gender" db:"Gender"`
	Mother    int64       `json:"mother" db:"Mother"`
	Status    string      `json:"status" db:"Status"`
}
