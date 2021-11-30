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

// NewUpdateClusterParams creates a new UpdateClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateClusterParams() *UpdateClusterParams {
	return &UpdateClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateClusterParamsWithTimeout creates a new UpdateClusterParams object
// with the ability to set a timeout on a request.
func NewUpdateClusterParamsWithTimeout(timeout time.Duration) *UpdateClusterParams {
	return &UpdateClusterParams{
		timeout: timeout,
	}
}

// NewUpdateClusterParamsWithContext creates a new UpdateClusterParams object
// with the ability to set a context for a request.
func NewUpdateClusterParamsWithContext(ctx context.Context) *UpdateClusterParams {
	return &UpdateClusterParams{
		Context: ctx,
	}
}

// NewUpdateClusterParamsWithHTTPClient creates a new UpdateClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateClusterParamsWithHTTPClient(client *http.Client) *UpdateClusterParams {
	return &UpdateClusterParams{
		HTTPClient: client,
	}
}

/* UpdateClusterParams contains all the parameters to send to the API endpoint
   for the update cluster operation.

   Typically these are written to a http.Request.
*/
type UpdateClusterParams struct {

	/* ClusterName.

	   The cluster name
	*/
	ClusterName string

	/* UpdateClusterParameters.

	   Update parameters
	*/
	UpdateClusterParameters *models.UpdateClusterParameters

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClusterParams) WithDefaults() *UpdateClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update cluster params
func (o *UpdateClusterParams) WithTimeout(timeout time.Duration) *UpdateClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update cluster params
func (o *UpdateClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update cluster params
func (o *UpdateClusterParams) WithContext(ctx context.Context) *UpdateClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update cluster params
func (o *UpdateClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update cluster params
func (o *UpdateClusterParams) WithHTTPClient(client *http.Client) *UpdateClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update cluster params
func (o *UpdateClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterName adds the clusterName to the update cluster params
func (o *UpdateClusterParams) WithClusterName(clusterName string) *UpdateClusterParams {
	o.SetClusterName(clusterName)
	return o
}

// SetClusterName adds the clusterName to the update cluster params
func (o *UpdateClusterParams) SetClusterName(clusterName string) {
	o.ClusterName = clusterName
}

// WithUpdateClusterParameters adds the updateClusterParameters to the update cluster params
func (o *UpdateClusterParams) WithUpdateClusterParameters(updateClusterParameters *models.UpdateClusterParameters) *UpdateClusterParams {
	o.SetUpdateClusterParameters(updateClusterParameters)
	return o
}

// SetUpdateClusterParameters adds the updateClusterParameters to the update cluster params
func (o *UpdateClusterParams) SetUpdateClusterParameters(updateClusterParameters *models.UpdateClusterParameters) {
	o.UpdateClusterParameters = updateClusterParameters
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param clusterName
	if err := r.SetPathParam("clusterName", o.ClusterName); err != nil {
		return err
	}
	if o.UpdateClusterParameters != nil {
		if err := r.SetBodyParam(o.UpdateClusterParameters); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
