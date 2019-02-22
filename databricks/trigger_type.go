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
	PERIODIC_TRIGGER_TYPE TriggerType = "PERIODIC_TRIGGER_TYPE"
	ONE_TIME_TRIGGER_TYPE TriggerType = "ONE_TIME_TRIGGER_TYPE"
	RETRY_TRIGGER_TYPE TriggerType = "RETRY_TRIGGER_TYPE"
)