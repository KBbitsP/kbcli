// Code generated by go-swagger; DO NOT EDIT.

package bundle

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// ResumeBundleReader is a Reader for the ResumeBundle structure.
type ResumeBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResumeBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewResumeBundleNoContent()
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

// NewResumeBundleNoContent creates a ResumeBundleNoContent with default headers values
func NewResumeBundleNoContent() *ResumeBundleNoContent {
	return &ResumeBundleNoContent{}
}

/*ResumeBundleNoContent handles this case with default header values.

Successful operation
*/
type ResumeBundleNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *ResumeBundleNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/resume][%d] resumeBundleNoContent ", 204)
}

func (o *ResumeBundleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResumeBundleBadRequest creates a ResumeBundleBadRequest with default headers values
func NewResumeBundleBadRequest() *ResumeBundleBadRequest {
	return &ResumeBundleBadRequest{}
}

/*ResumeBundleBadRequest handles this case with default header values.

Invalid bundle id supplied
*/
type ResumeBundleBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *ResumeBundleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/resume][%d] resumeBundleBadRequest ", 400)
}

func (o *ResumeBundleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewResumeBundleNotFound creates a ResumeBundleNotFound with default headers values
func NewResumeBundleNotFound() *ResumeBundleNotFound {
	return &ResumeBundleNotFound{}
}

/*ResumeBundleNotFound handles this case with default header values.

Bundle not found
*/
type ResumeBundleNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *ResumeBundleNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/bundles/{bundleId}/resume][%d] resumeBundleNotFound ", 404)
}

func (o *ResumeBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
