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

// TenancyContactGroupsCreateReader is a Reader for the TenancyContactGroupsCreate structure.
type TenancyContactGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyContactGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewTenancyContactGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTenancyContactGroupsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewTenancyContactGroupsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyContactGroupsCreateCreated creates a TenancyContactGroupsCreateCreated with default headers values
func NewTenancyContactGroupsCreateCreated() *TenancyContactGroupsCreateCreated {
	return &TenancyContactGroupsCreateCreated{}
}

/*
TenancyContactGroupsCreateCreated describes a response with status code 201, with default header values.

TenancyContactGroupsCreateCreated tenancy contact groups create created
*/
type TenancyContactGroupsCreateCreated struct {
	Payload *models.ContactGroup
}

// IsSuccess returns true when this tenancy contact groups create created response has a 2xx status code
func (o *TenancyContactGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy contact groups create created response has a 3xx status code
func (o *TenancyContactGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups create created response has a 4xx status code
func (o *TenancyContactGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy contact groups create created response has a 5xx status code
func (o *TenancyContactGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups create created response a status code equal to that given
func (o *TenancyContactGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *TenancyContactGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /tenancy/contact-groups/][%d] tenancyContactGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyContactGroupsCreateCreated) String() string {
	return fmt.Sprintf("[POST /tenancy/contact-groups/][%d] tenancyContactGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *TenancyContactGroupsCreateCreated) GetPayload() *models.ContactGroup {
	return o.Payload
}

func (o *TenancyContactGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactGroupsCreateBadRequest creates a TenancyContactGroupsCreateBadRequest with default headers values
func NewTenancyContactGroupsCreateBadRequest() *TenancyContactGroupsCreateBadRequest {
	return &TenancyContactGroupsCreateBadRequest{}
}

/*
TenancyContactGroupsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TenancyContactGroupsCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this tenancy contact groups create bad request response has a 2xx status code
func (o *TenancyContactGroupsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenancy contact groups create bad request response has a 3xx status code
func (o *TenancyContactGroupsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy contact groups create bad request response has a 4xx status code
func (o *TenancyContactGroupsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenancy contact groups create bad request response has a 5xx status code
func (o *TenancyContactGroupsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy contact groups create bad request response a status code equal to that given
func (o *TenancyContactGroupsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TenancyContactGroupsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /tenancy/contact-groups/][%d] tenancyContactGroupsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyContactGroupsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /tenancy/contact-groups/][%d] tenancyContactGroupsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyContactGroupsCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactGroupsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyContactGroupsCreateDefault creates a TenancyContactGroupsCreateDefault with default headers values
func NewTenancyContactGroupsCreateDefault(code int) *TenancyContactGroupsCreateDefault {
	return &TenancyContactGroupsCreateDefault{
		_statusCode: code,
	}
}

/*
TenancyContactGroupsCreateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type TenancyContactGroupsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy contact groups create default response
func (o *TenancyContactGroupsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy contact groups create default response has a 2xx status code
func (o *TenancyContactGroupsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy contact groups create default response has a 3xx status code
func (o *TenancyContactGroupsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy contact groups create default response has a 4xx status code
func (o *TenancyContactGroupsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy contact groups create default response has a 5xx status code
func (o *TenancyContactGroupsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy contact groups create default response a status code equal to that given
func (o *TenancyContactGroupsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyContactGroupsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /tenancy/contact-groups/][%d] tenancy_contact-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactGroupsCreateDefault) String() string {
	return fmt.Sprintf("[POST /tenancy/contact-groups/][%d] tenancy_contact-groups_create default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyContactGroupsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyContactGroupsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
