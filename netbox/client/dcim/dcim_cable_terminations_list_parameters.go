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

package dcim

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

// NewDcimCableTerminationsListParams creates a new DcimCableTerminationsListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimCableTerminationsListParams() *DcimCableTerminationsListParams {
	return &DcimCableTerminationsListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimCableTerminationsListParamsWithTimeout creates a new DcimCableTerminationsListParams object
// with the ability to set a timeout on a request.
func NewDcimCableTerminationsListParamsWithTimeout(timeout time.Duration) *DcimCableTerminationsListParams {
	return &DcimCableTerminationsListParams{
		timeout: timeout,
	}
}

// NewDcimCableTerminationsListParamsWithContext creates a new DcimCableTerminationsListParams object
// with the ability to set a context for a request.
func NewDcimCableTerminationsListParamsWithContext(ctx context.Context) *DcimCableTerminationsListParams {
	return &DcimCableTerminationsListParams{
		Context: ctx,
	}
}

// NewDcimCableTerminationsListParamsWithHTTPClient creates a new DcimCableTerminationsListParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimCableTerminationsListParamsWithHTTPClient(client *http.Client) *DcimCableTerminationsListParams {
	return &DcimCableTerminationsListParams{
		HTTPClient: client,
	}
}

/*
DcimCableTerminationsListParams contains all the parameters to send to the API endpoint

	for the dcim cable terminations list operation.

	Typically these are written to a http.Request.
*/
type DcimCableTerminationsListParams struct {

	// Cable.
	Cable *string

	// Cablen.
	Cablen *string

	// CableEnd.
	CableEnd *string

	// CableEndn.
	CableEndn *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	/* Ordering.

	   Which field to use when ordering the results.
	*/
	Ordering *string

	// TerminationID.
	TerminationID *string

	// TerminationIDGt.
	TerminationIDGt *string

	// TerminationIDGte.
	TerminationIDGte *string

	// TerminationIDLt.
	TerminationIDLt *string

	// TerminationIDLte.
	TerminationIDLte *string

	// TerminationIDn.
	TerminationIDn *string

	// TerminationType.
	TerminationType *string

	// TerminationTypen.
	TerminationTypen *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim cable terminations list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCableTerminationsListParams) WithDefaults() *DcimCableTerminationsListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim cable terminations list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimCableTerminationsListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithTimeout(timeout time.Duration) *DcimCableTerminationsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithContext(ctx context.Context) *DcimCableTerminationsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithHTTPClient(client *http.Client) *DcimCableTerminationsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCable adds the cable to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithCable(cable *string) *DcimCableTerminationsListParams {
	o.SetCable(cable)
	return o
}

// SetCable adds the cable to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetCable(cable *string) {
	o.Cable = cable
}

// WithCablen adds the cablen to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithCablen(cablen *string) *DcimCableTerminationsListParams {
	o.SetCablen(cablen)
	return o
}

// SetCablen adds the cableN to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetCablen(cablen *string) {
	o.Cablen = cablen
}

// WithCableEnd adds the cableEnd to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithCableEnd(cableEnd *string) *DcimCableTerminationsListParams {
	o.SetCableEnd(cableEnd)
	return o
}

// SetCableEnd adds the cableEnd to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetCableEnd(cableEnd *string) {
	o.CableEnd = cableEnd
}

// WithCableEndn adds the cableEndn to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithCableEndn(cableEndn *string) *DcimCableTerminationsListParams {
	o.SetCableEndn(cableEndn)
	return o
}

// SetCableEndn adds the cableEndN to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetCableEndn(cableEndn *string) {
	o.CableEndn = cableEndn
}

// WithID adds the id to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithID(id *string) *DcimCableTerminationsListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithIDGt(iDGt *string) *DcimCableTerminationsListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithIDGte(iDGte *string) *DcimCableTerminationsListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithIDLt(iDLt *string) *DcimCableTerminationsListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithIDLte(iDLte *string) *DcimCableTerminationsListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithIDn(iDn *string) *DcimCableTerminationsListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLimit adds the limit to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithLimit(limit *int64) *DcimCableTerminationsListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithOffset(offset *int64) *DcimCableTerminationsListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithOrdering(ordering *string) *DcimCableTerminationsListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithTerminationID adds the terminationID to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithTerminationID(terminationID *string) *DcimCableTerminationsListParams {
	o.SetTerminationID(terminationID)
	return o
}

// SetTerminationID adds the terminationId to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetTerminationID(terminationID *string) {
	o.TerminationID = terminationID
}

// WithTerminationIDGt adds the terminationIDGt to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithTerminationIDGt(terminationIDGt *string) *DcimCableTerminationsListParams {
	o.SetTerminationIDGt(terminationIDGt)
	return o
}

// SetTerminationIDGt adds the terminationIdGt to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetTerminationIDGt(terminationIDGt *string) {
	o.TerminationIDGt = terminationIDGt
}

// WithTerminationIDGte adds the terminationIDGte to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithTerminationIDGte(terminationIDGte *string) *DcimCableTerminationsListParams {
	o.SetTerminationIDGte(terminationIDGte)
	return o
}

// SetTerminationIDGte adds the terminationIdGte to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetTerminationIDGte(terminationIDGte *string) {
	o.TerminationIDGte = terminationIDGte
}

// WithTerminationIDLt adds the terminationIDLt to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithTerminationIDLt(terminationIDLt *string) *DcimCableTerminationsListParams {
	o.SetTerminationIDLt(terminationIDLt)
	return o
}

// SetTerminationIDLt adds the terminationIdLt to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetTerminationIDLt(terminationIDLt *string) {
	o.TerminationIDLt = terminationIDLt
}

// WithTerminationIDLte adds the terminationIDLte to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithTerminationIDLte(terminationIDLte *string) *DcimCableTerminationsListParams {
	o.SetTerminationIDLte(terminationIDLte)
	return o
}

// SetTerminationIDLte adds the terminationIdLte to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetTerminationIDLte(terminationIDLte *string) {
	o.TerminationIDLte = terminationIDLte
}

// WithTerminationIDn adds the terminationIDn to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithTerminationIDn(terminationIDn *string) *DcimCableTerminationsListParams {
	o.SetTerminationIDn(terminationIDn)
	return o
}

// SetTerminationIDn adds the terminationIdN to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetTerminationIDn(terminationIDn *string) {
	o.TerminationIDn = terminationIDn
}

// WithTerminationType adds the terminationType to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithTerminationType(terminationType *string) *DcimCableTerminationsListParams {
	o.SetTerminationType(terminationType)
	return o
}

// SetTerminationType adds the terminationType to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetTerminationType(terminationType *string) {
	o.TerminationType = terminationType
}

// WithTerminationTypen adds the terminationTypen to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) WithTerminationTypen(terminationTypen *string) *DcimCableTerminationsListParams {
	o.SetTerminationTypen(terminationTypen)
	return o
}

// SetTerminationTypen adds the terminationTypeN to the dcim cable terminations list params
func (o *DcimCableTerminationsListParams) SetTerminationTypen(terminationTypen *string) {
	o.TerminationTypen = terminationTypen
}

// WriteToRequest writes these params to a swagger request
func (o *DcimCableTerminationsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Cable != nil {

		// query param cable
		var qrCable string

		if o.Cable != nil {
			qrCable = *o.Cable
		}
		qCable := qrCable
		if qCable != "" {

			if err := r.SetQueryParam("cable", qCable); err != nil {
				return err
			}
		}
	}

	if o.Cablen != nil {

		// query param cable__n
		var qrCablen string

		if o.Cablen != nil {
			qrCablen = *o.Cablen
		}
		qCablen := qrCablen
		if qCablen != "" {

			if err := r.SetQueryParam("cable__n", qCablen); err != nil {
				return err
			}
		}
	}

	if o.CableEnd != nil {

		// query param cable_end
		var qrCableEnd string

		if o.CableEnd != nil {
			qrCableEnd = *o.CableEnd
		}
		qCableEnd := qrCableEnd
		if qCableEnd != "" {

			if err := r.SetQueryParam("cable_end", qCableEnd); err != nil {
				return err
			}
		}
	}

	if o.CableEndn != nil {

		// query param cable_end__n
		var qrCableEndn string

		if o.CableEndn != nil {
			qrCableEndn = *o.CableEndn
		}
		qCableEndn := qrCableEndn
		if qCableEndn != "" {

			if err := r.SetQueryParam("cable_end__n", qCableEndn); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string

		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {

			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}
	}

	if o.TerminationID != nil {

		// query param termination_id
		var qrTerminationID string

		if o.TerminationID != nil {
			qrTerminationID = *o.TerminationID
		}
		qTerminationID := qrTerminationID
		if qTerminationID != "" {

			if err := r.SetQueryParam("termination_id", qTerminationID); err != nil {
				return err
			}
		}
	}

	if o.TerminationIDGt != nil {

		// query param termination_id__gt
		var qrTerminationIDGt string

		if o.TerminationIDGt != nil {
			qrTerminationIDGt = *o.TerminationIDGt
		}
		qTerminationIDGt := qrTerminationIDGt
		if qTerminationIDGt != "" {

			if err := r.SetQueryParam("termination_id__gt", qTerminationIDGt); err != nil {
				return err
			}
		}
	}

	if o.TerminationIDGte != nil {

		// query param termination_id__gte
		var qrTerminationIDGte string

		if o.TerminationIDGte != nil {
			qrTerminationIDGte = *o.TerminationIDGte
		}
		qTerminationIDGte := qrTerminationIDGte
		if qTerminationIDGte != "" {

			if err := r.SetQueryParam("termination_id__gte", qTerminationIDGte); err != nil {
				return err
			}
		}
	}

	if o.TerminationIDLt != nil {

		// query param termination_id__lt
		var qrTerminationIDLt string

		if o.TerminationIDLt != nil {
			qrTerminationIDLt = *o.TerminationIDLt
		}
		qTerminationIDLt := qrTerminationIDLt
		if qTerminationIDLt != "" {

			if err := r.SetQueryParam("termination_id__lt", qTerminationIDLt); err != nil {
				return err
			}
		}
	}

	if o.TerminationIDLte != nil {

		// query param termination_id__lte
		var qrTerminationIDLte string

		if o.TerminationIDLte != nil {
			qrTerminationIDLte = *o.TerminationIDLte
		}
		qTerminationIDLte := qrTerminationIDLte
		if qTerminationIDLte != "" {

			if err := r.SetQueryParam("termination_id__lte", qTerminationIDLte); err != nil {
				return err
			}
		}
	}

	if o.TerminationIDn != nil {

		// query param termination_id__n
		var qrTerminationIDn string

		if o.TerminationIDn != nil {
			qrTerminationIDn = *o.TerminationIDn
		}
		qTerminationIDn := qrTerminationIDn
		if qTerminationIDn != "" {

			if err := r.SetQueryParam("termination_id__n", qTerminationIDn); err != nil {
				return err
			}
		}
	}

	if o.TerminationType != nil {

		// query param termination_type
		var qrTerminationType string

		if o.TerminationType != nil {
			qrTerminationType = *o.TerminationType
		}
		qTerminationType := qrTerminationType
		if qTerminationType != "" {

			if err := r.SetQueryParam("termination_type", qTerminationType); err != nil {
				return err
			}
		}
	}

	if o.TerminationTypen != nil {

		// query param termination_type__n
		var qrTerminationTypen string

		if o.TerminationTypen != nil {
			qrTerminationTypen = *o.TerminationTypen
		}
		qTerminationTypen := qrTerminationTypen
		if qTerminationTypen != "" {

			if err := r.SetQueryParam("termination_type__n", qTerminationTypen); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
