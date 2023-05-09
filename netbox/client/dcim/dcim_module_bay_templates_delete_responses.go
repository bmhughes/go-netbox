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
)

// DcimModuleBayTemplatesDeleteReader is a Reader for the DcimModuleBayTemplatesDelete structure.
type DcimModuleBayTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBayTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimModuleBayTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimModuleBayTemplatesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimModuleBayTemplatesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBayTemplatesDeleteNoContent creates a DcimModuleBayTemplatesDeleteNoContent with default headers values
func NewDcimModuleBayTemplatesDeleteNoContent() *DcimModuleBayTemplatesDeleteNoContent {
	return &DcimModuleBayTemplatesDeleteNoContent{}
}

/*
DcimModuleBayTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimModuleBayTemplatesDeleteNoContent dcim module bay templates delete no content
*/
type DcimModuleBayTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim module bay templates delete no content response has a 2xx status code
func (o *DcimModuleBayTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bay templates delete no content response has a 3xx status code
func (o *DcimModuleBayTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bay templates delete no content response has a 4xx status code
func (o *DcimModuleBayTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bay templates delete no content response has a 5xx status code
func (o *DcimModuleBayTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bay templates delete no content response a status code equal to that given
func (o *DcimModuleBayTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimModuleBayTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesDeleteNoContent ", 204)
}

func (o *DcimModuleBayTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesDeleteNoContent ", 204)
}

func (o *DcimModuleBayTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimModuleBayTemplatesDeleteBadRequest creates a DcimModuleBayTemplatesDeleteBadRequest with default headers values
func NewDcimModuleBayTemplatesDeleteBadRequest() *DcimModuleBayTemplatesDeleteBadRequest {
	return &DcimModuleBayTemplatesDeleteBadRequest{}
}

/*
DcimModuleBayTemplatesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimModuleBayTemplatesDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim module bay templates delete bad request response has a 2xx status code
func (o *DcimModuleBayTemplatesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim module bay templates delete bad request response has a 3xx status code
func (o *DcimModuleBayTemplatesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bay templates delete bad request response has a 4xx status code
func (o *DcimModuleBayTemplatesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim module bay templates delete bad request response has a 5xx status code
func (o *DcimModuleBayTemplatesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bay templates delete bad request response a status code equal to that given
func (o *DcimModuleBayTemplatesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimModuleBayTemplatesDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimModuleBayTemplatesDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimModuleBayTemplatesDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBayTemplatesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBayTemplatesDeleteDefault creates a DcimModuleBayTemplatesDeleteDefault with default headers values
func NewDcimModuleBayTemplatesDeleteDefault(code int) *DcimModuleBayTemplatesDeleteDefault {
	return &DcimModuleBayTemplatesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBayTemplatesDeleteDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimModuleBayTemplatesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bay templates delete default response
func (o *DcimModuleBayTemplatesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bay templates delete default response has a 2xx status code
func (o *DcimModuleBayTemplatesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bay templates delete default response has a 3xx status code
func (o *DcimModuleBayTemplatesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bay templates delete default response has a 4xx status code
func (o *DcimModuleBayTemplatesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bay templates delete default response has a 5xx status code
func (o *DcimModuleBayTemplatesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bay templates delete default response a status code equal to that given
func (o *DcimModuleBayTemplatesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBayTemplatesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/{id}/][%d] dcim_module-bay-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/{id}/][%d] dcim_module-bay-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBayTemplatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
