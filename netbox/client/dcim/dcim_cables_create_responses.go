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

// DcimCablesCreateReader is a Reader for the DcimCablesCreate structure.
type DcimCablesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimCablesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimCablesCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimCablesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCablesCreateCreated creates a DcimCablesCreateCreated with default headers values
func NewDcimCablesCreateCreated() *DcimCablesCreateCreated {
	return &DcimCablesCreateCreated{}
}

/*
DcimCablesCreateCreated describes a response with status code 201, with default header values.

DcimCablesCreateCreated dcim cables create created
*/
type DcimCablesCreateCreated struct {
	Payload *models.Cable
}

// IsSuccess returns true when this dcim cables create created response has a 2xx status code
func (o *DcimCablesCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cables create created response has a 3xx status code
func (o *DcimCablesCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cables create created response has a 4xx status code
func (o *DcimCablesCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cables create created response has a 5xx status code
func (o *DcimCablesCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cables create created response a status code equal to that given
func (o *DcimCablesCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimCablesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/cables/][%d] dcimCablesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimCablesCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/cables/][%d] dcimCablesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimCablesCreateCreated) GetPayload() *models.Cable {
	return o.Payload
}

func (o *DcimCablesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCablesCreateBadRequest creates a DcimCablesCreateBadRequest with default headers values
func NewDcimCablesCreateBadRequest() *DcimCablesCreateBadRequest {
	return &DcimCablesCreateBadRequest{}
}

/*
DcimCablesCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimCablesCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim cables create bad request response has a 2xx status code
func (o *DcimCablesCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim cables create bad request response has a 3xx status code
func (o *DcimCablesCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cables create bad request response has a 4xx status code
func (o *DcimCablesCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim cables create bad request response has a 5xx status code
func (o *DcimCablesCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cables create bad request response a status code equal to that given
func (o *DcimCablesCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimCablesCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /dcim/cables/][%d] dcimCablesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimCablesCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /dcim/cables/][%d] dcimCablesCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimCablesCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCablesCreateDefault creates a DcimCablesCreateDefault with default headers values
func NewDcimCablesCreateDefault(code int) *DcimCablesCreateDefault {
	return &DcimCablesCreateDefault{
		_statusCode: code,
	}
}

/*
DcimCablesCreateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimCablesCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cables create default response
func (o *DcimCablesCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim cables create default response has a 2xx status code
func (o *DcimCablesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cables create default response has a 3xx status code
func (o *DcimCablesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cables create default response has a 4xx status code
func (o *DcimCablesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cables create default response has a 5xx status code
func (o *DcimCablesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cables create default response a status code equal to that given
func (o *DcimCablesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimCablesCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/cables/][%d] dcim_cables_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/cables/][%d] dcim_cables_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
