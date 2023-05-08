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

// ExtrasScriptsListReader is a Reader for the ExtrasScriptsList structure.
type ExtrasScriptsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasScriptsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasScriptsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasScriptsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasScriptsListOK creates a ExtrasScriptsListOK with default headers values
func NewExtrasScriptsListOK() *ExtrasScriptsListOK {
	return &ExtrasScriptsListOK{}
}

/*
ExtrasScriptsListOK describes a response with status code 200, with default header values.

ExtrasScriptsListOK extras scripts list o k
*/
type ExtrasScriptsListOK struct {
}

// IsSuccess returns true when this extras scripts list o k response has a 2xx status code
func (o *ExtrasScriptsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras scripts list o k response has a 3xx status code
func (o *ExtrasScriptsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras scripts list o k response has a 4xx status code
func (o *ExtrasScriptsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras scripts list o k response has a 5xx status code
func (o *ExtrasScriptsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras scripts list o k response a status code equal to that given
func (o *ExtrasScriptsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasScriptsListOK) Error() string {
	return fmt.Sprintf("[GET /extras/scripts/][%d] extrasScriptsListOK ", 200)
}

func (o *ExtrasScriptsListOK) String() string {
	return fmt.Sprintf("[GET /extras/scripts/][%d] extrasScriptsListOK ", 200)
}

func (o *ExtrasScriptsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasScriptsListBadRequest creates a ExtrasScriptsListBadRequest with default headers values
func NewExtrasScriptsListBadRequest() *ExtrasScriptsListBadRequest {
	return &ExtrasScriptsListBadRequest{}
}

/*
ExtrasScriptsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasScriptsListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras scripts list bad request response has a 2xx status code
func (o *ExtrasScriptsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras scripts list bad request response has a 3xx status code
func (o *ExtrasScriptsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras scripts list bad request response has a 4xx status code
func (o *ExtrasScriptsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras scripts list bad request response has a 5xx status code
func (o *ExtrasScriptsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras scripts list bad request response a status code equal to that given
func (o *ExtrasScriptsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasScriptsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /extras/scripts/][%d] extrasScriptsListBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasScriptsListBadRequest) String() string {
	return fmt.Sprintf("[GET /extras/scripts/][%d] extrasScriptsListBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasScriptsListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasScriptsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
