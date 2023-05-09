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

// UsersTokensReadReader is a Reader for the UsersTokensRead structure.
type UsersTokensReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersTokensReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersTokensReadBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewUsersTokensReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersTokensReadOK creates a UsersTokensReadOK with default headers values
func NewUsersTokensReadOK() *UsersTokensReadOK {
	return &UsersTokensReadOK{}
}

/*
UsersTokensReadOK describes a response with status code 200, with default header values.

UsersTokensReadOK users tokens read o k
*/
type UsersTokensReadOK struct {
	Payload *models.Token
}

// IsSuccess returns true when this users tokens read o k response has a 2xx status code
func (o *UsersTokensReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens read o k response has a 3xx status code
func (o *UsersTokensReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens read o k response has a 4xx status code
func (o *UsersTokensReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens read o k response has a 5xx status code
func (o *UsersTokensReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens read o k response a status code equal to that given
func (o *UsersTokensReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersTokensReadOK) Error() string {
	return fmt.Sprintf("[GET /users/tokens/{id}/][%d] usersTokensReadOK  %+v", 200, o.Payload)
}

func (o *UsersTokensReadOK) String() string {
	return fmt.Sprintf("[GET /users/tokens/{id}/][%d] usersTokensReadOK  %+v", 200, o.Payload)
}

func (o *UsersTokensReadOK) GetPayload() *models.Token {
	return o.Payload
}

func (o *UsersTokensReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Token)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensReadBadRequest creates a UsersTokensReadBadRequest with default headers values
func NewUsersTokensReadBadRequest() *UsersTokensReadBadRequest {
	return &UsersTokensReadBadRequest{}
}

/*
UsersTokensReadBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersTokensReadBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this users tokens read bad request response has a 2xx status code
func (o *UsersTokensReadBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users tokens read bad request response has a 3xx status code
func (o *UsersTokensReadBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens read bad request response has a 4xx status code
func (o *UsersTokensReadBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users tokens read bad request response has a 5xx status code
func (o *UsersTokensReadBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens read bad request response a status code equal to that given
func (o *UsersTokensReadBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersTokensReadBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/tokens/{id}/][%d] usersTokensReadBadRequest  %+v", 400, o.Payload)
}

func (o *UsersTokensReadBadRequest) String() string {
	return fmt.Sprintf("[GET /users/tokens/{id}/][%d] usersTokensReadBadRequest  %+v", 400, o.Payload)
}

func (o *UsersTokensReadBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensReadBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersTokensReadDefault creates a UsersTokensReadDefault with default headers values
func NewUsersTokensReadDefault(code int) *UsersTokensReadDefault {
	return &UsersTokensReadDefault{
		_statusCode: code,
	}
}

/*
UsersTokensReadDefault describes a response with status code -1, with default header values.

Unexpected Response
*/
type UsersTokensReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the users tokens read default response
func (o *UsersTokensReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this users tokens read default response has a 2xx status code
func (o *UsersTokensReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users tokens read default response has a 3xx status code
func (o *UsersTokensReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users tokens read default response has a 4xx status code
func (o *UsersTokensReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users tokens read default response has a 5xx status code
func (o *UsersTokensReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users tokens read default response a status code equal to that given
func (o *UsersTokensReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UsersTokensReadDefault) Error() string {
	return fmt.Sprintf("[GET /users/tokens/{id}/][%d] users_tokens_read default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensReadDefault) String() string {
	return fmt.Sprintf("[GET /users/tokens/{id}/][%d] users_tokens_read default  %+v", o._statusCode, o.Payload)
}

func (o *UsersTokensReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
