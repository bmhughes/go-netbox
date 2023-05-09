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

// ExtrasTagsBulkUpdateReader is a Reader for the ExtrasTagsBulkUpdate structure.
type ExtrasTagsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasTagsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasTagsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasTagsBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasTagsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasTagsBulkUpdateOK creates a ExtrasTagsBulkUpdateOK with default headers values
func NewExtrasTagsBulkUpdateOK() *ExtrasTagsBulkUpdateOK {
	return &ExtrasTagsBulkUpdateOK{}
}

/*
ExtrasTagsBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasTagsBulkUpdateOK extras tags bulk update o k
*/
type ExtrasTagsBulkUpdateOK struct {
	Payload *models.Tag
}

// IsSuccess returns true when this extras tags bulk update o k response has a 2xx status code
func (o *ExtrasTagsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras tags bulk update o k response has a 3xx status code
func (o *ExtrasTagsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras tags bulk update o k response has a 4xx status code
func (o *ExtrasTagsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras tags bulk update o k response has a 5xx status code
func (o *ExtrasTagsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras tags bulk update o k response a status code equal to that given
func (o *ExtrasTagsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasTagsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extrasTagsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasTagsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extrasTagsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasTagsBulkUpdateOK) GetPayload() *models.Tag {
	return o.Payload
}

func (o *ExtrasTagsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasTagsBulkUpdateBadRequest creates a ExtrasTagsBulkUpdateBadRequest with default headers values
func NewExtrasTagsBulkUpdateBadRequest() *ExtrasTagsBulkUpdateBadRequest {
	return &ExtrasTagsBulkUpdateBadRequest{}
}

/*
ExtrasTagsBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasTagsBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras tags bulk update bad request response has a 2xx status code
func (o *ExtrasTagsBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras tags bulk update bad request response has a 3xx status code
func (o *ExtrasTagsBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras tags bulk update bad request response has a 4xx status code
func (o *ExtrasTagsBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras tags bulk update bad request response has a 5xx status code
func (o *ExtrasTagsBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras tags bulk update bad request response a status code equal to that given
func (o *ExtrasTagsBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasTagsBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extrasTagsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasTagsBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extrasTagsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasTagsBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasTagsBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasTagsBulkUpdateDefault creates a ExtrasTagsBulkUpdateDefault with default headers values
func NewExtrasTagsBulkUpdateDefault(code int) *ExtrasTagsBulkUpdateDefault {
	return &ExtrasTagsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasTagsBulkUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasTagsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras tags bulk update default response
func (o *ExtrasTagsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras tags bulk update default response has a 2xx status code
func (o *ExtrasTagsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras tags bulk update default response has a 3xx status code
func (o *ExtrasTagsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras tags bulk update default response has a 4xx status code
func (o *ExtrasTagsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras tags bulk update default response has a 5xx status code
func (o *ExtrasTagsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras tags bulk update default response a status code equal to that given
func (o *ExtrasTagsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasTagsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extras_tags_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasTagsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/tags/][%d] extras_tags_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasTagsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasTagsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
