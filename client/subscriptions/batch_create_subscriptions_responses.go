package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/authclub/billforward/models"
)

// BatchCreateSubscriptionsReader is a Reader for the BatchCreateSubscriptions structure.
type BatchCreateSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *BatchCreateSubscriptionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBatchCreateSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewBatchCreateSubscriptionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewBatchCreateSubscriptionsOK creates a BatchCreateSubscriptionsOK with default headers values
func NewBatchCreateSubscriptionsOK() *BatchCreateSubscriptionsOK {
	return &BatchCreateSubscriptionsOK{}
}

/*BatchCreateSubscriptionsOK handles this case with default header values.

success
*/
type BatchCreateSubscriptionsOK struct {
	Payload *models.SubscriptionPagedMetadata
}

func (o *BatchCreateSubscriptionsOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions/batch][%d] batchCreateSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *BatchCreateSubscriptionsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubscriptionPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBatchCreateSubscriptionsDefault creates a BatchCreateSubscriptionsDefault with default headers values
func NewBatchCreateSubscriptionsDefault(code int) *BatchCreateSubscriptionsDefault {
	return &BatchCreateSubscriptionsDefault{
		_statusCode: code,
	}
}

/*BatchCreateSubscriptionsDefault handles this case with default header values.

error
*/
type BatchCreateSubscriptionsDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the batch create subscriptions default response
func (o *BatchCreateSubscriptionsDefault) Code() int {
	return o._statusCode
}

func (o *BatchCreateSubscriptionsDefault) Error() string {
	return fmt.Sprintf("[POST /subscriptions/batch][%d] batchCreateSubscriptions default  %+v", o._statusCode, o.Payload)
}

func (o *BatchCreateSubscriptionsDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
