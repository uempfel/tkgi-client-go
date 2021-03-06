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

// NewGetQuotaParams creates a new GetQuotaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQuotaParams() *GetQuotaParams {
	return &GetQuotaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQuotaParamsWithTimeout creates a new GetQuotaParams object
// with the ability to set a timeout on a request.
func NewGetQuotaParamsWithTimeout(timeout time.Duration) *GetQuotaParams {
	return &GetQuotaParams{
		timeout: timeout,
	}
}

// NewGetQuotaParamsWithContext creates a new GetQuotaParams object
// with the ability to set a context for a request.
func NewGetQuotaParamsWithContext(ctx context.Context) *GetQuotaParams {
	return &GetQuotaParams{
		Context: ctx,
	}
}

// NewGetQuotaParamsWithHTTPClient creates a new GetQuotaParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQuotaParamsWithHTTPClient(client *http.Client) *GetQuotaParams {
	return &GetQuotaParams{
		HTTPClient: client,
	}
}

/* GetQuotaParams contains all the parameters to send to the API endpoint
   for the get quota operation.

   Typically these are written to a http.Request.
*/
type GetQuotaParams struct {

	/* Owner.

	   The owner name
	*/
	Owner string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQuotaParams) WithDefaults() *GetQuotaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQuotaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get quota params
func (o *GetQuotaParams) WithTimeout(timeout time.Duration) *GetQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quota params
func (o *GetQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quota params
func (o *GetQuotaParams) WithContext(ctx context.Context) *GetQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quota params
func (o *GetQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quota params
func (o *GetQuotaParams) WithHTTPClient(client *http.Client) *GetQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quota params
func (o *GetQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOwner adds the owner to the get quota params
func (o *GetQuotaParams) WithOwner(owner string) *GetQuotaParams {
	o.SetOwner(owner)
	return o
}

// SetOwner adds the owner to the get quota params
func (o *GetQuotaParams) SetOwner(owner string) {
	o.Owner = owner
}

// WriteToRequest writes these params to a swagger request
func (o *GetQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param owner
	if err := r.SetPathParam("owner", o.Owner); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
