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

// DcimRackRolesBulkUpdateReader is a Reader for the DcimRackRolesBulkUpdate structure.
type DcimRackRolesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackRolesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimRackRolesBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimRackRolesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackRolesBulkUpdateOK creates a DcimRackRolesBulkUpdateOK with default headers values
func NewDcimRackRolesBulkUpdateOK() *DcimRackRolesBulkUpdateOK {
	return &DcimRackRolesBulkUpdateOK{}
}

/*
DcimRackRolesBulkUpdateOK describes a response with status code 200, with default header values.

DcimRackRolesBulkUpdateOK dcim rack roles bulk update o k
*/
type DcimRackRolesBulkUpdateOK struct {
	Payload *models.RackRole
}

// IsSuccess returns true when this dcim rack roles bulk update o k response has a 2xx status code
func (o *DcimRackRolesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack roles bulk update o k response has a 3xx status code
func (o *DcimRackRolesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles bulk update o k response has a 4xx status code
func (o *DcimRackRolesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack roles bulk update o k response has a 5xx status code
func (o *DcimRackRolesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles bulk update o k response a status code equal to that given
func (o *DcimRackRolesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRackRolesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/][%d] dcimRackRolesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/][%d] dcimRackRolesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesBulkUpdateOK) GetPayload() *models.RackRole {
	return o.Payload
}

func (o *DcimRackRolesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesBulkUpdateBadRequest creates a DcimRackRolesBulkUpdateBadRequest with default headers values
func NewDcimRackRolesBulkUpdateBadRequest() *DcimRackRolesBulkUpdateBadRequest {
	return &DcimRackRolesBulkUpdateBadRequest{}
}

/*
DcimRackRolesBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimRackRolesBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim rack roles bulk update bad request response has a 2xx status code
func (o *DcimRackRolesBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim rack roles bulk update bad request response has a 3xx status code
func (o *DcimRackRolesBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack roles bulk update bad request response has a 4xx status code
func (o *DcimRackRolesBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim rack roles bulk update bad request response has a 5xx status code
func (o *DcimRackRolesBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack roles bulk update bad request response a status code equal to that given
func (o *DcimRackRolesBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimRackRolesBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/][%d] dcimRackRolesBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRackRolesBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/][%d] dcimRackRolesBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRackRolesBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackRolesBulkUpdateDefault creates a DcimRackRolesBulkUpdateDefault with default headers values
func NewDcimRackRolesBulkUpdateDefault(code int) *DcimRackRolesBulkUpdateDefault {
	return &DcimRackRolesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRackRolesBulkUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimRackRolesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rack roles bulk update default response
func (o *DcimRackRolesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim rack roles bulk update default response has a 2xx status code
func (o *DcimRackRolesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rack roles bulk update default response has a 3xx status code
func (o *DcimRackRolesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rack roles bulk update default response has a 4xx status code
func (o *DcimRackRolesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rack roles bulk update default response has a 5xx status code
func (o *DcimRackRolesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rack roles bulk update default response a status code equal to that given
func (o *DcimRackRolesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRackRolesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/][%d] dcim_rack-roles_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/][%d] dcim_rack-roles_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackRolesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackRolesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
