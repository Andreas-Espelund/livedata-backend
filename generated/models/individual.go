// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Individual individual
//
// swagger:model Individual
type Individual struct {

	// birth date
	// Format: date
	BirthDate strfmt.Date `json:"birth_date,omitempty"`

	// father
	Father int64 `json:"father,omitempty"`

	// gender
	// Enum: [MALE FEMALE]
	Gender string `json:"gender,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// mother
	Mother int64 `json:"mother,omitempty"`

	// status
	// Enum: [ACTIVE INACTIVE]
	Status string `json:"status,omitempty"`
}

// Validate validates this individual
func (m *Individual) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBirthDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Individual) validateBirthDate(formats strfmt.Registry) error {
	if swag.IsZero(m.BirthDate) { // not required
		return nil
	}

	if err := validate.FormatOf("birth_date", "body", "date", m.BirthDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var individualTypeGenderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MALE","FEMALE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		individualTypeGenderPropEnum = append(individualTypeGenderPropEnum, v)
	}
}

const (

	// IndividualGenderMALE captures enum value "MALE"
	IndividualGenderMALE string = "MALE"

	// IndividualGenderFEMALE captures enum value "FEMALE"
	IndividualGenderFEMALE string = "FEMALE"
)

// prop value enum
func (m *Individual) validateGenderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, individualTypeGenderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Individual) validateGender(formats strfmt.Registry) error {
	if swag.IsZero(m.Gender) { // not required
		return nil
	}

	// value enum
	if err := m.validateGenderEnum("gender", "body", m.Gender); err != nil {
		return err
	}

	return nil
}

var individualTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","INACTIVE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		individualTypeStatusPropEnum = append(individualTypeStatusPropEnum, v)
	}
}

const (

	// IndividualStatusACTIVE captures enum value "ACTIVE"
	IndividualStatusACTIVE string = "ACTIVE"

	// IndividualStatusINACTIVE captures enum value "INACTIVE"
	IndividualStatusINACTIVE string = "INACTIVE"
)

// prop value enum
func (m *Individual) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, individualTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Individual) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this individual based on context it is used
func (m *Individual) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Individual) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Individual) UnmarshalBinary(b []byte) error {
	var res Individual
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
