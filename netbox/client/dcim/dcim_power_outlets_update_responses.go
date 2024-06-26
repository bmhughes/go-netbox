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

// DcimPowerOutletsUpdateReader is a Reader for the DcimPowerOutletsUpdate structure.
type DcimPowerOutletsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerOutletsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerOutletsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimPowerOutletsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimPowerOutletsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerOutletsUpdateOK creates a DcimPowerOutletsUpdateOK with default headers values
func NewDcimPowerOutletsUpdateOK() *DcimPowerOutletsUpdateOK {
	return &DcimPowerOutletsUpdateOK{}
}

/*
DcimPowerOutletsUpdateOK describes a response with status code 200, with default header values.

DcimPowerOutletsUpdateOK dcim power outlets update o k
*/
type DcimPowerOutletsUpdateOK struct {
	Payload *models.PowerOutlet
}

// IsSuccess returns true when this dcim power outlets update o k response has a 2xx status code
func (o *DcimPowerOutletsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power outlets update o k response has a 3xx status code
func (o *DcimPowerOutletsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets update o k response has a 4xx status code
func (o *DcimPowerOutletsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power outlets update o k response has a 5xx status code
func (o *DcimPowerOutletsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets update o k response a status code equal to that given
func (o *DcimPowerOutletsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerOutletsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/{id}/][%d] dcimPowerOutletsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/{id}/][%d] dcimPowerOutletsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerOutletsUpdateOK) GetPayload() *models.PowerOutlet {
	return o.Payload
}

func (o *DcimPowerOutletsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerOutlet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletsUpdateBadRequest creates a DcimPowerOutletsUpdateBadRequest with default headers values
func NewDcimPowerOutletsUpdateBadRequest() *DcimPowerOutletsUpdateBadRequest {
	return &DcimPowerOutletsUpdateBadRequest{}
}

/*
DcimPowerOutletsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimPowerOutletsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim power outlets update bad request response has a 2xx status code
func (o *DcimPowerOutletsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim power outlets update bad request response has a 3xx status code
func (o *DcimPowerOutletsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power outlets update bad request response has a 4xx status code
func (o *DcimPowerOutletsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim power outlets update bad request response has a 5xx status code
func (o *DcimPowerOutletsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power outlets update bad request response a status code equal to that given
func (o *DcimPowerOutletsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimPowerOutletsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/{id}/][%d] dcimPowerOutletsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerOutletsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/{id}/][%d] dcimPowerOutletsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerOutletsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerOutletsUpdateDefault creates a DcimPowerOutletsUpdateDefault with default headers values
func NewDcimPowerOutletsUpdateDefault(code int) *DcimPowerOutletsUpdateDefault {
	return &DcimPowerOutletsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerOutletsUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimPowerOutletsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power outlets update default response
func (o *DcimPowerOutletsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power outlets update default response has a 2xx status code
func (o *DcimPowerOutletsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power outlets update default response has a 3xx status code
func (o *DcimPowerOutletsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power outlets update default response has a 4xx status code
func (o *DcimPowerOutletsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power outlets update default response has a 5xx status code
func (o *DcimPowerOutletsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power outlets update default response a status code equal to that given
func (o *DcimPowerOutletsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerOutletsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/{id}/][%d] dcim_power-outlets_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/power-outlets/{id}/][%d] dcim_power-outlets_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerOutletsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerOutletsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
