package main

import (
	"bufio"
	"fmt"
	"github.com/cattail/databricks-sdk-go/examples"
	"os"
)

func main() {
	prompt("CreateCluster")
	cluster := examples.CreateCluster()
	clusterId := cluster.ClusterId

	prompt("GetCluster")
	examples.JsonPrint(examples.GetCluster(clusterId))

	prompt("EditCluster")
	examples.EditCluster(clusterId)
	examples.JsonPrint(examples.GetCluster(clusterId))

	prompt("InstallLibrary")
	examples.InstallLibrary(clusterId)

	prompt("UninstallLibrary")
	examples.UninstallLibrary(clusterId)

	prompt("CreateJob")
	job := examples.CreateJob(clusterId)
	jobId := job.JobId

	prompt("GetJob")
	examples.JsonPrint(examples.GetJob(jobId))

	prompt("ResetJob")
	examples.ResetJob(jobId, clusterId)

	prompt("DeleteJob")
	examples.DeleteJob(jobId)

	prompt("DeleteCluster")
	examples.DeleteCluster(clusterId)
}

func prompt(title string) {
	fmt.Println()
	fmt.Println(title)
	pause()
}

func pause() {
	fmt.Print("Press 'Enter' to continue...")
	_, _ = bufio.NewReader(os.Stdin).ReadBytes('\n')
}
