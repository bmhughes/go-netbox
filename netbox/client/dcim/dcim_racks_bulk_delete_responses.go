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

// DcimRacksBulkDeleteReader is a Reader for the DcimRacksBulkDelete structure.
type DcimRacksBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRacksBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimRacksBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimRacksBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRacksBulkDeleteNoContent creates a DcimRacksBulkDeleteNoContent with default headers values
func NewDcimRacksBulkDeleteNoContent() *DcimRacksBulkDeleteNoContent {
	return &DcimRacksBulkDeleteNoContent{}
}

/*
DcimRacksBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRacksBulkDeleteNoContent dcim racks bulk delete no content
*/
type DcimRacksBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim racks bulk delete no content response has a 2xx status code
func (o *DcimRacksBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim racks bulk delete no content response has a 3xx status code
func (o *DcimRacksBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim racks bulk delete no content response has a 4xx status code
func (o *DcimRacksBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim racks bulk delete no content response has a 5xx status code
func (o *DcimRacksBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim racks bulk delete no content response a status code equal to that given
func (o *DcimRacksBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimRacksBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/racks/][%d] dcimRacksBulkDeleteNoContent ", 204)
}

func (o *DcimRacksBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/racks/][%d] dcimRacksBulkDeleteNoContent ", 204)
}

func (o *DcimRacksBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRacksBulkDeleteBadRequest creates a DcimRacksBulkDeleteBadRequest with default headers values
func NewDcimRacksBulkDeleteBadRequest() *DcimRacksBulkDeleteBadRequest {
	return &DcimRacksBulkDeleteBadRequest{}
}

/*
DcimRacksBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimRacksBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim racks bulk delete bad request response has a 2xx status code
func (o *DcimRacksBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim racks bulk delete bad request response has a 3xx status code
func (o *DcimRacksBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim racks bulk delete bad request response has a 4xx status code
func (o *DcimRacksBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim racks bulk delete bad request response has a 5xx status code
func (o *DcimRacksBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim racks bulk delete bad request response a status code equal to that given
func (o *DcimRacksBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimRacksBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/racks/][%d] dcimRacksBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRacksBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/racks/][%d] dcimRacksBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRacksBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRacksBulkDeleteDefault creates a DcimRacksBulkDeleteDefault with default headers values
func NewDcimRacksBulkDeleteDefault(code int) *DcimRacksBulkDeleteDefault {
	return &DcimRacksBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimRacksBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimRacksBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim racks bulk delete default response
func (o *DcimRacksBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim racks bulk delete default response has a 2xx status code
func (o *DcimRacksBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim racks bulk delete default response has a 3xx status code
func (o *DcimRacksBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim racks bulk delete default response has a 4xx status code
func (o *DcimRacksBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim racks bulk delete default response has a 5xx status code
func (o *DcimRacksBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim racks bulk delete default response a status code equal to that given
func (o *DcimRacksBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRacksBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/racks/][%d] dcim_racks_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/racks/][%d] dcim_racks_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRacksBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRacksBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
