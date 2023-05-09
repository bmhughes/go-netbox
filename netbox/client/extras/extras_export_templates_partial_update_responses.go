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

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// ExtrasExportTemplatesPartialUpdateReader is a Reader for the ExtrasExportTemplatesPartialUpdate structure.
type ExtrasExportTemplatesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasExportTemplatesPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasExportTemplatesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasExportTemplatesPartialUpdateOK creates a ExtrasExportTemplatesPartialUpdateOK with default headers values
func NewExtrasExportTemplatesPartialUpdateOK() *ExtrasExportTemplatesPartialUpdateOK {
	return &ExtrasExportTemplatesPartialUpdateOK{}
}

/*
ExtrasExportTemplatesPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasExportTemplatesPartialUpdateOK extras export templates partial update o k
*/
type ExtrasExportTemplatesPartialUpdateOK struct {
	Payload *models.ExportTemplate
}

// IsSuccess returns true when this extras export templates partial update o k response has a 2xx status code
func (o *ExtrasExportTemplatesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras export templates partial update o k response has a 3xx status code
func (o *ExtrasExportTemplatesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates partial update o k response has a 4xx status code
func (o *ExtrasExportTemplatesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras export templates partial update o k response has a 5xx status code
func (o *ExtrasExportTemplatesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates partial update o k response a status code equal to that given
func (o *ExtrasExportTemplatesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasExportTemplatesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/{id}/][%d] extrasExportTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/{id}/][%d] extrasExportTemplatesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesPartialUpdateOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesPartialUpdateBadRequest creates a ExtrasExportTemplatesPartialUpdateBadRequest with default headers values
func NewExtrasExportTemplatesPartialUpdateBadRequest() *ExtrasExportTemplatesPartialUpdateBadRequest {
	return &ExtrasExportTemplatesPartialUpdateBadRequest{}
}

/*
ExtrasExportTemplatesPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasExportTemplatesPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras export templates partial update bad request response has a 2xx status code
func (o *ExtrasExportTemplatesPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras export templates partial update bad request response has a 3xx status code
func (o *ExtrasExportTemplatesPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates partial update bad request response has a 4xx status code
func (o *ExtrasExportTemplatesPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras export templates partial update bad request response has a 5xx status code
func (o *ExtrasExportTemplatesPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates partial update bad request response a status code equal to that given
func (o *ExtrasExportTemplatesPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasExportTemplatesPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/{id}/][%d] extrasExportTemplatesPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasExportTemplatesPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/{id}/][%d] extrasExportTemplatesPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasExportTemplatesPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesPartialUpdateDefault creates a ExtrasExportTemplatesPartialUpdateDefault with default headers values
func NewExtrasExportTemplatesPartialUpdateDefault(code int) *ExtrasExportTemplatesPartialUpdateDefault {
	return &ExtrasExportTemplatesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasExportTemplatesPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasExportTemplatesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras export templates partial update default response
func (o *ExtrasExportTemplatesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras export templates partial update default response has a 2xx status code
func (o *ExtrasExportTemplatesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras export templates partial update default response has a 3xx status code
func (o *ExtrasExportTemplatesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras export templates partial update default response has a 4xx status code
func (o *ExtrasExportTemplatesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras export templates partial update default response has a 5xx status code
func (o *ExtrasExportTemplatesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras export templates partial update default response a status code equal to that given
func (o *ExtrasExportTemplatesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasExportTemplatesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/{id}/][%d] extras_export-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/export-templates/{id}/][%d] extras_export-templates_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
