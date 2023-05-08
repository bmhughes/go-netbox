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

// VirtualizationClusterGroupsReadReader is a Reader for the VirtualizationClusterGroupsRead structure.
type VirtualizationClusterGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClusterGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualizationClusterGroupsReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsReadOK creates a VirtualizationClusterGroupsReadOK with default headers values
func NewVirtualizationClusterGroupsReadOK() *VirtualizationClusterGroupsReadOK {
	return &VirtualizationClusterGroupsReadOK{}
}

/*
VirtualizationClusterGroupsReadOK describes a response with status code 200, with default header values.

VirtualizationClusterGroupsReadOK virtualization cluster groups read o k
*/
type VirtualizationClusterGroupsReadOK struct {
	Payload *models.ClusterGroup
}

// IsSuccess returns true when this virtualization cluster groups read o k response has a 2xx status code
func (o *VirtualizationClusterGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization cluster groups read o k response has a 3xx status code
func (o *VirtualizationClusterGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups read o k response has a 4xx status code
func (o *VirtualizationClusterGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization cluster groups read o k response has a 5xx status code
func (o *VirtualizationClusterGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups read o k response a status code equal to that given
func (o *VirtualizationClusterGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationClusterGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsReadOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsReadOK) String() string {
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsReadOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsReadOK) GetPayload() *models.ClusterGroup {
	return o.Payload
}

func (o *VirtualizationClusterGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClusterGroupsReadBadRequest creates a VirtualizationClusterGroupsReadBadRequest with default headers values
func NewVirtualizationClusterGroupsReadBadRequest() *VirtualizationClusterGroupsReadBadRequest {
	return &VirtualizationClusterGroupsReadBadRequest{}
}

/*
VirtualizationClusterGroupsReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VirtualizationClusterGroupsReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this virtualization cluster groups read bad request response has a 2xx status code
func (o *VirtualizationClusterGroupsReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this virtualization cluster groups read bad request response has a 3xx status code
func (o *VirtualizationClusterGroupsReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization cluster groups read bad request response has a 4xx status code
func (o *VirtualizationClusterGroupsReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this virtualization cluster groups read bad request response has a 5xx status code
func (o *VirtualizationClusterGroupsReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization cluster groups read bad request response a status code equal to that given
func (o *VirtualizationClusterGroupsReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *VirtualizationClusterGroupsReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsReadBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationClusterGroupsReadBadRequest) String() string {
	return fmt.Sprintf("[GET /virtualization/cluster-groups/{id}/][%d] virtualizationClusterGroupsReadBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationClusterGroupsReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClusterGroupsReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
