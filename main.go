package main

import (
	create_account "form3/client"
	account "form3/model"
)

func main() {
	country := "Brazil"
	accAttributes := &account.Attributes{
		Country: &country,
	}
	accData := &account.Data{Attributes: accAttributes}

	create_account.NewAccount(accData)
}
