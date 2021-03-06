/*
 * Databricks REST API
 *
 * The Databricks REST API 2.0 supports of a variety of services.
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package databricks

type RunState struct {

	LifeCycleState *RunLifeCycleState `json:"life_cycle_state,omitempty"`

	ResultState *RunResultState `json:"result_state,omitempty"`

	StateMessage string `json:"state_message,omitempty"`
}
