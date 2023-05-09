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

// DcimDeviceRolesDeleteReader is a Reader for the DcimDeviceRolesDelete structure.
type DcimDeviceRolesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceRolesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimDeviceRolesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimDeviceRolesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimDeviceRolesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceRolesDeleteNoContent creates a DcimDeviceRolesDeleteNoContent with default headers values
func NewDcimDeviceRolesDeleteNoContent() *DcimDeviceRolesDeleteNoContent {
	return &DcimDeviceRolesDeleteNoContent{}
}

/*
DcimDeviceRolesDeleteNoContent describes a response with status code 204, with default header values.

DcimDeviceRolesDeleteNoContent dcim device roles delete no content
*/
type DcimDeviceRolesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim device roles delete no content response has a 2xx status code
func (o *DcimDeviceRolesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device roles delete no content response has a 3xx status code
func (o *DcimDeviceRolesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device roles delete no content response has a 4xx status code
func (o *DcimDeviceRolesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device roles delete no content response has a 5xx status code
func (o *DcimDeviceRolesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device roles delete no content response a status code equal to that given
func (o *DcimDeviceRolesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimDeviceRolesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/{id}/][%d] dcimDeviceRolesDeleteNoContent ", 204)
}

func (o *DcimDeviceRolesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/{id}/][%d] dcimDeviceRolesDeleteNoContent ", 204)
}

func (o *DcimDeviceRolesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimDeviceRolesDeleteBadRequest creates a DcimDeviceRolesDeleteBadRequest with default headers values
func NewDcimDeviceRolesDeleteBadRequest() *DcimDeviceRolesDeleteBadRequest {
	return &DcimDeviceRolesDeleteBadRequest{}
}

/*
DcimDeviceRolesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimDeviceRolesDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim device roles delete bad request response has a 2xx status code
func (o *DcimDeviceRolesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim device roles delete bad request response has a 3xx status code
func (o *DcimDeviceRolesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device roles delete bad request response has a 4xx status code
func (o *DcimDeviceRolesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim device roles delete bad request response has a 5xx status code
func (o *DcimDeviceRolesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device roles delete bad request response a status code equal to that given
func (o *DcimDeviceRolesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimDeviceRolesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/{id}/][%d] dcimDeviceRolesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDeviceRolesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/{id}/][%d] dcimDeviceRolesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDeviceRolesDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceRolesDeleteDefault creates a DcimDeviceRolesDeleteDefault with default headers values
func NewDcimDeviceRolesDeleteDefault(code int) *DcimDeviceRolesDeleteDefault {
	return &DcimDeviceRolesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceRolesDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimDeviceRolesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device roles delete default response
func (o *DcimDeviceRolesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device roles delete default response has a 2xx status code
func (o *DcimDeviceRolesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device roles delete default response has a 3xx status code
func (o *DcimDeviceRolesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device roles delete default response has a 4xx status code
func (o *DcimDeviceRolesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device roles delete default response has a 5xx status code
func (o *DcimDeviceRolesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device roles delete default response a status code equal to that given
func (o *DcimDeviceRolesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceRolesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/{id}/][%d] dcim_device-roles_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/device-roles/{id}/][%d] dcim_device-roles_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceRolesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceRolesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
