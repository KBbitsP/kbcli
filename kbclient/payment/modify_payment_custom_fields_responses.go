// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ModifyPaymentCustomFieldsReader is a Reader for the ModifyPaymentCustomFields structure.
type ModifyPaymentCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyPaymentCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewModifyPaymentCustomFieldsNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewModifyPaymentCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyPaymentCustomFieldsNoContent creates a ModifyPaymentCustomFieldsNoContent with default headers values
func NewModifyPaymentCustomFieldsNoContent() *ModifyPaymentCustomFieldsNoContent {
	return &ModifyPaymentCustomFieldsNoContent{}
}

/*ModifyPaymentCustomFieldsNoContent handles this case with default header values.

Successful operation
*/
type ModifyPaymentCustomFieldsNoContent struct {
}

func (o *ModifyPaymentCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}/customFields][%d] modifyPaymentCustomFieldsNoContent ", 204)
}

func (o *ModifyPaymentCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewModifyPaymentCustomFieldsBadRequest creates a ModifyPaymentCustomFieldsBadRequest with default headers values
func NewModifyPaymentCustomFieldsBadRequest() *ModifyPaymentCustomFieldsBadRequest {
	return &ModifyPaymentCustomFieldsBadRequest{}
}

/*ModifyPaymentCustomFieldsBadRequest handles this case with default header values.

Invalid payment id supplied
*/
type ModifyPaymentCustomFieldsBadRequest struct {
}

func (o *ModifyPaymentCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments/{paymentId}/customFields][%d] modifyPaymentCustomFieldsBadRequest ", 400)
}

func (o *ModifyPaymentCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
