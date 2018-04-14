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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeletePaymentMethodParams creates a new DeletePaymentMethodParams object
// with the default values initialized.
func NewDeletePaymentMethodParams() *DeletePaymentMethodParams {
	var (
		deleteDefaultPmWithAutoPayOffDefault = bool(false)
		forceDefaultPmDeletionDefault        = bool(false)
	)
	return &DeletePaymentMethodParams{
		DeleteDefaultPmWithAutoPayOff: &deleteDefaultPmWithAutoPayOffDefault,
		ForceDefaultPmDeletion:        &forceDefaultPmDeletionDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePaymentMethodParamsWithTimeout creates a new DeletePaymentMethodParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePaymentMethodParamsWithTimeout(timeout time.Duration) *DeletePaymentMethodParams {
	var (
		deleteDefaultPmWithAutoPayOffDefault = bool(false)
		forceDefaultPmDeletionDefault        = bool(false)
	)
	return &DeletePaymentMethodParams{
		DeleteDefaultPmWithAutoPayOff: &deleteDefaultPmWithAutoPayOffDefault,
		ForceDefaultPmDeletion:        &forceDefaultPmDeletionDefault,

		timeout: timeout,
	}
}

// NewDeletePaymentMethodParamsWithContext creates a new DeletePaymentMethodParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePaymentMethodParamsWithContext(ctx context.Context) *DeletePaymentMethodParams {
	var (
		deleteDefaultPmWithAutoPayOffDefault = bool(false)
		forceDefaultPmDeletionDefault        = bool(false)
	)
	return &DeletePaymentMethodParams{
		DeleteDefaultPmWithAutoPayOff: &deleteDefaultPmWithAutoPayOffDefault,
		ForceDefaultPmDeletion:        &forceDefaultPmDeletionDefault,

		Context: ctx,
	}
}

// NewDeletePaymentMethodParamsWithHTTPClient creates a new DeletePaymentMethodParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePaymentMethodParamsWithHTTPClient(client *http.Client) *DeletePaymentMethodParams {
	var (
		deleteDefaultPmWithAutoPayOffDefault = bool(false)
		forceDefaultPmDeletionDefault        = bool(false)
	)
	return &DeletePaymentMethodParams{
		DeleteDefaultPmWithAutoPayOff: &deleteDefaultPmWithAutoPayOffDefault,
		ForceDefaultPmDeletion:        &forceDefaultPmDeletionDefault,
		HTTPClient:                    client,
	}
}

/*DeletePaymentMethodParams contains all the parameters to send to the API endpoint
for the delete payment method operation typically these are written to a http.Request
*/
type DeletePaymentMethodParams struct {

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
	/*DeleteDefaultPmWithAutoPayOff*/
	DeleteDefaultPmWithAutoPayOff *bool
	/*ForceDefaultPmDeletion*/
	ForceDefaultPmDeletion *bool
	/*PaymentMethodID*/
	PaymentMethodID strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete payment method params
func (o *DeletePaymentMethodParams) WithTimeout(timeout time.Duration) *DeletePaymentMethodParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete payment method params
func (o *DeletePaymentMethodParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete payment method params
func (o *DeletePaymentMethodParams) WithContext(ctx context.Context) *DeletePaymentMethodParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete payment method params
func (o *DeletePaymentMethodParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete payment method params
func (o *DeletePaymentMethodParams) WithHTTPClient(client *http.Client) *DeletePaymentMethodParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete payment method params
func (o *DeletePaymentMethodParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the delete payment method params
func (o *DeletePaymentMethodParams) WithXKillbillAPIKey(xKillbillAPIKey string) *DeletePaymentMethodParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the delete payment method params
func (o *DeletePaymentMethodParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the delete payment method params
func (o *DeletePaymentMethodParams) WithXKillbillAPISecret(xKillbillAPISecret string) *DeletePaymentMethodParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the delete payment method params
func (o *DeletePaymentMethodParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the delete payment method params
func (o *DeletePaymentMethodParams) WithXKillbillComment(xKillbillComment *string) *DeletePaymentMethodParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete payment method params
func (o *DeletePaymentMethodParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete payment method params
func (o *DeletePaymentMethodParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeletePaymentMethodParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete payment method params
func (o *DeletePaymentMethodParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete payment method params
func (o *DeletePaymentMethodParams) WithXKillbillReason(xKillbillReason *string) *DeletePaymentMethodParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete payment method params
func (o *DeletePaymentMethodParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithDeleteDefaultPmWithAutoPayOff adds the deleteDefaultPmWithAutoPayOff to the delete payment method params
func (o *DeletePaymentMethodParams) WithDeleteDefaultPmWithAutoPayOff(deleteDefaultPmWithAutoPayOff *bool) *DeletePaymentMethodParams {
	o.SetDeleteDefaultPmWithAutoPayOff(deleteDefaultPmWithAutoPayOff)
	return o
}

// SetDeleteDefaultPmWithAutoPayOff adds the deleteDefaultPmWithAutoPayOff to the delete payment method params
func (o *DeletePaymentMethodParams) SetDeleteDefaultPmWithAutoPayOff(deleteDefaultPmWithAutoPayOff *bool) {
	o.DeleteDefaultPmWithAutoPayOff = deleteDefaultPmWithAutoPayOff
}

// WithForceDefaultPmDeletion adds the forceDefaultPmDeletion to the delete payment method params
func (o *DeletePaymentMethodParams) WithForceDefaultPmDeletion(forceDefaultPmDeletion *bool) *DeletePaymentMethodParams {
	o.SetForceDefaultPmDeletion(forceDefaultPmDeletion)
	return o
}

// SetForceDefaultPmDeletion adds the forceDefaultPmDeletion to the delete payment method params
func (o *DeletePaymentMethodParams) SetForceDefaultPmDeletion(forceDefaultPmDeletion *bool) {
	o.ForceDefaultPmDeletion = forceDefaultPmDeletion
}

// WithPaymentMethodID adds the paymentMethodID to the delete payment method params
func (o *DeletePaymentMethodParams) WithPaymentMethodID(paymentMethodID strfmt.UUID) *DeletePaymentMethodParams {
	o.SetPaymentMethodID(paymentMethodID)
	return o
}

// SetPaymentMethodID adds the paymentMethodId to the delete payment method params
func (o *DeletePaymentMethodParams) SetPaymentMethodID(paymentMethodID strfmt.UUID) {
	o.PaymentMethodID = paymentMethodID
}

// WithPluginProperty adds the pluginProperty to the delete payment method params
func (o *DeletePaymentMethodParams) WithPluginProperty(pluginProperty []string) *DeletePaymentMethodParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the delete payment method params
func (o *DeletePaymentMethodParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePaymentMethodParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.DeleteDefaultPmWithAutoPayOff != nil {

		// query param deleteDefaultPmWithAutoPayOff
		var qrDeleteDefaultPmWithAutoPayOff bool
		if o.DeleteDefaultPmWithAutoPayOff != nil {
			qrDeleteDefaultPmWithAutoPayOff = *o.DeleteDefaultPmWithAutoPayOff
		}
		qDeleteDefaultPmWithAutoPayOff := swag.FormatBool(qrDeleteDefaultPmWithAutoPayOff)
		if qDeleteDefaultPmWithAutoPayOff != "" {
			if err := r.SetQueryParam("deleteDefaultPmWithAutoPayOff", qDeleteDefaultPmWithAutoPayOff); err != nil {
				return err
			}
		}

	}

	if o.ForceDefaultPmDeletion != nil {

		// query param forceDefaultPmDeletion
		var qrForceDefaultPmDeletion bool
		if o.ForceDefaultPmDeletion != nil {
			qrForceDefaultPmDeletion = *o.ForceDefaultPmDeletion
		}
		qForceDefaultPmDeletion := swag.FormatBool(qrForceDefaultPmDeletion)
		if qForceDefaultPmDeletion != "" {
			if err := r.SetQueryParam("forceDefaultPmDeletion", qForceDefaultPmDeletion); err != nil {
				return err
			}
		}

	}

	// path param paymentMethodId
	if err := r.SetPathParam("paymentMethodId", o.PaymentMethodID.String()); err != nil {
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
