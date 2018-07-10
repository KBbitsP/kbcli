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

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// CaptureAuthorizationByExternalKeyReader is a Reader for the CaptureAuthorizationByExternalKey structure.
type CaptureAuthorizationByExternalKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CaptureAuthorizationByExternalKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCaptureAuthorizationByExternalKeyCreated()
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

// NewCaptureAuthorizationByExternalKeyCreated creates a CaptureAuthorizationByExternalKeyCreated with default headers values
func NewCaptureAuthorizationByExternalKeyCreated() *CaptureAuthorizationByExternalKeyCreated {
	return &CaptureAuthorizationByExternalKeyCreated{}
}

/*CaptureAuthorizationByExternalKeyCreated handles this case with default header values.

Payment transaction created successfully
*/
type CaptureAuthorizationByExternalKeyCreated struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationByExternalKeyCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments][%d] captureAuthorizationByExternalKeyCreated  %+v", 201, o.Payload)
}

func (o *CaptureAuthorizationByExternalKeyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCaptureAuthorizationByExternalKeyPaymentRequired creates a CaptureAuthorizationByExternalKeyPaymentRequired with default headers values
func NewCaptureAuthorizationByExternalKeyPaymentRequired() *CaptureAuthorizationByExternalKeyPaymentRequired {
	return &CaptureAuthorizationByExternalKeyPaymentRequired{}
}

/*CaptureAuthorizationByExternalKeyPaymentRequired handles this case with default header values.

Transaction declined by gateway
*/
type CaptureAuthorizationByExternalKeyPaymentRequired struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationByExternalKeyPaymentRequired) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments][%d] captureAuthorizationByExternalKeyPaymentRequired ", 402)
}

func (o *CaptureAuthorizationByExternalKeyPaymentRequired) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationByExternalKeyNotFound creates a CaptureAuthorizationByExternalKeyNotFound with default headers values
func NewCaptureAuthorizationByExternalKeyNotFound() *CaptureAuthorizationByExternalKeyNotFound {
	return &CaptureAuthorizationByExternalKeyNotFound{}
}

/*CaptureAuthorizationByExternalKeyNotFound handles this case with default header values.

Account or payment not found
*/
type CaptureAuthorizationByExternalKeyNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationByExternalKeyNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments][%d] captureAuthorizationByExternalKeyNotFound ", 404)
}

func (o *CaptureAuthorizationByExternalKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationByExternalKeyUnprocessableEntity creates a CaptureAuthorizationByExternalKeyUnprocessableEntity with default headers values
func NewCaptureAuthorizationByExternalKeyUnprocessableEntity() *CaptureAuthorizationByExternalKeyUnprocessableEntity {
	return &CaptureAuthorizationByExternalKeyUnprocessableEntity{}
}

/*CaptureAuthorizationByExternalKeyUnprocessableEntity handles this case with default header values.

Payment is aborted by a control plugin
*/
type CaptureAuthorizationByExternalKeyUnprocessableEntity struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationByExternalKeyUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments][%d] captureAuthorizationByExternalKeyUnprocessableEntity ", 422)
}

func (o *CaptureAuthorizationByExternalKeyUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationByExternalKeyBadGateway creates a CaptureAuthorizationByExternalKeyBadGateway with default headers values
func NewCaptureAuthorizationByExternalKeyBadGateway() *CaptureAuthorizationByExternalKeyBadGateway {
	return &CaptureAuthorizationByExternalKeyBadGateway{}
}

/*CaptureAuthorizationByExternalKeyBadGateway handles this case with default header values.

Failed to submit payment transaction
*/
type CaptureAuthorizationByExternalKeyBadGateway struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationByExternalKeyBadGateway) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments][%d] captureAuthorizationByExternalKeyBadGateway ", 502)
}

func (o *CaptureAuthorizationByExternalKeyBadGateway) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationByExternalKeyServiceUnavailable creates a CaptureAuthorizationByExternalKeyServiceUnavailable with default headers values
func NewCaptureAuthorizationByExternalKeyServiceUnavailable() *CaptureAuthorizationByExternalKeyServiceUnavailable {
	return &CaptureAuthorizationByExternalKeyServiceUnavailable{}
}

/*CaptureAuthorizationByExternalKeyServiceUnavailable handles this case with default header values.

Payment in unknown status, failed to receive gateway response
*/
type CaptureAuthorizationByExternalKeyServiceUnavailable struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationByExternalKeyServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments][%d] captureAuthorizationByExternalKeyServiceUnavailable ", 503)
}

func (o *CaptureAuthorizationByExternalKeyServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCaptureAuthorizationByExternalKeyGatewayTimeout creates a CaptureAuthorizationByExternalKeyGatewayTimeout with default headers values
func NewCaptureAuthorizationByExternalKeyGatewayTimeout() *CaptureAuthorizationByExternalKeyGatewayTimeout {
	return &CaptureAuthorizationByExternalKeyGatewayTimeout{}
}

/*CaptureAuthorizationByExternalKeyGatewayTimeout handles this case with default header values.

Payment operation timeout
*/
type CaptureAuthorizationByExternalKeyGatewayTimeout struct {
	HttpResponse runtime.ClientResponse
}

func (o *CaptureAuthorizationByExternalKeyGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/payments][%d] captureAuthorizationByExternalKeyGatewayTimeout ", 504)
}

func (o *CaptureAuthorizationByExternalKeyGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
