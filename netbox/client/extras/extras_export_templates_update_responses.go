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

// ExtrasExportTemplatesUpdateReader is a Reader for the ExtrasExportTemplatesUpdate structure.
type ExtrasExportTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasExportTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasExportTemplatesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasExportTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasExportTemplatesUpdateOK creates a ExtrasExportTemplatesUpdateOK with default headers values
func NewExtrasExportTemplatesUpdateOK() *ExtrasExportTemplatesUpdateOK {
	return &ExtrasExportTemplatesUpdateOK{}
}

/*
ExtrasExportTemplatesUpdateOK describes a response with status code 200, with default header values.

ExtrasExportTemplatesUpdateOK extras export templates update o k
*/
type ExtrasExportTemplatesUpdateOK struct {
	Payload *models.ExportTemplate
}

// IsSuccess returns true when this extras export templates update o k response has a 2xx status code
func (o *ExtrasExportTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras export templates update o k response has a 3xx status code
func (o *ExtrasExportTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates update o k response has a 4xx status code
func (o *ExtrasExportTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras export templates update o k response has a 5xx status code
func (o *ExtrasExportTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates update o k response a status code equal to that given
func (o *ExtrasExportTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasExportTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/export-templates/{id}/][%d] extrasExportTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/export-templates/{id}/][%d] extrasExportTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesUpdateOK) GetPayload() *models.ExportTemplate {
	return o.Payload
}

func (o *ExtrasExportTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesUpdateBadRequest creates a ExtrasExportTemplatesUpdateBadRequest with default headers values
func NewExtrasExportTemplatesUpdateBadRequest() *ExtrasExportTemplatesUpdateBadRequest {
	return &ExtrasExportTemplatesUpdateBadRequest{}
}

/*
ExtrasExportTemplatesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasExportTemplatesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras export templates update bad request response has a 2xx status code
func (o *ExtrasExportTemplatesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras export templates update bad request response has a 3xx status code
func (o *ExtrasExportTemplatesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras export templates update bad request response has a 4xx status code
func (o *ExtrasExportTemplatesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras export templates update bad request response has a 5xx status code
func (o *ExtrasExportTemplatesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras export templates update bad request response a status code equal to that given
func (o *ExtrasExportTemplatesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasExportTemplatesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /extras/export-templates/{id}/][%d] extrasExportTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasExportTemplatesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /extras/export-templates/{id}/][%d] extrasExportTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasExportTemplatesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasExportTemplatesUpdateDefault creates a ExtrasExportTemplatesUpdateDefault with default headers values
func NewExtrasExportTemplatesUpdateDefault(code int) *ExtrasExportTemplatesUpdateDefault {
	return &ExtrasExportTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasExportTemplatesUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasExportTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras export templates update default response
func (o *ExtrasExportTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras export templates update default response has a 2xx status code
func (o *ExtrasExportTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras export templates update default response has a 3xx status code
func (o *ExtrasExportTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras export templates update default response has a 4xx status code
func (o *ExtrasExportTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras export templates update default response has a 5xx status code
func (o *ExtrasExportTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras export templates update default response a status code equal to that given
func (o *ExtrasExportTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasExportTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/export-templates/{id}/][%d] extras_export-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/export-templates/{id}/][%d] extras_export-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasExportTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasExportTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
