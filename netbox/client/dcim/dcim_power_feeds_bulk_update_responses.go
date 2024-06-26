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

// DcimPowerFeedsBulkUpdateReader is a Reader for the DcimPowerFeedsBulkUpdate structure.
type DcimPowerFeedsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerFeedsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimPowerFeedsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimPowerFeedsBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimPowerFeedsBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimPowerFeedsBulkUpdateOK creates a DcimPowerFeedsBulkUpdateOK with default headers values
func NewDcimPowerFeedsBulkUpdateOK() *DcimPowerFeedsBulkUpdateOK {
	return &DcimPowerFeedsBulkUpdateOK{}
}

/*
DcimPowerFeedsBulkUpdateOK describes a response with status code 200, with default header values.

DcimPowerFeedsBulkUpdateOK dcim power feeds bulk update o k
*/
type DcimPowerFeedsBulkUpdateOK struct {
	Payload *models.PowerFeed
}

// IsSuccess returns true when this dcim power feeds bulk update o k response has a 2xx status code
func (o *DcimPowerFeedsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim power feeds bulk update o k response has a 3xx status code
func (o *DcimPowerFeedsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power feeds bulk update o k response has a 4xx status code
func (o *DcimPowerFeedsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim power feeds bulk update o k response has a 5xx status code
func (o *DcimPowerFeedsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power feeds bulk update o k response a status code equal to that given
func (o *DcimPowerFeedsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimPowerFeedsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcimPowerFeedsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcimPowerFeedsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPowerFeedsBulkUpdateOK) GetPayload() *models.PowerFeed {
	return o.Payload
}

func (o *DcimPowerFeedsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PowerFeed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsBulkUpdateBadRequest creates a DcimPowerFeedsBulkUpdateBadRequest with default headers values
func NewDcimPowerFeedsBulkUpdateBadRequest() *DcimPowerFeedsBulkUpdateBadRequest {
	return &DcimPowerFeedsBulkUpdateBadRequest{}
}

/*
DcimPowerFeedsBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimPowerFeedsBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim power feeds bulk update bad request response has a 2xx status code
func (o *DcimPowerFeedsBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim power feeds bulk update bad request response has a 3xx status code
func (o *DcimPowerFeedsBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim power feeds bulk update bad request response has a 4xx status code
func (o *DcimPowerFeedsBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim power feeds bulk update bad request response has a 5xx status code
func (o *DcimPowerFeedsBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim power feeds bulk update bad request response a status code equal to that given
func (o *DcimPowerFeedsBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimPowerFeedsBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcimPowerFeedsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerFeedsBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcimPowerFeedsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPowerFeedsBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPowerFeedsBulkUpdateDefault creates a DcimPowerFeedsBulkUpdateDefault with default headers values
func NewDcimPowerFeedsBulkUpdateDefault(code int) *DcimPowerFeedsBulkUpdateDefault {
	return &DcimPowerFeedsBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimPowerFeedsBulkUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimPowerFeedsBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim power feeds bulk update default response
func (o *DcimPowerFeedsBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim power feeds bulk update default response has a 2xx status code
func (o *DcimPowerFeedsBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim power feeds bulk update default response has a 3xx status code
func (o *DcimPowerFeedsBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim power feeds bulk update default response has a 4xx status code
func (o *DcimPowerFeedsBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim power feeds bulk update default response has a 5xx status code
func (o *DcimPowerFeedsBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim power feeds bulk update default response a status code equal to that given
func (o *DcimPowerFeedsBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimPowerFeedsBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcim_power-feeds_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/power-feeds/][%d] dcim_power-feeds_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimPowerFeedsBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPowerFeedsBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
