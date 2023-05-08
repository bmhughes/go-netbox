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

// DcimPlatformsCreateReader is a Reader for the DcimPlatformsCreate structure.
type DcimPlatformsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewDcimPlatformsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimPlatformsCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimPlatformsCreateCreated creates a DcimPlatformsCreateCreated with default headers values
func NewDcimPlatformsCreateCreated() *DcimPlatformsCreateCreated {
	return &DcimPlatformsCreateCreated{}
}

/*
DcimPlatformsCreateCreated describes a response with status code 201, with default header values.

DcimPlatformsCreateCreated dcim platforms create created
*/
type DcimPlatformsCreateCreated struct {
	Payload *models.Platform
}

// IsSuccess returns true when this dcim platforms create created response has a 2xx status code
func (o *DcimPlatformsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim platforms create created response has a 3xx status code
func (o *DcimPlatformsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms create created response has a 4xx status code
func (o *DcimPlatformsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim platforms create created response has a 5xx status code
func (o *DcimPlatformsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms create created response a status code equal to that given
func (o *DcimPlatformsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *DcimPlatformsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/platforms/][%d] dcimPlatformsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPlatformsCreateCreated) String() string {
	return fmt.Sprintf("[POST /dcim/platforms/][%d] dcimPlatformsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPlatformsCreateCreated) GetPayload() *models.Platform {
	return o.Payload
}

func (o *DcimPlatformsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimPlatformsCreateBadRequest creates a DcimPlatformsCreateBadRequest with default headers values
func NewDcimPlatformsCreateBadRequest() *DcimPlatformsCreateBadRequest {
	return &DcimPlatformsCreateBadRequest{}
}

/*
DcimPlatformsCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimPlatformsCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim platforms create bad request response has a 2xx status code
func (o *DcimPlatformsCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim platforms create bad request response has a 3xx status code
func (o *DcimPlatformsCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim platforms create bad request response has a 4xx status code
func (o *DcimPlatformsCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim platforms create bad request response has a 5xx status code
func (o *DcimPlatformsCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim platforms create bad request response a status code equal to that given
func (o *DcimPlatformsCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimPlatformsCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /dcim/platforms/][%d] dcimPlatformsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPlatformsCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /dcim/platforms/][%d] dcimPlatformsCreateBadRequest  %+v", 400, o.Payload)
}

func (o *DcimPlatformsCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimPlatformsCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
