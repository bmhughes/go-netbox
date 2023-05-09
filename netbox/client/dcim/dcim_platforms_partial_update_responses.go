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

// DcimPlatformsPartialUpdateReader is a Reader for the DcimPlatformsPartialUpdate structure.
type DcimPlatformsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPlatformsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimPlatformsPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimPlatformsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPlatformsPartialUpdateOK creates a DcimPlatformsPartialUpdateOK with default headers values
func NewDcimPlatformsPartialUpdateOK() *DcimPlatformsPartialUpdateOK {
	return &DcimPlatformsPartialUpdateOK{}
}

/*
DcimPlatformsPartialUpdateOK describes a response with status code 200, with default header values.

DcimPlatformsPartialUpdateOK dcim platforms partial update o k
*/
type DcimPlatformsPartialUpdateOK struct {
	Payload *models.Platform
}

// IsSuccess returns true when this dcim platforms partial update o k response has a 2xx status code
func (o *DcimPlatformsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim platforms partial update o k response has a 3xx status code
func (o *DcimPlatformsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms partial update o k response has a 4xx status code
func (o *DcimPlatformsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim platforms partial update o k response has a 5xx status code
func (o *DcimPlatformsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms partial update o k response a status code equal to that given
func (o *DcimPlatformsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPlatformsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/{id}/][%d] dcimPlatformsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPlatformsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/{id}/][%d] dcimPlatformsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPlatformsPartialUpdateOK) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPlatformsPartialUpdateBadRequest creates a DcimPlatformsPartialUpdateBadRequest with default headers values
func NewDcimPlatformsPartialUpdateBadRequest() *DcimPlatformsPartialUpdateBadRequest {
	return &DcimPlatformsPartialUpdateBadRequest{}
}

/*
DcimPlatformsPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimPlatformsPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim platforms partial update bad request response has a 2xx status code
func (o *DcimPlatformsPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim platforms partial update bad request response has a 3xx status code
func (o *DcimPlatformsPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms partial update bad request response has a 4xx status code
func (o *DcimPlatformsPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim platforms partial update bad request response has a 5xx status code
func (o *DcimPlatformsPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms partial update bad request response a status code equal to that given
func (o *DcimPlatformsPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimPlatformsPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/{id}/][%d] dcimPlatformsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPlatformsPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/{id}/][%d] dcimPlatformsPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPlatformsPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPlatformsPartialUpdateDefault creates a DcimPlatformsPartialUpdateDefault with default headers values
func NewDcimPlatformsPartialUpdateDefault(code int) *DcimPlatformsPartialUpdateDefault {
	return &DcimPlatformsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPlatformsPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimPlatformsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim platforms partial update default response
func (o *DcimPlatformsPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim platforms partial update default response has a 2xx status code
func (o *DcimPlatformsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim platforms partial update default response has a 3xx status code
func (o *DcimPlatformsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim platforms partial update default response has a 4xx status code
func (o *DcimPlatformsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim platforms partial update default response has a 5xx status code
func (o *DcimPlatformsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim platforms partial update default response a status code equal to that given
func (o *DcimPlatformsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPlatformsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/{id}/][%d] dcim_platforms_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/{id}/][%d] dcim_platforms_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPlatformsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
