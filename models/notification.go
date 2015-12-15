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
Notification

swagger:model Notification
*/
type Notification struct {

	/* { "description" : "If true notifications will continue to be sent until an acknowledgment is received.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	AckEnabled bool `json:"ackEnabled,omitempty"`

	/* { "description" : "The UTC DateTime when the notification was acked if it is ack enabled.", "verbs":["POST","PUT","GET"] }
	 */
	Acked strfmt.DateTime `json:"acked,omitempty"`

	/* { "description" : "The action associated with the notification.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Action string `json:"action,omitempty"`

	/* BillingEntity billing entity
	 */
	BillingEntity string `json:"billingEntity,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* Changes changes
	 */
	Changes string `json:"changes,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "The URL the notification will be sent to.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	DestinationURL string `json:"destinationURL,omitempty"`

	/* { "description" : "The domain of the notification.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Domain string `json:"domain,omitempty"`

	/* Entity entity
	 */
	Entity string `json:"entity,omitempty"`

	/* { "description" : "The id of the entity associated with the notification.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	EntityID string `json:"entityID,omitempty"`

	/* Fields fields
	 */
	Fields map[string]string `json:"fields,omitempty"`

	/* { "description" : "The UTC DateTime of the notification's final send attempt.", "verbs":["POST","PUT","GET"] }
	 */
	FinalSendAttempt strfmt.DateTime `json:"finalSendAttempt,omitempty"`

	/* { "description" : "Format of the notification.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Format string `json:"format,omitempty"`

	/* { "description" : "ID of the notification.", "verbs":["POST","PUT","GET"] }
	 */
	ID string `json:"id,omitempty"`

	/* { "description" : "The UTC DateTime of the notifications's last send attempt.", "verbs":["POST","PUT","GET"] }
	 */
	LastSendAttempt strfmt.DateTime `json:"lastSendAttempt,omitempty"`

	/* { "description" : "The UTC DateTime of the notification's next send attempt.", "verbs":["POST","PUT","GET"] }
	 */
	NextSendAttempt strfmt.DateTime `json:"nextSendAttempt,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph string `json:"notificationObjectGraph,omitempty"`

	/* { "description" : "Organization associated with the notification.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* { "description" : "Whether the notification has been sent.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	State string `json:"state,omitempty"`

	/* { "description" : "The number of send attempts for this notification.", "verbs":["POST","PUT","GET"] }
	 */
	TotalSendAttempts int32 `json:"totalSendAttempts,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`

	/* { "description" : "Webhook associated with the notification.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	WebhookID string `json:"webhookID,omitempty"`
}

// Validate validates this notification
func (m *Notification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAckEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Notification) validateAckEnabled(formats strfmt.Registry) error {

	if err := validate.Required("ackEnabled", "body", bool(m.AckEnabled)); err != nil {
		return err
	}

	return nil
}

var notificationActionEnum []interface{}

func (m *Notification) validateActionEnum(path, location string, value string) error {
	if notificationActionEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Accept","Active","AwaitingPayment","AwaitingRefund","Cancelled","Completed","Created","Error","Expiring","Expired","Failed","Migrated","NeedsAmendments","Paid","Pending","Provisioned","Refunded","Reject","Trial","Unknown","Unpaid","Updated","Voided","PaymentFailed"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			notificationActionEnum = append(notificationActionEnum, v)
		}
	}
	return validate.Enum(path, location, value, notificationActionEnum)
}

func (m *Notification) validateAction(formats strfmt.Registry) error {

	if err := validate.Required("action", "body", string(m.Action)); err != nil {
		return err
	}

	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

var notificationBillingEntityEnum []interface{}

func (m *Notification) validateBillingEntityEnum(path, location string, value string) error {
	if notificationBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","Email","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","EmailProvider","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment","EmailSubscription","RevenueAttribution"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			notificationBillingEntityEnum = append(notificationBillingEntityEnum, v)
		}
	}
	return validate.Enum(path, location, value, notificationBillingEntityEnum)
}

func (m *Notification) validateBillingEntity(formats strfmt.Registry) error {

	if err := m.validateBillingEntityEnum("billingEntity", "body", m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateDestinationURL(formats strfmt.Registry) error {

	if err := validate.Required("destinationURL", "body", string(m.DestinationURL)); err != nil {
		return err
	}

	return nil
}

var notificationDomainEnum []interface{}

func (m *Notification) validateDomainEnum(path, location string, value string) error {
	if notificationDomainEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","Coupon","CouponDiscount","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","WebhookSubscription","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","AuthorizeNetToken","PaypalToken","SpreedlyToken","SagePayToken","TrustCommerceToken","PayVisionToken","SagePayOutstandingTransaction","SagePayEnabledCardType","SagePayTransaction","GatewayRevenue","Migration","AdhocSubscription","SubscriptionCharge","ComponentChange","Verification","UsageRoundingStrategies","PricingComponentValueMigrationChargeAmendmentMapping","AmendmentDiscardAmendment","EntityTime","AggregatingComponent","PricingComponentMigrationValue","MetadataKeyValue","Metadata","AggregationLink","BFPermission","Role","PermissionLink","PayVisionTransaction","KashToken","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","InvoiceLinePayment","EmailSubscription","EmailProvider","TimeResponse","Email","RevenueAttribution","Unknown"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			notificationDomainEnum = append(notificationDomainEnum, v)
		}
	}
	return validate.Enum(path, location, value, notificationDomainEnum)
}

func (m *Notification) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", string(m.Domain)); err != nil {
		return err
	}

	if err := m.validateDomainEnum("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateEntityID(formats strfmt.Registry) error {

	if err := validate.Required("entityID", "body", string(m.EntityID)); err != nil {
		return err
	}

	return nil
}

var notificationFormatEnum []interface{}

func (m *Notification) validateFormatEnum(path, location string, value string) error {
	if notificationFormatEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["JSON","XML"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			notificationFormatEnum = append(notificationFormatEnum, v)
		}
	}
	return validate.Enum(path, location, value, notificationFormatEnum)
}

func (m *Notification) validateFormat(formats strfmt.Registry) error {

	if err := validate.Required("format", "body", string(m.Format)); err != nil {
		return err
	}

	if err := m.validateFormatEnum("format", "body", m.Format); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}

var notificationStateEnum []interface{}

func (m *Notification) validateStateEnum(path, location string, value string) error {
	if notificationStateEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Unsent","Sending","Sent"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			notificationStateEnum = append(notificationStateEnum, v)
		}
	}
	return validate.Enum(path, location, value, notificationStateEnum)
}

func (m *Notification) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", string(m.State)); err != nil {
		return err
	}

	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Notification) validateWebhookID(formats strfmt.Registry) error {

	if err := validate.Required("webhookID", "body", string(m.WebhookID)); err != nil {
		return err
	}

	return nil
}