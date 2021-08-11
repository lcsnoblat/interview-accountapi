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

func buildAndMakeAccountRequest(body *account.RequestData, method string, url string) (*http.Response, error) {
	baseUrl := os.Getenv("BASE_URL")
	httpdeleteurl := baseUrl + url

	if body != nil {
		jsonBytes, _ := json.Marshal(body)
		request, _ := http.NewRequest(method, httpdeleteurl, bytes.NewBuffer(jsonBytes))
		request.Header.Add("Content-Type", "application/json")
		response, err := makeRequest(request)
		return response, err
	} else {
		request, _ := http.NewRequest(method, httpdeleteurl, nil)
		response, err := makeRequest(request)
		return response, err
	}
}

func makeRequest(request *http.Request) (*http.Response, error) {
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	return response, err
}

func Create(acc *account.Attributes) (string, error) {

	accData := &account.AccountInfo{
		Attributes:     acc,
		ID:             uuid.New().String(),
		Type:           "accounts",
		OrganisationID: uuid.New().String(),
	}

	requestData := &account.RequestData{
		RequestBody: accData,
	}

	url := "v1/organisation/accounts"
	response, err := buildAndMakeAccountRequest(requestData, "POST", url)

	if err != nil {
		return "", err
	}

	if response.StatusCode == 409 {
		return "", fmt.Errorf("account already exists")
	}

	logResponse(response)

	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

func FetchById(id string) (string, error) {

	url := "v1/organisation/accounts" + "/" + id
	response, err := buildAndMakeAccountRequest(nil, "GET", url)

	if err != nil {
		return "", err
	}

	if response.StatusCode == 404 {
		return "", fmt.Errorf("account not found")
	}

	logResponse(response)
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

func Delete(id string) (string, error) {

	url := "v1/organisation/accounts" + "/" + id
	response, err := buildAndMakeAccountRequest(nil, "DELETE", url)

	if err != nil {
		return "", err
	}

	if response.StatusCode == 404 {
		return "", fmt.Errorf("account not found")
	}

	logResponse(response)
	body, _ := ioutil.ReadAll(response.Body)
	return string(body), nil
}

func logResponse(response *http.Response) {
	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
}
