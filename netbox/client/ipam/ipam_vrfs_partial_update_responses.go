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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// IpamVrfsPartialUpdateReader is a Reader for the IpamVrfsPartialUpdate structure.
type IpamVrfsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVrfsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpamVrfsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamVrfsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVrfsPartialUpdateOK creates a IpamVrfsPartialUpdateOK with default headers values
func NewIpamVrfsPartialUpdateOK() *IpamVrfsPartialUpdateOK {
	return &IpamVrfsPartialUpdateOK{}
}

/*
IpamVrfsPartialUpdateOK describes a response with status code 200, with default header values.

IpamVrfsPartialUpdateOK ipam vrfs partial update o k
*/
type IpamVrfsPartialUpdateOK struct {
	Payload *models.VRF
}

// IsSuccess returns true when this ipam vrfs partial update o k response has a 2xx status code
func (o *IpamVrfsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vrfs partial update o k response has a 3xx status code
func (o *IpamVrfsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs partial update o k response has a 4xx status code
func (o *IpamVrfsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vrfs partial update o k response has a 5xx status code
func (o *IpamVrfsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs partial update o k response a status code equal to that given
func (o *IpamVrfsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpamVrfsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/{id}/][%d] ipamVrfsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVrfsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/{id}/][%d] ipamVrfsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IpamVrfsPartialUpdateOK) GetPayload() *models.VRF {
	return o.Payload
}

func (o *IpamVrfsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamVrfsPartialUpdateBadRequest creates a IpamVrfsPartialUpdateBadRequest with default headers values
func NewIpamVrfsPartialUpdateBadRequest() *IpamVrfsPartialUpdateBadRequest {
	return &IpamVrfsPartialUpdateBadRequest{}
}

/*
IpamVrfsPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamVrfsPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam vrfs partial update bad request response has a 2xx status code
func (o *IpamVrfsPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam vrfs partial update bad request response has a 3xx status code
func (o *IpamVrfsPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vrfs partial update bad request response has a 4xx status code
func (o *IpamVrfsPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam vrfs partial update bad request response has a 5xx status code
func (o *IpamVrfsPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vrfs partial update bad request response a status code equal to that given
func (o *IpamVrfsPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamVrfsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/{id}/][%d] ipamVrfsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamVrfsPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /ipam/vrfs/{id}/][%d] ipamVrfsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *IpamVrfsPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamVrfsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
