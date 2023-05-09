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

// IpamAsnsBulkPartialUpdateReader is a Reader for the IpamAsnsBulkPartialUpdate structure.
type IpamAsnsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAsnsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamAsnsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamAsnsBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamAsnsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamAsnsBulkPartialUpdateOK creates a IpamAsnsBulkPartialUpdateOK with default headers values
func NewIpamAsnsBulkPartialUpdateOK() *IpamAsnsBulkPartialUpdateOK {
	return &IpamAsnsBulkPartialUpdateOK{}
}

/*
IpamAsnsBulkPartialUpdateOK describes a response with status code 200, with default header values.

IpamAsnsBulkPartialUpdateOK ipam asns bulk partial update o k
*/
type IpamAsnsBulkPartialUpdateOK struct {
	Payload *models.ASN
}

// IsSuccess returns true when this ipam asns bulk partial update o k response has a 2xx status code
func (o *IpamAsnsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam asns bulk partial update o k response has a 3xx status code
func (o *IpamAsnsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam asns bulk partial update o k response has a 4xx status code
func (o *IpamAsnsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam asns bulk partial update o k response has a 5xx status code
func (o *IpamAsnsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam asns bulk partial update o k response a status code equal to that given
func (o *IpamAsnsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamAsnsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/asns/][%d] ipamAsnsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAsnsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/asns/][%d] ipamAsnsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamAsnsBulkPartialUpdateOK) GetPayload() *models.ASN {
	return o.Payload
}

func (o *IpamAsnsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ASN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAsnsBulkPartialUpdateBadRequest creates a IpamAsnsBulkPartialUpdateBadRequest with default headers values
func NewIpamAsnsBulkPartialUpdateBadRequest() *IpamAsnsBulkPartialUpdateBadRequest {
	return &IpamAsnsBulkPartialUpdateBadRequest{}
}

/*
IpamAsnsBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamAsnsBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam asns bulk partial update bad request response has a 2xx status code
func (o *IpamAsnsBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam asns bulk partial update bad request response has a 3xx status code
func (o *IpamAsnsBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam asns bulk partial update bad request response has a 4xx status code
func (o *IpamAsnsBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam asns bulk partial update bad request response has a 5xx status code
func (o *IpamAsnsBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam asns bulk partial update bad request response a status code equal to that given
func (o *IpamAsnsBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamAsnsBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /ipam/asns/][%d] ipamAsnsBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamAsnsBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /ipam/asns/][%d] ipamAsnsBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamAsnsBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAsnsBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAsnsBulkPartialUpdateDefault creates a IpamAsnsBulkPartialUpdateDefault with default headers values
func NewIpamAsnsBulkPartialUpdateDefault(code int) *IpamAsnsBulkPartialUpdateDefault {
	return &IpamAsnsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamAsnsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamAsnsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam asns bulk partial update default response
func (o *IpamAsnsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam asns bulk partial update default response has a 2xx status code
func (o *IpamAsnsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam asns bulk partial update default response has a 3xx status code
func (o *IpamAsnsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam asns bulk partial update default response has a 4xx status code
func (o *IpamAsnsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam asns bulk partial update default response has a 5xx status code
func (o *IpamAsnsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam asns bulk partial update default response a status code equal to that given
func (o *IpamAsnsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamAsnsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /ipam/asns/][%d] ipam_asns_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAsnsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /ipam/asns/][%d] ipam_asns_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAsnsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAsnsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
