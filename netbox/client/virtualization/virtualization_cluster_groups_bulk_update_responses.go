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

// VirtualizationClusterGroupsBulkUpdateReader is a Reader for the VirtualizationClusterGroupsBulkUpdate structure.
type VirtualizationClusterGroupsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualizationClusterGroupsBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsBulkUpdateOK creates a VirtualizationClusterGroupsBulkUpdateOK with default headers values
func NewVirtualizationClusterGroupsBulkUpdateOK() *VirtualizationClusterGroupsBulkUpdateOK {
	return &VirtualizationClusterGroupsBulkUpdateOK{}
}

/*
VirtualizationClusterGroupsBulkUpdateOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsBulkUpdateOK virtualization cluster groups bulk update o k
*/
type VirtualizationClusterGroupsBulkUpdateOK struct {
	Payload *models.ClusterGroup
}

// IsSuccess returns true when this virtualization cluster groups bulk update o k response has a 2xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster groups bulk update o k response has a 3xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups bulk update o k response has a 4xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster groups bulk update o k response has a 5xx status code
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups bulk update o k response a status code equal to that given
func (o *VirtualizationClusterGroupsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/][%d] virtualizationClusterGroupsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/][%d] virtualizationClusterGroupsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterGroupsBulkUpdateBadRequest creates a VirtualizationClusterGroupsBulkUpdateBadRequest with default headers values
func NewVirtualizationClusterGroupsBulkUpdateBadRequest() *VirtualizationClusterGroupsBulkUpdateBadRequest {
	return &VirtualizationClusterGroupsBulkUpdateBadRequest{}
}

/*
VirtualizationClusterGroupsBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VirtualizationClusterGroupsBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this virtualization cluster groups bulk update bad request response has a 2xx status code
func (o *VirtualizationClusterGroupsBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this virtualization cluster groups bulk update bad request response has a 3xx status code
func (o *VirtualizationClusterGroupsBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups bulk update bad request response has a 4xx status code
func (o *VirtualizationClusterGroupsBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this virtualization cluster groups bulk update bad request response has a 5xx status code
func (o *VirtualizationClusterGroupsBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups bulk update bad request response a status code equal to that given
func (o *VirtualizationClusterGroupsBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *VirtualizationClusterGroupsBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/][%d] virtualizationClusterGroupsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationClusterGroupsBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /virtualization/cluster-groups/][%d] virtualizationClusterGroupsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationClusterGroupsBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterGroupsBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
