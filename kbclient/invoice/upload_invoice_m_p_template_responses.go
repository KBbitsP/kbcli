// Code generated by go-swagger; DO NOT EDIT.

package invoice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UploadInvoiceMPTemplateReader is a Reader for the UploadInvoiceMPTemplate structure.
type UploadInvoiceMPTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadInvoiceMPTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUploadInvoiceMPTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadInvoiceMPTemplateOK creates a UploadInvoiceMPTemplateOK with default headers values
func NewUploadInvoiceMPTemplateOK() *UploadInvoiceMPTemplateOK {
	return &UploadInvoiceMPTemplateOK{}
}

/*UploadInvoiceMPTemplateOK handles this case with default header values.

successful operation
*/
type UploadInvoiceMPTemplateOK struct {
	Payload string
}

func (o *UploadInvoiceMPTemplateOK) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoices/manualPayTemplate][%d] uploadInvoiceMPTemplateOK  %+v", 200, o.Payload)
}

func (o *UploadInvoiceMPTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
