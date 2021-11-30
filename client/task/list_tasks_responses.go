// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/uempfel/tkgi-client-go/models"
)

// ListTasksReader is a Reader for the ListTasks structure.
type ListTasksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListTasksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListTasksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListTasksUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListTasksForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListTasksNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListTasksOK creates a ListTasksOK with default headers values
func NewListTasksOK() *ListTasksOK {
	return &ListTasksOK{}
}

/* ListTasksOK describes a response with status code 200, with default header values.

Listed successfully
*/
type ListTasksOK struct {
	Payload []*models.TaskInfo
}

func (o *ListTasksOK) Error() string {
	return fmt.Sprintf("[GET /v1/tasks][%d] listTasksOK  %+v", 200, o.Payload)
}
func (o *ListTasksOK) GetPayload() []*models.TaskInfo {
	return o.Payload
}

func (o *ListTasksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListTasksUnauthorized creates a ListTasksUnauthorized with default headers values
func NewListTasksUnauthorized() *ListTasksUnauthorized {
	return &ListTasksUnauthorized{}
}

/* ListTasksUnauthorized describes a response with status code 401, with default header values.

Authentication failure. Please check credentials and try again.
*/
type ListTasksUnauthorized struct {
}

func (o *ListTasksUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/tasks][%d] listTasksUnauthorized ", 401)
}

func (o *ListTasksUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListTasksForbidden creates a ListTasksForbidden with default headers values
func NewListTasksForbidden() *ListTasksForbidden {
	return &ListTasksForbidden{}
}

/* ListTasksForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListTasksForbidden struct {
}

func (o *ListTasksForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/tasks][%d] listTasksForbidden ", 403)
}

func (o *ListTasksForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListTasksNotFound creates a ListTasksNotFound with default headers values
func NewListTasksNotFound() *ListTasksNotFound {
	return &ListTasksNotFound{}
}

/* ListTasksNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListTasksNotFound struct {
}

func (o *ListTasksNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/tasks][%d] listTasksNotFound ", 404)
}

func (o *ListTasksNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
