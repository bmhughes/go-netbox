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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// ExtrasObjectChangesReadReader is a Reader for the ExtrasObjectChangesRead structure.
type ExtrasObjectChangesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasObjectChangesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasObjectChangesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasObjectChangesReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasObjectChangesReadOK creates a ExtrasObjectChangesReadOK with default headers values
func NewExtrasObjectChangesReadOK() *ExtrasObjectChangesReadOK {
	return &ExtrasObjectChangesReadOK{}
}

/*
ExtrasObjectChangesReadOK describes a response with status code 200, with default header values.

ExtrasObjectChangesReadOK extras object changes read o k
*/
type ExtrasObjectChangesReadOK struct {
	Payload *models.ObjectChange
}

// IsSuccess returns true when this extras object changes read o k response has a 2xx status code
func (o *ExtrasObjectChangesReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras object changes read o k response has a 3xx status code
func (o *ExtrasObjectChangesReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras object changes read o k response has a 4xx status code
func (o *ExtrasObjectChangesReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras object changes read o k response has a 5xx status code
func (o *ExtrasObjectChangesReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras object changes read o k response a status code equal to that given
func (o *ExtrasObjectChangesReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasObjectChangesReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extrasObjectChangesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasObjectChangesReadOK) String() string {
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extrasObjectChangesReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasObjectChangesReadOK) GetPayload() *models.ObjectChange {
	return o.Payload
}

func (o *ExtrasObjectChangesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectChange)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasObjectChangesReadBadRequest creates a ExtrasObjectChangesReadBadRequest with default headers values
func NewExtrasObjectChangesReadBadRequest() *ExtrasObjectChangesReadBadRequest {
	return &ExtrasObjectChangesReadBadRequest{}
}

/*
ExtrasObjectChangesReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasObjectChangesReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras object changes read bad request response has a 2xx status code
func (o *ExtrasObjectChangesReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras object changes read bad request response has a 3xx status code
func (o *ExtrasObjectChangesReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras object changes read bad request response has a 4xx status code
func (o *ExtrasObjectChangesReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras object changes read bad request response has a 5xx status code
func (o *ExtrasObjectChangesReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras object changes read bad request response a status code equal to that given
func (o *ExtrasObjectChangesReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasObjectChangesReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extrasObjectChangesReadBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasObjectChangesReadBadRequest) String() string {
	return fmt.Sprintf("[GET /extras/object-changes/{id}/][%d] extrasObjectChangesReadBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasObjectChangesReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasObjectChangesReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
