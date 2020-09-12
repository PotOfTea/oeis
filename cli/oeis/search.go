package oeis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"oeis/pkg/consts"
	"time"
)

func SearchAPI(queryData string) error {
	baseURL, err := buildURL(consts.EndpointURL, queryData)
	if err != nil {
		return err
	}
	body, err := httpGet(baseURL)
	if err != nil {
		return err
	}

	o, err := validateJSON(body)
	if err != nil {
		return err
	}

	displayResults(o)

	return nil
}

func displayResults(query *OeisQuery) {
	resultCount := query.Count
	results := query.Results

	if resultCount > 0 && results != nil {
		fmt.Printf("Found %v results. Showing first five:\n", resultCount)
		for i := 0; i < 5; i++ {
			fmt.Printf("%v) %v \n", i+1, results[i].Name)
		}
	} else if resultCount > 0 && results == nil {
		fmt.Printf("Found %v results, too many to show. Please refine your search.\n", resultCount)
	} else {
		fmt.Println("Sorry, but the terms do not match anything in the table.")
	}
}

func validateJSON(body []byte) (*OeisQuery, error) {
	var o = new(OeisQuery)
	err := json.Unmarshal(body, &o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

func httpGet(baseURL string) ([]byte, error) {
	var netClient = &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := netClient.Get(baseURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func buildURL(inputURL string, queryData string) (string, error) {
	baseURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return "", err
	}
	baseURL.Path += "search"
	params := url.Values{}
	params.Add("q", queryData)
	params.Add("fmt", "json")
	baseURL.RawQuery = params.Encode()

	//fmt.Printf("Encoded URL is %q\n", baseURL.String())
	return baseURL.String(), nil
}
