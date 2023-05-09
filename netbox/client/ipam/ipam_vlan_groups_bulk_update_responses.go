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

// IpamVlanGroupsBulkUpdateReader is a Reader for the IpamVlanGroupsBulkUpdate structure.
type IpamVlanGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVlanGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamVlanGroupsBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamVlanGroupsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamVlanGroupsBulkUpdateOK creates a IpamVlanGroupsBulkUpdateOK with default headers values
func NewIpamVlanGroupsBulkUpdateOK() *IpamVlanGroupsBulkUpdateOK {
	return &IpamVlanGroupsBulkUpdateOK{}
}

/*
IpamVlanGroupsBulkUpdateOK describes a response with status code 200, with default header values.

IpamVlanGroupsBulkUpdateOK ipam vlan groups bulk update o k
*/
type IpamVlanGroupsBulkUpdateOK struct {
	Payload *models.VLANGroup
}

// IsSuccess returns true when this ipam vlan groups bulk update o k response has a 2xx status code
func (o *IpamVlanGroupsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups bulk update o k response has a 3xx status code
func (o *IpamVlanGroupsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups bulk update o k response has a 4xx status code
func (o *IpamVlanGroupsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups bulk update o k response has a 5xx status code
func (o *IpamVlanGroupsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups bulk update o k response a status code equal to that given
func (o *IpamVlanGroupsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamVlanGroupsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/][%d] ipamVlanGroupsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/][%d] ipamVlanGroupsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVlanGroupsBulkUpdateOK) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsBulkUpdateBadRequest creates a IpamVlanGroupsBulkUpdateBadRequest with default headers values
func NewIpamVlanGroupsBulkUpdateBadRequest() *IpamVlanGroupsBulkUpdateBadRequest {
	return &IpamVlanGroupsBulkUpdateBadRequest{}
}

/*
IpamVlanGroupsBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamVlanGroupsBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam vlan groups bulk update bad request response has a 2xx status code
func (o *IpamVlanGroupsBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam vlan groups bulk update bad request response has a 3xx status code
func (o *IpamVlanGroupsBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups bulk update bad request response has a 4xx status code
func (o *IpamVlanGroupsBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam vlan groups bulk update bad request response has a 5xx status code
func (o *IpamVlanGroupsBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups bulk update bad request response a status code equal to that given
func (o *IpamVlanGroupsBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamVlanGroupsBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/][%d] ipamVlanGroupsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamVlanGroupsBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/][%d] ipamVlanGroupsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamVlanGroupsBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVlanGroupsBulkUpdateDefault creates a IpamVlanGroupsBulkUpdateDefault with default headers values
func NewIpamVlanGroupsBulkUpdateDefault(code int) *IpamVlanGroupsBulkUpdateDefault {
	return &IpamVlanGroupsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamVlanGroupsBulkUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamVlanGroupsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam vlan groups bulk update default response
func (o *IpamVlanGroupsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam vlan groups bulk update default response has a 2xx status code
func (o *IpamVlanGroupsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam vlan groups bulk update default response has a 3xx status code
func (o *IpamVlanGroupsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam vlan groups bulk update default response has a 4xx status code
func (o *IpamVlanGroupsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam vlan groups bulk update default response has a 5xx status code
func (o *IpamVlanGroupsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam vlan groups bulk update default response a status code equal to that given
func (o *IpamVlanGroupsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamVlanGroupsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/][%d] ipam_vlan-groups_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/vlan-groups/][%d] ipam_vlan-groups_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamVlanGroupsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVlanGroupsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
