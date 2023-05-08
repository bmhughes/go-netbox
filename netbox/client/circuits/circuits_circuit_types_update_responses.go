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

// CircuitsCircuitTypesUpdateReader is a Reader for the CircuitsCircuitTypesUpdate structure.
type CircuitsCircuitTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsCircuitTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsCircuitTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCircuitsCircuitTypesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsCircuitTypesUpdateOK creates a CircuitsCircuitTypesUpdateOK with default headers values
func NewCircuitsCircuitTypesUpdateOK() *CircuitsCircuitTypesUpdateOK {
	return &CircuitsCircuitTypesUpdateOK{}
}

/*
CircuitsCircuitTypesUpdateOK describes a response with status code 200, with default header values.

CircuitsCircuitTypesUpdateOK circuits circuit types update o k
*/
type CircuitsCircuitTypesUpdateOK struct {
	Payload *models.CircuitType
}

// IsSuccess returns true when this circuits circuit types update o k response has a 2xx status code
func (o *CircuitsCircuitTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits circuit types update o k response has a 3xx status code
func (o *CircuitsCircuitTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit types update o k response has a 4xx status code
func (o *CircuitsCircuitTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits circuit types update o k response has a 5xx status code
func (o *CircuitsCircuitTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit types update o k response a status code equal to that given
func (o *CircuitsCircuitTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsCircuitTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTypesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsCircuitTypesUpdateOK) GetPayload() *models.CircuitType {
	return o.Payload
}

func (o *CircuitsCircuitTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CircuitType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsCircuitTypesUpdateBadRequest creates a CircuitsCircuitTypesUpdateBadRequest with default headers values
func NewCircuitsCircuitTypesUpdateBadRequest() *CircuitsCircuitTypesUpdateBadRequest {
	return &CircuitsCircuitTypesUpdateBadRequest{}
}

/*
CircuitsCircuitTypesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CircuitsCircuitTypesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this circuits circuit types update bad request response has a 2xx status code
func (o *CircuitsCircuitTypesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this circuits circuit types update bad request response has a 3xx status code
func (o *CircuitsCircuitTypesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits circuit types update bad request response has a 4xx status code
func (o *CircuitsCircuitTypesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this circuits circuit types update bad request response has a 5xx status code
func (o *CircuitsCircuitTypesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits circuit types update bad request response a status code equal to that given
func (o *CircuitsCircuitTypesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CircuitsCircuitTypesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsCircuitTypesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /circuits/circuit-types/{id}/][%d] circuitsCircuitTypesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsCircuitTypesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsCircuitTypesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
