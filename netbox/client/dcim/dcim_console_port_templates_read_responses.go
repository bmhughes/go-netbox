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

// DcimConsolePortTemplatesReadReader is a Reader for the DcimConsolePortTemplatesRead structure.
type DcimConsolePortTemplatesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimConsolePortTemplatesReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimConsolePortTemplatesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortTemplatesReadOK creates a DcimConsolePortTemplatesReadOK with default headers values
func NewDcimConsolePortTemplatesReadOK() *DcimConsolePortTemplatesReadOK {
	return &DcimConsolePortTemplatesReadOK{}
}

/*
DcimConsolePortTemplatesReadOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesReadOK dcim console port templates read o k
*/
type DcimConsolePortTemplatesReadOK struct {
	Payload *models.ConsolePortTemplate
}

// IsSuccess returns true when this dcim console port templates read o k response has a 2xx status code
func (o *DcimConsolePortTemplatesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console port templates read o k response has a 3xx status code
func (o *DcimConsolePortTemplatesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates read o k response has a 4xx status code
func (o *DcimConsolePortTemplatesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console port templates read o k response has a 5xx status code
func (o *DcimConsolePortTemplatesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates read o k response a status code equal to that given
func (o *DcimConsolePortTemplatesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimConsolePortTemplatesReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesReadOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesReadOK) GetPayload() *models.ConsolePortTemplate {
	return o.Payload
}

func (o *DcimConsolePortTemplatesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConsolePortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesReadBadRequest creates a DcimConsolePortTemplatesReadBadRequest with default headers values
func NewDcimConsolePortTemplatesReadBadRequest() *DcimConsolePortTemplatesReadBadRequest {
	return &DcimConsolePortTemplatesReadBadRequest{}
}

/*
DcimConsolePortTemplatesReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimConsolePortTemplatesReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim console port templates read bad request response has a 2xx status code
func (o *DcimConsolePortTemplatesReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim console port templates read bad request response has a 3xx status code
func (o *DcimConsolePortTemplatesReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates read bad request response has a 4xx status code
func (o *DcimConsolePortTemplatesReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim console port templates read bad request response has a 5xx status code
func (o *DcimConsolePortTemplatesReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates read bad request response a status code equal to that given
func (o *DcimConsolePortTemplatesReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimConsolePortTemplatesReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsolePortTemplatesReadBadRequest) String() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/{id}/][%d] dcimConsolePortTemplatesReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsolePortTemplatesReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesReadDefault creates a DcimConsolePortTemplatesReadDefault with default headers values
func NewDcimConsolePortTemplatesReadDefault(code int) *DcimConsolePortTemplatesReadDefault {
	return &DcimConsolePortTemplatesReadDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortTemplatesReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimConsolePortTemplatesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console port templates read default response
func (o *DcimConsolePortTemplatesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console port templates read default response has a 2xx status code
func (o *DcimConsolePortTemplatesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console port templates read default response has a 3xx status code
func (o *DcimConsolePortTemplatesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console port templates read default response has a 4xx status code
func (o *DcimConsolePortTemplatesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console port templates read default response has a 5xx status code
func (o *DcimConsolePortTemplatesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console port templates read default response a status code equal to that given
func (o *DcimConsolePortTemplatesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsolePortTemplatesReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/{id}/][%d] dcim_console-port-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/{id}/][%d] dcim_console-port-templates_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
