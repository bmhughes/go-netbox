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

// DcimFrontPortTemplatesUpdateReader is a Reader for the DcimFrontPortTemplatesUpdate structure.
type DcimFrontPortTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimFrontPortTemplatesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimFrontPortTemplatesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortTemplatesUpdateOK creates a DcimFrontPortTemplatesUpdateOK with default headers values
func NewDcimFrontPortTemplatesUpdateOK() *DcimFrontPortTemplatesUpdateOK {
	return &DcimFrontPortTemplatesUpdateOK{}
}

/*
DcimFrontPortTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesUpdateOK dcim front port templates update o k
*/
type DcimFrontPortTemplatesUpdateOK struct {
	Payload *models.FrontPortTemplate
}

// IsSuccess returns true when this dcim front port templates update o k response has a 2xx status code
func (o *DcimFrontPortTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front port templates update o k response has a 3xx status code
func (o *DcimFrontPortTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front port templates update o k response has a 4xx status code
func (o *DcimFrontPortTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front port templates update o k response has a 5xx status code
func (o *DcimFrontPortTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front port templates update o k response a status code equal to that given
func (o *DcimFrontPortTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimFrontPortTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortTemplatesUpdateOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesUpdateBadRequest creates a DcimFrontPortTemplatesUpdateBadRequest with default headers values
func NewDcimFrontPortTemplatesUpdateBadRequest() *DcimFrontPortTemplatesUpdateBadRequest {
	return &DcimFrontPortTemplatesUpdateBadRequest{}
}

/*
DcimFrontPortTemplatesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimFrontPortTemplatesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim front port templates update bad request response has a 2xx status code
func (o *DcimFrontPortTemplatesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim front port templates update bad request response has a 3xx status code
func (o *DcimFrontPortTemplatesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front port templates update bad request response has a 4xx status code
func (o *DcimFrontPortTemplatesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim front port templates update bad request response has a 5xx status code
func (o *DcimFrontPortTemplatesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front port templates update bad request response a status code equal to that given
func (o *DcimFrontPortTemplatesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimFrontPortTemplatesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimFrontPortTemplatesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcimFrontPortTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimFrontPortTemplatesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesUpdateDefault creates a DcimFrontPortTemplatesUpdateDefault with default headers values
func NewDcimFrontPortTemplatesUpdateDefault(code int) *DcimFrontPortTemplatesUpdateDefault {
	return &DcimFrontPortTemplatesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimFrontPortTemplatesUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimFrontPortTemplatesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front port templates update default response
func (o *DcimFrontPortTemplatesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim front port templates update default response has a 2xx status code
func (o *DcimFrontPortTemplatesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim front port templates update default response has a 3xx status code
func (o *DcimFrontPortTemplatesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim front port templates update default response has a 4xx status code
func (o *DcimFrontPortTemplatesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim front port templates update default response has a 5xx status code
func (o *DcimFrontPortTemplatesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim front port templates update default response a status code equal to that given
func (o *DcimFrontPortTemplatesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimFrontPortTemplatesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcim_front-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortTemplatesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/front-port-templates/{id}/][%d] dcim_front-port-templates_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortTemplatesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
