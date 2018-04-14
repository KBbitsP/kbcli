// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteAccountTagsReader is a Reader for the DeleteAccountTags structure.
type DeleteAccountTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAccountTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteAccountTagsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteAccountTagsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAccountTagsNoContent creates a DeleteAccountTagsNoContent with default headers values
func NewDeleteAccountTagsNoContent() *DeleteAccountTagsNoContent {
	return &DeleteAccountTagsNoContent{}
}

/*DeleteAccountTagsNoContent handles this case with default header values.

Successful operation
*/
type DeleteAccountTagsNoContent struct {
}

func (o *DeleteAccountTagsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}/tags][%d] deleteAccountTagsNoContent ", 204)
}

func (o *DeleteAccountTagsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteAccountTagsBadRequest creates a DeleteAccountTagsBadRequest with default headers values
func NewDeleteAccountTagsBadRequest() *DeleteAccountTagsBadRequest {
	return &DeleteAccountTagsBadRequest{}
}

/*DeleteAccountTagsBadRequest handles this case with default header values.

Invalid account id supplied or account does not have a default payment method (AUTO_PAY_OFF tag only)
*/
type DeleteAccountTagsBadRequest struct {
}

func (o *DeleteAccountTagsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/accounts/{accountId}/tags][%d] deleteAccountTagsBadRequest ", 400)
}

func (o *DeleteAccountTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
