package redhat

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Inner struct {
	Date    string `json:"date"`
	Project string `json:"project"`
	Usage   struct {
		Value float64 `json:"value"`
		Units string  `json:"units"`
	} `json:"usage"`
}

type cresponse struct {
	Data []struct {
		Date     string `json:"date"`
		Projects []struct {
			Project string  `json:"project"`
			Values  []Inner `json:"values"`
		} `json:"projects"`
	} `json:"data"`
}

func GetComputeURL(token, clusterID, project string) ([]Inner, error) {
	baseURL := "https://console.redhat.com/api/cost-management/v1/reports/openshift/compute/"

	queryParams := url.Values{}
	queryParams.Set("filter[resolution]", "daily")
	queryParams.Set("filter[cluster]", clusterID)
	queryParams.Set("group_by[project]", project)
	queryString := queryParams.Encode()

	fullURL := fmt.Sprintf("%s?%s", baseURL, queryString)

	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res cresponse
	err = json.NewDecoder(resp.Body).Decode(&res)
	if err != nil {
		return nil, err
	}
	var result []Inner
	for _, val1 := range res.Data {
		for _, val2 := range val1.Projects {
			for _, val3 := range val2.Values {
				if (val3 != Inner{}) {
					result = append(result, val3)
				}
			}
		}
	}
	return result, nil
}
