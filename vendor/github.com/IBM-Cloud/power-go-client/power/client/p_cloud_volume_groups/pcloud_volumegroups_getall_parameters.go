// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volume_groups

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

// NewPcloudVolumegroupsGetallParams creates a new PcloudVolumegroupsGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudVolumegroupsGetallParams() *PcloudVolumegroupsGetallParams {
	return &PcloudVolumegroupsGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudVolumegroupsGetallParamsWithTimeout creates a new PcloudVolumegroupsGetallParams object
// with the ability to set a timeout on a request.
func NewPcloudVolumegroupsGetallParamsWithTimeout(timeout time.Duration) *PcloudVolumegroupsGetallParams {
	return &PcloudVolumegroupsGetallParams{
		timeout: timeout,
	}
}

// NewPcloudVolumegroupsGetallParamsWithContext creates a new PcloudVolumegroupsGetallParams object
// with the ability to set a context for a request.
func NewPcloudVolumegroupsGetallParamsWithContext(ctx context.Context) *PcloudVolumegroupsGetallParams {
	return &PcloudVolumegroupsGetallParams{
		Context: ctx,
	}
}

// NewPcloudVolumegroupsGetallParamsWithHTTPClient creates a new PcloudVolumegroupsGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudVolumegroupsGetallParamsWithHTTPClient(client *http.Client) *PcloudVolumegroupsGetallParams {
	return &PcloudVolumegroupsGetallParams{
		HTTPClient: client,
	}
}

/* PcloudVolumegroupsGetallParams contains all the parameters to send to the API endpoint
   for the pcloud volumegroups getall operation.

   Typically these are written to a http.Request.
*/
type PcloudVolumegroupsGetallParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud volumegroups getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVolumegroupsGetallParams) WithDefaults() *PcloudVolumegroupsGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud volumegroups getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudVolumegroupsGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud volumegroups getall params
func (o *PcloudVolumegroupsGetallParams) WithTimeout(timeout time.Duration) *PcloudVolumegroupsGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud volumegroups getall params
func (o *PcloudVolumegroupsGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud volumegroups getall params
func (o *PcloudVolumegroupsGetallParams) WithContext(ctx context.Context) *PcloudVolumegroupsGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud volumegroups getall params
func (o *PcloudVolumegroupsGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud volumegroups getall params
func (o *PcloudVolumegroupsGetallParams) WithHTTPClient(client *http.Client) *PcloudVolumegroupsGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud volumegroups getall params
func (o *PcloudVolumegroupsGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud volumegroups getall params
func (o *PcloudVolumegroupsGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudVolumegroupsGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud volumegroups getall params
func (o *PcloudVolumegroupsGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudVolumegroupsGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}