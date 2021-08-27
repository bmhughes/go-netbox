// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewIpamServicesDeleteParams creates a new IpamServicesDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServicesDeleteParams() *IpamServicesDeleteParams {
	return &IpamServicesDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServicesDeleteParamsWithTimeout creates a new IpamServicesDeleteParams object
// with the ability to set a timeout on a request.
func NewIpamServicesDeleteParamsWithTimeout(timeout time.Duration) *IpamServicesDeleteParams {
	return &IpamServicesDeleteParams{
		timeout: timeout,
	}
}

// NewIpamServicesDeleteParamsWithContext creates a new IpamServicesDeleteParams object
// with the ability to set a context for a request.
func NewIpamServicesDeleteParamsWithContext(ctx context.Context) *IpamServicesDeleteParams {
	return &IpamServicesDeleteParams{
		Context: ctx,
	}
}

// NewIpamServicesDeleteParamsWithHTTPClient creates a new IpamServicesDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServicesDeleteParamsWithHTTPClient(client *http.Client) *IpamServicesDeleteParams {
	return &IpamServicesDeleteParams{
		HTTPClient: client,
	}
}

/* IpamServicesDeleteParams contains all the parameters to send to the API endpoint
   for the ipam services delete operation.

   Typically these are written to a http.Request.
*/
type IpamServicesDeleteParams struct {

	/* ID.

	   A unique integer value identifying this service.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam services delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesDeleteParams) WithDefaults() *IpamServicesDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam services delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServicesDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam services delete params
func (o *IpamServicesDeleteParams) WithTimeout(timeout time.Duration) *IpamServicesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam services delete params
func (o *IpamServicesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam services delete params
func (o *IpamServicesDeleteParams) WithContext(ctx context.Context) *IpamServicesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam services delete params
func (o *IpamServicesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam services delete params
func (o *IpamServicesDeleteParams) WithHTTPClient(client *http.Client) *IpamServicesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam services delete params
func (o *IpamServicesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam services delete params
func (o *IpamServicesDeleteParams) WithID(id int64) *IpamServicesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam services delete params
func (o *IpamServicesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServicesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
