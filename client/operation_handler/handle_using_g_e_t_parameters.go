// Code generated by go-swagger; DO NOT EDIT.

package operation_handler

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

// NewHandleUsingGETParams creates a new HandleUsingGETParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHandleUsingGETParams() *HandleUsingGETParams {
	return &HandleUsingGETParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHandleUsingGETParamsWithTimeout creates a new HandleUsingGETParams object
// with the ability to set a timeout on a request.
func NewHandleUsingGETParamsWithTimeout(timeout time.Duration) *HandleUsingGETParams {
	return &HandleUsingGETParams{
		timeout: timeout,
	}
}

// NewHandleUsingGETParamsWithContext creates a new HandleUsingGETParams object
// with the ability to set a context for a request.
func NewHandleUsingGETParamsWithContext(ctx context.Context) *HandleUsingGETParams {
	return &HandleUsingGETParams{
		Context: ctx,
	}
}

// NewHandleUsingGETParamsWithHTTPClient creates a new HandleUsingGETParams object
// with the ability to set a custom HTTPClient for a request.
func NewHandleUsingGETParamsWithHTTPClient(client *http.Client) *HandleUsingGETParams {
	return &HandleUsingGETParams{
		HTTPClient: client,
	}
}

/* HandleUsingGETParams contains all the parameters to send to the API endpoint
   for the handle using g e t operation.

   Typically these are written to a http.Request.
*/
type HandleUsingGETParams struct {

	/* Body.

	   body
	*/
	Body map[string]string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the handle using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleUsingGETParams) WithDefaults() *HandleUsingGETParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the handle using g e t params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleUsingGETParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the handle using g e t params
func (o *HandleUsingGETParams) WithTimeout(timeout time.Duration) *HandleUsingGETParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the handle using g e t params
func (o *HandleUsingGETParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the handle using g e t params
func (o *HandleUsingGETParams) WithContext(ctx context.Context) *HandleUsingGETParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the handle using g e t params
func (o *HandleUsingGETParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the handle using g e t params
func (o *HandleUsingGETParams) WithHTTPClient(client *http.Client) *HandleUsingGETParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the handle using g e t params
func (o *HandleUsingGETParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the handle using g e t params
func (o *HandleUsingGETParams) WithBody(body map[string]string) *HandleUsingGETParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the handle using g e t params
func (o *HandleUsingGETParams) SetBody(body map[string]string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *HandleUsingGETParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}