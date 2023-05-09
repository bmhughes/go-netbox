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

// DcimCablesListReader is a Reader for the DcimCablesList structure.
type DcimCablesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimCablesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimCablesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimCablesListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewDcimCablesListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimCablesListOK creates a DcimCablesListOK with default headers values
func NewDcimCablesListOK() *DcimCablesListOK {
	return &DcimCablesListOK{}
}

/*
DcimCablesListOK describes a response with status code 200, with default header values.

DcimCablesListOK dcim cables list o k
*/
type DcimCablesListOK struct {
	Payload *DcimCablesListOKBody
}

// IsSuccess returns true when this dcim cables list o k response has a 2xx status code
func (o *DcimCablesListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim cables list o k response has a 3xx status code
func (o *DcimCablesListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cables list o k response has a 4xx status code
func (o *DcimCablesListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim cables list o k response has a 5xx status code
func (o *DcimCablesListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cables list o k response a status code equal to that given
func (o *DcimCablesListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimCablesListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/cables/][%d] dcimCablesListOK  %+v", 200, o.Payload)
}

func (o *DcimCablesListOK) String() string {
	return fmt.Sprintf("[GET /dcim/cables/][%d] dcimCablesListOK  %+v", 200, o.Payload)
}

func (o *DcimCablesListOK) GetPayload() *DcimCablesListOKBody {
	return o.Payload
}

func (o *DcimCablesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimCablesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCablesListBadRequest creates a DcimCablesListBadRequest with default headers values
func NewDcimCablesListBadRequest() *DcimCablesListBadRequest {
	return &DcimCablesListBadRequest{}
}

/*
DcimCablesListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimCablesListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim cables list bad request response has a 2xx status code
func (o *DcimCablesListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim cables list bad request response has a 3xx status code
func (o *DcimCablesListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim cables list bad request response has a 4xx status code
func (o *DcimCablesListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim cables list bad request response has a 5xx status code
func (o *DcimCablesListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim cables list bad request response a status code equal to that given
func (o *DcimCablesListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimCablesListBadRequest) Error() string {
	return fmt.Sprintf("[GET /dcim/cables/][%d] dcimCablesListBadRequest  %+v", 400, o.Payload)
}

func (o *DcimCablesListBadRequest) String() string {
	return fmt.Sprintf("[GET /dcim/cables/][%d] dcimCablesListBadRequest  %+v", 400, o.Payload)
}

func (o *DcimCablesListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimCablesListDefault creates a DcimCablesListDefault with default headers values
func NewDcimCablesListDefault(code int) *DcimCablesListDefault {
	return &DcimCablesListDefault{
		_statusCode: code,
	}
}

/*
DcimCablesListDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type DcimCablesListDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim cables list default response
func (o *DcimCablesListDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim cables list default response has a 2xx status code
func (o *DcimCablesListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim cables list default response has a 3xx status code
func (o *DcimCablesListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim cables list default response has a 4xx status code
func (o *DcimCablesListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim cables list default response has a 5xx status code
func (o *DcimCablesListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim cables list default response a status code equal to that given
func (o *DcimCablesListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimCablesListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/cables/][%d] dcim_cables_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/cables/][%d] dcim_cables_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimCablesListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimCablesListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimCablesListOKBody dcim cables list o k body
swagger:model DcimCablesListOKBody
*/
type DcimCablesListOKBody struct {

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
	Results []*models.Cable `json:"results"`
}

// Validate validates this dcim cables list o k body
func (o *DcimCablesListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimCablesListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimCablesListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimCablesListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimCablesListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimCablesListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimCablesListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimCablesListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimCablesListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimCablesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimCablesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim cables list o k body based on the context it is used
func (o *DcimCablesListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimCablesListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimCablesListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimCablesListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimCablesListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimCablesListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimCablesListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
