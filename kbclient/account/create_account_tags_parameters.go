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

// NewCreateAccountTagsParams creates a new CreateAccountTagsParams object
// with the default values initialized.
func NewCreateAccountTagsParams() *CreateAccountTagsParams {
	var ()
	return &CreateAccountTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccountTagsParamsWithTimeout creates a new CreateAccountTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAccountTagsParamsWithTimeout(timeout time.Duration) *CreateAccountTagsParams {
	var ()
	return &CreateAccountTagsParams{

		timeout: timeout,
	}
}

// NewCreateAccountTagsParamsWithContext creates a new CreateAccountTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAccountTagsParamsWithContext(ctx context.Context) *CreateAccountTagsParams {
	var ()
	return &CreateAccountTagsParams{

		Context: ctx,
	}
}

// NewCreateAccountTagsParamsWithHTTPClient creates a new CreateAccountTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAccountTagsParamsWithHTTPClient(client *http.Client) *CreateAccountTagsParams {
	var ()
	return &CreateAccountTagsParams{
		HTTPClient: client,
	}
}

/*CreateAccountTagsParams contains all the parameters to send to the API endpoint
for the create account tags operation typically these are written to a http.Request
*/
type CreateAccountTagsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*AccountID*/
	AccountID strfmt.UUID
	/*Body*/
	Body []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create account tags params
func (o *CreateAccountTagsParams) WithTimeout(timeout time.Duration) *CreateAccountTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create account tags params
func (o *CreateAccountTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create account tags params
func (o *CreateAccountTagsParams) WithContext(ctx context.Context) *CreateAccountTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create account tags params
func (o *CreateAccountTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create account tags params
func (o *CreateAccountTagsParams) WithHTTPClient(client *http.Client) *CreateAccountTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create account tags params
func (o *CreateAccountTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the create account tags params
func (o *CreateAccountTagsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *CreateAccountTagsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the create account tags params
func (o *CreateAccountTagsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the create account tags params
func (o *CreateAccountTagsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *CreateAccountTagsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the create account tags params
func (o *CreateAccountTagsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the create account tags params
func (o *CreateAccountTagsParams) WithXKillbillComment(xKillbillComment *string) *CreateAccountTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create account tags params
func (o *CreateAccountTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create account tags params
func (o *CreateAccountTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateAccountTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create account tags params
func (o *CreateAccountTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create account tags params
func (o *CreateAccountTagsParams) WithXKillbillReason(xKillbillReason *string) *CreateAccountTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create account tags params
func (o *CreateAccountTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the create account tags params
func (o *CreateAccountTagsParams) WithAccountID(accountID strfmt.UUID) *CreateAccountTagsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the create account tags params
func (o *CreateAccountTagsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the create account tags params
func (o *CreateAccountTagsParams) WithBody(body []strfmt.UUID) *CreateAccountTagsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create account tags params
func (o *CreateAccountTagsParams) SetBody(body []strfmt.UUID) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccountTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}

	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}

	}

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
