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
)

// IpamIPRangesDeleteReader is a Reader for the IpamIPRangesDelete structure.
type IpamIPRangesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamIPRangesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamIPRangesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamIPRangesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamIPRangesDeleteNoContent creates a IpamIPRangesDeleteNoContent with default headers values
func NewIpamIPRangesDeleteNoContent() *IpamIPRangesDeleteNoContent {
	return &IpamIPRangesDeleteNoContent{}
}

/*
IpamIPRangesDeleteNoContent describes a response with status code 204, with default header values.

IpamIPRangesDeleteNoContent ipam Ip ranges delete no content
*/
type IpamIPRangesDeleteNoContent struct {
}

// IsSuccess returns true when this ipam Ip ranges delete no content response has a 2xx status code
func (o *IpamIPRangesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges delete no content response has a 3xx status code
func (o *IpamIPRangesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges delete no content response has a 4xx status code
func (o *IpamIPRangesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges delete no content response has a 5xx status code
func (o *IpamIPRangesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges delete no content response a status code equal to that given
func (o *IpamIPRangesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamIPRangesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-ranges/{id}/][%d] ipamIpRangesDeleteNoContent ", 204)
}

func (o *IpamIPRangesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/ip-ranges/{id}/][%d] ipamIpRangesDeleteNoContent ", 204)
}

func (o *IpamIPRangesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamIPRangesDeleteBadRequest creates a IpamIPRangesDeleteBadRequest with default headers values
func NewIpamIPRangesDeleteBadRequest() *IpamIPRangesDeleteBadRequest {
	return &IpamIPRangesDeleteBadRequest{}
}

/*
IpamIPRangesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamIPRangesDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam Ip ranges delete bad request response has a 2xx status code
func (o *IpamIPRangesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam Ip ranges delete bad request response has a 3xx status code
func (o *IpamIPRangesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges delete bad request response has a 4xx status code
func (o *IpamIPRangesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam Ip ranges delete bad request response has a 5xx status code
func (o *IpamIPRangesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges delete bad request response a status code equal to that given
func (o *IpamIPRangesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamIPRangesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-ranges/{id}/][%d] ipamIpRangesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamIPRangesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /ipam/ip-ranges/{id}/][%d] ipamIpRangesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamIPRangesDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamIPRangesDeleteDefault creates a IpamIPRangesDeleteDefault with default headers values
func NewIpamIPRangesDeleteDefault(code int) *IpamIPRangesDeleteDefault {
	return &IpamIPRangesDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamIPRangesDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamIPRangesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam ip ranges delete default response
func (o *IpamIPRangesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam ip ranges delete default response has a 2xx status code
func (o *IpamIPRangesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam ip ranges delete default response has a 3xx status code
func (o *IpamIPRangesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam ip ranges delete default response has a 4xx status code
func (o *IpamIPRangesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam ip ranges delete default response has a 5xx status code
func (o *IpamIPRangesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam ip ranges delete default response a status code equal to that given
func (o *IpamIPRangesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamIPRangesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-ranges/{id}/][%d] ipam_ip-ranges_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPRangesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/ip-ranges/{id}/][%d] ipam_ip-ranges_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamIPRangesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamIPRangesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
