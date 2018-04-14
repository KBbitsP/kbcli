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

// AddEmailReader is a Reader for the AddEmail structure.
type AddEmailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddEmailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAddEmailCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewAddEmailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewAddEmailNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddEmailCreated creates a AddEmailCreated with default headers values
func NewAddEmailCreated() *AddEmailCreated {
	return &AddEmailCreated{}
}

/*AddEmailCreated handles this case with default header values.

Email created successfully
*/
type AddEmailCreated struct {
	Payload []*kbmodel.AccountEmail
}

func (o *AddEmailCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/emails][%d] addEmailCreated  %+v", 201, o.Payload)
}

func (o *AddEmailCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddEmailBadRequest creates a AddEmailBadRequest with default headers values
func NewAddEmailBadRequest() *AddEmailBadRequest {
	return &AddEmailBadRequest{}
}

/*AddEmailBadRequest handles this case with default header values.

Invalid account id supplied
*/
type AddEmailBadRequest struct {
}

func (o *AddEmailBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/emails][%d] addEmailBadRequest ", 400)
}

func (o *AddEmailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddEmailNotFound creates a AddEmailNotFound with default headers values
func NewAddEmailNotFound() *AddEmailNotFound {
	return &AddEmailNotFound{}
}

/*AddEmailNotFound handles this case with default header values.

Account not found
*/
type AddEmailNotFound struct {
}

func (o *AddEmailNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/emails][%d] addEmailNotFound ", 404)
}

func (o *AddEmailNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
