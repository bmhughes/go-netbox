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

// VirtualizationClustersPartialUpdateReader is a Reader for the VirtualizationClustersPartialUpdate structure.
type VirtualizationClustersPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualizationClustersPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVirtualizationClustersPartialUpdateOK creates a VirtualizationClustersPartialUpdateOK with default headers values
func NewVirtualizationClustersPartialUpdateOK() *VirtualizationClustersPartialUpdateOK {
	return &VirtualizationClustersPartialUpdateOK{}
}

/*
VirtualizationClustersPartialUpdateOK describes a response with status code 200, with default header values.

VirtualizationClustersPartialUpdateOK virtualization clusters partial update o k
*/
type VirtualizationClustersPartialUpdateOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this virtualization clusters partial update o k response has a 2xx status code
func (o *VirtualizationClustersPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters partial update o k response has a 3xx status code
func (o *VirtualizationClustersPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters partial update o k response has a 4xx status code
func (o *VirtualizationClustersPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters partial update o k response has a 5xx status code
func (o *VirtualizationClustersPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters partial update o k response a status code equal to that given
func (o *VirtualizationClustersPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationClustersPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/clusters/{id}/][%d] virtualizationClustersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /virtualization/clusters/{id}/][%d] virtualizationClustersPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersPartialUpdateOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClustersPartialUpdateBadRequest creates a VirtualizationClustersPartialUpdateBadRequest with default headers values
func NewVirtualizationClustersPartialUpdateBadRequest() *VirtualizationClustersPartialUpdateBadRequest {
	return &VirtualizationClustersPartialUpdateBadRequest{}
}

/*
VirtualizationClustersPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VirtualizationClustersPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this virtualization clusters partial update bad request response has a 2xx status code
func (o *VirtualizationClustersPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this virtualization clusters partial update bad request response has a 3xx status code
func (o *VirtualizationClustersPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters partial update bad request response has a 4xx status code
func (o *VirtualizationClustersPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this virtualization clusters partial update bad request response has a 5xx status code
func (o *VirtualizationClustersPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters partial update bad request response a status code equal to that given
func (o *VirtualizationClustersPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *VirtualizationClustersPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /virtualization/clusters/{id}/][%d] virtualizationClustersPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationClustersPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /virtualization/clusters/{id}/][%d] virtualizationClustersPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationClustersPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
