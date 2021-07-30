package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Simple interest calculation")

	httpposturl := "http://localhost:$PORT/v1/organisation/accounts/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc'"
	fmt.Println("HTTP JSON POST URL:", httpposturl)

	var jsonData = []byte(`{
		"name": "morpheus",
		"job": "leader"
	}`)
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

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
}
