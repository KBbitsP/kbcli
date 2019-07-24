// Code generated by go-swagger; DO NOT EDIT.

package debug

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

// GetClockReader is a Reader for the GetClock structure.
type GetClockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClockOK()
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

// NewGetClockOK creates a GetClockOK with default headers values
func NewGetClockOK() *GetClockOK {
	return &GetClockOK{}
}

/*GetClockOK handles this case with default header values.

successful operation
*/
type GetClockOK struct {
	Payload *kbmodel.Clock

	HttpResponse runtime.ClientResponse
}

func (o *GetClockOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/test/clock][%d] getClockOK  %+v", 200, o.Payload)
}

func (o *GetClockOK) GetPayload() *kbmodel.Clock {
	return o.Payload
}

func (o *GetClockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Clock)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
