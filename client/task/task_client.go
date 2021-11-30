// Code generated by go-swagger; DO NOT EDIT.

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new task API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for task API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelTask(params *CancelTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CancelTaskOK, *CancelTaskCreated, error)

	GetTask(params *GetTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTaskOK, error)

	ListTasks(params *ListTasksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTasksOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CancelTask cancels task
*/
func (a *Client) CancelTask(params *CancelTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CancelTaskOK, *CancelTaskCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "cancelTask",
		Method:             "PUT",
		PathPattern:        "/v1/tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CancelTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *CancelTaskOK:
		return value, nil, nil
	case *CancelTaskCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for task: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetTask gets task
*/
func (a *Client) GetTask(params *GetTaskParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetTaskOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetTaskParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getTask",
		Method:             "GET",
		PathPattern:        "/v1/tasks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTaskReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetTaskOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getTask: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListTasks lists all tasks
*/
func (a *Client) ListTasks(params *ListTasksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListTasksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListTasksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listTasks",
		Method:             "GET",
		PathPattern:        "/v1/tasks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListTasksReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListTasksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listTasks: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}