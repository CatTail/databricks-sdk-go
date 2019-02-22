/*
 * Databricks REST API
 *
 * The Databricks REST API 2.0 supports of a variety of services.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package databricks

type NotebookTask struct {

	NotebookPath string `json:"notebook_path"`

	BaseParameters []ParamPair `json:"base_parameters,omitempty"`
}
