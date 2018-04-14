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

// NewRefundPaymentParams creates a new RefundPaymentParams object
// with the default values initialized.
func NewRefundPaymentParams() *RefundPaymentParams {
	var ()
	return &RefundPaymentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRefundPaymentParamsWithTimeout creates a new RefundPaymentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRefundPaymentParamsWithTimeout(timeout time.Duration) *RefundPaymentParams {
	var ()
	return &RefundPaymentParams{

		timeout: timeout,
	}
}

// NewRefundPaymentParamsWithContext creates a new RefundPaymentParams object
// with the default values initialized, and the ability to set a context for a request
func NewRefundPaymentParamsWithContext(ctx context.Context) *RefundPaymentParams {
	var ()
	return &RefundPaymentParams{

		Context: ctx,
	}
}

// NewRefundPaymentParamsWithHTTPClient creates a new RefundPaymentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRefundPaymentParamsWithHTTPClient(client *http.Client) *RefundPaymentParams {
	var ()
	return &RefundPaymentParams{
		HTTPClient: client,
	}
}

/*RefundPaymentParams contains all the parameters to send to the API endpoint
for the refund payment operation typically these are written to a http.Request
*/
type RefundPaymentParams struct {

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

// WithTimeout adds the timeout to the refund payment params
func (o *RefundPaymentParams) WithTimeout(timeout time.Duration) *RefundPaymentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refund payment params
func (o *RefundPaymentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refund payment params
func (o *RefundPaymentParams) WithContext(ctx context.Context) *RefundPaymentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refund payment params
func (o *RefundPaymentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refund payment params
func (o *RefundPaymentParams) WithHTTPClient(client *http.Client) *RefundPaymentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refund payment params
func (o *RefundPaymentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the refund payment params
func (o *RefundPaymentParams) WithXKillbillAPIKey(xKillbillAPIKey string) *RefundPaymentParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the refund payment params
func (o *RefundPaymentParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the refund payment params
func (o *RefundPaymentParams) WithXKillbillAPISecret(xKillbillAPISecret string) *RefundPaymentParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the refund payment params
func (o *RefundPaymentParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the refund payment params
func (o *RefundPaymentParams) WithXKillbillComment(xKillbillComment *string) *RefundPaymentParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the refund payment params
func (o *RefundPaymentParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the refund payment params
func (o *RefundPaymentParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *RefundPaymentParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the refund payment params
func (o *RefundPaymentParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the refund payment params
func (o *RefundPaymentParams) WithXKillbillReason(xKillbillReason *string) *RefundPaymentParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the refund payment params
func (o *RefundPaymentParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the refund payment params
func (o *RefundPaymentParams) WithBody(body *kbmodel.PaymentTransaction) *RefundPaymentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the refund payment params
func (o *RefundPaymentParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the refund payment params
func (o *RefundPaymentParams) WithControlPluginName(controlPluginName []string) *RefundPaymentParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the refund payment params
func (o *RefundPaymentParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPaymentID adds the paymentID to the refund payment params
func (o *RefundPaymentParams) WithPaymentID(paymentID strfmt.UUID) *RefundPaymentParams {
	o.SetPaymentID(paymentID)
	return o
}

// SetPaymentID adds the paymentId to the refund payment params
func (o *RefundPaymentParams) SetPaymentID(paymentID strfmt.UUID) {
	o.PaymentID = paymentID
}

// WithPluginProperty adds the pluginProperty to the refund payment params
func (o *RefundPaymentParams) WithPluginProperty(pluginProperty []string) *RefundPaymentParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the refund payment params
func (o *RefundPaymentParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *RefundPaymentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
