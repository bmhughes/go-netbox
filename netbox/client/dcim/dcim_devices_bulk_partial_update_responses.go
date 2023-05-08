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

// DcimDevicesBulkPartialUpdateReader is a Reader for the DcimDevicesBulkPartialUpdate structure.
type DcimDevicesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDevicesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDevicesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimDevicesBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDevicesBulkPartialUpdateOK creates a DcimDevicesBulkPartialUpdateOK with default headers values
func NewDcimDevicesBulkPartialUpdateOK() *DcimDevicesBulkPartialUpdateOK {
	return &DcimDevicesBulkPartialUpdateOK{}
}

/*
DcimDevicesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDevicesBulkPartialUpdateOK dcim devices bulk partial update o k
*/
type DcimDevicesBulkPartialUpdateOK struct {
	Payload *models.DeviceWithConfigContext
}

// IsSuccess returns true when this dcim devices bulk partial update o k response has a 2xx status code
func (o *DcimDevicesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim devices bulk partial update o k response has a 3xx status code
func (o *DcimDevicesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices bulk partial update o k response has a 4xx status code
func (o *DcimDevicesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim devices bulk partial update o k response has a 5xx status code
func (o *DcimDevicesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices bulk partial update o k response a status code equal to that given
func (o *DcimDevicesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDevicesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/devices/][%d] dcimDevicesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/devices/][%d] dcimDevicesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDevicesBulkPartialUpdateOK) GetPayload() *models.DeviceWithConfigContext {
	return o.Payload
}

func (o *DcimDevicesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDevicesBulkPartialUpdateBadRequest creates a DcimDevicesBulkPartialUpdateBadRequest with default headers values
func NewDcimDevicesBulkPartialUpdateBadRequest() *DcimDevicesBulkPartialUpdateBadRequest {
	return &DcimDevicesBulkPartialUpdateBadRequest{}
}

/*
DcimDevicesBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimDevicesBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim devices bulk partial update bad request response has a 2xx status code
func (o *DcimDevicesBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim devices bulk partial update bad request response has a 3xx status code
func (o *DcimDevicesBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim devices bulk partial update bad request response has a 4xx status code
func (o *DcimDevicesBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim devices bulk partial update bad request response has a 5xx status code
func (o *DcimDevicesBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim devices bulk partial update bad request response a status code equal to that given
func (o *DcimDevicesBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimDevicesBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/devices/][%d] dcimDevicesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDevicesBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/devices/][%d] dcimDevicesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDevicesBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDevicesBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
