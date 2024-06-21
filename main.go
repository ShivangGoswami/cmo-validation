package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ShivangGoswami/cmo-validation/auth"
	"github.com/ShivangGoswami/cmo-validation/prom"
	"github.com/ShivangGoswami/cmo-validation/redhat"
)

const (
	//populate the clusterID of the openshift cluster here
	clusterID = ""
	//populate the project name for which metrics needs to be fetched here
	project = ""
	//populate service account id
	id = ""
	//populate service account secret
	secret = ""
)

func main() {
	if clusterID == "" || project == "" || id == "" || secret == "" {
		log.Fatal("Please populate the const values in main.go")
	}
	token, err := auth.AuthToken(id, secret)
	if err != nil {
		log.Fatal(err)
	}
	dataRESTAPI, err := redhat.GetComputeURL(token.AccessToken, clusterID, project)
	if err != nil {
		log.Fatal(err)
	}
	result, err := json.MarshalIndent(dataRESTAPI, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data from the console:\n\n")
	fmt.Println(string(result))
	fmt.Printf("\nData from the prometheus:\n\n")
	for _, val := range dataRESTAPI {
		result, err := prom.GetPromMatrix(val.Date, project)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("Prometheus Data for", val.Date, ":", result)
	}
}
