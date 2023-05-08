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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamRouteTargetsDeleteReader is a Reader for the IpamRouteTargetsDelete structure.
type IpamRouteTargetsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamRouteTargetsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamRouteTargetsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewIpamRouteTargetsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamRouteTargetsDeleteNoContent creates a IpamRouteTargetsDeleteNoContent with default headers values
func NewIpamRouteTargetsDeleteNoContent() *IpamRouteTargetsDeleteNoContent {
	return &IpamRouteTargetsDeleteNoContent{}
}

/*
IpamRouteTargetsDeleteNoContent describes a response with status code 204, with default header values.

IpamRouteTargetsDeleteNoContent ipam route targets delete no content
*/
type IpamRouteTargetsDeleteNoContent struct {
}

// IsSuccess returns true when this ipam route targets delete no content response has a 2xx status code
func (o *IpamRouteTargetsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam route targets delete no content response has a 3xx status code
func (o *IpamRouteTargetsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam route targets delete no content response has a 4xx status code
func (o *IpamRouteTargetsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam route targets delete no content response has a 5xx status code
func (o *IpamRouteTargetsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam route targets delete no content response a status code equal to that given
func (o *IpamRouteTargetsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamRouteTargetsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/{id}/][%d] ipamRouteTargetsDeleteNoContent ", 204)
}

func (o *IpamRouteTargetsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/{id}/][%d] ipamRouteTargetsDeleteNoContent ", 204)
}

func (o *IpamRouteTargetsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewIpamRouteTargetsDeleteBadRequest creates a IpamRouteTargetsDeleteBadRequest with default headers values
func NewIpamRouteTargetsDeleteBadRequest() *IpamRouteTargetsDeleteBadRequest {
	return &IpamRouteTargetsDeleteBadRequest{}
}

/*
IpamRouteTargetsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type IpamRouteTargetsDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this ipam route targets delete bad request response has a 2xx status code
func (o *IpamRouteTargetsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ipam route targets delete bad request response has a 3xx status code
func (o *IpamRouteTargetsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam route targets delete bad request response has a 4xx status code
func (o *IpamRouteTargetsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ipam route targets delete bad request response has a 5xx status code
func (o *IpamRouteTargetsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam route targets delete bad request response a status code equal to that given
func (o *IpamRouteTargetsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *IpamRouteTargetsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/{id}/][%d] ipamRouteTargetsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRouteTargetsDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /ipam/route-targets/{id}/][%d] ipamRouteTargetsDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *IpamRouteTargetsDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamRouteTargetsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
