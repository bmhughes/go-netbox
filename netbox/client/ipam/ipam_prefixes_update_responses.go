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

// IpamPrefixesUpdateReader is a Reader for the IpamPrefixesUpdate structure.
type IpamPrefixesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamPrefixesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamPrefixesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamPrefixesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamPrefixesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamPrefixesUpdateOK creates a IpamPrefixesUpdateOK with default headers values
func NewIpamPrefixesUpdateOK() *IpamPrefixesUpdateOK {
	return &IpamPrefixesUpdateOK{}
}

/*
IpamPrefixesUpdateOK describes a response with status code 200, with default header values.

IpamPrefixesUpdateOK ipam prefixes update o k
*/
type IpamPrefixesUpdateOK struct {
	Payload *models.Prefix
}

// IsSuccess returns true when this ipam prefixes update o k response has a 2xx status code
func (o *IpamPrefixesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam prefixes update o k response has a 3xx status code
func (o *IpamPrefixesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes update o k response has a 4xx status code
func (o *IpamPrefixesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam prefixes update o k response has a 5xx status code
func (o *IpamPrefixesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes update o k response a status code equal to that given
func (o *IpamPrefixesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamPrefixesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/prefixes/{id}/][%d] ipamPrefixesUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamPrefixesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/prefixes/{id}/][%d] ipamPrefixesUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamPrefixesUpdateOK) GetPayload() *models.Prefix {
	return o.Payload
}

func (o *IpamPrefixesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Prefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesUpdateBadRequest creates a IpamPrefixesUpdateBadRequest with default headers values
func NewIpamPrefixesUpdateBadRequest() *IpamPrefixesUpdateBadRequest {
	return &IpamPrefixesUpdateBadRequest{}
}

/*
IpamPrefixesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamPrefixesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam prefixes update bad request response has a 2xx status code
func (o *IpamPrefixesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam prefixes update bad request response has a 3xx status code
func (o *IpamPrefixesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam prefixes update bad request response has a 4xx status code
func (o *IpamPrefixesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam prefixes update bad request response has a 5xx status code
func (o *IpamPrefixesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam prefixes update bad request response a status code equal to that given
func (o *IpamPrefixesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamPrefixesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ipam/prefixes/{id}/][%d] ipamPrefixesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamPrefixesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /ipam/prefixes/{id}/][%d] ipamPrefixesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamPrefixesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamPrefixesUpdateDefault creates a IpamPrefixesUpdateDefault with default headers values
func NewIpamPrefixesUpdateDefault(code int) *IpamPrefixesUpdateDefault {
	return &IpamPrefixesUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamPrefixesUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamPrefixesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam prefixes update default response
func (o *IpamPrefixesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam prefixes update default response has a 2xx status code
func (o *IpamPrefixesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam prefixes update default response has a 3xx status code
func (o *IpamPrefixesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam prefixes update default response has a 4xx status code
func (o *IpamPrefixesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam prefixes update default response has a 5xx status code
func (o *IpamPrefixesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam prefixes update default response a status code equal to that given
func (o *IpamPrefixesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamPrefixesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/prefixes/{id}/][%d] ipam_prefixes_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/prefixes/{id}/][%d] ipam_prefixes_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamPrefixesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamPrefixesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
