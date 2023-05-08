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
)

// ExtrasReportsReadReader is a Reader for the ExtrasReportsRead structure.
type ExtrasReportsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasReportsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasReportsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasReportsReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasReportsReadOK creates a ExtrasReportsReadOK with default headers values
func NewExtrasReportsReadOK() *ExtrasReportsReadOK {
	return &ExtrasReportsReadOK{}
}

/*
ExtrasReportsReadOK describes a response with status code 200, with default header values.

ExtrasReportsReadOK extras reports read o k
*/
type ExtrasReportsReadOK struct {
}

// IsSuccess returns true when this extras reports read o k response has a 2xx status code
func (o *ExtrasReportsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras reports read o k response has a 3xx status code
func (o *ExtrasReportsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras reports read o k response has a 4xx status code
func (o *ExtrasReportsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras reports read o k response has a 5xx status code
func (o *ExtrasReportsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras reports read o k response a status code equal to that given
func (o *ExtrasReportsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasReportsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/reports/{id}/][%d] extrasReportsReadOK ", 200)
}

func (o *ExtrasReportsReadOK) String() string {
	return fmt.Sprintf("[GET /extras/reports/{id}/][%d] extrasReportsReadOK ", 200)
}

func (o *ExtrasReportsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasReportsReadBadRequest creates a ExtrasReportsReadBadRequest with default headers values
func NewExtrasReportsReadBadRequest() *ExtrasReportsReadBadRequest {
	return &ExtrasReportsReadBadRequest{}
}

/*
ExtrasReportsReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasReportsReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras reports read bad request response has a 2xx status code
func (o *ExtrasReportsReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras reports read bad request response has a 3xx status code
func (o *ExtrasReportsReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras reports read bad request response has a 4xx status code
func (o *ExtrasReportsReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras reports read bad request response has a 5xx status code
func (o *ExtrasReportsReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras reports read bad request response a status code equal to that given
func (o *ExtrasReportsReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasReportsReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /extras/reports/{id}/][%d] extrasReportsReadBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasReportsReadBadRequest) String() string {
	return fmt.Sprintf("[GET /extras/reports/{id}/][%d] extrasReportsReadBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasReportsReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasReportsReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
