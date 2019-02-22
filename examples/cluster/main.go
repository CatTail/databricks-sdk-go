package main

import (
	"encoding/json"
	"fmt"
	"os"

	ex "github.com/cattail/databricks-sdk-go/examples"
)

var client = ex.GetClient()
var auth = ex.GetAuth()
var secrets = ex.LoadSecrets()

func main() {
	clusterId := os.Args[1]
	res, _, err := client.ClusterApi.GetCluster(auth, clusterId)
	result, _ := json.Marshal(res)
	fmt.Println(err, string(result))
}
