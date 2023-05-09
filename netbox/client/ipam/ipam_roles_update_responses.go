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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// IpamRolesUpdateReader is a Reader for the IpamRolesUpdate structure.
type IpamRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamRolesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamRolesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRolesUpdateOK creates a IpamRolesUpdateOK with default headers values
func NewIpamRolesUpdateOK() *IpamRolesUpdateOK {
	return &IpamRolesUpdateOK{}
}

/*
IpamRolesUpdateOK describes a response with status code 200, with default header values.

IpamRolesUpdateOK ipam roles update o k
*/
type IpamRolesUpdateOK struct {
	Payload *models.Role
}

// IsSuccess returns true when this ipam roles update o k response has a 2xx status code
func (o *IpamRolesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam roles update o k response has a 3xx status code
func (o *IpamRolesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam roles update o k response has a 4xx status code
func (o *IpamRolesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam roles update o k response has a 5xx status code
func (o *IpamRolesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam roles update o k response a status code equal to that given
func (o *IpamRolesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamRolesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/roles/{id}/][%d] ipamRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRolesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/roles/{id}/][%d] ipamRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRolesUpdateOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRolesUpdateBadRequest creates a IpamRolesUpdateBadRequest with default headers values
func NewIpamRolesUpdateBadRequest() *IpamRolesUpdateBadRequest {
	return &IpamRolesUpdateBadRequest{}
}

/*
IpamRolesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamRolesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam roles update bad request response has a 2xx status code
func (o *IpamRolesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam roles update bad request response has a 3xx status code
func (o *IpamRolesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam roles update bad request response has a 4xx status code
func (o *IpamRolesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam roles update bad request response has a 5xx status code
func (o *IpamRolesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam roles update bad request response a status code equal to that given
func (o *IpamRolesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamRolesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ipam/roles/{id}/][%d] ipamRolesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRolesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /ipam/roles/{id}/][%d] ipamRolesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRolesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRolesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRolesUpdateDefault creates a IpamRolesUpdateDefault with default headers values
func NewIpamRolesUpdateDefault(code int) *IpamRolesUpdateDefault {
	return &IpamRolesUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamRolesUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamRolesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam roles update default response
func (o *IpamRolesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam roles update default response has a 2xx status code
func (o *IpamRolesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam roles update default response has a 3xx status code
func (o *IpamRolesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam roles update default response has a 4xx status code
func (o *IpamRolesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam roles update default response has a 5xx status code
func (o *IpamRolesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam roles update default response a status code equal to that given
func (o *IpamRolesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamRolesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/roles/{id}/][%d] ipam_roles_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRolesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/roles/{id}/][%d] ipam_roles_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRolesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRolesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
