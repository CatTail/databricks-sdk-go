/*
 * Databricks REST API
 *
 * The Databricks REST API 2.0 supports of a variety of services.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package databricks

type TriggerType string

// List of TriggerType
const (
	PERIODIC_TriggerType TriggerType = "PERIODIC"
	ONE_TIME_TriggerType TriggerType = "ONE_TIME"
	RETRY_TriggerType TriggerType = "RETRY"
)
