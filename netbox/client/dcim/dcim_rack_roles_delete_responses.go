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

// DcimRackRolesDeleteReader is a Reader for the DcimRackRolesDelete structure.
type DcimRackRolesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRackRolesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimRackRolesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimRackRolesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackRolesDeleteNoContent creates a DcimRackRolesDeleteNoContent with default headers values
func NewDcimRackRolesDeleteNoContent() *DcimRackRolesDeleteNoContent {
	return &DcimRackRolesDeleteNoContent{}
}

/*
DcimRackRolesDeleteNoContent describes a response with status code 204, with default header values.

DcimRackRolesDeleteNoContent dcim rack roles delete no content
*/
type DcimRackRolesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim rack roles delete no content response has a 2xx status code
func (o *DcimRackRolesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack roles delete no content response has a 3xx status code
func (o *DcimRackRolesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles delete no content response has a 4xx status code
func (o *DcimRackRolesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack roles delete no content response has a 5xx status code
func (o *DcimRackRolesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles delete no content response a status code equal to that given
func (o *DcimRackRolesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimRackRolesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-roles/{id}/][%d] dcimRackRolesDeleteNoContent ", 204)
}

func (o *DcimRackRolesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/rack-roles/{id}/][%d] dcimRackRolesDeleteNoContent ", 204)
}

func (o *DcimRackRolesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRackRolesDeleteBadRequest creates a DcimRackRolesDeleteBadRequest with default headers values
func NewDcimRackRolesDeleteBadRequest() *DcimRackRolesDeleteBadRequest {
	return &DcimRackRolesDeleteBadRequest{}
}

/*
DcimRackRolesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimRackRolesDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim rack roles delete bad request response has a 2xx status code
func (o *DcimRackRolesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim rack roles delete bad request response has a 3xx status code
func (o *DcimRackRolesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles delete bad request response has a 4xx status code
func (o *DcimRackRolesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim rack roles delete bad request response has a 5xx status code
func (o *DcimRackRolesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles delete bad request response a status code equal to that given
func (o *DcimRackRolesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimRackRolesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-roles/{id}/][%d] dcimRackRolesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRackRolesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/rack-roles/{id}/][%d] dcimRackRolesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRackRolesDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesDeleteDefault creates a DcimRackRolesDeleteDefault with default headers values
func NewDcimRackRolesDeleteDefault(code int) *DcimRackRolesDeleteDefault {
	return &DcimRackRolesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimRackRolesDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimRackRolesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack roles delete default response
func (o *DcimRackRolesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim rack roles delete default response has a 2xx status code
func (o *DcimRackRolesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rack roles delete default response has a 3xx status code
func (o *DcimRackRolesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rack roles delete default response has a 4xx status code
func (o *DcimRackRolesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rack roles delete default response has a 5xx status code
func (o *DcimRackRolesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rack roles delete default response a status code equal to that given
func (o *DcimRackRolesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRackRolesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-roles/{id}/][%d] dcim_rack-roles_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/rack-roles/{id}/][%d] dcim_rack-roles_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
