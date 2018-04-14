// Code generated by go-swagger; DO NOT EDIT.

package payment_method

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

// NewGetPaymentMethodCustomFieldsParams creates a new GetPaymentMethodCustomFieldsParams object
// with the default values initialized.
func NewGetPaymentMethodCustomFieldsParams() *GetPaymentMethodCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetPaymentMethodCustomFieldsParams{
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetPaymentMethodCustomFieldsParamsWithTimeout creates a new GetPaymentMethodCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetPaymentMethodCustomFieldsParamsWithTimeout(timeout time.Duration) *GetPaymentMethodCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetPaymentMethodCustomFieldsParams{
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetPaymentMethodCustomFieldsParamsWithContext creates a new GetPaymentMethodCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetPaymentMethodCustomFieldsParamsWithContext(ctx context.Context) *GetPaymentMethodCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetPaymentMethodCustomFieldsParams{
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetPaymentMethodCustomFieldsParamsWithHTTPClient creates a new GetPaymentMethodCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetPaymentMethodCustomFieldsParamsWithHTTPClient(client *http.Client) *GetPaymentMethodCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetPaymentMethodCustomFieldsParams{
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetPaymentMethodCustomFieldsParams contains all the parameters to send to the API endpoint
for the get payment method custom fields operation typically these are written to a http.Request
*/
type GetPaymentMethodCustomFieldsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*PaymentMethodID*/
	PaymentMethodID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithTimeout(timeout time.Duration) *GetPaymentMethodCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithContext(ctx context.Context) *GetPaymentMethodCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithHTTPClient(client *http.Client) *GetPaymentMethodCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetPaymentMethodCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetPaymentMethodCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithAudit(audit *string) *GetPaymentMethodCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithPaymentMethodID adds the paymentMethodID to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) WithPaymentMethodID(paymentMethodID strfmt.UUID) *GetPaymentMethodCustomFieldsParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the get payment method custom fields params
func (o *GetPaymentMethodCustomFieldsParams) SetPaymentMethodID(paymentMethodID strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WriteToRequest writes these params to a swagger request
func (o *GetPaymentMethodCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Audit != nil {

		// query param audit
		var qrAudit string
		if o.Audit != nil {
			qrAudit = *o.Audit
		}
		qAudit := qrAudit
		if qAudit != "" {
			if err := r.SetQueryParam("audit", qAudit); err != nil {
				return err
			}
		}

	}

	// path param paymentMethodId
	if err := r.SetPathParam("paymentMethodId", o.PaymentMethodID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
