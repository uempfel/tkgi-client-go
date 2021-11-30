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

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCreateUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateUserOK creates a CreateUserOK with default headers values
func NewCreateUserOK() *CreateUserOK {
	return &CreateUserOK{}
}

/* CreateUserOK describes a response with status code 200, with default header values.

OK
*/
type CreateUserOK struct {
	Payload interface{}
}

func (o *CreateUserOK) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterName}/binds][%d] createUserOK  %+v", 200, o.Payload)
}
func (o *CreateUserOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/* CreateUserCreated describes a response with status code 201, with default header values.

User created
*/
type CreateUserCreated struct {
	Payload interface{}
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterName}/binds][%d] createUserCreated  %+v", 201, o.Payload)
}
func (o *CreateUserCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserUnauthorized creates a CreateUserUnauthorized with default headers values
func NewCreateUserUnauthorized() *CreateUserUnauthorized {
	return &CreateUserUnauthorized{}
}

/* CreateUserUnauthorized describes a response with status code 401, with default header values.

Authentication failure. Please check credentials and try again.
*/
type CreateUserUnauthorized struct {
}

func (o *CreateUserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterName}/binds][%d] createUserUnauthorized ", 401)
}

func (o *CreateUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserForbidden creates a CreateUserForbidden with default headers values
func NewCreateUserForbidden() *CreateUserForbidden {
	return &CreateUserForbidden{}
}

/* CreateUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateUserForbidden struct {
}

func (o *CreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterName}/binds][%d] createUserForbidden ", 403)
}

func (o *CreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateUserNotFound creates a CreateUserNotFound with default headers values
func NewCreateUserNotFound() *CreateUserNotFound {
	return &CreateUserNotFound{}
}

/* CreateUserNotFound describes a response with status code 404, with default header values.

Cluster not found
*/
type CreateUserNotFound struct {
	Payload *models.ErrorResponse
}

func (o *CreateUserNotFound) Error() string {
	return fmt.Sprintf("[POST /v1/clusters/{clusterName}/binds][%d] createUserNotFound  %+v", 404, o.Payload)
}
func (o *CreateUserNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CreateUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}