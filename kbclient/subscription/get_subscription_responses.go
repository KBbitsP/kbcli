// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// GetSubscriptionReader is a Reader for the GetSubscription structure.
type GetSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionOK()
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

// NewGetSubscriptionOK creates a GetSubscriptionOK with default headers values
func NewGetSubscriptionOK() *GetSubscriptionOK {
	return &GetSubscriptionOK{}
}

/*GetSubscriptionOK handles this case with default header values.

successful operation
*/
type GetSubscriptionOK struct {
	Payload *kbmodel.Subscription

	HttpResponse runtime.ClientResponse
}

func (o *GetSubscriptionOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}][%d] getSubscriptionOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Subscription)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionBadRequest creates a GetSubscriptionBadRequest with default headers values
func NewGetSubscriptionBadRequest() *GetSubscriptionBadRequest {
	return &GetSubscriptionBadRequest{}
}

/*GetSubscriptionBadRequest handles this case with default header values.

Invalid subscription id supplied
*/
type GetSubscriptionBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetSubscriptionBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}][%d] getSubscriptionBadRequest ", 400)
}

func (o *GetSubscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSubscriptionNotFound creates a GetSubscriptionNotFound with default headers values
func NewGetSubscriptionNotFound() *GetSubscriptionNotFound {
	return &GetSubscriptionNotFound{}
}

/*GetSubscriptionNotFound handles this case with default header values.

Subscription not found
*/
type GetSubscriptionNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetSubscriptionNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/{subscriptionId}][%d] getSubscriptionNotFound ", 404)
}

func (o *GetSubscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
