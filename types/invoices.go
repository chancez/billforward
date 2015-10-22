package types

import (
	"time"
)

type Invoice struct {
	Id                            string           `json:"id,omitempty"`
	AccountId                     string           `json:"accountID"`
	ChangedBy                     string           `json:"changedBy,omitempty"`
	CostExcludingTax              float64          `json:"costExcludingTax"`
	Created                       time.Time        `json:"created,omitempty"`
	CreditRolledOver              float64          `json:"creditRolledOver"`
	CreditRolledOverExcludingTax  float64          `json:"creditRolledOverExcludingTax,omitempty"`
	Currency                      string           `json:"currency"`
	Deleted                       bool             `json:"deleted"`
	DiscountAmount                float64          `json:"discountAmount,omitempty"`
	DiscountAmountExcludingTax    float64          `json:"discountAmountExcludingTax,omitempty"`
	FinalExecutionAttempt         time.Time        `json:"finalExecutionAttempt,omitempty"`
	InitialInvoice                bool             `json:"initialInvoice"`
	InvoiceCost                   float64          `json:"invoiceCost"`
	InvoiceLines                  []InvoiceLine    `json:"invoiceLines,omitempty"`
	InvoicePaid                   float64          `json:"invoicePaid,omitempty"`
	InvoicePayments               []InvoicePayment `json:"invoicePayments,omitempty"`
	InvoiceRefunded               float64          `json:"invoiceRefunded,omitempty"`
	Issued                        time.Time        `json:"issued,omitempty"`
	LastExecutionAttempt          time.Time        `json:"lastExecutionAttempt,omitempty"`
	Locked                        string           `json:"locked,omitempty"`
	ManagedBy                     string           `json:"managedBy,omitempty"`
	Name                          string           `json:"name,omitempty"`
	NextExecutionAttempt          time.Time        `json:"nextExecutionAttempt,omitempty"`
	NonDiscountedCost             float64          `json:"nonDiscountedCost"`
	NonDiscountedCostExcludingTax float64          `json:"nonDiscountedCostExcludingTax"`
	Paid                          bool             `json:"paid,omitempty"`
	PaymentReceived               time.Time        `json:"paymentReceived,omitempty"`
	PeriodEnd                     time.Time        `json:"periodEnd,omitempty"`
	PeriodStart                   time.Time        `json:"periodStart,omitempty"`
	State                         string           `json:"state"`
	SubscriptionId                string           `json:"subscriptionID,omitempty"`
	SubscriptionVersionId         string           `json:"subscriptionVersionID"`
	TotalDiscountAmount           float64          `json:"totalDiscountAmount,omitempty"`
	TotalExecutionAttempts        int32            `json:"totalExecutionAttempts,omitempty"`
	TotalInvoiceCost              float64          `json:"totalInvoiceCost,omitempty"`
	TotalNominalPaid              float64          `json:"totalNominalPaid,omitempty"`
	TotalNominalUnpaid            float64          `json:"totalNominalUnpaid,omitempty"`
	Type                          string           `json:"type"`
	Updated                       time.Time        `json:"updated,omitempty"`
	VersionId                     string           `json:"versionID,omitempty"`
	VersionNumber                 int32            `json:"versionNumber"`
	ZeroCost                      bool             `json:"zeroCost,omitempty"`
}

type InvoiceLine struct {
	Id                   string    `json:"id,omitempty"`
	InvoiceId            string    `json:"invoiceID"`
	UnitOfMeasureId      string    `json:"unitOfMeasureID,omitempty"`
	Calculation          string    `json:"calculation"`
	ChargeType           string    `json:"chargeType"`
	ComponentValue       int32     `json:"componentValue"`
	Cost                 float64   `json:"cost"`
	Created              time.Time `json:"created,omitempty"`
	Updated              time.Time `json:"updated,omitempty"`
	Description          string    `json:"description"`
	Discount             float64   `json:"discount,omitempty"`
	DiscountExcludingTax float64   `json:"discountExcludingTax,omitempty"`
	Name                 string    `json:"name"`
	Tax                  float64   `json:"tax"`
	Type                 string    `json:"type"`
}

type InvoicePayment struct {
	Created        time.Time `json:"created,omitempty"`
	Id             string    `json:"id,omitempty"`
	InvoiceId      string    `json:"invoiceID"`
	PaymentId      string    `json:"paymentID"`
	Currency       string    `json:"currency"`
	Gateway        string    `json:"gateway"`
	RefundedAmount float64   `json:"refundedAmount"`
	NominalAmount  float64   `json:"nominalAmount"`
	ActualAmount   float64   `json:"actualAmount"`
}

type InvoicesResponse struct {
	PaginatedResponse
	Results []*Invoice `json:"results"`
}
