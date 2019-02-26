package examples

import (
	"fmt"
	"github.com/cattail/databricks-sdk-go/databricks"
	"time"
)

const (
	ClusterName  = "databricks-sdk-go example cluster name"
	SparkVersion = "4.2.x-scala2.11"
	NodeTypeId   = "r3.xlarge"
)

func CreateCluster() databricks.ClustersCreateResponse {
	res, _, err := client.ClusterApi.CreateCluster(ctx, databricks.NewCluster{
		ClusterName:  ClusterName,
		SparkVersion: SparkVersion,
		NodeTypeId:   NodeTypeId,
		NumWorkers:   1,
	})
	if err != nil {
		panic(err)
	}

	return res
}

func GetCluster(clusterId string) databricks.ClustersGetResponse {
	res, _, err := client.ClusterApi.GetCluster(ctx, clusterId)
	if err != nil {
		panic(err)
	}
	return res
}

func EditCluster(clusterId string) {
	WaitClusterState(clusterId, databricks.RUNNING_ClustersClusterState)

	_, err := client.ClusterApi.EditCluster(ctx, databricks.ClustersEditRequest{
		ClusterId:    clusterId,
		ClusterName:  ClusterName,
		SparkVersion: SparkVersion,
		NodeTypeId:   NodeTypeId,
		NumWorkers:   2,
	})
	if err != nil {
		panic(err)
	}
}

func DeleteCluster(clusterId string) {
	_, err := client.ClusterApi.PermanentDeleteCluster(ctx, databricks.ClustersPermanentDeleteRequest{
		ClusterId: clusterId,
	})
	if err != nil {
		panic(err)
	}
}

func WaitClusterState(clusterId string, state databricks.ClustersClusterState) {
	cluster := GetCluster(clusterId)
	for *cluster.State != state {
		cluster = GetCluster(clusterId)
		time.Sleep(5 * time.Second)
		fmt.Printf("Waiting cluster enter %s state from %s\n", state, *cluster.State)
	}
}
