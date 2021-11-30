// Code generated by go-swagger; DO NOT EDIT.

package upgradable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CheckUpgradableReader is a Reader for the CheckUpgradable structure.
type CheckUpgradableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CheckUpgradableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCheckUpgradableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewCheckUpgradableUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCheckUpgradableForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCheckUpgradableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCheckUpgradableOK creates a CheckUpgradableOK with default headers values
func NewCheckUpgradableOK() *CheckUpgradableOK {
	return &CheckUpgradableOK{}
}

/* CheckUpgradableOK describes a response with status code 200, with default header values.

Check if all clusters are upgraded to the most recent version and PKS API is upgradable
*/
type CheckUpgradableOK struct {
	Payload string
}

func (o *CheckUpgradableOK) Error() string {
	return fmt.Sprintf("[GET /v1/upgradable][%d] checkUpgradableOK  %+v", 200, o.Payload)
}
func (o *CheckUpgradableOK) GetPayload() string {
	return o.Payload
}

func (o *CheckUpgradableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCheckUpgradableUnauthorized creates a CheckUpgradableUnauthorized with default headers values
func NewCheckUpgradableUnauthorized() *CheckUpgradableUnauthorized {
	return &CheckUpgradableUnauthorized{}
}

/* CheckUpgradableUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CheckUpgradableUnauthorized struct {
}

func (o *CheckUpgradableUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/upgradable][%d] checkUpgradableUnauthorized ", 401)
}

func (o *CheckUpgradableUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCheckUpgradableForbidden creates a CheckUpgradableForbidden with default headers values
func NewCheckUpgradableForbidden() *CheckUpgradableForbidden {
	return &CheckUpgradableForbidden{}
}

/* CheckUpgradableForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CheckUpgradableForbidden struct {
}

func (o *CheckUpgradableForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/upgradable][%d] checkUpgradableForbidden ", 403)
}

func (o *CheckUpgradableForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCheckUpgradableNotFound creates a CheckUpgradableNotFound with default headers values
func NewCheckUpgradableNotFound() *CheckUpgradableNotFound {
	return &CheckUpgradableNotFound{}
}

/* CheckUpgradableNotFound describes a response with status code 404, with default header values.

Not Found
*/
type CheckUpgradableNotFound struct {
}

func (o *CheckUpgradableNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/upgradable][%d] checkUpgradableNotFound ", 404)
}

func (o *CheckUpgradableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
