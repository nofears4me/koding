package j_team_invitation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJTeamInvitationByCodeReader is a Reader for the PostRemoteAPIJTeamInvitationByCode structure.
type PostRemoteAPIJTeamInvitationByCodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJTeamInvitationByCodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJTeamInvitationByCodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJTeamInvitationByCodeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJTeamInvitationByCodeOK creates a PostRemoteAPIJTeamInvitationByCodeOK with default headers values
func NewPostRemoteAPIJTeamInvitationByCodeOK() *PostRemoteAPIJTeamInvitationByCodeOK {
	return &PostRemoteAPIJTeamInvitationByCodeOK{}
}

/*PostRemoteAPIJTeamInvitationByCodeOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJTeamInvitationByCodeOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJTeamInvitationByCodeOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTeamInvitation.byCode][%d] postRemoteApiJTeamInvitationByCodeOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJTeamInvitationByCodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJTeamInvitationByCodeUnauthorized creates a PostRemoteAPIJTeamInvitationByCodeUnauthorized with default headers values
func NewPostRemoteAPIJTeamInvitationByCodeUnauthorized() *PostRemoteAPIJTeamInvitationByCodeUnauthorized {
	return &PostRemoteAPIJTeamInvitationByCodeUnauthorized{}
}

/*PostRemoteAPIJTeamInvitationByCodeUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJTeamInvitationByCodeUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJTeamInvitationByCodeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JTeamInvitation.byCode][%d] postRemoteApiJTeamInvitationByCodeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJTeamInvitationByCodeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
