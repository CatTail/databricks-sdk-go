package main

import (
	"fmt"
	"github.com/cattail/databricks-sdk-go/databricks"
	"github.com/cattail/databricks-sdk-go/examples"
)

func main() {
	fmt.Println("CreateCluster")
	cluster := examples.CreateCluster()
	clusterId := cluster.ClusterId
	examples.JsonPrint(examples.GetCluster(clusterId))

	fmt.Println("")
	fmt.Println("EditCluster")
	examples.WaitClusterState(clusterId, databricks.RUNNING_CLUSTER_STATE)
	examples.EditCluster(clusterId)
	examples.JsonPrint(examples.GetCluster(clusterId))

	fmt.Println("")
	fmt.Println("DeleteCluster")
	examples.DeleteCluster(clusterId)
}
