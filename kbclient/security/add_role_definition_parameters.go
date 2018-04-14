// Code generated by go-swagger; DO NOT EDIT.

package security

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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// NewAddRoleDefinitionParams creates a new AddRoleDefinitionParams object
// with the default values initialized.
func NewAddRoleDefinitionParams() *AddRoleDefinitionParams {
	var ()
	return &AddRoleDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddRoleDefinitionParamsWithTimeout creates a new AddRoleDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddRoleDefinitionParamsWithTimeout(timeout time.Duration) *AddRoleDefinitionParams {
	var ()
	return &AddRoleDefinitionParams{

		timeout: timeout,
	}
}

// NewAddRoleDefinitionParamsWithContext creates a new AddRoleDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddRoleDefinitionParamsWithContext(ctx context.Context) *AddRoleDefinitionParams {
	var ()
	return &AddRoleDefinitionParams{

		Context: ctx,
	}
}

// NewAddRoleDefinitionParamsWithHTTPClient creates a new AddRoleDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddRoleDefinitionParamsWithHTTPClient(client *http.Client) *AddRoleDefinitionParams {
	var ()
	return &AddRoleDefinitionParams{
		HTTPClient: client,
	}
}

/*AddRoleDefinitionParams contains all the parameters to send to the API endpoint
for the add role definition operation typically these are written to a http.Request
*/
type AddRoleDefinitionParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.RoleDefinition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add role definition params
func (o *AddRoleDefinitionParams) WithTimeout(timeout time.Duration) *AddRoleDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add role definition params
func (o *AddRoleDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add role definition params
func (o *AddRoleDefinitionParams) WithContext(ctx context.Context) *AddRoleDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add role definition params
func (o *AddRoleDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add role definition params
func (o *AddRoleDefinitionParams) WithHTTPClient(client *http.Client) *AddRoleDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add role definition params
func (o *AddRoleDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the add role definition params
func (o *AddRoleDefinitionParams) WithXKillbillComment(xKillbillComment *string) *AddRoleDefinitionParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the add role definition params
func (o *AddRoleDefinitionParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the add role definition params
func (o *AddRoleDefinitionParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *AddRoleDefinitionParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the add role definition params
func (o *AddRoleDefinitionParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the add role definition params
func (o *AddRoleDefinitionParams) WithXKillbillReason(xKillbillReason *string) *AddRoleDefinitionParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the add role definition params
func (o *AddRoleDefinitionParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the add role definition params
func (o *AddRoleDefinitionParams) WithBody(body *kbmodel.RoleDefinition) *AddRoleDefinitionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add role definition params
func (o *AddRoleDefinitionParams) SetBody(body *kbmodel.RoleDefinition) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddRoleDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
