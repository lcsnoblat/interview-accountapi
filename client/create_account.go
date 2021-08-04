package client_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	account "form3/model"
	"io/ioutil"
	"net/http"
	"os"
)

func buildAndMakeCreateAccountRequest(acc *account.RequestData) *http.Response {

	baseUrl := os.Getenv("BASE_URL")
	httpposturl := baseUrl + "v1/organisation/accounts"

	accountJson, err := json.Marshal(acc)

	response, error := http.Post(httpposturl, "application/json", bytes.NewBuffer(accountJson))

	if err != nil {
		panic(error)
	}

	return response
}

func NewAccount(acc *account.RequestData) (string, error) {

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
