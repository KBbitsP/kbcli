// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v2/kbmodel"
)

// NewTransferBundleParams creates a new TransferBundleParams object
// with the default values initialized.
func NewTransferBundleParams() *TransferBundleParams {
	var (
		bcdTransferDefault   = string("USE_EXISTING")
		billingPolicyDefault = string("END_OF_TERM")
	)
	return &TransferBundleParams{
		BcdTransfer:   &bcdTransferDefault,
		BillingPolicy: &billingPolicyDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewTransferBundleParamsWithTimeout creates a new TransferBundleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTransferBundleParamsWithTimeout(timeout time.Duration) *TransferBundleParams {
	var (
		bcdTransferDefault   = string("USE_EXISTING")
		billingPolicyDefault = string("END_OF_TERM")
	)
	return &TransferBundleParams{
		BcdTransfer:   &bcdTransferDefault,
		BillingPolicy: &billingPolicyDefault,

		timeout: timeout,
	}
}

// NewTransferBundleParamsWithContext creates a new TransferBundleParams object
// with the default values initialized, and the ability to set a context for a request
func NewTransferBundleParamsWithContext(ctx context.Context) *TransferBundleParams {
	var (
		bcdTransferDefault   = string("USE_EXISTING")
		billingPolicyDefault = string("END_OF_TERM")
	)
	return &TransferBundleParams{
		BcdTransfer:   &bcdTransferDefault,
		BillingPolicy: &billingPolicyDefault,

		Context: ctx,
	}
}

// NewTransferBundleParamsWithHTTPClient creates a new TransferBundleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTransferBundleParamsWithHTTPClient(client *http.Client) *TransferBundleParams {
	var (
		bcdTransferDefault   = string("USE_EXISTING")
		billingPolicyDefault = string("END_OF_TERM")
	)
	return &TransferBundleParams{
		BcdTransfer:   &bcdTransferDefault,
		BillingPolicy: &billingPolicyDefault,
		HTTPClient:    client,
	}
}

/*TransferBundleParams contains all the parameters to send to the API endpoint
for the transfer bundle operation typically these are written to a http.Request
*/
type TransferBundleParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*BcdTransfer*/
	BcdTransfer *string
	/*BillingPolicy*/
	BillingPolicy *string
	/*Body*/
	Body *kbmodel.Bundle
	/*BundleID*/
	BundleID strfmt.UUID
	/*PluginProperty*/
	PluginProperty []string
	/*RequestedDate*/
	RequestedDate *strfmt.Date

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the transfer bundle params
func (o *TransferBundleParams) WithTimeout(timeout time.Duration) *TransferBundleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transfer bundle params
func (o *TransferBundleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transfer bundle params
func (o *TransferBundleParams) WithContext(ctx context.Context) *TransferBundleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transfer bundle params
func (o *TransferBundleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transfer bundle params
func (o *TransferBundleParams) WithHTTPClient(client *http.Client) *TransferBundleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transfer bundle params
func (o *TransferBundleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the transfer bundle params
func (o *TransferBundleParams) WithXKillbillComment(xKillbillComment *string) *TransferBundleParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the transfer bundle params
func (o *TransferBundleParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the transfer bundle params
func (o *TransferBundleParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *TransferBundleParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the transfer bundle params
func (o *TransferBundleParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the transfer bundle params
func (o *TransferBundleParams) WithXKillbillReason(xKillbillReason *string) *TransferBundleParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the transfer bundle params
func (o *TransferBundleParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBcdTransfer adds the bcdTransfer to the transfer bundle params
func (o *TransferBundleParams) WithBcdTransfer(bcdTransfer *string) *TransferBundleParams {
	o.SetBcdTransfer(bcdTransfer)
	return o
}

// SetBcdTransfer adds the bcdTransfer to the transfer bundle params
func (o *TransferBundleParams) SetBcdTransfer(bcdTransfer *string) {
	o.BcdTransfer = bcdTransfer
}

// WithBillingPolicy adds the billingPolicy to the transfer bundle params
func (o *TransferBundleParams) WithBillingPolicy(billingPolicy *string) *TransferBundleParams {
	o.SetBillingPolicy(billingPolicy)
	return o
}

// SetBillingPolicy adds the billingPolicy to the transfer bundle params
func (o *TransferBundleParams) SetBillingPolicy(billingPolicy *string) {
	o.BillingPolicy = billingPolicy
}

// WithBody adds the body to the transfer bundle params
func (o *TransferBundleParams) WithBody(body *kbmodel.Bundle) *TransferBundleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the transfer bundle params
func (o *TransferBundleParams) SetBody(body *kbmodel.Bundle) {
	o.Body = body
}

// WithBundleID adds the bundleID to the transfer bundle params
func (o *TransferBundleParams) WithBundleID(bundleID strfmt.UUID) *TransferBundleParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the transfer bundle params
func (o *TransferBundleParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WithPluginProperty adds the pluginProperty to the transfer bundle params
func (o *TransferBundleParams) WithPluginProperty(pluginProperty []string) *TransferBundleParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the transfer bundle params
func (o *TransferBundleParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRequestedDate adds the requestedDate to the transfer bundle params
func (o *TransferBundleParams) WithRequestedDate(requestedDate *strfmt.Date) *TransferBundleParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the transfer bundle params
func (o *TransferBundleParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WriteToRequest writes these params to a swagger request
func (o *TransferBundleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.BcdTransfer != nil {

		// query param bcdTransfer
		var qrBcdTransfer string
		if o.BcdTransfer != nil {
			qrBcdTransfer = *o.BcdTransfer
		}
		qBcdTransfer := qrBcdTransfer
		if qBcdTransfer != "" {
			if err := r.SetQueryParam("bcdTransfer", qBcdTransfer); err != nil {
				return err
			}
		}

	}

	if o.BillingPolicy != nil {

		// query param billingPolicy
		var qrBillingPolicy string
		if o.BillingPolicy != nil {
			qrBillingPolicy = *o.BillingPolicy
		}
		qBillingPolicy := qrBillingPolicy
		if qBillingPolicy != "" {
			if err := r.SetQueryParam("billingPolicy", qBillingPolicy); err != nil {
				return err
			}
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
		return err
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.Date
		if o.RequestedDate != nil {
			qrRequestedDate = *o.RequestedDate
		}
		qRequestedDate := qrRequestedDate.String()
		if qRequestedDate != "" {
			if err := r.SetQueryParam("requestedDate", qRequestedDate); err != nil {
				return err
			}
		}

	}

	// header param WithProfilingInfo
	if o.WithProfilingInfo != nil && len(*o.WithProfilingInfo) > 0 {
		if err := r.SetHeaderParam("X-Killbill-Profiling-Req", *o.WithProfilingInfo); err != nil {
			return err
		}
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
