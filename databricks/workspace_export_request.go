/*
 * Databricks REST API
 *
 * The Databricks REST API 2.0 supports of a variety of services.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package databricks

type WorkspaceExportRequest struct {

	Path string `json:"path"`

	Format *WorkspaceExportFormat `json:"format,omitempty"`

	DirectDownload bool `json:"direct_download,omitempty"`
}
