// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// InvalidatesCacheByTenantReader is a Reader for the InvalidatesCacheByTenant structure.
type InvalidatesCacheByTenantReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvalidatesCacheByTenantReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewInvalidatesCacheByTenantNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewInvalidatesCacheByTenantNoContent creates a InvalidatesCacheByTenantNoContent with default headers values
func NewInvalidatesCacheByTenantNoContent() *InvalidatesCacheByTenantNoContent {
	return &InvalidatesCacheByTenantNoContent{}
}

/*InvalidatesCacheByTenantNoContent handles this case with default header values.

Successful operation
*/
type InvalidatesCacheByTenantNoContent struct {
}

func (o *InvalidatesCacheByTenantNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/admin/cache/tenants][%d] invalidatesCacheByTenantNoContent ", 204)
}

func (o *InvalidatesCacheByTenantNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
