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

// IpamAggregatesBulkDeleteReader is a Reader for the IpamAggregatesBulkDelete structure.
type IpamAggregatesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAggregatesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamAggregatesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamAggregatesBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewIpamAggregatesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamAggregatesBulkDeleteNoContent creates a IpamAggregatesBulkDeleteNoContent with default headers values
func NewIpamAggregatesBulkDeleteNoContent() *IpamAggregatesBulkDeleteNoContent {
	return &IpamAggregatesBulkDeleteNoContent{}
}

/*
IpamAggregatesBulkDeleteNoContent describes a response with status code 204, with default header values.

IpamAggregatesBulkDeleteNoContent ipam aggregates bulk delete no content
*/
type IpamAggregatesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this ipam aggregates bulk delete no content response has a 2xx status code
func (o *IpamAggregatesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam aggregates bulk delete no content response has a 3xx status code
func (o *IpamAggregatesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam aggregates bulk delete no content response has a 4xx status code
func (o *IpamAggregatesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam aggregates bulk delete no content response has a 5xx status code
func (o *IpamAggregatesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam aggregates bulk delete no content response a status code equal to that given
func (o *IpamAggregatesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamAggregatesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/aggregates/][%d] ipamAggregatesBulkDeleteNoContent ", 204)
}

func (o *IpamAggregatesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/aggregates/][%d] ipamAggregatesBulkDeleteNoContent ", 204)
}

func (o *IpamAggregatesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamAggregatesBulkDeleteBadRequest creates a IpamAggregatesBulkDeleteBadRequest with default headers values
func NewIpamAggregatesBulkDeleteBadRequest() *IpamAggregatesBulkDeleteBadRequest {
	return &IpamAggregatesBulkDeleteBadRequest{}
}

/*
IpamAggregatesBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamAggregatesBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam aggregates bulk delete bad request response has a 2xx status code
func (o *IpamAggregatesBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam aggregates bulk delete bad request response has a 3xx status code
func (o *IpamAggregatesBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam aggregates bulk delete bad request response has a 4xx status code
func (o *IpamAggregatesBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam aggregates bulk delete bad request response has a 5xx status code
func (o *IpamAggregatesBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam aggregates bulk delete bad request response a status code equal to that given
func (o *IpamAggregatesBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamAggregatesBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ipam/aggregates/][%d] ipamAggregatesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamAggregatesBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /ipam/aggregates/][%d] ipamAggregatesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamAggregatesBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAggregatesBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAggregatesBulkDeleteDefault creates a IpamAggregatesBulkDeleteDefault with default headers values
func NewIpamAggregatesBulkDeleteDefault(code int) *IpamAggregatesBulkDeleteDefault {
	return &IpamAggregatesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
IpamAggregatesBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type IpamAggregatesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam aggregates bulk delete default response
func (o *IpamAggregatesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam aggregates bulk delete default response has a 2xx status code
func (o *IpamAggregatesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam aggregates bulk delete default response has a 3xx status code
func (o *IpamAggregatesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam aggregates bulk delete default response has a 4xx status code
func (o *IpamAggregatesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam aggregates bulk delete default response has a 5xx status code
func (o *IpamAggregatesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam aggregates bulk delete default response a status code equal to that given
func (o *IpamAggregatesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamAggregatesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /ipam/aggregates/][%d] ipam_aggregates_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /ipam/aggregates/][%d] ipam_aggregates_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAggregatesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAggregatesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
