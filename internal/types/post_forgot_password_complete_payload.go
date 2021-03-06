// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostForgotPasswordCompletePayload post forgot password complete payload
//
// swagger:model postForgotPasswordCompletePayload
type PostForgotPasswordCompletePayload struct {

	// New password to set for user
	// Required: true
	// Max Length: 500
	// Min Length: 1
	Password *string `json:"password"`

	// Password reset token sent via email
	// Required: true
	// Format: uuid4
	Token *strfmt.UUID4 `json:"token"`
}

// Validate validates this post forgot password complete payload
func (m *PostForgotPasswordCompletePayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePassword(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostForgotPasswordCompletePayload) validatePassword(formats strfmt.Registry) error {

	if err := validate.Required("password", "body", m.Password); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(*m.Password), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(*m.Password), 500); err != nil {
		return err
	}

	return nil
}

func (m *PostForgotPasswordCompletePayload) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	if err := validate.FormatOf("token", "body", "uuid4", m.Token.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostForgotPasswordCompletePayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostForgotPasswordCompletePayload) UnmarshalBinary(b []byte) error {
	var res PostForgotPasswordCompletePayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
