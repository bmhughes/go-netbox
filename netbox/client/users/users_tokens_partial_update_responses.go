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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// UsersTokensPartialUpdateReader is a Reader for the UsersTokensPartialUpdate structure.
type UsersTokensPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersTokensPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersTokensPartialUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersTokensPartialUpdateOK creates a UsersTokensPartialUpdateOK with default headers values
func NewUsersTokensPartialUpdateOK() *UsersTokensPartialUpdateOK {
	return &UsersTokensPartialUpdateOK{}
}

/*
UsersTokensPartialUpdateOK describes a response with status code 200, with default header values.

UsersTokensPartialUpdateOK users tokens partial update o k
*/
type UsersTokensPartialUpdateOK struct {
	Payload *models.Token
}

// IsSuccess returns true when this users tokens partial update o k response has a 2xx status code
func (o *UsersTokensPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens partial update o k response has a 3xx status code
func (o *UsersTokensPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens partial update o k response has a 4xx status code
func (o *UsersTokensPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens partial update o k response has a 5xx status code
func (o *UsersTokensPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens partial update o k response a status code equal to that given
func (o *UsersTokensPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersTokensPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /users/tokens/{id}/][%d] usersTokensPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersTokensPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /users/tokens/{id}/][%d] usersTokensPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersTokensPartialUpdateOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *UsersTokensPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensPartialUpdateBadRequest creates a UsersTokensPartialUpdateBadRequest with default headers values
func NewUsersTokensPartialUpdateBadRequest() *UsersTokensPartialUpdateBadRequest {
	return &UsersTokensPartialUpdateBadRequest{}
}

/*
UsersTokensPartialUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersTokensPartialUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this users tokens partial update bad request response has a 2xx status code
func (o *UsersTokensPartialUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users tokens partial update bad request response has a 3xx status code
func (o *UsersTokensPartialUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens partial update bad request response has a 4xx status code
func (o *UsersTokensPartialUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users tokens partial update bad request response has a 5xx status code
func (o *UsersTokensPartialUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens partial update bad request response a status code equal to that given
func (o *UsersTokensPartialUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersTokensPartialUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /users/tokens/{id}/][%d] usersTokensPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersTokensPartialUpdateBadRequest) String() string {
	return fmt.Sprintf("[PATCH /users/tokens/{id}/][%d] usersTokensPartialUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersTokensPartialUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensPartialUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
