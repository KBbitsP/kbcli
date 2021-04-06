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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteTransactionTagsParams creates a new DeleteTransactionTagsParams object
// with the default values initialized.
func NewDeleteTransactionTagsParams() *DeleteTransactionTagsParams {
	var ()
	return &DeleteTransactionTagsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTransactionTagsParamsWithTimeout creates a new DeleteTransactionTagsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteTransactionTagsParamsWithTimeout(timeout time.Duration) *DeleteTransactionTagsParams {
	var ()
	return &DeleteTransactionTagsParams{

		timeout: timeout,
	}
}

// NewDeleteTransactionTagsParamsWithContext creates a new DeleteTransactionTagsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteTransactionTagsParamsWithContext(ctx context.Context) *DeleteTransactionTagsParams {
	var ()
	return &DeleteTransactionTagsParams{

		Context: ctx,
	}
}

// NewDeleteTransactionTagsParamsWithHTTPClient creates a new DeleteTransactionTagsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteTransactionTagsParamsWithHTTPClient(client *http.Client) *DeleteTransactionTagsParams {
	var ()
	return &DeleteTransactionTagsParams{
		HTTPClient: client,
	}
}

/*DeleteTransactionTagsParams contains all the parameters to send to the API endpoint
for the delete transaction tags operation typically these are written to a http.Request
*/
type DeleteTransactionTagsParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*TagDef*/
	TagDef []strfmt.UUID
	/*TransactionID*/
	TransactionID strfmt.UUID

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the delete transaction tags params
func (o *DeleteTransactionTagsParams) WithTimeout(timeout time.Duration) *DeleteTransactionTagsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete transaction tags params
func (o *DeleteTransactionTagsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete transaction tags params
func (o *DeleteTransactionTagsParams) WithContext(ctx context.Context) *DeleteTransactionTagsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete transaction tags params
func (o *DeleteTransactionTagsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete transaction tags params
func (o *DeleteTransactionTagsParams) WithHTTPClient(client *http.Client) *DeleteTransactionTagsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete transaction tags params
func (o *DeleteTransactionTagsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the delete transaction tags params
func (o *DeleteTransactionTagsParams) WithXKillbillComment(xKillbillComment *string) *DeleteTransactionTagsParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the delete transaction tags params
func (o *DeleteTransactionTagsParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the delete transaction tags params
func (o *DeleteTransactionTagsParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *DeleteTransactionTagsParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the delete transaction tags params
func (o *DeleteTransactionTagsParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the delete transaction tags params
func (o *DeleteTransactionTagsParams) WithXKillbillReason(xKillbillReason *string) *DeleteTransactionTagsParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the delete transaction tags params
func (o *DeleteTransactionTagsParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithTagDef adds the tagDef to the delete transaction tags params
func (o *DeleteTransactionTagsParams) WithTagDef(tagDef []strfmt.UUID) *DeleteTransactionTagsParams {
	o.SetTagDef(tagDef)
	return o
}

// SetTagDef adds the tagDef to the delete transaction tags params
func (o *DeleteTransactionTagsParams) SetTagDef(tagDef []strfmt.UUID) {
	o.TagDef = tagDef
}

// WithTransactionID adds the transactionID to the delete transaction tags params
func (o *DeleteTransactionTagsParams) WithTransactionID(transactionID strfmt.UUID) *DeleteTransactionTagsParams {
	o.SetTransactionID(transactionID)
	return o
}

// SetTransactionID adds the transactionId to the delete transaction tags params
func (o *DeleteTransactionTagsParams) SetTransactionID(transactionID strfmt.UUID) {
	o.TransactionID = transactionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTransactionTagsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesTagDef []string
	for _, v := range o.TagDef {
		valuesTagDef = append(valuesTagDef, v.String())
	}

	joinedTagDef := swag.JoinByFormat(valuesTagDef, "multi")
	// query array param tagDef
	if err := r.SetQueryParam("tagDef", joinedTagDef...); err != nil {
		return err
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
