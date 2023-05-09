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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyTenantGroupsBulkDeleteReader is a Reader for the TenancyTenantGroupsBulkDelete structure.
type TenancyTenantGroupsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantGroupsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyTenantGroupsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTenancyTenantGroupsBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTenancyTenantGroupsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantGroupsBulkDeleteNoContent creates a TenancyTenantGroupsBulkDeleteNoContent with default headers values
func NewTenancyTenantGroupsBulkDeleteNoContent() *TenancyTenantGroupsBulkDeleteNoContent {
	return &TenancyTenantGroupsBulkDeleteNoContent{}
}

/*
TenancyTenantGroupsBulkDeleteNoContent describes a response with status code 204, with default header values.

TenancyTenantGroupsBulkDeleteNoContent tenancy tenant groups bulk delete no content
*/
type TenancyTenantGroupsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this tenancy tenant groups bulk delete no content response has a 2xx status code
func (o *TenancyTenantGroupsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenant groups bulk delete no content response has a 3xx status code
func (o *TenancyTenantGroupsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenant groups bulk delete no content response has a 4xx status code
func (o *TenancyTenantGroupsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenant groups bulk delete no content response has a 5xx status code
func (o *TenancyTenantGroupsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenant groups bulk delete no content response a status code equal to that given
func (o *TenancyTenantGroupsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *TenancyTenantGroupsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkDeleteNoContent ", 204)
}

func (o *TenancyTenantGroupsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkDeleteNoContent ", 204)
}

func (o *TenancyTenantGroupsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenancyTenantGroupsBulkDeleteBadRequest creates a TenancyTenantGroupsBulkDeleteBadRequest with default headers values
func NewTenancyTenantGroupsBulkDeleteBadRequest() *TenancyTenantGroupsBulkDeleteBadRequest {
	return &TenancyTenantGroupsBulkDeleteBadRequest{}
}

/*
TenancyTenantGroupsBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TenancyTenantGroupsBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this tenancy tenant groups bulk delete bad request response has a 2xx status code
func (o *TenancyTenantGroupsBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenancy tenant groups bulk delete bad request response has a 3xx status code
func (o *TenancyTenantGroupsBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenant groups bulk delete bad request response has a 4xx status code
func (o *TenancyTenantGroupsBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenancy tenant groups bulk delete bad request response has a 5xx status code
func (o *TenancyTenantGroupsBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenant groups bulk delete bad request response a status code equal to that given
func (o *TenancyTenantGroupsBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TenancyTenantGroupsBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyTenantGroupsBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/][%d] tenancyTenantGroupsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyTenantGroupsBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantGroupsBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantGroupsBulkDeleteDefault creates a TenancyTenantGroupsBulkDeleteDefault with default headers values
func NewTenancyTenantGroupsBulkDeleteDefault(code int) *TenancyTenantGroupsBulkDeleteDefault {
	return &TenancyTenantGroupsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
TenancyTenantGroupsBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type TenancyTenantGroupsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy tenant groups bulk delete default response
func (o *TenancyTenantGroupsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy tenant groups bulk delete default response has a 2xx status code
func (o *TenancyTenantGroupsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy tenant groups bulk delete default response has a 3xx status code
func (o *TenancyTenantGroupsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy tenant groups bulk delete default response has a 4xx status code
func (o *TenancyTenantGroupsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy tenant groups bulk delete default response has a 5xx status code
func (o *TenancyTenantGroupsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy tenant groups bulk delete default response a status code equal to that given
func (o *TenancyTenantGroupsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyTenantGroupsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/][%d] tenancy_tenant-groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantGroupsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /tenancy/tenant-groups/][%d] tenancy_tenant-groups_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantGroupsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantGroupsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
