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

// GetEmailsReader is a Reader for the GetEmails structure.
type GetEmailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEmailsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetEmailsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetEmailsOK creates a GetEmailsOK with default headers values
func NewGetEmailsOK() *GetEmailsOK {
	return &GetEmailsOK{}
}

/*GetEmailsOK handles this case with default header values.

successful operation
*/
type GetEmailsOK struct {
	Payload []*kbmodel.AccountEmail
}

func (o *GetEmailsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/emails][%d] getEmailsOK  %+v", 200, o.Payload)
}

func (o *GetEmailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmailsBadRequest creates a GetEmailsBadRequest with default headers values
func NewGetEmailsBadRequest() *GetEmailsBadRequest {
	return &GetEmailsBadRequest{}
}

/*GetEmailsBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetEmailsBadRequest struct {
}

func (o *GetEmailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/emails][%d] getEmailsBadRequest ", 400)
}

func (o *GetEmailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
