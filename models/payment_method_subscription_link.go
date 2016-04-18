package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*PaymentMethodSubscriptionLink PaymentMethodSubscriptionLink

swagger:model PaymentMethodSubscriptionLink
*/
type PaymentMethodSubscriptionLink struct {

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "CRM ID of the product-rate-plan.", "verbs":["POST","PUT","GET"] }
	 */
	CrmID string `json:"crmID,omitempty"`

	/* deleted
	 */
	Deleted *bool `json:"deleted,omitempty"`

	/* { "description" : "", "verbs":["GET"] }
	 */
	ID string `json:"id,omitempty"`

	/* organization ID
	 */
	OrganizationID string `json:"organizationID,omitempty"`

	/* payment method
	 */
	PaymentMethod *PaymentMethod `json:"paymentMethod,omitempty"`

	/* { "description" : "Payment method to add to subscription.", "verbs":["POST","GET"] }

	Required: true
	*/
	PaymentMethodID *string `json:"paymentMethodID"`

	/* { "description" : "Subscription to add payment method to.", "verbs":["POST","GET"] }

	Required: true
	*/
	SubscriptionID *string `json:"subscriptionID"`

	/* { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this payment method subscription link
func (m *PaymentMethodSubscriptionLink) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePaymentMethodID(formats); err != nil {
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

func (m *PaymentMethodSubscriptionLink) validatePaymentMethodID(formats strfmt.Registry) error {

	if err := validate.Required("paymentMethodID", "body", m.PaymentMethodID); err != nil {
		return err
	}

	return nil
}

func (m *PaymentMethodSubscriptionLink) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionID", "body", m.SubscriptionID); err != nil {
		return err
	}

	return nil
}
