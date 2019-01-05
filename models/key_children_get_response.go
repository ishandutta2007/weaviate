/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN: Bob van Luijt (bob@k10y.co)
 */
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// KeyChildrenGetResponse key children get response
// swagger:model KeyChildrenGetResponse
type KeyChildrenGetResponse struct {

	// children
	Children MultipleRef `json:"children,omitempty"`
}

// Validate validates this key children get response
func (m *KeyChildrenGetResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChildren(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyChildrenGetResponse) validateChildren(formats strfmt.Registry) error {

	if swag.IsZero(m.Children) { // not required
		return nil
	}

	if err := m.Children.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("children")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KeyChildrenGetResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyChildrenGetResponse) UnmarshalBinary(b []byte) error {
	var res KeyChildrenGetResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
