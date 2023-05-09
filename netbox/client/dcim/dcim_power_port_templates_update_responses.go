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

// DcimPowerPortTemplatesUpdateReader is a Reader for the DcimPowerPortTemplatesUpdate structure.
type DcimPowerPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimPowerPortTemplatesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimPowerPortTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPortTemplatesUpdateOK creates a DcimPowerPortTemplatesUpdateOK with default headers values
func NewDcimPowerPortTemplatesUpdateOK() *DcimPowerPortTemplatesUpdateOK {
	return &DcimPowerPortTemplatesUpdateOK{}
}

/*
DcimPowerPortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortTemplatesUpdateOK dcim power port templates update o k
*/
type DcimPowerPortTemplatesUpdateOK struct {
	Payload *models.PowerPortTemplate
}

// IsSuccess returns true when this dcim power port templates update o k response has a 2xx status code
func (o *DcimPowerPortTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power port templates update o k response has a 3xx status code
func (o *DcimPowerPortTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power port templates update o k response has a 4xx status code
func (o *DcimPowerPortTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power port templates update o k response has a 5xx status code
func (o *DcimPowerPortTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power port templates update o k response a status code equal to that given
func (o *DcimPowerPortTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortTemplatesUpdateOK) GetPayload() *models.PowerPortTemplate {
	return o.Payload
}

func (o *DcimPowerPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortTemplatesUpdateBadRequest creates a DcimPowerPortTemplatesUpdateBadRequest with default headers values
func NewDcimPowerPortTemplatesUpdateBadRequest() *DcimPowerPortTemplatesUpdateBadRequest {
	return &DcimPowerPortTemplatesUpdateBadRequest{}
}

/*
DcimPowerPortTemplatesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimPowerPortTemplatesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim power port templates update bad request response has a 2xx status code
func (o *DcimPowerPortTemplatesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim power port templates update bad request response has a 3xx status code
func (o *DcimPowerPortTemplatesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power port templates update bad request response has a 4xx status code
func (o *DcimPowerPortTemplatesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim power port templates update bad request response has a 5xx status code
func (o *DcimPowerPortTemplatesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power port templates update bad request response a status code equal to that given
func (o *DcimPowerPortTemplatesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimPowerPortTemplatesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerPortTemplatesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcimPowerPortTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerPortTemplatesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortTemplatesUpdateDefault creates a DcimPowerPortTemplatesUpdateDefault with default headers values
func NewDcimPowerPortTemplatesUpdateDefault(code int) *DcimPowerPortTemplatesUpdateDefault {
	return &DcimPowerPortTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPortTemplatesUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimPowerPortTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power port templates update default response
func (o *DcimPowerPortTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power port templates update default response has a 2xx status code
func (o *DcimPowerPortTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power port templates update default response has a 3xx status code
func (o *DcimPowerPortTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power port templates update default response has a 4xx status code
func (o *DcimPowerPortTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power port templates update default response has a 5xx status code
func (o *DcimPowerPortTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power port templates update default response a status code equal to that given
func (o *DcimPowerPortTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerPortTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcim_power-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/power-port-templates/{id}/][%d] dcim_power-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPortTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
