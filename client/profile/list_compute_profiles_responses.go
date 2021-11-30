// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/uempfel/tkgi-client-go/models"
)

// ListComputeProfilesReader is a Reader for the ListComputeProfiles structure.
type ListComputeProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListComputeProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListComputeProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListComputeProfilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListComputeProfilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListComputeProfilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListComputeProfilesOK creates a ListComputeProfilesOK with default headers values
func NewListComputeProfilesOK() *ListComputeProfilesOK {
	return &ListComputeProfilesOK{}
}

/* ListComputeProfilesOK describes a response with status code 200, with default header values.

Listed successfully
*/
type ListComputeProfilesOK struct {
	Payload []*models.ComputeProfile
}

func (o *ListComputeProfilesOK) Error() string {
	return fmt.Sprintf("[GET /v1/compute-profiles][%d] listComputeProfilesOK  %+v", 200, o.Payload)
}
func (o *ListComputeProfilesOK) GetPayload() []*models.ComputeProfile {
	return o.Payload
}

func (o *ListComputeProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListComputeProfilesUnauthorized creates a ListComputeProfilesUnauthorized with default headers values
func NewListComputeProfilesUnauthorized() *ListComputeProfilesUnauthorized {
	return &ListComputeProfilesUnauthorized{}
}

/* ListComputeProfilesUnauthorized describes a response with status code 401, with default header values.

Authentication failure. Please check credentials and try again.
*/
type ListComputeProfilesUnauthorized struct {
}

func (o *ListComputeProfilesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/compute-profiles][%d] listComputeProfilesUnauthorized ", 401)
}

func (o *ListComputeProfilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListComputeProfilesForbidden creates a ListComputeProfilesForbidden with default headers values
func NewListComputeProfilesForbidden() *ListComputeProfilesForbidden {
	return &ListComputeProfilesForbidden{}
}

/* ListComputeProfilesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListComputeProfilesForbidden struct {
}

func (o *ListComputeProfilesForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/compute-profiles][%d] listComputeProfilesForbidden ", 403)
}

func (o *ListComputeProfilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListComputeProfilesNotFound creates a ListComputeProfilesNotFound with default headers values
func NewListComputeProfilesNotFound() *ListComputeProfilesNotFound {
	return &ListComputeProfilesNotFound{}
}

/* ListComputeProfilesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListComputeProfilesNotFound struct {
}

func (o *ListComputeProfilesNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/compute-profiles][%d] listComputeProfilesNotFound ", 404)
}

func (o *ListComputeProfilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}