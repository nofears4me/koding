package j_reward_campaign

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJRewardCampaignCreateReader is a Reader for the PostRemoteAPIJRewardCampaignCreate structure.
type PostRemoteAPIJRewardCampaignCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJRewardCampaignCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJRewardCampaignCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostRemoteAPIJRewardCampaignCreateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJRewardCampaignCreateOK creates a PostRemoteAPIJRewardCampaignCreateOK with default headers values
func NewPostRemoteAPIJRewardCampaignCreateOK() *PostRemoteAPIJRewardCampaignCreateOK {
	return &PostRemoteAPIJRewardCampaignCreateOK{}
}

/*PostRemoteAPIJRewardCampaignCreateOK handles this case with default header values.

Request processed succesfully
*/
type PostRemoteAPIJRewardCampaignCreateOK struct {
	Payload *models.DefaultResponse
}

func (o *PostRemoteAPIJRewardCampaignCreateOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JRewardCampaign.create][%d] postRemoteApiJRewardCampaignCreateOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJRewardCampaignCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DefaultResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostRemoteAPIJRewardCampaignCreateUnauthorized creates a PostRemoteAPIJRewardCampaignCreateUnauthorized with default headers values
func NewPostRemoteAPIJRewardCampaignCreateUnauthorized() *PostRemoteAPIJRewardCampaignCreateUnauthorized {
	return &PostRemoteAPIJRewardCampaignCreateUnauthorized{}
}

/*PostRemoteAPIJRewardCampaignCreateUnauthorized handles this case with default header values.

Unauthorized request
*/
type PostRemoteAPIJRewardCampaignCreateUnauthorized struct {
	Payload *models.UnauthorizedRequest
}

func (o *PostRemoteAPIJRewardCampaignCreateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /remote.api/JRewardCampaign.create][%d] postRemoteApiJRewardCampaignCreateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostRemoteAPIJRewardCampaignCreateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UnauthorizedRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
