// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteCatalogReader is a Reader for the DeleteCatalog structure.
type DeleteCatalogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCatalogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteCatalogNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCatalogNoContent creates a DeleteCatalogNoContent with default headers values
func NewDeleteCatalogNoContent() *DeleteCatalogNoContent {
	return &DeleteCatalogNoContent{}
}

/*DeleteCatalogNoContent handles this case with default header values.

Successful operation
*/
type DeleteCatalogNoContent struct {
}

func (o *DeleteCatalogNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/catalog][%d] deleteCatalogNoContent ", 204)
}

func (o *DeleteCatalogNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
