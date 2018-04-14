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

// GetBundleCustomFieldsReader is a Reader for the GetBundleCustomFields structure.
type GetBundleCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetBundleCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetBundleCustomFieldsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetBundleCustomFieldsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetBundleCustomFieldsOK creates a GetBundleCustomFieldsOK with default headers values
func NewGetBundleCustomFieldsOK() *GetBundleCustomFieldsOK {
	return &GetBundleCustomFieldsOK{}
}

/*GetBundleCustomFieldsOK handles this case with default header values.

successful operation
*/
type GetBundleCustomFieldsOK struct {
	Payload []*kbmodel.CustomField
}

func (o *GetBundleCustomFieldsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/{bundleId}/customFields][%d] getBundleCustomFieldsOK  %+v", 200, o.Payload)
}

func (o *GetBundleCustomFieldsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetBundleCustomFieldsBadRequest creates a GetBundleCustomFieldsBadRequest with default headers values
func NewGetBundleCustomFieldsBadRequest() *GetBundleCustomFieldsBadRequest {
	return &GetBundleCustomFieldsBadRequest{}
}

/*GetBundleCustomFieldsBadRequest handles this case with default header values.

Invalid bundle id supplied
*/
type GetBundleCustomFieldsBadRequest struct {
}

func (o *GetBundleCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/bundles/{bundleId}/customFields][%d] getBundleCustomFieldsBadRequest ", 400)
}

func (o *GetBundleCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
