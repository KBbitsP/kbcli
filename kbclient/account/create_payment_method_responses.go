// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// CreatePaymentMethodReader is a Reader for the CreatePaymentMethod structure.
type CreatePaymentMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePaymentMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreatePaymentMethodCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreatePaymentMethodBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCreatePaymentMethodNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePaymentMethodCreated creates a CreatePaymentMethodCreated with default headers values
func NewCreatePaymentMethodCreated() *CreatePaymentMethodCreated {
	return &CreatePaymentMethodCreated{}
}

/*CreatePaymentMethodCreated handles this case with default header values.

Payment method created
*/
type CreatePaymentMethodCreated struct {
	Payload *kbmodel.PaymentMethod
}

func (o *CreatePaymentMethodCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/paymentMethods][%d] createPaymentMethodCreated  %+v", 201, o.Payload)
}

func (o *CreatePaymentMethodCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.PaymentMethod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreatePaymentMethodBadRequest creates a CreatePaymentMethodBadRequest with default headers values
func NewCreatePaymentMethodBadRequest() *CreatePaymentMethodBadRequest {
	return &CreatePaymentMethodBadRequest{}
}

/*CreatePaymentMethodBadRequest handles this case with default header values.

Invalid account id supplied
*/
type CreatePaymentMethodBadRequest struct {
}

func (o *CreatePaymentMethodBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/paymentMethods][%d] createPaymentMethodBadRequest ", 400)
}

func (o *CreatePaymentMethodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreatePaymentMethodNotFound creates a CreatePaymentMethodNotFound with default headers values
func NewCreatePaymentMethodNotFound() *CreatePaymentMethodNotFound {
	return &CreatePaymentMethodNotFound{}
}

/*CreatePaymentMethodNotFound handles this case with default header values.

Account not found
*/
type CreatePaymentMethodNotFound struct {
}

func (o *CreatePaymentMethodNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/paymentMethods][%d] createPaymentMethodNotFound ", 404)
}

func (o *CreatePaymentMethodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
