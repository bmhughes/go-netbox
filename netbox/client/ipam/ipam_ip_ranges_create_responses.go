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

// IpamIPRangesCreateReader is a Reader for the IpamIPRangesCreate structure.
type IpamIPRangesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamIPRangesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamIPRangesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamIPRangesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPRangesCreateCreated creates a IpamIPRangesCreateCreated with default headers values
func NewIpamIPRangesCreateCreated() *IpamIPRangesCreateCreated {
	return &IpamIPRangesCreateCreated{}
}

/*
IpamIPRangesCreateCreated describes a response with status code 201, with default header values.

IpamIPRangesCreateCreated ipam Ip ranges create created
*/
type IpamIPRangesCreateCreated struct {
	Payload *models.IPRange
}

// IsSuccess returns true when this ipam Ip ranges create created response has a 2xx status code
func (o *IpamIPRangesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges create created response has a 3xx status code
func (o *IpamIPRangesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges create created response has a 4xx status code
func (o *IpamIPRangesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges create created response has a 5xx status code
func (o *IpamIPRangesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges create created response a status code equal to that given
func (o *IpamIPRangesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamIPRangesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipamIpRangesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamIPRangesCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipamIpRangesCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamIPRangesCreateCreated) GetPayload() *models.IPRange {
	return o.Payload
}

func (o *IpamIPRangesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPRange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPRangesCreateBadRequest creates a IpamIPRangesCreateBadRequest with default headers values
func NewIpamIPRangesCreateBadRequest() *IpamIPRangesCreateBadRequest {
	return &IpamIPRangesCreateBadRequest{}
}

/*
IpamIPRangesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamIPRangesCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam Ip ranges create bad request response has a 2xx status code
func (o *IpamIPRangesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam Ip ranges create bad request response has a 3xx status code
func (o *IpamIPRangesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges create bad request response has a 4xx status code
func (o *IpamIPRangesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam Ip ranges create bad request response has a 5xx status code
func (o *IpamIPRangesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges create bad request response a status code equal to that given
func (o *IpamIPRangesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamIPRangesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipamIpRangesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamIPRangesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipamIpRangesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamIPRangesCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPRangesCreateDefault creates a IpamIPRangesCreateDefault with default headers values
func NewIpamIPRangesCreateDefault(code int) *IpamIPRangesCreateDefault {
	return &IpamIPRangesCreateDefault{
		_statusCode: code,
	}
}

/*
IpamIPRangesCreateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamIPRangesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip ranges create default response
func (o *IpamIPRangesCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam ip ranges create default response has a 2xx status code
func (o *IpamIPRangesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam ip ranges create default response has a 3xx status code
func (o *IpamIPRangesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam ip ranges create default response has a 4xx status code
func (o *IpamIPRangesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam ip ranges create default response has a 5xx status code
func (o *IpamIPRangesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam ip ranges create default response a status code equal to that given
func (o *IpamIPRangesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamIPRangesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipam_ip-ranges_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPRangesCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/ip-ranges/][%d] ipam_ip-ranges_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPRangesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
