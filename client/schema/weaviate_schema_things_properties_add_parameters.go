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

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/creativesoftwarefdn/weaviate/models"
)

// NewWeaviateSchemaThingsPropertiesAddParams creates a new WeaviateSchemaThingsPropertiesAddParams object
// with the default values initialized.
func NewWeaviateSchemaThingsPropertiesAddParams() *WeaviateSchemaThingsPropertiesAddParams {
	var ()
	return &WeaviateSchemaThingsPropertiesAddParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateSchemaThingsPropertiesAddParamsWithTimeout creates a new WeaviateSchemaThingsPropertiesAddParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateSchemaThingsPropertiesAddParamsWithTimeout(timeout time.Duration) *WeaviateSchemaThingsPropertiesAddParams {
	var ()
	return &WeaviateSchemaThingsPropertiesAddParams{

		timeout: timeout,
	}
}

// NewWeaviateSchemaThingsPropertiesAddParamsWithContext creates a new WeaviateSchemaThingsPropertiesAddParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateSchemaThingsPropertiesAddParamsWithContext(ctx context.Context) *WeaviateSchemaThingsPropertiesAddParams {
	var ()
	return &WeaviateSchemaThingsPropertiesAddParams{

		Context: ctx,
	}
}

// NewWeaviateSchemaThingsPropertiesAddParamsWithHTTPClient creates a new WeaviateSchemaThingsPropertiesAddParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateSchemaThingsPropertiesAddParamsWithHTTPClient(client *http.Client) *WeaviateSchemaThingsPropertiesAddParams {
	var ()
	return &WeaviateSchemaThingsPropertiesAddParams{
		HTTPClient: client,
	}
}

/*WeaviateSchemaThingsPropertiesAddParams contains all the parameters to send to the API endpoint
for the weaviate schema things properties add operation typically these are written to a http.Request
*/
type WeaviateSchemaThingsPropertiesAddParams struct {

	/*Body*/
	Body *models.SemanticSchemaClassProperty
	/*ClassName*/
	ClassName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate schema things properties add params
func (o *WeaviateSchemaThingsPropertiesAddParams) WithTimeout(timeout time.Duration) *WeaviateSchemaThingsPropertiesAddParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate schema things properties add params
func (o *WeaviateSchemaThingsPropertiesAddParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate schema things properties add params
func (o *WeaviateSchemaThingsPropertiesAddParams) WithContext(ctx context.Context) *WeaviateSchemaThingsPropertiesAddParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate schema things properties add params
func (o *WeaviateSchemaThingsPropertiesAddParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate schema things properties add params
func (o *WeaviateSchemaThingsPropertiesAddParams) WithHTTPClient(client *http.Client) *WeaviateSchemaThingsPropertiesAddParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate schema things properties add params
func (o *WeaviateSchemaThingsPropertiesAddParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the weaviate schema things properties add params
func (o *WeaviateSchemaThingsPropertiesAddParams) WithBody(body *models.SemanticSchemaClassProperty) *WeaviateSchemaThingsPropertiesAddParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the weaviate schema things properties add params
func (o *WeaviateSchemaThingsPropertiesAddParams) SetBody(body *models.SemanticSchemaClassProperty) {
	o.Body = body
}

// WithClassName adds the className to the weaviate schema things properties add params
func (o *WeaviateSchemaThingsPropertiesAddParams) WithClassName(className string) *WeaviateSchemaThingsPropertiesAddParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the weaviate schema things properties add params
func (o *WeaviateSchemaThingsPropertiesAddParams) SetClassName(className string) {
	o.ClassName = className
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateSchemaThingsPropertiesAddParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
