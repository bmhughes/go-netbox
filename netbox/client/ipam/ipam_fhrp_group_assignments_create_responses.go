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

// IpamFhrpGroupAssignmentsCreateReader is a Reader for the IpamFhrpGroupAssignmentsCreate structure.
type IpamFhrpGroupAssignmentsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamFhrpGroupAssignmentsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamFhrpGroupAssignmentsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamFhrpGroupAssignmentsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupAssignmentsCreateCreated creates a IpamFhrpGroupAssignmentsCreateCreated with default headers values
func NewIpamFhrpGroupAssignmentsCreateCreated() *IpamFhrpGroupAssignmentsCreateCreated {
	return &IpamFhrpGroupAssignmentsCreateCreated{}
}

/*
IpamFhrpGroupAssignmentsCreateCreated describes a response with status code 201, with default header values.

IpamFhrpGroupAssignmentsCreateCreated ipam fhrp group assignments create created
*/
type IpamFhrpGroupAssignmentsCreateCreated struct {
	Payload *models.FHRPGroupAssignment
}

// IsSuccess returns true when this ipam fhrp group assignments create created response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp group assignments create created response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments create created response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp group assignments create created response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments create created response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) GetPayload() *models.FHRPGroupAssignment {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FHRPGroupAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupAssignmentsCreateBadRequest creates a IpamFhrpGroupAssignmentsCreateBadRequest with default headers values
func NewIpamFhrpGroupAssignmentsCreateBadRequest() *IpamFhrpGroupAssignmentsCreateBadRequest {
	return &IpamFhrpGroupAssignmentsCreateBadRequest{}
}

/*
IpamFhrpGroupAssignmentsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamFhrpGroupAssignmentsCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam fhrp group assignments create bad request response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam fhrp group assignments create bad request response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments create bad request response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam fhrp group assignments create bad request response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments create bad request response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamFhrpGroupAssignmentsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipamFhrpGroupAssignmentsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupAssignmentsCreateDefault creates a IpamFhrpGroupAssignmentsCreateDefault with default headers values
func NewIpamFhrpGroupAssignmentsCreateDefault(code int) *IpamFhrpGroupAssignmentsCreateDefault {
	return &IpamFhrpGroupAssignmentsCreateDefault{
		_statusCode: code,
	}
}

/*
IpamFhrpGroupAssignmentsCreateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamFhrpGroupAssignmentsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam fhrp group assignments create default response
func (o *IpamFhrpGroupAssignmentsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam fhrp group assignments create default response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam fhrp group assignments create default response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam fhrp group assignments create default response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam fhrp group assignments create default response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam fhrp group assignments create default response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamFhrpGroupAssignmentsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipam_fhrp-group-assignments_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/fhrp-group-assignments/][%d] ipam_fhrp-group-assignments_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
