// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// LabelsResponse labels response
// swagger:model LabelsResponse
type LabelsResponse struct {

	// the ip of node
	IPAddress string `json:"ipAddress,omitempty"`

	// the labels of node
	Labels map[string]string `json:"labels,omitempty"`
}

// Validate validates this labels response
func (m *LabelsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LabelsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LabelsResponse) UnmarshalBinary(b []byte) error {
	var res LabelsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
