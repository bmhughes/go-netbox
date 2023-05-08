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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// TenancyTenantsUpdateReader is a Reader for the TenancyTenantsUpdate structure.
type TenancyTenantsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTenancyTenantsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTenancyTenantsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTenancyTenantsUpdateOK creates a TenancyTenantsUpdateOK with default headers values
func NewTenancyTenantsUpdateOK() *TenancyTenantsUpdateOK {
	return &TenancyTenantsUpdateOK{}
}

/*
TenancyTenantsUpdateOK describes a response with status code 200, with default header values.

TenancyTenantsUpdateOK tenancy tenants update o k
*/
type TenancyTenantsUpdateOK struct {
	Payload *models.Tenant
}

// IsSuccess returns true when this tenancy tenants update o k response has a 2xx status code
func (o *TenancyTenantsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenants update o k response has a 3xx status code
func (o *TenancyTenantsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenants update o k response has a 4xx status code
func (o *TenancyTenantsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenants update o k response has a 5xx status code
func (o *TenancyTenantsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenants update o k response a status code equal to that given
func (o *TenancyTenantsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *TenancyTenantsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenants/{id}/][%d] tenancyTenantsUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyTenantsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /tenancy/tenants/{id}/][%d] tenancyTenantsUpdateOK  %+v", 200, o.Payload)
}

func (o *TenancyTenantsUpdateOK) GetPayload() *models.Tenant {
	return o.Payload
}

func (o *TenancyTenantsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Tenant)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTenancyTenantsUpdateBadRequest creates a TenancyTenantsUpdateBadRequest with default headers values
func NewTenancyTenantsUpdateBadRequest() *TenancyTenantsUpdateBadRequest {
	return &TenancyTenantsUpdateBadRequest{}
}

/*
TenancyTenantsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type TenancyTenantsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this tenancy tenants update bad request response has a 2xx status code
func (o *TenancyTenantsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this tenancy tenants update bad request response has a 3xx status code
func (o *TenancyTenantsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenants update bad request response has a 4xx status code
func (o *TenancyTenantsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this tenancy tenants update bad request response has a 5xx status code
func (o *TenancyTenantsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenants update bad request response a status code equal to that given
func (o *TenancyTenantsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *TenancyTenantsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /tenancy/tenants/{id}/][%d] tenancyTenantsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyTenantsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /tenancy/tenants/{id}/][%d] tenancyTenantsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *TenancyTenantsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
