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

// DcimInventoryItemRolesBulkPartialUpdateReader is a Reader for the DcimInventoryItemRolesBulkPartialUpdate structure.
type DcimInventoryItemRolesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemRolesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemRolesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimInventoryItemRolesBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimInventoryItemRolesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemRolesBulkPartialUpdateOK creates a DcimInventoryItemRolesBulkPartialUpdateOK with default headers values
func NewDcimInventoryItemRolesBulkPartialUpdateOK() *DcimInventoryItemRolesBulkPartialUpdateOK {
	return &DcimInventoryItemRolesBulkPartialUpdateOK{}
}

/*
DcimInventoryItemRolesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemRolesBulkPartialUpdateOK dcim inventory item roles bulk partial update o k
*/
type DcimInventoryItemRolesBulkPartialUpdateOK struct {
	Payload *models.InventoryItemRole
}

// IsSuccess returns true when this dcim inventory item roles bulk partial update o k response has a 2xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory item roles bulk partial update o k response has a 3xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item roles bulk partial update o k response has a 4xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory item roles bulk partial update o k response has a 5xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item roles bulk partial update o k response a status code equal to that given
func (o *DcimInventoryItemRolesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInventoryItemRolesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-roles/][%d] dcimInventoryItemRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemRolesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-roles/][%d] dcimInventoryItemRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemRolesBulkPartialUpdateOK) GetPayload() *models.InventoryItemRole {
	return o.Payload
}

func (o *DcimInventoryItemRolesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemRolesBulkPartialUpdateBadRequest creates a DcimInventoryItemRolesBulkPartialUpdateBadRequest with default headers values
func NewDcimInventoryItemRolesBulkPartialUpdateBadRequest() *DcimInventoryItemRolesBulkPartialUpdateBadRequest {
	return &DcimInventoryItemRolesBulkPartialUpdateBadRequest{}
}

/*
DcimInventoryItemRolesBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimInventoryItemRolesBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim inventory item roles bulk partial update bad request response has a 2xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim inventory item roles bulk partial update bad request response has a 3xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item roles bulk partial update bad request response has a 4xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim inventory item roles bulk partial update bad request response has a 5xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item roles bulk partial update bad request response a status code equal to that given
func (o *DcimInventoryItemRolesBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimInventoryItemRolesBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-roles/][%d] dcimInventoryItemRolesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInventoryItemRolesBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-roles/][%d] dcimInventoryItemRolesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInventoryItemRolesBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemRolesBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemRolesBulkPartialUpdateDefault creates a DcimInventoryItemRolesBulkPartialUpdateDefault with default headers values
func NewDcimInventoryItemRolesBulkPartialUpdateDefault(code int) *DcimInventoryItemRolesBulkPartialUpdateDefault {
	return &DcimInventoryItemRolesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInventoryItemRolesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimInventoryItemRolesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item roles bulk partial update default response
func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim inventory item roles bulk partial update default response has a 2xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim inventory item roles bulk partial update default response has a 3xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim inventory item roles bulk partial update default response has a 4xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim inventory item roles bulk partial update default response has a 5xx status code
func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim inventory item roles bulk partial update default response a status code equal to that given
func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-roles/][%d] dcim_inventory-item-roles_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-roles/][%d] dcim_inventory-item-roles_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
