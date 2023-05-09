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

// IpamFhrpGroupAssignmentsDeleteReader is a Reader for the IpamFhrpGroupAssignmentsDelete structure.
type IpamFhrpGroupAssignmentsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamFhrpGroupAssignmentsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamFhrpGroupAssignmentsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamFhrpGroupAssignmentsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamFhrpGroupAssignmentsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamFhrpGroupAssignmentsDeleteNoContent creates a IpamFhrpGroupAssignmentsDeleteNoContent with default headers values
func NewIpamFhrpGroupAssignmentsDeleteNoContent() *IpamFhrpGroupAssignmentsDeleteNoContent {
	return &IpamFhrpGroupAssignmentsDeleteNoContent{}
}

/*
IpamFhrpGroupAssignmentsDeleteNoContent describes a response with status code 204, with default header values.

IpamFhrpGroupAssignmentsDeleteNoContent ipam fhrp group assignments delete no content
*/
type IpamFhrpGroupAssignmentsDeleteNoContent struct {
}

// IsSuccess returns true when this ipam fhrp group assignments delete no content response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam fhrp group assignments delete no content response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments delete no content response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam fhrp group assignments delete no content response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments delete no content response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamFhrpGroupAssignmentsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsDeleteNoContent ", 204)
}

func (o *IpamFhrpGroupAssignmentsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsDeleteNoContent ", 204)
}

func (o *IpamFhrpGroupAssignmentsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamFhrpGroupAssignmentsDeleteBadRequest creates a IpamFhrpGroupAssignmentsDeleteBadRequest with default headers values
func NewIpamFhrpGroupAssignmentsDeleteBadRequest() *IpamFhrpGroupAssignmentsDeleteBadRequest {
	return &IpamFhrpGroupAssignmentsDeleteBadRequest{}
}

/*
IpamFhrpGroupAssignmentsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamFhrpGroupAssignmentsDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam fhrp group assignments delete bad request response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam fhrp group assignments delete bad request response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam fhrp group assignments delete bad request response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam fhrp group assignments delete bad request response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam fhrp group assignments delete bad request response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamFhrpGroupAssignmentsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-group-assignments/{id}/][%d] ipamFhrpGroupAssignmentsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamFhrpGroupAssignmentsDeleteDefault creates a IpamFhrpGroupAssignmentsDeleteDefault with default headers values
func NewIpamFhrpGroupAssignmentsDeleteDefault(code int) *IpamFhrpGroupAssignmentsDeleteDefault {
	return &IpamFhrpGroupAssignmentsDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamFhrpGroupAssignmentsDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamFhrpGroupAssignmentsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam fhrp group assignments delete default response
func (o *IpamFhrpGroupAssignmentsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam fhrp group assignments delete default response has a 2xx status code
func (o *IpamFhrpGroupAssignmentsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam fhrp group assignments delete default response has a 3xx status code
func (o *IpamFhrpGroupAssignmentsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam fhrp group assignments delete default response has a 4xx status code
func (o *IpamFhrpGroupAssignmentsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam fhrp group assignments delete default response has a 5xx status code
func (o *IpamFhrpGroupAssignmentsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam fhrp group assignments delete default response a status code equal to that given
func (o *IpamFhrpGroupAssignmentsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamFhrpGroupAssignmentsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-group-assignments/{id}/][%d] ipam_fhrp-group-assignments_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/fhrp-group-assignments/{id}/][%d] ipam_fhrp-group-assignments_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamFhrpGroupAssignmentsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamFhrpGroupAssignmentsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
