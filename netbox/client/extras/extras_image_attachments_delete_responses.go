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

// ExtrasImageAttachmentsDeleteReader is a Reader for the ExtrasImageAttachmentsDelete structure.
type ExtrasImageAttachmentsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasImageAttachmentsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasImageAttachmentsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasImageAttachmentsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsDeleteNoContent creates a ExtrasImageAttachmentsDeleteNoContent with default headers values
func NewExtrasImageAttachmentsDeleteNoContent() *ExtrasImageAttachmentsDeleteNoContent {
	return &ExtrasImageAttachmentsDeleteNoContent{}
}

/*
ExtrasImageAttachmentsDeleteNoContent describes a response with status code 204, with default header values.

ExtrasImageAttachmentsDeleteNoContent extras image attachments delete no content
*/
type ExtrasImageAttachmentsDeleteNoContent struct {
}

// IsSuccess returns true when this extras image attachments delete no content response has a 2xx status code
func (o *ExtrasImageAttachmentsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments delete no content response has a 3xx status code
func (o *ExtrasImageAttachmentsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments delete no content response has a 4xx status code
func (o *ExtrasImageAttachmentsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments delete no content response has a 5xx status code
func (o *ExtrasImageAttachmentsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments delete no content response a status code equal to that given
func (o *ExtrasImageAttachmentsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *ExtrasImageAttachmentsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extrasImageAttachmentsDeleteNoContent ", 204)
}

func (o *ExtrasImageAttachmentsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extrasImageAttachmentsDeleteNoContent ", 204)
}

func (o *ExtrasImageAttachmentsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasImageAttachmentsDeleteBadRequest creates a ExtrasImageAttachmentsDeleteBadRequest with default headers values
func NewExtrasImageAttachmentsDeleteBadRequest() *ExtrasImageAttachmentsDeleteBadRequest {
	return &ExtrasImageAttachmentsDeleteBadRequest{}
}

/*
ExtrasImageAttachmentsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasImageAttachmentsDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras image attachments delete bad request response has a 2xx status code
func (o *ExtrasImageAttachmentsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras image attachments delete bad request response has a 3xx status code
func (o *ExtrasImageAttachmentsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments delete bad request response has a 4xx status code
func (o *ExtrasImageAttachmentsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras image attachments delete bad request response has a 5xx status code
func (o *ExtrasImageAttachmentsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments delete bad request response a status code equal to that given
func (o *ExtrasImageAttachmentsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasImageAttachmentsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extrasImageAttachmentsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasImageAttachmentsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extrasImageAttachmentsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasImageAttachmentsDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsDeleteDefault creates a ExtrasImageAttachmentsDeleteDefault with default headers values
func NewExtrasImageAttachmentsDeleteDefault(code int) *ExtrasImageAttachmentsDeleteDefault {
	return &ExtrasImageAttachmentsDeleteDefault{
		_statusCode: code,
	}
}

/*
ExtrasImageAttachmentsDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasImageAttachmentsDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras image attachments delete default response
func (o *ExtrasImageAttachmentsDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras image attachments delete default response has a 2xx status code
func (o *ExtrasImageAttachmentsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras image attachments delete default response has a 3xx status code
func (o *ExtrasImageAttachmentsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras image attachments delete default response has a 4xx status code
func (o *ExtrasImageAttachmentsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras image attachments delete default response has a 5xx status code
func (o *ExtrasImageAttachmentsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras image attachments delete default response a status code equal to that given
func (o *ExtrasImageAttachmentsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasImageAttachmentsDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extras_image-attachments_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/{id}/][%d] extras_image-attachments_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
