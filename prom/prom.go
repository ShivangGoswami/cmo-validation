package prom

import (
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
	"strconv"
)

type response struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				Namespace string `json:"namespace"`
				Node      string `json:"node"`
				Pod       string `json:"pod"`
			} `json:"metric"`
			Values [][]interface{} `json:"values"`
		} `json:"result"`
	} `json:"data"`
}

func GetPromMatrix(date, project string) (float64, error) {
	baseURL := "http://localhost:9090/api/v1/query_range"

	data := neturl.Values{}
	data.Add("query", `sum by (pod, namespace, node) (rate(container_cpu_usage_seconds_total{container!="", container!="POD", pod!="", namespace="`+project+`", node!=""}[5m]))`)
	data.Add("start", date+"T00:00:00.000Z")
	data.Add("end", date+"T23:59:59.000Z")
	data.Add("step", "1m")

	reqURL := fmt.Sprintf("%s?%s", baseURL, data.Encode())

	// Create a GET request
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return 0, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var respd response
	err = json.NewDecoder(resp.Body).Decode(&respd)
	if err != nil {
		return 0, err
	}

	result := float64(0)
	for _, val1 := range respd.Data.Result {
		for _, val2 := range val1.Values {
			temp, _ := strconv.ParseFloat(val2[1].(string), 64)
			result += (temp * 60)
		}
	}
	result = result / 3600
	return result, nil
}
