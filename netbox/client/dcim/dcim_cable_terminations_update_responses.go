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

// DcimCableTerminationsUpdateReader is a Reader for the DcimCableTerminationsUpdate structure.
type DcimCableTerminationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCableTerminationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCableTerminationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimCableTerminationsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimCableTerminationsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCableTerminationsUpdateOK creates a DcimCableTerminationsUpdateOK with default headers values
func NewDcimCableTerminationsUpdateOK() *DcimCableTerminationsUpdateOK {
	return &DcimCableTerminationsUpdateOK{}
}

/*
DcimCableTerminationsUpdateOK describes a response with status code 200, with default header values.

DcimCableTerminationsUpdateOK dcim cable terminations update o k
*/
type DcimCableTerminationsUpdateOK struct {
	Payload *models.CableTermination
}

// IsSuccess returns true when this dcim cable terminations update o k response has a 2xx status code
func (o *DcimCableTerminationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cable terminations update o k response has a 3xx status code
func (o *DcimCableTerminationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations update o k response has a 4xx status code
func (o *DcimCableTerminationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cable terminations update o k response has a 5xx status code
func (o *DcimCableTerminationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations update o k response a status code equal to that given
func (o *DcimCableTerminationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimCableTerminationsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimCableTerminationsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimCableTerminationsUpdateOK) GetPayload() *models.CableTermination {
	return o.Payload
}

func (o *DcimCableTerminationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CableTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCableTerminationsUpdateBadRequest creates a DcimCableTerminationsUpdateBadRequest with default headers values
func NewDcimCableTerminationsUpdateBadRequest() *DcimCableTerminationsUpdateBadRequest {
	return &DcimCableTerminationsUpdateBadRequest{}
}

/*
DcimCableTerminationsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimCableTerminationsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim cable terminations update bad request response has a 2xx status code
func (o *DcimCableTerminationsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim cable terminations update bad request response has a 3xx status code
func (o *DcimCableTerminationsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations update bad request response has a 4xx status code
func (o *DcimCableTerminationsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim cable terminations update bad request response has a 5xx status code
func (o *DcimCableTerminationsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations update bad request response a status code equal to that given
func (o *DcimCableTerminationsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimCableTerminationsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimCableTerminationsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimCableTerminationsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCableTerminationsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCableTerminationsUpdateDefault creates a DcimCableTerminationsUpdateDefault with default headers values
func NewDcimCableTerminationsUpdateDefault(code int) *DcimCableTerminationsUpdateDefault {
	return &DcimCableTerminationsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimCableTerminationsUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimCableTerminationsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cable terminations update default response
func (o *DcimCableTerminationsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim cable terminations update default response has a 2xx status code
func (o *DcimCableTerminationsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cable terminations update default response has a 3xx status code
func (o *DcimCableTerminationsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cable terminations update default response has a 4xx status code
func (o *DcimCableTerminationsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cable terminations update default response has a 5xx status code
func (o *DcimCableTerminationsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cable terminations update default response a status code equal to that given
func (o *DcimCableTerminationsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimCableTerminationsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/cable-terminations/{id}/][%d] dcim_cable-terminations_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCableTerminationsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/cable-terminations/{id}/][%d] dcim_cable-terminations_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCableTerminationsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCableTerminationsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
