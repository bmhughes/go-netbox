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

// DcimConsoleServerPortTemplatesReadReader is a Reader for the DcimConsoleServerPortTemplatesRead structure.
type DcimConsoleServerPortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsoleServerPortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimConsoleServerPortTemplatesReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimConsoleServerPortTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortTemplatesReadOK creates a DcimConsoleServerPortTemplatesReadOK with default headers values
func NewDcimConsoleServerPortTemplatesReadOK() *DcimConsoleServerPortTemplatesReadOK {
	return &DcimConsoleServerPortTemplatesReadOK{}
}

/*
DcimConsoleServerPortTemplatesReadOK describes a response with status code 200, with default header values.

DcimConsoleServerPortTemplatesReadOK dcim console server port templates read o k
*/
type DcimConsoleServerPortTemplatesReadOK struct {
	Payload *models.ConsoleServerPortTemplate
}

// IsSuccess returns true when this dcim console server port templates read o k response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server port templates read o k response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server port templates read o k response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server port templates read o k response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server port templates read o k response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimConsoleServerPortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesReadOK) GetPayload() *models.ConsoleServerPortTemplate {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsoleServerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortTemplatesReadBadRequest creates a DcimConsoleServerPortTemplatesReadBadRequest with default headers values
func NewDcimConsoleServerPortTemplatesReadBadRequest() *DcimConsoleServerPortTemplatesReadBadRequest {
	return &DcimConsoleServerPortTemplatesReadBadRequest{}
}

/*
DcimConsoleServerPortTemplatesReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimConsoleServerPortTemplatesReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim console server port templates read bad request response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim console server port templates read bad request response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server port templates read bad request response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim console server port templates read bad request response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server port templates read bad request response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimConsoleServerPortTemplatesReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesReadBadRequest) String() string {
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortTemplatesReadDefault creates a DcimConsoleServerPortTemplatesReadDefault with default headers values
func NewDcimConsoleServerPortTemplatesReadDefault(code int) *DcimConsoleServerPortTemplatesReadDefault {
	return &DcimConsoleServerPortTemplatesReadDefault{
		_statusCode: code,
	}
}

/*
DcimConsoleServerPortTemplatesReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimConsoleServerPortTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server port templates read default response
func (o *DcimConsoleServerPortTemplatesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console server port templates read default response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console server port templates read default response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console server port templates read default response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console server port templates read default response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console server port templates read default response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsoleServerPortTemplatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
