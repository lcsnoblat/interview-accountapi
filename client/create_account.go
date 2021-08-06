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

func buildAndMakeCreateAccountRequest(body *account.RequestData, method string, url string) *http.Response {

	client := &http.Client{}
	baseUrl := os.Getenv("BASE_URL")
	httpdeleteurl := baseUrl + url
	requestString, err := json.Marshal(body)
	byteBody := bytes.NewBuffer(requestString)

	req, err := http.NewRequest(method, httpdeleteurl, byteBody)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer resp.Body.Close()

	return resp
}

func Create(acc *account.RequestData) (string, error) {

	url := "v1/organisation/accounts"
	response := buildAndMakeCreateAccountRequest(acc, "POST", url)

	if response.StatusCode == 409 {
		return "", fmt.Errorf("account already exists")
	}

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return string(body), nil
}

func FetchById(id string) (string, error) {

	url := "v1/organisation/accounts" + "/" + id
	response := buildAndMakeCreateAccountRequest(nil, "GET", url)

	if response.StatusCode == 409 {
		return "", fmt.Errorf("account already exists")
	}

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return string(body), nil
}

func Delete(acc *account.RequestData) (string, error) {

	url := "v1/organisation/accounts"
	response := buildAndMakeCreateAccountRequest(acc, "DELETE")

	if response.StatusCode == 404 {
		return "", fmt.Errorf("account not found")
	}

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return string(body), nil
}
