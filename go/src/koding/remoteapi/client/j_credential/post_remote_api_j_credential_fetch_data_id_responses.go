package j_credential

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJCredentialFetchDataIDReader is a Reader for the PostRemoteAPIJCredentialFetchDataID structure.
type PostRemoteAPIJCredentialFetchDataIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJCredentialFetchDataIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJCredentialFetchDataIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJCredentialFetchDataIDOK creates a PostRemoteAPIJCredentialFetchDataIDOK with default headers values
func NewPostRemoteAPIJCredentialFetchDataIDOK() *PostRemoteAPIJCredentialFetchDataIDOK {
	return &PostRemoteAPIJCredentialFetchDataIDOK{}
}

/*PostRemoteAPIJCredentialFetchDataIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJCredentialFetchDataIDOK struct {
	Payload *models.JCredential
}

func (o *PostRemoteAPIJCredentialFetchDataIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JCredential.fetchData/{id}][%d] postRemoteApiJCredentialFetchDataIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJCredentialFetchDataIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JCredential)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
