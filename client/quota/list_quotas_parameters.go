// Code generated by go-swagger; DO NOT EDIT.

package quota

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

// NewListQuotasParams creates a new ListQuotasParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListQuotasParams() *ListQuotasParams {
	return &ListQuotasParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListQuotasParamsWithTimeout creates a new ListQuotasParams object
// with the ability to set a timeout on a request.
func NewListQuotasParamsWithTimeout(timeout time.Duration) *ListQuotasParams {
	return &ListQuotasParams{
		timeout: timeout,
	}
}

// NewListQuotasParamsWithContext creates a new ListQuotasParams object
// with the ability to set a context for a request.
func NewListQuotasParamsWithContext(ctx context.Context) *ListQuotasParams {
	return &ListQuotasParams{
		Context: ctx,
	}
}

// NewListQuotasParamsWithHTTPClient creates a new ListQuotasParams object
// with the ability to set a custom HTTPClient for a request.
func NewListQuotasParamsWithHTTPClient(client *http.Client) *ListQuotasParams {
	return &ListQuotasParams{
		HTTPClient: client,
	}
}

/* ListQuotasParams contains all the parameters to send to the API endpoint
   for the list quotas operation.

   Typically these are written to a http.Request.
*/
type ListQuotasParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list quotas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListQuotasParams) WithDefaults() *ListQuotasParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list quotas params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListQuotasParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list quotas params
func (o *ListQuotasParams) WithTimeout(timeout time.Duration) *ListQuotasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list quotas params
func (o *ListQuotasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list quotas params
func (o *ListQuotasParams) WithContext(ctx context.Context) *ListQuotasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list quotas params
func (o *ListQuotasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list quotas params
func (o *ListQuotasParams) WithHTTPClient(client *http.Client) *ListQuotasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list quotas params
func (o *ListQuotasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListQuotasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
