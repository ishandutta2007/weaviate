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

package keys

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
)

// NewWeaviateKeysMeGetParams creates a new WeaviateKeysMeGetParams object
// with the default values initialized.
func NewWeaviateKeysMeGetParams() *WeaviateKeysMeGetParams {

	return &WeaviateKeysMeGetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewWeaviateKeysMeGetParamsWithTimeout creates a new WeaviateKeysMeGetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewWeaviateKeysMeGetParamsWithTimeout(timeout time.Duration) *WeaviateKeysMeGetParams {

	return &WeaviateKeysMeGetParams{

		timeout: timeout,
	}
}

// NewWeaviateKeysMeGetParamsWithContext creates a new WeaviateKeysMeGetParams object
// with the default values initialized, and the ability to set a context for a request
func NewWeaviateKeysMeGetParamsWithContext(ctx context.Context) *WeaviateKeysMeGetParams {

	return &WeaviateKeysMeGetParams{

		Context: ctx,
	}
}

// NewWeaviateKeysMeGetParamsWithHTTPClient creates a new WeaviateKeysMeGetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewWeaviateKeysMeGetParamsWithHTTPClient(client *http.Client) *WeaviateKeysMeGetParams {

	return &WeaviateKeysMeGetParams{
		HTTPClient: client,
	}
}

/*WeaviateKeysMeGetParams contains all the parameters to send to the API endpoint
for the weaviate keys me get operation typically these are written to a http.Request
*/
type WeaviateKeysMeGetParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the weaviate keys me get params
func (o *WeaviateKeysMeGetParams) WithTimeout(timeout time.Duration) *WeaviateKeysMeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the weaviate keys me get params
func (o *WeaviateKeysMeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the weaviate keys me get params
func (o *WeaviateKeysMeGetParams) WithContext(ctx context.Context) *WeaviateKeysMeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the weaviate keys me get params
func (o *WeaviateKeysMeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the weaviate keys me get params
func (o *WeaviateKeysMeGetParams) WithHTTPClient(client *http.Client) *WeaviateKeysMeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the weaviate keys me get params
func (o *WeaviateKeysMeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *WeaviateKeysMeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
