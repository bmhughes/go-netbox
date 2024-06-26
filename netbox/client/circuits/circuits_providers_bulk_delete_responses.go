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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CircuitsProvidersBulkDeleteReader is a Reader for the CircuitsProvidersBulkDelete structure.
type CircuitsProvidersBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCircuitsProvidersBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCircuitsProvidersBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCircuitsProvidersBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProvidersBulkDeleteNoContent creates a CircuitsProvidersBulkDeleteNoContent with default headers values
func NewCircuitsProvidersBulkDeleteNoContent() *CircuitsProvidersBulkDeleteNoContent {
	return &CircuitsProvidersBulkDeleteNoContent{}
}

/*
CircuitsProvidersBulkDeleteNoContent describes a response with status code 204, with default header values.

CircuitsProvidersBulkDeleteNoContent circuits providers bulk delete no content
*/
type CircuitsProvidersBulkDeleteNoContent struct {
}

// IsSuccess returns true when this circuits providers bulk delete no content response has a 2xx status code
func (o *CircuitsProvidersBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits providers bulk delete no content response has a 3xx status code
func (o *CircuitsProvidersBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits providers bulk delete no content response has a 4xx status code
func (o *CircuitsProvidersBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits providers bulk delete no content response has a 5xx status code
func (o *CircuitsProvidersBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits providers bulk delete no content response a status code equal to that given
func (o *CircuitsProvidersBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *CircuitsProvidersBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /circuits/providers/][%d] circuitsProvidersBulkDeleteNoContent ", 204)
}

func (o *CircuitsProvidersBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /circuits/providers/][%d] circuitsProvidersBulkDeleteNoContent ", 204)
}

func (o *CircuitsProvidersBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCircuitsProvidersBulkDeleteBadRequest creates a CircuitsProvidersBulkDeleteBadRequest with default headers values
func NewCircuitsProvidersBulkDeleteBadRequest() *CircuitsProvidersBulkDeleteBadRequest {
	return &CircuitsProvidersBulkDeleteBadRequest{}
}

/*
CircuitsProvidersBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CircuitsProvidersBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this circuits providers bulk delete bad request response has a 2xx status code
func (o *CircuitsProvidersBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this circuits providers bulk delete bad request response has a 3xx status code
func (o *CircuitsProvidersBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits providers bulk delete bad request response has a 4xx status code
func (o *CircuitsProvidersBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this circuits providers bulk delete bad request response has a 5xx status code
func (o *CircuitsProvidersBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits providers bulk delete bad request response a status code equal to that given
func (o *CircuitsProvidersBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CircuitsProvidersBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /circuits/providers/][%d] circuitsProvidersBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsProvidersBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /circuits/providers/][%d] circuitsProvidersBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsProvidersBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProvidersBulkDeleteDefault creates a CircuitsProvidersBulkDeleteDefault with default headers values
func NewCircuitsProvidersBulkDeleteDefault(code int) *CircuitsProvidersBulkDeleteDefault {
	return &CircuitsProvidersBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
CircuitsProvidersBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type CircuitsProvidersBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits providers bulk delete default response
func (o *CircuitsProvidersBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits providers bulk delete default response has a 2xx status code
func (o *CircuitsProvidersBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits providers bulk delete default response has a 3xx status code
func (o *CircuitsProvidersBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits providers bulk delete default response has a 4xx status code
func (o *CircuitsProvidersBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits providers bulk delete default response has a 5xx status code
func (o *CircuitsProvidersBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits providers bulk delete default response a status code equal to that given
func (o *CircuitsProvidersBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsProvidersBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /circuits/providers/][%d] circuits_providers_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /circuits/providers/][%d] circuits_providers_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
