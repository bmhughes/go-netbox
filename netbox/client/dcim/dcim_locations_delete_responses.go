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

// DcimLocationsDeleteReader is a Reader for the DcimLocationsDelete structure.
type DcimLocationsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimLocationsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimLocationsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimLocationsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimLocationsDeleteNoContent creates a DcimLocationsDeleteNoContent with default headers values
func NewDcimLocationsDeleteNoContent() *DcimLocationsDeleteNoContent {
	return &DcimLocationsDeleteNoContent{}
}

/*
DcimLocationsDeleteNoContent describes a response with status code 204, with default header values.

DcimLocationsDeleteNoContent dcim locations delete no content
*/
type DcimLocationsDeleteNoContent struct {
}

// IsSuccess returns true when this dcim locations delete no content response has a 2xx status code
func (o *DcimLocationsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations delete no content response has a 3xx status code
func (o *DcimLocationsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations delete no content response has a 4xx status code
func (o *DcimLocationsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations delete no content response has a 5xx status code
func (o *DcimLocationsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations delete no content response a status code equal to that given
func (o *DcimLocationsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimLocationsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/locations/{id}/][%d] dcimLocationsDeleteNoContent ", 204)
}

func (o *DcimLocationsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/locations/{id}/][%d] dcimLocationsDeleteNoContent ", 204)
}

func (o *DcimLocationsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimLocationsDeleteBadRequest creates a DcimLocationsDeleteBadRequest with default headers values
func NewDcimLocationsDeleteBadRequest() *DcimLocationsDeleteBadRequest {
	return &DcimLocationsDeleteBadRequest{}
}

/*
DcimLocationsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimLocationsDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim locations delete bad request response has a 2xx status code
func (o *DcimLocationsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim locations delete bad request response has a 3xx status code
func (o *DcimLocationsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations delete bad request response has a 4xx status code
func (o *DcimLocationsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim locations delete bad request response has a 5xx status code
func (o *DcimLocationsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations delete bad request response a status code equal to that given
func (o *DcimLocationsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimLocationsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/locations/{id}/][%d] dcimLocationsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimLocationsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/locations/{id}/][%d] dcimLocationsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimLocationsDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimLocationsDeleteDefault creates a DcimLocationsDeleteDefault with default headers values
func NewDcimLocationsDeleteDefault(code int) *DcimLocationsDeleteDefault {
	return &DcimLocationsDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimLocationsDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimLocationsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim locations delete default response
func (o *DcimLocationsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim locations delete default response has a 2xx status code
func (o *DcimLocationsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim locations delete default response has a 3xx status code
func (o *DcimLocationsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim locations delete default response has a 4xx status code
func (o *DcimLocationsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim locations delete default response has a 5xx status code
func (o *DcimLocationsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim locations delete default response a status code equal to that given
func (o *DcimLocationsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimLocationsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/locations/{id}/][%d] dcim_locations_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/locations/{id}/][%d] dcim_locations_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
