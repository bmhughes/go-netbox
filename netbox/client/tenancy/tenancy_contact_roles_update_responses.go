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

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// TenancyContactRolesUpdateReader is a Reader for the TenancyContactRolesUpdate structure.
type TenancyContactRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTenancyContactRolesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTenancyContactRolesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactRolesUpdateOK creates a TenancyContactRolesUpdateOK with default headers values
func NewTenancyContactRolesUpdateOK() *TenancyContactRolesUpdateOK {
	return &TenancyContactRolesUpdateOK{}
}

/*
TenancyContactRolesUpdateOK describes a response with status code 200, with default header values.

TenancyContactRolesUpdateOK tenancy contact roles update o k
*/
type TenancyContactRolesUpdateOK struct {
	Payload *models.ContactRole
}

// IsSuccess returns true when this tenancy contact roles update o k response has a 2xx status code
func (o *TenancyContactRolesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact roles update o k response has a 3xx status code
func (o *TenancyContactRolesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact roles update o k response has a 4xx status code
func (o *TenancyContactRolesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact roles update o k response has a 5xx status code
func (o *TenancyContactRolesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact roles update o k response a status code equal to that given
func (o *TenancyContactRolesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyContactRolesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/{id}/][%d] tenancyContactRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactRolesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/{id}/][%d] tenancyContactRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactRolesUpdateOK) GetPayload() *models.ContactRole {
	return o.Payload
}

func (o *TenancyContactRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactRolesUpdateBadRequest creates a TenancyContactRolesUpdateBadRequest with default headers values
func NewTenancyContactRolesUpdateBadRequest() *TenancyContactRolesUpdateBadRequest {
	return &TenancyContactRolesUpdateBadRequest{}
}

/*
TenancyContactRolesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TenancyContactRolesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this tenancy contact roles update bad request response has a 2xx status code
func (o *TenancyContactRolesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenancy contact roles update bad request response has a 3xx status code
func (o *TenancyContactRolesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact roles update bad request response has a 4xx status code
func (o *TenancyContactRolesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenancy contact roles update bad request response has a 5xx status code
func (o *TenancyContactRolesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact roles update bad request response a status code equal to that given
func (o *TenancyContactRolesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TenancyContactRolesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/{id}/][%d] tenancyContactRolesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyContactRolesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/{id}/][%d] tenancyContactRolesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyContactRolesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactRolesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactRolesUpdateDefault creates a TenancyContactRolesUpdateDefault with default headers values
func NewTenancyContactRolesUpdateDefault(code int) *TenancyContactRolesUpdateDefault {
	return &TenancyContactRolesUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactRolesUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type TenancyContactRolesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact roles update default response
func (o *TenancyContactRolesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contact roles update default response has a 2xx status code
func (o *TenancyContactRolesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact roles update default response has a 3xx status code
func (o *TenancyContactRolesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact roles update default response has a 4xx status code
func (o *TenancyContactRolesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact roles update default response has a 5xx status code
func (o *TenancyContactRolesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact roles update default response a status code equal to that given
func (o *TenancyContactRolesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactRolesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/{id}/][%d] tenancy_contact-roles_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-roles/{id}/][%d] tenancy_contact-roles_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactRolesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactRolesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
