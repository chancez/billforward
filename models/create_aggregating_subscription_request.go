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

/*CreateAggregatingSubscriptionRequest Entity for requesting that an 'aggregating subscription' (i.e. a 'parent subscription' which collects the charges raised by many 'child subscriptions') be created.

swagger:model CreateAggregatingSubscriptionRequest
*/
type CreateAggregatingSubscriptionRequest struct {

	/* {"description":"ID of the BillForward Account who will own this aggregating subscription. You should ensure beforehand that the customer has had a BillForward Account created for them.","verbs":["POST"]}

	Required: true
	*/
	AccountID string `json:"accountID,omitempty"`

	/* {"default":false,"description":"Whether this 'aggregating subscription' should collect charges (starting now) from all other subscriptions (current and future) belonging to this BillForward Account.","verbs":["POST"]}
	 */
	AggregateAllSubscriptionsOnAccount *bool `json:"aggregateAllSubscriptionsOnAccount,omitempty"`

	/* {"default":"(empty list)","description":"[Required if and only if `productRatePlan` is omitted] List of components whose prices should be recalculated upon invoice aggregation. For example: two subscriptions' individual consumptions may neither of them be large enough to achieve bulk buy discounts. When aggregated, though, the same two subscriptions' consumption may add up to a quantity which does merit a bulk buy discount within your tiering system.","verbs":["POST"]}
	 */
	AggregatingComponents []*CreateAggregatingComponentRequest `json:"aggregatingComponents,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* {"description":"[Required if and only if `productRatePlan` is omitted] The currency of the product-rate-plan &mdash; as specified by a three-character ISO 4217 currency code (i.e. USD).","verbs":["POST"]}

	Required: true
	*/
	Currency string `json:"currency,omitempty"`

	/* {"default":"(null)","description":"Description of the created subscription. This is primarily for your benefit &mdash; for example, you could write here the mechanism through which you obtained this customer. (e.g. 'Business signed up using BUSYGUYS coupon, at management trade show').","verbs":["POST"]}
	 */
	Description *string `json:"description,omitempty"`

	/* {"description":"[Required if and only if `productRatePlan` is omitted] Number of length-measures which constitute the rate plan's period.","verbs":["POST"]}
	 */
	Duration int32 `json:"duration,omitempty"`

	/* {"description":"[Required if and only if `productRatePlan` is omitted] Measure describing the magnitude of the rate plan's period.","verbs":["POST"]}
	 */
	DurationPeriod *string `json:"durationPeriod,omitempty"`

	/* {"default":"(Subscription will be named after the rate plan to which the subscription subscribes)","description":"Name of the created subscription. This is primarily for your benefit &mdash; for example, to enable you to identify subscriptions at a glance in the BillForward web interface (e.g. 'BusinessCorp subscriptions, care of Mr Business (mr@busy.com)').","verbs":["POST"]}
	 */
	Name *string `json:"name,omitempty"`

	/* {"default":"(Auto-populated using your authentication credentials)","description":"ID of the BillForward Organization within which the requested Subscription should be created. If omitted, this will be auto-populated using your authentication credentials.","verbs":["POST"]}
	 */
	OrganizationID *string `json:"organizationID,omitempty"`

	/* {"description":"ID of the rate plan to which the subscription will be subscribing. If omitted: it will be assumed you wish to create a new rate plan as part of this request &mdash; this subscription will subscribe to that newly-created rate plan.","verbs":["POST"]}
	 */
	ProductRatePlan *string `json:"productRatePlan,omitempty"`

	/* {"description":"[Required if and only if `productRatePlan` is omitted] The frequency of the rate plan &mdash; either recurring or non-recurring.","verbs":["POST"]}
	 */
	ProductType *string `json:"productType,omitempty"`

	/* {"default":"(ServerNow upon receiving request)","description":"ISO 8601 UTC DateTime (e.g. 2015-06-16T11:58:41Z) describing the date at which the subscription should enter its first service period.","verbs":["POST"]}
	 */
	Start strfmt.DateTime `json:"start,omitempty"`

	/* {"default":"Provisioned","description":"The state in which the created subscription will begin.<br><span class=\"label label-default\">Provisioned</span> &mdash; The subscription will wait (without raising any invoices or beginning its service) until explicit action is taken to change its state.<br><span class=\"label label-default\">AwaitingPayment</span> &mdash; The subscription is activated. After `start` time is surpassed, it will begin service and raise its first invoice.","verbs":["POST"]}
	 */
	State *string `json:"state,omitempty"`
}

// Validate validates this create aggregating subscription request
func (m *CreateAggregatingSubscriptionRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAggregatingComponents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDurationPeriod(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateProductType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAggregatingSubscriptionRequest) validateAccountID(formats strfmt.Registry) error {

	if err := validate.RequiredString("accountID", "body", string(m.AccountID)); err != nil {
		return err
	}

	return nil
}

func (m *CreateAggregatingSubscriptionRequest) validateAggregatingComponents(formats strfmt.Registry) error {

	if swag.IsZero(m.AggregatingComponents) { // not required
		return nil
	}

	for i := 0; i < len(m.AggregatingComponents); i++ {

		if m.AggregatingComponents[i] != nil {

			if err := m.AggregatingComponents[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *CreateAggregatingSubscriptionRequest) validateCurrency(formats strfmt.Registry) error {

	if err := validate.RequiredString("currency", "body", string(m.Currency)); err != nil {
		return err
	}

	return nil
}

var createAggregatingSubscriptionRequestDurationPeriodEnum []interface{}

func (m *CreateAggregatingSubscriptionRequest) validateDurationPeriodEnum(path, location string, value string) error {
	if createAggregatingSubscriptionRequestDurationPeriodEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["minutes","days","months","years"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createAggregatingSubscriptionRequestDurationPeriodEnum = append(createAggregatingSubscriptionRequestDurationPeriodEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, createAggregatingSubscriptionRequestDurationPeriodEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateAggregatingSubscriptionRequest) validateDurationPeriod(formats strfmt.Registry) error {

	if swag.IsZero(m.DurationPeriod) { // not required
		return nil
	}

	if err := m.validateDurationPeriodEnum("durationPeriod", "body", *m.DurationPeriod); err != nil {
		return err
	}

	return nil
}

var createAggregatingSubscriptionRequestProductTypeEnum []interface{}

func (m *CreateAggregatingSubscriptionRequest) validateProductTypeEnum(path, location string, value string) error {
	if createAggregatingSubscriptionRequestProductTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["nonrecurring","recurring"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createAggregatingSubscriptionRequestProductTypeEnum = append(createAggregatingSubscriptionRequestProductTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, createAggregatingSubscriptionRequestProductTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateAggregatingSubscriptionRequest) validateProductType(formats strfmt.Registry) error {

	if swag.IsZero(m.ProductType) { // not required
		return nil
	}

	if err := m.validateProductTypeEnum("productType", "body", *m.ProductType); err != nil {
		return err
	}

	return nil
}

var createAggregatingSubscriptionRequestStateEnum []interface{}

func (m *CreateAggregatingSubscriptionRequest) validateStateEnum(path, location string, value string) error {
	if createAggregatingSubscriptionRequestStateEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Trial","Provisioned","Paid","AwaitingPayment","Cancelled","Failed","Expired"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			createAggregatingSubscriptionRequestStateEnum = append(createAggregatingSubscriptionRequestStateEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, createAggregatingSubscriptionRequestStateEnum); err != nil {
		return err
	}
	return nil
}

func (m *CreateAggregatingSubscriptionRequest) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}
