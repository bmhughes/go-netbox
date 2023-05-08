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

// UsersTokensCreateReader is a Reader for the UsersTokensCreate structure.
type UsersTokensCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersTokensCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersTokensCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersTokensCreateCreated creates a UsersTokensCreateCreated with default headers values
func NewUsersTokensCreateCreated() *UsersTokensCreateCreated {
	return &UsersTokensCreateCreated{}
}

/*
UsersTokensCreateCreated describes a response with status code 201, with default header values.

UsersTokensCreateCreated users tokens create created
*/
type UsersTokensCreateCreated struct {
	Payload *models.Token
}

// IsSuccess returns true when this users tokens create created response has a 2xx status code
func (o *UsersTokensCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens create created response has a 3xx status code
func (o *UsersTokensCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens create created response has a 4xx status code
func (o *UsersTokensCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens create created response has a 5xx status code
func (o *UsersTokensCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens create created response a status code equal to that given
func (o *UsersTokensCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UsersTokensCreateCreated) Error() string {
	return fmt.Sprintf("[POST /users/tokens/][%d] usersTokensCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersTokensCreateCreated) String() string {
	return fmt.Sprintf("[POST /users/tokens/][%d] usersTokensCreateCreated  %+v", 201, o.Payload)
}

func (o *UsersTokensCreateCreated) GetPayload() *models.Token {
	return o.Payload
}

func (o *UsersTokensCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensCreateBadRequest creates a UsersTokensCreateBadRequest with default headers values
func NewUsersTokensCreateBadRequest() *UsersTokensCreateBadRequest {
	return &UsersTokensCreateBadRequest{}
}

/*
UsersTokensCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersTokensCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this users tokens create bad request response has a 2xx status code
func (o *UsersTokensCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users tokens create bad request response has a 3xx status code
func (o *UsersTokensCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens create bad request response has a 4xx status code
func (o *UsersTokensCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users tokens create bad request response has a 5xx status code
func (o *UsersTokensCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens create bad request response a status code equal to that given
func (o *UsersTokensCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersTokensCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/tokens/][%d] usersTokensCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersTokensCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /users/tokens/][%d] usersTokensCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersTokensCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
