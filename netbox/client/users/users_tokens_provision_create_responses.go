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
)

// UsersTokensProvisionCreateReader is a Reader for the UsersTokensProvisionCreate structure.
type UsersTokensProvisionCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersTokensProvisionCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUsersTokensProvisionCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersTokensProvisionCreateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersTokensProvisionCreateCreated creates a UsersTokensProvisionCreateCreated with default headers values
func NewUsersTokensProvisionCreateCreated() *UsersTokensProvisionCreateCreated {
	return &UsersTokensProvisionCreateCreated{}
}

/*
UsersTokensProvisionCreateCreated describes a response with status code 201, with default header values.

UsersTokensProvisionCreateCreated users tokens provision create created
*/
type UsersTokensProvisionCreateCreated struct {
}

// IsSuccess returns true when this users tokens provision create created response has a 2xx status code
func (o *UsersTokensProvisionCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users tokens provision create created response has a 3xx status code
func (o *UsersTokensProvisionCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens provision create created response has a 4xx status code
func (o *UsersTokensProvisionCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this users tokens provision create created response has a 5xx status code
func (o *UsersTokensProvisionCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens provision create created response a status code equal to that given
func (o *UsersTokensProvisionCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UsersTokensProvisionCreateCreated) Error() string {
	return fmt.Sprintf("[POST /users/tokens/provision/][%d] usersTokensProvisionCreateCreated ", 201)
}

func (o *UsersTokensProvisionCreateCreated) String() string {
	return fmt.Sprintf("[POST /users/tokens/provision/][%d] usersTokensProvisionCreateCreated ", 201)
}

func (o *UsersTokensProvisionCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUsersTokensProvisionCreateBadRequest creates a UsersTokensProvisionCreateBadRequest with default headers values
func NewUsersTokensProvisionCreateBadRequest() *UsersTokensProvisionCreateBadRequest {
	return &UsersTokensProvisionCreateBadRequest{}
}

/*
UsersTokensProvisionCreateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersTokensProvisionCreateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this users tokens provision create bad request response has a 2xx status code
func (o *UsersTokensProvisionCreateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users tokens provision create bad request response has a 3xx status code
func (o *UsersTokensProvisionCreateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users tokens provision create bad request response has a 4xx status code
func (o *UsersTokensProvisionCreateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users tokens provision create bad request response has a 5xx status code
func (o *UsersTokensProvisionCreateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users tokens provision create bad request response a status code equal to that given
func (o *UsersTokensProvisionCreateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersTokensProvisionCreateBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/tokens/provision/][%d] usersTokensProvisionCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersTokensProvisionCreateBadRequest) String() string {
	return fmt.Sprintf("[POST /users/tokens/provision/][%d] usersTokensProvisionCreateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersTokensProvisionCreateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersTokensProvisionCreateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
