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

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTagAuditLogsWithHistoryParams creates a new GetTagAuditLogsWithHistoryParams object
// with the default values initialized.
func NewGetTagAuditLogsWithHistoryParams() *GetTagAuditLogsWithHistoryParams {
	var ()
	return &GetTagAuditLogsWithHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTagAuditLogsWithHistoryParamsWithTimeout creates a new GetTagAuditLogsWithHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTagAuditLogsWithHistoryParamsWithTimeout(timeout time.Duration) *GetTagAuditLogsWithHistoryParams {
	var ()
	return &GetTagAuditLogsWithHistoryParams{

		timeout: timeout,
	}
}

// NewGetTagAuditLogsWithHistoryParamsWithContext creates a new GetTagAuditLogsWithHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTagAuditLogsWithHistoryParamsWithContext(ctx context.Context) *GetTagAuditLogsWithHistoryParams {
	var ()
	return &GetTagAuditLogsWithHistoryParams{

		Context: ctx,
	}
}

// NewGetTagAuditLogsWithHistoryParamsWithHTTPClient creates a new GetTagAuditLogsWithHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTagAuditLogsWithHistoryParamsWithHTTPClient(client *http.Client) *GetTagAuditLogsWithHistoryParams {
	var ()
	return &GetTagAuditLogsWithHistoryParams{
		HTTPClient: client,
	}
}

/*GetTagAuditLogsWithHistoryParams contains all the parameters to send to the API endpoint
for the get tag audit logs with history operation typically these are written to a http.Request
*/
type GetTagAuditLogsWithHistoryParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*TagID*/
	TagID strfmt.UUID

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) WithTimeout(timeout time.Duration) *GetTagAuditLogsWithHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) WithContext(ctx context.Context) *GetTagAuditLogsWithHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) WithHTTPClient(client *http.Client) *GetTagAuditLogsWithHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetTagAuditLogsWithHistoryParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetTagAuditLogsWithHistoryParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithTagID adds the tagID to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) WithTagID(tagID strfmt.UUID) *GetTagAuditLogsWithHistoryParams {
	o.SetTagID(tagID)
	return o
}

// SetTagID adds the tagId to the get tag audit logs with history params
func (o *GetTagAuditLogsWithHistoryParams) SetTagID(tagID strfmt.UUID) {
	o.TagID = tagID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTagAuditLogsWithHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param tagId
	if err := r.SetPathParam("tagId", o.TagID.String()); err != nil {
		return err
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
