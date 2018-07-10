// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

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

// CreateTransactionTagsReader is a Reader for the CreateTransactionTags structure.
type CreateTransactionTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTransactionTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateTransactionTagsCreated()
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

// NewCreateTransactionTagsCreated creates a CreateTransactionTagsCreated with default headers values
func NewCreateTransactionTagsCreated() *CreateTransactionTagsCreated {
	return &CreateTransactionTagsCreated{}
}

/*CreateTransactionTagsCreated handles this case with default header values.

Tag created successfully
*/
type CreateTransactionTagsCreated struct {
	Payload []*kbmodel.Tag

	HttpResponse runtime.ClientResponse
}

func (o *CreateTransactionTagsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentTransactions/{transactionId}/tags][%d] createTransactionTagsCreated  %+v", 201, o.Payload)
}

func (o *CreateTransactionTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionTagsBadRequest creates a CreateTransactionTagsBadRequest with default headers values
func NewCreateTransactionTagsBadRequest() *CreateTransactionTagsBadRequest {
	return &CreateTransactionTagsBadRequest{}
}

/*CreateTransactionTagsBadRequest handles this case with default header values.

Invalid transaction id supplied
*/
type CreateTransactionTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *CreateTransactionTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentTransactions/{transactionId}/tags][%d] createTransactionTagsBadRequest ", 400)
}

func (o *CreateTransactionTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
