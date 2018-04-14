// Code generated by go-swagger; DO NOT EDIT.

package account

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

// NewDeleteAccountTagsParams creates a new DeleteAccountTagsParams object
// with the default values initialized.
func NewDeleteAccountTagsParams() *DeleteAccountTagsParams {
	var ()
	return &DeleteAccountTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAccountTagsParamsWithTimeout creates a new DeleteAccountTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAccountTagsParamsWithTimeout(timeout time.Duration) *DeleteAccountTagsParams {
	var ()
	return &DeleteAccountTagsParams{

		timeout: timeout,
	}
}

// NewDeleteAccountTagsParamsWithContext creates a new DeleteAccountTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAccountTagsParamsWithContext(ctx context.Context) *DeleteAccountTagsParams {
	var ()
	return &DeleteAccountTagsParams{

		Context: ctx,
	}
}

// NewDeleteAccountTagsParamsWithHTTPClient creates a new DeleteAccountTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAccountTagsParamsWithHTTPClient(client *http.Client) *DeleteAccountTagsParams {
	var ()
	return &DeleteAccountTagsParams{
		HTTPClient: client,
	}
}

/*DeleteAccountTagsParams contains all the parameters to send to the API endpoint
for the delete account tags operation typically these are written to a http.Request
*/
type DeleteAccountTagsParams struct {

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
	/*TagDef*/
	TagDef []strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete account tags params
func (o *DeleteAccountTagsParams) WithTimeout(timeout time.Duration) *DeleteAccountTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete account tags params
func (o *DeleteAccountTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete account tags params
func (o *DeleteAccountTagsParams) WithContext(ctx context.Context) *DeleteAccountTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete account tags params
func (o *DeleteAccountTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete account tags params
func (o *DeleteAccountTagsParams) WithHTTPClient(client *http.Client) *DeleteAccountTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete account tags params
func (o *DeleteAccountTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the delete account tags params
func (o *DeleteAccountTagsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *DeleteAccountTagsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the delete account tags params
func (o *DeleteAccountTagsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the delete account tags params
func (o *DeleteAccountTagsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *DeleteAccountTagsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the delete account tags params
func (o *DeleteAccountTagsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithXKillbillComment adds the xKillbillComment to the delete account tags params
func (o *DeleteAccountTagsParams) WithXKillbillComment(xKillbillComment *string) *DeleteAccountTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete account tags params
func (o *DeleteAccountTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete account tags params
func (o *DeleteAccountTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteAccountTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete account tags params
func (o *DeleteAccountTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete account tags params
func (o *DeleteAccountTagsParams) WithXKillbillReason(xKillbillReason *string) *DeleteAccountTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete account tags params
func (o *DeleteAccountTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithAccountID adds the accountID to the delete account tags params
func (o *DeleteAccountTagsParams) WithAccountID(accountID strfmt.UUID) *DeleteAccountTagsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the delete account tags params
func (o *DeleteAccountTagsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithTagDef adds the tagDef to the delete account tags params
func (o *DeleteAccountTagsParams) WithTagDef(tagDef []strfmt.UUID) *DeleteAccountTagsParams {
	o.SetTagDef(tagDef)
	return o
}

// SetTagDef adds the tagDef to the delete account tags params
func (o *DeleteAccountTagsParams) SetTagDef(tagDef []strfmt.UUID) {
	o.TagDef = tagDef
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAccountTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesTagDef []string
	for _, v := range o.TagDef {
		valuesTagDef = append(valuesTagDef, v.String())
	}

	joinedTagDef := swag.JoinByFormat(valuesTagDef, "multi")
	// query array param tagDef
	if err := r.SetQueryParam("tagDef", joinedTagDef...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
