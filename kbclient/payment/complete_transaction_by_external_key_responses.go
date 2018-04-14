// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CompleteTransactionByExternalKeyReader is a Reader for the CompleteTransactionByExternalKey structure.
type CompleteTransactionByExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteTransactionByExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewCompleteTransactionByExternalKeyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 402:
		result := NewCompleteTransactionByExternalKeyPaymentRequired()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewCompleteTransactionByExternalKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewCompleteTransactionByExternalKeyUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 502:
		result := NewCompleteTransactionByExternalKeyBadGateway()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewCompleteTransactionByExternalKeyServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewCompleteTransactionByExternalKeyGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompleteTransactionByExternalKeyNoContent creates a CompleteTransactionByExternalKeyNoContent with default headers values
func NewCompleteTransactionByExternalKeyNoContent() *CompleteTransactionByExternalKeyNoContent {
	return &CompleteTransactionByExternalKeyNoContent{}
}

/*CompleteTransactionByExternalKeyNoContent handles this case with default header values.

Successful operation
*/
type CompleteTransactionByExternalKeyNoContent struct {
}

func (o *CompleteTransactionByExternalKeyNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments][%d] completeTransactionByExternalKeyNoContent ", 204)
}

func (o *CompleteTransactionByExternalKeyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionByExternalKeyPaymentRequired creates a CompleteTransactionByExternalKeyPaymentRequired with default headers values
func NewCompleteTransactionByExternalKeyPaymentRequired() *CompleteTransactionByExternalKeyPaymentRequired {
	return &CompleteTransactionByExternalKeyPaymentRequired{}
}

/*CompleteTransactionByExternalKeyPaymentRequired handles this case with default header values.

Transaction declined by gateway
*/
type CompleteTransactionByExternalKeyPaymentRequired struct {
}

func (o *CompleteTransactionByExternalKeyPaymentRequired) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments][%d] completeTransactionByExternalKeyPaymentRequired ", 402)
}

func (o *CompleteTransactionByExternalKeyPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionByExternalKeyNotFound creates a CompleteTransactionByExternalKeyNotFound with default headers values
func NewCompleteTransactionByExternalKeyNotFound() *CompleteTransactionByExternalKeyNotFound {
	return &CompleteTransactionByExternalKeyNotFound{}
}

/*CompleteTransactionByExternalKeyNotFound handles this case with default header values.

Account or payment not found
*/
type CompleteTransactionByExternalKeyNotFound struct {
}

func (o *CompleteTransactionByExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments][%d] completeTransactionByExternalKeyNotFound ", 404)
}

func (o *CompleteTransactionByExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionByExternalKeyUnprocessableEntity creates a CompleteTransactionByExternalKeyUnprocessableEntity with default headers values
func NewCompleteTransactionByExternalKeyUnprocessableEntity() *CompleteTransactionByExternalKeyUnprocessableEntity {
	return &CompleteTransactionByExternalKeyUnprocessableEntity{}
}

/*CompleteTransactionByExternalKeyUnprocessableEntity handles this case with default header values.

Payment is aborted by a control plugin
*/
type CompleteTransactionByExternalKeyUnprocessableEntity struct {
}

func (o *CompleteTransactionByExternalKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments][%d] completeTransactionByExternalKeyUnprocessableEntity ", 422)
}

func (o *CompleteTransactionByExternalKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionByExternalKeyBadGateway creates a CompleteTransactionByExternalKeyBadGateway with default headers values
func NewCompleteTransactionByExternalKeyBadGateway() *CompleteTransactionByExternalKeyBadGateway {
	return &CompleteTransactionByExternalKeyBadGateway{}
}

/*CompleteTransactionByExternalKeyBadGateway handles this case with default header values.

Failed to submit payment transaction
*/
type CompleteTransactionByExternalKeyBadGateway struct {
}

func (o *CompleteTransactionByExternalKeyBadGateway) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments][%d] completeTransactionByExternalKeyBadGateway ", 502)
}

func (o *CompleteTransactionByExternalKeyBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionByExternalKeyServiceUnavailable creates a CompleteTransactionByExternalKeyServiceUnavailable with default headers values
func NewCompleteTransactionByExternalKeyServiceUnavailable() *CompleteTransactionByExternalKeyServiceUnavailable {
	return &CompleteTransactionByExternalKeyServiceUnavailable{}
}

/*CompleteTransactionByExternalKeyServiceUnavailable handles this case with default header values.

Payment in unknown status, failed to receive gateway response
*/
type CompleteTransactionByExternalKeyServiceUnavailable struct {
}

func (o *CompleteTransactionByExternalKeyServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments][%d] completeTransactionByExternalKeyServiceUnavailable ", 503)
}

func (o *CompleteTransactionByExternalKeyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteTransactionByExternalKeyGatewayTimeout creates a CompleteTransactionByExternalKeyGatewayTimeout with default headers values
func NewCompleteTransactionByExternalKeyGatewayTimeout() *CompleteTransactionByExternalKeyGatewayTimeout {
	return &CompleteTransactionByExternalKeyGatewayTimeout{}
}

/*CompleteTransactionByExternalKeyGatewayTimeout handles this case with default header values.

Payment operation timeout
*/
type CompleteTransactionByExternalKeyGatewayTimeout struct {
}

func (o *CompleteTransactionByExternalKeyGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/payments][%d] completeTransactionByExternalKeyGatewayTimeout ", 504)
}

func (o *CompleteTransactionByExternalKeyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
