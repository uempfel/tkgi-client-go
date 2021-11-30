// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ComputeProfileParameters ComputeProfileParameters
//
// swagger:model ComputeProfileParameters
type ComputeProfileParameters struct {

	// azs
	Azs []*AZ `json:"azs"`

	// cluster customization
	ClusterCustomization *ClusterCustomization `json:"cluster_customization,omitempty"`
}

// Validate validates this compute profile parameters
func (m *ComputeProfileParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAzs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterCustomization(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputeProfileParameters) validateAzs(formats strfmt.Registry) error {
	if swag.IsZero(m.Azs) { // not required
		return nil
	}

	for i := 0; i < len(m.Azs); i++ {
		if swag.IsZero(m.Azs[i]) { // not required
			continue
		}

		if m.Azs[i] != nil {
			if err := m.Azs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeProfileParameters) validateClusterCustomization(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterCustomization) { // not required
		return nil
	}

	if m.ClusterCustomization != nil {
		if err := m.ClusterCustomization.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_customization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_customization")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this compute profile parameters based on the context it is used
func (m *ComputeProfileParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAzs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterCustomization(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ComputeProfileParameters) contextValidateAzs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Azs); i++ {

		if m.Azs[i] != nil {
			if err := m.Azs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("azs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("azs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ComputeProfileParameters) contextValidateClusterCustomization(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterCustomization != nil {
		if err := m.ClusterCustomization.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_customization")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_customization")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ComputeProfileParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ComputeProfileParameters) UnmarshalBinary(b []byte) error {
	var res ComputeProfileParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}