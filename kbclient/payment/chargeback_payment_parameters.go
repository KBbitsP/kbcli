// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewChargebackPaymentParams creates a new ChargebackPaymentParams object
// with the default values initialized.
func NewChargebackPaymentParams() *ChargebackPaymentParams {
	var ()
	return &ChargebackPaymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChargebackPaymentParamsWithTimeout creates a new ChargebackPaymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChargebackPaymentParamsWithTimeout(timeout time.Duration) *ChargebackPaymentParams {
	var ()
	return &ChargebackPaymentParams{

		timeout: timeout,
	}
}

// NewChargebackPaymentParamsWithContext creates a new ChargebackPaymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewChargebackPaymentParamsWithContext(ctx context.Context) *ChargebackPaymentParams {
	var ()
	return &ChargebackPaymentParams{

		Context: ctx,
	}
}

// NewChargebackPaymentParamsWithHTTPClient creates a new ChargebackPaymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChargebackPaymentParamsWithHTTPClient(client *http.Client) *ChargebackPaymentParams {
	var ()
	return &ChargebackPaymentParams{
		HTTPClient: client,
	}
}

/*ChargebackPaymentParams contains all the parameters to send to the API endpoint
for the chargeback payment operation typically these are written to a http.Request
*/
type ChargebackPaymentParams struct {

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
	/*Body*/
	Body *kbmodel.PaymentTransaction
	/*ControlPluginName*/
	ControlPluginName []string
	/*PaymentID*/
	PaymentID strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the chargeback payment params
func (o *ChargebackPaymentParams) WithTimeout(timeout time.Duration) *ChargebackPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chargeback payment params
func (o *ChargebackPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chargeback payment params
func (o *ChargebackPaymentParams) WithContext(ctx context.Context) *ChargebackPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chargeback payment params
func (o *ChargebackPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chargeback payment params
func (o *ChargebackPaymentParams) WithHTTPClient(client *http.Client) *ChargebackPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chargeback payment params
func (o *ChargebackPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the chargeback payment params
func (o *ChargebackPaymentParams) WithXKillbillAPIKey(xKillbillAPIKey string) *ChargebackPaymentParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the chargeback payment params
func (o *ChargebackPaymentParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the chargeback payment params
func (o *ChargebackPaymentParams) WithXKillbillAPISecret(xKillbillAPISecret string) *ChargebackPaymentParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the chargeback payment params
func (o *ChargebackPaymentParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the chargeback payment params
func (o *ChargebackPaymentParams) WithXKillbillComment(xKillbillComment *string) *ChargebackPaymentParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the chargeback payment params
func (o *ChargebackPaymentParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the chargeback payment params
func (o *ChargebackPaymentParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ChargebackPaymentParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the chargeback payment params
func (o *ChargebackPaymentParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the chargeback payment params
func (o *ChargebackPaymentParams) WithXKillbillReason(xKillbillReason *string) *ChargebackPaymentParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the chargeback payment params
func (o *ChargebackPaymentParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the chargeback payment params
func (o *ChargebackPaymentParams) WithBody(body *kbmodel.PaymentTransaction) *ChargebackPaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the chargeback payment params
func (o *ChargebackPaymentParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the chargeback payment params
func (o *ChargebackPaymentParams) WithControlPluginName(controlPluginName []string) *ChargebackPaymentParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the chargeback payment params
func (o *ChargebackPaymentParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPaymentID adds the paymentID to the chargeback payment params
func (o *ChargebackPaymentParams) WithPaymentID(paymentID strfmt.UUID) *ChargebackPaymentParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the chargeback payment params
func (o *ChargebackPaymentParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the chargeback payment params
func (o *ChargebackPaymentParams) WithPluginProperty(pluginProperty []string) *ChargebackPaymentParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the chargeback payment params
func (o *ChargebackPaymentParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *ChargebackPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	valuesControlPluginName := o.ControlPluginName

	joinedControlPluginName := swag.JoinByFormat(valuesControlPluginName, "multi")
	// query array param controlPluginName
	if err := r.SetQueryParam("controlPluginName", joinedControlPluginName...); err != nil {
		return err
	}

	// path param paymentId
	if err := r.SetPathParam("paymentId", o.PaymentID.String()); err != nil {
		return err
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
