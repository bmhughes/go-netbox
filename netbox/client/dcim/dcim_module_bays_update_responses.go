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

// DcimModuleBaysUpdateReader is a Reader for the DcimModuleBaysUpdate structure.
type DcimModuleBaysUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBaysUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBaysUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimModuleBaysUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimModuleBaysUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBaysUpdateOK creates a DcimModuleBaysUpdateOK with default headers values
func NewDcimModuleBaysUpdateOK() *DcimModuleBaysUpdateOK {
	return &DcimModuleBaysUpdateOK{}
}

/*
DcimModuleBaysUpdateOK describes a response with status code 200, with default header values.

DcimModuleBaysUpdateOK dcim module bays update o k
*/
type DcimModuleBaysUpdateOK struct {
	Payload *models.ModuleBay
}

// IsSuccess returns true when this dcim module bays update o k response has a 2xx status code
func (o *DcimModuleBaysUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bays update o k response has a 3xx status code
func (o *DcimModuleBaysUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays update o k response has a 4xx status code
func (o *DcimModuleBaysUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bays update o k response has a 5xx status code
func (o *DcimModuleBaysUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays update o k response a status code equal to that given
func (o *DcimModuleBaysUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModuleBaysUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/{id}/][%d] dcimModuleBaysUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/{id}/][%d] dcimModuleBaysUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysUpdateOK) GetPayload() *models.ModuleBay {
	return o.Payload
}

func (o *DcimModuleBaysUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ModuleBay)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBaysUpdateBadRequest creates a DcimModuleBaysUpdateBadRequest with default headers values
func NewDcimModuleBaysUpdateBadRequest() *DcimModuleBaysUpdateBadRequest {
	return &DcimModuleBaysUpdateBadRequest{}
}

/*
DcimModuleBaysUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimModuleBaysUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim module bays update bad request response has a 2xx status code
func (o *DcimModuleBaysUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim module bays update bad request response has a 3xx status code
func (o *DcimModuleBaysUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays update bad request response has a 4xx status code
func (o *DcimModuleBaysUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim module bays update bad request response has a 5xx status code
func (o *DcimModuleBaysUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays update bad request response a status code equal to that given
func (o *DcimModuleBaysUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimModuleBaysUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/{id}/][%d] dcimModuleBaysUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimModuleBaysUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/{id}/][%d] dcimModuleBaysUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimModuleBaysUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBaysUpdateDefault creates a DcimModuleBaysUpdateDefault with default headers values
func NewDcimModuleBaysUpdateDefault(code int) *DcimModuleBaysUpdateDefault {
	return &DcimModuleBaysUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBaysUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimModuleBaysUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bays update default response
func (o *DcimModuleBaysUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bays update default response has a 2xx status code
func (o *DcimModuleBaysUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bays update default response has a 3xx status code
func (o *DcimModuleBaysUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bays update default response has a 4xx status code
func (o *DcimModuleBaysUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bays update default response has a 5xx status code
func (o *DcimModuleBaysUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bays update default response a status code equal to that given
func (o *DcimModuleBaysUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBaysUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/{id}/][%d] dcim_module-bays_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/module-bays/{id}/][%d] dcim_module-bays_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
