package main

import (
	"bytes"
	"log"
	"os"

	"github.com/authclub/billforward/client"
	"github.com/authclub/billforward/client/accounts"
	"github.com/authclub/billforward/client/product_rate_plans"
	"github.com/authclub/billforward/models"

	httpclient "github.com/go-swagger/go-swagger/httpkit/client"
	"github.com/go-swagger/go-swagger/spec"
)

type ReadWrapper struct {
	*bytes.Reader
}

func (r *ReadWrapper) Close() error {
	return nil
}

func main() {
	log.Println("starting")
	swaggerSpec, err := spec.New(client.SwaggerJSON, "")
	if err != nil {
		log.Fatal(err)
	}
	transport := httpclient.New(swaggerSpec)
	// transport.Debug = true
	transport.Host = "api-sandbox.billforward.net:443"
	transport.BasePath = "/v1"
	transport.DefaultAuthentication = httpclient.BearerToken(os.Getenv("BILLFORWARD_API_KEY"))
	// transport.Consumers["application/pdf"] = httpkit.ConsumerFunc(func(r io.Reader, data interface{}) error {
	// 	f := data.(*httpkit.File)
	// 	b, err := ioutil.ReadAll(r)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	f.Data = &ReadWrapper{bytes.NewReader(b)}
	// 	return nil
	// })
	bfClient := client.New(transport, nil)

	knownAccountID := "ACC-06E62D37-8246-4E54-B72A-22170478"

	var acct *models.Account
	log.Println("getting account by ID", knownAccountID)
	getAcctResp, err := bfClient.Accounts.GetAccountByID(accounts.GetAccountByIDParams{
		AccountID: knownAccountID,
	})
	if err != nil {
		log.Fatal(err)
	}
	acct = getAcctResp.Payload.Results[0]
	log.Println("acct:", acct)

	// log.Println("getting all accounts")
	// acctsResp, err := bfClient.Accounts.GetAllAccounts(accounts.GetAllAccountsParams{
	// 	Order:   "DESC",
	// 	OrderBy: "created",
	// })
	// if err != nil {
	// 	gerr := err.(accounts.APIError).Response.(*accounts.GetAllAccountsDefault).Payload
	// 	log.Fatal(gerr)
	// }
	// for _, result := range acctsResp.Payload.Results {
	// 	if !result.Deleted && result.ID == knownAccountID {
	// 		log.Println("found existing account", result)
	// 		acct = result
	// 		break
	// 	}
	// }

	// Only create the account if it doesnt exist
	if acct == nil {
		log.Println("creating new account")
		createAcctResp, err := bfClient.Accounts.CreateAccount(accounts.CreateAccountParams{
			Request: &models.CreateAccountRequest{
				Profile: &models.Profile{
					FirstName:   "chancetester3",
					LastName:    "tester3",
					Email:       "chance.zibolski+test3@coreos.com",
					CompanyName: "enterprise company",
				},
				AggregatingProductRatePlanID: "PRP-3DBAA1F3-0A14-40D1-80CF-1E4BC674",
			},
		})
		if err != nil {
			if apiErr, ok := err.(accounts.APIError); ok {
				if gerr, ok := apiErr.Response.(*accounts.CreateAccountDefault); ok {
					log.Fatal(gerr.Payload)
				}
			}
			log.Fatal(err)
		}
		acct = createAcctResp.Payload.Results[0]
	}

	// invoicesResp, err := bfClient.Invoices.GetInvoicesByAccountID(invoices.GetInvoicesByAccountIDParams{
	// 	AccountID: acct.ID,
	// })
	// if err != nil {
	// 	if apiErr, ok := err.(invoices.APIError); ok {
	// 		if gerr, ok := apiErr.Response.(*invoices.GetInvoicesByAccountIDDefault); ok {
	// 			log.Fatal(gerr.Payload)
	// 		}
	// 	}
	// 	log.Fatal(err)
	// }
	// for _, invoice := range invoicesResp.Payload.Results {
	// 	log.Printf("%+v\n", invoice)
	// }

	// invoiceID := "INV-DF41079A-40A4-4155-901D-9CB5E892"
	// invoicePDFResp, err := bfClient.Invoices.GetInvoiceAsPDF(invoices.GetInvoiceAsPDFParams{
	// 	ID: invoiceID,
	// })
	// if err != nil {
	// 	if apiErr, ok := err.(invoices.APIError); ok {
	// 		if gerr, ok := apiErr.Response.(*invoices.GetInvoiceAsPDFDefault); ok {
	// 			log.Fatal(gerr.Payload)
	// 		}
	// 	}
	// 	log.Fatal(err)
	// }
	// b, err := ioutil.ReadAll(invoicePDFResp.Payload.Data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("invoiceCsvResp", string(b))

	log.Println("GetAllRatePlans")
	ratePlansResp, err := bfClient.ProductRatePlans.GetAllRatePlans(product_rate_plans.GetAllRatePlansParams{
		OrderBy: "created",
		Records: 10,
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, rp := range ratePlansResp.Payload.Results {
		log.Println("product:", rp.Product.Name)
		log.Println("rate plan:", rp.Name)
		for _, pc := range rp.PricingComponents {
			log.Printf("Line item: %s: %d (%s)", pc.Name, pc.DefaultQuantity, pc.UnitOfMeasure.Name)
		}
	}

	// quayEnterprise := "PRO-8762520A-EC4F-48BC-8AA3-3AC1709F"
	// log.Println("GetRatePlanByProduct")
	// ratePlanResp, err := bfClient.ProductRatePlans.GetRatePlanByProduct(product_rate_plans.GetRatePlanByProductParams{
	// 	ProductID: quayEnterprise,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, rp := range ratePlanResp.Payload.Results {
	// 	log.Println("product:", rp.Product.Name)
	// 	log.Println("rate plan:", rp.Name)
	// 	for _, pc := range rp.PricingComponents {
	// 		log.Printf("Line item: %s: %d (%s)", pc.Name, pc.DefaultQuantity, pc.UnitOfMeasure.Name)
	// 	}
	// }

	// quayEnterprise := "PRO-8762520A-EC4F-48BC-8AA3-3AC1709F"
	// ratePlanID := "PRP-3DBAA1F3-0A14-40D1-80CF-1E4BC674"
	// log.Println("GetRatePlanByProductAndRatePlan")
	// ratePlanResp, err := bfClient.ProductRatePlans.GetRatePlanByProductAndRatePlan(product_rate_plans.GetRatePlanByProductAndRatePlanParams{
	// 	ProductID:  quayEnterprise,
	// 	RatePlanID: ratePlanID,
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, rp := range ratePlanResp.Payload.Results {
	// 	log.Println("product:", rp.Product.Name)
	// 	log.Println("rate plan:", rp.Name)
	// 	for _, pc := range rp.PricingComponents {
	// 		log.Printf("Line item: %s: %d (%s)", pc.Name, pc.DefaultQuantity, pc.UnitOfMeasure.Name)
	// 	}
	// }

	// productsResp, err := bfClient.Products.GetAllProducts(products.GetAllProductsParams{
	// 	Records: 10,
	// 	Order:   "DESC",
	// 	OrderBy: "created",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, product := range productsResp.Payload.Results {
	// 	log.Println("product", product.Name)
	// }

	// dob := time.Date(1993, time.July, 11, 0, 0, 0, 0, time.UTC)
	// updateProfileResp, err := bfClient.Profiles.UpdateProfile(profiles.UpdateProfileParams{
	// 	Request: &models.UpdateProfileRequest{
	// 		ID:        acct.Profile.ID,
	// 		AccountID: acct.ID,
	// 		// FirstName: "chancetest",
	// 		// LastName: "another tester",
	// 		Dob: &strfmt.DateTime{dob},
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("update profile response:", updateProfileResp.Payload.Results[0])

	// createAddrResp, err := bfClient.Addresses.CreateAddress(addresses.CreateAddressParams{
	// 	Request: &models.CreateAddressRequest{
	// 		ProfileID:    acct.Profile.ID,
	// 		AddressLine1: "555 nw kings st",
	// 		City:         "san francisco",
	// 		Postcode:     "94110",
	// 		Province:     "CA",
	// 		Country:      "US",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("create address response:", createAddrResp.Payload.Results[0])

	// updateAddrResp, err := bfClient.Addresses.UpdateAddress(addresses.UpdateAddressParams{
	// 	Request: &models.UpdateAddressRequest{
	// 		ID:           acct.Profile.Addresses[0].ID,
	// 		ProfileID:    acct.Profile.ID,
	// 		AddressLine1: "555 nw kings st",
	// 		City:         "san francisco",
	// 		Postcode:     "21000",
	// 		Province:     "CA",
	// 		Country:      "US",
	// 	},
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("update address response:", updateAddrResp.Payload.Results[0])

	// log.Println("creating aggregating subscription")
	// // TODO (remove). This can be automatically created when we create accounts
	// createSubResp, err := bfClient.Subscriptions.CreateAggregatingSubscription(subscriptions.CreateAggregatingSubscriptionParams{
	// 	Request: &models.CreateAggregatingSubscriptionRequest{
	// 		AccountID:      acct.ID,
	// 		Duration:       30,
	// 		DurationPeriod: "days",
	// 		ProductType:    "recurring",
	// 		Currency:       "USD",
	// 		AggregateAllSubscriptionsOnAccount: true,
	// 		Description:                        fmt.Sprintf("Aggregating subscription for %s", acct.Profile.CompanyName),
	// 	},
	// })
	// if err != nil {
	// 	gerr := err.(subscriptions.APIError).Response.(*subscriptions.CreateAggregatingSubscriptionDefault).Payload
	// 	log.Fatal(gerr)
	// }

	// log.Println("create aggregating subscription result:", createSubResp.Payload.Results[0])

	// log.Println("creating premium managed linux subscription")
	// _, err = bfClient.Subscriptions.CreateSubscription(subscriptions.CreateSubscriptionParams{
	// 	Request: &models.CreateSubscriptionRequest{
	// 		AccountID:       acct.ID,
	// 		Name:            "sub for enterprisecompany.com",
	// 		ProductRatePlan: "PRP-84EE41D7-09E9-4CAC-B0C1-C898D718", // Premium Managed Linux
	// 		State:           "AwaitingPayment",
	// 		Description:     "testing",
	// 	},
	// })
	// if err != nil {
	// 	apiErr, ok := err.(subscriptions.APIError)
	// 	if ok {
	// 		if gerr, ok := apiErr.Response.(*subscriptions.CreateSubscriptionDefault); ok {
	// 			log.Fatal(gerr.Payload)
	// 		}
	// 	}
	// 	log.Fatal(err)
	// }

	// log.Println("creating pro services subscription")
	// _, err = bfClient.Subscriptions.CreateSubscription(subscriptions.CreateSubscriptionParams{
	// 	Request: &models.CreateSubscriptionRequest{
	// 		AccountID:       acct.ID,
	// 		Name:            "sub for enterprisecompany.com",
	// 		ProductRatePlan: "PRP-F1E1C03E-4C41-47ED-BB9B-CABADC72", // pro services
	// 		State:           "AwaitingPayment",
	// 		Description:     "test",
	// 	},
	// })
	// if err != nil {
	// 	apiErr, ok := err.(subscriptions.APIError)
	// 	if ok {
	// 		if gerr, ok := apiErr.Response.(*subscriptions.CreateSubscriptionDefault); ok {
	// 			log.Fatal(gerr.Payload)
	// 		}
	// 	}
	// 	log.Fatal(err)
	// }
}

func String(s string) *string {
	return &s
}
