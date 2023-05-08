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

// DcimInventoryItemRolesUpdateReader is a Reader for the DcimInventoryItemRolesUpdate structure.
type DcimInventoryItemRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimInventoryItemRolesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimInventoryItemRolesUpdateOK creates a DcimInventoryItemRolesUpdateOK with default headers values
func NewDcimInventoryItemRolesUpdateOK() *DcimInventoryItemRolesUpdateOK {
	return &DcimInventoryItemRolesUpdateOK{}
}

/*
DcimInventoryItemRolesUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemRolesUpdateOK dcim inventory item roles update o k
*/
type DcimInventoryItemRolesUpdateOK struct {
	Payload *models.InventoryItemRole
}

// IsSuccess returns true when this dcim inventory item roles update o k response has a 2xx status code
func (o *DcimInventoryItemRolesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim inventory item roles update o k response has a 3xx status code
func (o *DcimInventoryItemRolesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item roles update o k response has a 4xx status code
func (o *DcimInventoryItemRolesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim inventory item roles update o k response has a 5xx status code
func (o *DcimInventoryItemRolesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item roles update o k response a status code equal to that given
func (o *DcimInventoryItemRolesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInventoryItemRolesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-roles/{id}/][%d] dcimInventoryItemRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemRolesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-roles/{id}/][%d] dcimInventoryItemRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInventoryItemRolesUpdateOK) GetPayload() *models.InventoryItemRole {
	return o.Payload
}

func (o *DcimInventoryItemRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemRolesUpdateBadRequest creates a DcimInventoryItemRolesUpdateBadRequest with default headers values
func NewDcimInventoryItemRolesUpdateBadRequest() *DcimInventoryItemRolesUpdateBadRequest {
	return &DcimInventoryItemRolesUpdateBadRequest{}
}

/*
DcimInventoryItemRolesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimInventoryItemRolesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim inventory item roles update bad request response has a 2xx status code
func (o *DcimInventoryItemRolesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim inventory item roles update bad request response has a 3xx status code
func (o *DcimInventoryItemRolesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim inventory item roles update bad request response has a 4xx status code
func (o *DcimInventoryItemRolesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim inventory item roles update bad request response has a 5xx status code
func (o *DcimInventoryItemRolesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim inventory item roles update bad request response a status code equal to that given
func (o *DcimInventoryItemRolesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimInventoryItemRolesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-roles/{id}/][%d] dcimInventoryItemRolesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInventoryItemRolesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/inventory-item-roles/{id}/][%d] dcimInventoryItemRolesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInventoryItemRolesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemRolesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
