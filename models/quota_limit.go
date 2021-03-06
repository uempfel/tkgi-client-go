// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// QuotaLimit QuotaLimit
//
// swagger:model QuotaLimit
type QuotaLimit struct {

	// cluster
	Cluster int32 `json:"cluster,omitempty"`

	// cpu
	CPU int64 `json:"cpu,omitempty"`

	// memory
	Memory float64 `json:"memory,omitempty"`
}

// Validate validates this quota limit
func (m *QuotaLimit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this quota limit based on context it is used
func (m *QuotaLimit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *QuotaLimit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaLimit) UnmarshalBinary(b []byte) error {
	var res QuotaLimit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
