/*
 * Databricks REST API
 *
 * The Databricks REST API 2.0 supports of a variety of services.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package databricks

type WorkspaceExportFormat string

// List of WorkspaceExportFormat
const (
	SOURCE_WorkspaceExportFormat WorkspaceExportFormat = "SOURCE"
	HTML_WorkspaceExportFormat WorkspaceExportFormat = "HTML"
	JUPYTER_WorkspaceExportFormat WorkspaceExportFormat = "JUPYTER"
	DBC_WorkspaceExportFormat WorkspaceExportFormat = "DBC"
)
