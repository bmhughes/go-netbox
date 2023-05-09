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

// ExtrasSavedFiltersUpdateReader is a Reader for the ExtrasSavedFiltersUpdate structure.
type ExtrasSavedFiltersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasSavedFiltersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasSavedFiltersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasSavedFiltersUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasSavedFiltersUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasSavedFiltersUpdateOK creates a ExtrasSavedFiltersUpdateOK with default headers values
func NewExtrasSavedFiltersUpdateOK() *ExtrasSavedFiltersUpdateOK {
	return &ExtrasSavedFiltersUpdateOK{}
}

/*
ExtrasSavedFiltersUpdateOK describes a response with status code 200, with default header values.

ExtrasSavedFiltersUpdateOK extras saved filters update o k
*/
type ExtrasSavedFiltersUpdateOK struct {
	Payload *models.SavedFilter
}

// IsSuccess returns true when this extras saved filters update o k response has a 2xx status code
func (o *ExtrasSavedFiltersUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras saved filters update o k response has a 3xx status code
func (o *ExtrasSavedFiltersUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras saved filters update o k response has a 4xx status code
func (o *ExtrasSavedFiltersUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras saved filters update o k response has a 5xx status code
func (o *ExtrasSavedFiltersUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras saved filters update o k response a status code equal to that given
func (o *ExtrasSavedFiltersUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasSavedFiltersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/saved-filters/{id}/][%d] extrasSavedFiltersUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasSavedFiltersUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/saved-filters/{id}/][%d] extrasSavedFiltersUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasSavedFiltersUpdateOK) GetPayload() *models.SavedFilter {
	return o.Payload
}

func (o *ExtrasSavedFiltersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SavedFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasSavedFiltersUpdateBadRequest creates a ExtrasSavedFiltersUpdateBadRequest with default headers values
func NewExtrasSavedFiltersUpdateBadRequest() *ExtrasSavedFiltersUpdateBadRequest {
	return &ExtrasSavedFiltersUpdateBadRequest{}
}

/*
ExtrasSavedFiltersUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasSavedFiltersUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras saved filters update bad request response has a 2xx status code
func (o *ExtrasSavedFiltersUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras saved filters update bad request response has a 3xx status code
func (o *ExtrasSavedFiltersUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras saved filters update bad request response has a 4xx status code
func (o *ExtrasSavedFiltersUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras saved filters update bad request response has a 5xx status code
func (o *ExtrasSavedFiltersUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras saved filters update bad request response a status code equal to that given
func (o *ExtrasSavedFiltersUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasSavedFiltersUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /extras/saved-filters/{id}/][%d] extrasSavedFiltersUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasSavedFiltersUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /extras/saved-filters/{id}/][%d] extrasSavedFiltersUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasSavedFiltersUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasSavedFiltersUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasSavedFiltersUpdateDefault creates a ExtrasSavedFiltersUpdateDefault with default headers values
func NewExtrasSavedFiltersUpdateDefault(code int) *ExtrasSavedFiltersUpdateDefault {
	return &ExtrasSavedFiltersUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasSavedFiltersUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasSavedFiltersUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras saved filters update default response
func (o *ExtrasSavedFiltersUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras saved filters update default response has a 2xx status code
func (o *ExtrasSavedFiltersUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras saved filters update default response has a 3xx status code
func (o *ExtrasSavedFiltersUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras saved filters update default response has a 4xx status code
func (o *ExtrasSavedFiltersUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras saved filters update default response has a 5xx status code
func (o *ExtrasSavedFiltersUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras saved filters update default response a status code equal to that given
func (o *ExtrasSavedFiltersUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasSavedFiltersUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/saved-filters/{id}/][%d] extras_saved-filters_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasSavedFiltersUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/saved-filters/{id}/][%d] extras_saved-filters_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasSavedFiltersUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasSavedFiltersUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
