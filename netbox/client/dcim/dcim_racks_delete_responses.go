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

// DcimRacksDeleteReader is a Reader for the DcimRacksDelete structure.
type DcimRacksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRacksDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimRacksDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimRacksDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRacksDeleteNoContent creates a DcimRacksDeleteNoContent with default headers values
func NewDcimRacksDeleteNoContent() *DcimRacksDeleteNoContent {
	return &DcimRacksDeleteNoContent{}
}

/*
DcimRacksDeleteNoContent describes a response with status code 204, with default header values.

DcimRacksDeleteNoContent dcim racks delete no content
*/
type DcimRacksDeleteNoContent struct {
}

// IsSuccess returns true when this dcim racks delete no content response has a 2xx status code
func (o *DcimRacksDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim racks delete no content response has a 3xx status code
func (o *DcimRacksDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim racks delete no content response has a 4xx status code
func (o *DcimRacksDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim racks delete no content response has a 5xx status code
func (o *DcimRacksDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim racks delete no content response a status code equal to that given
func (o *DcimRacksDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimRacksDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcimRacksDeleteNoContent ", 204)
}

func (o *DcimRacksDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcimRacksDeleteNoContent ", 204)
}

func (o *DcimRacksDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRacksDeleteBadRequest creates a DcimRacksDeleteBadRequest with default headers values
func NewDcimRacksDeleteBadRequest() *DcimRacksDeleteBadRequest {
	return &DcimRacksDeleteBadRequest{}
}

/*
DcimRacksDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimRacksDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim racks delete bad request response has a 2xx status code
func (o *DcimRacksDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim racks delete bad request response has a 3xx status code
func (o *DcimRacksDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim racks delete bad request response has a 4xx status code
func (o *DcimRacksDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim racks delete bad request response has a 5xx status code
func (o *DcimRacksDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim racks delete bad request response a status code equal to that given
func (o *DcimRacksDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimRacksDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcimRacksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRacksDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcimRacksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRacksDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRacksDeleteDefault creates a DcimRacksDeleteDefault with default headers values
func NewDcimRacksDeleteDefault(code int) *DcimRacksDeleteDefault {
	return &DcimRacksDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimRacksDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimRacksDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim racks delete default response
func (o *DcimRacksDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim racks delete default response has a 2xx status code
func (o *DcimRacksDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim racks delete default response has a 3xx status code
func (o *DcimRacksDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim racks delete default response has a 4xx status code
func (o *DcimRacksDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim racks delete default response has a 5xx status code
func (o *DcimRacksDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim racks delete default response a status code equal to that given
func (o *DcimRacksDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRacksDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcim_racks_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcim_racks_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
