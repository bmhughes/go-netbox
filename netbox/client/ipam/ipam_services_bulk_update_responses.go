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

// IpamServicesBulkUpdateReader is a Reader for the IpamServicesBulkUpdate structure.
type IpamServicesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamServicesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamServicesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamServicesBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamServicesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamServicesBulkUpdateOK creates a IpamServicesBulkUpdateOK with default headers values
func NewIpamServicesBulkUpdateOK() *IpamServicesBulkUpdateOK {
	return &IpamServicesBulkUpdateOK{}
}

/*
IpamServicesBulkUpdateOK describes a response with status code 200, with default header values.

IpamServicesBulkUpdateOK ipam services bulk update o k
*/
type IpamServicesBulkUpdateOK struct {
	Payload *models.Service
}

// IsSuccess returns true when this ipam services bulk update o k response has a 2xx status code
func (o *IpamServicesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam services bulk update o k response has a 3xx status code
func (o *IpamServicesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam services bulk update o k response has a 4xx status code
func (o *IpamServicesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam services bulk update o k response has a 5xx status code
func (o *IpamServicesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam services bulk update o k response a status code equal to that given
func (o *IpamServicesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamServicesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/services/][%d] ipamServicesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServicesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/services/][%d] ipamServicesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamServicesBulkUpdateOK) GetPayload() *models.Service {
	return o.Payload
}

func (o *IpamServicesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Service)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServicesBulkUpdateBadRequest creates a IpamServicesBulkUpdateBadRequest with default headers values
func NewIpamServicesBulkUpdateBadRequest() *IpamServicesBulkUpdateBadRequest {
	return &IpamServicesBulkUpdateBadRequest{}
}

/*
IpamServicesBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamServicesBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam services bulk update bad request response has a 2xx status code
func (o *IpamServicesBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam services bulk update bad request response has a 3xx status code
func (o *IpamServicesBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam services bulk update bad request response has a 4xx status code
func (o *IpamServicesBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam services bulk update bad request response has a 5xx status code
func (o *IpamServicesBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam services bulk update bad request response a status code equal to that given
func (o *IpamServicesBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamServicesBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ipam/services/][%d] ipamServicesBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamServicesBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /ipam/services/][%d] ipamServicesBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamServicesBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServicesBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamServicesBulkUpdateDefault creates a IpamServicesBulkUpdateDefault with default headers values
func NewIpamServicesBulkUpdateDefault(code int) *IpamServicesBulkUpdateDefault {
	return &IpamServicesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamServicesBulkUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamServicesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam services bulk update default response
func (o *IpamServicesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam services bulk update default response has a 2xx status code
func (o *IpamServicesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam services bulk update default response has a 3xx status code
func (o *IpamServicesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam services bulk update default response has a 4xx status code
func (o *IpamServicesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam services bulk update default response has a 5xx status code
func (o *IpamServicesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam services bulk update default response a status code equal to that given
func (o *IpamServicesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamServicesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/services/][%d] ipam_services_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServicesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/services/][%d] ipam_services_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamServicesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamServicesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
