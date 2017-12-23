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

type id struct {
	Type  string `json:"type"`
	Value string `json:"value"`
	Uri   string `json:"uri"`
	Name  string `json:"name"`
}

type job struct {
	Id                           id     `json:"id"`
	Name                         string `json:"name"`
	Description                  string `json:"description"`
	Generation                   int    `json:"generation"`
	JobRunState                  string `json:"jobRunState"`
	JobSummaryState              string `json:"jobSummaryState"`
	ProgressMessage              string `json:"progressMessage"`
	LatestSummaryProgressMessage string `json:"latestSummaryProgressMessage"`
	StartTime                    int64  `json:"startTime"`
	EndTime                      int64  `json:"endTime"`
	User                         string `json:"user"`
	SummaryDone                  bool   `json:"summaryDone"`
}

func main() {

	username := os.Getenv("OVM_USERNAME")
	password := os.Getenv("OVM_PASSWORD")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	baseUri := "https://10.64.78.100:7002/ovm/core/wsapi/rest/Job/1514050814239"
	req, err := http.NewRequest("GET", baseUri, nil)
	req.SetBasicAuth(username, password)
	req.Header.Add("content-type", "application/json; charset=utf-8")
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)

	job1 := job{}

	jsonErr := json.Unmarshal(body, &job1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	fmt.Println(job1)

	//bodyString := string(body)

	//	fmt.Println(bodyString)
}
