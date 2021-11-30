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

// NewHandleUsingGET3Params creates a new HandleUsingGET3Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHandleUsingGET3Params() *HandleUsingGET3Params {
	return &HandleUsingGET3Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewHandleUsingGET3ParamsWithTimeout creates a new HandleUsingGET3Params object
// with the ability to set a timeout on a request.
func NewHandleUsingGET3ParamsWithTimeout(timeout time.Duration) *HandleUsingGET3Params {
	return &HandleUsingGET3Params{
		timeout: timeout,
	}
}

// NewHandleUsingGET3ParamsWithContext creates a new HandleUsingGET3Params object
// with the ability to set a context for a request.
func NewHandleUsingGET3ParamsWithContext(ctx context.Context) *HandleUsingGET3Params {
	return &HandleUsingGET3Params{
		Context: ctx,
	}
}

// NewHandleUsingGET3ParamsWithHTTPClient creates a new HandleUsingGET3Params object
// with the ability to set a custom HTTPClient for a request.
func NewHandleUsingGET3ParamsWithHTTPClient(client *http.Client) *HandleUsingGET3Params {
	return &HandleUsingGET3Params{
		HTTPClient: client,
	}
}

/* HandleUsingGET3Params contains all the parameters to send to the API endpoint
   for the handle using g e t 3 operation.

   Typically these are written to a http.Request.
*/
type HandleUsingGET3Params struct {

	/* Body.

	   body
	*/
	Body map[string]string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the handle using g e t 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleUsingGET3Params) WithDefaults() *HandleUsingGET3Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the handle using g e t 3 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HandleUsingGET3Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the handle using g e t 3 params
func (o *HandleUsingGET3Params) WithTimeout(timeout time.Duration) *HandleUsingGET3Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the handle using g e t 3 params
func (o *HandleUsingGET3Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the handle using g e t 3 params
func (o *HandleUsingGET3Params) WithContext(ctx context.Context) *HandleUsingGET3Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the handle using g e t 3 params
func (o *HandleUsingGET3Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the handle using g e t 3 params
func (o *HandleUsingGET3Params) WithHTTPClient(client *http.Client) *HandleUsingGET3Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the handle using g e t 3 params
func (o *HandleUsingGET3Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the handle using g e t 3 params
func (o *HandleUsingGET3Params) WithBody(body map[string]string) *HandleUsingGET3Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the handle using g e t 3 params
func (o *HandleUsingGET3Params) SetBody(body map[string]string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *HandleUsingGET3Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
