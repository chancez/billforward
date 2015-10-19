package types

type Account struct {
	Id                      string           `json:"id"`
	SuccessfulSubscriptions int              `json:"successfulSubscriptions,omitempty"`
	Deleted                 bool             `json:"deleted,omitempty"`
	Profile                 *Profile         `json:"profile"`
	PaymentMethods          []*PaymentMethod `json:"paymentMethods"`
}

type AccountsResponse struct {
	ExecutionTime int        `json:"executionTime"`
	Results       []*Account `json:"results"`
}
