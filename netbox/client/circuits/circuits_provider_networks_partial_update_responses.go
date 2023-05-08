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

package circuits

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// CircuitsProviderNetworksPartialUpdateReader is a Reader for the CircuitsProviderNetworksPartialUpdate structure.
type CircuitsProviderNetworksPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CircuitsProviderNetworksPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCircuitsProviderNetworksPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCircuitsProviderNetworksPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCircuitsProviderNetworksPartialUpdateOK creates a CircuitsProviderNetworksPartialUpdateOK with default headers values
func NewCircuitsProviderNetworksPartialUpdateOK() *CircuitsProviderNetworksPartialUpdateOK {
	return &CircuitsProviderNetworksPartialUpdateOK{}
}

/*
CircuitsProviderNetworksPartialUpdateOK describes a response with status code 200, with default header values.

CircuitsProviderNetworksPartialUpdateOK circuits provider networks partial update o k
*/
type CircuitsProviderNetworksPartialUpdateOK struct {
	Payload *models.ProviderNetwork
}

// IsSuccess returns true when this circuits provider networks partial update o k response has a 2xx status code
func (o *CircuitsProviderNetworksPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this circuits provider networks partial update o k response has a 3xx status code
func (o *CircuitsProviderNetworksPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits provider networks partial update o k response has a 4xx status code
func (o *CircuitsProviderNetworksPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this circuits provider networks partial update o k response has a 5xx status code
func (o *CircuitsProviderNetworksPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits provider networks partial update o k response a status code equal to that given
func (o *CircuitsProviderNetworksPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *CircuitsProviderNetworksPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /circuits/provider-networks/{id}/][%d] circuitsProviderNetworksPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProviderNetworksPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /circuits/provider-networks/{id}/][%d] circuitsProviderNetworksPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *CircuitsProviderNetworksPartialUpdateOK) GetPayload() *models.ProviderNetwork {
	return o.Payload
}

func (o *CircuitsProviderNetworksPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProviderNetwork)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCircuitsProviderNetworksPartialUpdateBadRequest creates a CircuitsProviderNetworksPartialUpdateBadRequest with default headers values
func NewCircuitsProviderNetworksPartialUpdateBadRequest() *CircuitsProviderNetworksPartialUpdateBadRequest {
	return &CircuitsProviderNetworksPartialUpdateBadRequest{}
}

/*
CircuitsProviderNetworksPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CircuitsProviderNetworksPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this circuits provider networks partial update bad request response has a 2xx status code
func (o *CircuitsProviderNetworksPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this circuits provider networks partial update bad request response has a 3xx status code
func (o *CircuitsProviderNetworksPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this circuits provider networks partial update bad request response has a 4xx status code
func (o *CircuitsProviderNetworksPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this circuits provider networks partial update bad request response has a 5xx status code
func (o *CircuitsProviderNetworksPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this circuits provider networks partial update bad request response a status code equal to that given
func (o *CircuitsProviderNetworksPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CircuitsProviderNetworksPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /circuits/provider-networks/{id}/][%d] circuitsProviderNetworksPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsProviderNetworksPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /circuits/provider-networks/{id}/][%d] circuitsProviderNetworksPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *CircuitsProviderNetworksPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *CircuitsProviderNetworksPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
