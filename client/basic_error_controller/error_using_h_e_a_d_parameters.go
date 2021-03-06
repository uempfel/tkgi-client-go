// Code generated by go-swagger; DO NOT EDIT.

package basic_error_controller

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

// NewErrorUsingHEADParams creates a new ErrorUsingHEADParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewErrorUsingHEADParams() *ErrorUsingHEADParams {
	return &ErrorUsingHEADParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewErrorUsingHEADParamsWithTimeout creates a new ErrorUsingHEADParams object
// with the ability to set a timeout on a request.
func NewErrorUsingHEADParamsWithTimeout(timeout time.Duration) *ErrorUsingHEADParams {
	return &ErrorUsingHEADParams{
		timeout: timeout,
	}
}

// NewErrorUsingHEADParamsWithContext creates a new ErrorUsingHEADParams object
// with the ability to set a context for a request.
func NewErrorUsingHEADParamsWithContext(ctx context.Context) *ErrorUsingHEADParams {
	return &ErrorUsingHEADParams{
		Context: ctx,
	}
}

// NewErrorUsingHEADParamsWithHTTPClient creates a new ErrorUsingHEADParams object
// with the ability to set a custom HTTPClient for a request.
func NewErrorUsingHEADParamsWithHTTPClient(client *http.Client) *ErrorUsingHEADParams {
	return &ErrorUsingHEADParams{
		HTTPClient: client,
	}
}

/* ErrorUsingHEADParams contains all the parameters to send to the API endpoint
   for the error using h e a d operation.

   Typically these are written to a http.Request.
*/
type ErrorUsingHEADParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the error using h e a d params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ErrorUsingHEADParams) WithDefaults() *ErrorUsingHEADParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the error using h e a d params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ErrorUsingHEADParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the error using h e a d params
func (o *ErrorUsingHEADParams) WithTimeout(timeout time.Duration) *ErrorUsingHEADParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the error using h e a d params
func (o *ErrorUsingHEADParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the error using h e a d params
func (o *ErrorUsingHEADParams) WithContext(ctx context.Context) *ErrorUsingHEADParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the error using h e a d params
func (o *ErrorUsingHEADParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the error using h e a d params
func (o *ErrorUsingHEADParams) WithHTTPClient(client *http.Client) *ErrorUsingHEADParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the error using h e a d params
func (o *ErrorUsingHEADParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ErrorUsingHEADParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
