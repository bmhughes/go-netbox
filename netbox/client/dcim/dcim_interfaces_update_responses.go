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

// DcimInterfacesUpdateReader is a Reader for the DcimInterfacesUpdate structure.
type DcimInterfacesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInterfacesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimInterfacesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimInterfacesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInterfacesUpdateOK creates a DcimInterfacesUpdateOK with default headers values
func NewDcimInterfacesUpdateOK() *DcimInterfacesUpdateOK {
	return &DcimInterfacesUpdateOK{}
}

/*
DcimInterfacesUpdateOK describes a response with status code 200, with default header values.

DcimInterfacesUpdateOK dcim interfaces update o k
*/
type DcimInterfacesUpdateOK struct {
	Payload *models.Interface
}

// IsSuccess returns true when this dcim interfaces update o k response has a 2xx status code
func (o *DcimInterfacesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim interfaces update o k response has a 3xx status code
func (o *DcimInterfacesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interfaces update o k response has a 4xx status code
func (o *DcimInterfacesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim interfaces update o k response has a 5xx status code
func (o *DcimInterfacesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interfaces update o k response a status code equal to that given
func (o *DcimInterfacesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimInterfacesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/interfaces/{id}/][%d] dcimInterfacesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/interfaces/{id}/][%d] dcimInterfacesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimInterfacesUpdateOK) GetPayload() *models.Interface {
	return o.Payload
}

func (o *DcimInterfacesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Interface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfacesUpdateBadRequest creates a DcimInterfacesUpdateBadRequest with default headers values
func NewDcimInterfacesUpdateBadRequest() *DcimInterfacesUpdateBadRequest {
	return &DcimInterfacesUpdateBadRequest{}
}

/*
DcimInterfacesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimInterfacesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim interfaces update bad request response has a 2xx status code
func (o *DcimInterfacesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim interfaces update bad request response has a 3xx status code
func (o *DcimInterfacesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim interfaces update bad request response has a 4xx status code
func (o *DcimInterfacesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim interfaces update bad request response has a 5xx status code
func (o *DcimInterfacesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim interfaces update bad request response a status code equal to that given
func (o *DcimInterfacesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimInterfacesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/interfaces/{id}/][%d] dcimInterfacesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInterfacesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/interfaces/{id}/][%d] dcimInterfacesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimInterfacesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfacesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInterfacesUpdateDefault creates a DcimInterfacesUpdateDefault with default headers values
func NewDcimInterfacesUpdateDefault(code int) *DcimInterfacesUpdateDefault {
	return &DcimInterfacesUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimInterfacesUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimInterfacesUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim interfaces update default response
func (o *DcimInterfacesUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim interfaces update default response has a 2xx status code
func (o *DcimInterfacesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim interfaces update default response has a 3xx status code
func (o *DcimInterfacesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim interfaces update default response has a 4xx status code
func (o *DcimInterfacesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim interfaces update default response has a 5xx status code
func (o *DcimInterfacesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim interfaces update default response a status code equal to that given
func (o *DcimInterfacesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimInterfacesUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /dcim/interfaces/{id}/][%d] dcim_interfaces_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /dcim/interfaces/{id}/][%d] dcim_interfaces_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimInterfacesUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInterfacesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
