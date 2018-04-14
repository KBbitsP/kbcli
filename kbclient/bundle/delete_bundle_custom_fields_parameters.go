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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteBundleCustomFieldsParams creates a new DeleteBundleCustomFieldsParams object
// with the default values initialized.
func NewDeleteBundleCustomFieldsParams() *DeleteBundleCustomFieldsParams {
	var ()
	return &DeleteBundleCustomFieldsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteBundleCustomFieldsParamsWithTimeout creates a new DeleteBundleCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteBundleCustomFieldsParamsWithTimeout(timeout time.Duration) *DeleteBundleCustomFieldsParams {
	var ()
	return &DeleteBundleCustomFieldsParams{

		timeout: timeout,
	}
}

// NewDeleteBundleCustomFieldsParamsWithContext creates a new DeleteBundleCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteBundleCustomFieldsParamsWithContext(ctx context.Context) *DeleteBundleCustomFieldsParams {
	var ()
	return &DeleteBundleCustomFieldsParams{

		Context: ctx,
	}
}

// NewDeleteBundleCustomFieldsParamsWithHTTPClient creates a new DeleteBundleCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteBundleCustomFieldsParamsWithHTTPClient(client *http.Client) *DeleteBundleCustomFieldsParams {
	var ()
	return &DeleteBundleCustomFieldsParams{
		HTTPClient: client,
	}
}

/*DeleteBundleCustomFieldsParams contains all the parameters to send to the API endpoint
for the delete bundle custom fields operation typically these are written to a http.Request
*/
type DeleteBundleCustomFieldsParams struct {

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
	/*BundleID*/
	BundleID strfmt.UUID
	/*CustomField*/
	CustomField []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithTimeout(timeout time.Duration) *DeleteBundleCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithContext(ctx context.Context) *DeleteBundleCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithHTTPClient(client *http.Client) *DeleteBundleCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *DeleteBundleCustomFieldsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *DeleteBundleCustomFieldsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithXKillbillComment(xKillbillComment *string) *DeleteBundleCustomFieldsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteBundleCustomFieldsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithXKillbillReason(xKillbillReason *string) *DeleteBundleCustomFieldsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBundleID adds the bundleID to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithBundleID(bundleID strfmt.UUID) *DeleteBundleCustomFieldsParams {
	o.SetBundleID(bundleID)
	return o
}

// SetBundleID adds the bundleId to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetBundleID(bundleID strfmt.UUID) {
	o.BundleID = bundleID
}

// WithCustomField adds the customField to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) WithCustomField(customField []strfmt.UUID) *DeleteBundleCustomFieldsParams {
	o.SetCustomField(customField)
	return o
}

// SetCustomField adds the customField to the delete bundle custom fields params
func (o *DeleteBundleCustomFieldsParams) SetCustomField(customField []strfmt.UUID) {
	o.CustomField = customField
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteBundleCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param bundleId
	if err := r.SetPathParam("bundleId", o.BundleID.String()); err != nil {
		return err
	}

	var valuesCustomField []string
	for _, v := range o.CustomField {
		valuesCustomField = append(valuesCustomField, v.String())
	}

	joinedCustomField := swag.JoinByFormat(valuesCustomField, "multi")
	// query array param customField
	if err := r.SetQueryParam("customField", joinedCustomField...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
