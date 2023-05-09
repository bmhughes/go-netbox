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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ExtrasConfigContextsBulkDeleteReader is a Reader for the ExtrasConfigContextsBulkDelete structure.
type ExtrasConfigContextsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasConfigContextsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasConfigContextsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasConfigContextsBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasConfigContextsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasConfigContextsBulkDeleteNoContent creates a ExtrasConfigContextsBulkDeleteNoContent with default headers values
func NewExtrasConfigContextsBulkDeleteNoContent() *ExtrasConfigContextsBulkDeleteNoContent {
	return &ExtrasConfigContextsBulkDeleteNoContent{}
}

/*
ExtrasConfigContextsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasConfigContextsBulkDeleteNoContent extras config contexts bulk delete no content
*/
type ExtrasConfigContextsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this extras config contexts bulk delete no content response has a 2xx status code
func (o *ExtrasConfigContextsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras config contexts bulk delete no content response has a 3xx status code
func (o *ExtrasConfigContextsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config contexts bulk delete no content response has a 4xx status code
func (o *ExtrasConfigContextsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras config contexts bulk delete no content response has a 5xx status code
func (o *ExtrasConfigContextsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config contexts bulk delete no content response a status code equal to that given
func (o *ExtrasConfigContextsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ExtrasConfigContextsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/config-contexts/][%d] extrasConfigContextsBulkDeleteNoContent ", 204)
}

func (o *ExtrasConfigContextsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/config-contexts/][%d] extrasConfigContextsBulkDeleteNoContent ", 204)
}

func (o *ExtrasConfigContextsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasConfigContextsBulkDeleteBadRequest creates a ExtrasConfigContextsBulkDeleteBadRequest with default headers values
func NewExtrasConfigContextsBulkDeleteBadRequest() *ExtrasConfigContextsBulkDeleteBadRequest {
	return &ExtrasConfigContextsBulkDeleteBadRequest{}
}

/*
ExtrasConfigContextsBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasConfigContextsBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras config contexts bulk delete bad request response has a 2xx status code
func (o *ExtrasConfigContextsBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras config contexts bulk delete bad request response has a 3xx status code
func (o *ExtrasConfigContextsBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras config contexts bulk delete bad request response has a 4xx status code
func (o *ExtrasConfigContextsBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras config contexts bulk delete bad request response has a 5xx status code
func (o *ExtrasConfigContextsBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras config contexts bulk delete bad request response a status code equal to that given
func (o *ExtrasConfigContextsBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasConfigContextsBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /extras/config-contexts/][%d] extrasConfigContextsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasConfigContextsBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /extras/config-contexts/][%d] extrasConfigContextsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasConfigContextsBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasConfigContextsBulkDeleteDefault creates a ExtrasConfigContextsBulkDeleteDefault with default headers values
func NewExtrasConfigContextsBulkDeleteDefault(code int) *ExtrasConfigContextsBulkDeleteDefault {
	return &ExtrasConfigContextsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
ExtrasConfigContextsBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasConfigContextsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras config contexts bulk delete default response
func (o *ExtrasConfigContextsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras config contexts bulk delete default response has a 2xx status code
func (o *ExtrasConfigContextsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras config contexts bulk delete default response has a 3xx status code
func (o *ExtrasConfigContextsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras config contexts bulk delete default response has a 4xx status code
func (o *ExtrasConfigContextsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras config contexts bulk delete default response has a 5xx status code
func (o *ExtrasConfigContextsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras config contexts bulk delete default response a status code equal to that given
func (o *ExtrasConfigContextsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasConfigContextsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/config-contexts/][%d] extras_config-contexts_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigContextsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /extras/config-contexts/][%d] extras_config-contexts_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasConfigContextsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasConfigContextsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
