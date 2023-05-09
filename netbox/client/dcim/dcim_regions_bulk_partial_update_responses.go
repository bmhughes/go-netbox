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

// DcimRegionsBulkPartialUpdateReader is a Reader for the DcimRegionsBulkPartialUpdate structure.
type DcimRegionsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimRegionsBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimRegionsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRegionsBulkPartialUpdateOK creates a DcimRegionsBulkPartialUpdateOK with default headers values
func NewDcimRegionsBulkPartialUpdateOK() *DcimRegionsBulkPartialUpdateOK {
	return &DcimRegionsBulkPartialUpdateOK{}
}

/*
DcimRegionsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRegionsBulkPartialUpdateOK dcim regions bulk partial update o k
*/
type DcimRegionsBulkPartialUpdateOK struct {
	Payload *models.Region
}

// IsSuccess returns true when this dcim regions bulk partial update o k response has a 2xx status code
func (o *DcimRegionsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim regions bulk partial update o k response has a 3xx status code
func (o *DcimRegionsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions bulk partial update o k response has a 4xx status code
func (o *DcimRegionsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim regions bulk partial update o k response has a 5xx status code
func (o *DcimRegionsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions bulk partial update o k response a status code equal to that given
func (o *DcimRegionsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRegionsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcimRegionsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcimRegionsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsBulkPartialUpdateOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRegionsBulkPartialUpdateBadRequest creates a DcimRegionsBulkPartialUpdateBadRequest with default headers values
func NewDcimRegionsBulkPartialUpdateBadRequest() *DcimRegionsBulkPartialUpdateBadRequest {
	return &DcimRegionsBulkPartialUpdateBadRequest{}
}

/*
DcimRegionsBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimRegionsBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim regions bulk partial update bad request response has a 2xx status code
func (o *DcimRegionsBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim regions bulk partial update bad request response has a 3xx status code
func (o *DcimRegionsBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions bulk partial update bad request response has a 4xx status code
func (o *DcimRegionsBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim regions bulk partial update bad request response has a 5xx status code
func (o *DcimRegionsBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions bulk partial update bad request response a status code equal to that given
func (o *DcimRegionsBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimRegionsBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcimRegionsBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRegionsBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcimRegionsBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRegionsBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRegionsBulkPartialUpdateDefault creates a DcimRegionsBulkPartialUpdateDefault with default headers values
func NewDcimRegionsBulkPartialUpdateDefault(code int) *DcimRegionsBulkPartialUpdateDefault {
	return &DcimRegionsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRegionsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimRegionsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim regions bulk partial update default response
func (o *DcimRegionsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim regions bulk partial update default response has a 2xx status code
func (o *DcimRegionsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim regions bulk partial update default response has a 3xx status code
func (o *DcimRegionsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim regions bulk partial update default response has a 4xx status code
func (o *DcimRegionsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim regions bulk partial update default response has a 5xx status code
func (o *DcimRegionsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim regions bulk partial update default response a status code equal to that given
func (o *DcimRegionsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRegionsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcim_regions_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/regions/][%d] dcim_regions_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
