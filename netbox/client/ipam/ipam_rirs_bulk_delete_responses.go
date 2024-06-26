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

// IpamRirsBulkDeleteReader is a Reader for the IpamRirsBulkDelete structure.
type IpamRirsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRirsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRirsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamRirsBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamRirsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamRirsBulkDeleteNoContent creates a IpamRirsBulkDeleteNoContent with default headers values
func NewIpamRirsBulkDeleteNoContent() *IpamRirsBulkDeleteNoContent {
	return &IpamRirsBulkDeleteNoContent{}
}

/*
IpamRirsBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamRirsBulkDeleteNoContent ipam rirs bulk delete no content
*/
type IpamRirsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this ipam rirs bulk delete no content response has a 2xx status code
func (o *IpamRirsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam rirs bulk delete no content response has a 3xx status code
func (o *IpamRirsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs bulk delete no content response has a 4xx status code
func (o *IpamRirsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam rirs bulk delete no content response has a 5xx status code
func (o *IpamRirsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs bulk delete no content response a status code equal to that given
func (o *IpamRirsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamRirsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipamRirsBulkDeleteNoContent ", 204)
}

func (o *IpamRirsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipamRirsBulkDeleteNoContent ", 204)
}

func (o *IpamRirsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamRirsBulkDeleteBadRequest creates a IpamRirsBulkDeleteBadRequest with default headers values
func NewIpamRirsBulkDeleteBadRequest() *IpamRirsBulkDeleteBadRequest {
	return &IpamRirsBulkDeleteBadRequest{}
}

/*
IpamRirsBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamRirsBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam rirs bulk delete bad request response has a 2xx status code
func (o *IpamRirsBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam rirs bulk delete bad request response has a 3xx status code
func (o *IpamRirsBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam rirs bulk delete bad request response has a 4xx status code
func (o *IpamRirsBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam rirs bulk delete bad request response has a 5xx status code
func (o *IpamRirsBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam rirs bulk delete bad request response a status code equal to that given
func (o *IpamRirsBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamRirsBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipamRirsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRirsBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipamRirsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRirsBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRirsBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamRirsBulkDeleteDefault creates a IpamRirsBulkDeleteDefault with default headers values
func NewIpamRirsBulkDeleteDefault(code int) *IpamRirsBulkDeleteDefault {
	return &IpamRirsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamRirsBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamRirsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam rirs bulk delete default response
func (o *IpamRirsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam rirs bulk delete default response has a 2xx status code
func (o *IpamRirsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam rirs bulk delete default response has a 3xx status code
func (o *IpamRirsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam rirs bulk delete default response has a 4xx status code
func (o *IpamRirsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam rirs bulk delete default response has a 5xx status code
func (o *IpamRirsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam rirs bulk delete default response a status code equal to that given
func (o *IpamRirsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamRirsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipam_rirs_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/rirs/][%d] ipam_rirs_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamRirsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRirsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
