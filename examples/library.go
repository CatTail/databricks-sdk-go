package examples

import (
	"github.com/cattail/databricks-sdk-go/databricks"
)

func InstallLibrary(clusterId string) {
	_, err := client.LibraryApi.InstallLibrary(auth, databricks.LibrariesInstallRequest{
		ClusterId: clusterId,
	})
	if err != nil {
		panic(err)
	}
}

func UninstallLibrary(clusterId string) {
	_, err := client.LibraryApi.UninstallLibrary(auth, databricks.LibrariesUninstallRequest{
		ClusterId: clusterId,
	})
	if err != nil {
		panic(err)
	}
}
