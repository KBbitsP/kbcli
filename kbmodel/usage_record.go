// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UsageRecord usage record
// swagger:model UsageRecord
type UsageRecord struct {

	// amount
	Amount float64 `json:"amount,omitempty"`

	// record date
	// Format: date-time
	RecordDate strfmt.DateTime `json:"recordDate,omitempty"`
}

// Validate validates this usage record
func (m *UsageRecord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecordDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UsageRecord) validateRecordDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RecordDate) { // not required
		return nil
	}

	if err := validate.FormatOf("recordDate", "body", "date-time", m.RecordDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UsageRecord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UsageRecord) UnmarshalBinary(b []byte) error {
	var res UsageRecord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
