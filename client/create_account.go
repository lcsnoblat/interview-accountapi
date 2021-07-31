package create_account

import (
	"bytes"
	"encoding/json"
	"fmt"
	account "form3/model"
	"io/ioutil"
	"net/http"
	"os"
)

func BuildRequest(acc *account.Data) (*http.Request) {

	httpposturl := "http://localhost:8080/v1/organisation/accounts"
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(acc)

	request, error := http.NewRequest("POST", httpposturl, payloadBuf)
	request.Header.Set("Signature", os.Getenv("SIGNATURE"))
	request.Header.Set("Authorization", os.Getenv("AUTHORIZATION_HEADER"))
	request.Header.Set("Content-Type", "application/vnd.api+json; charset=UTF-8")

	return request
}

func NewAccount(acc *account.Data) (string, error) {
	httpposturl := "http://localhost:8080/v1/organisation/accounts"
	
	client := &http.Client{}
	response, error := client.Do(request)

	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	return string(body), nil
}
