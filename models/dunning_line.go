package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*DunningLine A dunning-line specifies when the next execution attempt of an unpaid invoice should take place given the number of existing execution attempts.

swagger:model DunningLine
*/
type DunningLine struct {

	/* { "description" : "The payment attempt this dunning line applies to, specified as a positive integer. Dunning lines are ZERO indexed.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	AttemptIx int32 `json:"attemptIx,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy *string `json:"changedBy,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "Has the dunning-line been deleted?", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Deleted bool `json:"deleted,omitempty"`

	/* { "description" : "The payment gateway to use for this payment attempt.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Gateway string `json:"gateway,omitempty"`

	/* { "description" : "ID of the dunning-line.", "verbs":["GET"] }
	 */
	ID *string `json:"id,omitempty"`

	/* { "description" : "The time before the next payment attempt in minutes.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	MinutesDelay int32 `json:"minutesDelay,omitempty"`

	/* { "description" : "organization associated with the dunning-line.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Organization *Organization `json:"organization,omitempty"`

	/* { "description" : "ID of the organization associated with the dunning-line.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* { "description" : "ID of the organization associated with the dunning-line.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Type string `json:"type,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this dunning line
func (m *DunningLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttemptIx(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateMinutesDelay(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganization(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DunningLine) validateAttemptIx(formats strfmt.Registry) error {

	if err := validate.Required("attemptIx", "body", int32(m.AttemptIx)); err != nil {
		return err
	}

	return nil
}

func (m *DunningLine) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", bool(m.Deleted)); err != nil {
		return err
	}

	return nil
}

var dunningLineGatewayEnum []interface{}

func (m *DunningLine) validateGatewayEnum(path, location string, value string) error {
	if dunningLineGatewayEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["cybersource_token","card_vault","paypal_simple","locustworld","free","coupon","credit_note","stripe","braintree","balanced","paypal","billforward_test","offline","trial","stripeACH","authorizeNet","spreedly","sagePay","trustCommerce","payvision","kash"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			dunningLineGatewayEnum = append(dunningLineGatewayEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, dunningLineGatewayEnum); err != nil {
		return err
	}
	return nil
}

func (m *DunningLine) validateGateway(formats strfmt.Registry) error {

	if err := validate.RequiredString("gateway", "body", string(m.Gateway)); err != nil {
		return err
	}

	if err := m.validateGatewayEnum("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *DunningLine) validateMinutesDelay(formats strfmt.Registry) error {

	if err := validate.Required("minutesDelay", "body", int32(m.MinutesDelay)); err != nil {
		return err
	}

	return nil
}

func (m *DunningLine) validateOrganization(formats strfmt.Registry) error {

	if m.Organization != nil {

		if err := m.Organization.Validate(formats); err != nil {
			return err
		}
	}

	return nil
}

func (m *DunningLine) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.RequiredString("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}

var dunningLineTypeEnum []interface{}

func (m *DunningLine) validateTypeEnum(path, location string, value string) error {
	if dunningLineTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Invoice","Notification"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			dunningLineTypeEnum = append(dunningLineTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, dunningLineTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *DunningLine) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("type", "body", string(m.Type)); err != nil {
		return err
	}

	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}
