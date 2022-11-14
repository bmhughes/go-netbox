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

// NewIpamL2vpnsReadParams creates a new IpamL2vpnsReadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamL2vpnsReadParams() *IpamL2vpnsReadParams {
	return &IpamL2vpnsReadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamL2vpnsReadParamsWithTimeout creates a new IpamL2vpnsReadParams object
// with the ability to set a timeout on a request.
func NewIpamL2vpnsReadParamsWithTimeout(timeout time.Duration) *IpamL2vpnsReadParams {
	return &IpamL2vpnsReadParams{
		timeout: timeout,
	}
}

// NewIpamL2vpnsReadParamsWithContext creates a new IpamL2vpnsReadParams object
// with the ability to set a context for a request.
func NewIpamL2vpnsReadParamsWithContext(ctx context.Context) *IpamL2vpnsReadParams {
	return &IpamL2vpnsReadParams{
		Context: ctx,
	}
}

// NewIpamL2vpnsReadParamsWithHTTPClient creates a new IpamL2vpnsReadParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamL2vpnsReadParamsWithHTTPClient(client *http.Client) *IpamL2vpnsReadParams {
	return &IpamL2vpnsReadParams{
		HTTPClient: client,
	}
}

/*
IpamL2vpnsReadParams contains all the parameters to send to the API endpoint

	for the ipam l2vpns read operation.

	Typically these are written to a http.Request.
*/
type IpamL2vpnsReadParams struct {

	/* ID.

	   A unique integer value identifying this L2VPN.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam l2vpns read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnsReadParams) WithDefaults() *IpamL2vpnsReadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam l2vpns read params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnsReadParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam l2vpns read params
func (o *IpamL2vpnsReadParams) WithTimeout(timeout time.Duration) *IpamL2vpnsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam l2vpns read params
func (o *IpamL2vpnsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam l2vpns read params
func (o *IpamL2vpnsReadParams) WithContext(ctx context.Context) *IpamL2vpnsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam l2vpns read params
func (o *IpamL2vpnsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam l2vpns read params
func (o *IpamL2vpnsReadParams) WithHTTPClient(client *http.Client) *IpamL2vpnsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam l2vpns read params
func (o *IpamL2vpnsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the ipam l2vpns read params
func (o *IpamL2vpnsReadParams) WithID(id int64) *IpamL2vpnsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam l2vpns read params
func (o *IpamL2vpnsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamL2vpnsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
