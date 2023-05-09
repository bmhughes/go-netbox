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

// DcimRearPortsReadReader is a Reader for the DcimRearPortsRead structure.
type DcimRearPortsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimRearPortsReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimRearPortsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortsReadOK creates a DcimRearPortsReadOK with default headers values
func NewDcimRearPortsReadOK() *DcimRearPortsReadOK {
	return &DcimRearPortsReadOK{}
}

/*
DcimRearPortsReadOK describes a response with status code 200, with default header values.

DcimRearPortsReadOK dcim rear ports read o k
*/
type DcimRearPortsReadOK struct {
	Payload *models.RearPort
}

// IsSuccess returns true when this dcim rear ports read o k response has a 2xx status code
func (o *DcimRearPortsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear ports read o k response has a 3xx status code
func (o *DcimRearPortsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear ports read o k response has a 4xx status code
func (o *DcimRearPortsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear ports read o k response has a 5xx status code
func (o *DcimRearPortsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear ports read o k response a status code equal to that given
func (o *DcimRearPortsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRearPortsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/][%d] dcimRearPortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/][%d] dcimRearPortsReadOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortsReadOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortsReadBadRequest creates a DcimRearPortsReadBadRequest with default headers values
func NewDcimRearPortsReadBadRequest() *DcimRearPortsReadBadRequest {
	return &DcimRearPortsReadBadRequest{}
}

/*
DcimRearPortsReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimRearPortsReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim rear ports read bad request response has a 2xx status code
func (o *DcimRearPortsReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim rear ports read bad request response has a 3xx status code
func (o *DcimRearPortsReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear ports read bad request response has a 4xx status code
func (o *DcimRearPortsReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim rear ports read bad request response has a 5xx status code
func (o *DcimRearPortsReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear ports read bad request response a status code equal to that given
func (o *DcimRearPortsReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimRearPortsReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/][%d] dcimRearPortsReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRearPortsReadBadRequest) String() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/][%d] dcimRearPortsReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRearPortsReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortsReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortsReadDefault creates a DcimRearPortsReadDefault with default headers values
func NewDcimRearPortsReadDefault(code int) *DcimRearPortsReadDefault {
	return &DcimRearPortsReadDefault{
		_statusCode: code,
	}
}

/*
DcimRearPortsReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimRearPortsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rear ports read default response
func (o *DcimRearPortsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim rear ports read default response has a 2xx status code
func (o *DcimRearPortsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rear ports read default response has a 3xx status code
func (o *DcimRearPortsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rear ports read default response has a 4xx status code
func (o *DcimRearPortsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rear ports read default response has a 5xx status code
func (o *DcimRearPortsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rear ports read default response a status code equal to that given
func (o *DcimRearPortsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRearPortsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/][%d] dcim_rear-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/][%d] dcim_rear-ports_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
