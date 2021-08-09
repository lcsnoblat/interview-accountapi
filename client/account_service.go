package account_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	account "form3/model"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/google/uuid"
)

func buildAndMakeAccountRequest(body *account.RequestData, method string, url string) *http.Response {

	client := &http.Client{}
	baseUrl := os.Getenv("BASE_URL")
	httpdeleteurl := baseUrl + url

	if body != nil {
		requestString, _ := json.Marshal(body)
		byteBody := bytes.NewBuffer(requestString)
		req, _ := http.NewRequest(method, httpdeleteurl, byteBody)
		resp, err := client.Do(req)
	} else {
		req, err := http.NewRequest(method, httpdeleteurl, nil)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()

	return resp
}

func configureHttpClient() {
	client := Client{}
	c := config.Schema{}
}

func Create(acc *account.RequestData) (string, error) {

	info := account.AccountInfo{
		Attributes:     accAttributes,
		ID:             uuid.New(),
		Type:           "accounts",
		OrganisationID: uuid.New(),
	}

	url := "v1/organisation/accounts"
	response := buildAndMakeAccountRequest(acc, "POST", url)

	if response.StatusCode == 409 {
		return "", fmt.Errorf("account already exists")
	}

	logResponse(response)

	return string(body), error
}

func FetchById(id string) (string, error) {

	url := "v1/organisation/accounts" + "/" + id
	response := buildAndMakeAccountRequest(nil, "GET", url)

	if response.StatusCode == 404 {
		return "", fmt.Errorf("account not found")
	}

	logResponse(response)
	return string(body), nil
}

func Delete(id string) (string, error) {

	url := "v1/organisation/accounts" + "/" + id
	response := buildAndMakeAccountRequest(nil, "DELETE", url)

	if response.StatusCode == 404 {
		return "", fmt.Errorf("account not found")
	}

	logResponse(response)

	return string(body), nil
}

func logResponse(resp *http.Response) {
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}
