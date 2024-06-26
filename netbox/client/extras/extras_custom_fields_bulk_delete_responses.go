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

// ExtrasCustomFieldsBulkDeleteReader is a Reader for the ExtrasCustomFieldsBulkDelete structure.
type ExtrasCustomFieldsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasCustomFieldsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasCustomFieldsBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasCustomFieldsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldsBulkDeleteNoContent creates a ExtrasCustomFieldsBulkDeleteNoContent with default headers values
func NewExtrasCustomFieldsBulkDeleteNoContent() *ExtrasCustomFieldsBulkDeleteNoContent {
	return &ExtrasCustomFieldsBulkDeleteNoContent{}
}

/*
ExtrasCustomFieldsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasCustomFieldsBulkDeleteNoContent extras custom fields bulk delete no content
*/
type ExtrasCustomFieldsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this extras custom fields bulk delete no content response has a 2xx status code
func (o *ExtrasCustomFieldsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom fields bulk delete no content response has a 3xx status code
func (o *ExtrasCustomFieldsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom fields bulk delete no content response has a 4xx status code
func (o *ExtrasCustomFieldsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom fields bulk delete no content response has a 5xx status code
func (o *ExtrasCustomFieldsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom fields bulk delete no content response a status code equal to that given
func (o *ExtrasCustomFieldsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ExtrasCustomFieldsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-fields/][%d] extrasCustomFieldsBulkDeleteNoContent ", 204)
}

func (o *ExtrasCustomFieldsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/custom-fields/][%d] extrasCustomFieldsBulkDeleteNoContent ", 204)
}

func (o *ExtrasCustomFieldsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasCustomFieldsBulkDeleteBadRequest creates a ExtrasCustomFieldsBulkDeleteBadRequest with default headers values
func NewExtrasCustomFieldsBulkDeleteBadRequest() *ExtrasCustomFieldsBulkDeleteBadRequest {
	return &ExtrasCustomFieldsBulkDeleteBadRequest{}
}

/*
ExtrasCustomFieldsBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasCustomFieldsBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras custom fields bulk delete bad request response has a 2xx status code
func (o *ExtrasCustomFieldsBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras custom fields bulk delete bad request response has a 3xx status code
func (o *ExtrasCustomFieldsBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom fields bulk delete bad request response has a 4xx status code
func (o *ExtrasCustomFieldsBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras custom fields bulk delete bad request response has a 5xx status code
func (o *ExtrasCustomFieldsBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom fields bulk delete bad request response a status code equal to that given
func (o *ExtrasCustomFieldsBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasCustomFieldsBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-fields/][%d] extrasCustomFieldsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasCustomFieldsBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /extras/custom-fields/][%d] extrasCustomFieldsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasCustomFieldsBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldsBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldsBulkDeleteDefault creates a ExtrasCustomFieldsBulkDeleteDefault with default headers values
func NewExtrasCustomFieldsBulkDeleteDefault(code int) *ExtrasCustomFieldsBulkDeleteDefault {
	return &ExtrasCustomFieldsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomFieldsBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasCustomFieldsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom fields bulk delete default response
func (o *ExtrasCustomFieldsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras custom fields bulk delete default response has a 2xx status code
func (o *ExtrasCustomFieldsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom fields bulk delete default response has a 3xx status code
func (o *ExtrasCustomFieldsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom fields bulk delete default response has a 4xx status code
func (o *ExtrasCustomFieldsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom fields bulk delete default response has a 5xx status code
func (o *ExtrasCustomFieldsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom fields bulk delete default response a status code equal to that given
func (o *ExtrasCustomFieldsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasCustomFieldsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/custom-fields/][%d] extras_custom-fields_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /extras/custom-fields/][%d] extras_custom-fields_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
