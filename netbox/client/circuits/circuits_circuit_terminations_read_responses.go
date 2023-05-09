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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// CircuitsCircuitTerminationsReadReader is a Reader for the CircuitsCircuitTerminationsRead structure.
type CircuitsCircuitTerminationsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCircuitsCircuitTerminationsReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCircuitsCircuitTerminationsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsReadOK creates a CircuitsCircuitTerminationsReadOK with default headers values
func NewCircuitsCircuitTerminationsReadOK() *CircuitsCircuitTerminationsReadOK {
	return &CircuitsCircuitTerminationsReadOK{}
}

/*
CircuitsCircuitTerminationsReadOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsReadOK circuits circuit terminations read o k
*/
type CircuitsCircuitTerminationsReadOK struct {
	Payload *models.CircuitTermination
}

// IsSuccess returns true when this circuits circuit terminations read o k response has a 2xx status code
func (o *CircuitsCircuitTerminationsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit terminations read o k response has a 3xx status code
func (o *CircuitsCircuitTerminationsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit terminations read o k response has a 4xx status code
func (o *CircuitsCircuitTerminationsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit terminations read o k response has a 5xx status code
func (o *CircuitsCircuitTerminationsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit terminations read o k response a status code equal to that given
func (o *CircuitsCircuitTerminationsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsCircuitTerminationsReadOK) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsReadOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsReadOK) String() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsReadOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsReadOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsReadBadRequest creates a CircuitsCircuitTerminationsReadBadRequest with default headers values
func NewCircuitsCircuitTerminationsReadBadRequest() *CircuitsCircuitTerminationsReadBadRequest {
	return &CircuitsCircuitTerminationsReadBadRequest{}
}

/*
CircuitsCircuitTerminationsReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CircuitsCircuitTerminationsReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this circuits circuit terminations read bad request response has a 2xx status code
func (o *CircuitsCircuitTerminationsReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this circuits circuit terminations read bad request response has a 3xx status code
func (o *CircuitsCircuitTerminationsReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit terminations read bad request response has a 4xx status code
func (o *CircuitsCircuitTerminationsReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this circuits circuit terminations read bad request response has a 5xx status code
func (o *CircuitsCircuitTerminationsReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit terminations read bad request response a status code equal to that given
func (o *CircuitsCircuitTerminationsReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CircuitsCircuitTerminationsReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsReadBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsCircuitTerminationsReadBadRequest) String() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/][%d] circuitsCircuitTerminationsReadBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsCircuitTerminationsReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsReadDefault creates a CircuitsCircuitTerminationsReadDefault with default headers values
func NewCircuitsCircuitTerminationsReadDefault(code int) *CircuitsCircuitTerminationsReadDefault {
	return &CircuitsCircuitTerminationsReadDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitTerminationsReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type CircuitsCircuitTerminationsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit terminations read default response
func (o *CircuitsCircuitTerminationsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits circuit terminations read default response has a 2xx status code
func (o *CircuitsCircuitTerminationsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuit terminations read default response has a 3xx status code
func (o *CircuitsCircuitTerminationsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuit terminations read default response has a 4xx status code
func (o *CircuitsCircuitTerminationsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuit terminations read default response has a 5xx status code
func (o *CircuitsCircuitTerminationsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuit terminations read default response a status code equal to that given
func (o *CircuitsCircuitTerminationsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsCircuitTerminationsReadDefault) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/][%d] circuits_circuit-terminations_read default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTerminationsReadDefault) String() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/][%d] circuits_circuit-terminations_read default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTerminationsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
