// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// GetInvoicePaymentCustomFieldsReader is a Reader for the GetInvoicePaymentCustomFields structure.
type GetInvoicePaymentCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoicePaymentCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoicePaymentCustomFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetInvoicePaymentCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInvoicePaymentCustomFieldsOK creates a GetInvoicePaymentCustomFieldsOK with default headers values
func NewGetInvoicePaymentCustomFieldsOK() *GetInvoicePaymentCustomFieldsOK {
	return &GetInvoicePaymentCustomFieldsOK{}
}

/*GetInvoicePaymentCustomFieldsOK handles this case with default header values.

successful operation
*/
type GetInvoicePaymentCustomFieldsOK struct {
	Payload []*kbmodel.CustomField
}

func (o *GetInvoicePaymentCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoicePayments/{paymentId}/customFields][%d] getInvoicePaymentCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetInvoicePaymentCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicePaymentCustomFieldsBadRequest creates a GetInvoicePaymentCustomFieldsBadRequest with default headers values
func NewGetInvoicePaymentCustomFieldsBadRequest() *GetInvoicePaymentCustomFieldsBadRequest {
	return &GetInvoicePaymentCustomFieldsBadRequest{}
}

/*GetInvoicePaymentCustomFieldsBadRequest handles this case with default header values.

Invalid payment id supplied
*/
type GetInvoicePaymentCustomFieldsBadRequest struct {
}

func (o *GetInvoicePaymentCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoicePayments/{paymentId}/customFields][%d] getInvoicePaymentCustomFieldsBadRequest ", 400)
}

func (o *GetInvoicePaymentCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
