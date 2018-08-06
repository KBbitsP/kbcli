// Code generated by go-swagger; DO NOT EDIT.

package credit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new credit API client.
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *Client {

	return &Client{transport: transport, formats: formats, authInfo: authInfo, defaults: defaults}
}

// killbill default values. When a call is made to an operation, these values are used
// if params doesn't specify them.
type KillbillDefaults interface {
	// Default CreatedBy. If not set explicitly in params, this will be used.
	XKillbillCreatedBy() *string
	// Default Comment. If not set explicitly in params, this will be used.
	XKillbillComment() *string
	// Default Reason. If not set explicitly in params, this will be used.
	XKillbillReason() *string
	// Default WithStackTrace. If not set explicitly in params, this will be used.
	KillbillWithStackTrace() *bool
}

/*
Client for credit API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	defaults  KillbillDefaults
}

// ICredit - interface for Credit client.
type ICredit interface {
	/*
		CreateCredit creates a credit
	*/
	CreateCredit(ctx context.Context, params *CreateCreditParams) (*CreateCreditCreated, error)

	/*
		GetCredit retrieves a credit by id
	*/
	GetCredit(ctx context.Context, params *GetCreditParams) (*GetCreditOK, error)
}

/*
CreateCredit creates a credit
*/
func (a *Client) CreateCredit(ctx context.Context, params *CreateCreditParams) (*CreateCreditCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateCreditParams()
	}
	getParams := NewCreateCreditParams()
	getParams.Context = ctx
	params.Context = ctx
	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}
	getParams.XKillbillComment = params.XKillbillComment
	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}
	getParams.XKillbillCreatedBy = params.XKillbillCreatedBy
	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}
	getParams.XKillbillReason = params.XKillbillReason
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}
	getParams.WithStackTrace = params.WithStackTrace

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCredit",
		Method:             "POST",
		PathPattern:        "/1.0/kb/credits",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateCreditReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*CreateCreditCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createCredit",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &CreateCreditReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*CreateCreditCreated), nil

}

/*
GetCredit retrieves a credit by id
*/
func (a *Client) GetCredit(ctx context.Context, params *GetCreditParams) (*GetCreditOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCreditParams()
	}
	params.Context = ctx
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCredit",
		Method:             "GET",
		PathPattern:        "/1.0/kb/credits/{creditId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCreditReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCreditOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
