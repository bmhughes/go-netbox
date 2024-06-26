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

// DcimConsolePortTemplatesUpdateReader is a Reader for the DcimConsolePortTemplatesUpdate structure.
type DcimConsolePortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimConsolePortTemplatesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimConsolePortTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortTemplatesUpdateOK creates a DcimConsolePortTemplatesUpdateOK with default headers values
func NewDcimConsolePortTemplatesUpdateOK() *DcimConsolePortTemplatesUpdateOK {
	return &DcimConsolePortTemplatesUpdateOK{}
}

/*
DcimConsolePortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesUpdateOK dcim console port templates update o k
*/
type DcimConsolePortTemplatesUpdateOK struct {
	Payload *models.ConsolePortTemplate
}

// IsSuccess returns true when this dcim console port templates update o k response has a 2xx status code
func (o *DcimConsolePortTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console port templates update o k response has a 3xx status code
func (o *DcimConsolePortTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates update o k response has a 4xx status code
func (o *DcimConsolePortTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console port templates update o k response has a 5xx status code
func (o *DcimConsolePortTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates update o k response a status code equal to that given
func (o *DcimConsolePortTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimConsolePortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesUpdateOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesUpdateBadRequest creates a DcimConsolePortTemplatesUpdateBadRequest with default headers values
func NewDcimConsolePortTemplatesUpdateBadRequest() *DcimConsolePortTemplatesUpdateBadRequest {
	return &DcimConsolePortTemplatesUpdateBadRequest{}
}

/*
DcimConsolePortTemplatesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimConsolePortTemplatesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim console port templates update bad request response has a 2xx status code
func (o *DcimConsolePortTemplatesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim console port templates update bad request response has a 3xx status code
func (o *DcimConsolePortTemplatesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates update bad request response has a 4xx status code
func (o *DcimConsolePortTemplatesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim console port templates update bad request response has a 5xx status code
func (o *DcimConsolePortTemplatesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates update bad request response a status code equal to that given
func (o *DcimConsolePortTemplatesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimConsolePortTemplatesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsolePortTemplatesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsolePortTemplatesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesUpdateDefault creates a DcimConsolePortTemplatesUpdateDefault with default headers values
func NewDcimConsolePortTemplatesUpdateDefault(code int) *DcimConsolePortTemplatesUpdateDefault {
	return &DcimConsolePortTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortTemplatesUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimConsolePortTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console port templates update default response
func (o *DcimConsolePortTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console port templates update default response has a 2xx status code
func (o *DcimConsolePortTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console port templates update default response has a 3xx status code
func (o *DcimConsolePortTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console port templates update default response has a 4xx status code
func (o *DcimConsolePortTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console port templates update default response has a 5xx status code
func (o *DcimConsolePortTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console port templates update default response a status code equal to that given
func (o *DcimConsolePortTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsolePortTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcim_console-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/console-port-templates/{id}/][%d] dcim_console-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
