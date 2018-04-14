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

// NewGetEmailNotificationsForAccountParams creates a new GetEmailNotificationsForAccountParams object
// with the default values initialized.
func NewGetEmailNotificationsForAccountParams() *GetEmailNotificationsForAccountParams {
	var ()
	return &GetEmailNotificationsForAccountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmailNotificationsForAccountParamsWithTimeout creates a new GetEmailNotificationsForAccountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEmailNotificationsForAccountParamsWithTimeout(timeout time.Duration) *GetEmailNotificationsForAccountParams {
	var ()
	return &GetEmailNotificationsForAccountParams{

		timeout: timeout,
	}
}

// NewGetEmailNotificationsForAccountParamsWithContext creates a new GetEmailNotificationsForAccountParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEmailNotificationsForAccountParamsWithContext(ctx context.Context) *GetEmailNotificationsForAccountParams {
	var ()
	return &GetEmailNotificationsForAccountParams{

		Context: ctx,
	}
}

// NewGetEmailNotificationsForAccountParamsWithHTTPClient creates a new GetEmailNotificationsForAccountParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEmailNotificationsForAccountParamsWithHTTPClient(client *http.Client) *GetEmailNotificationsForAccountParams {
	var ()
	return &GetEmailNotificationsForAccountParams{
		HTTPClient: client,
	}
}

/*GetEmailNotificationsForAccountParams contains all the parameters to send to the API endpoint
for the get email notifications for account operation typically these are written to a http.Request
*/
type GetEmailNotificationsForAccountParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*AccountID*/
	AccountID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) WithTimeout(timeout time.Duration) *GetEmailNotificationsForAccountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) WithContext(ctx context.Context) *GetEmailNotificationsForAccountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) WithHTTPClient(client *http.Client) *GetEmailNotificationsForAccountParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetEmailNotificationsForAccountParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetEmailNotificationsForAccountParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAccountID adds the accountID to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) WithAccountID(accountID strfmt.UUID) *GetEmailNotificationsForAccountParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get email notifications for account params
func (o *GetEmailNotificationsForAccountParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmailNotificationsForAccountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
