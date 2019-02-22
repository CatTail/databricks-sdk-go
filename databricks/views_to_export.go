/*
 * Databricks REST API
 *
 * The Databricks REST API 2.0 supports of a variety of services.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package databricks

type ViewsToExport string

// List of ViewsToExport
const (
	CODE_EXPORT ViewsToExport = "CODE_EXPORT"
	DASHBOARDS_EXPORT ViewsToExport = "DASHBOARDS_EXPORT"
	ALL_EXPORT ViewsToExport = "ALL_EXPORT"
)
