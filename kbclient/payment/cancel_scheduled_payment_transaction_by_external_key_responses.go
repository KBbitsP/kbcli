// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// CancelScheduledPaymentTransactionByExternalKeyReader is a Reader for the CancelScheduledPaymentTransactionByExternalKey structure.
type CancelScheduledPaymentTransactionByExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelScheduledPaymentTransactionByExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCancelScheduledPaymentTransactionByExternalKeyNoContent()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewCancelScheduledPaymentTransactionByExternalKeyNoContent creates a CancelScheduledPaymentTransactionByExternalKeyNoContent with default headers values
func NewCancelScheduledPaymentTransactionByExternalKeyNoContent() *CancelScheduledPaymentTransactionByExternalKeyNoContent {
	return &CancelScheduledPaymentTransactionByExternalKeyNoContent{}
}

/*CancelScheduledPaymentTransactionByExternalKeyNoContent handles this case with default header values.

Successful operation
*/
type CancelScheduledPaymentTransactionByExternalKeyNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *CancelScheduledPaymentTransactionByExternalKeyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/payments/cancelScheduledPaymentTransaction][%d] cancelScheduledPaymentTransactionByExternalKeyNoContent ", 204)
}

func (o *CancelScheduledPaymentTransactionByExternalKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCancelScheduledPaymentTransactionByExternalKeyBadRequest creates a CancelScheduledPaymentTransactionByExternalKeyBadRequest with default headers values
func NewCancelScheduledPaymentTransactionByExternalKeyBadRequest() *CancelScheduledPaymentTransactionByExternalKeyBadRequest {
	return &CancelScheduledPaymentTransactionByExternalKeyBadRequest{}
}

/*CancelScheduledPaymentTransactionByExternalKeyBadRequest handles this case with default header values.

Invalid paymentTransactionExternalKey supplied
*/
type CancelScheduledPaymentTransactionByExternalKeyBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CancelScheduledPaymentTransactionByExternalKeyBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/payments/cancelScheduledPaymentTransaction][%d] cancelScheduledPaymentTransactionByExternalKeyBadRequest ", 400)
}

func (o *CancelScheduledPaymentTransactionByExternalKeyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
