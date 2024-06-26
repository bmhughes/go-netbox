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

// DcimLocationsUpdateReader is a Reader for the DcimLocationsUpdate structure.
type DcimLocationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimLocationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimLocationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimLocationsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimLocationsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimLocationsUpdateOK creates a DcimLocationsUpdateOK with default headers values
func NewDcimLocationsUpdateOK() *DcimLocationsUpdateOK {
	return &DcimLocationsUpdateOK{}
}

/*
DcimLocationsUpdateOK describes a response with status code 200, with default header values.

DcimLocationsUpdateOK dcim locations update o k
*/
type DcimLocationsUpdateOK struct {
	Payload *models.Location
}

// IsSuccess returns true when this dcim locations update o k response has a 2xx status code
func (o *DcimLocationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim locations update o k response has a 3xx status code
func (o *DcimLocationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations update o k response has a 4xx status code
func (o *DcimLocationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim locations update o k response has a 5xx status code
func (o *DcimLocationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations update o k response a status code equal to that given
func (o *DcimLocationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimLocationsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcimLocationsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimLocationsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcimLocationsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimLocationsUpdateOK) GetPayload() *models.Location {
	return o.Payload
}

func (o *DcimLocationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Location)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimLocationsUpdateBadRequest creates a DcimLocationsUpdateBadRequest with default headers values
func NewDcimLocationsUpdateBadRequest() *DcimLocationsUpdateBadRequest {
	return &DcimLocationsUpdateBadRequest{}
}

/*
DcimLocationsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimLocationsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim locations update bad request response has a 2xx status code
func (o *DcimLocationsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim locations update bad request response has a 3xx status code
func (o *DcimLocationsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim locations update bad request response has a 4xx status code
func (o *DcimLocationsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim locations update bad request response has a 5xx status code
func (o *DcimLocationsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim locations update bad request response a status code equal to that given
func (o *DcimLocationsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimLocationsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcimLocationsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimLocationsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcimLocationsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimLocationsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimLocationsUpdateDefault creates a DcimLocationsUpdateDefault with default headers values
func NewDcimLocationsUpdateDefault(code int) *DcimLocationsUpdateDefault {
	return &DcimLocationsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimLocationsUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimLocationsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim locations update default response
func (o *DcimLocationsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim locations update default response has a 2xx status code
func (o *DcimLocationsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim locations update default response has a 3xx status code
func (o *DcimLocationsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim locations update default response has a 4xx status code
func (o *DcimLocationsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim locations update default response has a 5xx status code
func (o *DcimLocationsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim locations update default response a status code equal to that given
func (o *DcimLocationsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimLocationsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcim_locations_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/locations/{id}/][%d] dcim_locations_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimLocationsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimLocationsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
