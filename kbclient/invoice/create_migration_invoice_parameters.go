// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewCreateMigrationInvoiceParams creates a new CreateMigrationInvoiceParams object
// with the default values initialized.
func NewCreateMigrationInvoiceParams() *CreateMigrationInvoiceParams {
	var ()
	return &CreateMigrationInvoiceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateMigrationInvoiceParamsWithTimeout creates a new CreateMigrationInvoiceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateMigrationInvoiceParamsWithTimeout(timeout time.Duration) *CreateMigrationInvoiceParams {
	var ()
	return &CreateMigrationInvoiceParams{

		timeout: timeout,
	}
}

// NewCreateMigrationInvoiceParamsWithContext creates a new CreateMigrationInvoiceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateMigrationInvoiceParamsWithContext(ctx context.Context) *CreateMigrationInvoiceParams {
	var ()
	return &CreateMigrationInvoiceParams{

		Context: ctx,
	}
}

// NewCreateMigrationInvoiceParamsWithHTTPClient creates a new CreateMigrationInvoiceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateMigrationInvoiceParamsWithHTTPClient(client *http.Client) *CreateMigrationInvoiceParams {
	var ()
	return &CreateMigrationInvoiceParams{
		HTTPClient: client,
	}
}

/*CreateMigrationInvoiceParams contains all the parameters to send to the API endpoint
for the create migration invoice operation typically these are written to a http.Request
*/
type CreateMigrationInvoiceParams struct {

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
	Body []*kbmodel.InvoiceItem
	/*TargetDate*/
	TargetDate *strfmt.Date

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithTimeout(timeout time.Duration) *CreateMigrationInvoiceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithContext(ctx context.Context) *CreateMigrationInvoiceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithHTTPClient(client *http.Client) *CreateMigrationInvoiceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithXKillbillAPIKey(xKillbillAPIKey string) *CreateMigrationInvoiceParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithXKillbillAPISecret(xKillbillAPISecret string) *CreateMigrationInvoiceParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithXKillbillComment(xKillbillComment *string) *CreateMigrationInvoiceParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CreateMigrationInvoiceParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithXKillbillReason(xKillbillReason *string) *CreateMigrationInvoiceParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithAccountID(accountID strfmt.UUID) *CreateMigrationInvoiceParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithBody adds the body to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithBody(body []*kbmodel.InvoiceItem) *CreateMigrationInvoiceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetBody(body []*kbmodel.InvoiceItem) {
	o.Body = body
}

// WithTargetDate adds the targetDate to the create migration invoice params
func (o *CreateMigrationInvoiceParams) WithTargetDate(targetDate *strfmt.Date) *CreateMigrationInvoiceParams {
	o.SetTargetDate(targetDate)
	return o
}

// SetTargetDate adds the targetDate to the create migration invoice params
func (o *CreateMigrationInvoiceParams) SetTargetDate(targetDate *strfmt.Date) {
	o.TargetDate = targetDate
}

// WriteToRequest writes these params to a swagger request
func (o *CreateMigrationInvoiceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.TargetDate != nil {

		// query param targetDate
		var qrTargetDate strfmt.Date
		if o.TargetDate != nil {
			qrTargetDate = *o.TargetDate
		}
		qTargetDate := qrTargetDate.String()
		if qTargetDate != "" {
			if err := r.SetQueryParam("targetDate", qTargetDate); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
