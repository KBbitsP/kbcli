// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Bundle bundle
// swagger:model Bundle
type Bundle struct {

	// account Id
	// Required: true
	AccountID *strfmt.UUID `json:"accountId"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// bundle Id
	BundleID strfmt.UUID `json:"bundleId,omitempty"`

	// external key
	ExternalKey string `json:"externalKey,omitempty"`

	// subscriptions
	Subscriptions []*Subscription `json:"subscriptions"`

	// timeline
	Timeline *BundleTimeline `json:"timeline,omitempty"`
}

// Validate validates this bundle
func (m *Bundle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBundleID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubscriptions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTimeline(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Bundle) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountId", "body", m.AccountID); err != nil {
		return err
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Bundle) validateAuditLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {

		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {

			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *Bundle) validateBundleID(formats strfmt.Registry) error {

	if swag.IsZero(m.BundleID) { // not required
		return nil
	}

	if err := validate.FormatOf("bundleId", "body", "uuid", m.BundleID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Bundle) validateSubscriptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Subscriptions) { // not required
		return nil
	}

	for i := 0; i < len(m.Subscriptions); i++ {

		if swag.IsZero(m.Subscriptions[i]) { // not required
			continue
		}

		if m.Subscriptions[i] != nil {

			if err := m.Subscriptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subscriptions" + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

func (m *Bundle) validateTimeline(formats strfmt.Registry) error {

	if swag.IsZero(m.Timeline) { // not required
		return nil
	}

	if m.Timeline != nil {

		if err := m.Timeline.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeline")
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Bundle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Bundle) UnmarshalBinary(b []byte) error {
	var res Bundle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
