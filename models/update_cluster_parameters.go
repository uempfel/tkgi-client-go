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

// UpdateClusterParameters UpdateClusterParameters
//
// swagger:model UpdateClusterParameters
type UpdateClusterParameters struct {

	// cluster config parameters
	ClusterConfigParameters *ClusterConfigParameters `json:"cluster_config_parameters,omitempty"`

	// cluster tags
	ClusterTags []*ClusterTag `json:"cluster_tags"`

	// compute profile name
	ComputeProfileName string `json:"compute_profile_name,omitempty"`

	// custom ca certs
	CustomCaCerts []*CustomCaCert `json:"custom_ca_certs"`

	// delete all tags
	DeleteAllTags bool `json:"delete_all_tags,omitempty"`

	// insecure registries
	InsecureRegistries []string `json:"insecure_registries"`

	// kubelet drain delete local data
	KubeletDrainDeleteLocalData string `json:"kubelet_drain_delete_local_data,omitempty"`

	// kubelet drain force
	KubeletDrainForce string `json:"kubelet_drain_force,omitempty"`

	// kubelet drain force node
	KubeletDrainForceNode string `json:"kubelet_drain_force_node,omitempty"`

	// kubelet drain grace period
	KubeletDrainGracePeriod string `json:"kubelet_drain_grace_period,omitempty"`

	// kubelet drain ignore daemonsets
	KubeletDrainIgnoreDaemonsets string `json:"kubelet_drain_ignore_daemonsets,omitempty"`

	// kubelet drain timeout
	KubeletDrainTimeout string `json:"kubelet_drain_timeout,omitempty"`

	// kubernetes profile name
	KubernetesProfileName string `json:"kubernetes_profile_name,omitempty"`

	// kubernetes worker instances
	KubernetesWorkerInstances int32 `json:"kubernetes_worker_instances,omitempty"`

	// network profile name
	NetworkProfileName string `json:"network_profile_name,omitempty"`

	// node pool instances
	NodePoolInstances string `json:"node_pool_instances,omitempty"`
}

// Validate validates this update cluster parameters
func (m *UpdateClusterParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterConfigParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClusterTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomCaCerts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateClusterParameters) validateClusterConfigParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterConfigParameters) { // not required
		return nil
	}

	if m.ClusterConfigParameters != nil {
		if err := m.ClusterConfigParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_config_parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_config_parameters")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateClusterParameters) validateClusterTags(formats strfmt.Registry) error {
	if swag.IsZero(m.ClusterTags) { // not required
		return nil
	}

	for i := 0; i < len(m.ClusterTags); i++ {
		if swag.IsZero(m.ClusterTags[i]) { // not required
			continue
		}

		if m.ClusterTags[i] != nil {
			if err := m.ClusterTags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateClusterParameters) validateCustomCaCerts(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomCaCerts) { // not required
		return nil
	}

	for i := 0; i < len(m.CustomCaCerts); i++ {
		if swag.IsZero(m.CustomCaCerts[i]) { // not required
			continue
		}

		if m.CustomCaCerts[i] != nil {
			if err := m.CustomCaCerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_ca_certs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_ca_certs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update cluster parameters based on the context it is used
func (m *UpdateClusterParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClusterConfigParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateClusterTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomCaCerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateClusterParameters) contextValidateClusterConfigParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.ClusterConfigParameters != nil {
		if err := m.ClusterConfigParameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster_config_parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cluster_config_parameters")
			}
			return err
		}
	}

	return nil
}

func (m *UpdateClusterParameters) contextValidateClusterTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ClusterTags); i++ {

		if m.ClusterTags[i] != nil {
			if err := m.ClusterTags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cluster_tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("cluster_tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *UpdateClusterParameters) contextValidateCustomCaCerts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CustomCaCerts); i++ {

		if m.CustomCaCerts[i] != nil {
			if err := m.CustomCaCerts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("custom_ca_certs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("custom_ca_certs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateClusterParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateClusterParameters) UnmarshalBinary(b []byte) error {
	var res UpdateClusterParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
