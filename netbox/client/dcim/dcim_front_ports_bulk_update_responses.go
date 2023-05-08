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

// DcimFrontPortsBulkUpdateReader is a Reader for the DcimFrontPortsBulkUpdate structure.
type DcimFrontPortsBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimFrontPortsBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimFrontPortsBulkUpdateOK creates a DcimFrontPortsBulkUpdateOK with default headers values
func NewDcimFrontPortsBulkUpdateOK() *DcimFrontPortsBulkUpdateOK {
	return &DcimFrontPortsBulkUpdateOK{}
}

/*
DcimFrontPortsBulkUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortsBulkUpdateOK dcim front ports bulk update o k
*/
type DcimFrontPortsBulkUpdateOK struct {
	Payload *models.FrontPort
}

// IsSuccess returns true when this dcim front ports bulk update o k response has a 2xx status code
func (o *DcimFrontPortsBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front ports bulk update o k response has a 3xx status code
func (o *DcimFrontPortsBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front ports bulk update o k response has a 4xx status code
func (o *DcimFrontPortsBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front ports bulk update o k response has a 5xx status code
func (o *DcimFrontPortsBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front ports bulk update o k response a status code equal to that given
func (o *DcimFrontPortsBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimFrontPortsBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/][%d] dcimFrontPortsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortsBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/][%d] dcimFrontPortsBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortsBulkUpdateOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsBulkUpdateBadRequest creates a DcimFrontPortsBulkUpdateBadRequest with default headers values
func NewDcimFrontPortsBulkUpdateBadRequest() *DcimFrontPortsBulkUpdateBadRequest {
	return &DcimFrontPortsBulkUpdateBadRequest{}
}

/*
DcimFrontPortsBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimFrontPortsBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim front ports bulk update bad request response has a 2xx status code
func (o *DcimFrontPortsBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim front ports bulk update bad request response has a 3xx status code
func (o *DcimFrontPortsBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front ports bulk update bad request response has a 4xx status code
func (o *DcimFrontPortsBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim front ports bulk update bad request response has a 5xx status code
func (o *DcimFrontPortsBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front ports bulk update bad request response a status code equal to that given
func (o *DcimFrontPortsBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimFrontPortsBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/][%d] dcimFrontPortsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimFrontPortsBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /dcim/front-ports/][%d] dcimFrontPortsBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimFrontPortsBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
