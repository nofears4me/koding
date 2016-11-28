package j_custom_partials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJCustomPartialsUpdateIDReader is a Reader for the PostRemoteAPIJCustomPartialsUpdateID structure.
type PostRemoteAPIJCustomPartialsUpdateIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJCustomPartialsUpdateIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJCustomPartialsUpdateIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJCustomPartialsUpdateIDOK creates a PostRemoteAPIJCustomPartialsUpdateIDOK with default headers values
func NewPostRemoteAPIJCustomPartialsUpdateIDOK() *PostRemoteAPIJCustomPartialsUpdateIDOK {
	return &PostRemoteAPIJCustomPartialsUpdateIDOK{}
}

/*PostRemoteAPIJCustomPartialsUpdateIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJCustomPartialsUpdateIDOK struct {
	Payload *models.JCustomPartials
}

func (o *PostRemoteAPIJCustomPartialsUpdateIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JCustomPartials.update/{id}][%d] postRemoteApiJCustomPartialsUpdateIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJCustomPartialsUpdateIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JCustomPartials)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
