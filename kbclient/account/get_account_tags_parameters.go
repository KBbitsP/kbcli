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

// NewGetAccountTagsParams creates a new GetAccountTagsParams object
// with the default values initialized.
func NewGetAccountTagsParams() *GetAccountTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetAccountTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAccountTagsParamsWithTimeout creates a new GetAccountTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAccountTagsParamsWithTimeout(timeout time.Duration) *GetAccountTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetAccountTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		timeout: timeout,
	}
}

// NewGetAccountTagsParamsWithContext creates a new GetAccountTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAccountTagsParamsWithContext(ctx context.Context) *GetAccountTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetAccountTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,

		Context: ctx,
	}
}

// NewGetAccountTagsParamsWithHTTPClient creates a new GetAccountTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAccountTagsParamsWithHTTPClient(client *http.Client) *GetAccountTagsParams {
	var (
		auditDefault           = string("NONE")
		includedDeletedDefault = bool(false)
	)
	return &GetAccountTagsParams{
		Audit:           &auditDefault,
		IncludedDeleted: &includedDeletedDefault,
		HTTPClient:      client,
	}
}

/*GetAccountTagsParams contains all the parameters to send to the API endpoint
for the get account tags operation typically these are written to a http.Request
*/
type GetAccountTagsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*AccountID*/
	AccountID strfmt.UUID
	/*Audit*/
	Audit *string
	/*IncludedDeleted*/
	IncludedDeleted *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get account tags params
func (o *GetAccountTagsParams) WithTimeout(timeout time.Duration) *GetAccountTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get account tags params
func (o *GetAccountTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get account tags params
func (o *GetAccountTagsParams) WithContext(ctx context.Context) *GetAccountTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get account tags params
func (o *GetAccountTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get account tags params
func (o *GetAccountTagsParams) WithHTTPClient(client *http.Client) *GetAccountTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get account tags params
func (o *GetAccountTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get account tags params
func (o *GetAccountTagsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetAccountTagsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get account tags params
func (o *GetAccountTagsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get account tags params
func (o *GetAccountTagsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetAccountTagsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get account tags params
func (o *GetAccountTagsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAccountID adds the accountID to the get account tags params
func (o *GetAccountTagsParams) WithAccountID(accountID strfmt.UUID) *GetAccountTagsParams {
	o.SetAccountID(accountID)
	return o
}

// SetAccountID adds the accountId to the get account tags params
func (o *GetAccountTagsParams) SetAccountID(accountID strfmt.UUID) {
	o.AccountID = accountID
}

// WithAudit adds the audit to the get account tags params
func (o *GetAccountTagsParams) WithAudit(audit *string) *GetAccountTagsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get account tags params
func (o *GetAccountTagsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithIncludedDeleted adds the includedDeleted to the get account tags params
func (o *GetAccountTagsParams) WithIncludedDeleted(includedDeleted *bool) *GetAccountTagsParams {
	o.SetIncludedDeleted(includedDeleted)
	return o
}

// SetIncludedDeleted adds the includedDeleted to the get account tags params
func (o *GetAccountTagsParams) SetIncludedDeleted(includedDeleted *bool) {
	o.IncludedDeleted = includedDeleted
}

// WriteToRequest writes these params to a swagger request
func (o *GetAccountTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param accountId
	if err := r.SetPathParam("accountId", o.AccountID.String()); err != nil {
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

	if o.IncludedDeleted != nil {

		// query param includedDeleted
		var qrIncludedDeleted bool
		if o.IncludedDeleted != nil {
			qrIncludedDeleted = *o.IncludedDeleted
		}
		qIncludedDeleted := swag.FormatBool(qrIncludedDeleted)
		if qIncludedDeleted != "" {
			if err := r.SetQueryParam("includedDeleted", qIncludedDeleted); err != nil {
				return err
			}
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
