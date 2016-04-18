package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/authclub/billforward/models"
)

// DeleteMetadataForSubscriptionReader is a Reader for the DeleteMetadataForSubscription structure.
type DeleteMetadataForSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteMetadataForSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMetadataForSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteMetadataForSubscriptionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteMetadataForSubscriptionOK creates a DeleteMetadataForSubscriptionOK with default headers values
func NewDeleteMetadataForSubscriptionOK() *DeleteMetadataForSubscriptionOK {
	return &DeleteMetadataForSubscriptionOK{}
}

/*DeleteMetadataForSubscriptionOK handles this case with default header values.

success
*/
type DeleteMetadataForSubscriptionOK struct {
	Payload models.DynamicMetadata
}

func (o *DeleteMetadataForSubscriptionOK) Error() string {
	return fmt.Sprintf("[DELETE /subscriptions/{subscription-ID}/metadata][%d] deleteMetadataForSubscriptionOK  %+v", 200, o.Payload)
}

func (o *DeleteMetadataForSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMetadataForSubscriptionDefault creates a DeleteMetadataForSubscriptionDefault with default headers values
func NewDeleteMetadataForSubscriptionDefault(code int) *DeleteMetadataForSubscriptionDefault {
	return &DeleteMetadataForSubscriptionDefault{
		_statusCode: code,
	}
}

/*DeleteMetadataForSubscriptionDefault handles this case with default header values.

error
*/
type DeleteMetadataForSubscriptionDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the delete metadata for subscription default response
func (o *DeleteMetadataForSubscriptionDefault) Code() int {
	return o._statusCode
}

func (o *DeleteMetadataForSubscriptionDefault) Error() string {
	return fmt.Sprintf("[DELETE /subscriptions/{subscription-ID}/metadata][%d] deleteMetadataForSubscription default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteMetadataForSubscriptionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
