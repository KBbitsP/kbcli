// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// NewSearchPaymentsParams creates a new SearchPaymentsParams object
// with the default values initialized.
func NewSearchPaymentsParams() *SearchPaymentsParams {
	var (
		auditDefault          = string("NONE")
		limitDefault          = int64(100)
		offsetDefault         = int64(0)
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &SearchPaymentsParams{
		Audit:          &auditDefault,
		Limit:          &limitDefault,
		Offset:         &offsetDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewSearchPaymentsParamsWithTimeout creates a new SearchPaymentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSearchPaymentsParamsWithTimeout(timeout time.Duration) *SearchPaymentsParams {
	var (
		auditDefault          = string("NONE")
		limitDefault          = int64(100)
		offsetDefault         = int64(0)
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &SearchPaymentsParams{
		Audit:          &auditDefault,
		Limit:          &limitDefault,
		Offset:         &offsetDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		timeout: timeout,
	}
}

// NewSearchPaymentsParamsWithContext creates a new SearchPaymentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewSearchPaymentsParamsWithContext(ctx context.Context) *SearchPaymentsParams {
	var (
		auditDefault          = string("NONE")
		limitDefault          = int64(100)
		offsetDefault         = int64(0)
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &SearchPaymentsParams{
		Audit:          &auditDefault,
		Limit:          &limitDefault,
		Offset:         &offsetDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,

		Context: ctx,
	}
}

// NewSearchPaymentsParamsWithHTTPClient creates a new SearchPaymentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSearchPaymentsParamsWithHTTPClient(client *http.Client) *SearchPaymentsParams {
	var (
		auditDefault          = string("NONE")
		limitDefault          = int64(100)
		offsetDefault         = int64(0)
		withAttemptsDefault   = bool(false)
		withPluginInfoDefault = bool(false)
	)
	return &SearchPaymentsParams{
		Audit:          &auditDefault,
		Limit:          &limitDefault,
		Offset:         &offsetDefault,
		WithAttempts:   &withAttemptsDefault,
		WithPluginInfo: &withPluginInfoDefault,
		HTTPClient:     client,
	}
}

/*SearchPaymentsParams contains all the parameters to send to the API endpoint
for the search payments operation typically these are written to a http.Request
*/
type SearchPaymentsParams struct {

	/*Audit*/
	Audit *string
	/*Limit*/
	Limit *int64
	/*Offset*/
	Offset *int64
	/*PluginName*/
	PluginName *string
	/*PluginProperty*/
	PluginProperty []string
	/*SearchKey*/
	SearchKey string
	/*WithAttempts*/
	WithAttempts *bool
	/*WithPluginInfo*/
	WithPluginInfo *bool

	WithStackTrace        *bool // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the search payments params
func (o *SearchPaymentsParams) WithTimeout(timeout time.Duration) *SearchPaymentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the search payments params
func (o *SearchPaymentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the search payments params
func (o *SearchPaymentsParams) WithContext(ctx context.Context) *SearchPaymentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the search payments params
func (o *SearchPaymentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the search payments params
func (o *SearchPaymentsParams) WithHTTPClient(client *http.Client) *SearchPaymentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the search payments params
func (o *SearchPaymentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudit adds the audit to the search payments params
func (o *SearchPaymentsParams) WithAudit(audit *string) *SearchPaymentsParams {
	o.SetAudit(audit)
	return o
}

// SetAudit adds the audit to the search payments params
func (o *SearchPaymentsParams) SetAudit(audit *string) {
	o.Audit = audit
}

// WithLimit adds the limit to the search payments params
func (o *SearchPaymentsParams) WithLimit(limit *int64) *SearchPaymentsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the search payments params
func (o *SearchPaymentsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the search payments params
func (o *SearchPaymentsParams) WithOffset(offset *int64) *SearchPaymentsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the search payments params
func (o *SearchPaymentsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithPluginName adds the pluginName to the search payments params
func (o *SearchPaymentsParams) WithPluginName(pluginName *string) *SearchPaymentsParams {
	o.SetPluginName(pluginName)
	return o
}

// SetPluginName adds the pluginName to the search payments params
func (o *SearchPaymentsParams) SetPluginName(pluginName *string) {
	o.PluginName = pluginName
}

// WithPluginProperty adds the pluginProperty to the search payments params
func (o *SearchPaymentsParams) WithPluginProperty(pluginProperty []string) *SearchPaymentsParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the search payments params
func (o *SearchPaymentsParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithSearchKey adds the searchKey to the search payments params
func (o *SearchPaymentsParams) WithSearchKey(searchKey string) *SearchPaymentsParams {
	o.SetSearchKey(searchKey)
	return o
}

// SetSearchKey adds the searchKey to the search payments params
func (o *SearchPaymentsParams) SetSearchKey(searchKey string) {
	o.SearchKey = searchKey
}

// WithWithAttempts adds the withAttempts to the search payments params
func (o *SearchPaymentsParams) WithWithAttempts(withAttempts *bool) *SearchPaymentsParams {
	o.SetWithAttempts(withAttempts)
	return o
}

// SetWithAttempts adds the withAttempts to the search payments params
func (o *SearchPaymentsParams) SetWithAttempts(withAttempts *bool) {
	o.WithAttempts = withAttempts
}

// WithWithPluginInfo adds the withPluginInfo to the search payments params
func (o *SearchPaymentsParams) WithWithPluginInfo(withPluginInfo *bool) *SearchPaymentsParams {
	o.SetWithPluginInfo(withPluginInfo)
	return o
}

// SetWithPluginInfo adds the withPluginInfo to the search payments params
func (o *SearchPaymentsParams) SetWithPluginInfo(withPluginInfo *bool) {
	o.WithPluginInfo = withPluginInfo
}

// WriteToRequest writes these params to a swagger request
func (o *SearchPaymentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.PluginName != nil {

		// query param pluginName
		var qrPluginName string
		if o.PluginName != nil {
			qrPluginName = *o.PluginName
		}
		qPluginName := qrPluginName
		if qPluginName != "" {
			if err := r.SetQueryParam("pluginName", qPluginName); err != nil {
				return err
			}
		}

	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	// path param searchKey
	if err := r.SetPathParam("searchKey", o.SearchKey); err != nil {
		return err
	}

	if o.WithAttempts != nil {

		// query param withAttempts
		var qrWithAttempts bool
		if o.WithAttempts != nil {
			qrWithAttempts = *o.WithAttempts
		}
		qWithAttempts := swag.FormatBool(qrWithAttempts)
		if qWithAttempts != "" {
			if err := r.SetQueryParam("withAttempts", qWithAttempts); err != nil {
				return err
			}
		}

	}

	if o.WithPluginInfo != nil {

		// query param withPluginInfo
		var qrWithPluginInfo bool
		if o.WithPluginInfo != nil {
			qrWithPluginInfo = *o.WithPluginInfo
		}
		qWithPluginInfo := swag.FormatBool(qrWithPluginInfo)
		if qWithPluginInfo != "" {
			if err := r.SetQueryParam("withPluginInfo", qWithPluginInfo); err != nil {
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
