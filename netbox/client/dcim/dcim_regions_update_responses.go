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

// DcimRegionsUpdateReader is a Reader for the DcimRegionsUpdate structure.
type DcimRegionsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRegionsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimRegionsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimRegionsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRegionsUpdateOK creates a DcimRegionsUpdateOK with default headers values
func NewDcimRegionsUpdateOK() *DcimRegionsUpdateOK {
	return &DcimRegionsUpdateOK{}
}

/*
DcimRegionsUpdateOK describes a response with status code 200, with default header values.

DcimRegionsUpdateOK dcim regions update o k
*/
type DcimRegionsUpdateOK struct {
	Payload *models.Region
}

// IsSuccess returns true when this dcim regions update o k response has a 2xx status code
func (o *DcimRegionsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim regions update o k response has a 3xx status code
func (o *DcimRegionsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions update o k response has a 4xx status code
func (o *DcimRegionsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim regions update o k response has a 5xx status code
func (o *DcimRegionsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions update o k response a status code equal to that given
func (o *DcimRegionsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRegionsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcimRegionsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcimRegionsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsUpdateOK) GetPayload() *models.Region {
	return o.Payload
}

func (o *DcimRegionsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Region)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRegionsUpdateBadRequest creates a DcimRegionsUpdateBadRequest with default headers values
func NewDcimRegionsUpdateBadRequest() *DcimRegionsUpdateBadRequest {
	return &DcimRegionsUpdateBadRequest{}
}

/*
DcimRegionsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimRegionsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim regions update bad request response has a 2xx status code
func (o *DcimRegionsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim regions update bad request response has a 3xx status code
func (o *DcimRegionsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions update bad request response has a 4xx status code
func (o *DcimRegionsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim regions update bad request response has a 5xx status code
func (o *DcimRegionsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions update bad request response a status code equal to that given
func (o *DcimRegionsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimRegionsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcimRegionsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRegionsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcimRegionsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRegionsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRegionsUpdateDefault creates a DcimRegionsUpdateDefault with default headers values
func NewDcimRegionsUpdateDefault(code int) *DcimRegionsUpdateDefault {
	return &DcimRegionsUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRegionsUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimRegionsUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim regions update default response
func (o *DcimRegionsUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim regions update default response has a 2xx status code
func (o *DcimRegionsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim regions update default response has a 3xx status code
func (o *DcimRegionsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim regions update default response has a 4xx status code
func (o *DcimRegionsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim regions update default response has a 5xx status code
func (o *DcimRegionsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim regions update default response a status code equal to that given
func (o *DcimRegionsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRegionsUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcim_regions_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcim_regions_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRegionsUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
