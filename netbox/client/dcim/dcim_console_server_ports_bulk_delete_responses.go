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

// DcimConsoleServerPortsBulkDeleteReader is a Reader for the DcimConsoleServerPortsBulkDelete structure.
type DcimConsoleServerPortsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsoleServerPortsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimConsoleServerPortsBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimConsoleServerPortsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortsBulkDeleteNoContent creates a DcimConsoleServerPortsBulkDeleteNoContent with default headers values
func NewDcimConsoleServerPortsBulkDeleteNoContent() *DcimConsoleServerPortsBulkDeleteNoContent {
	return &DcimConsoleServerPortsBulkDeleteNoContent{}
}

/*
DcimConsoleServerPortsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimConsoleServerPortsBulkDeleteNoContent dcim console server ports bulk delete no content
*/
type DcimConsoleServerPortsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim console server ports bulk delete no content response has a 2xx status code
func (o *DcimConsoleServerPortsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server ports bulk delete no content response has a 3xx status code
func (o *DcimConsoleServerPortsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server ports bulk delete no content response has a 4xx status code
func (o *DcimConsoleServerPortsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server ports bulk delete no content response has a 5xx status code
func (o *DcimConsoleServerPortsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server ports bulk delete no content response a status code equal to that given
func (o *DcimConsoleServerPortsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimConsoleServerPortsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-ports/][%d] dcimConsoleServerPortsBulkDeleteNoContent ", 204)
}

func (o *DcimConsoleServerPortsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-ports/][%d] dcimConsoleServerPortsBulkDeleteNoContent ", 204)
}

func (o *DcimConsoleServerPortsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimConsoleServerPortsBulkDeleteBadRequest creates a DcimConsoleServerPortsBulkDeleteBadRequest with default headers values
func NewDcimConsoleServerPortsBulkDeleteBadRequest() *DcimConsoleServerPortsBulkDeleteBadRequest {
	return &DcimConsoleServerPortsBulkDeleteBadRequest{}
}

/*
DcimConsoleServerPortsBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimConsoleServerPortsBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim console server ports bulk delete bad request response has a 2xx status code
func (o *DcimConsoleServerPortsBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim console server ports bulk delete bad request response has a 3xx status code
func (o *DcimConsoleServerPortsBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server ports bulk delete bad request response has a 4xx status code
func (o *DcimConsoleServerPortsBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim console server ports bulk delete bad request response has a 5xx status code
func (o *DcimConsoleServerPortsBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server ports bulk delete bad request response a status code equal to that given
func (o *DcimConsoleServerPortsBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimConsoleServerPortsBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-ports/][%d] dcimConsoleServerPortsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsoleServerPortsBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-ports/][%d] dcimConsoleServerPortsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsoleServerPortsBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsoleServerPortsBulkDeleteDefault creates a DcimConsoleServerPortsBulkDeleteDefault with default headers values
func NewDcimConsoleServerPortsBulkDeleteDefault(code int) *DcimConsoleServerPortsBulkDeleteDefault {
	return &DcimConsoleServerPortsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimConsoleServerPortsBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimConsoleServerPortsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server ports bulk delete default response
func (o *DcimConsoleServerPortsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console server ports bulk delete default response has a 2xx status code
func (o *DcimConsoleServerPortsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console server ports bulk delete default response has a 3xx status code
func (o *DcimConsoleServerPortsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console server ports bulk delete default response has a 4xx status code
func (o *DcimConsoleServerPortsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console server ports bulk delete default response has a 5xx status code
func (o *DcimConsoleServerPortsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console server ports bulk delete default response a status code equal to that given
func (o *DcimConsoleServerPortsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsoleServerPortsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-ports/][%d] dcim_console-server-ports_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-ports/][%d] dcim_console-server-ports_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
