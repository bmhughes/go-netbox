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

// CircuitsProvidersReadReader is a Reader for the CircuitsProvidersRead structure.
type CircuitsProvidersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProvidersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProvidersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCircuitsProvidersReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCircuitsProvidersReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProvidersReadOK creates a CircuitsProvidersReadOK with default headers values
func NewCircuitsProvidersReadOK() *CircuitsProvidersReadOK {
	return &CircuitsProvidersReadOK{}
}

/*
CircuitsProvidersReadOK describes a response with status code 200, with default header values.

CircuitsProvidersReadOK circuits providers read o k
*/
type CircuitsProvidersReadOK struct {
	Payload *models.Provider
}

// IsSuccess returns true when this circuits providers read o k response has a 2xx status code
func (o *CircuitsProvidersReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits providers read o k response has a 3xx status code
func (o *CircuitsProvidersReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits providers read o k response has a 4xx status code
func (o *CircuitsProvidersReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits providers read o k response has a 5xx status code
func (o *CircuitsProvidersReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits providers read o k response a status code equal to that given
func (o *CircuitsProvidersReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsProvidersReadOK) Error() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuitsProvidersReadOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersReadOK) String() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuitsProvidersReadOK  %+v", 200, o.Payload)
}

func (o *CircuitsProvidersReadOK) GetPayload() *models.Provider {
	return o.Payload
}

func (o *CircuitsProvidersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Provider)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProvidersReadBadRequest creates a CircuitsProvidersReadBadRequest with default headers values
func NewCircuitsProvidersReadBadRequest() *CircuitsProvidersReadBadRequest {
	return &CircuitsProvidersReadBadRequest{}
}

/*
CircuitsProvidersReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CircuitsProvidersReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this circuits providers read bad request response has a 2xx status code
func (o *CircuitsProvidersReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this circuits providers read bad request response has a 3xx status code
func (o *CircuitsProvidersReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits providers read bad request response has a 4xx status code
func (o *CircuitsProvidersReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this circuits providers read bad request response has a 5xx status code
func (o *CircuitsProvidersReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits providers read bad request response a status code equal to that given
func (o *CircuitsProvidersReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CircuitsProvidersReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuitsProvidersReadBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsProvidersReadBadRequest) String() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuitsProvidersReadBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsProvidersReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProvidersReadDefault creates a CircuitsProvidersReadDefault with default headers values
func NewCircuitsProvidersReadDefault(code int) *CircuitsProvidersReadDefault {
	return &CircuitsProvidersReadDefault{
		_statusCode: code,
	}
}

/*
CircuitsProvidersReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type CircuitsProvidersReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits providers read default response
func (o *CircuitsProvidersReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits providers read default response has a 2xx status code
func (o *CircuitsProvidersReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits providers read default response has a 3xx status code
func (o *CircuitsProvidersReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits providers read default response has a 4xx status code
func (o *CircuitsProvidersReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits providers read default response has a 5xx status code
func (o *CircuitsProvidersReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits providers read default response a status code equal to that given
func (o *CircuitsProvidersReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsProvidersReadDefault) Error() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuits_providers_read default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersReadDefault) String() string {
	return fmt.Sprintf("[GET /circuits/providers/{id}/][%d] circuits_providers_read default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProvidersReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProvidersReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
