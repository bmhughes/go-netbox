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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// ExtrasContentTypesReadReader is a Reader for the ExtrasContentTypesRead structure.
type ExtrasContentTypesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasContentTypesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasContentTypesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasContentTypesReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasContentTypesReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasContentTypesReadOK creates a ExtrasContentTypesReadOK with default headers values
func NewExtrasContentTypesReadOK() *ExtrasContentTypesReadOK {
	return &ExtrasContentTypesReadOK{}
}

/*
ExtrasContentTypesReadOK describes a response with status code 200, with default header values.

ExtrasContentTypesReadOK extras content types read o k
*/
type ExtrasContentTypesReadOK struct {
	Payload *models.ContentType
}

// IsSuccess returns true when this extras content types read o k response has a 2xx status code
func (o *ExtrasContentTypesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras content types read o k response has a 3xx status code
func (o *ExtrasContentTypesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras content types read o k response has a 4xx status code
func (o *ExtrasContentTypesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras content types read o k response has a 5xx status code
func (o *ExtrasContentTypesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras content types read o k response a status code equal to that given
func (o *ExtrasContentTypesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasContentTypesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extrasContentTypesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasContentTypesReadOK) String() string {
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extrasContentTypesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasContentTypesReadOK) GetPayload() *models.ContentType {
	return o.Payload
}

func (o *ExtrasContentTypesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasContentTypesReadBadRequest creates a ExtrasContentTypesReadBadRequest with default headers values
func NewExtrasContentTypesReadBadRequest() *ExtrasContentTypesReadBadRequest {
	return &ExtrasContentTypesReadBadRequest{}
}

/*
ExtrasContentTypesReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasContentTypesReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras content types read bad request response has a 2xx status code
func (o *ExtrasContentTypesReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras content types read bad request response has a 3xx status code
func (o *ExtrasContentTypesReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras content types read bad request response has a 4xx status code
func (o *ExtrasContentTypesReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras content types read bad request response has a 5xx status code
func (o *ExtrasContentTypesReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras content types read bad request response a status code equal to that given
func (o *ExtrasContentTypesReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasContentTypesReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extrasContentTypesReadBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasContentTypesReadBadRequest) String() string {
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extrasContentTypesReadBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasContentTypesReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasContentTypesReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasContentTypesReadDefault creates a ExtrasContentTypesReadDefault with default headers values
func NewExtrasContentTypesReadDefault(code int) *ExtrasContentTypesReadDefault {
	return &ExtrasContentTypesReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasContentTypesReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasContentTypesReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras content types read default response
func (o *ExtrasContentTypesReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras content types read default response has a 2xx status code
func (o *ExtrasContentTypesReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras content types read default response has a 3xx status code
func (o *ExtrasContentTypesReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras content types read default response has a 4xx status code
func (o *ExtrasContentTypesReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras content types read default response has a 5xx status code
func (o *ExtrasContentTypesReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras content types read default response a status code equal to that given
func (o *ExtrasContentTypesReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasContentTypesReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extras_content-types_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasContentTypesReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/content-types/{id}/][%d] extras_content-types_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasContentTypesReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasContentTypesReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
