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

// DcimConsoleServerPortTemplatesUpdateReader is a Reader for the DcimConsoleServerPortTemplatesUpdate structure.
type DcimConsoleServerPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimConsoleServerPortTemplatesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimConsoleServerPortTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortTemplatesUpdateOK creates a DcimConsoleServerPortTemplatesUpdateOK with default headers values
func NewDcimConsoleServerPortTemplatesUpdateOK() *DcimConsoleServerPortTemplatesUpdateOK {
	return &DcimConsoleServerPortTemplatesUpdateOK{}
}

/*
DcimConsoleServerPortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesUpdateOK dcim console server port templates update o k
*/
type DcimConsoleServerPortTemplatesUpdateOK struct {
	Payload *models.ConsoleServerPortTemplate
}

// IsSuccess returns true when this dcim console server port templates update o k response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server port templates update o k response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server port templates update o k response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server port templates update o k response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server port templates update o k response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimConsoleServerPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesUpdateOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortTemplatesUpdateBadRequest creates a DcimConsoleServerPortTemplatesUpdateBadRequest with default headers values
func NewDcimConsoleServerPortTemplatesUpdateBadRequest() *DcimConsoleServerPortTemplatesUpdateBadRequest {
	return &DcimConsoleServerPortTemplatesUpdateBadRequest{}
}

/*
DcimConsoleServerPortTemplatesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimConsoleServerPortTemplatesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim console server port templates update bad request response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim console server port templates update bad request response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server port templates update bad request response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim console server port templates update bad request response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server port templates update bad request response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimConsoleServerPortTemplatesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortTemplatesUpdateDefault creates a DcimConsoleServerPortTemplatesUpdateDefault with default headers values
func NewDcimConsoleServerPortTemplatesUpdateDefault(code int) *DcimConsoleServerPortTemplatesUpdateDefault {
	return &DcimConsoleServerPortTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimConsoleServerPortTemplatesUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimConsoleServerPortTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server port templates update default response
func (o *DcimConsoleServerPortTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console server port templates update default response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console server port templates update default response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console server port templates update default response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console server port templates update default response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console server port templates update default response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsoleServerPortTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
