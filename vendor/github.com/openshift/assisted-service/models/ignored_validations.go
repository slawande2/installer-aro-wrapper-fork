// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IgnoredValidations ignored validations
//
// swagger:model ignored-validations
type IgnoredValidations struct {

	// JSON-formatted list of cluster validation IDs that will be ignored for all hosts that belong to this cluster. It may also contain a list with a single string "all" to ignore all cluster validations. Some validations cannot be ignored.
	ClusterValidationIds string `json:"cluster-validation-ids,omitempty"`

	// JSON-formatted list of host validation IDs that will be ignored for all hosts that belong to this cluster. It may also contain a list with a single string "all" to ignore all host validations. Some validations cannot be ignored.
	HostValidationIds string `json:"host-validation-ids,omitempty"`
}

// Validate validates this ignored validations
func (m *IgnoredValidations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this ignored validations based on context it is used
func (m *IgnoredValidations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IgnoredValidations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgnoredValidations) UnmarshalBinary(b []byte) error {
	var res IgnoredValidations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}