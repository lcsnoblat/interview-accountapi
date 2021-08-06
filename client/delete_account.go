package client_service

import (
	"fmt"
	account "form3/model"
	"io/ioutil"
	"net/http"
	"os"
)

func buildAndMakeAccountRequest(accountId string) *http.Response {

	client := &http.Client{}
	baseUrl := os.Getenv("BASE_URL")
	httpdeleteurl := baseUrl + "v1/organisation/accounts/" + accountId

	req, err := http.NewRequest("DELETE", httpdeleteurl, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	return resp
}

func DeleteAccount(acc *account.RequestData) (string, error) {

	response := buildAndMakeCreateAccountRequest(acc)

	if response.StatusCode == 404 {
		return "", fmt.Errorf("account not found")
	}

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return string(body), nil
}
