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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// WirelessWirelessLanGroupsReadReader is a Reader for the WirelessWirelessLanGroupsRead structure.
type WirelessWirelessLanGroupsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLanGroupsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLanGroupsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewWirelessWirelessLanGroupsReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewWirelessWirelessLanGroupsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLanGroupsReadOK creates a WirelessWirelessLanGroupsReadOK with default headers values
func NewWirelessWirelessLanGroupsReadOK() *WirelessWirelessLanGroupsReadOK {
	return &WirelessWirelessLanGroupsReadOK{}
}

/*
WirelessWirelessLanGroupsReadOK describes a response with status code 200, with default header values.

WirelessWirelessLanGroupsReadOK wireless wireless lan groups read o k
*/
type WirelessWirelessLanGroupsReadOK struct {
	Payload *models.WirelessLANGroup
}

// IsSuccess returns true when this wireless wireless lan groups read o k response has a 2xx status code
func (o *WirelessWirelessLanGroupsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lan groups read o k response has a 3xx status code
func (o *WirelessWirelessLanGroupsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lan groups read o k response has a 4xx status code
func (o *WirelessWirelessLanGroupsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lan groups read o k response has a 5xx status code
func (o *WirelessWirelessLanGroupsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lan groups read o k response a status code equal to that given
func (o *WirelessWirelessLanGroupsReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *WirelessWirelessLanGroupsReadOK) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/{id}/][%d] wirelessWirelessLanGroupsReadOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLanGroupsReadOK) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/{id}/][%d] wirelessWirelessLanGroupsReadOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLanGroupsReadOK) GetPayload() *models.WirelessLANGroup {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLanGroupsReadBadRequest creates a WirelessWirelessLanGroupsReadBadRequest with default headers values
func NewWirelessWirelessLanGroupsReadBadRequest() *WirelessWirelessLanGroupsReadBadRequest {
	return &WirelessWirelessLanGroupsReadBadRequest{}
}

/*
WirelessWirelessLanGroupsReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type WirelessWirelessLanGroupsReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this wireless wireless lan groups read bad request response has a 2xx status code
func (o *WirelessWirelessLanGroupsReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this wireless wireless lan groups read bad request response has a 3xx status code
func (o *WirelessWirelessLanGroupsReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lan groups read bad request response has a 4xx status code
func (o *WirelessWirelessLanGroupsReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this wireless wireless lan groups read bad request response has a 5xx status code
func (o *WirelessWirelessLanGroupsReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lan groups read bad request response a status code equal to that given
func (o *WirelessWirelessLanGroupsReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *WirelessWirelessLanGroupsReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/{id}/][%d] wirelessWirelessLanGroupsReadBadRequest  %+v", 400, o.Payload)
}

func (o *WirelessWirelessLanGroupsReadBadRequest) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/{id}/][%d] wirelessWirelessLanGroupsReadBadRequest  %+v", 400, o.Payload)
}

func (o *WirelessWirelessLanGroupsReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLanGroupsReadDefault creates a WirelessWirelessLanGroupsReadDefault with default headers values
func NewWirelessWirelessLanGroupsReadDefault(code int) *WirelessWirelessLanGroupsReadDefault {
	return &WirelessWirelessLanGroupsReadDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLanGroupsReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type WirelessWirelessLanGroupsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the wireless wireless lan groups read default response
func (o *WirelessWirelessLanGroupsReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this wireless wireless lan groups read default response has a 2xx status code
func (o *WirelessWirelessLanGroupsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless lan groups read default response has a 3xx status code
func (o *WirelessWirelessLanGroupsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless lan groups read default response has a 4xx status code
func (o *WirelessWirelessLanGroupsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless lan groups read default response has a 5xx status code
func (o *WirelessWirelessLanGroupsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless lan groups read default response a status code equal to that given
func (o *WirelessWirelessLanGroupsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WirelessWirelessLanGroupsReadDefault) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/{id}/][%d] wireless_wireless-lan-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLanGroupsReadDefault) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-lan-groups/{id}/][%d] wireless_wireless-lan-groups_read default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLanGroupsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLanGroupsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
