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

// ExtrasJournalEntriesBulkUpdateReader is a Reader for the ExtrasJournalEntriesBulkUpdate structure.
type ExtrasJournalEntriesBulkUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasJournalEntriesBulkUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasJournalEntriesBulkUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewExtrasJournalEntriesBulkUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewExtrasJournalEntriesBulkUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasJournalEntriesBulkUpdateOK creates a ExtrasJournalEntriesBulkUpdateOK with default headers values
func NewExtrasJournalEntriesBulkUpdateOK() *ExtrasJournalEntriesBulkUpdateOK {
	return &ExtrasJournalEntriesBulkUpdateOK{}
}

/*
ExtrasJournalEntriesBulkUpdateOK describes a response with status code 200, with default header values.

ExtrasJournalEntriesBulkUpdateOK extras journal entries bulk update o k
*/
type ExtrasJournalEntriesBulkUpdateOK struct {
	Payload *models.JournalEntry
}

// IsSuccess returns true when this extras journal entries bulk update o k response has a 2xx status code
func (o *ExtrasJournalEntriesBulkUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras journal entries bulk update o k response has a 3xx status code
func (o *ExtrasJournalEntriesBulkUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras journal entries bulk update o k response has a 4xx status code
func (o *ExtrasJournalEntriesBulkUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras journal entries bulk update o k response has a 5xx status code
func (o *ExtrasJournalEntriesBulkUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras journal entries bulk update o k response a status code equal to that given
func (o *ExtrasJournalEntriesBulkUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *ExtrasJournalEntriesBulkUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/journal-entries/][%d] extrasJournalEntriesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasJournalEntriesBulkUpdateOK) String() string {
	return fmt.Sprintf("[PUT /extras/journal-entries/][%d] extrasJournalEntriesBulkUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasJournalEntriesBulkUpdateOK) GetPayload() *models.JournalEntry {
	return o.Payload
}

func (o *ExtrasJournalEntriesBulkUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JournalEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasJournalEntriesBulkUpdateBadRequest creates a ExtrasJournalEntriesBulkUpdateBadRequest with default headers values
func NewExtrasJournalEntriesBulkUpdateBadRequest() *ExtrasJournalEntriesBulkUpdateBadRequest {
	return &ExtrasJournalEntriesBulkUpdateBadRequest{}
}

/*
ExtrasJournalEntriesBulkUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type ExtrasJournalEntriesBulkUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this extras journal entries bulk update bad request response has a 2xx status code
func (o *ExtrasJournalEntriesBulkUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this extras journal entries bulk update bad request response has a 3xx status code
func (o *ExtrasJournalEntriesBulkUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras journal entries bulk update bad request response has a 4xx status code
func (o *ExtrasJournalEntriesBulkUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this extras journal entries bulk update bad request response has a 5xx status code
func (o *ExtrasJournalEntriesBulkUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this extras journal entries bulk update bad request response a status code equal to that given
func (o *ExtrasJournalEntriesBulkUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *ExtrasJournalEntriesBulkUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /extras/journal-entries/][%d] extrasJournalEntriesBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasJournalEntriesBulkUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /extras/journal-entries/][%d] extrasJournalEntriesBulkUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *ExtrasJournalEntriesBulkUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasJournalEntriesBulkUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasJournalEntriesBulkUpdateDefault creates a ExtrasJournalEntriesBulkUpdateDefault with default headers values
func NewExtrasJournalEntriesBulkUpdateDefault(code int) *ExtrasJournalEntriesBulkUpdateDefault {
	return &ExtrasJournalEntriesBulkUpdateDefault{
		_statusCode: code,
	}
}

/*
ExtrasJournalEntriesBulkUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type ExtrasJournalEntriesBulkUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the extras journal entries bulk update default response
func (o *ExtrasJournalEntriesBulkUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this extras journal entries bulk update default response has a 2xx status code
func (o *ExtrasJournalEntriesBulkUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras journal entries bulk update default response has a 3xx status code
func (o *ExtrasJournalEntriesBulkUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras journal entries bulk update default response has a 4xx status code
func (o *ExtrasJournalEntriesBulkUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras journal entries bulk update default response has a 5xx status code
func (o *ExtrasJournalEntriesBulkUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras journal entries bulk update default response a status code equal to that given
func (o *ExtrasJournalEntriesBulkUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ExtrasJournalEntriesBulkUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /extras/journal-entries/][%d] extras_journal-entries_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasJournalEntriesBulkUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /extras/journal-entries/][%d] extras_journal-entries_bulk_update default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasJournalEntriesBulkUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasJournalEntriesBulkUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
