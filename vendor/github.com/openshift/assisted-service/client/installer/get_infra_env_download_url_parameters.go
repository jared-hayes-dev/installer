// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewGetInfraEnvDownloadURLParams creates a new GetInfraEnvDownloadURLParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetInfraEnvDownloadURLParams() *GetInfraEnvDownloadURLParams {
	return &GetInfraEnvDownloadURLParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetInfraEnvDownloadURLParamsWithTimeout creates a new GetInfraEnvDownloadURLParams object
// with the ability to set a timeout on a request.
func NewGetInfraEnvDownloadURLParamsWithTimeout(timeout time.Duration) *GetInfraEnvDownloadURLParams {
	return &GetInfraEnvDownloadURLParams{
		timeout: timeout,
	}
}

// NewGetInfraEnvDownloadURLParamsWithContext creates a new GetInfraEnvDownloadURLParams object
// with the ability to set a context for a request.
func NewGetInfraEnvDownloadURLParamsWithContext(ctx context.Context) *GetInfraEnvDownloadURLParams {
	return &GetInfraEnvDownloadURLParams{
		Context: ctx,
	}
}

// NewGetInfraEnvDownloadURLParamsWithHTTPClient creates a new GetInfraEnvDownloadURLParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetInfraEnvDownloadURLParamsWithHTTPClient(client *http.Client) *GetInfraEnvDownloadURLParams {
	return &GetInfraEnvDownloadURLParams{
		HTTPClient: client,
	}
}

/* GetInfraEnvDownloadURLParams contains all the parameters to send to the API endpoint
   for the get infra env download URL operation.

   Typically these are written to a http.Request.
*/
type GetInfraEnvDownloadURLParams struct {

	/* InfraEnvID.

	   The infra-env to be retrieved.

	   Format: uuid
	*/
	InfraEnvID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get infra env download URL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInfraEnvDownloadURLParams) WithDefaults() *GetInfraEnvDownloadURLParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get infra env download URL params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetInfraEnvDownloadURLParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get infra env download URL params
func (o *GetInfraEnvDownloadURLParams) WithTimeout(timeout time.Duration) *GetInfraEnvDownloadURLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get infra env download URL params
func (o *GetInfraEnvDownloadURLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get infra env download URL params
func (o *GetInfraEnvDownloadURLParams) WithContext(ctx context.Context) *GetInfraEnvDownloadURLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get infra env download URL params
func (o *GetInfraEnvDownloadURLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get infra env download URL params
func (o *GetInfraEnvDownloadURLParams) WithHTTPClient(client *http.Client) *GetInfraEnvDownloadURLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get infra env download URL params
func (o *GetInfraEnvDownloadURLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfraEnvID adds the infraEnvID to the get infra env download URL params
func (o *GetInfraEnvDownloadURLParams) WithInfraEnvID(infraEnvID strfmt.UUID) *GetInfraEnvDownloadURLParams {
	o.SetInfraEnvID(infraEnvID)
	return o
}

// SetInfraEnvID adds the infraEnvId to the get infra env download URL params
func (o *GetInfraEnvDownloadURLParams) SetInfraEnvID(infraEnvID strfmt.UUID) {
	o.InfraEnvID = infraEnvID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInfraEnvDownloadURLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param infra_env_id
	if err := r.SetPathParam("infra_env_id", o.InfraEnvID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
