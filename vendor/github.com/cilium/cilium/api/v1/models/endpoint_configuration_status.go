// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EndpointConfigurationStatus An endpoint's configuration
//
// swagger:model EndpointConfigurationStatus
type EndpointConfigurationStatus struct {

	// Most recent error, if applicable
	Error Error `json:"error,omitempty"`

	// Immutable configuration (read-only)
	Immutable ConfigurationMap `json:"immutable,omitempty"`

	// currently applied changeable configuration
	Realized *EndpointConfigurationSpec `json:"realized,omitempty"`
}

// Validate validates this endpoint configuration status
func (m *EndpointConfigurationStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImmutable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRealized(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointConfigurationStatus) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if err := m.Error.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("error")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("error")
		}
		return err
	}

	return nil
}

func (m *EndpointConfigurationStatus) validateImmutable(formats strfmt.Registry) error {
	if swag.IsZero(m.Immutable) { // not required
		return nil
	}

	if m.Immutable != nil {
		if err := m.Immutable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("immutable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("immutable")
			}
			return err
		}
	}

	return nil
}

func (m *EndpointConfigurationStatus) validateRealized(formats strfmt.Registry) error {
	if swag.IsZero(m.Realized) { // not required
		return nil
	}

	if m.Realized != nil {
		if err := m.Realized.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realized")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realized")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this endpoint configuration status based on the context it is used
func (m *EndpointConfigurationStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImmutable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRealized(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EndpointConfigurationStatus) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if err := m.Error.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("error")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("error")
		}
		return err
	}

	return nil
}

func (m *EndpointConfigurationStatus) contextValidateImmutable(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Immutable) { // not required
		return nil
	}

	if err := m.Immutable.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("immutable")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("immutable")
		}
		return err
	}

	return nil
}

func (m *EndpointConfigurationStatus) contextValidateRealized(ctx context.Context, formats strfmt.Registry) error {

	if m.Realized != nil {

		if swag.IsZero(m.Realized) { // not required
			return nil
		}

		if err := m.Realized.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("realized")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("realized")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointConfigurationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointConfigurationStatus) UnmarshalBinary(b []byte) error {
	var res EndpointConfigurationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
