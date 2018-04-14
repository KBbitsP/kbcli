// Code generated by go-swagger; DO NOT EDIT.

package payment_method

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// GetPaymentMethodByKeyReader is a Reader for the GetPaymentMethodByKey structure.
type GetPaymentMethodByKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentMethodByKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentMethodByKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetPaymentMethodByKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPaymentMethodByKeyOK creates a GetPaymentMethodByKeyOK with default headers values
func NewGetPaymentMethodByKeyOK() *GetPaymentMethodByKeyOK {
	return &GetPaymentMethodByKeyOK{}
}

/*GetPaymentMethodByKeyOK handles this case with default header values.

successful operation
*/
type GetPaymentMethodByKeyOK struct {
	Payload *kbmodel.PaymentMethod
}

func (o *GetPaymentMethodByKeyOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods][%d] getPaymentMethodByKeyOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodByKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.PaymentMethod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentMethodByKeyNotFound creates a GetPaymentMethodByKeyNotFound with default headers values
func NewGetPaymentMethodByKeyNotFound() *GetPaymentMethodByKeyNotFound {
	return &GetPaymentMethodByKeyNotFound{}
}

/*GetPaymentMethodByKeyNotFound handles this case with default header values.

Account or payment method not found
*/
type GetPaymentMethodByKeyNotFound struct {
}

func (o *GetPaymentMethodByKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/paymentMethods][%d] getPaymentMethodByKeyNotFound ", 404)
}

func (o *GetPaymentMethodByKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
