// Code generated by go-swagger; DO NOT EDIT.

package plugin_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/kbmodel"
)

// GetPluginsInfoReader is a Reader for the GetPluginsInfo structure.
type GetPluginsInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPluginsInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPluginsInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPluginsInfoOK creates a GetPluginsInfoOK with default headers values
func NewGetPluginsInfoOK() *GetPluginsInfoOK {
	return &GetPluginsInfoOK{}
}

/*GetPluginsInfoOK handles this case with default header values.

successful operation
*/
type GetPluginsInfoOK struct {
	Payload []*kbmodel.PluginInfo
}

func (o *GetPluginsInfoOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/pluginsInfo][%d] getPluginsInfoOK  %+v", 200, o.Payload)
}

func (o *GetPluginsInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
