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

// TenancyContactGroupsReadReader is a Reader for the TenancyContactGroupsRead structure.
type TenancyContactGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTenancyContactGroupsReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTenancyContactGroupsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactGroupsReadOK creates a TenancyContactGroupsReadOK with default headers values
func NewTenancyContactGroupsReadOK() *TenancyContactGroupsReadOK {
	return &TenancyContactGroupsReadOK{}
}

/*
TenancyContactGroupsReadOK describes a response with status code 200, with default header values.

TenancyContactGroupsReadOK tenancy contact groups read o k
*/
type TenancyContactGroupsReadOK struct {
	Payload *models.ContactGroup
}

// IsSuccess returns true when this tenancy contact groups read o k response has a 2xx status code
func (o *TenancyContactGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact groups read o k response has a 3xx status code
func (o *TenancyContactGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups read o k response has a 4xx status code
func (o *TenancyContactGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact groups read o k response has a 5xx status code
func (o *TenancyContactGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups read o k response a status code equal to that given
func (o *TenancyContactGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyContactGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsReadOK  %+v", 200, o.Payload)
}

func (o *TenancyContactGroupsReadOK) String() string {
	return fmt.Sprintf("[GET /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsReadOK  %+v", 200, o.Payload)
}

func (o *TenancyContactGroupsReadOK) GetPayload() *models.ContactGroup {
	return o.Payload
}

func (o *TenancyContactGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactGroupsReadBadRequest creates a TenancyContactGroupsReadBadRequest with default headers values
func NewTenancyContactGroupsReadBadRequest() *TenancyContactGroupsReadBadRequest {
	return &TenancyContactGroupsReadBadRequest{}
}

/*
TenancyContactGroupsReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TenancyContactGroupsReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this tenancy contact groups read bad request response has a 2xx status code
func (o *TenancyContactGroupsReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenancy contact groups read bad request response has a 3xx status code
func (o *TenancyContactGroupsReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups read bad request response has a 4xx status code
func (o *TenancyContactGroupsReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenancy contact groups read bad request response has a 5xx status code
func (o *TenancyContactGroupsReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups read bad request response a status code equal to that given
func (o *TenancyContactGroupsReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TenancyContactGroupsReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsReadBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyContactGroupsReadBadRequest) String() string {
	return fmt.Sprintf("[GET /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsReadBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyContactGroupsReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactGroupsReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactGroupsReadDefault creates a TenancyContactGroupsReadDefault with default headers values
func NewTenancyContactGroupsReadDefault(code int) *TenancyContactGroupsReadDefault {
	return &TenancyContactGroupsReadDefault{
		_statusCode: code,
	}
}

/*
TenancyContactGroupsReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type TenancyContactGroupsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact groups read default response
func (o *TenancyContactGroupsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contact groups read default response has a 2xx status code
func (o *TenancyContactGroupsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact groups read default response has a 3xx status code
func (o *TenancyContactGroupsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact groups read default response has a 4xx status code
func (o *TenancyContactGroupsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact groups read default response has a 5xx status code
func (o *TenancyContactGroupsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact groups read default response a status code equal to that given
func (o *TenancyContactGroupsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactGroupsReadDefault) Error() string {
	return fmt.Sprintf("[GET /tenancy/contact-groups/{id}/][%d] tenancy_contact-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactGroupsReadDefault) String() string {
	return fmt.Sprintf("[GET /tenancy/contact-groups/{id}/][%d] tenancy_contact-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactGroupsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactGroupsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
