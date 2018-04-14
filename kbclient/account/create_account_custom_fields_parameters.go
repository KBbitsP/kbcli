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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewCreateAccountCustomFieldsParams creates a new CreateAccountCustomFieldsParams object
// with the default values initialized.
func NewCreateAccountCustomFieldsParams() *CreateAccountCustomFieldsParams {
	var ()
	return &CreateAccountCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAccountCustomFieldsParamsWithTimeout creates a new CreateAccountCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAccountCustomFieldsParamsWithTimeout(timeout time.Duration) *CreateAccountCustomFieldsParams {
	var ()
	return &CreateAccountCustomFieldsParams{

		timeout: timeout,
	}
}

// NewCreateAccountCustomFieldsParamsWithContext creates a new CreateAccountCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAccountCustomFieldsParamsWithContext(ctx context.Context) *CreateAccountCustomFieldsParams {
	var ()
	return &CreateAccountCustomFieldsParams{

		Context: ctx,
	}
}

// NewCreateAccountCustomFieldsParamsWithHTTPClient creates a new CreateAccountCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAccountCustomFieldsParamsWithHTTPClient(client *http.Client) *CreateAccountCustomFieldsParams {
	var ()
	return &CreateAccountCustomFieldsParams{
		HTTPClient: client,
	}
}

/*CreateAccountCustomFieldsParams contains all the parameters to send to the API endpoint
for the create account custom fields operation typically these are written to a http.Request
*/
type CreateAccountCustomFieldsParams struct {

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
	Body []*kbmodel.CustomField

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) WithTimeout(timeout time.Duration) *CreateAccountCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) WithContext(ctx context.Context) *CreateAccountCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) WithHTTPClient(client *http.Client) *CreateAccountCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *CreateAccountCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *CreateAccountCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *CreateAccountCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateAccountCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *CreateAccountCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) WithAccountID(accountID strfmt.UUID) *CreateAccountCustomFieldsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *CreateAccountCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create account custom fields params
func (o *CreateAccountCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAccountCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
