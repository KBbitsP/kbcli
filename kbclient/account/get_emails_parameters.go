// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewGetEmailsParams creates a new GetEmailsParams object
// with the default values initialized.
func NewGetEmailsParams() *GetEmailsParams {
	var ()
	return &GetEmailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmailsParamsWithTimeout creates a new GetEmailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmailsParamsWithTimeout(timeout time.Duration) *GetEmailsParams {
	var ()
	return &GetEmailsParams{

		timeout: timeout,
	}
}

// NewGetEmailsParamsWithContext creates a new GetEmailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmailsParamsWithContext(ctx context.Context) *GetEmailsParams {
	var ()
	return &GetEmailsParams{

		Context: ctx,
	}
}

// NewGetEmailsParamsWithHTTPClient creates a new GetEmailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmailsParamsWithHTTPClient(client *http.Client) *GetEmailsParams {
	var ()
	return &GetEmailsParams{
		HTTPClient: client,
	}
}

/*GetEmailsParams contains all the parameters to send to the API endpoint
for the get emails operation typically these are written to a http.Request
*/
type GetEmailsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*AccountID*/
	AccountID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get emails params
func (o *GetEmailsParams) WithTimeout(timeout time.Duration) *GetEmailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get emails params
func (o *GetEmailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get emails params
func (o *GetEmailsParams) WithContext(ctx context.Context) *GetEmailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get emails params
func (o *GetEmailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get emails params
func (o *GetEmailsParams) WithHTTPClient(client *http.Client) *GetEmailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get emails params
func (o *GetEmailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get emails params
func (o *GetEmailsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetEmailsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get emails params
func (o *GetEmailsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get emails params
func (o *GetEmailsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetEmailsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get emails params
func (o *GetEmailsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAccountID adds the accountID to the get emails params
func (o *GetEmailsParams) WithAccountID(accountID strfmt.UUID) *GetEmailsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get emails params
func (o *GetEmailsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
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
