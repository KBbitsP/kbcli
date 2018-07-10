// Code generated by go-swagger; DO NOT EDIT.

package account

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

// GetAccountAuditLogsReader is a Reader for the GetAccountAuditLogs structure.
type GetAccountAuditLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountAuditLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountAuditLogsOK()
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

// NewGetAccountAuditLogsOK creates a GetAccountAuditLogsOK with default headers values
func NewGetAccountAuditLogsOK() *GetAccountAuditLogsOK {
	return &GetAccountAuditLogsOK{}
}

/*GetAccountAuditLogsOK handles this case with default header values.

successful operation
*/
type GetAccountAuditLogsOK struct {
	Payload []*kbmodel.AuditLog

	HttpResponse runtime.ClientResponse
}

func (o *GetAccountAuditLogsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/auditLogs][%d] getAccountAuditLogsOK  %+v", 200, o.Payload)
}

func (o *GetAccountAuditLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountAuditLogsNotFound creates a GetAccountAuditLogsNotFound with default headers values
func NewGetAccountAuditLogsNotFound() *GetAccountAuditLogsNotFound {
	return &GetAccountAuditLogsNotFound{}
}

/*GetAccountAuditLogsNotFound handles this case with default header values.

Account not found
*/
type GetAccountAuditLogsNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAccountAuditLogsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/auditLogs][%d] getAccountAuditLogsNotFound ", 404)
}

func (o *GetAccountAuditLogsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
