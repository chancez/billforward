package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*CancelSubscriptionRequest CancelSubscriptionRequest

swagger:model CancelSubscriptionRequest
*/
type CancelSubscriptionRequest struct {

	/* cancel children
	 */
	CancelChildren bool `json:"cancelChildren,omitempty"`

	/* cancel empty parent
	 */
	CancelEmptyParent bool `json:"cancelEmptyParent,omitempty"`

	/* Specifies whether the service will end immediately on cancellation or if it will continue until the end of the current period. Default: AtPeriodEnd
	 */
	CancellationCredit string `json:"cancellationCredit,omitempty"`

	/* service end
	 */
	ServiceEnd string `json:"serviceEnd,omitempty"`

	/* source

	Required: true
	*/
	Source *string `json:"source"`

	/* state
	 */
	State string `json:"state,omitempty"`

	/* subscription ID

	Required: true
	*/
	SubscriptionID *string `json:"subscriptionID"`
}

// Validate validates this cancel subscription request
func (m *CancelSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCancellationCredit(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateServiceEnd(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubscriptionID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cancelSubscriptionRequestTypeCancellationCreditPropEnum []interface{}

// prop value enum
func (m *CancelSubscriptionRequest) validateCancellationCreditEnum(path, location string, value string) error {
	if cancelSubscriptionRequestTypeCancellationCreditPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Credit","None"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			cancelSubscriptionRequestTypeCancellationCreditPropEnum = append(cancelSubscriptionRequestTypeCancellationCreditPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, cancelSubscriptionRequestTypeCancellationCreditPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CancelSubscriptionRequest) validateCancellationCredit(formats strfmt.Registry) error {

	if swag.IsZero(m.CancellationCredit) { // not required
		return nil
	}

	// value enum
	if err := m.validateCancellationCreditEnum("cancellationCredit", "body", m.CancellationCredit); err != nil {
		return err
	}

	return nil
}

var cancelSubscriptionRequestTypeServiceEndPropEnum []interface{}

// prop value enum
func (m *CancelSubscriptionRequest) validateServiceEndEnum(path, location string, value string) error {
	if cancelSubscriptionRequestTypeServiceEndPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Immediate","AtPeriodEnd"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			cancelSubscriptionRequestTypeServiceEndPropEnum = append(cancelSubscriptionRequestTypeServiceEndPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, cancelSubscriptionRequestTypeServiceEndPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CancelSubscriptionRequest) validateServiceEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceEnd) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceEndEnum("serviceEnd", "body", m.ServiceEnd); err != nil {
		return err
	}

	return nil
}

func (m *CancelSubscriptionRequest) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *CancelSubscriptionRequest) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionID", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}
