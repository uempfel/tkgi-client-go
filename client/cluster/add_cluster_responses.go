// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/uempfel/tkgi-client-go/models"
)

// AddClusterReader is a Reader for the AddCluster structure.
type AddClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewAddClusterCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewAddClusterAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddClusterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddClusterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewAddClusterMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddClusterConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewAddClusterUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddClusterOK creates a AddClusterOK with default headers values
func NewAddClusterOK() *AddClusterOK {
	return &AddClusterOK{}
}

/* AddClusterOK describes a response with status code 200, with default header values.

OK
*/
type AddClusterOK struct {
	Payload *models.Cluster
}

func (o *AddClusterOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] addClusterOK  %+v", 200, o.Payload)
}
func (o *AddClusterOK) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *AddClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddClusterCreated creates a AddClusterCreated with default headers values
func NewAddClusterCreated() *AddClusterCreated {
	return &AddClusterCreated{}
}

/* AddClusterCreated describes a response with status code 201, with default header values.

Created
*/
type AddClusterCreated struct {
}

func (o *AddClusterCreated) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] addClusterCreated ", 201)
}

func (o *AddClusterCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddClusterAccepted creates a AddClusterAccepted with default headers values
func NewAddClusterAccepted() *AddClusterAccepted {
	return &AddClusterAccepted{}
}

/* AddClusterAccepted describes a response with status code 202, with default header values.

Successful operation
*/
type AddClusterAccepted struct {
	Payload *models.Cluster
}

func (o *AddClusterAccepted) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] addClusterAccepted  %+v", 202, o.Payload)
}
func (o *AddClusterAccepted) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *AddClusterAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddClusterBadRequest creates a AddClusterBadRequest with default headers values
func NewAddClusterBadRequest() *AddClusterBadRequest {
	return &AddClusterBadRequest{}
}

/* AddClusterBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddClusterBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *AddClusterBadRequest) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] addClusterBadRequest  %+v", 400, o.Payload)
}
func (o *AddClusterBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddClusterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddClusterUnauthorized creates a AddClusterUnauthorized with default headers values
func NewAddClusterUnauthorized() *AddClusterUnauthorized {
	return &AddClusterUnauthorized{}
}

/* AddClusterUnauthorized describes a response with status code 401, with default header values.

Authentication failure. Please check credentials and try again.
*/
type AddClusterUnauthorized struct {
}

func (o *AddClusterUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] addClusterUnauthorized ", 401)
}

func (o *AddClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddClusterForbidden creates a AddClusterForbidden with default headers values
func NewAddClusterForbidden() *AddClusterForbidden {
	return &AddClusterForbidden{}
}

/* AddClusterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddClusterForbidden struct {
}

func (o *AddClusterForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] addClusterForbidden ", 403)
}

func (o *AddClusterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddClusterNotFound creates a AddClusterNotFound with default headers values
func NewAddClusterNotFound() *AddClusterNotFound {
	return &AddClusterNotFound{}
}

/* AddClusterNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AddClusterNotFound struct {
}

func (o *AddClusterNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] addClusterNotFound ", 404)
}

func (o *AddClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddClusterMethodNotAllowed creates a AddClusterMethodNotAllowed with default headers values
func NewAddClusterMethodNotAllowed() *AddClusterMethodNotAllowed {
	return &AddClusterMethodNotAllowed{}
}

/* AddClusterMethodNotAllowed describes a response with status code 405, with default header values.

Invalid input
*/
type AddClusterMethodNotAllowed struct {
}

func (o *AddClusterMethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] addClusterMethodNotAllowed ", 405)
}

func (o *AddClusterMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddClusterConflict creates a AddClusterConflict with default headers values
func NewAddClusterConflict() *AddClusterConflict {
	return &AddClusterConflict{}
}

/* AddClusterConflict describes a response with status code 409, with default header values.

Cluster with that name already exists
*/
type AddClusterConflict struct {
	Payload *models.ErrorResponse
}

func (o *AddClusterConflict) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] addClusterConflict  %+v", 409, o.Payload)
}
func (o *AddClusterConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddClusterConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddClusterUnprocessableEntity creates a AddClusterUnprocessableEntity with default headers values
func NewAddClusterUnprocessableEntity() *AddClusterUnprocessableEntity {
	return &AddClusterUnprocessableEntity{}
}

/* AddClusterUnprocessableEntity describes a response with status code 422, with default header values.

Error adding cluster
*/
type AddClusterUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *AddClusterUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /v1/clusters][%d] addClusterUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *AddClusterUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AddClusterUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
