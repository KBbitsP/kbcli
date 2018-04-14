// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteBundleCustomFieldsReader is a Reader for the DeleteBundleCustomFields structure.
type DeleteBundleCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteBundleCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteBundleCustomFieldsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteBundleCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteBundleCustomFieldsNoContent creates a DeleteBundleCustomFieldsNoContent with default headers values
func NewDeleteBundleCustomFieldsNoContent() *DeleteBundleCustomFieldsNoContent {
	return &DeleteBundleCustomFieldsNoContent{}
}

/*DeleteBundleCustomFieldsNoContent handles this case with default header values.

Successful operation
*/
type DeleteBundleCustomFieldsNoContent struct {
}

func (o *DeleteBundleCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/bundles/{bundleId}/customFields][%d] deleteBundleCustomFieldsNoContent ", 204)
}

func (o *DeleteBundleCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteBundleCustomFieldsBadRequest creates a DeleteBundleCustomFieldsBadRequest with default headers values
func NewDeleteBundleCustomFieldsBadRequest() *DeleteBundleCustomFieldsBadRequest {
	return &DeleteBundleCustomFieldsBadRequest{}
}

/*DeleteBundleCustomFieldsBadRequest handles this case with default header values.

Invalid bundle id supplied
*/
type DeleteBundleCustomFieldsBadRequest struct {
}

func (o *DeleteBundleCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/bundles/{bundleId}/customFields][%d] deleteBundleCustomFieldsBadRequest ", 400)
}

func (o *DeleteBundleCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
