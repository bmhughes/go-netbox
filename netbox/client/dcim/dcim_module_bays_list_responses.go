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

// DcimModuleBaysListReader is a Reader for the DcimModuleBaysList structure.
type DcimModuleBaysListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBaysListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimModuleBaysListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimModuleBaysListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimModuleBaysListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBaysListOK creates a DcimModuleBaysListOK with default headers values
func NewDcimModuleBaysListOK() *DcimModuleBaysListOK {
	return &DcimModuleBaysListOK{}
}

/*
DcimModuleBaysListOK describes a response with status code 200, with default header values.

DcimModuleBaysListOK dcim module bays list o k
*/
type DcimModuleBaysListOK struct {
	Payload *DcimModuleBaysListOKBody
}

// IsSuccess returns true when this dcim module bays list o k response has a 2xx status code
func (o *DcimModuleBaysListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bays list o k response has a 3xx status code
func (o *DcimModuleBaysListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays list o k response has a 4xx status code
func (o *DcimModuleBaysListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bays list o k response has a 5xx status code
func (o *DcimModuleBaysListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays list o k response a status code equal to that given
func (o *DcimModuleBaysListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimModuleBaysListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/module-bays/][%d] dcimModuleBaysListOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysListOK) String() string {
	return fmt.Sprintf("[GET /dcim/module-bays/][%d] dcimModuleBaysListOK  %+v", 200, o.Payload)
}

func (o *DcimModuleBaysListOK) GetPayload() *DcimModuleBaysListOKBody {
	return o.Payload
}

func (o *DcimModuleBaysListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimModuleBaysListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBaysListBadRequest creates a DcimModuleBaysListBadRequest with default headers values
func NewDcimModuleBaysListBadRequest() *DcimModuleBaysListBadRequest {
	return &DcimModuleBaysListBadRequest{}
}

/*
DcimModuleBaysListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimModuleBaysListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim module bays list bad request response has a 2xx status code
func (o *DcimModuleBaysListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim module bays list bad request response has a 3xx status code
func (o *DcimModuleBaysListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bays list bad request response has a 4xx status code
func (o *DcimModuleBaysListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim module bays list bad request response has a 5xx status code
func (o *DcimModuleBaysListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bays list bad request response a status code equal to that given
func (o *DcimModuleBaysListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimModuleBaysListBadRequest) Error() string {
	return fmt.Sprintf("[GET /dcim/module-bays/][%d] dcimModuleBaysListBadRequest  %+v", 400, o.Payload)
}

func (o *DcimModuleBaysListBadRequest) String() string {
	return fmt.Sprintf("[GET /dcim/module-bays/][%d] dcimModuleBaysListBadRequest  %+v", 400, o.Payload)
}

func (o *DcimModuleBaysListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimModuleBaysListDefault creates a DcimModuleBaysListDefault with default headers values
func NewDcimModuleBaysListDefault(code int) *DcimModuleBaysListDefault {
	return &DcimModuleBaysListDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBaysListDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimModuleBaysListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim module bays list default response
func (o *DcimModuleBaysListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim module bays list default response has a 2xx status code
func (o *DcimModuleBaysListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bays list default response has a 3xx status code
func (o *DcimModuleBaysListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bays list default response has a 4xx status code
func (o *DcimModuleBaysListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bays list default response has a 5xx status code
func (o *DcimModuleBaysListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bays list default response a status code equal to that given
func (o *DcimModuleBaysListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimModuleBaysListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/module-bays/][%d] dcim_module-bays_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/module-bays/][%d] dcim_module-bays_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBaysListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBaysListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimModuleBaysListOKBody dcim module bays list o k body
swagger:model DcimModuleBaysListOKBody
*/
type DcimModuleBaysListOKBody struct {

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
	Results []*models.ModuleBay `json:"results"`
}

// Validate validates this dcim module bays list o k body
func (o *DcimModuleBaysListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimModuleBaysListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimModuleBaysListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimModuleBaysListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimModuleBaysListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimModuleBaysListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimModuleBaysListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimModuleBaysListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimModuleBaysListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimModuleBaysListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimModuleBaysListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim module bays list o k body based on the context it is used
func (o *DcimModuleBaysListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimModuleBaysListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimModuleBaysListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimModuleBaysListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimModuleBaysListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimModuleBaysListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimModuleBaysListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
