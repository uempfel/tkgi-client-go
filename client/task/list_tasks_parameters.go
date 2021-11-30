// Code generated by go-swagger; DO NOT EDIT.

package task

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

// NewListTasksParams creates a new ListTasksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListTasksParams() *ListTasksParams {
	return &ListTasksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListTasksParamsWithTimeout creates a new ListTasksParams object
// with the ability to set a timeout on a request.
func NewListTasksParamsWithTimeout(timeout time.Duration) *ListTasksParams {
	return &ListTasksParams{
		timeout: timeout,
	}
}

// NewListTasksParamsWithContext creates a new ListTasksParams object
// with the ability to set a context for a request.
func NewListTasksParamsWithContext(ctx context.Context) *ListTasksParams {
	return &ListTasksParams{
		Context: ctx,
	}
}

// NewListTasksParamsWithHTTPClient creates a new ListTasksParams object
// with the ability to set a custom HTTPClient for a request.
func NewListTasksParamsWithHTTPClient(client *http.Client) *ListTasksParams {
	return &ListTasksParams{
		HTTPClient: client,
	}
}

/* ListTasksParams contains all the parameters to send to the API endpoint
   for the list tasks operation.

   Typically these are written to a http.Request.
*/
type ListTasksParams struct {

	/* Limit.

	   Number of tasks to be returned

	   Default: "10"
	*/
	Limit *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTasksParams) WithDefaults() *ListTasksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list tasks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListTasksParams) SetDefaults() {
	var (
		limitDefault = string("10")
	)

	val := ListTasksParams{
		Limit: &limitDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the list tasks params
func (o *ListTasksParams) WithTimeout(timeout time.Duration) *ListTasksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list tasks params
func (o *ListTasksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list tasks params
func (o *ListTasksParams) WithContext(ctx context.Context) *ListTasksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list tasks params
func (o *ListTasksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list tasks params
func (o *ListTasksParams) WithHTTPClient(client *http.Client) *ListTasksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list tasks params
func (o *ListTasksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list tasks params
func (o *ListTasksParams) WithLimit(limit *string) *ListTasksParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list tasks params
func (o *ListTasksParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *ListTasksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit string

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
