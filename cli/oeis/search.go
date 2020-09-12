package oeis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"oeis/pkg/consts"
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

	//https://stackoverflow.com/questions/59759095/error-interface-conversion-interface-is-interface-not-mapstringinter

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

// func displayResults(data map[string]interface{}) {
// 	resultCount := data["count"].(float64)
// 	results := data["results"]

// 	if resultCount > 0 && results != nil {
// 		fmt.Printf("Found %v results. Showing first five:\n", resultCount)
// 		var filtredResults []string
// 		for _, item := range results.([]interface{}) {
// 			filtredResults = append(filtredResults, item.(map[string]interface{})["name"].(string))
// 		}

// 		for i := 0; i < 5; i++ {
// 			fmt.Printf("%v) %v \n", i+1, filtredResults[i])
// 		}

// 	} else if resultCount > 0 && results == nil {
// 		fmt.Printf("Found %v results, too many to show. Please refine your search.\n", resultCount)
// 	} else {
// 		fmt.Println("Sorry, but the terms do not match anything in the table.")
// 	}
// }

func validateJSON(body []byte) (*OeisQuery, error) {
	var o = new(OeisQuery)
	err := json.Unmarshal(body, &o)
	if err != nil {
		return nil, err
	}
	return o, nil
}

// func validateJSON(body []byte) (map[string]inte()rface{}, error) {

// 	//fmt.Println(string(body))
// 	var jsonKeys = []string{"greeting", "query", "count", "start", "results"}

// 	var objmap map[string]interface{}
// 	json.Unmarshal([]byte(body), &objmap)

// 	for _, jsonKey := range jsonKeys {
// 		_, ok := objmap[jsonKey]
// 		if !ok {
// 			//errorMsg := "Invalid JSON response should contain keys: " + strings.Join(jsonKeys, ",") + ". Response didn't contain key - " + jsonKey
// 			errorMsg := "Error: OEIS.org could not parse search query"
// 			return nil, errors.New(errorMsg)
// 		}
// 	}
// 	return objmap, nil

// }

func httpGet(baseURL *url.URL) ([]byte, error) {
	resp, err := http.Get(baseURL.String())

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

func buildURL(inputURL string, queryData string) (*url.URL, error) {
	baseURL, err := url.Parse(inputURL)
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return nil, err
	}
	baseURL.Path += "search"
	params := url.Values{}
	params.Add("q", queryData)
	params.Add("fmt", "json")
	baseURL.RawQuery = params.Encode()

	fmt.Printf("Encoded URL is %q\n", baseURL.String())
	return baseURL, err
}
