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

// DcimManufacturersBulkPartialUpdateReader is a Reader for the DcimManufacturersBulkPartialUpdate structure.
type DcimManufacturersBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimManufacturersBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimManufacturersBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimManufacturersBulkPartialUpdateOK creates a DcimManufacturersBulkPartialUpdateOK with default headers values
func NewDcimManufacturersBulkPartialUpdateOK() *DcimManufacturersBulkPartialUpdateOK {
	return &DcimManufacturersBulkPartialUpdateOK{}
}

/*
DcimManufacturersBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimManufacturersBulkPartialUpdateOK dcim manufacturers bulk partial update o k
*/
type DcimManufacturersBulkPartialUpdateOK struct {
	Payload *models.Manufacturer
}

// IsSuccess returns true when this dcim manufacturers bulk partial update o k response has a 2xx status code
func (o *DcimManufacturersBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim manufacturers bulk partial update o k response has a 3xx status code
func (o *DcimManufacturersBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim manufacturers bulk partial update o k response has a 4xx status code
func (o *DcimManufacturersBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim manufacturers bulk partial update o k response has a 5xx status code
func (o *DcimManufacturersBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim manufacturers bulk partial update o k response a status code equal to that given
func (o *DcimManufacturersBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimManufacturersBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/][%d] dcimManufacturersBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/][%d] dcimManufacturersBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersBulkPartialUpdateOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersBulkPartialUpdateBadRequest creates a DcimManufacturersBulkPartialUpdateBadRequest with default headers values
func NewDcimManufacturersBulkPartialUpdateBadRequest() *DcimManufacturersBulkPartialUpdateBadRequest {
	return &DcimManufacturersBulkPartialUpdateBadRequest{}
}

/*
DcimManufacturersBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimManufacturersBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim manufacturers bulk partial update bad request response has a 2xx status code
func (o *DcimManufacturersBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim manufacturers bulk partial update bad request response has a 3xx status code
func (o *DcimManufacturersBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim manufacturers bulk partial update bad request response has a 4xx status code
func (o *DcimManufacturersBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim manufacturers bulk partial update bad request response has a 5xx status code
func (o *DcimManufacturersBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim manufacturers bulk partial update bad request response a status code equal to that given
func (o *DcimManufacturersBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimManufacturersBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/][%d] dcimManufacturersBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimManufacturersBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/][%d] dcimManufacturersBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimManufacturersBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersBulkPartialUpdateDefault creates a DcimManufacturersBulkPartialUpdateDefault with default headers values
func NewDcimManufacturersBulkPartialUpdateDefault(code int) *DcimManufacturersBulkPartialUpdateDefault {
	return &DcimManufacturersBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimManufacturersBulkPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimManufacturersBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim manufacturers bulk partial update default response
func (o *DcimManufacturersBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim manufacturers bulk partial update default response has a 2xx status code
func (o *DcimManufacturersBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim manufacturers bulk partial update default response has a 3xx status code
func (o *DcimManufacturersBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim manufacturers bulk partial update default response has a 4xx status code
func (o *DcimManufacturersBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim manufacturers bulk partial update default response has a 5xx status code
func (o *DcimManufacturersBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim manufacturers bulk partial update default response a status code equal to that given
func (o *DcimManufacturersBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimManufacturersBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/][%d] dcim_manufacturers_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimManufacturersBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/manufacturers/][%d] dcim_manufacturers_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimManufacturersBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
