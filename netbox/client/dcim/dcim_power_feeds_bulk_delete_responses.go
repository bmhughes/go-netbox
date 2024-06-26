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

// DcimPowerFeedsBulkDeleteReader is a Reader for the DcimPowerFeedsBulkDelete structure.
type DcimPowerFeedsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimPowerFeedsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimPowerFeedsBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimPowerFeedsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsBulkDeleteNoContent creates a DcimPowerFeedsBulkDeleteNoContent with default headers values
func NewDcimPowerFeedsBulkDeleteNoContent() *DcimPowerFeedsBulkDeleteNoContent {
	return &DcimPowerFeedsBulkDeleteNoContent{}
}

/*
DcimPowerFeedsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimPowerFeedsBulkDeleteNoContent dcim power feeds bulk delete no content
*/
type DcimPowerFeedsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim power feeds bulk delete no content response has a 2xx status code
func (o *DcimPowerFeedsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power feeds bulk delete no content response has a 3xx status code
func (o *DcimPowerFeedsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power feeds bulk delete no content response has a 4xx status code
func (o *DcimPowerFeedsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power feeds bulk delete no content response has a 5xx status code
func (o *DcimPowerFeedsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power feeds bulk delete no content response a status code equal to that given
func (o *DcimPowerFeedsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimPowerFeedsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-feeds/][%d] dcimPowerFeedsBulkDeleteNoContent ", 204)
}

func (o *DcimPowerFeedsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-feeds/][%d] dcimPowerFeedsBulkDeleteNoContent ", 204)
}

func (o *DcimPowerFeedsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimPowerFeedsBulkDeleteBadRequest creates a DcimPowerFeedsBulkDeleteBadRequest with default headers values
func NewDcimPowerFeedsBulkDeleteBadRequest() *DcimPowerFeedsBulkDeleteBadRequest {
	return &DcimPowerFeedsBulkDeleteBadRequest{}
}

/*
DcimPowerFeedsBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimPowerFeedsBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim power feeds bulk delete bad request response has a 2xx status code
func (o *DcimPowerFeedsBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim power feeds bulk delete bad request response has a 3xx status code
func (o *DcimPowerFeedsBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power feeds bulk delete bad request response has a 4xx status code
func (o *DcimPowerFeedsBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim power feeds bulk delete bad request response has a 5xx status code
func (o *DcimPowerFeedsBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power feeds bulk delete bad request response a status code equal to that given
func (o *DcimPowerFeedsBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimPowerFeedsBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-feeds/][%d] dcimPowerFeedsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerFeedsBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-feeds/][%d] dcimPowerFeedsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerFeedsBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsBulkDeleteDefault creates a DcimPowerFeedsBulkDeleteDefault with default headers values
func NewDcimPowerFeedsBulkDeleteDefault(code int) *DcimPowerFeedsBulkDeleteDefault {
	return &DcimPowerFeedsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimPowerFeedsBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimPowerFeedsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds bulk delete default response
func (o *DcimPowerFeedsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power feeds bulk delete default response has a 2xx status code
func (o *DcimPowerFeedsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power feeds bulk delete default response has a 3xx status code
func (o *DcimPowerFeedsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power feeds bulk delete default response has a 4xx status code
func (o *DcimPowerFeedsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power feeds bulk delete default response has a 5xx status code
func (o *DcimPowerFeedsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power feeds bulk delete default response a status code equal to that given
func (o *DcimPowerFeedsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerFeedsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/power-feeds/][%d] dcim_power-feeds_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/power-feeds/][%d] dcim_power-feeds_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
