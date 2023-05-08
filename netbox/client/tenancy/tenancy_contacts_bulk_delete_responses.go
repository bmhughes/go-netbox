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

// TenancyContactsBulkDeleteReader is a Reader for the TenancyContactsBulkDelete structure.
type TenancyContactsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyContactsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTenancyContactsBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyContactsBulkDeleteNoContent creates a TenancyContactsBulkDeleteNoContent with default headers values
func NewTenancyContactsBulkDeleteNoContent() *TenancyContactsBulkDeleteNoContent {
	return &TenancyContactsBulkDeleteNoContent{}
}

/*
TenancyContactsBulkDeleteNoContent describes a response with status code 204, with default header values.

TenancyContactsBulkDeleteNoContent tenancy contacts bulk delete no content
*/
type TenancyContactsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this tenancy contacts bulk delete no content response has a 2xx status code
func (o *TenancyContactsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contacts bulk delete no content response has a 3xx status code
func (o *TenancyContactsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contacts bulk delete no content response has a 4xx status code
func (o *TenancyContactsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contacts bulk delete no content response has a 5xx status code
func (o *TenancyContactsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contacts bulk delete no content response a status code equal to that given
func (o *TenancyContactsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *TenancyContactsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contacts/][%d] tenancyContactsBulkDeleteNoContent ", 204)
}

func (o *TenancyContactsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /tenancy/contacts/][%d] tenancyContactsBulkDeleteNoContent ", 204)
}

func (o *TenancyContactsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenancyContactsBulkDeleteBadRequest creates a TenancyContactsBulkDeleteBadRequest with default headers values
func NewTenancyContactsBulkDeleteBadRequest() *TenancyContactsBulkDeleteBadRequest {
	return &TenancyContactsBulkDeleteBadRequest{}
}

/*
TenancyContactsBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TenancyContactsBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this tenancy contacts bulk delete bad request response has a 2xx status code
func (o *TenancyContactsBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenancy contacts bulk delete bad request response has a 3xx status code
func (o *TenancyContactsBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contacts bulk delete bad request response has a 4xx status code
func (o *TenancyContactsBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenancy contacts bulk delete bad request response has a 5xx status code
func (o *TenancyContactsBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contacts bulk delete bad request response a status code equal to that given
func (o *TenancyContactsBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TenancyContactsBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/contacts/][%d] tenancyContactsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyContactsBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /tenancy/contacts/][%d] tenancyContactsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyContactsBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactsBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
