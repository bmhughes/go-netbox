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

// ExtrasExportTemplatesDeleteReader is a Reader for the ExtrasExportTemplatesDelete structure.
type ExtrasExportTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasExportTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasExportTemplatesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasExportTemplatesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasExportTemplatesDeleteNoContent creates a ExtrasExportTemplatesDeleteNoContent with default headers values
func NewExtrasExportTemplatesDeleteNoContent() *ExtrasExportTemplatesDeleteNoContent {
	return &ExtrasExportTemplatesDeleteNoContent{}
}

/*
ExtrasExportTemplatesDeleteNoContent describes a response with status code 204, with default header values.

ExtrasExportTemplatesDeleteNoContent extras export templates delete no content
*/
type ExtrasExportTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this extras export templates delete no content response has a 2xx status code
func (o *ExtrasExportTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras export templates delete no content response has a 3xx status code
func (o *ExtrasExportTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates delete no content response has a 4xx status code
func (o *ExtrasExportTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras export templates delete no content response has a 5xx status code
func (o *ExtrasExportTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates delete no content response a status code equal to that given
func (o *ExtrasExportTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ExtrasExportTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/{id}/][%d] extrasExportTemplatesDeleteNoContent ", 204)
}

func (o *ExtrasExportTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/{id}/][%d] extrasExportTemplatesDeleteNoContent ", 204)
}

func (o *ExtrasExportTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasExportTemplatesDeleteBadRequest creates a ExtrasExportTemplatesDeleteBadRequest with default headers values
func NewExtrasExportTemplatesDeleteBadRequest() *ExtrasExportTemplatesDeleteBadRequest {
	return &ExtrasExportTemplatesDeleteBadRequest{}
}

/*
ExtrasExportTemplatesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasExportTemplatesDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras export templates delete bad request response has a 2xx status code
func (o *ExtrasExportTemplatesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras export templates delete bad request response has a 3xx status code
func (o *ExtrasExportTemplatesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates delete bad request response has a 4xx status code
func (o *ExtrasExportTemplatesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras export templates delete bad request response has a 5xx status code
func (o *ExtrasExportTemplatesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates delete bad request response a status code equal to that given
func (o *ExtrasExportTemplatesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasExportTemplatesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/{id}/][%d] extrasExportTemplatesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasExportTemplatesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/{id}/][%d] extrasExportTemplatesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasExportTemplatesDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesDeleteDefault creates a ExtrasExportTemplatesDeleteDefault with default headers values
func NewExtrasExportTemplatesDeleteDefault(code int) *ExtrasExportTemplatesDeleteDefault {
	return &ExtrasExportTemplatesDeleteDefault{
		_statusCode: code,
	}
}

/*
ExtrasExportTemplatesDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasExportTemplatesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras export templates delete default response
func (o *ExtrasExportTemplatesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras export templates delete default response has a 2xx status code
func (o *ExtrasExportTemplatesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras export templates delete default response has a 3xx status code
func (o *ExtrasExportTemplatesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras export templates delete default response has a 4xx status code
func (o *ExtrasExportTemplatesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras export templates delete default response has a 5xx status code
func (o *ExtrasExportTemplatesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras export templates delete default response a status code equal to that given
func (o *ExtrasExportTemplatesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasExportTemplatesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/{id}/][%d] extras_export-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/{id}/][%d] extras_export-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
