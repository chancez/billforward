package types

type Account struct {
	Id                      string           `json:"id"`
	SuccessfulSubscriptions int              `json:"successfulSubscriptions,omitempty"`
	Deleted                 bool             `json:"deleted,omitempty"`
	Profile                 *Profile         `json:"profile"`
	PaymentMethods          []*PaymentMethod `json:"paymentMethods"`
}

/*
   "nextPage" : "/accounts?organizations=ORG-82939914-0AF3-4B74-B515-9204EA4C&offset=10&records=10&order_by=created&order=DESC&include_retired=true",
   "currentPage" : 0,
   "currentOffset" : 0,
   "recordsRequested" : 10,
   "recordsReturned" : 4,
   "executionTime" : 100798,
*/
type AccountsResponse struct {
	NextPage         string     `json:"nextPage"`
	CurrentPage      int64      `json:"currentPage"`
	CurrentOffset    int64      `json:"currentOffset"`
	RecordsRequested int64      `json:"recordsRequested"`
	RecordsReturned  int64      `json:"recordsReturned"`
	ExecutionTime    int        `json:"executionTime"`
	Results          []*Account `json:"results"`
}
