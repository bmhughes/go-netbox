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

// UsersGroupsUpdateReader is a Reader for the UsersGroupsUpdate structure.
type UsersGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUsersGroupsUpdateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUsersGroupsUpdateOK creates a UsersGroupsUpdateOK with default headers values
func NewUsersGroupsUpdateOK() *UsersGroupsUpdateOK {
	return &UsersGroupsUpdateOK{}
}

/*
UsersGroupsUpdateOK describes a response with status code 200, with default header values.

UsersGroupsUpdateOK users groups update o k
*/
type UsersGroupsUpdateOK struct {
	Payload *models.Group
}

// IsSuccess returns true when this users groups update o k response has a 2xx status code
func (o *UsersGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users groups update o k response has a 3xx status code
func (o *UsersGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users groups update o k response has a 4xx status code
func (o *UsersGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users groups update o k response has a 5xx status code
func (o *UsersGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users groups update o k response a status code equal to that given
func (o *UsersGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *UsersGroupsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /users/groups/{id}/][%d] usersGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersGroupsUpdateOK) String() string {
	return fmt.Sprintf("[PUT /users/groups/{id}/][%d] usersGroupsUpdateOK  %+v", 200, o.Payload)
}

func (o *UsersGroupsUpdateOK) GetPayload() *models.Group {
	return o.Payload
}

func (o *UsersGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersGroupsUpdateBadRequest creates a UsersGroupsUpdateBadRequest with default headers values
func NewUsersGroupsUpdateBadRequest() *UsersGroupsUpdateBadRequest {
	return &UsersGroupsUpdateBadRequest{}
}

/*
UsersGroupsUpdateBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type UsersGroupsUpdateBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this users groups update bad request response has a 2xx status code
func (o *UsersGroupsUpdateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this users groups update bad request response has a 3xx status code
func (o *UsersGroupsUpdateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users groups update bad request response has a 4xx status code
func (o *UsersGroupsUpdateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this users groups update bad request response has a 5xx status code
func (o *UsersGroupsUpdateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this users groups update bad request response a status code equal to that given
func (o *UsersGroupsUpdateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UsersGroupsUpdateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /users/groups/{id}/][%d] usersGroupsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersGroupsUpdateBadRequest) String() string {
	return fmt.Sprintf("[PUT /users/groups/{id}/][%d] usersGroupsUpdateBadRequest  %+v", 400, o.Payload)
}

func (o *UsersGroupsUpdateBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *UsersGroupsUpdateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
