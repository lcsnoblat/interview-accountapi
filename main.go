package main

import (
	clientservice "form3/client"
	account "form3/model"
)

func main() {
	country := "GB"
	accAttributes := &account.Attributes{
		Country: &country,
		Name:    []string{"United Kingdom", "UK"},
	}
	accData := &account.AccountInfo{
		Attributes:     accAttributes,
		ID:             "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		Type:           "accounts",
		OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
	}
	requestData := &account.RequestData{
		RequestBody: accData,
	}

	clientservice.NewAccount(requestData)
}
