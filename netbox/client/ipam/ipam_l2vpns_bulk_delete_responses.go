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

// IpamL2vpnsBulkDeleteReader is a Reader for the IpamL2vpnsBulkDelete structure.
type IpamL2vpnsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamL2vpnsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamL2vpnsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamL2vpnsBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamL2vpnsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamL2vpnsBulkDeleteNoContent creates a IpamL2vpnsBulkDeleteNoContent with default headers values
func NewIpamL2vpnsBulkDeleteNoContent() *IpamL2vpnsBulkDeleteNoContent {
	return &IpamL2vpnsBulkDeleteNoContent{}
}

/*
IpamL2vpnsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamL2vpnsBulkDeleteNoContent ipam l2vpns bulk delete no content
*/
type IpamL2vpnsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this ipam l2vpns bulk delete no content response has a 2xx status code
func (o *IpamL2vpnsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam l2vpns bulk delete no content response has a 3xx status code
func (o *IpamL2vpnsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpns bulk delete no content response has a 4xx status code
func (o *IpamL2vpnsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam l2vpns bulk delete no content response has a 5xx status code
func (o *IpamL2vpnsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpns bulk delete no content response a status code equal to that given
func (o *IpamL2vpnsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamL2vpnsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/l2vpns/][%d] ipamL2vpnsBulkDeleteNoContent ", 204)
}

func (o *IpamL2vpnsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/l2vpns/][%d] ipamL2vpnsBulkDeleteNoContent ", 204)
}

func (o *IpamL2vpnsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamL2vpnsBulkDeleteBadRequest creates a IpamL2vpnsBulkDeleteBadRequest with default headers values
func NewIpamL2vpnsBulkDeleteBadRequest() *IpamL2vpnsBulkDeleteBadRequest {
	return &IpamL2vpnsBulkDeleteBadRequest{}
}

/*
IpamL2vpnsBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamL2vpnsBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam l2vpns bulk delete bad request response has a 2xx status code
func (o *IpamL2vpnsBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam l2vpns bulk delete bad request response has a 3xx status code
func (o *IpamL2vpnsBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam l2vpns bulk delete bad request response has a 4xx status code
func (o *IpamL2vpnsBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam l2vpns bulk delete bad request response has a 5xx status code
func (o *IpamL2vpnsBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam l2vpns bulk delete bad request response a status code equal to that given
func (o *IpamL2vpnsBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamL2vpnsBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ipam/l2vpns/][%d] ipamL2vpnsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamL2vpnsBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /ipam/l2vpns/][%d] ipamL2vpnsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamL2vpnsBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamL2vpnsBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamL2vpnsBulkDeleteDefault creates a IpamL2vpnsBulkDeleteDefault with default headers values
func NewIpamL2vpnsBulkDeleteDefault(code int) *IpamL2vpnsBulkDeleteDefault {
	return &IpamL2vpnsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamL2vpnsBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamL2vpnsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam l2vpns bulk delete default response
func (o *IpamL2vpnsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam l2vpns bulk delete default response has a 2xx status code
func (o *IpamL2vpnsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam l2vpns bulk delete default response has a 3xx status code
func (o *IpamL2vpnsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam l2vpns bulk delete default response has a 4xx status code
func (o *IpamL2vpnsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam l2vpns bulk delete default response has a 5xx status code
func (o *IpamL2vpnsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam l2vpns bulk delete default response a status code equal to that given
func (o *IpamL2vpnsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamL2vpnsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/l2vpns/][%d] ipam_l2vpns_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/l2vpns/][%d] ipam_l2vpns_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamL2vpnsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamL2vpnsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
