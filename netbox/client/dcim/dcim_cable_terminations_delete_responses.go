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
)

// DcimCableTerminationsDeleteReader is a Reader for the DcimCableTerminationsDelete structure.
type DcimCableTerminationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCableTerminationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimCableTerminationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimCableTerminationsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimCableTerminationsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCableTerminationsDeleteNoContent creates a DcimCableTerminationsDeleteNoContent with default headers values
func NewDcimCableTerminationsDeleteNoContent() *DcimCableTerminationsDeleteNoContent {
	return &DcimCableTerminationsDeleteNoContent{}
}

/*
DcimCableTerminationsDeleteNoContent describes a response with status code 204, with default header values.

DcimCableTerminationsDeleteNoContent dcim cable terminations delete no content
*/
type DcimCableTerminationsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim cable terminations delete no content response has a 2xx status code
func (o *DcimCableTerminationsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cable terminations delete no content response has a 3xx status code
func (o *DcimCableTerminationsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations delete no content response has a 4xx status code
func (o *DcimCableTerminationsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cable terminations delete no content response has a 5xx status code
func (o *DcimCableTerminationsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations delete no content response a status code equal to that given
func (o *DcimCableTerminationsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimCableTerminationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsDeleteNoContent ", 204)
}

func (o *DcimCableTerminationsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsDeleteNoContent ", 204)
}

func (o *DcimCableTerminationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimCableTerminationsDeleteBadRequest creates a DcimCableTerminationsDeleteBadRequest with default headers values
func NewDcimCableTerminationsDeleteBadRequest() *DcimCableTerminationsDeleteBadRequest {
	return &DcimCableTerminationsDeleteBadRequest{}
}

/*
DcimCableTerminationsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimCableTerminationsDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim cable terminations delete bad request response has a 2xx status code
func (o *DcimCableTerminationsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim cable terminations delete bad request response has a 3xx status code
func (o *DcimCableTerminationsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cable terminations delete bad request response has a 4xx status code
func (o *DcimCableTerminationsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim cable terminations delete bad request response has a 5xx status code
func (o *DcimCableTerminationsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cable terminations delete bad request response a status code equal to that given
func (o *DcimCableTerminationsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimCableTerminationsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimCableTerminationsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/cable-terminations/{id}/][%d] dcimCableTerminationsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimCableTerminationsDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCableTerminationsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCableTerminationsDeleteDefault creates a DcimCableTerminationsDeleteDefault with default headers values
func NewDcimCableTerminationsDeleteDefault(code int) *DcimCableTerminationsDeleteDefault {
	return &DcimCableTerminationsDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimCableTerminationsDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimCableTerminationsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cable terminations delete default response
func (o *DcimCableTerminationsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim cable terminations delete default response has a 2xx status code
func (o *DcimCableTerminationsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cable terminations delete default response has a 3xx status code
func (o *DcimCableTerminationsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cable terminations delete default response has a 4xx status code
func (o *DcimCableTerminationsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cable terminations delete default response has a 5xx status code
func (o *DcimCableTerminationsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cable terminations delete default response a status code equal to that given
func (o *DcimCableTerminationsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimCableTerminationsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/cable-terminations/{id}/][%d] dcim_cable-terminations_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCableTerminationsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/cable-terminations/{id}/][%d] dcim_cable-terminations_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCableTerminationsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCableTerminationsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
