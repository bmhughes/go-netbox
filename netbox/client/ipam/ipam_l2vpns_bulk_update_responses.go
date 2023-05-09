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

// IpamL2vpnsBulkUpdateReader is a Reader for the IpamL2vpnsBulkUpdate structure.
type IpamL2vpnsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamL2vpnsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamL2vpnsBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamL2vpnsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamL2vpnsBulkUpdateOK creates a IpamL2vpnsBulkUpdateOK with default headers values
func NewIpamL2vpnsBulkUpdateOK() *IpamL2vpnsBulkUpdateOK {
	return &IpamL2vpnsBulkUpdateOK{}
}

/*
IpamL2vpnsBulkUpdateOK describes a response with status code 200, with default header values.

IpamL2vpnsBulkUpdateOK ipam l2vpns bulk update o k
*/
type IpamL2vpnsBulkUpdateOK struct {
	Payload *models.L2VPN
}

// IsSuccess returns true when this ipam l2vpns bulk update o k response has a 2xx status code
func (o *IpamL2vpnsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpns bulk update o k response has a 3xx status code
func (o *IpamL2vpnsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpns bulk update o k response has a 4xx status code
func (o *IpamL2vpnsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpns bulk update o k response has a 5xx status code
func (o *IpamL2vpnsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpns bulk update o k response a status code equal to that given
func (o *IpamL2vpnsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamL2vpnsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/l2vpns/][%d] ipamL2vpnsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/l2vpns/][%d] ipamL2vpnsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamL2vpnsBulkUpdateOK) GetPayload() *models.L2VPN {
	return o.Payload
}

func (o *IpamL2vpnsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.L2VPN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamL2vpnsBulkUpdateBadRequest creates a IpamL2vpnsBulkUpdateBadRequest with default headers values
func NewIpamL2vpnsBulkUpdateBadRequest() *IpamL2vpnsBulkUpdateBadRequest {
	return &IpamL2vpnsBulkUpdateBadRequest{}
}

/*
IpamL2vpnsBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamL2vpnsBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam l2vpns bulk update bad request response has a 2xx status code
func (o *IpamL2vpnsBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam l2vpns bulk update bad request response has a 3xx status code
func (o *IpamL2vpnsBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpns bulk update bad request response has a 4xx status code
func (o *IpamL2vpnsBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam l2vpns bulk update bad request response has a 5xx status code
func (o *IpamL2vpnsBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpns bulk update bad request response a status code equal to that given
func (o *IpamL2vpnsBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamL2vpnsBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ipam/l2vpns/][%d] ipamL2vpnsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamL2vpnsBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /ipam/l2vpns/][%d] ipamL2vpnsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamL2vpnsBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamL2vpnsBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamL2vpnsBulkUpdateDefault creates a IpamL2vpnsBulkUpdateDefault with default headers values
func NewIpamL2vpnsBulkUpdateDefault(code int) *IpamL2vpnsBulkUpdateDefault {
	return &IpamL2vpnsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamL2vpnsBulkUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamL2vpnsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam l2vpns bulk update default response
func (o *IpamL2vpnsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam l2vpns bulk update default response has a 2xx status code
func (o *IpamL2vpnsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam l2vpns bulk update default response has a 3xx status code
func (o *IpamL2vpnsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam l2vpns bulk update default response has a 4xx status code
func (o *IpamL2vpnsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam l2vpns bulk update default response has a 5xx status code
func (o *IpamL2vpnsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam l2vpns bulk update default response a status code equal to that given
func (o *IpamL2vpnsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamL2vpnsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/l2vpns/][%d] ipam_l2vpns_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/l2vpns/][%d] ipam_l2vpns_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamL2vpnsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
