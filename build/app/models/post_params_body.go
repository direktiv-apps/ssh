// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/direktiv/apps/go/pkg/apps"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostParamsBody post params body
//
// swagger:model postParamsBody
type PostParamsBody struct {

	// auth
	// Required: true
	Auth *PostParamsBodyAuth `json:"auth"`

	// Array of SSH commands.
	Commands []*PostParamsBodyCommandsItems `json:"commands"`

	// File to create before running commands.
	Files []apps.DirektivFile `json:"files"`

	// host
	// Required: true
	Host *PostParamsBodyHost `json:"host"`
}

// Validate validates this post params body
func (m *PostParamsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommands(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHost(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBody) validateAuth(formats strfmt.Registry) error {

	if err := validate.Required("auth", "body", m.Auth); err != nil {
		return err
	}

	if m.Auth != nil {
		if err := m.Auth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

func (m *PostParamsBody) validateCommands(formats strfmt.Registry) error {
	if swag.IsZero(m.Commands) { // not required
		return nil
	}

	for i := 0; i < len(m.Commands); i++ {
		if swag.IsZero(m.Commands[i]) { // not required
			continue
		}

		if m.Commands[i] != nil {
			if err := m.Commands[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostParamsBody) validateFiles(formats strfmt.Registry) error {
	if swag.IsZero(m.Files) { // not required
		return nil
	}

	for i := 0; i < len(m.Files); i++ {

		if err := m.Files[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("files" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("files" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PostParamsBody) validateHost(formats strfmt.Registry) error {

	if err := validate.Required("host", "body", m.Host); err != nil {
		return err
	}

	if m.Host != nil {
		if err := m.Host.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this post params body based on the context it is used
func (m *PostParamsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCommands(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFiles(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHost(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostParamsBody) contextValidateAuth(ctx context.Context, formats strfmt.Registry) error {

	if m.Auth != nil {
		if err := m.Auth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("auth")
			}
			return err
		}
	}

	return nil
}

func (m *PostParamsBody) contextValidateCommands(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Commands); i++ {

		if m.Commands[i] != nil {
			if err := m.Commands[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("commands" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("commands" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *PostParamsBody) contextValidateFiles(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Files); i++ {

		if err := m.Files[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("files" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("files" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *PostParamsBody) contextValidateHost(ctx context.Context, formats strfmt.Registry) error {

	if m.Host != nil {
		if err := m.Host.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("host")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("host")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostParamsBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostParamsBody) UnmarshalBinary(b []byte) error {
	var res PostParamsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
