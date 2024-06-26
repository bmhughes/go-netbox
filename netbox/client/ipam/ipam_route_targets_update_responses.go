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

// IpamRouteTargetsUpdateReader is a Reader for the IpamRouteTargetsUpdate structure.
type IpamRouteTargetsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamRouteTargetsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamRouteTargetsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamRouteTargetsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRouteTargetsUpdateOK creates a IpamRouteTargetsUpdateOK with default headers values
func NewIpamRouteTargetsUpdateOK() *IpamRouteTargetsUpdateOK {
	return &IpamRouteTargetsUpdateOK{}
}

/*
IpamRouteTargetsUpdateOK describes a response with status code 200, with default header values.

IpamRouteTargetsUpdateOK ipam route targets update o k
*/
type IpamRouteTargetsUpdateOK struct {
	Payload *models.RouteTarget
}

// IsSuccess returns true when this ipam route targets update o k response has a 2xx status code
func (o *IpamRouteTargetsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam route targets update o k response has a 3xx status code
func (o *IpamRouteTargetsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam route targets update o k response has a 4xx status code
func (o *IpamRouteTargetsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam route targets update o k response has a 5xx status code
func (o *IpamRouteTargetsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam route targets update o k response a status code equal to that given
func (o *IpamRouteTargetsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamRouteTargetsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/route-targets/{id}/][%d] ipamRouteTargetsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRouteTargetsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ipam/route-targets/{id}/][%d] ipamRouteTargetsUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamRouteTargetsUpdateOK) GetPayload() *models.RouteTarget {
	return o.Payload
}

func (o *IpamRouteTargetsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RouteTarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRouteTargetsUpdateBadRequest creates a IpamRouteTargetsUpdateBadRequest with default headers values
func NewIpamRouteTargetsUpdateBadRequest() *IpamRouteTargetsUpdateBadRequest {
	return &IpamRouteTargetsUpdateBadRequest{}
}

/*
IpamRouteTargetsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamRouteTargetsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam route targets update bad request response has a 2xx status code
func (o *IpamRouteTargetsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam route targets update bad request response has a 3xx status code
func (o *IpamRouteTargetsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam route targets update bad request response has a 4xx status code
func (o *IpamRouteTargetsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam route targets update bad request response has a 5xx status code
func (o *IpamRouteTargetsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam route targets update bad request response a status code equal to that given
func (o *IpamRouteTargetsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamRouteTargetsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ipam/route-targets/{id}/][%d] ipamRouteTargetsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRouteTargetsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /ipam/route-targets/{id}/][%d] ipamRouteTargetsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRouteTargetsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRouteTargetsUpdateDefault creates a IpamRouteTargetsUpdateDefault with default headers values
func NewIpamRouteTargetsUpdateDefault(code int) *IpamRouteTargetsUpdateDefault {
	return &IpamRouteTargetsUpdateDefault{
		_statusCode: code,
	}
}

/*
IpamRouteTargetsUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamRouteTargetsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam route targets update default response
func (o *IpamRouteTargetsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam route targets update default response has a 2xx status code
func (o *IpamRouteTargetsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam route targets update default response has a 3xx status code
func (o *IpamRouteTargetsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam route targets update default response has a 4xx status code
func (o *IpamRouteTargetsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam route targets update default response has a 5xx status code
func (o *IpamRouteTargetsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam route targets update default response a status code equal to that given
func (o *IpamRouteTargetsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamRouteTargetsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /ipam/route-targets/{id}/][%d] ipam_route-targets_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /ipam/route-targets/{id}/][%d] ipam_route-targets_update default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRouteTargetsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
