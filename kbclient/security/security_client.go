// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new security API client.
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *Client {

	return &Client{transport: transport, formats: formats, authInfo: authInfo, defaults: defaults}
}

// killbill default values. When a call is made to an operation, these values are used
// if params doesn't specify them.
type KillbillDefaults interface {
	// Default API Key. If not set explicitly in params, this will be used.
	XKillbillAPIKey() *string
	// Default API Secret. If not set explicitly in params, this will be used.
	XKillbillAPISecret() *string
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
Client for security API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	defaults  KillbillDefaults
}

// ISecurity - interface for Security client.
type ISecurity interface {
	/*
		AddRoleDefinition adds a new role definition
	*/
	AddRoleDefinition(ctx context.Context, params *AddRoleDefinitionParams) (*AddRoleDefinitionCreated, error)

	/*
		AddUserRoles adds a new user with roles to make api requests
	*/
	AddUserRoles(ctx context.Context, params *AddUserRolesParams) (*AddUserRolesCreated, error)

	/*
		GetCurrentUserPermissions lists user permissions
	*/
	GetCurrentUserPermissions(ctx context.Context, params *GetCurrentUserPermissionsParams) (*GetCurrentUserPermissionsOK, error)

	/*
		GetCurrentUserSubject gets user information
	*/
	GetCurrentUserSubject(ctx context.Context, params *GetCurrentUserSubjectParams) (*GetCurrentUserSubjectOK, error)

	/*
		GetRoleDefinition gets role definition
	*/
	GetRoleDefinition(ctx context.Context, params *GetRoleDefinitionParams) (*GetRoleDefinitionOK, error)

	/*
		GetUserRoles gets roles associated to a user
	*/
	GetUserRoles(ctx context.Context, params *GetUserRolesParams) (*GetUserRolesOK, error)

	/*
		InvalidateUser invalidates an existing user
	*/
	InvalidateUser(ctx context.Context, params *InvalidateUserParams) (*InvalidateUserNoContent, error)

	/*
		UpdateRoleDefinition updates a new role definition
	*/
	UpdateRoleDefinition(ctx context.Context, params *UpdateRoleDefinitionParams) (*UpdateRoleDefinitionNoContent, error)

	/*
		UpdateUserPassword updates a user password
	*/
	UpdateUserPassword(ctx context.Context, params *UpdateUserPasswordParams) (*UpdateUserPasswordNoContent, error)

	/*
		UpdateUserRoles updates roles associated to a user
	*/
	UpdateUserRoles(ctx context.Context, params *UpdateUserRolesParams) (*UpdateUserRolesNoContent, error)
}

/*
AddRoleDefinition adds a new role definition
*/
func (a *Client) AddRoleDefinition(ctx context.Context, params *AddRoleDefinitionParams) (*AddRoleDefinitionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddRoleDefinitionParams()
	}
	getParams := NewAddRoleDefinitionParams()
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
		ID:                 "addRoleDefinition",
		Method:             "POST",
		PathPattern:        "/1.0/kb/security/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddRoleDefinitionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*AddRoleDefinitionCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addRoleDefinition",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &AddRoleDefinitionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*AddRoleDefinitionCreated), nil

}

/*
AddUserRoles adds a new user with roles to make api requests
*/
func (a *Client) AddUserRoles(ctx context.Context, params *AddUserRolesParams) (*AddUserRolesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddUserRolesParams()
	}
	getParams := NewAddUserRolesParams()
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
		ID:                 "addUserRoles",
		Method:             "POST",
		PathPattern:        "/1.0/kb/security/users",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddUserRolesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	createdResult := result.(*AddUserRolesCreated)
	location := kbcommon.ParseLocationHeader(createdResult.HttpResponse.GetHeader("Location"))
	if !params.ProcessLocationHeader || location == "" {
		return createdResult, nil
	}

	getResult, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addUserRoles",
		Method:             "GET",
		PathPattern:        location,
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             getParams,
		Reader:             &AddUserRolesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            getParams.Context,
		Client:             getParams.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return getResult.(*AddUserRolesCreated), nil

}

/*
GetCurrentUserPermissions lists user permissions
*/
func (a *Client) GetCurrentUserPermissions(ctx context.Context, params *GetCurrentUserPermissionsParams) (*GetCurrentUserPermissionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentUserPermissionsParams()
	}
	params.Context = ctx
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrentUserPermissions",
		Method:             "GET",
		PathPattern:        "/1.0/kb/security/permissions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCurrentUserPermissionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCurrentUserPermissionsOK), nil

}

/*
GetCurrentUserSubject gets user information
*/
func (a *Client) GetCurrentUserSubject(ctx context.Context, params *GetCurrentUserSubjectParams) (*GetCurrentUserSubjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentUserSubjectParams()
	}
	params.Context = ctx
	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrentUserSubject",
		Method:             "GET",
		PathPattern:        "/1.0/kb/security/subject",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetCurrentUserSubjectReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCurrentUserSubjectOK), nil

}

/*
GetRoleDefinition gets role definition
*/
func (a *Client) GetRoleDefinition(ctx context.Context, params *GetRoleDefinitionParams) (*GetRoleDefinitionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRoleDefinitionParams()
	}
	params.Context = ctx

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRoleDefinition",
		Method:             "GET",
		PathPattern:        "/1.0/kb/security/roles/{role}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRoleDefinitionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRoleDefinitionOK), nil

}

/*
GetUserRoles gets roles associated to a user
*/
func (a *Client) GetUserRoles(ctx context.Context, params *GetUserRolesParams) (*GetUserRolesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserRolesParams()
	}
	params.Context = ctx

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUserRoles",
		Method:             "GET",
		PathPattern:        "/1.0/kb/security/users/{username}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetUserRolesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserRolesOK), nil

}

/*
InvalidateUser invalidates an existing user
*/
func (a *Client) InvalidateUser(ctx context.Context, params *InvalidateUserParams) (*InvalidateUserNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInvalidateUserParams()
	}
	params.Context = ctx

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "invalidateUser",
		Method:             "DELETE",
		PathPattern:        "/1.0/kb/security/users/{username}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InvalidateUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InvalidateUserNoContent), nil

}

/*
UpdateRoleDefinition updates a new role definition
*/
func (a *Client) UpdateRoleDefinition(ctx context.Context, params *UpdateRoleDefinitionParams) (*UpdateRoleDefinitionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRoleDefinitionParams()
	}
	params.Context = ctx

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateRoleDefinition",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/security/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRoleDefinitionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateRoleDefinitionNoContent), nil

}

/*
UpdateUserPassword updates a user password
*/
func (a *Client) UpdateUserPassword(ctx context.Context, params *UpdateUserPasswordParams) (*UpdateUserPasswordNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserPasswordParams()
	}
	params.Context = ctx

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserPassword",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/security/users/{username}/password",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateUserPasswordReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateUserPasswordNoContent), nil

}

/*
UpdateUserRoles updates roles associated to a user
*/
func (a *Client) UpdateUserRoles(ctx context.Context, params *UpdateUserRolesParams) (*UpdateUserRolesNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateUserRolesParams()
	}
	params.Context = ctx

	if params.XKillbillComment == nil && a.defaults.XKillbillComment() != nil {
		params.XKillbillComment = a.defaults.XKillbillComment()
	}

	if params.XKillbillCreatedBy == "" && a.defaults.XKillbillCreatedBy() != nil {
		params.XKillbillCreatedBy = *a.defaults.XKillbillCreatedBy()
	}

	if params.XKillbillReason == nil && a.defaults.XKillbillReason() != nil {
		params.XKillbillReason = a.defaults.XKillbillReason()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateUserRoles",
		Method:             "PUT",
		PathPattern:        "/1.0/kb/security/users/{username}/roles",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateUserRolesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateUserRolesNoContent), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
