// Code generated by go-swagger; DO NOT EDIT.

package admin

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

// NewPutOutOfRotationParams creates a new PutOutOfRotationParams object
// with the default values initialized.
func NewPutOutOfRotationParams() *PutOutOfRotationParams {
	var ()
	return &PutOutOfRotationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutOutOfRotationParamsWithTimeout creates a new PutOutOfRotationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutOutOfRotationParamsWithTimeout(timeout time.Duration) *PutOutOfRotationParams {
	var ()
	return &PutOutOfRotationParams{

		timeout: timeout,
	}
}

// NewPutOutOfRotationParamsWithContext creates a new PutOutOfRotationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutOutOfRotationParamsWithContext(ctx context.Context) *PutOutOfRotationParams {
	var ()
	return &PutOutOfRotationParams{

		Context: ctx,
	}
}

// NewPutOutOfRotationParamsWithHTTPClient creates a new PutOutOfRotationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutOutOfRotationParamsWithHTTPClient(client *http.Client) *PutOutOfRotationParams {
	var ()
	return &PutOutOfRotationParams{
		HTTPClient: client,
	}
}

/*PutOutOfRotationParams contains all the parameters to send to the API endpoint
for the put out of rotation operation typically these are written to a http.Request
*/
type PutOutOfRotationParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put out of rotation params
func (o *PutOutOfRotationParams) WithTimeout(timeout time.Duration) *PutOutOfRotationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put out of rotation params
func (o *PutOutOfRotationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put out of rotation params
func (o *PutOutOfRotationParams) WithContext(ctx context.Context) *PutOutOfRotationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put out of rotation params
func (o *PutOutOfRotationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put out of rotation params
func (o *PutOutOfRotationParams) WithHTTPClient(client *http.Client) *PutOutOfRotationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put out of rotation params
func (o *PutOutOfRotationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the put out of rotation params
func (o *PutOutOfRotationParams) WithXKillbillAPIKey(xKillbillAPIKey string) *PutOutOfRotationParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the put out of rotation params
func (o *PutOutOfRotationParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the put out of rotation params
func (o *PutOutOfRotationParams) WithXKillbillAPISecret(xKillbillAPISecret string) *PutOutOfRotationParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the put out of rotation params
func (o *PutOutOfRotationParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WriteToRequest writes these params to a swagger request
func (o *PutOutOfRotationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Killbill-ApiKey
	if err := r.SetHeaderParam("X-Killbill-ApiKey", o.XKillbillAPIKey); err != nil {
		return err
	}

	// header param X-Killbill-ApiSecret
	if err := r.SetHeaderParam("X-Killbill-ApiSecret", o.XKillbillAPISecret); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
