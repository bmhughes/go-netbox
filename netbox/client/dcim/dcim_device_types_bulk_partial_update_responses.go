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

// DcimDeviceTypesBulkPartialUpdateReader is a Reader for the DcimDeviceTypesBulkPartialUpdate structure.
type DcimDeviceTypesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceTypesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimDeviceTypesBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimDeviceTypesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceTypesBulkPartialUpdateOK creates a DcimDeviceTypesBulkPartialUpdateOK with default headers values
func NewDcimDeviceTypesBulkPartialUpdateOK() *DcimDeviceTypesBulkPartialUpdateOK {
	return &DcimDeviceTypesBulkPartialUpdateOK{}
}

/*
DcimDeviceTypesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceTypesBulkPartialUpdateOK dcim device types bulk partial update o k
*/
type DcimDeviceTypesBulkPartialUpdateOK struct {
	Payload *models.DeviceType
}

// IsSuccess returns true when this dcim device types bulk partial update o k response has a 2xx status code
func (o *DcimDeviceTypesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device types bulk partial update o k response has a 3xx status code
func (o *DcimDeviceTypesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device types bulk partial update o k response has a 4xx status code
func (o *DcimDeviceTypesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device types bulk partial update o k response has a 5xx status code
func (o *DcimDeviceTypesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device types bulk partial update o k response a status code equal to that given
func (o *DcimDeviceTypesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcimDeviceTypesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcimDeviceTypesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) GetPayload() *models.DeviceType {
	return o.Payload
}

func (o *DcimDeviceTypesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceTypesBulkPartialUpdateBadRequest creates a DcimDeviceTypesBulkPartialUpdateBadRequest with default headers values
func NewDcimDeviceTypesBulkPartialUpdateBadRequest() *DcimDeviceTypesBulkPartialUpdateBadRequest {
	return &DcimDeviceTypesBulkPartialUpdateBadRequest{}
}

/*
DcimDeviceTypesBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimDeviceTypesBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim device types bulk partial update bad request response has a 2xx status code
func (o *DcimDeviceTypesBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim device types bulk partial update bad request response has a 3xx status code
func (o *DcimDeviceTypesBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device types bulk partial update bad request response has a 4xx status code
func (o *DcimDeviceTypesBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim device types bulk partial update bad request response has a 5xx status code
func (o *DcimDeviceTypesBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device types bulk partial update bad request response a status code equal to that given
func (o *DcimDeviceTypesBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimDeviceTypesBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcimDeviceTypesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDeviceTypesBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcimDeviceTypesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDeviceTypesBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceTypesBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceTypesBulkPartialUpdateDefault creates a DcimDeviceTypesBulkPartialUpdateDefault with default headers values
func NewDcimDeviceTypesBulkPartialUpdateDefault(code int) *DcimDeviceTypesBulkPartialUpdateDefault {
	return &DcimDeviceTypesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceTypesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimDeviceTypesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device types bulk partial update default response
func (o *DcimDeviceTypesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device types bulk partial update default response has a 2xx status code
func (o *DcimDeviceTypesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device types bulk partial update default response has a 3xx status code
func (o *DcimDeviceTypesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device types bulk partial update default response has a 4xx status code
func (o *DcimDeviceTypesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device types bulk partial update default response has a 5xx status code
func (o *DcimDeviceTypesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device types bulk partial update default response a status code equal to that given
func (o *DcimDeviceTypesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceTypesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcim_device-types_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceTypesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-types/][%d] dcim_device-types_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceTypesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceTypesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
