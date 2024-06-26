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

// UsersUsersUpdateReader is a Reader for the UsersUsersUpdate structure.
type UsersUsersUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersUsersUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersUsersUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersUsersUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUsersUsersUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersUsersUpdateOK creates a UsersUsersUpdateOK with default headers values
func NewUsersUsersUpdateOK() *UsersUsersUpdateOK {
	return &UsersUsersUpdateOK{}
}

/*
UsersUsersUpdateOK describes a response with status code 200, with default header values.

UsersUsersUpdateOK users users update o k
*/
type UsersUsersUpdateOK struct {
	Payload *models.User
}

// IsSuccess returns true when this users users update o k response has a 2xx status code
func (o *UsersUsersUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users users update o k response has a 3xx status code
func (o *UsersUsersUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users update o k response has a 4xx status code
func (o *UsersUsersUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users users update o k response has a 5xx status code
func (o *UsersUsersUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users users update o k response a status code equal to that given
func (o *UsersUsersUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersUsersUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/users/{id}/][%d] usersUsersUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersUsersUpdateOK) String() string {
	return fmt.Sprintf("[PUT /users/users/{id}/][%d] usersUsersUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersUsersUpdateOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersUsersUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUsersUpdateBadRequest creates a UsersUsersUpdateBadRequest with default headers values
func NewUsersUsersUpdateBadRequest() *UsersUsersUpdateBadRequest {
	return &UsersUsersUpdateBadRequest{}
}

/*
UsersUsersUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersUsersUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this users users update bad request response has a 2xx status code
func (o *UsersUsersUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users users update bad request response has a 3xx status code
func (o *UsersUsersUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users users update bad request response has a 4xx status code
func (o *UsersUsersUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users users update bad request response has a 5xx status code
func (o *UsersUsersUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users users update bad request response a status code equal to that given
func (o *UsersUsersUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersUsersUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/users/{id}/][%d] usersUsersUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersUsersUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /users/users/{id}/][%d] usersUsersUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersUsersUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersUsersUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersUsersUpdateDefault creates a UsersUsersUpdateDefault with default headers values
func NewUsersUsersUpdateDefault(code int) *UsersUsersUpdateDefault {
	return &UsersUsersUpdateDefault{
		_statusCode: code,
	}
}

/*
UsersUsersUpdateDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type UsersUsersUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users users update default response
func (o *UsersUsersUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users users update default response has a 2xx status code
func (o *UsersUsersUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users users update default response has a 3xx status code
func (o *UsersUsersUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users users update default response has a 4xx status code
func (o *UsersUsersUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users users update default response has a 5xx status code
func (o *UsersUsersUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users users update default response a status code equal to that given
func (o *UsersUsersUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersUsersUpdateDefault) Error() string {
	return fmt.Sprintf("[PUT /users/users/{id}/][%d] users_users_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersUpdateDefault) String() string {
	return fmt.Sprintf("[PUT /users/users/{id}/][%d] users_users_update default  %+v", o._statusCode, o.Payload)
}

func (o *UsersUsersUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersUsersUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
