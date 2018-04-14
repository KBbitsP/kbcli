// Code generated by go-swagger; DO NOT EDIT.

package invoice

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

// NewGetInvoiceAsHTMLParams creates a new GetInvoiceAsHTMLParams object
// with the default values initialized.
func NewGetInvoiceAsHTMLParams() *GetInvoiceAsHTMLParams {
	var ()
	return &GetInvoiceAsHTMLParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceAsHTMLParamsWithTimeout creates a new GetInvoiceAsHTMLParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoiceAsHTMLParamsWithTimeout(timeout time.Duration) *GetInvoiceAsHTMLParams {
	var ()
	return &GetInvoiceAsHTMLParams{

		timeout: timeout,
	}
}

// NewGetInvoiceAsHTMLParamsWithContext creates a new GetInvoiceAsHTMLParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoiceAsHTMLParamsWithContext(ctx context.Context) *GetInvoiceAsHTMLParams {
	var ()
	return &GetInvoiceAsHTMLParams{

		Context: ctx,
	}
}

// NewGetInvoiceAsHTMLParamsWithHTTPClient creates a new GetInvoiceAsHTMLParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoiceAsHTMLParamsWithHTTPClient(client *http.Client) *GetInvoiceAsHTMLParams {
	var ()
	return &GetInvoiceAsHTMLParams{
		HTTPClient: client,
	}
}

/*GetInvoiceAsHTMLParams contains all the parameters to send to the API endpoint
for the get invoice as HTML operation typically these are written to a http.Request
*/
type GetInvoiceAsHTMLParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*InvoiceID*/
	InvoiceID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) WithTimeout(timeout time.Duration) *GetInvoiceAsHTMLParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) WithContext(ctx context.Context) *GetInvoiceAsHTMLParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) WithHTTPClient(client *http.Client) *GetInvoiceAsHTMLParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetInvoiceAsHTMLParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetInvoiceAsHTMLParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithInvoiceID adds the invoiceID to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) WithInvoiceID(invoiceID strfmt.UUID) *GetInvoiceAsHTMLParams {
	o.SetInvoiceID(invoiceID)
	return o
}

// SetInvoiceID adds the invoiceId to the get invoice as HTML params
func (o *GetInvoiceAsHTMLParams) SetInvoiceID(invoiceID strfmt.UUID) {
	o.InvoiceID = invoiceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceAsHTMLParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param invoiceId
	if err := r.SetPathParam("invoiceId", o.InvoiceID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
