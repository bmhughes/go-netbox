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

// DcimVirtualChassisPartialUpdateReader is a Reader for the DcimVirtualChassisPartialUpdate structure.
type DcimVirtualChassisPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimVirtualChassisPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimVirtualChassisPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualChassisPartialUpdateOK creates a DcimVirtualChassisPartialUpdateOK with default headers values
func NewDcimVirtualChassisPartialUpdateOK() *DcimVirtualChassisPartialUpdateOK {
	return &DcimVirtualChassisPartialUpdateOK{}
}

/*
DcimVirtualChassisPartialUpdateOK describes a response with status code 200, with default header values.

DcimVirtualChassisPartialUpdateOK dcim virtual chassis partial update o k
*/
type DcimVirtualChassisPartialUpdateOK struct {
	Payload *models.VirtualChassis
}

// IsSuccess returns true when this dcim virtual chassis partial update o k response has a 2xx status code
func (o *DcimVirtualChassisPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual chassis partial update o k response has a 3xx status code
func (o *DcimVirtualChassisPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual chassis partial update o k response has a 4xx status code
func (o *DcimVirtualChassisPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual chassis partial update o k response has a 5xx status code
func (o *DcimVirtualChassisPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual chassis partial update o k response a status code equal to that given
func (o *DcimVirtualChassisPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimVirtualChassisPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualChassisPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualChassisPartialUpdateOK) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualChassisPartialUpdateBadRequest creates a DcimVirtualChassisPartialUpdateBadRequest with default headers values
func NewDcimVirtualChassisPartialUpdateBadRequest() *DcimVirtualChassisPartialUpdateBadRequest {
	return &DcimVirtualChassisPartialUpdateBadRequest{}
}

/*
DcimVirtualChassisPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimVirtualChassisPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim virtual chassis partial update bad request response has a 2xx status code
func (o *DcimVirtualChassisPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim virtual chassis partial update bad request response has a 3xx status code
func (o *DcimVirtualChassisPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual chassis partial update bad request response has a 4xx status code
func (o *DcimVirtualChassisPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim virtual chassis partial update bad request response has a 5xx status code
func (o *DcimVirtualChassisPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual chassis partial update bad request response a status code equal to that given
func (o *DcimVirtualChassisPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimVirtualChassisPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimVirtualChassisPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimVirtualChassisPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualChassisPartialUpdateDefault creates a DcimVirtualChassisPartialUpdateDefault with default headers values
func NewDcimVirtualChassisPartialUpdateDefault(code int) *DcimVirtualChassisPartialUpdateDefault {
	return &DcimVirtualChassisPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimVirtualChassisPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimVirtualChassisPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim virtual chassis partial update default response
func (o *DcimVirtualChassisPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim virtual chassis partial update default response has a 2xx status code
func (o *DcimVirtualChassisPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim virtual chassis partial update default response has a 3xx status code
func (o *DcimVirtualChassisPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim virtual chassis partial update default response has a 4xx status code
func (o *DcimVirtualChassisPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim virtual chassis partial update default response has a 5xx status code
func (o *DcimVirtualChassisPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim virtual chassis partial update default response a status code equal to that given
func (o *DcimVirtualChassisPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimVirtualChassisPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-chassis/{id}/][%d] dcim_virtual-chassis_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualChassisPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/virtual-chassis/{id}/][%d] dcim_virtual-chassis_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualChassisPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
