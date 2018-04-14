// Code generated by go-swagger; DO NOT EDIT.

package catalog

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// UploadCatalogXMLReader is a Reader for the UploadCatalogXML structure.
type UploadCatalogXMLReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadCatalogXMLReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewUploadCatalogXMLCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadCatalogXMLCreated creates a UploadCatalogXMLCreated with default headers values
func NewUploadCatalogXMLCreated() *UploadCatalogXMLCreated {
	return &UploadCatalogXMLCreated{}
}

/*UploadCatalogXMLCreated handles this case with default header values.

Catalog XML created successfully
*/
type UploadCatalogXMLCreated struct {
	Payload string
}

func (o *UploadCatalogXMLCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/catalog/xml][%d] uploadCatalogXmlCreated  %+v", 201, o.Payload)
}

func (o *UploadCatalogXMLCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
