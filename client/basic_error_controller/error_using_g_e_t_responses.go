// Code generated by go-swagger; DO NOT EDIT.

package basic_error_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ErrorUsingGETReader is a Reader for the ErrorUsingGET structure.
type ErrorUsingGETReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ErrorUsingGETReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewErrorUsingGETOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewErrorUsingGETUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewErrorUsingGETForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewErrorUsingGETNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewErrorUsingGETOK creates a ErrorUsingGETOK with default headers values
func NewErrorUsingGETOK() *ErrorUsingGETOK {
	return &ErrorUsingGETOK{}
}

/* ErrorUsingGETOK describes a response with status code 200, with default header values.

OK
*/
type ErrorUsingGETOK struct {
	Payload map[string]interface{}
}

func (o *ErrorUsingGETOK) Error() string {
	return fmt.Sprintf("[GET /error][%d] errorUsingGETOK  %+v", 200, o.Payload)
}
func (o *ErrorUsingGETOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *ErrorUsingGETOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewErrorUsingGETUnauthorized creates a ErrorUsingGETUnauthorized with default headers values
func NewErrorUsingGETUnauthorized() *ErrorUsingGETUnauthorized {
	return &ErrorUsingGETUnauthorized{}
}

/* ErrorUsingGETUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ErrorUsingGETUnauthorized struct {
}

func (o *ErrorUsingGETUnauthorized) Error() string {
	return fmt.Sprintf("[GET /error][%d] errorUsingGETUnauthorized ", 401)
}

func (o *ErrorUsingGETUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorUsingGETForbidden creates a ErrorUsingGETForbidden with default headers values
func NewErrorUsingGETForbidden() *ErrorUsingGETForbidden {
	return &ErrorUsingGETForbidden{}
}

/* ErrorUsingGETForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ErrorUsingGETForbidden struct {
}

func (o *ErrorUsingGETForbidden) Error() string {
	return fmt.Sprintf("[GET /error][%d] errorUsingGETForbidden ", 403)
}

func (o *ErrorUsingGETForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorUsingGETNotFound creates a ErrorUsingGETNotFound with default headers values
func NewErrorUsingGETNotFound() *ErrorUsingGETNotFound {
	return &ErrorUsingGETNotFound{}
}

/* ErrorUsingGETNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ErrorUsingGETNotFound struct {
}

func (o *ErrorUsingGETNotFound) Error() string {
	return fmt.Sprintf("[GET /error][%d] errorUsingGETNotFound ", 404)
}

func (o *ErrorUsingGETNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
