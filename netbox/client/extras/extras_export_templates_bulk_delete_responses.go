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

// ExtrasExportTemplatesBulkDeleteReader is a Reader for the ExtrasExportTemplatesBulkDelete structure.
type ExtrasExportTemplatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasExportTemplatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasExportTemplatesBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasExportTemplatesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasExportTemplatesBulkDeleteNoContent creates a ExtrasExportTemplatesBulkDeleteNoContent with default headers values
func NewExtrasExportTemplatesBulkDeleteNoContent() *ExtrasExportTemplatesBulkDeleteNoContent {
	return &ExtrasExportTemplatesBulkDeleteNoContent{}
}

/*
ExtrasExportTemplatesBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasExportTemplatesBulkDeleteNoContent extras export templates bulk delete no content
*/
type ExtrasExportTemplatesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this extras export templates bulk delete no content response has a 2xx status code
func (o *ExtrasExportTemplatesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras export templates bulk delete no content response has a 3xx status code
func (o *ExtrasExportTemplatesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates bulk delete no content response has a 4xx status code
func (o *ExtrasExportTemplatesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras export templates bulk delete no content response has a 5xx status code
func (o *ExtrasExportTemplatesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates bulk delete no content response a status code equal to that given
func (o *ExtrasExportTemplatesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ExtrasExportTemplatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/][%d] extrasExportTemplatesBulkDeleteNoContent ", 204)
}

func (o *ExtrasExportTemplatesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/][%d] extrasExportTemplatesBulkDeleteNoContent ", 204)
}

func (o *ExtrasExportTemplatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasExportTemplatesBulkDeleteBadRequest creates a ExtrasExportTemplatesBulkDeleteBadRequest with default headers values
func NewExtrasExportTemplatesBulkDeleteBadRequest() *ExtrasExportTemplatesBulkDeleteBadRequest {
	return &ExtrasExportTemplatesBulkDeleteBadRequest{}
}

/*
ExtrasExportTemplatesBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasExportTemplatesBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras export templates bulk delete bad request response has a 2xx status code
func (o *ExtrasExportTemplatesBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras export templates bulk delete bad request response has a 3xx status code
func (o *ExtrasExportTemplatesBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates bulk delete bad request response has a 4xx status code
func (o *ExtrasExportTemplatesBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras export templates bulk delete bad request response has a 5xx status code
func (o *ExtrasExportTemplatesBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates bulk delete bad request response a status code equal to that given
func (o *ExtrasExportTemplatesBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasExportTemplatesBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/][%d] extrasExportTemplatesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasExportTemplatesBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/][%d] extrasExportTemplatesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasExportTemplatesBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesBulkDeleteDefault creates a ExtrasExportTemplatesBulkDeleteDefault with default headers values
func NewExtrasExportTemplatesBulkDeleteDefault(code int) *ExtrasExportTemplatesBulkDeleteDefault {
	return &ExtrasExportTemplatesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
ExtrasExportTemplatesBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasExportTemplatesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras export templates bulk delete default response
func (o *ExtrasExportTemplatesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras export templates bulk delete default response has a 2xx status code
func (o *ExtrasExportTemplatesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras export templates bulk delete default response has a 3xx status code
func (o *ExtrasExportTemplatesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras export templates bulk delete default response has a 4xx status code
func (o *ExtrasExportTemplatesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras export templates bulk delete default response has a 5xx status code
func (o *ExtrasExportTemplatesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras export templates bulk delete default response a status code equal to that given
func (o *ExtrasExportTemplatesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasExportTemplatesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/][%d] extras_export-templates_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /extras/export-templates/][%d] extras_export-templates_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
