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

// DcimConsolePortsReadReader is a Reader for the DcimConsolePortsRead structure.
type DcimConsolePortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimConsolePortsReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimConsolePortsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortsReadOK creates a DcimConsolePortsReadOK with default headers values
func NewDcimConsolePortsReadOK() *DcimConsolePortsReadOK {
	return &DcimConsolePortsReadOK{}
}

/*
DcimConsolePortsReadOK describes a response with status code 200, with default header values.

DcimConsolePortsReadOK dcim console ports read o k
*/
type DcimConsolePortsReadOK struct {
	Payload *models.ConsolePort
}

// IsSuccess returns true when this dcim console ports read o k response has a 2xx status code
func (o *DcimConsolePortsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console ports read o k response has a 3xx status code
func (o *DcimConsolePortsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console ports read o k response has a 4xx status code
func (o *DcimConsolePortsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console ports read o k response has a 5xx status code
func (o *DcimConsolePortsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console ports read o k response a status code equal to that given
func (o *DcimConsolePortsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimConsolePortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcimConsolePortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcimConsolePortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortsReadOK) GetPayload() *models.ConsolePort {
	return o.Payload
}

func (o *DcimConsolePortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortsReadBadRequest creates a DcimConsolePortsReadBadRequest with default headers values
func NewDcimConsolePortsReadBadRequest() *DcimConsolePortsReadBadRequest {
	return &DcimConsolePortsReadBadRequest{}
}

/*
DcimConsolePortsReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimConsolePortsReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim console ports read bad request response has a 2xx status code
func (o *DcimConsolePortsReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim console ports read bad request response has a 3xx status code
func (o *DcimConsolePortsReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console ports read bad request response has a 4xx status code
func (o *DcimConsolePortsReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim console ports read bad request response has a 5xx status code
func (o *DcimConsolePortsReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console ports read bad request response a status code equal to that given
func (o *DcimConsolePortsReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimConsolePortsReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcimConsolePortsReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsolePortsReadBadRequest) String() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcimConsolePortsReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsolePortsReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortsReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortsReadDefault creates a DcimConsolePortsReadDefault with default headers values
func NewDcimConsolePortsReadDefault(code int) *DcimConsolePortsReadDefault {
	return &DcimConsolePortsReadDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortsReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimConsolePortsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console ports read default response
func (o *DcimConsolePortsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console ports read default response has a 2xx status code
func (o *DcimConsolePortsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console ports read default response has a 3xx status code
func (o *DcimConsolePortsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console ports read default response has a 4xx status code
func (o *DcimConsolePortsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console ports read default response has a 5xx status code
func (o *DcimConsolePortsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console ports read default response a status code equal to that given
func (o *DcimConsolePortsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsolePortsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcim_console-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/console-ports/{id}/][%d] dcim_console-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
