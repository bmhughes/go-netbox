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
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// DcimConsolePortTemplatesListReader is a Reader for the DcimConsolePortTemplatesList structure.
type DcimConsolePortTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimConsolePortTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimConsolePortTemplatesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimConsolePortTemplatesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsolePortTemplatesListOK creates a DcimConsolePortTemplatesListOK with default headers values
func NewDcimConsolePortTemplatesListOK() *DcimConsolePortTemplatesListOK {
	return &DcimConsolePortTemplatesListOK{}
}

/*
DcimConsolePortTemplatesListOK describes a response with status code 200, with default header values.

DcimConsolePortTemplatesListOK dcim console port templates list o k
*/
type DcimConsolePortTemplatesListOK struct {
	Payload *DcimConsolePortTemplatesListOKBody
}

// IsSuccess returns true when this dcim console port templates list o k response has a 2xx status code
func (o *DcimConsolePortTemplatesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console port templates list o k response has a 3xx status code
func (o *DcimConsolePortTemplatesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates list o k response has a 4xx status code
func (o *DcimConsolePortTemplatesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console port templates list o k response has a 5xx status code
func (o *DcimConsolePortTemplatesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates list o k response a status code equal to that given
func (o *DcimConsolePortTemplatesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimConsolePortTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/][%d] dcimConsolePortTemplatesListOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesListOK) String() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/][%d] dcimConsolePortTemplatesListOK  %+v", 200, o.Payload)
}

func (o *DcimConsolePortTemplatesListOK) GetPayload() *DcimConsolePortTemplatesListOKBody {
	return o.Payload
}

func (o *DcimConsolePortTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimConsolePortTemplatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesListBadRequest creates a DcimConsolePortTemplatesListBadRequest with default headers values
func NewDcimConsolePortTemplatesListBadRequest() *DcimConsolePortTemplatesListBadRequest {
	return &DcimConsolePortTemplatesListBadRequest{}
}

/*
DcimConsolePortTemplatesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimConsolePortTemplatesListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim console port templates list bad request response has a 2xx status code
func (o *DcimConsolePortTemplatesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim console port templates list bad request response has a 3xx status code
func (o *DcimConsolePortTemplatesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console port templates list bad request response has a 4xx status code
func (o *DcimConsolePortTemplatesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim console port templates list bad request response has a 5xx status code
func (o *DcimConsolePortTemplatesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console port templates list bad request response a status code equal to that given
func (o *DcimConsolePortTemplatesListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimConsolePortTemplatesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/][%d] dcimConsolePortTemplatesListBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsolePortTemplatesListBadRequest) String() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/][%d] dcimConsolePortTemplatesListBadRequest  %+v", 400, o.Payload)
}

func (o *DcimConsolePortTemplatesListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimConsolePortTemplatesListDefault creates a DcimConsolePortTemplatesListDefault with default headers values
func NewDcimConsolePortTemplatesListDefault(code int) *DcimConsolePortTemplatesListDefault {
	return &DcimConsolePortTemplatesListDefault{
		_statusCode: code,
	}
}

/*
DcimConsolePortTemplatesListDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimConsolePortTemplatesListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console port templates list default response
func (o *DcimConsolePortTemplatesListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console port templates list default response has a 2xx status code
func (o *DcimConsolePortTemplatesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console port templates list default response has a 3xx status code
func (o *DcimConsolePortTemplatesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console port templates list default response has a 4xx status code
func (o *DcimConsolePortTemplatesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console port templates list default response has a 5xx status code
func (o *DcimConsolePortTemplatesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console port templates list default response a status code equal to that given
func (o *DcimConsolePortTemplatesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsolePortTemplatesListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/][%d] dcim_console-port-templates_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/console-port-templates/][%d] dcim_console-port-templates_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsolePortTemplatesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsolePortTemplatesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimConsolePortTemplatesListOKBody dcim console port templates list o k body
swagger:model DcimConsolePortTemplatesListOKBody
*/
type DcimConsolePortTemplatesListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.ConsolePortTemplate `json:"results"`
}

// Validate validates this dcim console port templates list o k body
func (o *DcimConsolePortTemplatesListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimConsolePortTemplatesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimConsolePortTemplatesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimConsolePortTemplatesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimConsolePortTemplatesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimConsolePortTemplatesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimConsolePortTemplatesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimConsolePortTemplatesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimConsolePortTemplatesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimConsolePortTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimConsolePortTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim console port templates list o k body based on the context it is used
func (o *DcimConsolePortTemplatesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimConsolePortTemplatesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimConsolePortTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimConsolePortTemplatesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimConsolePortTemplatesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimConsolePortTemplatesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimConsolePortTemplatesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
