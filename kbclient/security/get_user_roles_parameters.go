// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewGetUserRolesParams creates a new GetUserRolesParams object
// with the default values initialized.
func NewGetUserRolesParams() *GetUserRolesParams {
	var ()
	return &GetUserRolesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserRolesParamsWithTimeout creates a new GetUserRolesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserRolesParamsWithTimeout(timeout time.Duration) *GetUserRolesParams {
	var ()
	return &GetUserRolesParams{

		timeout: timeout,
	}
}

// NewGetUserRolesParamsWithContext creates a new GetUserRolesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserRolesParamsWithContext(ctx context.Context) *GetUserRolesParams {
	var ()
	return &GetUserRolesParams{

		Context: ctx,
	}
}

// NewGetUserRolesParamsWithHTTPClient creates a new GetUserRolesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserRolesParamsWithHTTPClient(client *http.Client) *GetUserRolesParams {
	var ()
	return &GetUserRolesParams{
		HTTPClient: client,
	}
}

/*GetUserRolesParams contains all the parameters to send to the API endpoint
for the get user roles operation typically these are written to a http.Request
*/
type GetUserRolesParams struct {

	/*Username*/
	Username string

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get user roles params
func (o *GetUserRolesParams) WithTimeout(timeout time.Duration) *GetUserRolesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user roles params
func (o *GetUserRolesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user roles params
func (o *GetUserRolesParams) WithContext(ctx context.Context) *GetUserRolesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user roles params
func (o *GetUserRolesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user roles params
func (o *GetUserRolesParams) WithHTTPClient(client *http.Client) *GetUserRolesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user roles params
func (o *GetUserRolesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUsername adds the username to the get user roles params
func (o *GetUserRolesParams) WithUsername(username string) *GetUserRolesParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the get user roles params
func (o *GetUserRolesParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserRolesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
