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

// DcimManufacturersReadReader is a Reader for the DcimManufacturersRead structure.
type DcimManufacturersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimManufacturersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimManufacturersReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimManufacturersReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimManufacturersReadOK creates a DcimManufacturersReadOK with default headers values
func NewDcimManufacturersReadOK() *DcimManufacturersReadOK {
	return &DcimManufacturersReadOK{}
}

/*
DcimManufacturersReadOK describes a response with status code 200, with default header values.

DcimManufacturersReadOK dcim manufacturers read o k
*/
type DcimManufacturersReadOK struct {
	Payload *models.Manufacturer
}

// IsSuccess returns true when this dcim manufacturers read o k response has a 2xx status code
func (o *DcimManufacturersReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim manufacturers read o k response has a 3xx status code
func (o *DcimManufacturersReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim manufacturers read o k response has a 4xx status code
func (o *DcimManufacturersReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim manufacturers read o k response has a 5xx status code
func (o *DcimManufacturersReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim manufacturers read o k response a status code equal to that given
func (o *DcimManufacturersReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimManufacturersReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/{id}/][%d] dcimManufacturersReadOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersReadOK) String() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/{id}/][%d] dcimManufacturersReadOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersReadOK) GetPayload() *models.Manufacturer {
	return o.Payload
}

func (o *DcimManufacturersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersReadBadRequest creates a DcimManufacturersReadBadRequest with default headers values
func NewDcimManufacturersReadBadRequest() *DcimManufacturersReadBadRequest {
	return &DcimManufacturersReadBadRequest{}
}

/*
DcimManufacturersReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimManufacturersReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim manufacturers read bad request response has a 2xx status code
func (o *DcimManufacturersReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim manufacturers read bad request response has a 3xx status code
func (o *DcimManufacturersReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim manufacturers read bad request response has a 4xx status code
func (o *DcimManufacturersReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim manufacturers read bad request response has a 5xx status code
func (o *DcimManufacturersReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim manufacturers read bad request response a status code equal to that given
func (o *DcimManufacturersReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimManufacturersReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/{id}/][%d] dcimManufacturersReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimManufacturersReadBadRequest) String() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/{id}/][%d] dcimManufacturersReadBadRequest  %+v", 400, o.Payload)
}

func (o *DcimManufacturersReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimManufacturersReadDefault creates a DcimManufacturersReadDefault with default headers values
func NewDcimManufacturersReadDefault(code int) *DcimManufacturersReadDefault {
	return &DcimManufacturersReadDefault{
		_statusCode: code,
	}
}

/*
DcimManufacturersReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimManufacturersReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim manufacturers read default response
func (o *DcimManufacturersReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim manufacturers read default response has a 2xx status code
func (o *DcimManufacturersReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim manufacturers read default response has a 3xx status code
func (o *DcimManufacturersReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim manufacturers read default response has a 4xx status code
func (o *DcimManufacturersReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim manufacturers read default response has a 5xx status code
func (o *DcimManufacturersReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim manufacturers read default response a status code equal to that given
func (o *DcimManufacturersReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimManufacturersReadDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/{id}/][%d] dcim_manufacturers_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimManufacturersReadDefault) String() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/{id}/][%d] dcim_manufacturers_read default  %+v", o._statusCode, o.Payload)
}

func (o *DcimManufacturersReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimManufacturersReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
