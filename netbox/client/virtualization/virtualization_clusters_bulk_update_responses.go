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

// VirtualizationClustersBulkUpdateReader is a Reader for the VirtualizationClustersBulkUpdate structure.
type VirtualizationClustersBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualizationClustersBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVirtualizationClustersBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClustersBulkUpdateOK creates a VirtualizationClustersBulkUpdateOK with default headers values
func NewVirtualizationClustersBulkUpdateOK() *VirtualizationClustersBulkUpdateOK {
	return &VirtualizationClustersBulkUpdateOK{}
}

/*
VirtualizationClustersBulkUpdateOK describes a response with status code 200, with default header values.

VirtualizationClustersBulkUpdateOK virtualization clusters bulk update o k
*/
type VirtualizationClustersBulkUpdateOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this virtualization clusters bulk update o k response has a 2xx status code
func (o *VirtualizationClustersBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters bulk update o k response has a 3xx status code
func (o *VirtualizationClustersBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters bulk update o k response has a 4xx status code
func (o *VirtualizationClustersBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters bulk update o k response has a 5xx status code
func (o *VirtualizationClustersBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters bulk update o k response a status code equal to that given
func (o *VirtualizationClustersBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationClustersBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualizationClustersBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualizationClustersBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersBulkUpdateOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClustersBulkUpdateBadRequest creates a VirtualizationClustersBulkUpdateBadRequest with default headers values
func NewVirtualizationClustersBulkUpdateBadRequest() *VirtualizationClustersBulkUpdateBadRequest {
	return &VirtualizationClustersBulkUpdateBadRequest{}
}

/*
VirtualizationClustersBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VirtualizationClustersBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this virtualization clusters bulk update bad request response has a 2xx status code
func (o *VirtualizationClustersBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this virtualization clusters bulk update bad request response has a 3xx status code
func (o *VirtualizationClustersBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters bulk update bad request response has a 4xx status code
func (o *VirtualizationClustersBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this virtualization clusters bulk update bad request response has a 5xx status code
func (o *VirtualizationClustersBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters bulk update bad request response a status code equal to that given
func (o *VirtualizationClustersBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *VirtualizationClustersBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualizationClustersBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationClustersBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualizationClustersBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationClustersBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClustersBulkUpdateDefault creates a VirtualizationClustersBulkUpdateDefault with default headers values
func NewVirtualizationClustersBulkUpdateDefault(code int) *VirtualizationClustersBulkUpdateDefault {
	return &VirtualizationClustersBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClustersBulkUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type VirtualizationClustersBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization clusters bulk update default response
func (o *VirtualizationClustersBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization clusters bulk update default response has a 2xx status code
func (o *VirtualizationClustersBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization clusters bulk update default response has a 3xx status code
func (o *VirtualizationClustersBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization clusters bulk update default response has a 4xx status code
func (o *VirtualizationClustersBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization clusters bulk update default response has a 5xx status code
func (o *VirtualizationClustersBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization clusters bulk update default response a status code equal to that given
func (o *VirtualizationClustersBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationClustersBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualization_clusters_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /virtualization/clusters/][%d] virtualization_clusters_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
