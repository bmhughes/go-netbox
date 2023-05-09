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

// DcimPowerPanelsPartialUpdateReader is a Reader for the DcimPowerPanelsPartialUpdate structure.
type DcimPowerPanelsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPanelsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerPanelsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimPowerPanelsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimPowerPanelsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerPanelsPartialUpdateOK creates a DcimPowerPanelsPartialUpdateOK with default headers values
func NewDcimPowerPanelsPartialUpdateOK() *DcimPowerPanelsPartialUpdateOK {
	return &DcimPowerPanelsPartialUpdateOK{}
}

/*
DcimPowerPanelsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPowerPanelsPartialUpdateOK dcim power panels partial update o k
*/
type DcimPowerPanelsPartialUpdateOK struct {
	Payload *models.PowerPanel
}

// IsSuccess returns true when this dcim power panels partial update o k response has a 2xx status code
func (o *DcimPowerPanelsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power panels partial update o k response has a 3xx status code
func (o *DcimPowerPanelsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power panels partial update o k response has a 4xx status code
func (o *DcimPowerPanelsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power panels partial update o k response has a 5xx status code
func (o *DcimPowerPanelsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power panels partial update o k response a status code equal to that given
func (o *DcimPowerPanelsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerPanelsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/{id}/][%d] dcimPowerPanelsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPanelsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/{id}/][%d] dcimPowerPanelsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerPanelsPartialUpdateOK) GetPayload() *models.PowerPanel {
	return o.Payload
}

func (o *DcimPowerPanelsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerPanel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPanelsPartialUpdateBadRequest creates a DcimPowerPanelsPartialUpdateBadRequest with default headers values
func NewDcimPowerPanelsPartialUpdateBadRequest() *DcimPowerPanelsPartialUpdateBadRequest {
	return &DcimPowerPanelsPartialUpdateBadRequest{}
}

/*
DcimPowerPanelsPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimPowerPanelsPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim power panels partial update bad request response has a 2xx status code
func (o *DcimPowerPanelsPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim power panels partial update bad request response has a 3xx status code
func (o *DcimPowerPanelsPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power panels partial update bad request response has a 4xx status code
func (o *DcimPowerPanelsPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim power panels partial update bad request response has a 5xx status code
func (o *DcimPowerPanelsPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power panels partial update bad request response a status code equal to that given
func (o *DcimPowerPanelsPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimPowerPanelsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/{id}/][%d] dcimPowerPanelsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerPanelsPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/{id}/][%d] dcimPowerPanelsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerPanelsPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerPanelsPartialUpdateDefault creates a DcimPowerPanelsPartialUpdateDefault with default headers values
func NewDcimPowerPanelsPartialUpdateDefault(code int) *DcimPowerPanelsPartialUpdateDefault {
	return &DcimPowerPanelsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerPanelsPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimPowerPanelsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power panels partial update default response
func (o *DcimPowerPanelsPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power panels partial update default response has a 2xx status code
func (o *DcimPowerPanelsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power panels partial update default response has a 3xx status code
func (o *DcimPowerPanelsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power panels partial update default response has a 4xx status code
func (o *DcimPowerPanelsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power panels partial update default response has a 5xx status code
func (o *DcimPowerPanelsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power panels partial update default response a status code equal to that given
func (o *DcimPowerPanelsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerPanelsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/{id}/][%d] dcim_power-panels_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/power-panels/{id}/][%d] dcim_power-panels_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerPanelsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerPanelsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
