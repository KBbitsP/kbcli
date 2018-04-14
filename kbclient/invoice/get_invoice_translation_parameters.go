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

// NewGetInvoiceTranslationParams creates a new GetInvoiceTranslationParams object
// with the default values initialized.
func NewGetInvoiceTranslationParams() *GetInvoiceTranslationParams {
	var ()
	return &GetInvoiceTranslationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInvoiceTranslationParamsWithTimeout creates a new GetInvoiceTranslationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInvoiceTranslationParamsWithTimeout(timeout time.Duration) *GetInvoiceTranslationParams {
	var ()
	return &GetInvoiceTranslationParams{

		timeout: timeout,
	}
}

// NewGetInvoiceTranslationParamsWithContext creates a new GetInvoiceTranslationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInvoiceTranslationParamsWithContext(ctx context.Context) *GetInvoiceTranslationParams {
	var ()
	return &GetInvoiceTranslationParams{

		Context: ctx,
	}
}

// NewGetInvoiceTranslationParamsWithHTTPClient creates a new GetInvoiceTranslationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInvoiceTranslationParamsWithHTTPClient(client *http.Client) *GetInvoiceTranslationParams {
	var ()
	return &GetInvoiceTranslationParams{
		HTTPClient: client,
	}
}

/*GetInvoiceTranslationParams contains all the parameters to send to the API endpoint
for the get invoice translation operation typically these are written to a http.Request
*/
type GetInvoiceTranslationParams struct {

	/*XKillbillAPIKey*/
	XKillbillAPIKey string
	/*XKillbillAPISecret*/
	XKillbillAPISecret string
	/*Locale*/
	Locale string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get invoice translation params
func (o *GetInvoiceTranslationParams) WithTimeout(timeout time.Duration) *GetInvoiceTranslationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get invoice translation params
func (o *GetInvoiceTranslationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get invoice translation params
func (o *GetInvoiceTranslationParams) WithContext(ctx context.Context) *GetInvoiceTranslationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get invoice translation params
func (o *GetInvoiceTranslationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get invoice translation params
func (o *GetInvoiceTranslationParams) WithHTTPClient(client *http.Client) *GetInvoiceTranslationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get invoice translation params
func (o *GetInvoiceTranslationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillAPIKey adds the xKillbillAPIKey to the get invoice translation params
func (o *GetInvoiceTranslationParams) WithXKillbillAPIKey(xKillbillAPIKey string) *GetInvoiceTranslationParams {
	o.SetXKillbillAPIKey(xKillbillAPIKey)
	return o
}

// SetXKillbillAPIKey adds the xKillbillApiKey to the get invoice translation params
func (o *GetInvoiceTranslationParams) SetXKillbillAPIKey(xKillbillAPIKey string) {
	o.XKillbillAPIKey = xKillbillAPIKey
}

// WithXKillbillAPISecret adds the xKillbillAPISecret to the get invoice translation params
func (o *GetInvoiceTranslationParams) WithXKillbillAPISecret(xKillbillAPISecret string) *GetInvoiceTranslationParams {
	o.SetXKillbillAPISecret(xKillbillAPISecret)
	return o
}

// SetXKillbillAPISecret adds the xKillbillApiSecret to the get invoice translation params
func (o *GetInvoiceTranslationParams) SetXKillbillAPISecret(xKillbillAPISecret string) {
	o.XKillbillAPISecret = xKillbillAPISecret
}

// WithLocale adds the locale to the get invoice translation params
func (o *GetInvoiceTranslationParams) WithLocale(locale string) *GetInvoiceTranslationParams {
	o.SetLocale(locale)
	return o
}

// SetLocale adds the locale to the get invoice translation params
func (o *GetInvoiceTranslationParams) SetLocale(locale string) {
	o.Locale = locale
}

// WriteToRequest writes these params to a swagger request
func (o *GetInvoiceTranslationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param locale
	if err := r.SetPathParam("locale", o.Locale); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
