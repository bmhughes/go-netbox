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

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// DcimInventoryItemTemplatesBulkPartialUpdateReader is a Reader for the DcimInventoryItemTemplatesBulkPartialUpdate structure.
type DcimInventoryItemTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimInventoryItemTemplatesBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimInventoryItemTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemTemplatesBulkPartialUpdateOK creates a DcimInventoryItemTemplatesBulkPartialUpdateOK with default headers values
func NewDcimInventoryItemTemplatesBulkPartialUpdateOK() *DcimInventoryItemTemplatesBulkPartialUpdateOK {
	return &DcimInventoryItemTemplatesBulkPartialUpdateOK{}
}

/*
DcimInventoryItemTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemTemplatesBulkPartialUpdateOK dcim inventory item templates bulk partial update o k
*/
type DcimInventoryItemTemplatesBulkPartialUpdateOK struct {
	Payload *models.InventoryItemTemplate
}

// IsSuccess returns true when this dcim inventory item templates bulk partial update o k response has a 2xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory item templates bulk partial update o k response has a 3xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item templates bulk partial update o k response has a 4xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory item templates bulk partial update o k response has a 5xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item templates bulk partial update o k response a status code equal to that given
func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) GetPayload() *models.InventoryItemTemplate {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemTemplatesBulkPartialUpdateBadRequest creates a DcimInventoryItemTemplatesBulkPartialUpdateBadRequest with default headers values
func NewDcimInventoryItemTemplatesBulkPartialUpdateBadRequest() *DcimInventoryItemTemplatesBulkPartialUpdateBadRequest {
	return &DcimInventoryItemTemplatesBulkPartialUpdateBadRequest{}
}

/*
DcimInventoryItemTemplatesBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimInventoryItemTemplatesBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim inventory item templates bulk partial update bad request response has a 2xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim inventory item templates bulk partial update bad request response has a 3xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item templates bulk partial update bad request response has a 4xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim inventory item templates bulk partial update bad request response has a 5xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item templates bulk partial update bad request response a status code equal to that given
func (o *DcimInventoryItemTemplatesBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemTemplatesBulkPartialUpdateDefault creates a DcimInventoryItemTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimInventoryItemTemplatesBulkPartialUpdateDefault(code int) *DcimInventoryItemTemplatesBulkPartialUpdateDefault {
	return &DcimInventoryItemTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimInventoryItemTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item templates bulk partial update default response
func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim inventory item templates bulk partial update default response has a 2xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory item templates bulk partial update default response has a 3xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory item templates bulk partial update default response has a 4xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory item templates bulk partial update default response has a 5xx status code
func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory item templates bulk partial update default response a status code equal to that given
func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/][%d] dcim_inventory-item-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-templates/][%d] dcim_inventory-item-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
