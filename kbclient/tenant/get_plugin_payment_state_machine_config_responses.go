// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// GetPluginPaymentStateMachineConfigReader is a Reader for the GetPluginPaymentStateMachineConfig structure.
type GetPluginPaymentStateMachineConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPluginPaymentStateMachineConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPluginPaymentStateMachineConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetPluginPaymentStateMachineConfigBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPluginPaymentStateMachineConfigOK creates a GetPluginPaymentStateMachineConfigOK with default headers values
func NewGetPluginPaymentStateMachineConfigOK() *GetPluginPaymentStateMachineConfigOK {
	return &GetPluginPaymentStateMachineConfigOK{}
}

/*GetPluginPaymentStateMachineConfigOK handles this case with default header values.

successful operation
*/
type GetPluginPaymentStateMachineConfigOK struct {
	Payload *kbmodel.TenantKeyValue
}

func (o *GetPluginPaymentStateMachineConfigOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}][%d] getPluginPaymentStateMachineConfigOK  %+v", 200, o.Payload)
}

func (o *GetPluginPaymentStateMachineConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.TenantKeyValue)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginPaymentStateMachineConfigBadRequest creates a GetPluginPaymentStateMachineConfigBadRequest with default headers values
func NewGetPluginPaymentStateMachineConfigBadRequest() *GetPluginPaymentStateMachineConfigBadRequest {
	return &GetPluginPaymentStateMachineConfigBadRequest{}
}

/*GetPluginPaymentStateMachineConfigBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type GetPluginPaymentStateMachineConfigBadRequest struct {
}

func (o *GetPluginPaymentStateMachineConfigBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/tenants/uploadPluginPaymentStateMachineConfig/{pluginName}][%d] getPluginPaymentStateMachineConfigBadRequest ", 400)
}

func (o *GetPluginPaymentStateMachineConfigBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
