package examples

import (
	"github.com/cattail/databricks-sdk-go/databricks"
)

func CreateJob() databricks.JobsCreateResponse {
	res, _, err := client.JobApi.CreateJob(auth, databricks.JobsCreateRequest{})
	if err != nil {
		panic(err)
	}
	return res
}

func GetJob(jobId int32) databricks.JobsGetResponse {
	res, _, err := client.JobApi.GetJob(auth, jobId)
	if err != nil {
		panic(err)
	}
	return res
}

func ResetJob() {
	_, err := client.JobApi.ResetJob(auth, databricks.JobsResetRequest{})
	if err != nil {
		panic(err)
	}
}

func DeleteJob() {
	_, err := client.JobApi.DeleteJob(auth, databricks.JobsDeleteRequest{})
	if err != nil {
		panic(err)
	}
}
