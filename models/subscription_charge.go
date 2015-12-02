package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
	"github.com/go-swagger/go-swagger/strfmt"
)

/*
Represents a monetary amount &mdash; or quantity consumed &mdash; attributed to some invoice or subscription.

swagger:model SubscriptionCharge
*/
type SubscriptionCharge struct {

	/* {"description":"ID of the account owning the subscription for which the charge was generated.","verbs":["GET","POST"]}
	 */
	AccountID string `json:"accountID,omitempty"`

	/* {"description":"Monetary amount of the charge &mdash; including any tax applied to the final amount.","verbs":["POST","GET"]}
	 */
	Amount float64 `json:"amount,omitempty"`

	/* {"description":"Monetary amount of the charge &mdash; excluding any tax applied to the final amount.","verbs":["GET"]}
	 */
	AmountExcludingTax float64 `json:"amountExcludingTax,omitempty"`

	/* BillingEntity billing entity
	 */
	BillingEntity string `json:"billingEntity,omitempty"`

	/* {"default":"(Empty string)","description":"A human-readable explanation of how the value of the charge was calculated.","verbs":["GET"]}
	 */
	Calculation string `json:"calculation,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* {"default":"<span class=\"label label-default\">Debit</span>","description":"Whether this charge represents money given to or taken from the customer.<br><span class=\"label label-default\">Credit</span> &mdash; This is a charge for money given to the customer.<br><span class=\"label label-default\">Debit</span> &mdash; This is a charge for money taken from the customer.","verbs":["POST","GET"]}

	Required: true
	*/
	ChargeType string `json:"chargeType,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* {"description":"Customer-relationship-management ID of the charge.","verbs":["GET","PUT","POST"]}
	 */
	CrmID string `json:"crmID,omitempty"`

	/* { "description" : "Currency of the invoice specified by a three character ISO 4217 currency code.", "verbs":["GET"] }
	 */
	Currency string `json:"currency,omitempty"`

	/* {"description":"Description given to the charge.","verbs":["POST","GET","PUT"]}
	 */
	Description string `json:"description,omitempty"`

	/* Fields fields
	 */
	Fields map[string]string `json:"fields,omitempty"`

	/* { "description" : "", "verbs":["GET", "PUT"] }
	 */
	ID string `json:"id,omitempty"`

	/* {"description":"The invoice to which this charge applies (if the charge targets a specific invoice).","verbs":["GET"]}
	 */
	Invoice *Invoice `json:"invoice,omitempty"`

	/* {"description":"ID of the invoice to which this charge applies (if the charge targets a specific invoice).","verbs":["POST","GET"]}
	 */
	InvoiceID string `json:"invoiceID,omitempty"`

	/* {"default":"<span class=\"label label-default\">Aggregated</span>","description":"The strategy for how this charge will raise invoices.<br><span class=\"label label-default\">Immediate</span> &mdash; Generate straight-away an invoice containing this charge.<br><span class=\"label label-default\">Aggregated</span> &mdash; Add this charge to the next invoice which is generated naturally &mdash; i.e. the invoice raised at the current period's end.","verbs":["POST","GET"]}

	Required: true
	*/
	InvoicingType string `json:"invoicingType,omitempty"`

	/* {"description":"Friendly name given to the charge to help identify it.","verbs":["POST","GET","PUT"]}
	 */
	Name string `json:"name,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph string `json:"notificationObjectGraph,omitempty"`

	/* {"default":"(End of current period)","description":"The time-ending of the interval to which the charge applies. This can be used to apply a charge across partial or multiple periods,to pro-rate its price.","verbs":["POST","GET"]}
	 */
	PeriodEnd strfmt.DateTime `json:"periodEnd,omitempty"`

	/* {"default":"(Now)","description":"The time-beginning of the interval to which the charge applies. This can be used to apply a charge across partial or multiple periods,to pro-rate its price.","verbs":["POST","GET"]}
	 */
	PeriodStart strfmt.DateTime `json:"periodStart,omitempty"`

	/* {"default":"<span class=\"label label-default\">Rollover</span>","description":"Defines the behaviour applied to any outstanding credit resulting from the application of the charge..<br><span class=\"label label-default\">Rollover</span> &mdash; Outstanding credit is returned to the accounts credit pool.<br><span class=\"label label-default\">Discard</span> &mdash; Outstanding credit is lost.","verbs":["POST","GET"]}

	Required: true
	*/
	RemainingCreditBehaviour string `json:"remainingCreditBehaviour,omitempty"`

	/* {"default":"<span class=\"label label-default\">Pending</span>","description":"The current state of the charge.","verbs":["POST","GET","PUT"]}

	Required: true
	*/
	State string `json:"state,omitempty"`

	/* {"description":"ID of the subscription for which the charge was generated.","verbs":["GET","POST"]}
	 */
	SubscriptionID string `json:"subscriptionID,omitempty"`

	/* {"description":"Version ID of the subscription for which the charge was generated.","verbs":["GET","POST"]}
	 */
	SubscriptionVersionID string `json:"subscriptionVersionID,omitempty"`

	/* {"default":"false","description":"(Applicable only if any of [`pricingComponentName`, `pricingComponentID`] are defined)<br>Whether the charge was created for a subscription whilst in a trial period.<br><span class=\"label label-default\">false</span> &mdash; This is a non-trial charge, so funds will be sought from the customer.<br><span class=\"label label-default\">true</span> &mdash; This is a trial charge, soThe charge can be considered 'Paid' without taking any funds from the customer.","verbs":["POST","GET"]}
	 */
	Trial bool `json:"trial,omitempty"`

	/* {"description":"A type describing the nature of the charge.","verbs":["POST","GET"]}

	Required: true
	*/
	Type string `json:"type,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`

	/* {"description":"The version ID of the charge.","verbs":["GET"]}
	 */
	VersionID string `json:"versionID,omitempty"`

	/* {"description":"The version number of the charge. The first version of a charge is version number 1.","verbs":["GET"]}

	Required: true
	*/
	VersionNumber int32 `json:"versionNumber,omitempty"`
}

// Validate validates this subscription charge
func (m *SubscriptionCharge) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChargeType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoicingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemainingCreditBehaviour(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersionNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var subscriptionChargeBillingEntityEnum []interface{}

func (m *SubscriptionCharge) validateBillingEntityEnum(path, location string, value string) error {
	if subscriptionChargeBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","Email","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","EmailProvider","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment","EmailSubscription","RevenueAttribution"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			subscriptionChargeBillingEntityEnum = append(subscriptionChargeBillingEntityEnum, v)
		}
	}
	return validate.Enum(path, location, value, subscriptionChargeBillingEntityEnum)
}

func (m *SubscriptionCharge) validateBillingEntity(formats strfmt.Registry) error {

	if err := m.validateBillingEntityEnum("billingEntity", "body", m.BillingEntity); err != nil {
		return err
	}

	return nil
}

var subscriptionChargeChargeTypeEnum []interface{}

func (m *SubscriptionCharge) validateChargeTypeEnum(path, location string, value string) error {
	if subscriptionChargeChargeTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Credit","Debit"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			subscriptionChargeChargeTypeEnum = append(subscriptionChargeChargeTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, subscriptionChargeChargeTypeEnum)
}

func (m *SubscriptionCharge) validateChargeType(formats strfmt.Registry) error {

	if err := validate.Required("chargeType", "body", string(m.ChargeType)); err != nil {
		return err
	}

	if err := m.validateChargeTypeEnum("chargeType", "body", m.ChargeType); err != nil {
		return err
	}

	return nil
}

var subscriptionChargeInvoicingTypeEnum []interface{}

func (m *SubscriptionCharge) validateInvoicingTypeEnum(path, location string, value string) error {
	if subscriptionChargeInvoicingTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Immediate","Aggregated"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			subscriptionChargeInvoicingTypeEnum = append(subscriptionChargeInvoicingTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, subscriptionChargeInvoicingTypeEnum)
}

func (m *SubscriptionCharge) validateInvoicingType(formats strfmt.Registry) error {

	if err := validate.Required("invoicingType", "body", string(m.InvoicingType)); err != nil {
		return err
	}

	if err := m.validateInvoicingTypeEnum("invoicingType", "body", m.InvoicingType); err != nil {
		return err
	}

	return nil
}

var subscriptionChargeRemainingCreditBehaviourEnum []interface{}

func (m *SubscriptionCharge) validateRemainingCreditBehaviourEnum(path, location string, value string) error {
	if subscriptionChargeRemainingCreditBehaviourEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Rollover","Discard"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			subscriptionChargeRemainingCreditBehaviourEnum = append(subscriptionChargeRemainingCreditBehaviourEnum, v)
		}
	}
	return validate.Enum(path, location, value, subscriptionChargeRemainingCreditBehaviourEnum)
}

func (m *SubscriptionCharge) validateRemainingCreditBehaviour(formats strfmt.Registry) error {

	if err := validate.Required("remainingCreditBehaviour", "body", string(m.RemainingCreditBehaviour)); err != nil {
		return err
	}

	if err := m.validateRemainingCreditBehaviourEnum("remainingCreditBehaviour", "body", m.RemainingCreditBehaviour); err != nil {
		return err
	}

	return nil
}

var subscriptionChargeStateEnum []interface{}

func (m *SubscriptionCharge) validateStateEnum(path, location string, value string) error {
	if subscriptionChargeStateEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Voided","Pending","AwaitingPayment","Paid","Failed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			subscriptionChargeStateEnum = append(subscriptionChargeStateEnum, v)
		}
	}
	return validate.Enum(path, location, value, subscriptionChargeStateEnum)
}

func (m *SubscriptionCharge) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", string(m.State)); err != nil {
		return err
	}

	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var subscriptionChargeTypeEnum []interface{}

func (m *SubscriptionCharge) validateTypeEnum(path, location string, value string) error {
	if subscriptionChargeTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Setup","Upgrade","Manual","ProductRatePlanMigration","Arrears","Advance","Coupon","Usage","PricingComponent"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			subscriptionChargeTypeEnum = append(subscriptionChargeTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, subscriptionChargeTypeEnum)
}

func (m *SubscriptionCharge) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", string(m.Type)); err != nil {
		return err
	}

	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionCharge) validateVersionNumber(formats strfmt.Registry) error {

	if err := validate.Required("versionNumber", "body", int32(m.VersionNumber)); err != nil {
		return err
	}

	return nil
}
