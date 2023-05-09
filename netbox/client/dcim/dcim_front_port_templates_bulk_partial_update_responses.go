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

// DcimFrontPortTemplatesBulkPartialUpdateReader is a Reader for the DcimFrontPortTemplatesBulkPartialUpdate structure.
type DcimFrontPortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimFrontPortTemplatesBulkPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimFrontPortTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortTemplatesBulkPartialUpdateOK creates a DcimFrontPortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimFrontPortTemplatesBulkPartialUpdateOK() *DcimFrontPortTemplatesBulkPartialUpdateOK {
	return &DcimFrontPortTemplatesBulkPartialUpdateOK{}
}

/*
DcimFrontPortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortTemplatesBulkPartialUpdateOK dcim front port templates bulk partial update o k
*/
type DcimFrontPortTemplatesBulkPartialUpdateOK struct {
	Payload *models.FrontPortTemplate
}

// IsSuccess returns true when this dcim front port templates bulk partial update o k response has a 2xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front port templates bulk partial update o k response has a 3xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front port templates bulk partial update o k response has a 4xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front port templates bulk partial update o k response has a 5xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front port templates bulk partial update o k response a status code equal to that given
func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/][%d] dcimFrontPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/][%d] dcimFrontPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) GetPayload() *models.FrontPortTemplate {
	return o.Payload
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesBulkPartialUpdateBadRequest creates a DcimFrontPortTemplatesBulkPartialUpdateBadRequest with default headers values
func NewDcimFrontPortTemplatesBulkPartialUpdateBadRequest() *DcimFrontPortTemplatesBulkPartialUpdateBadRequest {
	return &DcimFrontPortTemplatesBulkPartialUpdateBadRequest{}
}

/*
DcimFrontPortTemplatesBulkPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimFrontPortTemplatesBulkPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim front port templates bulk partial update bad request response has a 2xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim front port templates bulk partial update bad request response has a 3xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front port templates bulk partial update bad request response has a 4xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim front port templates bulk partial update bad request response has a 5xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front port templates bulk partial update bad request response a status code equal to that given
func (o *DcimFrontPortTemplatesBulkPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/][%d] dcimFrontPortTemplatesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/][%d] dcimFrontPortTemplatesBulkPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortTemplatesBulkPartialUpdateDefault creates a DcimFrontPortTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimFrontPortTemplatesBulkPartialUpdateDefault(code int) *DcimFrontPortTemplatesBulkPartialUpdateDefault {
	return &DcimFrontPortTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimFrontPortTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimFrontPortTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front port templates bulk partial update default response
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim front port templates bulk partial update default response has a 2xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim front port templates bulk partial update default response has a 3xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim front port templates bulk partial update default response has a 4xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim front port templates bulk partial update default response has a 5xx status code
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim front port templates bulk partial update default response a status code equal to that given
func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/][%d] dcim_front-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-port-templates/][%d] dcim_front-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
