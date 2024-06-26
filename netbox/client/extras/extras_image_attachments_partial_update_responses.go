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

// ExtrasImageAttachmentsPartialUpdateReader is a Reader for the ExtrasImageAttachmentsPartialUpdate structure.
type ExtrasImageAttachmentsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasImageAttachmentsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasImageAttachmentsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsPartialUpdateOK creates a ExtrasImageAttachmentsPartialUpdateOK with default headers values
func NewExtrasImageAttachmentsPartialUpdateOK() *ExtrasImageAttachmentsPartialUpdateOK {
	return &ExtrasImageAttachmentsPartialUpdateOK{}
}

/*
ExtrasImageAttachmentsPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsPartialUpdateOK extras image attachments partial update o k
*/
type ExtrasImageAttachmentsPartialUpdateOK struct {
	Payload *models.ImageAttachment
}

// IsSuccess returns true when this extras image attachments partial update o k response has a 2xx status code
func (o *ExtrasImageAttachmentsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments partial update o k response has a 3xx status code
func (o *ExtrasImageAttachmentsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments partial update o k response has a 4xx status code
func (o *ExtrasImageAttachmentsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments partial update o k response has a 5xx status code
func (o *ExtrasImageAttachmentsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments partial update o k response a status code equal to that given
func (o *ExtrasImageAttachmentsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extrasImageAttachmentsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extrasImageAttachmentsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsPartialUpdateBadRequest creates a ExtrasImageAttachmentsPartialUpdateBadRequest with default headers values
func NewExtrasImageAttachmentsPartialUpdateBadRequest() *ExtrasImageAttachmentsPartialUpdateBadRequest {
	return &ExtrasImageAttachmentsPartialUpdateBadRequest{}
}

/*
ExtrasImageAttachmentsPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasImageAttachmentsPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras image attachments partial update bad request response has a 2xx status code
func (o *ExtrasImageAttachmentsPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras image attachments partial update bad request response has a 3xx status code
func (o *ExtrasImageAttachmentsPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments partial update bad request response has a 4xx status code
func (o *ExtrasImageAttachmentsPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras image attachments partial update bad request response has a 5xx status code
func (o *ExtrasImageAttachmentsPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments partial update bad request response a status code equal to that given
func (o *ExtrasImageAttachmentsPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasImageAttachmentsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extrasImageAttachmentsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extrasImageAttachmentsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsPartialUpdateDefault creates a ExtrasImageAttachmentsPartialUpdateDefault with default headers values
func NewExtrasImageAttachmentsPartialUpdateDefault(code int) *ExtrasImageAttachmentsPartialUpdateDefault {
	return &ExtrasImageAttachmentsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasImageAttachmentsPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasImageAttachmentsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras image attachments partial update default response
func (o *ExtrasImageAttachmentsPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras image attachments partial update default response has a 2xx status code
func (o *ExtrasImageAttachmentsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras image attachments partial update default response has a 3xx status code
func (o *ExtrasImageAttachmentsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras image attachments partial update default response has a 4xx status code
func (o *ExtrasImageAttachmentsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras image attachments partial update default response has a 5xx status code
func (o *ExtrasImageAttachmentsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras image attachments partial update default response a status code equal to that given
func (o *ExtrasImageAttachmentsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extras_image-attachments_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/{id}/][%d] extras_image-attachments_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
