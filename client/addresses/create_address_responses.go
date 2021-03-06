package addresses

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

type CreateAddressReader struct {
	formats strfmt.Registry
}

func (o *CreateAddressReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateAddressOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateAddressDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewCreateAddressOK creates a CreateAddressOK with default headers values
func NewCreateAddressOK() *CreateAddressOK {
	return &CreateAddressOK{}
}

/*CreateAddressOK

success
*/
type CreateAddressOK struct {
	Payload *models.AddressPagedMetadata
}

func (o *CreateAddressOK) Error() string {
	return fmt.Sprintf("[POST /addresses][%d] createAddressOK  %+v", 200, o.Payload)
}

func (o *CreateAddressOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressPagedMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAddressDefault creates a CreateAddressDefault with default headers values
func NewCreateAddressDefault(code int) *CreateAddressDefault {
	return &CreateAddressDefault{
		_statusCode: code,
	}
}

/*CreateAddressDefault

error
*/
type CreateAddressDefault struct {
	_statusCode int

	Payload *models.BFError
}

// Code gets the status code for the create address default response
func (o *CreateAddressDefault) Code() int {
	return o._statusCode
}

func (o *CreateAddressDefault) Error() string {
	return fmt.Sprintf("[POST /addresses][%d] createAddress default  %+v", o._statusCode, o.Payload)
}

func (o *CreateAddressDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BFError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
