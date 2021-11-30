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

// ClusterUsage ClusterUsage
//
// swagger:model ClusterUsage
type ClusterUsage struct {

	// cpu
	// Required: true
	CPU *int64 `json:"cpu"`

	// memory
	// Required: true
	Memory *float64 `json:"memory"`

	// name
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this cluster usage
func (m *ClusterUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCPU(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterUsage) validateCPU(formats strfmt.Registry) error {

	if err := validate.Required("cpu", "body", m.CPU); err != nil {
		return err
	}

	return nil
}

func (m *ClusterUsage) validateMemory(formats strfmt.Registry) error {

	if err := validate.Required("memory", "body", m.Memory); err != nil {
		return err
	}

	return nil
}

func (m *ClusterUsage) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cluster usage based on context it is used
func (m *ClusterUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ClusterUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterUsage) UnmarshalBinary(b []byte) error {
	var res ClusterUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
