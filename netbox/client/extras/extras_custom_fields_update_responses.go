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

// ExtrasCustomFieldsUpdateReader is a Reader for the ExtrasCustomFieldsUpdate structure.
type ExtrasCustomFieldsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasCustomFieldsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasCustomFieldsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldsUpdateOK creates a ExtrasCustomFieldsUpdateOK with default headers values
func NewExtrasCustomFieldsUpdateOK() *ExtrasCustomFieldsUpdateOK {
	return &ExtrasCustomFieldsUpdateOK{}
}

/*
ExtrasCustomFieldsUpdateOK describes a response with status code 200, with default header values.

ExtrasCustomFieldsUpdateOK extras custom fields update o k
*/
type ExtrasCustomFieldsUpdateOK struct {
	Payload *models.CustomField
}

// IsSuccess returns true when this extras custom fields update o k response has a 2xx status code
func (o *ExtrasCustomFieldsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom fields update o k response has a 3xx status code
func (o *ExtrasCustomFieldsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom fields update o k response has a 4xx status code
func (o *ExtrasCustomFieldsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom fields update o k response has a 5xx status code
func (o *ExtrasCustomFieldsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom fields update o k response a status code equal to that given
func (o *ExtrasCustomFieldsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasCustomFieldsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-fields/{id}/][%d] extrasCustomFieldsUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/custom-fields/{id}/][%d] extrasCustomFieldsUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldsUpdateOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldsUpdateBadRequest creates a ExtrasCustomFieldsUpdateBadRequest with default headers values
func NewExtrasCustomFieldsUpdateBadRequest() *ExtrasCustomFieldsUpdateBadRequest {
	return &ExtrasCustomFieldsUpdateBadRequest{}
}

/*
ExtrasCustomFieldsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasCustomFieldsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras custom fields update bad request response has a 2xx status code
func (o *ExtrasCustomFieldsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras custom fields update bad request response has a 3xx status code
func (o *ExtrasCustomFieldsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom fields update bad request response has a 4xx status code
func (o *ExtrasCustomFieldsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras custom fields update bad request response has a 5xx status code
func (o *ExtrasCustomFieldsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom fields update bad request response a status code equal to that given
func (o *ExtrasCustomFieldsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasCustomFieldsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-fields/{id}/][%d] extrasCustomFieldsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasCustomFieldsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /extras/custom-fields/{id}/][%d] extrasCustomFieldsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasCustomFieldsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldsUpdateDefault creates a ExtrasCustomFieldsUpdateDefault with default headers values
func NewExtrasCustomFieldsUpdateDefault(code int) *ExtrasCustomFieldsUpdateDefault {
	return &ExtrasCustomFieldsUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomFieldsUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasCustomFieldsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras custom fields update default response
func (o *ExtrasCustomFieldsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras custom fields update default response has a 2xx status code
func (o *ExtrasCustomFieldsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom fields update default response has a 3xx status code
func (o *ExtrasCustomFieldsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom fields update default response has a 4xx status code
func (o *ExtrasCustomFieldsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom fields update default response has a 5xx status code
func (o *ExtrasCustomFieldsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom fields update default response a status code equal to that given
func (o *ExtrasCustomFieldsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasCustomFieldsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/custom-fields/{id}/][%d] extras_custom-fields_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/custom-fields/{id}/][%d] extras_custom-fields_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
