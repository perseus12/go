package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func apiStreams(stream string) string {

	operations := map[string]string{
		"psc":        "persons-with-significant-control",
		"officers":   "officers",
		"roa":        "registered-office-address",
		"insolvency": "insolvency",
		"charges":    "charges",
		"exemptions": "exemptions",
		"registers":  "registers",
	}

	value, ok := operations[stream]
	if ok {
		return value
	} else {
		return "fail"
	}
}

func apiCall(stream, companyNo string) string {
	s := apiStreams(stream)
	url := "https://api.companieshouse.gov.uk/company/" + companyNo + "/" + s

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	chsapikey := (os.Getenv("CHS_API_KEY"))
	req.SetBasicAuth(chsapikey, "")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}
