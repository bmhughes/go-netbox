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

// DcimVirtualDeviceContextsReadReader is a Reader for the DcimVirtualDeviceContextsRead structure.
type DcimVirtualDeviceContextsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualDeviceContextsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualDeviceContextsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimVirtualDeviceContextsReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimVirtualDeviceContextsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualDeviceContextsReadOK creates a DcimVirtualDeviceContextsReadOK with default headers values
func NewDcimVirtualDeviceContextsReadOK() *DcimVirtualDeviceContextsReadOK {
	return &DcimVirtualDeviceContextsReadOK{}
}

/*
DcimVirtualDeviceContextsReadOK describes a response with status code 200, with default header values.

DcimVirtualDeviceContextsReadOK dcim virtual device contexts read o k
*/
type DcimVirtualDeviceContextsReadOK struct {
	Payload *models.VirtualDeviceContext
}

// IsSuccess returns true when this dcim virtual device contexts read o k response has a 2xx status code
func (o *DcimVirtualDeviceContextsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual device contexts read o k response has a 3xx status code
func (o *DcimVirtualDeviceContextsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual device contexts read o k response has a 4xx status code
func (o *DcimVirtualDeviceContextsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual device contexts read o k response has a 5xx status code
func (o *DcimVirtualDeviceContextsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual device contexts read o k response a status code equal to that given
func (o *DcimVirtualDeviceContextsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimVirtualDeviceContextsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/virtual-device-contexts/{id}/][%d] dcimVirtualDeviceContextsReadOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualDeviceContextsReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/virtual-device-contexts/{id}/][%d] dcimVirtualDeviceContextsReadOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualDeviceContextsReadOK) GetPayload() *models.VirtualDeviceContext {
	return o.Payload
}

func (o *DcimVirtualDeviceContextsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualDeviceContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualDeviceContextsReadBadRequest creates a DcimVirtualDeviceContextsReadBadRequest with default headers values
func NewDcimVirtualDeviceContextsReadBadRequest() *DcimVirtualDeviceContextsReadBadRequest {
	return &DcimVirtualDeviceContextsReadBadRequest{}
}

/*
DcimVirtualDeviceContextsReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimVirtualDeviceContextsReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim virtual device contexts read bad request response has a 2xx status code
func (o *DcimVirtualDeviceContextsReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim virtual device contexts read bad request response has a 3xx status code
func (o *DcimVirtualDeviceContextsReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual device contexts read bad request response has a 4xx status code
func (o *DcimVirtualDeviceContextsReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim virtual device contexts read bad request response has a 5xx status code
func (o *DcimVirtualDeviceContextsReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual device contexts read bad request response a status code equal to that given
func (o *DcimVirtualDeviceContextsReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimVirtualDeviceContextsReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /dcim/virtual-device-contexts/{id}/][%d] dcimVirtualDeviceContextsReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimVirtualDeviceContextsReadBadRequest) String() string {
	return fmt.Sprintf("[GET /dcim/virtual-device-contexts/{id}/][%d] dcimVirtualDeviceContextsReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimVirtualDeviceContextsReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualDeviceContextsReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualDeviceContextsReadDefault creates a DcimVirtualDeviceContextsReadDefault with default headers values
func NewDcimVirtualDeviceContextsReadDefault(code int) *DcimVirtualDeviceContextsReadDefault {
	return &DcimVirtualDeviceContextsReadDefault{
		_statusCode: code,
	}
}

/*
DcimVirtualDeviceContextsReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimVirtualDeviceContextsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim virtual device contexts read default response
func (o *DcimVirtualDeviceContextsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim virtual device contexts read default response has a 2xx status code
func (o *DcimVirtualDeviceContextsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim virtual device contexts read default response has a 3xx status code
func (o *DcimVirtualDeviceContextsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim virtual device contexts read default response has a 4xx status code
func (o *DcimVirtualDeviceContextsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim virtual device contexts read default response has a 5xx status code
func (o *DcimVirtualDeviceContextsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim virtual device contexts read default response a status code equal to that given
func (o *DcimVirtualDeviceContextsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimVirtualDeviceContextsReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/virtual-device-contexts/{id}/][%d] dcim_virtual-device-contexts_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualDeviceContextsReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/virtual-device-contexts/{id}/][%d] dcim_virtual-device-contexts_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualDeviceContextsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualDeviceContextsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
