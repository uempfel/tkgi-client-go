// Code generated by go-swagger; DO NOT EDIT.

package cert

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

// NewListCertificatesParams creates a new ListCertificatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListCertificatesParams() *ListCertificatesParams {
	return &ListCertificatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListCertificatesParamsWithTimeout creates a new ListCertificatesParams object
// with the ability to set a timeout on a request.
func NewListCertificatesParamsWithTimeout(timeout time.Duration) *ListCertificatesParams {
	return &ListCertificatesParams{
		timeout: timeout,
	}
}

// NewListCertificatesParamsWithContext creates a new ListCertificatesParams object
// with the ability to set a context for a request.
func NewListCertificatesParamsWithContext(ctx context.Context) *ListCertificatesParams {
	return &ListCertificatesParams{
		Context: ctx,
	}
}

// NewListCertificatesParamsWithHTTPClient creates a new ListCertificatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListCertificatesParamsWithHTTPClient(client *http.Client) *ListCertificatesParams {
	return &ListCertificatesParams{
		HTTPClient: client,
	}
}

/* ListCertificatesParams contains all the parameters to send to the API endpoint
   for the list certificates operation.

   Typically these are written to a http.Request.
*/
type ListCertificatesParams struct {

	/* ClusterName.

	   The cluster name
	*/
	ClusterName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list certificates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCertificatesParams) WithDefaults() *ListCertificatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list certificates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListCertificatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list certificates params
func (o *ListCertificatesParams) WithTimeout(timeout time.Duration) *ListCertificatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list certificates params
func (o *ListCertificatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list certificates params
func (o *ListCertificatesParams) WithContext(ctx context.Context) *ListCertificatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list certificates params
func (o *ListCertificatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list certificates params
func (o *ListCertificatesParams) WithHTTPClient(client *http.Client) *ListCertificatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list certificates params
func (o *ListCertificatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterName adds the clusterName to the list certificates params
func (o *ListCertificatesParams) WithClusterName(clusterName string) *ListCertificatesParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the list certificates params
func (o *ListCertificatesParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WriteToRequest writes these params to a swagger request
func (o *ListCertificatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}