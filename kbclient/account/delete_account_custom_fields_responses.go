// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAccountCustomFieldsReader is a Reader for the DeleteAccountCustomFields structure.
type DeleteAccountCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccountCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAccountCustomFieldsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteAccountCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAccountCustomFieldsNoContent creates a DeleteAccountCustomFieldsNoContent with default headers values
func NewDeleteAccountCustomFieldsNoContent() *DeleteAccountCustomFieldsNoContent {
	return &DeleteAccountCustomFieldsNoContent{}
}

/*DeleteAccountCustomFieldsNoContent handles this case with default header values.

Successful operation
*/
type DeleteAccountCustomFieldsNoContent struct {
}

func (o *DeleteAccountCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}/customFields][%d] deleteAccountCustomFieldsNoContent ", 204)
}

func (o *DeleteAccountCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAccountCustomFieldsBadRequest creates a DeleteAccountCustomFieldsBadRequest with default headers values
func NewDeleteAccountCustomFieldsBadRequest() *DeleteAccountCustomFieldsBadRequest {
	return &DeleteAccountCustomFieldsBadRequest{}
}

/*DeleteAccountCustomFieldsBadRequest handles this case with default header values.

Invalid account id supplied
*/
type DeleteAccountCustomFieldsBadRequest struct {
}

func (o *DeleteAccountCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}/customFields][%d] deleteAccountCustomFieldsBadRequest ", 400)
}

func (o *DeleteAccountCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
