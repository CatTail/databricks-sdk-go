package examples

import (
	"github.com/cattail/databricks-sdk-go/databricks"
)

const jar = "dbfs:/mnt/libraries/library.jar"

func InstallLibrary(clusterId string) {
	_, err := client.LibraryApi.InstallLibrary(ctx, databricks.LibrariesInstallRequest{
		ClusterId: clusterId,
		Libraries: []databricks.Library{{Jar: jar}},
	})
	if err != nil {
		panic(err)
	}
}

func UninstallLibrary(clusterId string) {
	_, err := client.LibraryApi.UninstallLibrary(ctx, databricks.LibrariesUninstallRequest{
		ClusterId: clusterId,
		Libraries: []databricks.Library{{Jar: jar}},
	})
	if err != nil {
		panic(err)
	}
}
