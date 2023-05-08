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

// DcimDeviceBayTemplatesUpdateReader is a Reader for the DcimDeviceBayTemplatesUpdate structure.
type DcimDeviceBayTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceBayTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimDeviceBayTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimDeviceBayTemplatesUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimDeviceBayTemplatesUpdateOK creates a DcimDeviceBayTemplatesUpdateOK with default headers values
func NewDcimDeviceBayTemplatesUpdateOK() *DcimDeviceBayTemplatesUpdateOK {
	return &DcimDeviceBayTemplatesUpdateOK{}
}

/*
DcimDeviceBayTemplatesUpdateOK describes a response with status code 200, with default header values.

DcimDeviceBayTemplatesUpdateOK dcim device bay templates update o k
*/
type DcimDeviceBayTemplatesUpdateOK struct {
	Payload *models.DeviceBayTemplate
}

// IsSuccess returns true when this dcim device bay templates update o k response has a 2xx status code
func (o *DcimDeviceBayTemplatesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim device bay templates update o k response has a 3xx status code
func (o *DcimDeviceBayTemplatesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates update o k response has a 4xx status code
func (o *DcimDeviceBayTemplatesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim device bay templates update o k response has a 5xx status code
func (o *DcimDeviceBayTemplatesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates update o k response a status code equal to that given
func (o *DcimDeviceBayTemplatesUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimDeviceBayTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceBayTemplatesUpdateOK) GetPayload() *models.DeviceBayTemplate {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceBayTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimDeviceBayTemplatesUpdateBadRequest creates a DcimDeviceBayTemplatesUpdateBadRequest with default headers values
func NewDcimDeviceBayTemplatesUpdateBadRequest() *DcimDeviceBayTemplatesUpdateBadRequest {
	return &DcimDeviceBayTemplatesUpdateBadRequest{}
}

/*
DcimDeviceBayTemplatesUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimDeviceBayTemplatesUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim device bay templates update bad request response has a 2xx status code
func (o *DcimDeviceBayTemplatesUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim device bay templates update bad request response has a 3xx status code
func (o *DcimDeviceBayTemplatesUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim device bay templates update bad request response has a 4xx status code
func (o *DcimDeviceBayTemplatesUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim device bay templates update bad request response has a 5xx status code
func (o *DcimDeviceBayTemplatesUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim device bay templates update bad request response a status code equal to that given
func (o *DcimDeviceBayTemplatesUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimDeviceBayTemplatesUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDeviceBayTemplatesUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/device-bay-templates/{id}/][%d] dcimDeviceBayTemplatesUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimDeviceBayTemplatesUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimDeviceBayTemplatesUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
