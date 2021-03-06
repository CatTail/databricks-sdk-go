/*
 * Databricks REST API
 *
 * The Databricks REST API 2.0 supports of a variety of services.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package databricks

type RunResultState string

// List of RunResultState
const (
	SUCCESS_RunResultState RunResultState = "SUCCESS"
	FAILED_RunResultState RunResultState = "FAILED"
	TIMEDOUT_RunResultState RunResultState = "TIMEDOUT"
	CANCELED_RunResultState RunResultState = "CANCELED"
)
