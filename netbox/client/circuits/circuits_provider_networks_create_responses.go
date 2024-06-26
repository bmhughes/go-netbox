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

// CircuitsProviderNetworksCreateReader is a Reader for the CircuitsProviderNetworksCreate structure.
type CircuitsProviderNetworksCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProviderNetworksCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCircuitsProviderNetworksCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCircuitsProviderNetworksCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewCircuitsProviderNetworksCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCircuitsProviderNetworksCreateCreated creates a CircuitsProviderNetworksCreateCreated with default headers values
func NewCircuitsProviderNetworksCreateCreated() *CircuitsProviderNetworksCreateCreated {
	return &CircuitsProviderNetworksCreateCreated{}
}

/*
CircuitsProviderNetworksCreateCreated describes a response with status code 201, with default header values.

CircuitsProviderNetworksCreateCreated circuits provider networks create created
*/
type CircuitsProviderNetworksCreateCreated struct {
	Payload *models.ProviderNetwork
}

// IsSuccess returns true when this circuits provider networks create created response has a 2xx status code
func (o *CircuitsProviderNetworksCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits provider networks create created response has a 3xx status code
func (o *CircuitsProviderNetworksCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits provider networks create created response has a 4xx status code
func (o *CircuitsProviderNetworksCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits provider networks create created response has a 5xx status code
func (o *CircuitsProviderNetworksCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits provider networks create created response a status code equal to that given
func (o *CircuitsProviderNetworksCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CircuitsProviderNetworksCreateCreated) Error() string {
	return fmt.Sprintf("[POST /circuits/provider-networks/][%d] circuitsProviderNetworksCreateCreated  %+v", 201, o.Payload)
}

func (o *CircuitsProviderNetworksCreateCreated) String() string {
	return fmt.Sprintf("[POST /circuits/provider-networks/][%d] circuitsProviderNetworksCreateCreated  %+v", 201, o.Payload)
}

func (o *CircuitsProviderNetworksCreateCreated) GetPayload() *models.ProviderNetwork {
	return o.Payload
}

func (o *CircuitsProviderNetworksCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProviderNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProviderNetworksCreateBadRequest creates a CircuitsProviderNetworksCreateBadRequest with default headers values
func NewCircuitsProviderNetworksCreateBadRequest() *CircuitsProviderNetworksCreateBadRequest {
	return &CircuitsProviderNetworksCreateBadRequest{}
}

/*
CircuitsProviderNetworksCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CircuitsProviderNetworksCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this circuits provider networks create bad request response has a 2xx status code
func (o *CircuitsProviderNetworksCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this circuits provider networks create bad request response has a 3xx status code
func (o *CircuitsProviderNetworksCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits provider networks create bad request response has a 4xx status code
func (o *CircuitsProviderNetworksCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this circuits provider networks create bad request response has a 5xx status code
func (o *CircuitsProviderNetworksCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits provider networks create bad request response a status code equal to that given
func (o *CircuitsProviderNetworksCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CircuitsProviderNetworksCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /circuits/provider-networks/][%d] circuitsProviderNetworksCreateBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsProviderNetworksCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /circuits/provider-networks/][%d] circuitsProviderNetworksCreateBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsProviderNetworksCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProviderNetworksCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProviderNetworksCreateDefault creates a CircuitsProviderNetworksCreateDefault with default headers values
func NewCircuitsProviderNetworksCreateDefault(code int) *CircuitsProviderNetworksCreateDefault {
	return &CircuitsProviderNetworksCreateDefault{
		_statusCode: code,
	}
}

/*
CircuitsProviderNetworksCreateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type CircuitsProviderNetworksCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the circuits provider networks create default response
func (o *CircuitsProviderNetworksCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this circuits provider networks create default response has a 2xx status code
func (o *CircuitsProviderNetworksCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this circuits provider networks create default response has a 3xx status code
func (o *CircuitsProviderNetworksCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this circuits provider networks create default response has a 4xx status code
func (o *CircuitsProviderNetworksCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this circuits provider networks create default response has a 5xx status code
func (o *CircuitsProviderNetworksCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this circuits provider networks create default response a status code equal to that given
func (o *CircuitsProviderNetworksCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CircuitsProviderNetworksCreateDefault) Error() string {
	return fmt.Sprintf("[POST /circuits/provider-networks/][%d] circuits_provider-networks_create default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProviderNetworksCreateDefault) String() string {
	return fmt.Sprintf("[POST /circuits/provider-networks/][%d] circuits_provider-networks_create default  %+v", o._statusCode, o.Payload)
}

func (o *CircuitsProviderNetworksCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProviderNetworksCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
