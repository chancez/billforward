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

/*Account Account

swagger:model Account
*/
type Account struct {

	/* { "description" : "If present, this will be the product rate plan to use when creating an aggregating subscription.  An account level aggregating subscription will be created when the first subscription is created against the account.", "verbs":[] }
	 */
	AggregatingProductRatePlanID *string `json:"aggregatingProductRatePlanID,omitempty"`

	/* { "description" : "The consistent ID of the account level aggregating subscription, if one exists.", "verbs":[] }
	 */
	AggregatingSubscriptionID *string `json:"aggregatingSubscriptionID,omitempty"`

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

	/* {  "default" : "false",  "description" : "Indicates if an account has been retired. If an account has been retired it can still be retrieved using the appropriate flag on API requests.", "verbs":["GET"] }
	 */
	Deleted *bool `json:"deleted,omitempty"`

	/* { "description" : "ID of the account.", "verbs":["POST","PUT","GET"] }

	Required: true
	*/
	ID string `json:"id,omitempty"`

	/* NotificationObjectGraph notification object graph
	 */
	NotificationObjectGraph *string `json:"notificationObjectGraph,omitempty"`

	/* { "description" : "Organization associated with the account.", "verbs":[] }

	Required: true
	*/
	OrganizationID string `json:"organizationID,omitempty"`

	/* { "description" : "The payment-methods associated with the account.", "verbs":["GET"] }
	 */
	PaymentMethods []*PaymentMethod `json:"paymentMethods,omitempty"`

	/* Profile profile
	 */
	Profile *Profile `json:"profile,omitempty"`

	/* { "description" : "Roles associated with the account. These are used to govern access privileges.", "verbs":[] }

	Required: true
	*/
	Roles []*InsertableBillingEntity `json:"roles,omitempty"`

	/* { "description" : "Number of distinct, paid subscriptions associated with this account. Initially the value will be 0 when no successful subscriptions have been taken. A subscription cancelled whilst in trial is counted as successful.", "verbs":["GET"] }
	 */
	SuccessfulSubscriptions int32 `json:"successfulSubscriptions,omitempty"`

	/* { "description" : "The UTC DateTime when the object was last updated. ", "verbs":[] }
	 */
	Updated strfmt.DateTime `json:"updated,omitempty"`

	/* { "description" : "User associated with the account. If this is null, no user is currently assocaited with the account. A user is only set when an account is associated with a user account.", "verbs":[] }
	 */
	UserID *string `json:"userID,omitempty"`
}

// Validate validates this account
func (m *Account) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillingEntity(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePaymentMethods(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var accountBillingEntityEnum []interface{}

func (m *Account) validateBillingEntityEnum(path, location string, value string) error {
	if accountBillingEntityEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["Notification","Organization","OrganizationGateway","Product","User","Subscription","Profile","ProductRatePlan","Client","Invoice","PricingComponentValue","Account","PricingComponentValueChange","PricingComponentTier","PricingComponent","PricingCalculation","CouponDefinition","CouponInstance","CouponModifier","CouponRule","CouponBookDefinition","CouponBook","InvoiceLine","Webhook","SubscriptionCancellation","NotificationSnapshot","InvoicePayment","InvoiceLinePayment","Payment","PaymentMethod","PaymentMethodSubscriptionLink","DunningLine","CybersourceToken","Card","Alias","PaypalSimplePaymentReconciliation","FreePaymentReconciliation","LocustworldPaymentReconciliation","CouponInstanceExistingValue","TaxLine","TaxationStrategy","TaxationLink","Address","AmendmentPriceNTime","Authority","UnitOfMeasure","SearchResult","Amendment","AuditLog","Password","Username","FixedTermDefinition","FixedTerm","Refund","CreditNote","Receipt","AmendmentCompoundConstituent","APIConfiguration","StripeToken","BraintreeToken","BalancedToken","PaypalToken","AuthorizeNetToken","SpreedlyToken","GatewayRevenue","AmendmentDiscardAmendment","CancellationAmendment","CompoundAmendment","CompoundAmendmentConstituent","FixedTermExpiryAmendment","InvoiceNextExecutionAttemptAmendment","PricingComponentValueAmendment","BraintreeMerchantAccount","WebhookSubscription","Migration","CassResult","CassPaymentResult","CassProductRatePlanResult","CassChurnResult","CassUpgradeResult","SubscriptionCharge","CassPaymentPProductResult","ProductPaymentsArgs","StripeACHToken","UsageAmount","UsageSession","Usage","UsagePeriod","Period","OfflinePayment","CreditNotePayment","CardVaultPayment","FreePayment","BraintreePayment","BalancedPayment","CybersourcePayment","PaypalPayment","PaypalSimplePayment","LocustWorldPayment","StripeOnlyPayment","ProductPaymentsResult","StripeACHPayment","AuthorizeNetPayment","CompoundUsageSession","CompoundUsage","UsageRoundingStrategies","BillforwardManagedPaymentsResult","PricingComponentValueMigrationChargeAmendmentMapping","SubscriptionLTVResult","AccountLTVResult","ProductRatePlanPaymentsResult","DebtsResult","AccountPaymentsResult","ComponentChange","QuoteRequest","Quote","CouponCharge","CouponInstanceInvoiceLink","Coupon","CouponDiscount","CouponUniqueCodesRequest","CouponUniqueCodesResponse","GetCouponsResponse","AddCouponCodeRequest","AddCouponCodeResponse","RemoveCouponFromSubscriptionRequest","TokenizationPreAuth","StripeTokenizationPreAuth","BraintreeTokenizationPreAuth","SpreedlyTokenizationPreAuth","SagePayTokenizationPreAuth","PayVisionTokenizationPreAuth","TokenizationPreAuthRequest","AuthCaptureRequest","StripeACHBankAccountVerification","PasswordReset","PricingRequest","AddTaxationStrategyRequest","AddPaymentMethodRequest","APIRequest","SagePayToken","SagePayNotificationRequest","SagePayNotificationResponse","SagePayOutstandingTransaction","SagePayEnabledCardType","TrustCommerceToken","SagePayTransaction","PricingComponentValueResponse","MigrationResponse","TimeResponse","EntityTime","Email","AggregationLink","BFPermission","Role","PermissionLink","PayVisionToken","PayVisionTransaction","KashToken","EmailProvider","DataSynchronizationJob","DataSynchronizationJobError","DataSynchronizationConfiguration","DataSynchronizationAppConfiguration","AggregationChildrenResponse","MetadataKeyValue","Metadata","AggregatingComponent","PricingComponentMigrationValue","InvoiceRecalculationAmendment","IssueInvoiceAmendment","EmailSubscription","RevenueAttribution"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			accountBillingEntityEnum = append(accountBillingEntityEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, accountBillingEntityEnum); err != nil {
		return err
	}
	return nil
}

func (m *Account) validateBillingEntity(formats strfmt.Registry) error {

	if swag.IsZero(m.BillingEntity) { // not required
		return nil
	}

	if err := m.validateBillingEntityEnum("billingEntity", "body", *m.BillingEntity); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Account) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organizationID", "body", string(m.OrganizationID)); err != nil {
		return err
	}

	return nil
}

func (m *Account) validatePaymentMethods(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentMethods) { // not required
		return nil
	}

	for i := 0; i < len(m.PaymentMethods); i++ {

		if m.PaymentMethods[i] != nil {

			if err := m.PaymentMethods[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *Account) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", m.Roles); err != nil {
		return err
	}

	for i := 0; i < len(m.Roles); i++ {

		if m.Roles[i] != nil {

			if err := m.Roles[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}
