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

// TenancyContactGroupsUpdateReader is a Reader for the TenancyContactGroupsUpdate structure.
type TenancyContactGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyContactGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTenancyContactGroupsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTenancyContactGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactGroupsUpdateOK creates a TenancyContactGroupsUpdateOK with default headers values
func NewTenancyContactGroupsUpdateOK() *TenancyContactGroupsUpdateOK {
	return &TenancyContactGroupsUpdateOK{}
}

/*
TenancyContactGroupsUpdateOK describes a response with status code 200, with default header values.

TenancyContactGroupsUpdateOK tenancy contact groups update o k
*/
type TenancyContactGroupsUpdateOK struct {
	Payload *models.ContactGroup
}

// IsSuccess returns true when this tenancy contact groups update o k response has a 2xx status code
func (o *TenancyContactGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact groups update o k response has a 3xx status code
func (o *TenancyContactGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups update o k response has a 4xx status code
func (o *TenancyContactGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact groups update o k response has a 5xx status code
func (o *TenancyContactGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups update o k response a status code equal to that given
func (o *TenancyContactGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyContactGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactGroupsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyContactGroupsUpdateOK) GetPayload() *models.ContactGroup {
	return o.Payload
}

func (o *TenancyContactGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactGroupsUpdateBadRequest creates a TenancyContactGroupsUpdateBadRequest with default headers values
func NewTenancyContactGroupsUpdateBadRequest() *TenancyContactGroupsUpdateBadRequest {
	return &TenancyContactGroupsUpdateBadRequest{}
}

/*
TenancyContactGroupsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TenancyContactGroupsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this tenancy contact groups update bad request response has a 2xx status code
func (o *TenancyContactGroupsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenancy contact groups update bad request response has a 3xx status code
func (o *TenancyContactGroupsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups update bad request response has a 4xx status code
func (o *TenancyContactGroupsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenancy contact groups update bad request response has a 5xx status code
func (o *TenancyContactGroupsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups update bad request response a status code equal to that given
func (o *TenancyContactGroupsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TenancyContactGroupsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyContactGroupsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-groups/{id}/][%d] tenancyContactGroupsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyContactGroupsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactGroupsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactGroupsUpdateDefault creates a TenancyContactGroupsUpdateDefault with default headers values
func NewTenancyContactGroupsUpdateDefault(code int) *TenancyContactGroupsUpdateDefault {
	return &TenancyContactGroupsUpdateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactGroupsUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type TenancyContactGroupsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact groups update default response
func (o *TenancyContactGroupsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contact groups update default response has a 2xx status code
func (o *TenancyContactGroupsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact groups update default response has a 3xx status code
func (o *TenancyContactGroupsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact groups update default response has a 4xx status code
func (o *TenancyContactGroupsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact groups update default response has a 5xx status code
func (o *TenancyContactGroupsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact groups update default response a status code equal to that given
func (o *TenancyContactGroupsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactGroupsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /tenancy/contact-groups/{id}/][%d] tenancy_contact-groups_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactGroupsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /tenancy/contact-groups/{id}/][%d] tenancy_contact-groups_update default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactGroupsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
