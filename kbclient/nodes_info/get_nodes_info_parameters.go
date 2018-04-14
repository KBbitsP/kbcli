// Code generated by go-swagger; DO NOT EDIT.

package nodes_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetNodesInfoParams creates a new GetNodesInfoParams object
// with the default values initialized.
func NewGetNodesInfoParams() *GetNodesInfoParams {

	return &GetNodesInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodesInfoParamsWithTimeout creates a new GetNodesInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodesInfoParamsWithTimeout(timeout time.Duration) *GetNodesInfoParams {

	return &GetNodesInfoParams{

		timeout: timeout,
	}
}

// NewGetNodesInfoParamsWithContext creates a new GetNodesInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNodesInfoParamsWithContext(ctx context.Context) *GetNodesInfoParams {

	return &GetNodesInfoParams{

		Context: ctx,
	}
}

// NewGetNodesInfoParamsWithHTTPClient creates a new GetNodesInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNodesInfoParamsWithHTTPClient(client *http.Client) *GetNodesInfoParams {

	return &GetNodesInfoParams{
		HTTPClient: client,
	}
}

/*GetNodesInfoParams contains all the parameters to send to the API endpoint
for the get nodes info operation typically these are written to a http.Request
*/
type GetNodesInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get nodes info params
func (o *GetNodesInfoParams) WithTimeout(timeout time.Duration) *GetNodesInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get nodes info params
func (o *GetNodesInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get nodes info params
func (o *GetNodesInfoParams) WithContext(ctx context.Context) *GetNodesInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get nodes info params
func (o *GetNodesInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get nodes info params
func (o *GetNodesInfoParams) WithHTTPClient(client *http.Client) *GetNodesInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get nodes info params
func (o *GetNodesInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
