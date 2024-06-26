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

// DcimPowerPanelsCreateReader is a Reader for the DcimPowerPanelsCreate structure.
type DcimPowerPanelsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPowerPanelsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimPowerPanelsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimPowerPanelsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPanelsCreateCreated creates a DcimPowerPanelsCreateCreated with default headers values
func NewDcimPowerPanelsCreateCreated() *DcimPowerPanelsCreateCreated {
	return &DcimPowerPanelsCreateCreated{}
}

/*
DcimPowerPanelsCreateCreated describes a response with status code 201, with default header values.

DcimPowerPanelsCreateCreated dcim power panels create created
*/
type DcimPowerPanelsCreateCreated struct {
	Payload *models.PowerPanel
}

// IsSuccess returns true when this dcim power panels create created response has a 2xx status code
func (o *DcimPowerPanelsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power panels create created response has a 3xx status code
func (o *DcimPowerPanelsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power panels create created response has a 4xx status code
func (o *DcimPowerPanelsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power panels create created response has a 5xx status code
func (o *DcimPowerPanelsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power panels create created response a status code equal to that given
func (o *DcimPowerPanelsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimPowerPanelsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-panels/][%d] dcimPowerPanelsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerPanelsCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/power-panels/][%d] dcimPowerPanelsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerPanelsCreateCreated) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPanelsCreateBadRequest creates a DcimPowerPanelsCreateBadRequest with default headers values
func NewDcimPowerPanelsCreateBadRequest() *DcimPowerPanelsCreateBadRequest {
	return &DcimPowerPanelsCreateBadRequest{}
}

/*
DcimPowerPanelsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimPowerPanelsCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim power panels create bad request response has a 2xx status code
func (o *DcimPowerPanelsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim power panels create bad request response has a 3xx status code
func (o *DcimPowerPanelsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power panels create bad request response has a 4xx status code
func (o *DcimPowerPanelsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim power panels create bad request response has a 5xx status code
func (o *DcimPowerPanelsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power panels create bad request response a status code equal to that given
func (o *DcimPowerPanelsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimPowerPanelsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /dcim/power-panels/][%d] dcimPowerPanelsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerPanelsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /dcim/power-panels/][%d] dcimPowerPanelsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerPanelsCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPanelsCreateDefault creates a DcimPowerPanelsCreateDefault with default headers values
func NewDcimPowerPanelsCreateDefault(code int) *DcimPowerPanelsCreateDefault {
	return &DcimPowerPanelsCreateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPanelsCreateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimPowerPanelsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power panels create default response
func (o *DcimPowerPanelsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power panels create default response has a 2xx status code
func (o *DcimPowerPanelsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power panels create default response has a 3xx status code
func (o *DcimPowerPanelsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power panels create default response has a 4xx status code
func (o *DcimPowerPanelsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power panels create default response has a 5xx status code
func (o *DcimPowerPanelsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power panels create default response a status code equal to that given
func (o *DcimPowerPanelsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerPanelsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /dcim/power-panels/][%d] dcim_power-panels_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsCreateDefault) String() string {
	return fmt.Sprintf("[POST /dcim/power-panels/][%d] dcim_power-panels_create default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
