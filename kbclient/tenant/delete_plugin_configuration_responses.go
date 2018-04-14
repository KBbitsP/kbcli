// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePluginConfigurationReader is a Reader for the DeletePluginConfiguration structure.
type DeletePluginConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePluginConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeletePluginConfigurationNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeletePluginConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeletePluginConfigurationNoContent creates a DeletePluginConfigurationNoContent with default headers values
func NewDeletePluginConfigurationNoContent() *DeletePluginConfigurationNoContent {
	return &DeletePluginConfigurationNoContent{}
}

/*DeletePluginConfigurationNoContent handles this case with default header values.

Successful operation
*/
type DeletePluginConfigurationNoContent struct {
}

func (o *DeletePluginConfigurationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/uploadPluginConfig/{pluginName}][%d] deletePluginConfigurationNoContent ", 204)
}

func (o *DeletePluginConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePluginConfigurationBadRequest creates a DeletePluginConfigurationBadRequest with default headers values
func NewDeletePluginConfigurationBadRequest() *DeletePluginConfigurationBadRequest {
	return &DeletePluginConfigurationBadRequest{}
}

/*DeletePluginConfigurationBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type DeletePluginConfigurationBadRequest struct {
}

func (o *DeletePluginConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/uploadPluginConfig/{pluginName}][%d] deletePluginConfigurationBadRequest ", 400)
}

func (o *DeletePluginConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
