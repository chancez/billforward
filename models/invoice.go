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
An invoice encapsulates the cumulative charges of a subscription's pricing-components and their respective pricing-component-values for a specific period of time.

swagger:model Invoice
*/
type Invoice struct {

	/* { "description" : "", "verbs":["GET"] }

	Required: true
	*/
	AccountID string `json:"accountID,omitempty"`

	/* BillingEntity billing entity
	 */
	BillingEntity string `json:"billingEntity,omitempty"`

	/* { "description" : "ID of the user who last updated the entity.", "verbs":[] }
	 */
	ChangedBy string `json:"changedBy,omitempty"`

	/* Charges charges
	 */
	Charges []*SubscriptionCharge `json:"charges,omitempty"`

	/* Children children
	 */
	Children []*Invoice `json:"children,omitempty"`

	/* { "description" : "Cost of the invoice exclusive of tax.", "verbs":["GET"] }

	Required: true
	*/
	CostExcludingTax float64 `json:"costExcludingTax,omitempty"`

	/* { "description" : "The UTC DateTime when the object was created.", "verbs":[] }
	 */
	Created strfmt.DateTime `json:"created,omitempty"`

	/* { "description" : "The amount of credit to be returned to the pool.", "verbs":["GET"] }

	Required: true
	*/
	CreditRolledOver float64 `json:"creditRolledOver,omitempty"`

	/* { "description" : "The amount of credit to be returned to the pool excluding tax.", "verbs":["GET"] }
	 */
	CreditRolledOverExcludingTax float64 `json:"creditRolledOverExcludingTax,omitempty"`

	/* Crmid crmid
	 */
	Crmid string `json:"crmid,omitempty"`

	/* { "description" : "Currency of the invoice specified by a three character ISO 4217 currency code.", "verbs":["GET"] }

	Required: true
	*/
	Currency string `json:"currency,omitempty"`

	/* { "description" : "Indicates if an invoice has been retired. If an invoice has been retired it can still be retrieved using the appropriate flag on API requests. Generally invoices will not be retired.", "verbs":[] }

	Required: true
	*/
	Deleted bool `json:"deleted,omitempty"`

	/* { "description" : "The description of the invoice", "verbs":[] }
	 */
	Description string `json:"description,omitempty"`

	/* { "description" : "The amount of discounts for the invoice.", "verbs":["GET"] }
	 */
	DiscountAmount float64 `json:"discountAmount,omitempty"`

	/* { "description" : "The amount of discounts for the invoice excluding tax.", "verbs":["GET"] }
	 */
	DiscountAmountExcludingTax float64 `json:"discountAmountExcludingTax,omitempty"`

	/* Fields fields
	 */
	Fields map[string]string `json:"fields,omitempty"`

	/* { "description" : "UTC DateTime of the invoice's final execution attempt. The final execution attempt. This may be less than the next execution attempt if the next execution attempt never occurred.", "verbs":["GET"] }
	 */
	FinalExecutionAttempt strfmt.DateTime `json:"finalExecutionAttempt,omitempty"`

	/* { "description" : "ID of the invoice.", "verbs":["POST","PUT","GET"] }
	 */
	ID string `json:"id,omitempty"`

	/* { "description" : "Is this an initial invoice. An initial invoice is the first invoice generated for a subscription. Initial invoices will not have dunning applied to them and as such will only have a single payment attempt. For trial periods, the trial invoice is the initial invoice.", "verbs":["GET"] }

	Required: true
	*/
	InitialInvoice bool `json:"initialInvoice,omitempty"`

	/* { "description" : "Cost of the invoice inclusive of tax.", "verbs":["GET"] }

	Required: true
	*/
	InvoiceCost float64 `json:"invoiceCost,omitempty"`

	/* { "description" : "The collection of invoice-lines associated with the invoice.", "verbs":["GET"] }
	 */
	InvoiceLines []*InvoiceLine `json:"invoiceLines,omitempty"`

	/* { "description" : "Total amount of the invoice currently paid for. As the invoice may be paid by multiple payments, for example partly paid for by a voucher and then paid for by a card, this indicates the current paid amount of the invoice.", "verbs":["GET"] }
	 */
	InvoicePaid float64 `json:"invoicePaid,omitempty"`

	/* { "description" : "Payments used for this invoice. Multiple payments may be associated with the invoice.", "verbs":["GET"] }
	 */
	InvoicePayments []*InvoicePayment `json:"invoicePayments,omitempty"`

	/* { "description" : "Total amount of the invoice refunded.", "verbs":["GET"] }
	 */
	InvoiceRefunded float64 `json:"invoiceRefunded,omitempty"`

	/* { "description" : "The UTC DateTime when the invoice was first issued.", "verbs":["GET"] }
	 */
	Issued strfmt.DateTime `json:"issued,omitempty"`

	/* { "description" : "UTC DateTime of the invoice's last execution attempt. This was the last time an attempt was made to take payment for this invoice.", "verbs":["GET"] }
	 */
	LastExecutionAttempt strfmt.DateTime `json:"lastExecutionAttempt,omitempty"`

	/* { "description" : "If the subscription is locked, it will not be processed by the system", "verbs":[] }
	 */
	Locked string `json:"locked,omitempty"`

	/* { "description" : "Which system is responsible for managing the subscription.", "verbs":[] }
	 */
	ManagedBy string `json:"managedBy,omitempty"`

	/* { "description" : "The name of the invoice", "verbs":[] }
	 */
	Name string `json:"name,omitempty"`

	/* { "description" : "UTC DateTime of the invoice's next execution attempt. If the next execution attempt is greater than the period end for this invoice, the invoice will not receive another automatic execution attempt.", "verbs":["GET"] }
	 */
	NextExecutionAttempt strfmt.DateTime `json:"nextExecutionAttempt,omitempty"`

	/* Cost of the invoice before discounts, inclusive of tax.

	Required: true
	*/
	NonDiscountedCost float64 `json:"nonDiscountedCost,omitempty"`

	/* Cost of the invoice before discounts, inclusive of tax.

	Required: true
	*/
	NonDiscountedCostExcludingTax float64 `json:"nonDiscountedCostExcludingTax,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph string `json:"notificationObjectGraph,omitempty"`

	/* { "description" : "", "verbs":[] }

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* Paid paid
	 */
	Paid bool `json:"paid,omitempty"`

	/* Parent parent
	 */
	Parent *Invoice `json:"parent,omitempty"`

	/* { "description" : "", "verbs":[] }
	 */
	ParentInvoiceID string `json:"parentInvoiceID,omitempty"`

	/* { "description" : "UTC DateTime specifying when payment was received for the invoice.", "verbs":["GET"] }
	 */
	PaymentReceived strfmt.DateTime `json:"paymentReceived,omitempty"`

	/* { "description" : "End of the period being billed by this invoice, UTC DateTime.", "verbs":["GET"] }
	 */
	PeriodEnd strfmt.DateTime `json:"periodEnd,omitempty"`

	/* { "description" : "Start of the period being billed by this invoice, UTC DateTime", "verbs":["GET"] }
	 */
	PeriodStart strfmt.DateTime `json:"periodStart,omitempty"`

	/* ShortID short ID
	 */
	ShortID string `json:"shortID,omitempty"`

	/* { "description" : "Initially an invoice is set as unpaid. Once payment for the full value of the invoice has been received it is marked as paid. An invoice may be paid from various sources including cards, coupons or manual payments.", "verbs":["GET"] }

	Required: true
	*/
	State string `json:"state,omitempty"`

	/* { "description" : "", "verbs":["GET"] }
	 */
	SubscriptionID string `json:"subscriptionID,omitempty"`

	/* { "description" : "", "verbs":["GET"] }

	Required: true
	*/
	SubscriptionVersionID string `json:"subscriptionVersionID,omitempty"`

	/* TaxLines tax lines
	 */
	TaxLines []*InsertableBillingEntity `json:"taxLines,omitempty"`

	/* TotalDiscountAmount total discount amount
	 */
	TotalDiscountAmount float64 `json:"totalDiscountAmount,omitempty"`

	/* { "description" : "Number of payment attempts for this invoice. This includes any manual execution attempts.", "verbs":["GET"] }
	 */
	TotalExecutionAttempts int32 `json:"totalExecutionAttempts,omitempty"`

	/* TotalInvoiceCost total invoice cost
	 */
	TotalInvoiceCost float64 `json:"totalInvoiceCost,omitempty"`

	/* TotalNominalPaid total nominal paid
	 */
	TotalNominalPaid float64 `json:"totalNominalPaid,omitempty"`

	/* TotalNominalUnpaid total nominal unpaid
	 */
	TotalNominalUnpaid float64 `json:"totalNominalUnpaid,omitempty"`

	/* { "description" : "The type of the invoice. A subscription invoice is raised every time a subscription recurs. An amendment is created for intra-contract changes. An Adhoc invoice is created for payment that is taken out-of-band of a subscription. Finally the invoice generated for a trial period is marked as Trial.", "verbs":["GET"] }

	Required: true
	*/
	Type string `json:"type,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`

	/* { "description" : "Version identifier of the invoice.", "verbs":["GET"] }
	 */
	VersionID string `json:"versionID,omitempty"`

	/* { "description" : "The version number of the Invoice.  The first version of an Invoice is version number 1", "verbs":["GET"] }

	Required: true
	*/
	VersionNumber int32 `json:"versionNumber,omitempty"`

	/* ZeroCost zero cost
	 */
	ZeroCost bool `json:"zeroCost,omitempty"`
}

// Validate validates this invoice
func (m *Invoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBillingEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCostExcludingTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreditRolledOver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitialInvoice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoiceCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonDiscountedCost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNonDiscountedCostExcludingTax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptionVersionID(formats); err != nil {
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

func (m *Invoice) validateAccountID(formats strfmt.Registry) error {

	if err := validate.Required("accountID", "body", string(m.AccountID)); err != nil {
		return err
	}

	return nil
}

var invoiceBillingEntityEnum []interface{}

func (m *Invoice) validateBillingEntityEnum(path, location string, value string) error {
	if invoiceBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","Email","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","EmailProvider","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment","EmailSubscription","RevenueAttribution"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			invoiceBillingEntityEnum = append(invoiceBillingEntityEnum, v)
		}
	}
	return validate.Enum(path, location, value, invoiceBillingEntityEnum)
}

func (m *Invoice) validateBillingEntity(formats strfmt.Registry) error {

	if err := m.validateBillingEntityEnum("billingEntity", "body", m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateCostExcludingTax(formats strfmt.Registry) error {

	if err := validate.Required("costExcludingTax", "body", float64(m.CostExcludingTax)); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateCreditRolledOver(formats strfmt.Registry) error {

	if err := validate.Required("creditRolledOver", "body", float64(m.CreditRolledOver)); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateCurrency(formats strfmt.Registry) error {

	if err := validate.Required("currency", "body", string(m.Currency)); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateDeleted(formats strfmt.Registry) error {

	if err := validate.Required("deleted", "body", bool(m.Deleted)); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateInitialInvoice(formats strfmt.Registry) error {

	if err := validate.Required("initialInvoice", "body", bool(m.InitialInvoice)); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateInvoiceCost(formats strfmt.Registry) error {

	if err := validate.Required("invoiceCost", "body", float64(m.InvoiceCost)); err != nil {
		return err
	}

	return nil
}

var invoiceManagedByEnum []interface{}

func (m *Invoice) validateManagedByEnum(path, location string, value string) error {
	if invoiceManagedByEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["BillForward","Stripe"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			invoiceManagedByEnum = append(invoiceManagedByEnum, v)
		}
	}
	return validate.Enum(path, location, value, invoiceManagedByEnum)
}

func (m *Invoice) validateManagedBy(formats strfmt.Registry) error {

	if err := m.validateManagedByEnum("managedBy", "body", m.ManagedBy); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateNonDiscountedCost(formats strfmt.Registry) error {

	if err := validate.Required("nonDiscountedCost", "body", float64(m.NonDiscountedCost)); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateNonDiscountedCostExcludingTax(formats strfmt.Registry) error {

	if err := validate.Required("nonDiscountedCostExcludingTax", "body", float64(m.NonDiscountedCostExcludingTax)); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}

var invoiceStateEnum []interface{}

func (m *Invoice) validateStateEnum(path, location string, value string) error {
	if invoiceStateEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Paid","Unpaid","Pending","Voided"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			invoiceStateEnum = append(invoiceStateEnum, v)
		}
	}
	return validate.Enum(path, location, value, invoiceStateEnum)
}

func (m *Invoice) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", string(m.State)); err != nil {
		return err
	}

	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateSubscriptionVersionID(formats strfmt.Registry) error {

	if err := validate.Required("subscriptionVersionID", "body", string(m.SubscriptionVersionID)); err != nil {
		return err
	}

	return nil
}

var invoiceTypeEnum []interface{}

func (m *Invoice) validateTypeEnum(path, location string, value string) error {
	if invoiceTypeEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Subscription","Trial","Charge","FinalArrears","Amendment","Aggregated"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			invoiceTypeEnum = append(invoiceTypeEnum, v)
		}
	}
	return validate.Enum(path, location, value, invoiceTypeEnum)
}

func (m *Invoice) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", string(m.Type)); err != nil {
		return err
	}

	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateVersionNumber(formats strfmt.Registry) error {

	if err := validate.Required("versionNumber", "body", int32(m.VersionNumber)); err != nil {
		return err
	}

	return nil
}
