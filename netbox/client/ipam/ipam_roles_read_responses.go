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

// IpamRolesReadReader is a Reader for the IpamRolesRead structure.
type IpamRolesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRolesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRolesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamRolesReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamRolesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRolesReadOK creates a IpamRolesReadOK with default headers values
func NewIpamRolesReadOK() *IpamRolesReadOK {
	return &IpamRolesReadOK{}
}

/*
IpamRolesReadOK describes a response with status code 200, with default header values.

IpamRolesReadOK ipam roles read o k
*/
type IpamRolesReadOK struct {
	Payload *models.Role
}

// IsSuccess returns true when this ipam roles read o k response has a 2xx status code
func (o *IpamRolesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam roles read o k response has a 3xx status code
func (o *IpamRolesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam roles read o k response has a 4xx status code
func (o *IpamRolesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam roles read o k response has a 5xx status code
func (o *IpamRolesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam roles read o k response a status code equal to that given
func (o *IpamRolesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamRolesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipamRolesReadOK  %+v", 200, o.Payload)
}

func (o *IpamRolesReadOK) String() string {
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipamRolesReadOK  %+v", 200, o.Payload)
}

func (o *IpamRolesReadOK) GetPayload() *models.Role {
	return o.Payload
}

func (o *IpamRolesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRolesReadBadRequest creates a IpamRolesReadBadRequest with default headers values
func NewIpamRolesReadBadRequest() *IpamRolesReadBadRequest {
	return &IpamRolesReadBadRequest{}
}

/*
IpamRolesReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamRolesReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam roles read bad request response has a 2xx status code
func (o *IpamRolesReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam roles read bad request response has a 3xx status code
func (o *IpamRolesReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam roles read bad request response has a 4xx status code
func (o *IpamRolesReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam roles read bad request response has a 5xx status code
func (o *IpamRolesReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam roles read bad request response a status code equal to that given
func (o *IpamRolesReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamRolesReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipamRolesReadBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRolesReadBadRequest) String() string {
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipamRolesReadBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRolesReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRolesReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRolesReadDefault creates a IpamRolesReadDefault with default headers values
func NewIpamRolesReadDefault(code int) *IpamRolesReadDefault {
	return &IpamRolesReadDefault{
		_statusCode: code,
	}
}

/*
IpamRolesReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamRolesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam roles read default response
func (o *IpamRolesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam roles read default response has a 2xx status code
func (o *IpamRolesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam roles read default response has a 3xx status code
func (o *IpamRolesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam roles read default response has a 4xx status code
func (o *IpamRolesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam roles read default response has a 5xx status code
func (o *IpamRolesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam roles read default response a status code equal to that given
func (o *IpamRolesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamRolesReadDefault) Error() string {
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipam_roles_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRolesReadDefault) String() string {
	return fmt.Sprintf("[GET /ipam/roles/{id}/][%d] ipam_roles_read default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRolesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRolesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
