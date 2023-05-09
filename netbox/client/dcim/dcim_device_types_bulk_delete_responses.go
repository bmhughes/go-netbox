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

// DcimDeviceTypesBulkDeleteReader is a Reader for the DcimDeviceTypesBulkDelete structure.
type DcimDeviceTypesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceTypesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimDeviceTypesBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimDeviceTypesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceTypesBulkDeleteNoContent creates a DcimDeviceTypesBulkDeleteNoContent with default headers values
func NewDcimDeviceTypesBulkDeleteNoContent() *DcimDeviceTypesBulkDeleteNoContent {
	return &DcimDeviceTypesBulkDeleteNoContent{}
}

/*
DcimDeviceTypesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceTypesBulkDeleteNoContent dcim device types bulk delete no content
*/
type DcimDeviceTypesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim device types bulk delete no content response has a 2xx status code
func (o *DcimDeviceTypesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device types bulk delete no content response has a 3xx status code
func (o *DcimDeviceTypesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device types bulk delete no content response has a 4xx status code
func (o *DcimDeviceTypesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device types bulk delete no content response has a 5xx status code
func (o *DcimDeviceTypesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device types bulk delete no content response a status code equal to that given
func (o *DcimDeviceTypesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimDeviceTypesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/][%d] dcimDeviceTypesBulkDeleteNoContent ", 204)
}

func (o *DcimDeviceTypesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/][%d] dcimDeviceTypesBulkDeleteNoContent ", 204)
}

func (o *DcimDeviceTypesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimDeviceTypesBulkDeleteBadRequest creates a DcimDeviceTypesBulkDeleteBadRequest with default headers values
func NewDcimDeviceTypesBulkDeleteBadRequest() *DcimDeviceTypesBulkDeleteBadRequest {
	return &DcimDeviceTypesBulkDeleteBadRequest{}
}

/*
DcimDeviceTypesBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimDeviceTypesBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim device types bulk delete bad request response has a 2xx status code
func (o *DcimDeviceTypesBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim device types bulk delete bad request response has a 3xx status code
func (o *DcimDeviceTypesBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device types bulk delete bad request response has a 4xx status code
func (o *DcimDeviceTypesBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim device types bulk delete bad request response has a 5xx status code
func (o *DcimDeviceTypesBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device types bulk delete bad request response a status code equal to that given
func (o *DcimDeviceTypesBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimDeviceTypesBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/][%d] dcimDeviceTypesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDeviceTypesBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/][%d] dcimDeviceTypesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDeviceTypesBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceTypesBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceTypesBulkDeleteDefault creates a DcimDeviceTypesBulkDeleteDefault with default headers values
func NewDcimDeviceTypesBulkDeleteDefault(code int) *DcimDeviceTypesBulkDeleteDefault {
	return &DcimDeviceTypesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceTypesBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimDeviceTypesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device types bulk delete default response
func (o *DcimDeviceTypesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device types bulk delete default response has a 2xx status code
func (o *DcimDeviceTypesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device types bulk delete default response has a 3xx status code
func (o *DcimDeviceTypesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device types bulk delete default response has a 4xx status code
func (o *DcimDeviceTypesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device types bulk delete default response has a 5xx status code
func (o *DcimDeviceTypesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device types bulk delete default response a status code equal to that given
func (o *DcimDeviceTypesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceTypesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/][%d] dcim_device-types_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceTypesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-types/][%d] dcim_device-types_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceTypesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceTypesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
