// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteUserKeyValueReader is a Reader for the DeleteUserKeyValue structure.
type DeleteUserKeyValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserKeyValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserKeyValueNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteUserKeyValueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteUserKeyValueNoContent creates a DeleteUserKeyValueNoContent with default headers values
func NewDeleteUserKeyValueNoContent() *DeleteUserKeyValueNoContent {
	return &DeleteUserKeyValueNoContent{}
}

/*DeleteUserKeyValueNoContent handles this case with default header values.

Successful operation
*/
type DeleteUserKeyValueNoContent struct {
}

func (o *DeleteUserKeyValueNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/userKeyValue/{keyName}][%d] deleteUserKeyValueNoContent ", 204)
}

func (o *DeleteUserKeyValueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserKeyValueBadRequest creates a DeleteUserKeyValueBadRequest with default headers values
func NewDeleteUserKeyValueBadRequest() *DeleteUserKeyValueBadRequest {
	return &DeleteUserKeyValueBadRequest{}
}

/*DeleteUserKeyValueBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type DeleteUserKeyValueBadRequest struct {
}

func (o *DeleteUserKeyValueBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/userKeyValue/{keyName}][%d] deleteUserKeyValueBadRequest ", 400)
}

func (o *DeleteUserKeyValueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
