// Code generated by go-swagger; DO NOT EDIT.

package tag

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

// NewGetTagsParams creates a new GetTagsParams object
// with the default values initialized.
func NewGetTagsParams() *GetTagsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetTagsParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTagsParamsWithTimeout creates a new GetTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTagsParamsWithTimeout(timeout time.Duration) *GetTagsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetTagsParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		timeout: timeout,
	}
}

// NewGetTagsParamsWithContext creates a new GetTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTagsParamsWithContext(ctx context.Context) *GetTagsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetTagsParams{
		Audit:  &auditDefault,
		Limit:  &limitDefault,
		Offset: &offsetDefault,

		Context: ctx,
	}
}

// NewGetTagsParamsWithHTTPClient creates a new GetTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTagsParamsWithHTTPClient(client *http.Client) *GetTagsParams {
	var (
		auditDefault  = string("NONE")
		limitDefault  = int64(100)
		offsetDefault = int64(0)
	)
	return &GetTagsParams{
		Audit:      &auditDefault,
		Limit:      &limitDefault,
		Offset:     &offsetDefault,
		HTTPClient: client,
	}
}

/*GetTagsParams contains all the parameters to send to the API endpoint
for the get tags operation typically these are written to a http.Request
*/
type GetTagsParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Audit*/
	Audit *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get tags params
func (o *GetTagsParams) WithTimeout(timeout time.Duration) *GetTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tags params
func (o *GetTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tags params
func (o *GetTagsParams) WithContext(ctx context.Context) *GetTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tags params
func (o *GetTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tags params
func (o *GetTagsParams) WithHTTPClient(client *http.Client) *GetTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tags params
func (o *GetTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get tags params
func (o *GetTagsParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetTagsParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get tags params
func (o *GetTagsParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get tags params
func (o *GetTagsParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetTagsParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get tags params
func (o *GetTagsParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithAudit adds the audit to the get tags params
func (o *GetTagsParams) WithAudit(audit *string) *GetTagsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get tags params
func (o *GetTagsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the get tags params
func (o *GetTagsParams) WithLimit(limit *int64) *GetTagsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get tags params
func (o *GetTagsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get tags params
func (o *GetTagsParams) WithOffset(offset *int64) *GetTagsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get tags params
func (o *GetTagsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WriteToRequest writes these params to a swagger request
func (o *GetTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
