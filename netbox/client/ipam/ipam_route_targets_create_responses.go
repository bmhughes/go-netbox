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

// IpamRouteTargetsCreateReader is a Reader for the IpamRouteTargetsCreate structure.
type IpamRouteTargetsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamRouteTargetsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamRouteTargetsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamRouteTargetsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRouteTargetsCreateCreated creates a IpamRouteTargetsCreateCreated with default headers values
func NewIpamRouteTargetsCreateCreated() *IpamRouteTargetsCreateCreated {
	return &IpamRouteTargetsCreateCreated{}
}

/*
IpamRouteTargetsCreateCreated describes a response with status code 201, with default header values.

IpamRouteTargetsCreateCreated ipam route targets create created
*/
type IpamRouteTargetsCreateCreated struct {
	Payload *models.RouteTarget
}

// IsSuccess returns true when this ipam route targets create created response has a 2xx status code
func (o *IpamRouteTargetsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam route targets create created response has a 3xx status code
func (o *IpamRouteTargetsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam route targets create created response has a 4xx status code
func (o *IpamRouteTargetsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam route targets create created response has a 5xx status code
func (o *IpamRouteTargetsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam route targets create created response a status code equal to that given
func (o *IpamRouteTargetsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamRouteTargetsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/route-targets/][%d] ipamRouteTargetsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamRouteTargetsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/route-targets/][%d] ipamRouteTargetsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamRouteTargetsCreateCreated) GetPayload() *models.RouteTarget {
	return o.Payload
}

func (o *IpamRouteTargetsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRouteTargetsCreateBadRequest creates a IpamRouteTargetsCreateBadRequest with default headers values
func NewIpamRouteTargetsCreateBadRequest() *IpamRouteTargetsCreateBadRequest {
	return &IpamRouteTargetsCreateBadRequest{}
}

/*
IpamRouteTargetsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamRouteTargetsCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam route targets create bad request response has a 2xx status code
func (o *IpamRouteTargetsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam route targets create bad request response has a 3xx status code
func (o *IpamRouteTargetsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam route targets create bad request response has a 4xx status code
func (o *IpamRouteTargetsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam route targets create bad request response has a 5xx status code
func (o *IpamRouteTargetsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam route targets create bad request response a status code equal to that given
func (o *IpamRouteTargetsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamRouteTargetsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /ipam/route-targets/][%d] ipamRouteTargetsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRouteTargetsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /ipam/route-targets/][%d] ipamRouteTargetsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRouteTargetsCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRouteTargetsCreateDefault creates a IpamRouteTargetsCreateDefault with default headers values
func NewIpamRouteTargetsCreateDefault(code int) *IpamRouteTargetsCreateDefault {
	return &IpamRouteTargetsCreateDefault{
		_statusCode: code,
	}
}

/*
IpamRouteTargetsCreateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamRouteTargetsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam route targets create default response
func (o *IpamRouteTargetsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam route targets create default response has a 2xx status code
func (o *IpamRouteTargetsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam route targets create default response has a 3xx status code
func (o *IpamRouteTargetsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam route targets create default response has a 4xx status code
func (o *IpamRouteTargetsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam route targets create default response has a 5xx status code
func (o *IpamRouteTargetsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam route targets create default response a status code equal to that given
func (o *IpamRouteTargetsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamRouteTargetsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/route-targets/][%d] ipam_route-targets_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/route-targets/][%d] ipam_route-targets_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
