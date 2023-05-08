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

// DcimSitesBulkDeleteReader is a Reader for the DcimSitesBulkDelete structure.
type DcimSitesBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimSitesBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimSitesBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDcimSitesBulkDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimSitesBulkDeleteNoContent creates a DcimSitesBulkDeleteNoContent with default headers values
func NewDcimSitesBulkDeleteNoContent() *DcimSitesBulkDeleteNoContent {
	return &DcimSitesBulkDeleteNoContent{}
}

/*
DcimSitesBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimSitesBulkDeleteNoContent dcim sites bulk delete no content
*/
type DcimSitesBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim sites bulk delete no content response has a 2xx status code
func (o *DcimSitesBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim sites bulk delete no content response has a 3xx status code
func (o *DcimSitesBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites bulk delete no content response has a 4xx status code
func (o *DcimSitesBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim sites bulk delete no content response has a 5xx status code
func (o *DcimSitesBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites bulk delete no content response a status code equal to that given
func (o *DcimSitesBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimSitesBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/sites/][%d] dcimSitesBulkDeleteNoContent ", 204)
}

func (o *DcimSitesBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/sites/][%d] dcimSitesBulkDeleteNoContent ", 204)
}

func (o *DcimSitesBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimSitesBulkDeleteBadRequest creates a DcimSitesBulkDeleteBadRequest with default headers values
func NewDcimSitesBulkDeleteBadRequest() *DcimSitesBulkDeleteBadRequest {
	return &DcimSitesBulkDeleteBadRequest{}
}

/*
DcimSitesBulkDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type DcimSitesBulkDeleteBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this dcim sites bulk delete bad request response has a 2xx status code
func (o *DcimSitesBulkDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this dcim sites bulk delete bad request response has a 3xx status code
func (o *DcimSitesBulkDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim sites bulk delete bad request response has a 4xx status code
func (o *DcimSitesBulkDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this dcim sites bulk delete bad request response has a 5xx status code
func (o *DcimSitesBulkDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim sites bulk delete bad request response a status code equal to that given
func (o *DcimSitesBulkDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DcimSitesBulkDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /dcim/sites/][%d] dcimSitesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimSitesBulkDeleteBadRequest) String() string {
	return fmt.Sprintf("[DELETE /dcim/sites/][%d] dcimSitesBulkDeleteBadRequest  %+v", 400, o.Payload)
}

func (o *DcimSitesBulkDeleteBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimSitesBulkDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
