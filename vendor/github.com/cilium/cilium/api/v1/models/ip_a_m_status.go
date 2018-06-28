// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IPAMStatus Status of IP address management
// swagger:model IPAMStatus

type IPAMStatus struct {

	// ipv4
	IPV4 []string `json:"ipv4"`

	// ipv6
	IPV6 []string `json:"ipv6"`
}

/* polymorph IPAMStatus ipv4 false */

/* polymorph IPAMStatus ipv6 false */

// Validate validates this IP a m status
func (m *IPAMStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPV4(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIPV6(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IPAMStatus) validateIPV4(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV4) { // not required
		return nil
	}

	return nil
}

func (m *IPAMStatus) validateIPV6(formats strfmt.Registry) error {

	if swag.IsZero(m.IPV6) { // not required
		return nil
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IPAMStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IPAMStatus) UnmarshalBinary(b []byte) error {
	var res IPAMStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}