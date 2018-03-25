/*
 * Screenwriter API
 *
 * Data API for Screenwriter: an online collaborative screenwriting software.
 *
 * API version: 0.0.1
 * Contact: logan@theinitiativepg.ocm
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Project struct {

	Id int64 `json:"id,omitempty"`

	Title string `json:"title,omitempty"`

	Owner int64 `json:"owner,omitempty"`

	CreatedDate int64 `json:"created_date,omitempty"`

	UpdatedDate int64 `json:"updated_date,omitempty"`

	// Project Status
	Status string `json:"status,omitempty"`
}
