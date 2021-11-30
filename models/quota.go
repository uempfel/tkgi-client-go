// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Quota Quota
//
// swagger:model Quota
type Quota struct {

	// limit
	// Required: true
	Limit *QuotaLimit `json:"limit"`

	// owner
	// Required: true
	Owner *string `json:"owner"`
}

// Validate validates this quota
func (m *Quota) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLimit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Quota) validateLimit(formats strfmt.Registry) error {

	if err := validate.Required("limit", "body", m.Limit); err != nil {
		return err
	}

	if m.Limit != nil {
		if err := m.Limit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("limit")
			}
			return err
		}
	}

	return nil
}

func (m *Quota) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this quota based on the context it is used
func (m *Quota) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLimit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Quota) contextValidateLimit(ctx context.Context, formats strfmt.Registry) error {

	if m.Limit != nil {
		if err := m.Limit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("limit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("limit")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Quota) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Quota) UnmarshalBinary(b []byte) error {
	var res Quota
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}