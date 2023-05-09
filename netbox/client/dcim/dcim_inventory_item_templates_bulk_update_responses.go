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

// DcimInventoryItemTemplatesBulkUpdateReader is a Reader for the DcimInventoryItemTemplatesBulkUpdate structure.
type DcimInventoryItemTemplatesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemTemplatesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemTemplatesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimInventoryItemTemplatesBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimInventoryItemTemplatesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemTemplatesBulkUpdateOK creates a DcimInventoryItemTemplatesBulkUpdateOK with default headers values
func NewDcimInventoryItemTemplatesBulkUpdateOK() *DcimInventoryItemTemplatesBulkUpdateOK {
	return &DcimInventoryItemTemplatesBulkUpdateOK{}
}

/*
DcimInventoryItemTemplatesBulkUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemTemplatesBulkUpdateOK dcim inventory item templates bulk update o k
*/
type DcimInventoryItemTemplatesBulkUpdateOK struct {
	Payload *models.InventoryItemTemplate
}

// IsSuccess returns true when this dcim inventory item templates bulk update o k response has a 2xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory item templates bulk update o k response has a 3xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item templates bulk update o k response has a 4xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory item templates bulk update o k response has a 5xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item templates bulk update o k response a status code equal to that given
func (o *DcimInventoryItemTemplatesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) GetPayload() *models.InventoryItemTemplate {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemTemplatesBulkUpdateBadRequest creates a DcimInventoryItemTemplatesBulkUpdateBadRequest with default headers values
func NewDcimInventoryItemTemplatesBulkUpdateBadRequest() *DcimInventoryItemTemplatesBulkUpdateBadRequest {
	return &DcimInventoryItemTemplatesBulkUpdateBadRequest{}
}

/*
DcimInventoryItemTemplatesBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimInventoryItemTemplatesBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim inventory item templates bulk update bad request response has a 2xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim inventory item templates bulk update bad request response has a 3xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item templates bulk update bad request response has a 4xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim inventory item templates bulk update bad request response has a 5xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item templates bulk update bad request response a status code equal to that given
func (o *DcimInventoryItemTemplatesBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimInventoryItemTemplatesBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcimInventoryItemTemplatesBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemTemplatesBulkUpdateDefault creates a DcimInventoryItemTemplatesBulkUpdateDefault with default headers values
func NewDcimInventoryItemTemplatesBulkUpdateDefault(code int) *DcimInventoryItemTemplatesBulkUpdateDefault {
	return &DcimInventoryItemTemplatesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemTemplatesBulkUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimInventoryItemTemplatesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item templates bulk update default response
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim inventory item templates bulk update default response has a 2xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory item templates bulk update default response has a 3xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory item templates bulk update default response has a 4xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory item templates bulk update default response has a 5xx status code
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory item templates bulk update default response a status code equal to that given
func (o *DcimInventoryItemTemplatesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInventoryItemTemplatesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcim_inventory-item-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-templates/][%d] dcim_inventory-item-templates_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemTemplatesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemTemplatesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
