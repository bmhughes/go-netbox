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

// IpamIPAddressesPartialUpdateReader is a Reader for the IpamIPAddressesPartialUpdate structure.
type IpamIPAddressesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPAddressesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamIPAddressesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamIPAddressesPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamIPAddressesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPAddressesPartialUpdateOK creates a IpamIPAddressesPartialUpdateOK with default headers values
func NewIpamIPAddressesPartialUpdateOK() *IpamIPAddressesPartialUpdateOK {
	return &IpamIPAddressesPartialUpdateOK{}
}

/*
IpamIPAddressesPartialUpdateOK describes a response with status code 200, with default header values.

IpamIPAddressesPartialUpdateOK ipam Ip addresses partial update o k
*/
type IpamIPAddressesPartialUpdateOK struct {
	Payload *models.IPAddress
}

// IsSuccess returns true when this ipam Ip addresses partial update o k response has a 2xx status code
func (o *IpamIPAddressesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip addresses partial update o k response has a 3xx status code
func (o *IpamIPAddressesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip addresses partial update o k response has a 4xx status code
func (o *IpamIPAddressesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip addresses partial update o k response has a 5xx status code
func (o *IpamIPAddressesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip addresses partial update o k response a status code equal to that given
func (o *IpamIPAddressesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamIPAddressesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipamIpAddressesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPAddressesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipamIpAddressesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamIPAddressesPartialUpdateOK) GetPayload() *models.IPAddress {
	return o.Payload
}

func (o *IpamIPAddressesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPAddressesPartialUpdateBadRequest creates a IpamIPAddressesPartialUpdateBadRequest with default headers values
func NewIpamIPAddressesPartialUpdateBadRequest() *IpamIPAddressesPartialUpdateBadRequest {
	return &IpamIPAddressesPartialUpdateBadRequest{}
}

/*
IpamIPAddressesPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamIPAddressesPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam Ip addresses partial update bad request response has a 2xx status code
func (o *IpamIPAddressesPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam Ip addresses partial update bad request response has a 3xx status code
func (o *IpamIPAddressesPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip addresses partial update bad request response has a 4xx status code
func (o *IpamIPAddressesPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam Ip addresses partial update bad request response has a 5xx status code
func (o *IpamIPAddressesPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip addresses partial update bad request response a status code equal to that given
func (o *IpamIPAddressesPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamIPAddressesPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipamIpAddressesPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamIPAddressesPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipamIpAddressesPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamIPAddressesPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPAddressesPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPAddressesPartialUpdateDefault creates a IpamIPAddressesPartialUpdateDefault with default headers values
func NewIpamIPAddressesPartialUpdateDefault(code int) *IpamIPAddressesPartialUpdateDefault {
	return &IpamIPAddressesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamIPAddressesPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamIPAddressesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip addresses partial update default response
func (o *IpamIPAddressesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam ip addresses partial update default response has a 2xx status code
func (o *IpamIPAddressesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam ip addresses partial update default response has a 3xx status code
func (o *IpamIPAddressesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam ip addresses partial update default response has a 4xx status code
func (o *IpamIPAddressesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam ip addresses partial update default response has a 5xx status code
func (o *IpamIPAddressesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam ip addresses partial update default response a status code equal to that given
func (o *IpamIPAddressesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamIPAddressesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipam_ip-addresses_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPAddressesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/ip-addresses/{id}/][%d] ipam_ip-addresses_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPAddressesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPAddressesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
