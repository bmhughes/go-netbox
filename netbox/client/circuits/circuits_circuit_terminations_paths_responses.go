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

// CircuitsCircuitTerminationsPathsReader is a Reader for the CircuitsCircuitTerminationsPaths structure.
type CircuitsCircuitTerminationsPathsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTerminationsPathsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTerminationsPathsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCircuitsCircuitTerminationsPathsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCircuitsCircuitTerminationsPathsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsCircuitTerminationsPathsOK creates a CircuitsCircuitTerminationsPathsOK with default headers values
func NewCircuitsCircuitTerminationsPathsOK() *CircuitsCircuitTerminationsPathsOK {
	return &CircuitsCircuitTerminationsPathsOK{}
}

/*
CircuitsCircuitTerminationsPathsOK describes a response with status code 200, with default header values.

CircuitsCircuitTerminationsPathsOK circuits circuit terminations paths o k
*/
type CircuitsCircuitTerminationsPathsOK struct {
	Payload *models.CircuitTermination
}

// IsSuccess returns true when this circuits circuit terminations paths o k response has a 2xx status code
func (o *CircuitsCircuitTerminationsPathsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit terminations paths o k response has a 3xx status code
func (o *CircuitsCircuitTerminationsPathsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit terminations paths o k response has a 4xx status code
func (o *CircuitsCircuitTerminationsPathsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit terminations paths o k response has a 5xx status code
func (o *CircuitsCircuitTerminationsPathsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit terminations paths o k response a status code equal to that given
func (o *CircuitsCircuitTerminationsPathsOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsCircuitTerminationsPathsOK) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/paths/][%d] circuitsCircuitTerminationsPathsOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsPathsOK) String() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/paths/][%d] circuitsCircuitTerminationsPathsOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTerminationsPathsOK) GetPayload() *models.CircuitTermination {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsPathsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitTermination)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsPathsBadRequest creates a CircuitsCircuitTerminationsPathsBadRequest with default headers values
func NewCircuitsCircuitTerminationsPathsBadRequest() *CircuitsCircuitTerminationsPathsBadRequest {
	return &CircuitsCircuitTerminationsPathsBadRequest{}
}

/*
CircuitsCircuitTerminationsPathsBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CircuitsCircuitTerminationsPathsBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this circuits circuit terminations paths bad request response has a 2xx status code
func (o *CircuitsCircuitTerminationsPathsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this circuits circuit terminations paths bad request response has a 3xx status code
func (o *CircuitsCircuitTerminationsPathsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit terminations paths bad request response has a 4xx status code
func (o *CircuitsCircuitTerminationsPathsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this circuits circuit terminations paths bad request response has a 5xx status code
func (o *CircuitsCircuitTerminationsPathsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit terminations paths bad request response a status code equal to that given
func (o *CircuitsCircuitTerminationsPathsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CircuitsCircuitTerminationsPathsBadRequest) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/paths/][%d] circuitsCircuitTerminationsPathsBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsCircuitTerminationsPathsBadRequest) String() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/paths/][%d] circuitsCircuitTerminationsPathsBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsCircuitTerminationsPathsBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsPathsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTerminationsPathsDefault creates a CircuitsCircuitTerminationsPathsDefault with default headers values
func NewCircuitsCircuitTerminationsPathsDefault(code int) *CircuitsCircuitTerminationsPathsDefault {
	return &CircuitsCircuitTerminationsPathsDefault{
		_statusCode: code,
	}
}

/*
CircuitsCircuitTerminationsPathsDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type CircuitsCircuitTerminationsPathsDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits circuit terminations paths default response
func (o *CircuitsCircuitTerminationsPathsDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits circuit terminations paths default response has a 2xx status code
func (o *CircuitsCircuitTerminationsPathsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits circuit terminations paths default response has a 3xx status code
func (o *CircuitsCircuitTerminationsPathsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits circuit terminations paths default response has a 4xx status code
func (o *CircuitsCircuitTerminationsPathsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits circuit terminations paths default response has a 5xx status code
func (o *CircuitsCircuitTerminationsPathsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits circuit terminations paths default response a status code equal to that given
func (o *CircuitsCircuitTerminationsPathsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsCircuitTerminationsPathsDefault) Error() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/paths/][%d] circuits_circuit-terminations_paths default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTerminationsPathsDefault) String() string {
	return fmt.Sprintf("[GET /circuits/circuit-terminations/{id}/paths/][%d] circuits_circuit-terminations_paths default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsCircuitTerminationsPathsDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTerminationsPathsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
