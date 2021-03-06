package oeis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"oeis/pkg/consts"
	"time"

	"github.com/briandowns/spinner"
)

func SearchAPI(queryData string) ([]string, int, error) {
	baseURL, err := buildURL(consts.EndpointURL, queryData)

	if err != nil {
		return nil, 0, err
	}
	body, err := httpGet(baseURL)
	if err != nil {
		return nil, 0, err
	}

	queryResponse, err := validateJSON(body)
	if err != nil {
		return nil, 0, err
	}

	return getData(queryResponse)
}

func getData(query *OeisQuery) ([]string, int, error) {
	resultCount := query.Count
	queryResults := query.Results
	var results []string

	if resultCount > 0 && queryResults != nil {
		var counter int
		if consts.SearchResults >= len(queryResults) {
			counter = len(queryResults)
		} else {
			counter = consts.SearchResults
		}

		for i := 0; i < counter; i++ {
			results = append(results, queryResults[i].Name)
		}
	} else if resultCount > 0 && results == nil {
		return nil, 0, fmt.Errorf("found %v results, too many to show. Please refine your search", resultCount)
	} else {
		return nil, 0, fmt.Errorf("sorry, but the terms do not match anything in the table")
	}
	return results, resultCount, nil
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
		Timeout: time.Second * 25,
	}
	s := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	s.Prefix = "Waiting for OEIS.ORG to respond: "
	s.Start()
	resp, err := netClient.Get(baseURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	defer s.Stop()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func buildURL(inputURL string, queryData string) (string, error) {
	baseURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}
	baseURL.Path += "search"
	params := url.Values{}
	params.Add("q", queryData)
	params.Add("fmt", "json")
	baseURL.RawQuery = params.Encode()

	return baseURL.String(), nil
}
