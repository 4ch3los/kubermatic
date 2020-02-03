// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ClusterSpec ClusterSpec defines the cluster specification
// swagger:model ClusterSpec
type ClusterSpec struct {

	// MachineNetworks optionally specifies the parameters for IPAM.
	MachineNetworks []*MachineNetworkingConfig `json:"machineNetworks"`

	// If active the PodNodeSelector admission plugin is configured at the apiserver
	UsePodNodeSelectorAdmissionPlugin bool `json:"usePodNodeSelectorAdmissionPlugin,omitempty"`

	// If active the PodSecurityPolicy admission plugin is configured at the apiserver
	UsePodSecurityPolicyAdmissionPlugin bool `json:"usePodSecurityPolicyAdmissionPlugin,omitempty"`

	// audit logging
	AuditLogging *AuditLoggingSettings `json:"auditLogging,omitempty"`

	// cloud
	Cloud *CloudSpec `json:"cloud,omitempty"`

	// oidc
	Oidc *OIDCSettings `json:"oidc,omitempty"`

	// openshift
	Openshift *Openshift `json:"openshift,omitempty"`

	// version
	Version Semver `json:"version,omitempty"`
}

// Validate validates this cluster spec
func (m *ClusterSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMachineNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogging(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloud(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOidc(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenshift(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterSpec) validateMachineNetworks(formats strfmt.Registry) error {

	if swag.IsZero(m.MachineNetworks) { // not required
		return nil
	}

	for i := 0; i < len(m.MachineNetworks); i++ {
		if swag.IsZero(m.MachineNetworks[i]) { // not required
			continue
		}

		if m.MachineNetworks[i] != nil {
			if err := m.MachineNetworks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("machineNetworks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterSpec) validateAuditLogging(formats strfmt.Registry) error {

	if swag.IsZero(m.AuditLogging) { // not required
		return nil
	}

	if m.AuditLogging != nil {
		if err := m.AuditLogging.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auditLogging")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateCloud(formats strfmt.Registry) error {

	if swag.IsZero(m.Cloud) { // not required
		return nil
	}

	if m.Cloud != nil {
		if err := m.Cloud.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloud")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateOidc(formats strfmt.Registry) error {

	if swag.IsZero(m.Oidc) { // not required
		return nil
	}

	if m.Oidc != nil {
		if err := m.Oidc.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("oidc")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterSpec) validateOpenshift(formats strfmt.Registry) error {

	if swag.IsZero(m.Openshift) { // not required
		return nil
	}

	if m.Openshift != nil {
		if err := m.Openshift.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openshift")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterSpec) UnmarshalBinary(b []byte) error {
	var res ClusterSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
