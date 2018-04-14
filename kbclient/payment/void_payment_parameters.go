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

// NewVoidPaymentParams creates a new VoidPaymentParams object
// with the default values initialized.
func NewVoidPaymentParams() *VoidPaymentParams {
	var ()
	return &VoidPaymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVoidPaymentParamsWithTimeout creates a new VoidPaymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVoidPaymentParamsWithTimeout(timeout time.Duration) *VoidPaymentParams {
	var ()
	return &VoidPaymentParams{

		timeout: timeout,
	}
}

// NewVoidPaymentParamsWithContext creates a new VoidPaymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewVoidPaymentParamsWithContext(ctx context.Context) *VoidPaymentParams {
	var ()
	return &VoidPaymentParams{

		Context: ctx,
	}
}

// NewVoidPaymentParamsWithHTTPClient creates a new VoidPaymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVoidPaymentParamsWithHTTPClient(client *http.Client) *VoidPaymentParams {
	var ()
	return &VoidPaymentParams{
		HTTPClient: client,
	}
}

/*VoidPaymentParams contains all the parameters to send to the API endpoint
for the void payment operation typically these are written to a http.Request
*/
type VoidPaymentParams struct {

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

// WithTimeout adds the timeout to the void payment params
func (o *VoidPaymentParams) WithTimeout(timeout time.Duration) *VoidPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the void payment params
func (o *VoidPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the void payment params
func (o *VoidPaymentParams) WithContext(ctx context.Context) *VoidPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the void payment params
func (o *VoidPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the void payment params
func (o *VoidPaymentParams) WithHTTPClient(client *http.Client) *VoidPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the void payment params
func (o *VoidPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the void payment params
func (o *VoidPaymentParams) WithXKillbillAPIKey(xKillbillAPIKey string) *VoidPaymentParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the void payment params
func (o *VoidPaymentParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the void payment params
func (o *VoidPaymentParams) WithXKillbillAPISecret(xKillbillAPISecret string) *VoidPaymentParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the void payment params
func (o *VoidPaymentParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the void payment params
func (o *VoidPaymentParams) WithXKillbillComment(xKillbillComment *string) *VoidPaymentParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the void payment params
func (o *VoidPaymentParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the void payment params
func (o *VoidPaymentParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *VoidPaymentParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the void payment params
func (o *VoidPaymentParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the void payment params
func (o *VoidPaymentParams) WithXKillbillReason(xKillbillReason *string) *VoidPaymentParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the void payment params
func (o *VoidPaymentParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the void payment params
func (o *VoidPaymentParams) WithBody(body *kbmodel.PaymentTransaction) *VoidPaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the void payment params
func (o *VoidPaymentParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the void payment params
func (o *VoidPaymentParams) WithControlPluginName(controlPluginName []string) *VoidPaymentParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the void payment params
func (o *VoidPaymentParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPaymentID adds the paymentID to the void payment params
func (o *VoidPaymentParams) WithPaymentID(paymentID strfmt.UUID) *VoidPaymentParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the void payment params
func (o *VoidPaymentParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the void payment params
func (o *VoidPaymentParams) WithPluginProperty(pluginProperty []string) *VoidPaymentParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the void payment params
func (o *VoidPaymentParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *VoidPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
