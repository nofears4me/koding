package j_workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJWorkspaceDeleteByIDReader is a Reader for the PostRemoteAPIJWorkspaceDeleteByID structure.
type PostRemoteAPIJWorkspaceDeleteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJWorkspaceDeleteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJWorkspaceDeleteByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJWorkspaceDeleteByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJWorkspaceDeleteByIDOK creates a PostRemoteAPIJWorkspaceDeleteByIDOK with default headers values
func NewPostRemoteAPIJWorkspaceDeleteByIDOK() *PostRemoteAPIJWorkspaceDeleteByIDOK {
	return &PostRemoteAPIJWorkspaceDeleteByIDOK{}
}

/*PostRemoteAPIJWorkspaceDeleteByIDOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJWorkspaceDeleteByIDOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJWorkspaceDeleteByIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JWorkspace.deleteById][%d] postRemoteApiJWorkspaceDeleteByIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJWorkspaceDeleteByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJWorkspaceDeleteByIDUnauthorized creates a PostRemoteAPIJWorkspaceDeleteByIDUnauthorized with default headers values
func NewPostRemoteAPIJWorkspaceDeleteByIDUnauthorized() *PostRemoteAPIJWorkspaceDeleteByIDUnauthorized {
	return &PostRemoteAPIJWorkspaceDeleteByIDUnauthorized{}
}

/*PostRemoteAPIJWorkspaceDeleteByIDUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJWorkspaceDeleteByIDUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJWorkspaceDeleteByIDUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JWorkspace.deleteById][%d] postRemoteApiJWorkspaceDeleteByIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJWorkspaceDeleteByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
