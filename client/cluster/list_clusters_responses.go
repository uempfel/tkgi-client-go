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

// ListClustersReader is a Reader for the ListClusters structure.
type ListClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListClustersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListClustersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListClustersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListClustersOK creates a ListClustersOK with default headers values
func NewListClustersOK() *ListClustersOK {
	return &ListClustersOK{}
}

/* ListClustersOK describes a response with status code 200, with default header values.

Listed successfully
*/
type ListClustersOK struct {
	Payload []*models.Cluster
}

func (o *ListClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/clusters][%d] listClustersOK  %+v", 200, o.Payload)
}
func (o *ListClustersOK) GetPayload() []*models.Cluster {
	return o.Payload
}

func (o *ListClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListClustersUnauthorized creates a ListClustersUnauthorized with default headers values
func NewListClustersUnauthorized() *ListClustersUnauthorized {
	return &ListClustersUnauthorized{}
}

/* ListClustersUnauthorized describes a response with status code 401, with default header values.

Authentication failure. Please check credentials and try again.
*/
type ListClustersUnauthorized struct {
}

func (o *ListClustersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/clusters][%d] listClustersUnauthorized ", 401)
}

func (o *ListClustersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClustersForbidden creates a ListClustersForbidden with default headers values
func NewListClustersForbidden() *ListClustersForbidden {
	return &ListClustersForbidden{}
}

/* ListClustersForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListClustersForbidden struct {
}

func (o *ListClustersForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/clusters][%d] listClustersForbidden ", 403)
}

func (o *ListClustersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewListClustersNotFound creates a ListClustersNotFound with default headers values
func NewListClustersNotFound() *ListClustersNotFound {
	return &ListClustersNotFound{}
}

/* ListClustersNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ListClustersNotFound struct {
}

func (o *ListClustersNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/clusters][%d] listClustersNotFound ", 404)
}

func (o *ListClustersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}