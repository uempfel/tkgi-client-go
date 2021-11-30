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

// ErrorUsingOPTIONSReader is a Reader for the ErrorUsingOPTIONS structure.
type ErrorUsingOPTIONSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ErrorUsingOPTIONSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewErrorUsingOPTIONSOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 204:
		result := NewErrorUsingOPTIONSNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewErrorUsingOPTIONSUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewErrorUsingOPTIONSForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewErrorUsingOPTIONSOK creates a ErrorUsingOPTIONSOK with default headers values
func NewErrorUsingOPTIONSOK() *ErrorUsingOPTIONSOK {
	return &ErrorUsingOPTIONSOK{}
}

/* ErrorUsingOPTIONSOK describes a response with status code 200, with default header values.

OK
*/
type ErrorUsingOPTIONSOK struct {
	Payload map[string]interface{}
}

func (o *ErrorUsingOPTIONSOK) Error() string {
	return fmt.Sprintf("[OPTIONS /error][%d] errorUsingOPTIONSOK  %+v", 200, o.Payload)
}
func (o *ErrorUsingOPTIONSOK) GetPayload() map[string]interface{} {
	return o.Payload
}

func (o *ErrorUsingOPTIONSOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewErrorUsingOPTIONSNoContent creates a ErrorUsingOPTIONSNoContent with default headers values
func NewErrorUsingOPTIONSNoContent() *ErrorUsingOPTIONSNoContent {
	return &ErrorUsingOPTIONSNoContent{}
}

/* ErrorUsingOPTIONSNoContent describes a response with status code 204, with default header values.

No Content
*/
type ErrorUsingOPTIONSNoContent struct {
}

func (o *ErrorUsingOPTIONSNoContent) Error() string {
	return fmt.Sprintf("[OPTIONS /error][%d] errorUsingOPTIONSNoContent ", 204)
}

func (o *ErrorUsingOPTIONSNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorUsingOPTIONSUnauthorized creates a ErrorUsingOPTIONSUnauthorized with default headers values
func NewErrorUsingOPTIONSUnauthorized() *ErrorUsingOPTIONSUnauthorized {
	return &ErrorUsingOPTIONSUnauthorized{}
}

/* ErrorUsingOPTIONSUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ErrorUsingOPTIONSUnauthorized struct {
}

func (o *ErrorUsingOPTIONSUnauthorized) Error() string {
	return fmt.Sprintf("[OPTIONS /error][%d] errorUsingOPTIONSUnauthorized ", 401)
}

func (o *ErrorUsingOPTIONSUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewErrorUsingOPTIONSForbidden creates a ErrorUsingOPTIONSForbidden with default headers values
func NewErrorUsingOPTIONSForbidden() *ErrorUsingOPTIONSForbidden {
	return &ErrorUsingOPTIONSForbidden{}
}

/* ErrorUsingOPTIONSForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ErrorUsingOPTIONSForbidden struct {
}

func (o *ErrorUsingOPTIONSForbidden) Error() string {
	return fmt.Sprintf("[OPTIONS /error][%d] errorUsingOPTIONSForbidden ", 403)
}

func (o *ErrorUsingOPTIONSForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
