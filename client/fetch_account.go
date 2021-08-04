package client_service

import (
	"fmt"
	account "form3/model"
	"io/ioutil"
	"net/http"
	"os"
)

func buildAndMakeFetchAccountRequest(accountId string) *http.Response {

	baseUrl := os.Getenv("BASE_URL")
	httpgeturl := baseUrl + "v1/organisation/accounts/" + accountId

	response, error := http.Get(httpgeturl)

	if error != nil {
		panic(error)
	}

	return response
}

func FetchAccount(acc *account.RequestData) (string, error) {

	response := buildAndMakeCreateAccountRequest(acc)

	if response.StatusCode == 409 {
		return "", fmt.Errorf("account already exists")
	}

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return string(body), nil
}
