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

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// DcimRackRolesCreateReader is a Reader for the DcimRackRolesCreate structure.
type DcimRackRolesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimRackRolesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimRackRolesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimRackRolesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackRolesCreateCreated creates a DcimRackRolesCreateCreated with default headers values
func NewDcimRackRolesCreateCreated() *DcimRackRolesCreateCreated {
	return &DcimRackRolesCreateCreated{}
}

/*
DcimRackRolesCreateCreated describes a response with status code 201, with default header values.

DcimRackRolesCreateCreated dcim rack roles create created
*/
type DcimRackRolesCreateCreated struct {
	Payload *models.RackRole
}

// IsSuccess returns true when this dcim rack roles create created response has a 2xx status code
func (o *DcimRackRolesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack roles create created response has a 3xx status code
func (o *DcimRackRolesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles create created response has a 4xx status code
func (o *DcimRackRolesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack roles create created response has a 5xx status code
func (o *DcimRackRolesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles create created response a status code equal to that given
func (o *DcimRackRolesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimRackRolesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/rack-roles/][%d] dcimRackRolesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRackRolesCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/rack-roles/][%d] dcimRackRolesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimRackRolesCreateCreated) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesCreateBadRequest creates a DcimRackRolesCreateBadRequest with default headers values
func NewDcimRackRolesCreateBadRequest() *DcimRackRolesCreateBadRequest {
	return &DcimRackRolesCreateBadRequest{}
}

/*
DcimRackRolesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimRackRolesCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim rack roles create bad request response has a 2xx status code
func (o *DcimRackRolesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim rack roles create bad request response has a 3xx status code
func (o *DcimRackRolesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles create bad request response has a 4xx status code
func (o *DcimRackRolesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim rack roles create bad request response has a 5xx status code
func (o *DcimRackRolesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles create bad request response a status code equal to that given
func (o *DcimRackRolesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimRackRolesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /dcim/rack-roles/][%d] dcimRackRolesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRackRolesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /dcim/rack-roles/][%d] dcimRackRolesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRackRolesCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesCreateDefault creates a DcimRackRolesCreateDefault with default headers values
func NewDcimRackRolesCreateDefault(code int) *DcimRackRolesCreateDefault {
	return &DcimRackRolesCreateDefault{
		_statusCode: code,
	}
}

/*
DcimRackRolesCreateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimRackRolesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack roles create default response
func (o *DcimRackRolesCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim rack roles create default response has a 2xx status code
func (o *DcimRackRolesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rack roles create default response has a 3xx status code
func (o *DcimRackRolesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rack roles create default response has a 4xx status code
func (o *DcimRackRolesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rack roles create default response has a 5xx status code
func (o *DcimRackRolesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rack roles create default response a status code equal to that given
func (o *DcimRackRolesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRackRolesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/rack-roles/][%d] dcim_rack-roles_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/rack-roles/][%d] dcim_rack-roles_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
