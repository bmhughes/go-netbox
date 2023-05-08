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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimRegionsBulkDeleteReader is a Reader for the DcimRegionsBulkDelete structure.
type DcimRegionsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRegionsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimRegionsBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRegionsBulkDeleteNoContent creates a DcimRegionsBulkDeleteNoContent with default headers values
func NewDcimRegionsBulkDeleteNoContent() *DcimRegionsBulkDeleteNoContent {
	return &DcimRegionsBulkDeleteNoContent{}
}

/*
DcimRegionsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRegionsBulkDeleteNoContent dcim regions bulk delete no content
*/
type DcimRegionsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim regions bulk delete no content response has a 2xx status code
func (o *DcimRegionsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim regions bulk delete no content response has a 3xx status code
func (o *DcimRegionsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions bulk delete no content response has a 4xx status code
func (o *DcimRegionsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim regions bulk delete no content response has a 5xx status code
func (o *DcimRegionsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions bulk delete no content response a status code equal to that given
func (o *DcimRegionsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimRegionsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/regions/][%d] dcimRegionsBulkDeleteNoContent ", 204)
}

func (o *DcimRegionsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/regions/][%d] dcimRegionsBulkDeleteNoContent ", 204)
}

func (o *DcimRegionsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRegionsBulkDeleteBadRequest creates a DcimRegionsBulkDeleteBadRequest with default headers values
func NewDcimRegionsBulkDeleteBadRequest() *DcimRegionsBulkDeleteBadRequest {
	return &DcimRegionsBulkDeleteBadRequest{}
}

/*
DcimRegionsBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimRegionsBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim regions bulk delete bad request response has a 2xx status code
func (o *DcimRegionsBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim regions bulk delete bad request response has a 3xx status code
func (o *DcimRegionsBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim regions bulk delete bad request response has a 4xx status code
func (o *DcimRegionsBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim regions bulk delete bad request response has a 5xx status code
func (o *DcimRegionsBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim regions bulk delete bad request response a status code equal to that given
func (o *DcimRegionsBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimRegionsBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/regions/][%d] dcimRegionsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRegionsBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/regions/][%d] dcimRegionsBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimRegionsBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRegionsBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
