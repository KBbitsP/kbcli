// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// GetBundleByKeyReader is a Reader for the GetBundleByKey structure.
type GetBundleByKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBundleByKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBundleByKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetBundleByKeyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBundleByKeyOK creates a GetBundleByKeyOK with default headers values
func NewGetBundleByKeyOK() *GetBundleByKeyOK {
	return &GetBundleByKeyOK{}
}

/*GetBundleByKeyOK handles this case with default header values.

successful operation
*/
type GetBundleByKeyOK struct {
	Payload []*kbmodel.Bundle
}

func (o *GetBundleByKeyOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles][%d] getBundleByKeyOK  %+v", 200, o.Payload)
}

func (o *GetBundleByKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBundleByKeyNotFound creates a GetBundleByKeyNotFound with default headers values
func NewGetBundleByKeyNotFound() *GetBundleByKeyNotFound {
	return &GetBundleByKeyNotFound{}
}

/*GetBundleByKeyNotFound handles this case with default header values.

Bundle not found
*/
type GetBundleByKeyNotFound struct {
}

func (o *GetBundleByKeyNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles][%d] getBundleByKeyNotFound ", 404)
}

func (o *GetBundleByKeyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
