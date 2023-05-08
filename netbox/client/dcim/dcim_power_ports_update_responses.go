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

// DcimPowerPortsUpdateReader is a Reader for the DcimPowerPortsUpdate structure.
type DcimPowerPortsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPortsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimPowerPortsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPowerPortsUpdateOK creates a DcimPowerPortsUpdateOK with default headers values
func NewDcimPowerPortsUpdateOK() *DcimPowerPortsUpdateOK {
	return &DcimPowerPortsUpdateOK{}
}

/*
DcimPowerPortsUpdateOK describes a response with status code 200, with default header values.

DcimPowerPortsUpdateOK dcim power ports update o k
*/
type DcimPowerPortsUpdateOK struct {
	Payload *models.PowerPort
}

// IsSuccess returns true when this dcim power ports update o k response has a 2xx status code
func (o *DcimPowerPortsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power ports update o k response has a 3xx status code
func (o *DcimPowerPortsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports update o k response has a 4xx status code
func (o *DcimPowerPortsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power ports update o k response has a 5xx status code
func (o *DcimPowerPortsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports update o k response a status code equal to that given
func (o *DcimPowerPortsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerPortsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-ports/{id}/][%d] dcimPowerPortsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/power-ports/{id}/][%d] dcimPowerPortsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPortsUpdateOK) GetPayload() *models.PowerPort {
	return o.Payload
}

func (o *DcimPowerPortsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPortsUpdateBadRequest creates a DcimPowerPortsUpdateBadRequest with default headers values
func NewDcimPowerPortsUpdateBadRequest() *DcimPowerPortsUpdateBadRequest {
	return &DcimPowerPortsUpdateBadRequest{}
}

/*
DcimPowerPortsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimPowerPortsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim power ports update bad request response has a 2xx status code
func (o *DcimPowerPortsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim power ports update bad request response has a 3xx status code
func (o *DcimPowerPortsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power ports update bad request response has a 4xx status code
func (o *DcimPowerPortsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim power ports update bad request response has a 5xx status code
func (o *DcimPowerPortsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power ports update bad request response a status code equal to that given
func (o *DcimPowerPortsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimPowerPortsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-ports/{id}/][%d] dcimPowerPortsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerPortsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/power-ports/{id}/][%d] dcimPowerPortsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerPortsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPortsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
