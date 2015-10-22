package types

/*
   "nextPage" : "/accounts?organizations=ORG-82939914-0AF3-4B74-B515-9204EA4C&offset=10&records=10&order_by=created&order=DESC&include_retired=true",
   "currentPage" : 0,
   "currentOffset" : 0,
   "recordsRequested" : 10,
   "recordsReturned" : 4,
   "executionTime" : 100798,
*/

type PaginatedResponse struct {
	NextPage         string `json:"nextPage"`
	CurrentPage      int64  `json:"currentPage"`
	CurrentOffset    int64  `json:"currentOffset"`
	RecordsRequested int64  `json:"recordsRequested"`
	RecordsReturned  int64  `json:"recordsReturned"`
	ExecutionTime    int    `json:"executionTime"`
}
