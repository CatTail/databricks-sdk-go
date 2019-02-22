package examples

import (
	"github.com/cattail/databricks-sdk-go/databricks"
)

const (
	ClusterName  = "databricks-sdk-go generated cluster name"
	SparkVersion = "4.2.x-scala2.11"
	NodeTypeId   = "r3.xlarge"
)

func CreateCluster() databricks.ClustersCreateResponse {
	res, _, err := client.ClusterApi.CreateCluster(auth, databricks.ClustersCreateRequest{
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
	res, _, err := client.ClusterApi.GetCluster(auth, clusterId)
	if err != nil {
		panic(err)
	}
	return res
}

func EditCluster(clusterId string) {
	_, err := client.ClusterApi.EditCluster(auth, databricks.ClustersEditRequest{
		ClusterId: clusterId,
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
	_, err := client.ClusterApi.PermanentDeleteCluster(auth, databricks.ClustersPermanentDeleteRequest{
		ClusterId: clusterId,
	})
	if err != nil {
		panic(err)
	}
}
