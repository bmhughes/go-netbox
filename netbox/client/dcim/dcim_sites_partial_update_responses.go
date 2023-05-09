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

// DcimSitesPartialUpdateReader is a Reader for the DcimSitesPartialUpdate structure.
type DcimSitesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimSitesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimSitesPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimSitesPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimSitesPartialUpdateOK creates a DcimSitesPartialUpdateOK with default headers values
func NewDcimSitesPartialUpdateOK() *DcimSitesPartialUpdateOK {
	return &DcimSitesPartialUpdateOK{}
}

/*
DcimSitesPartialUpdateOK describes a response with status code 200, with default header values.

DcimSitesPartialUpdateOK dcim sites partial update o k
*/
type DcimSitesPartialUpdateOK struct {
	Payload *models.Site
}

// IsSuccess returns true when this dcim sites partial update o k response has a 2xx status code
func (o *DcimSitesPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim sites partial update o k response has a 3xx status code
func (o *DcimSitesPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites partial update o k response has a 4xx status code
func (o *DcimSitesPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim sites partial update o k response has a 5xx status code
func (o *DcimSitesPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites partial update o k response a status code equal to that given
func (o *DcimSitesPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimSitesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/{id}/][%d] dcimSitesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimSitesPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/sites/{id}/][%d] dcimSitesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimSitesPartialUpdateOK) GetPayload() *models.Site {
	return o.Payload
}

func (o *DcimSitesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Site)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesPartialUpdateBadRequest creates a DcimSitesPartialUpdateBadRequest with default headers values
func NewDcimSitesPartialUpdateBadRequest() *DcimSitesPartialUpdateBadRequest {
	return &DcimSitesPartialUpdateBadRequest{}
}

/*
DcimSitesPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimSitesPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim sites partial update bad request response has a 2xx status code
func (o *DcimSitesPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim sites partial update bad request response has a 3xx status code
func (o *DcimSitesPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites partial update bad request response has a 4xx status code
func (o *DcimSitesPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim sites partial update bad request response has a 5xx status code
func (o *DcimSitesPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites partial update bad request response a status code equal to that given
func (o *DcimSitesPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimSitesPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/{id}/][%d] dcimSitesPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimSitesPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /dcim/sites/{id}/][%d] dcimSitesPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimSitesPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimSitesPartialUpdateDefault creates a DcimSitesPartialUpdateDefault with default headers values
func NewDcimSitesPartialUpdateDefault(code int) *DcimSitesPartialUpdateDefault {
	return &DcimSitesPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimSitesPartialUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimSitesPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim sites partial update default response
func (o *DcimSitesPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim sites partial update default response has a 2xx status code
func (o *DcimSitesPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim sites partial update default response has a 3xx status code
func (o *DcimSitesPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim sites partial update default response has a 4xx status code
func (o *DcimSitesPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim sites partial update default response has a 5xx status code
func (o *DcimSitesPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim sites partial update default response a status code equal to that given
func (o *DcimSitesPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimSitesPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/sites/{id}/][%d] dcim_sites_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/sites/{id}/][%d] dcim_sites_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimSitesPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
