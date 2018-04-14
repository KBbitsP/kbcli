// Code generated by go-swagger; DO NOT EDIT.

package tenant

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

// NewDeletePluginPaymentStateMachineConfigParams creates a new DeletePluginPaymentStateMachineConfigParams object
// with the default values initialized.
func NewDeletePluginPaymentStateMachineConfigParams() *DeletePluginPaymentStateMachineConfigParams {
	var ()
	return &DeletePluginPaymentStateMachineConfigParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePluginPaymentStateMachineConfigParamsWithTimeout creates a new DeletePluginPaymentStateMachineConfigParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeletePluginPaymentStateMachineConfigParamsWithTimeout(timeout time.Duration) *DeletePluginPaymentStateMachineConfigParams {
	var ()
	return &DeletePluginPaymentStateMachineConfigParams{

		timeout: timeout,
	}
}

// NewDeletePluginPaymentStateMachineConfigParamsWithContext creates a new DeletePluginPaymentStateMachineConfigParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeletePluginPaymentStateMachineConfigParamsWithContext(ctx context.Context) *DeletePluginPaymentStateMachineConfigParams {
	var ()
	return &DeletePluginPaymentStateMachineConfigParams{

		Context: ctx,
	}
}

// NewDeletePluginPaymentStateMachineConfigParamsWithHTTPClient creates a new DeletePluginPaymentStateMachineConfigParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeletePluginPaymentStateMachineConfigParamsWithHTTPClient(client *http.Client) *DeletePluginPaymentStateMachineConfigParams {
	var ()
	return &DeletePluginPaymentStateMachineConfigParams{
		HTTPClient: client,
	}
}

/*DeletePluginPaymentStateMachineConfigParams contains all the parameters to send to the API endpoint
for the delete plugin payment state machine config operation typically these are written to a http.Request
*/
type DeletePluginPaymentStateMachineConfigParams struct {

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
	/*PluginName*/
	PluginName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithTimeout(timeout time.Duration) *DeletePluginPaymentStateMachineConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithContext(ctx context.Context) *DeletePluginPaymentStateMachineConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithHTTPClient(client *http.Client) *DeletePluginPaymentStateMachineConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithXKillbillAPIKey(xKillbillAPIKey string) *DeletePluginPaymentStateMachineConfigParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithXKillbillAPISecret(xKillbillAPISecret string) *DeletePluginPaymentStateMachineConfigParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithXKillbillComment(xKillbillComment *string) *DeletePluginPaymentStateMachineConfigParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeletePluginPaymentStateMachineConfigParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithXKillbillReason(xKillbillReason *string) *DeletePluginPaymentStateMachineConfigParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithPluginName adds the pluginName to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) WithPluginName(pluginName string) *DeletePluginPaymentStateMachineConfigParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the delete plugin payment state machine config params
func (o *DeletePluginPaymentStateMachineConfigParams) SetPluginName(pluginName string) {
	o.PluginName = pluginName
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePluginPaymentStateMachineConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param pluginName
	if err := r.SetPathParam("pluginName", o.PluginName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
