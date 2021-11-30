// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

	"github.com/uempfel/tkgi-client-go/models"
)

// NewAddClusterParams creates a new AddClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddClusterParams() *AddClusterParams {
	return &AddClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddClusterParamsWithTimeout creates a new AddClusterParams object
// with the ability to set a timeout on a request.
func NewAddClusterParamsWithTimeout(timeout time.Duration) *AddClusterParams {
	return &AddClusterParams{
		timeout: timeout,
	}
}

// NewAddClusterParamsWithContext creates a new AddClusterParams object
// with the ability to set a context for a request.
func NewAddClusterParamsWithContext(ctx context.Context) *AddClusterParams {
	return &AddClusterParams{
		Context: ctx,
	}
}

// NewAddClusterParamsWithHTTPClient creates a new AddClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddClusterParamsWithHTTPClient(client *http.Client) *AddClusterParams {
	return &AddClusterParams{
		HTTPClient: client,
	}
}

/* AddClusterParams contains all the parameters to send to the API endpoint
   for the add cluster operation.

   Typically these are written to a http.Request.
*/
type AddClusterParams struct {

	/* ClusterRequest.

	   Cluster info
	*/
	ClusterRequest *models.ClusterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddClusterParams) WithDefaults() *AddClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add cluster params
func (o *AddClusterParams) WithTimeout(timeout time.Duration) *AddClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add cluster params
func (o *AddClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add cluster params
func (o *AddClusterParams) WithContext(ctx context.Context) *AddClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add cluster params
func (o *AddClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add cluster params
func (o *AddClusterParams) WithHTTPClient(client *http.Client) *AddClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add cluster params
func (o *AddClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterRequest adds the clusterRequest to the add cluster params
func (o *AddClusterParams) WithClusterRequest(clusterRequest *models.ClusterRequest) *AddClusterParams {
	o.SetClusterRequest(clusterRequest)
	return o
}

// SetClusterRequest adds the clusterRequest to the add cluster params
func (o *AddClusterParams) SetClusterRequest(clusterRequest *models.ClusterRequest) {
	o.ClusterRequest = clusterRequest
}

// WriteToRequest writes these params to a swagger request
func (o *AddClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ClusterRequest != nil {
		if err := r.SetBodyParam(o.ClusterRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
