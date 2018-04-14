// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// InvalidateUserReader is a Reader for the InvalidateUser structure.
type InvalidateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvalidateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewInvalidateUserNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInvalidateUserNoContent creates a InvalidateUserNoContent with default headers values
func NewInvalidateUserNoContent() *InvalidateUserNoContent {
	return &InvalidateUserNoContent{}
}

/*InvalidateUserNoContent handles this case with default header values.

Successful operation
*/
type InvalidateUserNoContent struct {
}

func (o *InvalidateUserNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/security/users/{username}][%d] invalidateUserNoContent ", 204)
}

func (o *InvalidateUserNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
