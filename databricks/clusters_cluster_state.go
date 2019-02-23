/*
 * Databricks REST API
 *
 * The Databricks REST API 2.0 supports of a variety of services.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package databricks

type ClustersClusterState string

// List of ClustersClusterState
const (
	PENDING_ClustersClusterState ClustersClusterState = "PENDING"
	RUNNING_ClustersClusterState ClustersClusterState = "RUNNING"
	RESTARTING_ClustersClusterState ClustersClusterState = "RESTARTING"
	RESIZING_ClustersClusterState ClustersClusterState = "RESIZING"
	TERMINATING_ClustersClusterState ClustersClusterState = "TERMINATING"
	TERMINATED_ClustersClusterState ClustersClusterState = "TERMINATED"
	ERROR__ClustersClusterState ClustersClusterState = "ERROR"
	UNKNOWN_ClustersClusterState ClustersClusterState = "UNKNOWN"
)
