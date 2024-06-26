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

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// VirtualizationVirtualMachinesUpdateReader is a Reader for the VirtualizationVirtualMachinesUpdate structure.
type VirtualizationVirtualMachinesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationVirtualMachinesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationVirtualMachinesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualizationVirtualMachinesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVirtualizationVirtualMachinesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationVirtualMachinesUpdateOK creates a VirtualizationVirtualMachinesUpdateOK with default headers values
func NewVirtualizationVirtualMachinesUpdateOK() *VirtualizationVirtualMachinesUpdateOK {
	return &VirtualizationVirtualMachinesUpdateOK{}
}

/*
VirtualizationVirtualMachinesUpdateOK describes a response with status code 200, with default header values.

VirtualizationVirtualMachinesUpdateOK virtualization virtual machines update o k
*/
type VirtualizationVirtualMachinesUpdateOK struct {
	Payload *models.VirtualMachineWithConfigContext
}

// IsSuccess returns true when this virtualization virtual machines update o k response has a 2xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization virtual machines update o k response has a 3xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual machines update o k response has a 4xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization virtual machines update o k response has a 5xx status code
func (o *VirtualizationVirtualMachinesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual machines update o k response a status code equal to that given
func (o *VirtualizationVirtualMachinesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationVirtualMachinesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateOK) GetPayload() *models.VirtualMachineWithConfigContext {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualMachineWithConfigContext)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationVirtualMachinesUpdateBadRequest creates a VirtualizationVirtualMachinesUpdateBadRequest with default headers values
func NewVirtualizationVirtualMachinesUpdateBadRequest() *VirtualizationVirtualMachinesUpdateBadRequest {
	return &VirtualizationVirtualMachinesUpdateBadRequest{}
}

/*
VirtualizationVirtualMachinesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VirtualizationVirtualMachinesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this virtualization virtual machines update bad request response has a 2xx status code
func (o *VirtualizationVirtualMachinesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this virtualization virtual machines update bad request response has a 3xx status code
func (o *VirtualizationVirtualMachinesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization virtual machines update bad request response has a 4xx status code
func (o *VirtualizationVirtualMachinesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this virtualization virtual machines update bad request response has a 5xx status code
func (o *VirtualizationVirtualMachinesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization virtual machines update bad request response a status code equal to that given
func (o *VirtualizationVirtualMachinesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *VirtualizationVirtualMachinesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualizationVirtualMachinesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationVirtualMachinesUpdateDefault creates a VirtualizationVirtualMachinesUpdateDefault with default headers values
func NewVirtualizationVirtualMachinesUpdateDefault(code int) *VirtualizationVirtualMachinesUpdateDefault {
	return &VirtualizationVirtualMachinesUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationVirtualMachinesUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type VirtualizationVirtualMachinesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization virtual machines update default response
func (o *VirtualizationVirtualMachinesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization virtual machines update default response has a 2xx status code
func (o *VirtualizationVirtualMachinesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization virtual machines update default response has a 3xx status code
func (o *VirtualizationVirtualMachinesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization virtual machines update default response has a 4xx status code
func (o *VirtualizationVirtualMachinesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization virtual machines update default response has a 5xx status code
func (o *VirtualizationVirtualMachinesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization virtual machines update default response a status code equal to that given
func (o *VirtualizationVirtualMachinesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationVirtualMachinesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualization_virtual-machines_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /virtualization/virtual-machines/{id}/][%d] virtualization_virtual-machines_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationVirtualMachinesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationVirtualMachinesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
