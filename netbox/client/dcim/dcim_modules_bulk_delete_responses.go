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

// DcimModulesBulkDeleteReader is a Reader for the DcimModulesBulkDelete structure.
type DcimModulesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModulesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimModulesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimModulesBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimModulesBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModulesBulkDeleteNoContent creates a DcimModulesBulkDeleteNoContent with default headers values
func NewDcimModulesBulkDeleteNoContent() *DcimModulesBulkDeleteNoContent {
	return &DcimModulesBulkDeleteNoContent{}
}

/*
DcimModulesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimModulesBulkDeleteNoContent dcim modules bulk delete no content
*/
type DcimModulesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim modules bulk delete no content response has a 2xx status code
func (o *DcimModulesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim modules bulk delete no content response has a 3xx status code
func (o *DcimModulesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim modules bulk delete no content response has a 4xx status code
func (o *DcimModulesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim modules bulk delete no content response has a 5xx status code
func (o *DcimModulesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim modules bulk delete no content response a status code equal to that given
func (o *DcimModulesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimModulesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/modules/][%d] dcimModulesBulkDeleteNoContent ", 204)
}

func (o *DcimModulesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/modules/][%d] dcimModulesBulkDeleteNoContent ", 204)
}

func (o *DcimModulesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimModulesBulkDeleteBadRequest creates a DcimModulesBulkDeleteBadRequest with default headers values
func NewDcimModulesBulkDeleteBadRequest() *DcimModulesBulkDeleteBadRequest {
	return &DcimModulesBulkDeleteBadRequest{}
}

/*
DcimModulesBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimModulesBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim modules bulk delete bad request response has a 2xx status code
func (o *DcimModulesBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim modules bulk delete bad request response has a 3xx status code
func (o *DcimModulesBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim modules bulk delete bad request response has a 4xx status code
func (o *DcimModulesBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim modules bulk delete bad request response has a 5xx status code
func (o *DcimModulesBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim modules bulk delete bad request response a status code equal to that given
func (o *DcimModulesBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimModulesBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/modules/][%d] dcimModulesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimModulesBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/modules/][%d] dcimModulesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimModulesBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModulesBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModulesBulkDeleteDefault creates a DcimModulesBulkDeleteDefault with default headers values
func NewDcimModulesBulkDeleteDefault(code int) *DcimModulesBulkDeleteDefault {
	return &DcimModulesBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimModulesBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimModulesBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim modules bulk delete default response
func (o *DcimModulesBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim modules bulk delete default response has a 2xx status code
func (o *DcimModulesBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim modules bulk delete default response has a 3xx status code
func (o *DcimModulesBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim modules bulk delete default response has a 4xx status code
func (o *DcimModulesBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim modules bulk delete default response has a 5xx status code
func (o *DcimModulesBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim modules bulk delete default response a status code equal to that given
func (o *DcimModulesBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModulesBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/modules/][%d] dcim_modules_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModulesBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/modules/][%d] dcim_modules_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModulesBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModulesBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
