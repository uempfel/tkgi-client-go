// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new profile API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for profile API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AddComputeProfile(params *AddComputeProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddComputeProfileOK, *AddComputeProfileCreated, error)

	AddKubernetesProfile(params *AddKubernetesProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddKubernetesProfileOK, *AddKubernetesProfileCreated, error)

	AddNetworkProfile(params *AddNetworkProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddNetworkProfileOK, *AddNetworkProfileCreated, error)

	DeleteComputeProfile(params *DeleteComputeProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteComputeProfileOK, *DeleteComputeProfileNoContent, error)

	DeleteKubernetesProfile(params *DeleteKubernetesProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteKubernetesProfileOK, *DeleteKubernetesProfileNoContent, error)

	DeleteNetworkProfile(params *DeleteNetworkProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNetworkProfileOK, *DeleteNetworkProfileNoContent, error)

	GetComputeProfile(params *GetComputeProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComputeProfileOK, error)

	GetKubernetesProfile(params *GetKubernetesProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKubernetesProfileOK, error)

	GetNetworkProfile(params *GetNetworkProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNetworkProfileOK, error)

	ListComputeProfiles(params *ListComputeProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListComputeProfilesOK, error)

	ListKubernetesProfiles(params *ListKubernetesProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubernetesProfilesOK, error)

	ListNetworkProfiles(params *ListNetworkProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNetworkProfilesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AddComputeProfile creates a new compute profile
*/
func (a *Client) AddComputeProfile(params *AddComputeProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddComputeProfileOK, *AddComputeProfileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddComputeProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addComputeProfile",
		Method:             "POST",
		PathPattern:        "/v1/compute-profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddComputeProfileReader{formats: a.formats},
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
	case *AddComputeProfileOK:
		return value, nil, nil
	case *AddComputeProfileCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for profile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddKubernetesProfile creates a new kubernetes profile
*/
func (a *Client) AddKubernetesProfile(params *AddKubernetesProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddKubernetesProfileOK, *AddKubernetesProfileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddKubernetesProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addKubernetesProfile",
		Method:             "POST",
		PathPattern:        "/v1/kubernetes-profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddKubernetesProfileReader{formats: a.formats},
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
	case *AddKubernetesProfileOK:
		return value, nil, nil
	case *AddKubernetesProfileCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for profile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  AddNetworkProfile creates a new network profile
*/
func (a *Client) AddNetworkProfile(params *AddNetworkProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AddNetworkProfileOK, *AddNetworkProfileCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddNetworkProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "addNetworkProfile",
		Method:             "POST",
		PathPattern:        "/v1/network-profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AddNetworkProfileReader{formats: a.formats},
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
	case *AddNetworkProfileOK:
		return value, nil, nil
	case *AddNetworkProfileCreated:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for profile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteComputeProfile deletes compute profile
*/
func (a *Client) DeleteComputeProfile(params *DeleteComputeProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteComputeProfileOK, *DeleteComputeProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteComputeProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteComputeProfile",
		Method:             "DELETE",
		PathPattern:        "/v1/compute-profiles/{profileName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteComputeProfileReader{formats: a.formats},
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
	case *DeleteComputeProfileOK:
		return value, nil, nil
	case *DeleteComputeProfileNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for profile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteKubernetesProfile deletes kubernetes profile
*/
func (a *Client) DeleteKubernetesProfile(params *DeleteKubernetesProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteKubernetesProfileOK, *DeleteKubernetesProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteKubernetesProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteKubernetesProfile",
		Method:             "DELETE",
		PathPattern:        "/v1/kubernetes-profiles/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteKubernetesProfileReader{formats: a.formats},
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
	case *DeleteKubernetesProfileOK:
		return value, nil, nil
	case *DeleteKubernetesProfileNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for profile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteNetworkProfile deletes network profile
*/
func (a *Client) DeleteNetworkProfile(params *DeleteNetworkProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNetworkProfileOK, *DeleteNetworkProfileNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteNetworkProfile",
		Method:             "DELETE",
		PathPattern:        "/v1/network-profiles/{profileName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNetworkProfileReader{formats: a.formats},
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
	case *DeleteNetworkProfileOK:
		return value, nil, nil
	case *DeleteNetworkProfileNoContent:
		return nil, value, nil
	}
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for profile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetComputeProfile gets compute profile
*/
func (a *Client) GetComputeProfile(params *GetComputeProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetComputeProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetComputeProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getComputeProfile",
		Method:             "GET",
		PathPattern:        "/v1/compute-profiles/{profileName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetComputeProfileReader{formats: a.formats},
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
	success, ok := result.(*GetComputeProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getComputeProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetKubernetesProfile gets kubernetes profile
*/
func (a *Client) GetKubernetesProfile(params *GetKubernetesProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetKubernetesProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetKubernetesProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getKubernetesProfile",
		Method:             "GET",
		PathPattern:        "/v1/kubernetes-profiles/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetKubernetesProfileReader{formats: a.formats},
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
	success, ok := result.(*GetKubernetesProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getKubernetesProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  GetNetworkProfile gets network profile
*/
func (a *Client) GetNetworkProfile(params *GetNetworkProfileParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetNetworkProfileOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetworkProfileParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getNetworkProfile",
		Method:             "GET",
		PathPattern:        "/v1/network-profiles/{profileName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetworkProfileReader{formats: a.formats},
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
	success, ok := result.(*GetNetworkProfileOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetworkProfile: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListComputeProfiles lists all compute profiles
*/
func (a *Client) ListComputeProfiles(params *ListComputeProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListComputeProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListComputeProfilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listComputeProfiles",
		Method:             "GET",
		PathPattern:        "/v1/compute-profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListComputeProfilesReader{formats: a.formats},
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
	success, ok := result.(*ListComputeProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listComputeProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListKubernetesProfiles lists all kubernetes profiles
*/
func (a *Client) ListKubernetesProfiles(params *ListKubernetesProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListKubernetesProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListKubernetesProfilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listKubernetesProfiles",
		Method:             "GET",
		PathPattern:        "/v1/kubernetes-profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListKubernetesProfilesReader{formats: a.formats},
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
	success, ok := result.(*ListKubernetesProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listKubernetesProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListNetworkProfiles lists all network profiles
*/
func (a *Client) ListNetworkProfiles(params *ListNetworkProfilesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNetworkProfilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNetworkProfilesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNetworkProfiles",
		Method:             "GET",
		PathPattern:        "/v1/network-profiles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ListNetworkProfilesReader{formats: a.formats},
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
	success, ok := result.(*ListNetworkProfilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listNetworkProfiles: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}