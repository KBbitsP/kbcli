// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// AddRoleDefinitionReader is a Reader for the AddRoleDefinition structure.
type AddRoleDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddRoleDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddRoleDefinitionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddRoleDefinitionCreated creates a AddRoleDefinitionCreated with default headers values
func NewAddRoleDefinitionCreated() *AddRoleDefinitionCreated {
	return &AddRoleDefinitionCreated{}
}

/*AddRoleDefinitionCreated handles this case with default header values.

Role definition created successfully
*/
type AddRoleDefinitionCreated struct {
	Payload *kbmodel.RoleDefinition
}

func (o *AddRoleDefinitionCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/security/roles][%d] addRoleDefinitionCreated  %+v", 201, o.Payload)
}

func (o *AddRoleDefinitionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.RoleDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
