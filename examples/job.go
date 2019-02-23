package examples

import (
	"github.com/cattail/databricks-sdk-go/databricks"
)

func CreateJob(clusterId string) databricks.JobsCreateResponse {
	res, _, err := client.JobApi.CreateJob(auth, databricks.JobsCreateRequest{
		Name:              "databricks-sdk-go job name",
		ExistingClusterId: clusterId,
	})
	if err != nil {
		panic(err)
	}
	return res
}

func GetJob(jobId int64) databricks.JobsGetResponse {
	res, _, err := client.JobApi.GetJob(auth, jobId)
	if err != nil {
		panic(err)
	}
	return res
}

func ResetJob(jobId int64, clusterId string) {
	_, err := client.JobApi.ResetJob(auth, databricks.JobsResetRequest{
		JobId: jobId,
		NewSettings: &databricks.JobSettings{
			Name:              "databricks-sdk-go new job name",
			ExistingClusterId: clusterId,
		},
	})
	if err != nil {
		panic(err)
	}
}

func DeleteJob(jobId int64) {
	_, err := client.JobApi.DeleteJob(auth, databricks.JobsDeleteRequest{JobId: jobId})
	if err != nil {
		panic(err)
	}
}
