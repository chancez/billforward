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

/*Payment Payment

swagger:model Payment
*/
type Payment struct {

	/* { "description" : "Actual monetary value of the payment. This is real value of the payment amount, thus the amount of currency received for this payment. A coupon has a real value of zero, even though the nominal amount it pays may be greater than zero.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ActualValue float64 `json:"actualValue,omitempty"`

	/* BillingEntity billing entity
	 */
	BillingEntity *string `json:"billingEntity,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy *string `json:"changedBy,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* Crmid crmid
	 */
	Crmid *string `json:"crmid,omitempty"`

	/* { "description" : "The currency of the payment specified by a three character ISO 4217 currency code.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Currency string `json:"currency,omitempty"`

	/* Fields fields
	 */
	Fields map[string]string `json:"fields,omitempty"`

	/* { "description" : "Payment gateway associated with the payment", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Gateway string `json:"gateway,omitempty"`

	/* { "description" : "ID of the payment.", "verbs":["POST","PUT","GET"] }
	 */
	ID *string `json:"id,omitempty"`

	/* { "description" : "ID of the invoice associated with the payment. This may be null when a payment is not explicitly associated with an invoice. Payments may be used across multiple invoices.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	InvoiceID string `json:"invoiceID,omitempty"`

	/* { "description" : "Nominal value of the payment. This is the theoretical value of the payment, thus the value this payment can pay off an invoice. For example a coupon has a nominal value of the discount, and the actual value is zero.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	NominalValue float64 `json:"nominalValue,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph *string `json:"notificationObjectGraph,omitempty"`

	/* { "description" : "ID of the organization associated with the payment.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* { "description" : "ID of the payment method associated with the payment.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	PaymentMethodID string `json:"paymentMethodID,omitempty"`

	/* { "description" : "UTC DateTime specifying when payment was received for the invoice.", "verbs":["POST","PUT","GET"] }
	 */
	PaymentReceived strfmt.DateTime `json:"paymentReceived,omitempty"`

	/* { "description" : "Refunded nominal value of the payment.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	RefundedValue float64 `json:"refundedValue,omitempty"`

	/* { "description" : "Remaining nominal value of the payment not used. In the case when a payment is used across a range of invoices a payment may be used multiple times, each use reducing the available blance of the payment.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	RemainingValue float64 `json:"remainingValue,omitempty" xml:"remainingNominalValue"`

	/* { "description" : "Type of payment.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	Type string `json:"type,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this payment
func (m *Payment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBillingEntity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFields(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInvoiceID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNominalValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethodID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRefundedValue(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRemainingValue(formats); err != nil {
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

func (m *Payment) validateActualValue(formats strfmt.Registry) error {

	if err := validate.Required("actualValue", "body", float64(m.ActualValue)); err != nil {
		return err
	}

	return nil
}

var paymentBillingEntityEnum []interface{}

func (m *Payment) validateBillingEntityEnum(path, location string, value string) error {
	if paymentBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","Email","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","EmailProvider","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment","EmailSubscription","RevenueAttribution"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			paymentBillingEntityEnum = append(paymentBillingEntityEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, paymentBillingEntityEnum); err != nil {
		return err
	}
	return nil
}

func (m *Payment) validateBillingEntity(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingEntity) { // not required
		return nil
	}

	if err := m.validateBillingEntityEnum("billingEntity", "body", *m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", string(m.Currency)); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateFields(formats strfmt.Registry) error {

	if swag.IsZero(m.Fields) { // not required
		return nil
	}

	if err := validate.Required("fields", "body", m.Fields); err != nil {
		return err
	}

	return nil
}

var paymentGatewayEnum []interface{}

func (m *Payment) validateGatewayEnum(path, location string, value string) error {
	if paymentGatewayEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["cybersource_token","card_vault","paypal_simple","locustworld","free","coupon","credit_note","stripe","braintree","balanced","paypal","billforward_test","offline","trial","stripeACH","authorizeNet","spreedly","sagePay","trustCommerce","payvision","kash"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			paymentGatewayEnum = append(paymentGatewayEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, paymentGatewayEnum); err != nil {
		return err
	}
	return nil
}

func (m *Payment) validateGateway(formats strfmt.Registry) error {

	if err := validate.Required("gateway", "body", string(m.Gateway)); err != nil {
		return err
	}

	if err := m.validateGatewayEnum("gateway", "body", m.Gateway); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateInvoiceID(formats strfmt.Registry) error {

	if err := validate.Required("invoiceID", "body", string(m.InvoiceID)); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateNominalValue(formats strfmt.Registry) error {

	if err := validate.Required("nominalValue", "body", float64(m.NominalValue)); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validatePaymentMethodID(formats strfmt.Registry) error {

	if err := validate.Required("paymentMethodID", "body", string(m.PaymentMethodID)); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateRefundedValue(formats strfmt.Registry) error {

	if err := validate.Required("refundedValue", "body", float64(m.RefundedValue)); err != nil {
		return err
	}

	return nil
}

func (m *Payment) validateRemainingValue(formats strfmt.Registry) error {

	if err := validate.Required("remainingValue", "body", float64(m.RemainingValue)); err != nil {
		return err
	}

	return nil
}

var paymentTypeEnum []interface{}

func (m *Payment) validateTypeEnum(path, location string, value string) error {
	if paymentTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["credit","debit"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			paymentTypeEnum = append(paymentTypeEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, paymentTypeEnum); err != nil {
		return err
	}
	return nil
}

func (m *Payment) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", string(m.Type)); err != nil {
		return err
	}

	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}