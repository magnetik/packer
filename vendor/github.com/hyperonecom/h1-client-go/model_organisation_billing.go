/*
 * HyperOne API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// OrganisationBilling struct for OrganisationBilling
type OrganisationBilling struct {
	Currency string                     `json:"currency,omitempty"`
	Company  string                     `json:"company,omitempty"`
	Email    string                     `json:"email,omitempty"`
	Address  OrganisationBillingAddress `json:"address,omitempty"`
	Nip      string                     `json:"nip,omitempty"`
}
