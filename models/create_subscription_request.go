package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

/*CreateSubscriptionRequest Entity for requesting that a subscription be created.

swagger:model CreateSubscriptionRequest
*/
type CreateSubscriptionRequest struct {

	/* {"description":"ID of the BillForward Account who will own this subscription. You should ensure beforehand that the customer has had a BillForward Account created for them.","verbs":["POST"]}

	Required: true
	*/
	AccountID string `json:"accountID,omitempty"`

	/* {"default":false,"description":"Whether this subscription should become an 'aggregating subscription', collecting charges (starting now) from all other subscriptions (current and future) belonging to this BillForward Account.","verbs":["POST"]}
	 */
	AggregateAllSubscriptionsOnAccount *bool `json:"aggregateAllSubscriptionsOnAccount,omitempty"`

	/* {"default":true,"description":"Whether to override the `end` date to line up with the current period end of the 'aggregating subscription' to which this subscription belongs.","verbs":["POST"]}
	 */
	AlignPeriodWithAggregatingSubscription *bool `json:"alignPeriodWithAggregatingSubscription,omitempty"`

	/* {"default":"(null)","description":"Description of the created subscription. This is primarily for your benefit &mdash; for example, you could write here the mechanism through which you obtained this customer. (e.g. 'Customer obtained through Lazy Wednesdays promotion').","verbs":["POST"]}
	 */
	Description *string `json:"description,omitempty"`

	/* {"default":"(1 period ahead of the `start` time)","description":"ISO 8601 UTC DateTime (e.g. 2015-06-16T11:58:41Z) describing the date at which the subscription should finish its first service period.","verbs":["POST"]}
	 */
	End *strfmt.DateTime `json:"end,omitempty"`

	/* { "description" : "Add metadata.", "verbs":["POST"] }
	 */
	Metadata DynamicMetadata `json:"metadata,omitempty"`

	/* {"default":"(Subscription will be named after the rate plan to which the subscription subscribes)","description":"Name of the created subscription. This is primarily for your benefit &mdash; for example, to enable you to identify subscriptions at a glance in the BillForward web interface (e.g. 'Customer 1425, guy@mail.com, Premium membership').","verbs":["POST"]}
	 */
	Name *string `json:"name,omitempty"`

	/* {"default":"(Auto-populated using your authentication credentials)","description":"ID of the BillForward Organization within which the requested Subscription should be created. If omitted, this will be auto-populated using your authentication credentials.","verbs":["POST"]}
	 */
	OrganizationID *string `json:"organizationID,omitempty"`

	/* {"default":"(If a subscription exists which 'aggregates all subscriptions belonging to this BillForward Account', refer to the ID of that subscription. Otherwise: null)","description":"ID of a parent subscription which will collect the charges raised by this subscription. The parent becomes responsible for paying those charges. If a subscription exists which 'aggregates all subscriptions belonging to this BillForward Account', then that parent will override any parent specified here.","verbs":["POST"]}
	 */
	ParentID *string `json:"parentID,omitempty"`

	/* {"default":"(empty list)","description":"Quantities that this subscription possesses (upon beginning service), of pricing components upon the subscription's rate plan. For example: you can set the subscription to begin its service with '5 widgets' consumed. Otherwise the 'default quantity' will be observed instead, for each pricing component upon the rate plan.","verbs":["POST"]}
	 */
	PricingComponentQuantities []*PricingComponentQuantityRequest `json:"pricingComponentQuantities,omitempty"`

	/* {"description":"Name or ID of the product.","verbs":["POST"]}

	Required: true
	*/
	Product string `json:"product,omitempty"`

	/* {"description":"ID or name of the rate plan to which the subscription will be subscribing. Lookup by name is only possible if a `productID` is specified.","verbs":["POST"]}

	Required: true
	*/
	ProductRatePlan string `json:"productRatePlan,omitempty"`

	/* {"default":"(ServerNow upon receiving request)","description":"ISO 8601 UTC DateTime (e.g. 2015-06-16T11:58:41Z) describing the date at which the subscription should enter its first service period.","verbs":["POST"]}
	 */
	Start *strfmt.DateTime `json:"start,omitempty"`

	/* {"default":"Provisioned","description":"The state in which the created subscription will begin.<br><span class=\"label label-default\">Provisioned</span> &mdash; The subscription will wait (without raising any invoices or beginning its service) until explicit action is taken to change its state.<br><span class=\"label label-default\">AwaitingPayment</span> &mdash; The subscription is activated. After `start` time is surpassed, it will begin service and raise its first invoice.","verbs":["POST"]}
	 */
	State *string `json:"state,omitempty"`

	/* {"default":"(null)","description":"ISO 8601 UTC DateTime (e.g. 2015-06-16T11:58:41Z) describing the date at which the subscription should leave the trial period.","verbs":["POST"]}
	 */
	TrialEnd *strfmt.DateTime `json:"trialEnd,omitempty"`

	/* Type type
	 */
	Type *string `json:"type,omitempty"`
}

// Validate validates this create subscription request
func (m *CreateSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePricingComponentQuantities(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProductRatePlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *CreateSubscriptionRequest) validateAccountID(formats strfmt.Registry) error {

	if err := validate.RequiredString("accountID", "body", string(m.AccountID)); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubscriptionRequest) validatePricingComponentQuantities(formats strfmt.Registry) error {

	if swag.IsZero(m.PricingComponentQuantities) { // not required
		return nil
	}

	for i := 0; i < len(m.PricingComponentQuantities); i++ {

		if m.PricingComponentQuantities[i] != nil {

			if err := m.PricingComponentQuantities[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *CreateSubscriptionRequest) validateProduct(formats strfmt.Registry) error {

	if err := validate.RequiredString("product", "body", string(m.Product)); err != nil {
		return err
	}

	return nil
}

func (m *CreateSubscriptionRequest) validateProductRatePlan(formats strfmt.Registry) error {

	if err := validate.RequiredString("productRatePlan", "body", string(m.ProductRatePlan)); err != nil {
		return err
	}

	return nil
}

var createSubscriptionRequestStateEnum []interface{}

func (m *CreateSubscriptionRequest) validateStateEnum(path, location string, value string) error {
	if createSubscriptionRequestStateEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Trial","Provisioned","Paid","AwaitingPayment","Cancelled","Failed","Expired"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createSubscriptionRequestStateEnum = append(createSubscriptionRequestStateEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, createSubscriptionRequestStateEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateSubscriptionRequest) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

var createSubscriptionRequestTypeEnum []interface{}

func (m *CreateSubscriptionRequest) validateTypeEnum(path, location string, value string) error {
	if createSubscriptionRequestTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Subscription","FixedTerm","Trial"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createSubscriptionRequestTypeEnum = append(createSubscriptionRequestTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, createSubscriptionRequestTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateSubscriptionRequest) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}
