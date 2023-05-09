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

// DcimInterfaceTemplatesCreateReader is a Reader for the DcimInterfaceTemplatesCreate structure.
type DcimInterfaceTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfaceTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimInterfaceTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimInterfaceTemplatesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimInterfaceTemplatesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfaceTemplatesCreateCreated creates a DcimInterfaceTemplatesCreateCreated with default headers values
func NewDcimInterfaceTemplatesCreateCreated() *DcimInterfaceTemplatesCreateCreated {
	return &DcimInterfaceTemplatesCreateCreated{}
}

/*
DcimInterfaceTemplatesCreateCreated describes a response with status code 201, with default header values.

DcimInterfaceTemplatesCreateCreated dcim interface templates create created
*/
type DcimInterfaceTemplatesCreateCreated struct {
	Payload *models.InterfaceTemplate
}

// IsSuccess returns true when this dcim interface templates create created response has a 2xx status code
func (o *DcimInterfaceTemplatesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interface templates create created response has a 3xx status code
func (o *DcimInterfaceTemplatesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interface templates create created response has a 4xx status code
func (o *DcimInterfaceTemplatesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interface templates create created response has a 5xx status code
func (o *DcimInterfaceTemplatesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interface templates create created response a status code equal to that given
func (o *DcimInterfaceTemplatesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimInterfaceTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/interface-templates/][%d] dcimInterfaceTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimInterfaceTemplatesCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/interface-templates/][%d] dcimInterfaceTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimInterfaceTemplatesCreateCreated) GetPayload() *models.InterfaceTemplate {
	return o.Payload
}

func (o *DcimInterfaceTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InterfaceTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfaceTemplatesCreateBadRequest creates a DcimInterfaceTemplatesCreateBadRequest with default headers values
func NewDcimInterfaceTemplatesCreateBadRequest() *DcimInterfaceTemplatesCreateBadRequest {
	return &DcimInterfaceTemplatesCreateBadRequest{}
}

/*
DcimInterfaceTemplatesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimInterfaceTemplatesCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim interface templates create bad request response has a 2xx status code
func (o *DcimInterfaceTemplatesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim interface templates create bad request response has a 3xx status code
func (o *DcimInterfaceTemplatesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interface templates create bad request response has a 4xx status code
func (o *DcimInterfaceTemplatesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim interface templates create bad request response has a 5xx status code
func (o *DcimInterfaceTemplatesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interface templates create bad request response a status code equal to that given
func (o *DcimInterfaceTemplatesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimInterfaceTemplatesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /dcim/interface-templates/][%d] dcimInterfaceTemplatesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInterfaceTemplatesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /dcim/interface-templates/][%d] dcimInterfaceTemplatesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInterfaceTemplatesCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfaceTemplatesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfaceTemplatesCreateDefault creates a DcimInterfaceTemplatesCreateDefault with default headers values
func NewDcimInterfaceTemplatesCreateDefault(code int) *DcimInterfaceTemplatesCreateDefault {
	return &DcimInterfaceTemplatesCreateDefault{
		_statusCode: code,
	}
}

/*
DcimInterfaceTemplatesCreateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimInterfaceTemplatesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interface templates create default response
func (o *DcimInterfaceTemplatesCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim interface templates create default response has a 2xx status code
func (o *DcimInterfaceTemplatesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interface templates create default response has a 3xx status code
func (o *DcimInterfaceTemplatesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interface templates create default response has a 4xx status code
func (o *DcimInterfaceTemplatesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interface templates create default response has a 5xx status code
func (o *DcimInterfaceTemplatesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interface templates create default response a status code equal to that given
func (o *DcimInterfaceTemplatesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInterfaceTemplatesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/interface-templates/][%d] dcim_interface-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfaceTemplatesCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/interface-templates/][%d] dcim_interface-templates_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfaceTemplatesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfaceTemplatesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
