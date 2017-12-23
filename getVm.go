package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type vm struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Locked      bool   `json:"locked"`
	ReadOnly    bool   `json:"readOnly"`
	VmRunState  string `json: "vmRunState"`
}

func main() {

	username := os.Getenv("OVM_USERNAME")
	password := os.Getenv("OVM_PASSWORD")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	baseUri := "https://10.64.78.100:7002/ovm/core/wsapi/rest/Vm"
	req, err := http.NewRequest("GET", baseUri, nil)
	req.SetBasicAuth(username, password)
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	}

	vm1 := []vm{}
	body, err := ioutil.ReadAll(resp.Body)
	jsonErr := json.Unmarshal(body, &vm1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(vm1)

}
