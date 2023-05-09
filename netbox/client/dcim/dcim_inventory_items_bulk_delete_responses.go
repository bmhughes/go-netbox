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

// DcimInventoryItemsBulkDeleteReader is a Reader for the DcimInventoryItemsBulkDelete structure.
type DcimInventoryItemsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimInventoryItemsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimInventoryItemsBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimInventoryItemsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemsBulkDeleteNoContent creates a DcimInventoryItemsBulkDeleteNoContent with default headers values
func NewDcimInventoryItemsBulkDeleteNoContent() *DcimInventoryItemsBulkDeleteNoContent {
	return &DcimInventoryItemsBulkDeleteNoContent{}
}

/*
DcimInventoryItemsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimInventoryItemsBulkDeleteNoContent dcim inventory items bulk delete no content
*/
type DcimInventoryItemsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim inventory items bulk delete no content response has a 2xx status code
func (o *DcimInventoryItemsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory items bulk delete no content response has a 3xx status code
func (o *DcimInventoryItemsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory items bulk delete no content response has a 4xx status code
func (o *DcimInventoryItemsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory items bulk delete no content response has a 5xx status code
func (o *DcimInventoryItemsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory items bulk delete no content response a status code equal to that given
func (o *DcimInventoryItemsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimInventoryItemsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/][%d] dcimInventoryItemsBulkDeleteNoContent ", 204)
}

func (o *DcimInventoryItemsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/][%d] dcimInventoryItemsBulkDeleteNoContent ", 204)
}

func (o *DcimInventoryItemsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimInventoryItemsBulkDeleteBadRequest creates a DcimInventoryItemsBulkDeleteBadRequest with default headers values
func NewDcimInventoryItemsBulkDeleteBadRequest() *DcimInventoryItemsBulkDeleteBadRequest {
	return &DcimInventoryItemsBulkDeleteBadRequest{}
}

/*
DcimInventoryItemsBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimInventoryItemsBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim inventory items bulk delete bad request response has a 2xx status code
func (o *DcimInventoryItemsBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim inventory items bulk delete bad request response has a 3xx status code
func (o *DcimInventoryItemsBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory items bulk delete bad request response has a 4xx status code
func (o *DcimInventoryItemsBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim inventory items bulk delete bad request response has a 5xx status code
func (o *DcimInventoryItemsBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory items bulk delete bad request response a status code equal to that given
func (o *DcimInventoryItemsBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimInventoryItemsBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/][%d] dcimInventoryItemsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInventoryItemsBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/][%d] dcimInventoryItemsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInventoryItemsBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemsBulkDeleteDefault creates a DcimInventoryItemsBulkDeleteDefault with default headers values
func NewDcimInventoryItemsBulkDeleteDefault(code int) *DcimInventoryItemsBulkDeleteDefault {
	return &DcimInventoryItemsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemsBulkDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimInventoryItemsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory items bulk delete default response
func (o *DcimInventoryItemsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim inventory items bulk delete default response has a 2xx status code
func (o *DcimInventoryItemsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory items bulk delete default response has a 3xx status code
func (o *DcimInventoryItemsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory items bulk delete default response has a 4xx status code
func (o *DcimInventoryItemsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory items bulk delete default response has a 5xx status code
func (o *DcimInventoryItemsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory items bulk delete default response a status code equal to that given
func (o *DcimInventoryItemsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInventoryItemsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/][%d] dcim_inventory-items_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/inventory-items/][%d] dcim_inventory-items_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
