// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// CreateFutureInvoiceReader is a Reader for the CreateFutureInvoice structure.
type CreateFutureInvoiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateFutureInvoiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateFutureInvoiceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCreateFutureInvoiceBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateFutureInvoiceCreated creates a CreateFutureInvoiceCreated with default headers values
func NewCreateFutureInvoiceCreated() *CreateFutureInvoiceCreated {
	return &CreateFutureInvoiceCreated{}
}

/*CreateFutureInvoiceCreated handles this case with default header values.

Created invoice successfully
*/
type CreateFutureInvoiceCreated struct {
	Payload *kbmodel.Invoice
}

func (o *CreateFutureInvoiceCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices][%d] createFutureInvoiceCreated  %+v", 201, o.Payload)
}

func (o *CreateFutureInvoiceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Invoice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateFutureInvoiceBadRequest creates a CreateFutureInvoiceBadRequest with default headers values
func NewCreateFutureInvoiceBadRequest() *CreateFutureInvoiceBadRequest {
	return &CreateFutureInvoiceBadRequest{}
}

/*CreateFutureInvoiceBadRequest handles this case with default header values.

Invalid account id or target datetime supplied
*/
type CreateFutureInvoiceBadRequest struct {
}

func (o *CreateFutureInvoiceBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices][%d] createFutureInvoiceBadRequest ", 400)
}

func (o *CreateFutureInvoiceBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
