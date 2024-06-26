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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimDevicesBulkDeleteReader is a Reader for the DcimDevicesBulkDelete structure.
type DcimDevicesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDevicesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimDevicesBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimDevicesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDevicesBulkDeleteNoContent creates a DcimDevicesBulkDeleteNoContent with default headers values
func NewDcimDevicesBulkDeleteNoContent() *DcimDevicesBulkDeleteNoContent {
	return &DcimDevicesBulkDeleteNoContent{}
}

/*
DcimDevicesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimDevicesBulkDeleteNoContent dcim devices bulk delete no content
*/
type DcimDevicesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim devices bulk delete no content response has a 2xx status code
func (o *DcimDevicesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim devices bulk delete no content response has a 3xx status code
func (o *DcimDevicesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices bulk delete no content response has a 4xx status code
func (o *DcimDevicesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim devices bulk delete no content response has a 5xx status code
func (o *DcimDevicesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices bulk delete no content response a status code equal to that given
func (o *DcimDevicesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimDevicesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/devices/][%d] dcimDevicesBulkDeleteNoContent ", 204)
}

func (o *DcimDevicesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/devices/][%d] dcimDevicesBulkDeleteNoContent ", 204)
}

func (o *DcimDevicesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimDevicesBulkDeleteBadRequest creates a DcimDevicesBulkDeleteBadRequest with default headers values
func NewDcimDevicesBulkDeleteBadRequest() *DcimDevicesBulkDeleteBadRequest {
	return &DcimDevicesBulkDeleteBadRequest{}
}

/*
DcimDevicesBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimDevicesBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim devices bulk delete bad request response has a 2xx status code
func (o *DcimDevicesBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim devices bulk delete bad request response has a 3xx status code
func (o *DcimDevicesBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices bulk delete bad request response has a 4xx status code
func (o *DcimDevicesBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim devices bulk delete bad request response has a 5xx status code
func (o *DcimDevicesBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices bulk delete bad request response a status code equal to that given
func (o *DcimDevicesBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimDevicesBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/devices/][%d] dcimDevicesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDevicesBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/devices/][%d] dcimDevicesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDevicesBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDevicesBulkDeleteDefault creates a DcimDevicesBulkDeleteDefault with default headers values
func NewDcimDevicesBulkDeleteDefault(code int) *DcimDevicesBulkDeleteDefault {
	return &DcimDevicesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimDevicesBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimDevicesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim devices bulk delete default response
func (o *DcimDevicesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim devices bulk delete default response has a 2xx status code
func (o *DcimDevicesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim devices bulk delete default response has a 3xx status code
func (o *DcimDevicesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim devices bulk delete default response has a 4xx status code
func (o *DcimDevicesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim devices bulk delete default response has a 5xx status code
func (o *DcimDevicesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim devices bulk delete default response a status code equal to that given
func (o *DcimDevicesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDevicesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/devices/][%d] dcim_devices_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/devices/][%d] dcim_devices_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDevicesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
