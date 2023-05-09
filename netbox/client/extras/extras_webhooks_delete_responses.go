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

// ExtrasWebhooksDeleteReader is a Reader for the ExtrasWebhooksDelete structure.
type ExtrasWebhooksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasWebhooksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasWebhooksDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasWebhooksDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasWebhooksDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasWebhooksDeleteNoContent creates a ExtrasWebhooksDeleteNoContent with default headers values
func NewExtrasWebhooksDeleteNoContent() *ExtrasWebhooksDeleteNoContent {
	return &ExtrasWebhooksDeleteNoContent{}
}

/*
ExtrasWebhooksDeleteNoContent describes a response with status code 204, with default header values.

ExtrasWebhooksDeleteNoContent extras webhooks delete no content
*/
type ExtrasWebhooksDeleteNoContent struct {
}

// IsSuccess returns true when this extras webhooks delete no content response has a 2xx status code
func (o *ExtrasWebhooksDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras webhooks delete no content response has a 3xx status code
func (o *ExtrasWebhooksDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras webhooks delete no content response has a 4xx status code
func (o *ExtrasWebhooksDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras webhooks delete no content response has a 5xx status code
func (o *ExtrasWebhooksDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras webhooks delete no content response a status code equal to that given
func (o *ExtrasWebhooksDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ExtrasWebhooksDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/webhooks/{id}/][%d] extrasWebhooksDeleteNoContent ", 204)
}

func (o *ExtrasWebhooksDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/webhooks/{id}/][%d] extrasWebhooksDeleteNoContent ", 204)
}

func (o *ExtrasWebhooksDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasWebhooksDeleteBadRequest creates a ExtrasWebhooksDeleteBadRequest with default headers values
func NewExtrasWebhooksDeleteBadRequest() *ExtrasWebhooksDeleteBadRequest {
	return &ExtrasWebhooksDeleteBadRequest{}
}

/*
ExtrasWebhooksDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasWebhooksDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras webhooks delete bad request response has a 2xx status code
func (o *ExtrasWebhooksDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras webhooks delete bad request response has a 3xx status code
func (o *ExtrasWebhooksDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras webhooks delete bad request response has a 4xx status code
func (o *ExtrasWebhooksDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras webhooks delete bad request response has a 5xx status code
func (o *ExtrasWebhooksDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras webhooks delete bad request response a status code equal to that given
func (o *ExtrasWebhooksDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasWebhooksDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /extras/webhooks/{id}/][%d] extrasWebhooksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasWebhooksDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /extras/webhooks/{id}/][%d] extrasWebhooksDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasWebhooksDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasWebhooksDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasWebhooksDeleteDefault creates a ExtrasWebhooksDeleteDefault with default headers values
func NewExtrasWebhooksDeleteDefault(code int) *ExtrasWebhooksDeleteDefault {
	return &ExtrasWebhooksDeleteDefault{
		_statusCode: code,
	}
}

/*
ExtrasWebhooksDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasWebhooksDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras webhooks delete default response
func (o *ExtrasWebhooksDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras webhooks delete default response has a 2xx status code
func (o *ExtrasWebhooksDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras webhooks delete default response has a 3xx status code
func (o *ExtrasWebhooksDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras webhooks delete default response has a 4xx status code
func (o *ExtrasWebhooksDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras webhooks delete default response has a 5xx status code
func (o *ExtrasWebhooksDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras webhooks delete default response a status code equal to that given
func (o *ExtrasWebhooksDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasWebhooksDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/webhooks/{id}/][%d] extras_webhooks_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasWebhooksDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /extras/webhooks/{id}/][%d] extras_webhooks_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasWebhooksDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasWebhooksDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
