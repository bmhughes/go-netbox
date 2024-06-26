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

// VirtualizationClustersReadReader is a Reader for the VirtualizationClustersRead structure.
type VirtualizationClustersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVirtualizationClustersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewVirtualizationClustersReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewVirtualizationClustersReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationClustersReadOK creates a VirtualizationClustersReadOK with default headers values
func NewVirtualizationClustersReadOK() *VirtualizationClustersReadOK {
	return &VirtualizationClustersReadOK{}
}

/*
VirtualizationClustersReadOK describes a response with status code 200, with default header values.

VirtualizationClustersReadOK virtualization clusters read o k
*/
type VirtualizationClustersReadOK struct {
	Payload *models.Cluster
}

// IsSuccess returns true when this virtualization clusters read o k response has a 2xx status code
func (o *VirtualizationClustersReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization clusters read o k response has a 3xx status code
func (o *VirtualizationClustersReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters read o k response has a 4xx status code
func (o *VirtualizationClustersReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization clusters read o k response has a 5xx status code
func (o *VirtualizationClustersReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters read o k response a status code equal to that given
func (o *VirtualizationClustersReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *VirtualizationClustersReadOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualizationClustersReadOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersReadOK) String() string {
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualizationClustersReadOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClustersReadOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *VirtualizationClustersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClustersReadBadRequest creates a VirtualizationClustersReadBadRequest with default headers values
func NewVirtualizationClustersReadBadRequest() *VirtualizationClustersReadBadRequest {
	return &VirtualizationClustersReadBadRequest{}
}

/*
VirtualizationClustersReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type VirtualizationClustersReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this virtualization clusters read bad request response has a 2xx status code
func (o *VirtualizationClustersReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this virtualization clusters read bad request response has a 3xx status code
func (o *VirtualizationClustersReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization clusters read bad request response has a 4xx status code
func (o *VirtualizationClustersReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this virtualization clusters read bad request response has a 5xx status code
func (o *VirtualizationClustersReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization clusters read bad request response a status code equal to that given
func (o *VirtualizationClustersReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *VirtualizationClustersReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualizationClustersReadBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationClustersReadBadRequest) String() string {
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualizationClustersReadBadRequest  %+v", 400, o.Payload)
}

func (o *VirtualizationClustersReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVirtualizationClustersReadDefault creates a VirtualizationClustersReadDefault with default headers values
func NewVirtualizationClustersReadDefault(code int) *VirtualizationClustersReadDefault {
	return &VirtualizationClustersReadDefault{
		_statusCode: code,
	}
}

/*
VirtualizationClustersReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type VirtualizationClustersReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization clusters read default response
func (o *VirtualizationClustersReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization clusters read default response has a 2xx status code
func (o *VirtualizationClustersReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization clusters read default response has a 3xx status code
func (o *VirtualizationClustersReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization clusters read default response has a 4xx status code
func (o *VirtualizationClustersReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization clusters read default response has a 5xx status code
func (o *VirtualizationClustersReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization clusters read default response a status code equal to that given
func (o *VirtualizationClustersReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationClustersReadDefault) Error() string {
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualization_clusters_read default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersReadDefault) String() string {
	return fmt.Sprintf("[GET /virtualization/clusters/{id}/][%d] virtualization_clusters_read default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationClustersReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationClustersReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
