package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"io"
	"io/ioutil"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*PricingComponent PricingComponent pricing component

swagger:discriminator PricingComponent @type
*/
type PricingComponent interface {
	httpkit.Validatable

	/* { "description" : "", "default" : "", "verbs":["POST","GET"] }

	Required: true
	*/
	Type() string
	SetType(string)

	/* BillingEntity billing entity
	 */
	BillingEntity() *string
	SetBillingEntity(*string)

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy() *string
	SetChangedBy(*string)

	/* { "description" : "The charge model of the pricing-component.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ChargeModel() string
	SetChargeModel(string)

	/* { "description" : "The charge type of the pricing-component.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ChargeType() string
	SetChargeType(string)

	/* ComponentValue component value
	 */
	ComponentValue() int32
	SetComponentValue(int32)

	/* { "description" : "ID of the pricing-component. This ID does not change with new versions of the pricing component.", "verbs":["POST","PUT","GET"] } When associating a pricing component with a product rate plan, this ID should be used.

	Required: true
	*/
	ConsistentID() string
	SetConsistentID(string)

	/* Cost cost
	 */
	Cost() *float64
	SetCost(*float64)

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created() strfmt.DateTime
	SetCreated(strfmt.DateTime)

	/* Crmid crmid
	 */
	Crmid() *string
	SetCrmid(*string)

	/* { "description" : "The default quantity of the pricing-component. If no value is supplied on a subscription this value will be used. This is useful for setting an expected purchase level of for introducing new pricing components to existing subscriptions and not having to back-fill values", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	DefaultQuantity() int32
	SetDefaultQuantity(int32)

	/* { "description" : "", "verbs":["POST","PUT","GET"] }
	 */
	Description() *string
	SetDescription(*string)

	/* {"default":"<span class=\"label label-default\">delayed</span>","description":"Default behaviour when a value is downgraded using this pricing component, this behaviour can be overridden when changing the value.<br><span class=\"label label-default\">immediate</span> &mdash; Downgrade will apply at the time the request is made.<br><span class=\"label label-default\">delayed</span> &mdash; Downgrade will apply at the end of the current billing cycle.","verbs":["POST","GET"]}
	 */
	DowngradeMode() *string
	SetDowngradeMode(*string)

	/* { "description" : "Version ID of the pricing-component. Unique ID for each version of a pricing-component.", "verbs":["POST","PUT","GET"] }
	 */
	ID() *string
	SetID(*string)

	/* { "default" : "Aggregated",  "description" : "For <span class=\"label label-default\">setup</span> pricing components <span class=\"label label-default\">Immediate</span> invoicing will result in an invoice being issued on subscription being set to the AwaitingPayment state, irrespective of the subscription start date. <span class=\"label label-default\">Aggregated</span> invoicing will add a charge to the first invoice of the subscription.", "verbs":["POST","PUT","GET"] }
	 */
	InvoicingType() *string
	SetInvoicingType(*string)

	/* { "description" : "The maximum quantity of the pricing-component.", "verbs":[] }
	 */
	MaxQuantity() int32
	SetMaxQuantity(int32)

	/* { "default" : "0", "description" : "The minimum quantity of the pricing-component.", "verbs":[] }
	 */
	MinQuantity() int32
	SetMinQuantity(int32)

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Name() string
	SetName(string)

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph() *string
	SetNotificationObjectGraph(*string)

	/* { "description" : "", "verbs":[] }

	Required: true
	*/
	OrganizationID() string
	SetOrganizationID(string)

	/* PriceExplanation price explanation
	 */
	PriceExplanation() []string
	SetPriceExplanation([]string)

	/* PriceExplanationString price explanation string
	 */
	PriceExplanationString() *string
	SetPriceExplanationString(*string)

	/* { "description" : "The product-rate-plan associated with the pricing-component.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ProductRatePlan() *ProductRatePlan
	SetProductRatePlan(*ProductRatePlan)

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ProductRatePlanID() string
	SetProductRatePlanID(string)

	/* {"description":"A friendly non-unique name used to identify this pricing-component","verbs":["POST","PUT","GET"]}
	 */
	PublicName() *string
	SetPublicName(*string)

	/* {  "default" : "[]", "description" : "The pricing-component-tiers associated with the pricing-component.", "verbs":["POST","PUT","GET"] }
	 */
	Tiers() []*PricingComponentTier
	SetTiers([]*PricingComponentTier)

	/* UnitOfMeasure unit of measure
	 */
	UnitOfMeasure() *UnitOfMeasure
	SetUnitOfMeasure(*UnitOfMeasure)

	/* { "description" : "", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	UnitOfMeasureID() string
	SetUnitOfMeasureID(string)

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated() strfmt.DateTime
	SetUpdated(strfmt.DateTime)

	/* {"default":"<span class=\"label label-default\">immediate</span>","description":"Default behaviour when a value is upgraded using this pricing component, this behaviour can be overridden when changing the value.<br><span class=\"label label-default\">immediate</span> &mdash; Upgrade will apply at the time the request is made.<br><span class=\"label label-default\">delayed</span> &mdash; Upgrade will apply at the end of the current billing cycle.","verbs":["POST","GET"]}
	 */
	UpgradeMode() *string
	SetUpgradeMode(*string)

	/* { "default" : "current time", "description" : "The DateTime specifying when the pricing-component is valid from.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ValidFrom() strfmt.DateTime
	SetValidFrom(strfmt.DateTime)

	/* {  "default" : "null", "description" : "The UTC DateTime specifying when the pricing-component is valid till.", "verbs":["POST","PUT","GET"] }
	 */
	ValidTill() strfmt.DateTime
	SetValidTill(strfmt.DateTime)
}

// UnmarshalPricingComponentSlice unmarshals polymorphic slices of PricingComponent
func UnmarshalPricingComponentSlice(reader io.Reader, consumer httpkit.Consumer) ([]PricingComponent, error) {
	var elements [][]byte
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []PricingComponent
	for _, element := range elements {
		obj, err := unmarshalPricingComponent(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalPricingComponent unmarshals polymorphic PricingComponent
func UnmarshalPricingComponent(reader io.Reader, consumer httpkit.Consumer) (PricingComponent, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalPricingComponent(data, consumer)
}

func unmarshalPricingComponent(data []byte, consumer httpkit.Consumer) (PricingComponent, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the @type property.
	var getType struct {
		Type string `json:"@type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("@type", "body", getType.Type); err != nil {
		return nil, err
	}

	// The value of @type is used to determine which type to create and unmarshal the data into
	switch getType.Type {
	case "flatPricingComponent":
		var result FlatPricingComponent
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "tieredPricingComponent":
		var result TieredPricingComponent
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	case "tieredVolumePricingComponent":
		var result TieredVolumePricingComponent
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid @type value: %q", getType.Type)

}
