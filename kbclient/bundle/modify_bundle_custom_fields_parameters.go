// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// NewModifyBundleCustomFieldsParams creates a new ModifyBundleCustomFieldsParams object
// with the default values initialized.
func NewModifyBundleCustomFieldsParams() *ModifyBundleCustomFieldsParams {
	var ()
	return &ModifyBundleCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyBundleCustomFieldsParamsWithTimeout creates a new ModifyBundleCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyBundleCustomFieldsParamsWithTimeout(timeout time.Duration) *ModifyBundleCustomFieldsParams {
	var ()
	return &ModifyBundleCustomFieldsParams{

		timeout: timeout,
	}
}

// NewModifyBundleCustomFieldsParamsWithContext creates a new ModifyBundleCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyBundleCustomFieldsParamsWithContext(ctx context.Context) *ModifyBundleCustomFieldsParams {
	var ()
	return &ModifyBundleCustomFieldsParams{

		Context: ctx,
	}
}

// NewModifyBundleCustomFieldsParamsWithHTTPClient creates a new ModifyBundleCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyBundleCustomFieldsParamsWithHTTPClient(client *http.Client) *ModifyBundleCustomFieldsParams {
	var ()
	return &ModifyBundleCustomFieldsParams{
		HTTPClient: client,
	}
}

/*ModifyBundleCustomFieldsParams contains all the parameters to send to the API endpoint
for the modify bundle custom fields operation typically these are written to a http.Request
*/
type ModifyBundleCustomFieldsParams struct {

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
	Body []*kbmodel.CustomField
	/*BundleID*/
	BundleID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) WithTimeout(timeout time.Duration) *ModifyBundleCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) WithContext(ctx context.Context) *ModifyBundleCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) WithHTTPClient(client *http.Client) *ModifyBundleCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *ModifyBundleCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *ModifyBundleCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *ModifyBundleCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ModifyBundleCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *ModifyBundleCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) WithBody(body []*kbmodel.CustomField) *ModifyBundleCustomFieldsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) SetBody(body []*kbmodel.CustomField) {
	o.Body = body
}

// WithBundleID adds the bundleID to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) WithBundleID(bundleID strfmt.UUID) *ModifyBundleCustomFieldsParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the modify bundle custom fields params
func (o *ModifyBundleCustomFieldsParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyBundleCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
