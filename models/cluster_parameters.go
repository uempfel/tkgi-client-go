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
	"github.com/go-openapi/validate"
)

// ClusterParameters ClusterParameters
//
// swagger:model ClusterParameters
type ClusterParameters struct {

	// authorization mode
	AuthorizationMode string `json:"authorization_mode,omitempty"`

	// cluster config parameters
	ClusterConfigParameters *ClusterConfigParameters `json:"cluster_config_parameters,omitempty"`

	// cluster tags
	ClusterTags []*ClusterTag `json:"cluster_tags"`

	// compute profile
	ComputeProfile string `json:"compute_profile,omitempty"`

	// custom ca certs
	CustomCaCerts []*CustomCaCert `json:"custom_ca_certs"`

	// k8s customization parameters
	K8sCustomizationParameters *KubernetesCustomizationParameters `json:"k8s_customization_parameters,omitempty"`

	// kubernetes master host
	// Required: true
	KubernetesMasterHost *string `json:"kubernetes_master_host"`

	// kubernetes master port
	// Example: 8443
	KubernetesMasterPort int32 `json:"kubernetes_master_port,omitempty"`

	// kubernetes setting cluster details
	KubernetesSettingClusterDetails *KubernetesSettings `json:"kubernetes_setting_cluster_details,omitempty"`

	// kubernetes setting plan details
	KubernetesSettingPlanDetails *KubernetesSettings `json:"kubernetes_setting_plan_details,omitempty"`

	// kubernetes worker instances
	KubernetesWorkerInstances int32 `json:"kubernetes_worker_instances,omitempty"`

	// node pool instances
	NodePoolInstances string `json:"node_pool_instances,omitempty"`

	// nsxt network details
	NsxtNetworkDetails string `json:"nsxt_network_details,omitempty"`

	// nsxt network profile
	NsxtNetworkProfile string `json:"nsxt_network_profile,omitempty"`

	// worker haproxy ip addresses
	WorkerHaproxyIPAddresses string `json:"worker_haproxy_ip_addresses,omitempty"`
}

// Validate validates this cluster parameters
func (m *ClusterParameters) Validate(formats strfmt.Registry) error {
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

	if err := m.validateK8sCustomizationParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesMasterHost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesSettingClusterDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetesSettingPlanDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterParameters) validateClusterConfigParameters(formats strfmt.Registry) error {
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

func (m *ClusterParameters) validateClusterTags(formats strfmt.Registry) error {
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

func (m *ClusterParameters) validateCustomCaCerts(formats strfmt.Registry) error {
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

func (m *ClusterParameters) validateK8sCustomizationParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.K8sCustomizationParameters) { // not required
		return nil
	}

	if m.K8sCustomizationParameters != nil {
		if err := m.K8sCustomizationParameters.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("k8s_customization_parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("k8s_customization_parameters")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterParameters) validateKubernetesMasterHost(formats strfmt.Registry) error {

	if err := validate.Required("kubernetes_master_host", "body", m.KubernetesMasterHost); err != nil {
		return err
	}

	return nil
}

func (m *ClusterParameters) validateKubernetesSettingClusterDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.KubernetesSettingClusterDetails) { // not required
		return nil
	}

	if m.KubernetesSettingClusterDetails != nil {
		if err := m.KubernetesSettingClusterDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes_setting_cluster_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetes_setting_cluster_details")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterParameters) validateKubernetesSettingPlanDetails(formats strfmt.Registry) error {
	if swag.IsZero(m.KubernetesSettingPlanDetails) { // not required
		return nil
	}

	if m.KubernetesSettingPlanDetails != nil {
		if err := m.KubernetesSettingPlanDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes_setting_plan_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetes_setting_plan_details")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cluster parameters based on the context it is used
func (m *ClusterParameters) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateK8sCustomizationParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubernetesSettingClusterDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateKubernetesSettingPlanDetails(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterParameters) contextValidateClusterConfigParameters(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterParameters) contextValidateClusterTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterParameters) contextValidateCustomCaCerts(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ClusterParameters) contextValidateK8sCustomizationParameters(ctx context.Context, formats strfmt.Registry) error {

	if m.K8sCustomizationParameters != nil {
		if err := m.K8sCustomizationParameters.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("k8s_customization_parameters")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("k8s_customization_parameters")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterParameters) contextValidateKubernetesSettingClusterDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.KubernetesSettingClusterDetails != nil {
		if err := m.KubernetesSettingClusterDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes_setting_cluster_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetes_setting_cluster_details")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterParameters) contextValidateKubernetesSettingPlanDetails(ctx context.Context, formats strfmt.Registry) error {

	if m.KubernetesSettingPlanDetails != nil {
		if err := m.KubernetesSettingPlanDetails.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubernetes_setting_plan_details")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("kubernetes_setting_plan_details")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterParameters) UnmarshalBinary(b []byte) error {
	var res ClusterParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
