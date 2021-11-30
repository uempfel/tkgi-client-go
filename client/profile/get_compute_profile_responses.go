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

// GetComputeProfileReader is a Reader for the GetComputeProfile structure.
type GetComputeProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetComputeProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetComputeProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetComputeProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetComputeProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetComputeProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetComputeProfileOK creates a GetComputeProfileOK with default headers values
func NewGetComputeProfileOK() *GetComputeProfileOK {
	return &GetComputeProfileOK{}
}

/* GetComputeProfileOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetComputeProfileOK struct {
	Payload *models.ComputeProfile
}

func (o *GetComputeProfileOK) Error() string {
	return fmt.Sprintf("[GET /v1/compute-profiles/{profileName}][%d] getComputeProfileOK  %+v", 200, o.Payload)
}
func (o *GetComputeProfileOK) GetPayload() *models.ComputeProfile {
	return o.Payload
}

func (o *GetComputeProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ComputeProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetComputeProfileUnauthorized creates a GetComputeProfileUnauthorized with default headers values
func NewGetComputeProfileUnauthorized() *GetComputeProfileUnauthorized {
	return &GetComputeProfileUnauthorized{}
}

/* GetComputeProfileUnauthorized describes a response with status code 401, with default header values.

Authentication failure. Please check credentials and try again.
*/
type GetComputeProfileUnauthorized struct {
}

func (o *GetComputeProfileUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/compute-profiles/{profileName}][%d] getComputeProfileUnauthorized ", 401)
}

func (o *GetComputeProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComputeProfileForbidden creates a GetComputeProfileForbidden with default headers values
func NewGetComputeProfileForbidden() *GetComputeProfileForbidden {
	return &GetComputeProfileForbidden{}
}

/* GetComputeProfileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetComputeProfileForbidden struct {
}

func (o *GetComputeProfileForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/compute-profiles/{profileName}][%d] getComputeProfileForbidden ", 403)
}

func (o *GetComputeProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetComputeProfileNotFound creates a GetComputeProfileNotFound with default headers values
func NewGetComputeProfileNotFound() *GetComputeProfileNotFound {
	return &GetComputeProfileNotFound{}
}

/* GetComputeProfileNotFound describes a response with status code 404, with default header values.

Compute profile does not exist
*/
type GetComputeProfileNotFound struct {
	Payload *models.ErrorResponse
}

func (o *GetComputeProfileNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/compute-profiles/{profileName}][%d] getComputeProfileNotFound  %+v", 404, o.Payload)
}
func (o *GetComputeProfileNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetComputeProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
