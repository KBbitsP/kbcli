// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTransactionCustomFieldsParams creates a new GetTransactionCustomFieldsParams object
// with the default values initialized.
func NewGetTransactionCustomFieldsParams() *GetTransactionCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetTransactionCustomFieldsParams{
		Audit: &auditDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransactionCustomFieldsParamsWithTimeout creates a new GetTransactionCustomFieldsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTransactionCustomFieldsParamsWithTimeout(timeout time.Duration) *GetTransactionCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetTransactionCustomFieldsParams{
		Audit: &auditDefault,

		timeout: timeout,
	}
}

// NewGetTransactionCustomFieldsParamsWithContext creates a new GetTransactionCustomFieldsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTransactionCustomFieldsParamsWithContext(ctx context.Context) *GetTransactionCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetTransactionCustomFieldsParams{
		Audit: &auditDefault,

		Context: ctx,
	}
}

// NewGetTransactionCustomFieldsParamsWithHTTPClient creates a new GetTransactionCustomFieldsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTransactionCustomFieldsParamsWithHTTPClient(client *http.Client) *GetTransactionCustomFieldsParams {
	var (
		auditDefault = string("NONE")
	)
	return &GetTransactionCustomFieldsParams{
		Audit:      &auditDefault,
		HTTPClient: client,
	}
}

/*GetTransactionCustomFieldsParams contains all the parameters to send to the API endpoint
for the get transaction custom fields operation typically these are written to a http.Request
*/
type GetTransactionCustomFieldsParams struct {

	/*Audit*/
	Audit *string
	/*TransactionID*/
	TransactionID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get transaction custom fields params
func (o *GetTransactionCustomFieldsParams) WithTimeout(timeout time.Duration) *GetTransactionCustomFieldsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transaction custom fields params
func (o *GetTransactionCustomFieldsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transaction custom fields params
func (o *GetTransactionCustomFieldsParams) WithContext(ctx context.Context) *GetTransactionCustomFieldsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transaction custom fields params
func (o *GetTransactionCustomFieldsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transaction custom fields params
func (o *GetTransactionCustomFieldsParams) WithHTTPClient(client *http.Client) *GetTransactionCustomFieldsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transaction custom fields params
func (o *GetTransactionCustomFieldsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the get transaction custom fields params
func (o *GetTransactionCustomFieldsParams) WithAudit(audit *string) *GetTransactionCustomFieldsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the get transaction custom fields params
func (o *GetTransactionCustomFieldsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithTransactionID adds the transactionID to the get transaction custom fields params
func (o *GetTransactionCustomFieldsParams) WithTransactionID(transactionID strfmt.UUID) *GetTransactionCustomFieldsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the get transaction custom fields params
func (o *GetTransactionCustomFieldsParams) SetTransactionID(transactionID strfmt.UUID) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransactionCustomFieldsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param transactionId
	if err := r.SetPathParam("transactionId", o.TransactionID.String()); err != nil {
		return err
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
