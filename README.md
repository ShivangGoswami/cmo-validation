# Cost Management Data Validation 

This app is written in golang and require golang environment to be present in the system <br>
refer this link for instructions [Go Downloads](https://go.dev/dl/)

### NOTE: For access to prometheus data oc port foward is used
```sh
oc login -u **username** -p **password** **console-url**
oc port-forward -n openshift-monitoring pod/prometheus-k8s-0 9090:9090
```
Execute these command in a seperate terminal session because oc port forward is a blocking call

#### NOTE: Few constant values needs to be populated in code(main.go)

```golang
//populate the clusterID of the openshift cluster here
clusterID = ""
//populate the project name for which metrics needs to be fetched here
project = ""
//populate service account id
id = ""
//populate service account secret
secret = ""
```

* ClusterID: cluster id of the openshift cluster
* project: The openshift project for which data needs to be queried
* id & secret: client id and secret. This is obtained by creating service account on the redhat console

For more information on service accounts [Redhat Service Accounts](https://access.redhat.com/articles/7036194#step-4-update-your-api-integration-9)

### To Execute the code 
```sh
go run main.go
```

Sample Output:
```sh
(base) shivanggoswami@Shivangs-MacBook-Pro cmo-validation-github % go run main.go
Data from the console:

[
 {
  "date": "2024-06-13",
  "project": "costmanagement-metrics-operator",
  "usage": {
   "value": 0.006976286666667,
   "units": "Core-Hours"
  }
 },
 {
  "date": "2024-06-14",
  "project": "costmanagement-metrics-operator",
  "usage": {
   "value": 0.039031766388889,
   "units": "Core-Hours"
  }
 }
]

Data from the prometheus:

Prometheus Data for 2024-06-13 : 0.006948658973104022
Prometheus Data for 2024-06-14 : 0.03928629528140759
```
