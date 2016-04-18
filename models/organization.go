package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/validate"
)

/*Organization Organization within which all your BillForward interactions are scoped.

swagger:model Organization
*/
type Organization struct {

	/* alias
	 */
	Alias []*Alias `json:"alias,omitempty"`

	/* { "description" : "Stores the organizations 3rd party API keys which may be used by the system for payment gateway integration etc.", "verbs":["POST","PUT","GET"] }
	 */
	APIConfigurations []*MutableBillingEntity `json:"apiConfigurations,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* { "description" : "The OAuth2 clients associated with the organization. In most cases an organization would not have any clients associated with their account. In the case of an APP developer, a clients would exist per an application they have developed. To further understand clients please see the client, OAuth2 API and APP development documentation.", "verbs":["POST","PUT","GET"] }
	 */
	Clients []*Client `json:"clients,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "A shortcode for the organization. This is used as a short reference code for use when referring to the organization, by default this is set to the organizations name.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	CustomerCode *string `json:"customerCode"`

	/* { "description" : "Indicates if an organization has been retired. If an organization has been retired it can still be retrieved using the appropriate flag on API requests.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Deleted bool `json:"deleted"`

	/* { "description" : "The dunning-lines associated with the organization. Dunning lines are used as re-try logic for invoices to attempt to reconcile the payment.", "verbs":["POST","PUT","GET"] }
	 */
	DunningLines []*DunningLine `json:"dunningLines,omitempty"`

	/* { "description" : "ID of the organization.", "verbs":["POST","PUT","GET"] }
	 */
	ID string `json:"id,omitempty"`

	/* { "description" : "The name of the organization.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Name *string `json:"name"`

	/* { "description" : "The card-vault gateways associated with the organization.", "verbs":["POST","PUT","GET"] }
	 */
	OrganizationGateways []*MutableBillingEntity `json:"organizationGateways,omitempty"`

	/* { "description" : "The taxation-strategies associated with the organization. Taxation-strategies may be linked to product-rate-plans to specify how their tax should be calculated.", "verbs":["POST","PUT","GET"] }
	 */
	TaxationStrategies []*MutableBillingEntity `json:"taxationStrategies,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated.", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`

	/* { "description" : "The WebHooks associated with the organization. These are the end-points where notifications are sent. WebHooks are added, updated and removed from the organization. Thus to add a WebHook, the webhook must be defined on this property of the organization and then the organization updated. To update a WebHook the same procedure must be followed, first retrieving the organization followed by updating the appropriate WebHook, finally the organization is updated.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Webhooks []*MutableBillingEntity `json:"webhooks"`
}

// Validate validates this organization
func (m *Organization) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlias(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAPIConfigurations(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClients(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCustomerCode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDunningLines(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationGateways(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTaxationStrategies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWebhooks(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Organization) validateAlias(formats strfmt.Registry) error {

	if swag.IsZero(m.Alias) { // not required
		return nil
	}

	return nil
}

func (m *Organization) validateAPIConfigurations(formats strfmt.Registry) error {

	if swag.IsZero(m.APIConfigurations) { // not required
		return nil
	}

	return nil
}

func (m *Organization) validateClients(formats strfmt.Registry) error {

	if swag.IsZero(m.Clients) { // not required
		return nil
	}

	return nil
}

func (m *Organization) validateCustomerCode(formats strfmt.Registry) error {

	if err := validate.Required("customerCode", "body", m.CustomerCode); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", bool(m.Deleted)); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateDunningLines(formats strfmt.Registry) error {

	if swag.IsZero(m.DunningLines) { // not required
		return nil
	}

	return nil
}

func (m *Organization) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Organization) validateOrganizationGateways(formats strfmt.Registry) error {

	if swag.IsZero(m.OrganizationGateways) { // not required
		return nil
	}

	return nil
}

func (m *Organization) validateTaxationStrategies(formats strfmt.Registry) error {

	if swag.IsZero(m.TaxationStrategies) { // not required
		return nil
	}

	return nil
}

func (m *Organization) validateWebhooks(formats strfmt.Registry) error {

	if err := validate.Required("webhooks", "body", m.Webhooks); err != nil {
		return err
	}

	return nil
}
