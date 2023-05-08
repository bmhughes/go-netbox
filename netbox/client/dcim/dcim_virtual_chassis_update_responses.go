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

// DcimVirtualChassisUpdateReader is a Reader for the DcimVirtualChassisUpdate structure.
type DcimVirtualChassisUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimVirtualChassisUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimVirtualChassisUpdateOK creates a DcimVirtualChassisUpdateOK with default headers values
func NewDcimVirtualChassisUpdateOK() *DcimVirtualChassisUpdateOK {
	return &DcimVirtualChassisUpdateOK{}
}

/*
DcimVirtualChassisUpdateOK describes a response with status code 200, with default header values.

DcimVirtualChassisUpdateOK dcim virtual chassis update o k
*/
type DcimVirtualChassisUpdateOK struct {
	Payload *models.VirtualChassis
}

// IsSuccess returns true when this dcim virtual chassis update o k response has a 2xx status code
func (o *DcimVirtualChassisUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual chassis update o k response has a 3xx status code
func (o *DcimVirtualChassisUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual chassis update o k response has a 4xx status code
func (o *DcimVirtualChassisUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual chassis update o k response has a 5xx status code
func (o *DcimVirtualChassisUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual chassis update o k response a status code equal to that given
func (o *DcimVirtualChassisUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimVirtualChassisUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualChassisUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualChassisUpdateOK) GetPayload() *models.VirtualChassis {
	return o.Payload
}

func (o *DcimVirtualChassisUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualChassis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualChassisUpdateBadRequest creates a DcimVirtualChassisUpdateBadRequest with default headers values
func NewDcimVirtualChassisUpdateBadRequest() *DcimVirtualChassisUpdateBadRequest {
	return &DcimVirtualChassisUpdateBadRequest{}
}

/*
DcimVirtualChassisUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimVirtualChassisUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim virtual chassis update bad request response has a 2xx status code
func (o *DcimVirtualChassisUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim virtual chassis update bad request response has a 3xx status code
func (o *DcimVirtualChassisUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual chassis update bad request response has a 4xx status code
func (o *DcimVirtualChassisUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim virtual chassis update bad request response has a 5xx status code
func (o *DcimVirtualChassisUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual chassis update bad request response a status code equal to that given
func (o *DcimVirtualChassisUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimVirtualChassisUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimVirtualChassisUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/virtual-chassis/{id}/][%d] dcimVirtualChassisUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimVirtualChassisUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
