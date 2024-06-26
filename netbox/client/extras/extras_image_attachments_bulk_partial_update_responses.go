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

// ExtrasImageAttachmentsBulkPartialUpdateReader is a Reader for the ExtrasImageAttachmentsBulkPartialUpdate structure.
type ExtrasImageAttachmentsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasImageAttachmentsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasImageAttachmentsBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasImageAttachmentsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsBulkPartialUpdateOK creates a ExtrasImageAttachmentsBulkPartialUpdateOK with default headers values
func NewExtrasImageAttachmentsBulkPartialUpdateOK() *ExtrasImageAttachmentsBulkPartialUpdateOK {
	return &ExtrasImageAttachmentsBulkPartialUpdateOK{}
}

/*
ExtrasImageAttachmentsBulkPartialUpdateOK describes a response with status code 200, with default header values.

ExtrasImageAttachmentsBulkPartialUpdateOK extras image attachments bulk partial update o k
*/
type ExtrasImageAttachmentsBulkPartialUpdateOK struct {
	Payload *models.ImageAttachment
}

// IsSuccess returns true when this extras image attachments bulk partial update o k response has a 2xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments bulk partial update o k response has a 3xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments bulk partial update o k response has a 4xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments bulk partial update o k response has a 5xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments bulk partial update o k response a status code equal to that given
func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/][%d] extrasImageAttachmentsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/][%d] extrasImageAttachmentsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) GetPayload() *models.ImageAttachment {
	return o.Payload
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsBulkPartialUpdateBadRequest creates a ExtrasImageAttachmentsBulkPartialUpdateBadRequest with default headers values
func NewExtrasImageAttachmentsBulkPartialUpdateBadRequest() *ExtrasImageAttachmentsBulkPartialUpdateBadRequest {
	return &ExtrasImageAttachmentsBulkPartialUpdateBadRequest{}
}

/*
ExtrasImageAttachmentsBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasImageAttachmentsBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras image attachments bulk partial update bad request response has a 2xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras image attachments bulk partial update bad request response has a 3xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments bulk partial update bad request response has a 4xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras image attachments bulk partial update bad request response has a 5xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments bulk partial update bad request response a status code equal to that given
func (o *ExtrasImageAttachmentsBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/][%d] extrasImageAttachmentsBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/][%d] extrasImageAttachmentsBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasImageAttachmentsBulkPartialUpdateDefault creates a ExtrasImageAttachmentsBulkPartialUpdateDefault with default headers values
func NewExtrasImageAttachmentsBulkPartialUpdateDefault(code int) *ExtrasImageAttachmentsBulkPartialUpdateDefault {
	return &ExtrasImageAttachmentsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasImageAttachmentsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasImageAttachmentsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras image attachments bulk partial update default response
func (o *ExtrasImageAttachmentsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras image attachments bulk partial update default response has a 2xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras image attachments bulk partial update default response has a 3xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras image attachments bulk partial update default response has a 4xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras image attachments bulk partial update default response has a 5xx status code
func (o *ExtrasImageAttachmentsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras image attachments bulk partial update default response a status code equal to that given
func (o *ExtrasImageAttachmentsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/][%d] extras_image-attachments_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /extras/image-attachments/][%d] extras_image-attachments_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
