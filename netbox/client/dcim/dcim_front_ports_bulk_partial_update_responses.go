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

// DcimFrontPortsBulkPartialUpdateReader is a Reader for the DcimFrontPortsBulkPartialUpdate structure.
type DcimFrontPortsBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimFrontPortsBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimFrontPortsBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortsBulkPartialUpdateOK creates a DcimFrontPortsBulkPartialUpdateOK with default headers values
func NewDcimFrontPortsBulkPartialUpdateOK() *DcimFrontPortsBulkPartialUpdateOK {
	return &DcimFrontPortsBulkPartialUpdateOK{}
}

/*
DcimFrontPortsBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortsBulkPartialUpdateOK dcim front ports bulk partial update o k
*/
type DcimFrontPortsBulkPartialUpdateOK struct {
	Payload *models.FrontPort
}

// IsSuccess returns true when this dcim front ports bulk partial update o k response has a 2xx status code
func (o *DcimFrontPortsBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front ports bulk partial update o k response has a 3xx status code
func (o *DcimFrontPortsBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front ports bulk partial update o k response has a 4xx status code
func (o *DcimFrontPortsBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front ports bulk partial update o k response has a 5xx status code
func (o *DcimFrontPortsBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front ports bulk partial update o k response a status code equal to that given
func (o *DcimFrontPortsBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimFrontPortsBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/][%d] dcimFrontPortsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortsBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/][%d] dcimFrontPortsBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortsBulkPartialUpdateOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsBulkPartialUpdateBadRequest creates a DcimFrontPortsBulkPartialUpdateBadRequest with default headers values
func NewDcimFrontPortsBulkPartialUpdateBadRequest() *DcimFrontPortsBulkPartialUpdateBadRequest {
	return &DcimFrontPortsBulkPartialUpdateBadRequest{}
}

/*
DcimFrontPortsBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimFrontPortsBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim front ports bulk partial update bad request response has a 2xx status code
func (o *DcimFrontPortsBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim front ports bulk partial update bad request response has a 3xx status code
func (o *DcimFrontPortsBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front ports bulk partial update bad request response has a 4xx status code
func (o *DcimFrontPortsBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim front ports bulk partial update bad request response has a 5xx status code
func (o *DcimFrontPortsBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front ports bulk partial update bad request response a status code equal to that given
func (o *DcimFrontPortsBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimFrontPortsBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/][%d] dcimFrontPortsBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimFrontPortsBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/][%d] dcimFrontPortsBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimFrontPortsBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsBulkPartialUpdateDefault creates a DcimFrontPortsBulkPartialUpdateDefault with default headers values
func NewDcimFrontPortsBulkPartialUpdateDefault(code int) *DcimFrontPortsBulkPartialUpdateDefault {
	return &DcimFrontPortsBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimFrontPortsBulkPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimFrontPortsBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front ports bulk partial update default response
func (o *DcimFrontPortsBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim front ports bulk partial update default response has a 2xx status code
func (o *DcimFrontPortsBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim front ports bulk partial update default response has a 3xx status code
func (o *DcimFrontPortsBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim front ports bulk partial update default response has a 4xx status code
func (o *DcimFrontPortsBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim front ports bulk partial update default response has a 5xx status code
func (o *DcimFrontPortsBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim front ports bulk partial update default response a status code equal to that given
func (o *DcimFrontPortsBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimFrontPortsBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/][%d] dcim_front-ports_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortsBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/][%d] dcim_front-ports_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortsBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
