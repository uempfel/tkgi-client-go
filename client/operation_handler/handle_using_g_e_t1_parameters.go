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

// NewHandleUsingGET1Params creates a new HandleUsingGET1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHandleUsingGET1Params() *HandleUsingGET1Params {
	return &HandleUsingGET1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewHandleUsingGET1ParamsWithTimeout creates a new HandleUsingGET1Params object
// with the ability to set a timeout on a request.
func NewHandleUsingGET1ParamsWithTimeout(timeout time.Duration) *HandleUsingGET1Params {
	return &HandleUsingGET1Params{
		timeout: timeout,
	}
}

// NewHandleUsingGET1ParamsWithContext creates a new HandleUsingGET1Params object
// with the ability to set a context for a request.
func NewHandleUsingGET1ParamsWithContext(ctx context.Context) *HandleUsingGET1Params {
	return &HandleUsingGET1Params{
		Context: ctx,
	}
}

// NewHandleUsingGET1ParamsWithHTTPClient creates a new HandleUsingGET1Params object
// with the ability to set a custom HTTPClient for a request.
func NewHandleUsingGET1ParamsWithHTTPClient(client *http.Client) *HandleUsingGET1Params {
	return &HandleUsingGET1Params{
		HTTPClient: client,
	}
}

/* HandleUsingGET1Params contains all the parameters to send to the API endpoint
   for the handle using g e t 1 operation.

   Typically these are written to a http.Request.
*/
type HandleUsingGET1Params struct {

	/* Body.

	   body
	*/
	Body map[string]string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the handle using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleUsingGET1Params) WithDefaults() *HandleUsingGET1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the handle using g e t 1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleUsingGET1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the handle using g e t 1 params
func (o *HandleUsingGET1Params) WithTimeout(timeout time.Duration) *HandleUsingGET1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the handle using g e t 1 params
func (o *HandleUsingGET1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the handle using g e t 1 params
func (o *HandleUsingGET1Params) WithContext(ctx context.Context) *HandleUsingGET1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the handle using g e t 1 params
func (o *HandleUsingGET1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the handle using g e t 1 params
func (o *HandleUsingGET1Params) WithHTTPClient(client *http.Client) *HandleUsingGET1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the handle using g e t 1 params
func (o *HandleUsingGET1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the handle using g e t 1 params
func (o *HandleUsingGET1Params) WithBody(body map[string]string) *HandleUsingGET1Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the handle using g e t 1 params
func (o *HandleUsingGET1Params) SetBody(body map[string]string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *HandleUsingGET1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
