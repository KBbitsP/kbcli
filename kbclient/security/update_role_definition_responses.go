// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// UpdateRoleDefinitionReader is a Reader for the UpdateRoleDefinition structure.
type UpdateRoleDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoleDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewUpdateRoleDefinitionNoContent()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewUpdateRoleDefinitionNoContent creates a UpdateRoleDefinitionNoContent with default headers values
func NewUpdateRoleDefinitionNoContent() *UpdateRoleDefinitionNoContent {
	return &UpdateRoleDefinitionNoContent{}
}

/*UpdateRoleDefinitionNoContent handles this case with default header values.

Successful operation
*/
type UpdateRoleDefinitionNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *UpdateRoleDefinitionNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/security/roles][%d] updateRoleDefinitionNoContent ", 204)
}

func (o *UpdateRoleDefinitionNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
