// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// RenameExternalKeyReader is a Reader for the RenameExternalKey structure.
type RenameExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RenameExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRenameExternalKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewRenameExternalKeyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewRenameExternalKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRenameExternalKeyNoContent creates a RenameExternalKeyNoContent with default headers values
func NewRenameExternalKeyNoContent() *RenameExternalKeyNoContent {
	return &RenameExternalKeyNoContent{}
}

/*RenameExternalKeyNoContent handles this case with default header values.

Successful operation
*/
type RenameExternalKeyNoContent struct {
}

func (o *RenameExternalKeyNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/renameKey][%d] renameExternalKeyNoContent ", 204)
}

func (o *RenameExternalKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenameExternalKeyBadRequest creates a RenameExternalKeyBadRequest with default headers values
func NewRenameExternalKeyBadRequest() *RenameExternalKeyBadRequest {
	return &RenameExternalKeyBadRequest{}
}

/*RenameExternalKeyBadRequest handles this case with default header values.

Invalid argumnent supplied
*/
type RenameExternalKeyBadRequest struct {
}

func (o *RenameExternalKeyBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/renameKey][%d] renameExternalKeyBadRequest ", 400)
}

func (o *RenameExternalKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRenameExternalKeyNotFound creates a RenameExternalKeyNotFound with default headers values
func NewRenameExternalKeyNotFound() *RenameExternalKeyNotFound {
	return &RenameExternalKeyNotFound{}
}

/*RenameExternalKeyNotFound handles this case with default header values.

Bundle not found
*/
type RenameExternalKeyNotFound struct {
}

func (o *RenameExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/renameKey][%d] renameExternalKeyNotFound ", 404)
}

func (o *RenameExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
