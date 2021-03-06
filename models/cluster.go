// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Cluster Cluster
//
// swagger:model Cluster
type Cluster struct {

	// available upgrades
	AvailableUpgrades []*AvailableClusterUpgrade `json:"available_upgrades"`

	// compute profile name
	ComputeProfileName string `json:"compute_profile_name,omitempty"`

	// k8s version
	K8sVersion string `json:"k8s_version,omitempty"`

	// kubernetes master ips
	KubernetesMasterIps []string `json:"kubernetes_master_ips"`

	// kubernetes profile name
	KubernetesProfileName string `json:"kubernetes_profile_name,omitempty"`

	// last action
	// Enum: [CREATE UPDATE DELETE UPGRADE ROTATE_NSX_CERTIFICATES ADD_CUSTOM_CERTIFICATE_AUTHORITY SIGN_WITH_CUSTOM_CERTIFICATE_AUTHORITY PURGE_OLD_CUSTOM_CERTIFICATE_AUTHORITY]
	LastAction string `json:"last_action,omitempty"`

	// last action description
	LastActionDescription string `json:"last_action_description,omitempty"`

	// last action state
	// Enum: [in progress succeeded failed queued]
	LastActionState string `json:"last_action_state,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// network profile name
	NetworkProfileName string `json:"network_profile_name,omitempty"`

	// parameters
	Parameters *ClusterParameters `json:"parameters,omitempty"`

	// pks version
	PksVersion string `json:"pks_version,omitempty"`

	// plan name
	// Required: true
	PlanName *string `json:"plan_name"`

	// uuid
	// Required: true
	UUID *string `json:"uuid"`
}

// Validate validates this cluster
func (m *Cluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAvailableUpgrades(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastActionState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) validateAvailableUpgrades(formats strfmt.Registry) error {
	if swag.IsZero(m.AvailableUpgrades) { // not required
		return nil
	}

	for i := 0; i < len(m.AvailableUpgrades); i++ {
		if swag.IsZero(m.AvailableUpgrades[i]) { // not required
			continue
		}

		if m.AvailableUpgrades[i] != nil {
			if err := m.AvailableUpgrades[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("available_upgrades" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("available_upgrades" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var clusterTypeLastActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREATE","UPDATE","DELETE","UPGRADE","ROTATE_NSX_CERTIFICATES","ADD_CUSTOM_CERTIFICATE_AUTHORITY","SIGN_WITH_CUSTOM_CERTIFICATE_AUTHORITY","PURGE_OLD_CUSTOM_CERTIFICATE_AUTHORITY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTypeLastActionPropEnum = append(clusterTypeLastActionPropEnum, v)
	}
}

const (

	// ClusterLastActionCREATE captures enum value "CREATE"
	ClusterLastActionCREATE string = "CREATE"

	// ClusterLastActionUPDATE captures enum value "UPDATE"
	ClusterLastActionUPDATE string = "UPDATE"

	// ClusterLastActionDELETE captures enum value "DELETE"
	ClusterLastActionDELETE string = "DELETE"

	// ClusterLastActionUPGRADE captures enum value "UPGRADE"
	ClusterLastActionUPGRADE string = "UPGRADE"

	// ClusterLastActionROTATENSXCERTIFICATES captures enum value "ROTATE_NSX_CERTIFICATES"
	ClusterLastActionROTATENSXCERTIFICATES string = "ROTATE_NSX_CERTIFICATES"

	// ClusterLastActionADDCUSTOMCERTIFICATEAUTHORITY captures enum value "ADD_CUSTOM_CERTIFICATE_AUTHORITY"
	ClusterLastActionADDCUSTOMCERTIFICATEAUTHORITY string = "ADD_CUSTOM_CERTIFICATE_AUTHORITY"

	// ClusterLastActionSIGNWITHCUSTOMCERTIFICATEAUTHORITY captures enum value "SIGN_WITH_CUSTOM_CERTIFICATE_AUTHORITY"
	ClusterLastActionSIGNWITHCUSTOMCERTIFICATEAUTHORITY string = "SIGN_WITH_CUSTOM_CERTIFICATE_AUTHORITY"

	// ClusterLastActionPURGEOLDCUSTOMCERTIFICATEAUTHORITY captures enum value "PURGE_OLD_CUSTOM_CERTIFICATE_AUTHORITY"
	ClusterLastActionPURGEOLDCUSTOMCERTIFICATEAUTHORITY string = "PURGE_OLD_CUSTOM_CERTIFICATE_AUTHORITY"
)

// prop value enum
func (m *Cluster) validateLastActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterTypeLastActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Cluster) validateLastAction(formats strfmt.Registry) error {
	if swag.IsZero(m.LastAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateLastActionEnum("last_action", "body", m.LastAction); err != nil {
		return err
	}

	return nil
}

var clusterTypeLastActionStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["in progress","succeeded","failed","queued"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterTypeLastActionStatePropEnum = append(clusterTypeLastActionStatePropEnum, v)
	}
}

const (

	// ClusterLastActionStateInProgress captures enum value "in progress"
	ClusterLastActionStateInProgress string = "in progress"

	// ClusterLastActionStateSucceeded captures enum value "succeeded"
	ClusterLastActionStateSucceeded string = "succeeded"

	// ClusterLastActionStateFailed captures enum value "failed"
	ClusterLastActionStateFailed string = "failed"

	// ClusterLastActionStateQueued captures enum value "queued"
	ClusterLastActionStateQueued string = "queued"
)

// prop value enum
func (m *Cluster) validateLastActionStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, clusterTypeLastActionStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Cluster) validateLastActionState(formats strfmt.Registry) error {
	if swag.IsZero(m.LastActionState) { // not required
		return nil
	}

	// value enum
	if err := m.validateLastActionStateEnum("last_action_state", "body", m.LastActionState); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if m.Parameters != nil {
		if err := m.Parameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

func (m *Cluster) validatePlanName(formats strfmt.Registry) error {

	if err := validate.Required("plan_name", "body", m.PlanName); err != nil {
		return err
	}

	return nil
}

func (m *Cluster) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this cluster based on the context it is used
func (m *Cluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAvailableUpgrades(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cluster) contextValidateAvailableUpgrades(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AvailableUpgrades); i++ {

		if m.AvailableUpgrades[i] != nil {
			if err := m.AvailableUpgrades[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("available_upgrades" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("available_upgrades" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Cluster) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.Parameters != nil {
		if err := m.Parameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("parameters")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cluster) UnmarshalBinary(b []byte) error {
	var res Cluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
