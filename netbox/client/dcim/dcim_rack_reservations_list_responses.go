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

// DcimRackReservationsListReader is a Reader for the DcimRackReservationsList structure.
type DcimRackReservationsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRackReservationsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimRackReservationsListBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRackReservationsListOK creates a DcimRackReservationsListOK with default headers values
func NewDcimRackReservationsListOK() *DcimRackReservationsListOK {
	return &DcimRackReservationsListOK{}
}

/*
DcimRackReservationsListOK describes a response with status code 200, with default header values.

DcimRackReservationsListOK dcim rack reservations list o k
*/
type DcimRackReservationsListOK struct {
	Payload *DcimRackReservationsListOKBody
}

// IsSuccess returns true when this dcim rack reservations list o k response has a 2xx status code
func (o *DcimRackReservationsListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack reservations list o k response has a 3xx status code
func (o *DcimRackReservationsListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack reservations list o k response has a 4xx status code
func (o *DcimRackReservationsListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack reservations list o k response has a 5xx status code
func (o *DcimRackReservationsListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack reservations list o k response a status code equal to that given
func (o *DcimRackReservationsListOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRackReservationsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-reservations/][%d] dcimRackReservationsListOK  %+v", 200, o.Payload)
}

func (o *DcimRackReservationsListOK) String() string {
	return fmt.Sprintf("[GET /dcim/rack-reservations/][%d] dcimRackReservationsListOK  %+v", 200, o.Payload)
}

func (o *DcimRackReservationsListOK) GetPayload() *DcimRackReservationsListOKBody {
	return o.Payload
}

func (o *DcimRackReservationsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimRackReservationsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRackReservationsListBadRequest creates a DcimRackReservationsListBadRequest with default headers values
func NewDcimRackReservationsListBadRequest() *DcimRackReservationsListBadRequest {
	return &DcimRackReservationsListBadRequest{}
}

/*
DcimRackReservationsListBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimRackReservationsListBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim rack reservations list bad request response has a 2xx status code
func (o *DcimRackReservationsListBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim rack reservations list bad request response has a 3xx status code
func (o *DcimRackReservationsListBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack reservations list bad request response has a 4xx status code
func (o *DcimRackReservationsListBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim rack reservations list bad request response has a 5xx status code
func (o *DcimRackReservationsListBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack reservations list bad request response a status code equal to that given
func (o *DcimRackReservationsListBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimRackReservationsListBadRequest) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-reservations/][%d] dcimRackReservationsListBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRackReservationsListBadRequest) String() string {
	return fmt.Sprintf("[GET /dcim/rack-reservations/][%d] dcimRackReservationsListBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRackReservationsListBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackReservationsListBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimRackReservationsListOKBody dcim rack reservations list o k body
swagger:model DcimRackReservationsListOKBody
*/
type DcimRackReservationsListOKBody struct {

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
	Results []*models.RackReservation `json:"results"`
}

// Validate validates this dcim rack reservations list o k body
func (o *DcimRackReservationsListOKBody) Validate(formats strfmt.Registry) error {
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

func (o *DcimRackReservationsListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimRackReservationsListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimRackReservationsListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimRackReservationsListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimRackReservationsListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimRackReservationsListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimRackReservationsListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimRackReservationsListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimRackReservationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimRackReservationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim rack reservations list o k body based on the context it is used
func (o *DcimRackReservationsListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimRackReservationsListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimRackReservationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimRackReservationsListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimRackReservationsListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimRackReservationsListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimRackReservationsListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
