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

// DcimDeviceBayTemplatesBulkPartialUpdateReader is a Reader for the DcimDeviceBayTemplatesBulkPartialUpdate structure.
type DcimDeviceBayTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimDeviceBayTemplatesBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimDeviceBayTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimDeviceBayTemplatesBulkPartialUpdateOK creates a DcimDeviceBayTemplatesBulkPartialUpdateOK with default headers values
func NewDcimDeviceBayTemplatesBulkPartialUpdateOK() *DcimDeviceBayTemplatesBulkPartialUpdateOK {
	return &DcimDeviceBayTemplatesBulkPartialUpdateOK{}
}

/*
DcimDeviceBayTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesBulkPartialUpdateOK dcim device bay templates bulk partial update o k
*/
type DcimDeviceBayTemplatesBulkPartialUpdateOK struct {
	Payload *models.DeviceBayTemplate
}

// IsSuccess returns true when this dcim device bay templates bulk partial update o k response has a 2xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bay templates bulk partial update o k response has a 3xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates bulk partial update o k response has a 4xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bay templates bulk partial update o k response has a 5xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates bulk partial update o k response a status code equal to that given
func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBayTemplatesBulkPartialUpdateBadRequest creates a DcimDeviceBayTemplatesBulkPartialUpdateBadRequest with default headers values
func NewDcimDeviceBayTemplatesBulkPartialUpdateBadRequest() *DcimDeviceBayTemplatesBulkPartialUpdateBadRequest {
	return &DcimDeviceBayTemplatesBulkPartialUpdateBadRequest{}
}

/*
DcimDeviceBayTemplatesBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimDeviceBayTemplatesBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim device bay templates bulk partial update bad request response has a 2xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim device bay templates bulk partial update bad request response has a 3xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates bulk partial update bad request response has a 4xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim device bay templates bulk partial update bad request response has a 5xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates bulk partial update bad request response a status code equal to that given
func (o *DcimDeviceBayTemplatesBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/][%d] dcimDeviceBayTemplatesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBayTemplatesBulkPartialUpdateDefault creates a DcimDeviceBayTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimDeviceBayTemplatesBulkPartialUpdateDefault(code int) *DcimDeviceBayTemplatesBulkPartialUpdateDefault {
	return &DcimDeviceBayTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimDeviceBayTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimDeviceBayTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim device bay templates bulk partial update default response
func (o *DcimDeviceBayTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim device bay templates bulk partial update default response has a 2xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim device bay templates bulk partial update default response has a 3xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim device bay templates bulk partial update default response has a 4xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim device bay templates bulk partial update default response has a 5xx status code
func (o *DcimDeviceBayTemplatesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim device bay templates bulk partial update default response a status code equal to that given
func (o *DcimDeviceBayTemplatesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/][%d] dcim_device-bay-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/device-bay-templates/][%d] dcim_device-bay-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
