// Code generated by go-swagger; DO NOT EDIT.

package profile

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListKubernetesProfilesParams creates a new ListKubernetesProfilesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListKubernetesProfilesParams() *ListKubernetesProfilesParams {
	return &ListKubernetesProfilesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListKubernetesProfilesParamsWithTimeout creates a new ListKubernetesProfilesParams object
// with the ability to set a timeout on a request.
func NewListKubernetesProfilesParamsWithTimeout(timeout time.Duration) *ListKubernetesProfilesParams {
	return &ListKubernetesProfilesParams{
		timeout: timeout,
	}
}

// NewListKubernetesProfilesParamsWithContext creates a new ListKubernetesProfilesParams object
// with the ability to set a context for a request.
func NewListKubernetesProfilesParamsWithContext(ctx context.Context) *ListKubernetesProfilesParams {
	return &ListKubernetesProfilesParams{
		Context: ctx,
	}
}

// NewListKubernetesProfilesParamsWithHTTPClient creates a new ListKubernetesProfilesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListKubernetesProfilesParamsWithHTTPClient(client *http.Client) *ListKubernetesProfilesParams {
	return &ListKubernetesProfilesParams{
		HTTPClient: client,
	}
}

/* ListKubernetesProfilesParams contains all the parameters to send to the API endpoint
   for the list kubernetes profiles operation.

   Typically these are written to a http.Request.
*/
type ListKubernetesProfilesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list kubernetes profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListKubernetesProfilesParams) WithDefaults() *ListKubernetesProfilesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list kubernetes profiles params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListKubernetesProfilesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list kubernetes profiles params
func (o *ListKubernetesProfilesParams) WithTimeout(timeout time.Duration) *ListKubernetesProfilesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list kubernetes profiles params
func (o *ListKubernetesProfilesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list kubernetes profiles params
func (o *ListKubernetesProfilesParams) WithContext(ctx context.Context) *ListKubernetesProfilesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list kubernetes profiles params
func (o *ListKubernetesProfilesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list kubernetes profiles params
func (o *ListKubernetesProfilesParams) WithHTTPClient(client *http.Client) *ListKubernetesProfilesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list kubernetes profiles params
func (o *ListKubernetesProfilesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListKubernetesProfilesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
