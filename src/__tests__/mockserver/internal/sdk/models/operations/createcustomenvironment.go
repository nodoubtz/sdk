// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/models/components"
)

// CreateCustomEnvironmentType - Type of matcher. One of \"equals\", \"startsWith\", or \"endsWith\".
type CreateCustomEnvironmentType string

const (
	CreateCustomEnvironmentTypeEquals     CreateCustomEnvironmentType = "equals"
	CreateCustomEnvironmentTypeStartsWith CreateCustomEnvironmentType = "startsWith"
	CreateCustomEnvironmentTypeEndsWith   CreateCustomEnvironmentType = "endsWith"
)

func (e CreateCustomEnvironmentType) ToPointer() *CreateCustomEnvironmentType {
	return &e
}
func (e *CreateCustomEnvironmentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "equals":
		fallthrough
	case "startsWith":
		fallthrough
	case "endsWith":
		*e = CreateCustomEnvironmentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCustomEnvironmentType: %v", v)
	}
}

// BranchMatcher - How we want to determine a matching branch. This is optional.
type BranchMatcher struct {
	// Type of matcher. One of \"equals\", \"startsWith\", or \"endsWith\".
	Type CreateCustomEnvironmentType `json:"type"`
	// Git branch name or portion thereof.
	Pattern string `json:"pattern"`
}

func (o *BranchMatcher) GetType() CreateCustomEnvironmentType {
	if o == nil {
		return CreateCustomEnvironmentType("")
	}
	return o.Type
}

func (o *BranchMatcher) GetPattern() string {
	if o == nil {
		return ""
	}
	return o.Pattern
}

type CreateCustomEnvironmentRequestBody struct {
	// The slug of the custom environment to create.
	Slug *string `json:"slug,omitempty"`
	// Description of the custom environment. This is optional.
	Description *string `json:"description,omitempty"`
	// How we want to determine a matching branch. This is optional.
	BranchMatcher *BranchMatcher `json:"branchMatcher,omitempty"`
	// Where to copy environment variables from. This is optional.
	CopyEnvVarsFrom *string `json:"copyEnvVarsFrom,omitempty"`
}

func (o *CreateCustomEnvironmentRequestBody) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateCustomEnvironmentRequestBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateCustomEnvironmentRequestBody) GetBranchMatcher() *BranchMatcher {
	if o == nil {
		return nil
	}
	return o.BranchMatcher
}

func (o *CreateCustomEnvironmentRequestBody) GetCopyEnvVarsFrom() *string {
	if o == nil {
		return nil
	}
	return o.CopyEnvVarsFrom
}

type CreateCustomEnvironmentRequest struct {
	// The unique project identifier or the project name
	IDOrName string `pathParam:"style=simple,explode=false,name=idOrName"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                             `queryParam:"style=form,explode=true,name=slug"`
	RequestBody *CreateCustomEnvironmentRequestBody `request:"mediaType=application/json"`
}

func (o *CreateCustomEnvironmentRequest) GetIDOrName() string {
	if o == nil {
		return ""
	}
	return o.IDOrName
}

func (o *CreateCustomEnvironmentRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *CreateCustomEnvironmentRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateCustomEnvironmentRequest) GetRequestBody() *CreateCustomEnvironmentRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

// CreateCustomEnvironmentEnvironmentType - The type of environment (production, preview, or development)
type CreateCustomEnvironmentEnvironmentType string

const (
	CreateCustomEnvironmentEnvironmentTypeProduction  CreateCustomEnvironmentEnvironmentType = "production"
	CreateCustomEnvironmentEnvironmentTypePreview     CreateCustomEnvironmentEnvironmentType = "preview"
	CreateCustomEnvironmentEnvironmentTypeDevelopment CreateCustomEnvironmentEnvironmentType = "development"
)

func (e CreateCustomEnvironmentEnvironmentType) ToPointer() *CreateCustomEnvironmentEnvironmentType {
	return &e
}
func (e *CreateCustomEnvironmentEnvironmentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "production":
		fallthrough
	case "preview":
		fallthrough
	case "development":
		*e = CreateCustomEnvironmentEnvironmentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCustomEnvironmentEnvironmentType: %v", v)
	}
}

// CreateCustomEnvironmentEnvironmentResponseType - The type of matching to perform
type CreateCustomEnvironmentEnvironmentResponseType string

const (
	CreateCustomEnvironmentEnvironmentResponseTypeEndsWith   CreateCustomEnvironmentEnvironmentResponseType = "endsWith"
	CreateCustomEnvironmentEnvironmentResponseTypeStartsWith CreateCustomEnvironmentEnvironmentResponseType = "startsWith"
	CreateCustomEnvironmentEnvironmentResponseTypeEquals     CreateCustomEnvironmentEnvironmentResponseType = "equals"
)

func (e CreateCustomEnvironmentEnvironmentResponseType) ToPointer() *CreateCustomEnvironmentEnvironmentResponseType {
	return &e
}
func (e *CreateCustomEnvironmentEnvironmentResponseType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "endsWith":
		fallthrough
	case "startsWith":
		fallthrough
	case "equals":
		*e = CreateCustomEnvironmentEnvironmentResponseType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateCustomEnvironmentEnvironmentResponseType: %v", v)
	}
}

// CreateCustomEnvironmentBranchMatcher - Configuration for matching git branches to this environment
type CreateCustomEnvironmentBranchMatcher struct {
	// The type of matching to perform
	Type CreateCustomEnvironmentEnvironmentResponseType `json:"type"`
	// The pattern to match against branch names
	Pattern string `json:"pattern"`
}

func (o *CreateCustomEnvironmentBranchMatcher) GetType() CreateCustomEnvironmentEnvironmentResponseType {
	if o == nil {
		return CreateCustomEnvironmentEnvironmentResponseType("")
	}
	return o.Type
}

func (o *CreateCustomEnvironmentBranchMatcher) GetPattern() string {
	if o == nil {
		return ""
	}
	return o.Pattern
}

// CreateCustomEnvironmentVerification - A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
type CreateCustomEnvironmentVerification struct {
	Type   string `json:"type"`
	Domain string `json:"domain"`
	Value  string `json:"value"`
	Reason string `json:"reason"`
}

func (o *CreateCustomEnvironmentVerification) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *CreateCustomEnvironmentVerification) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *CreateCustomEnvironmentVerification) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

func (o *CreateCustomEnvironmentVerification) GetReason() string {
	if o == nil {
		return ""
	}
	return o.Reason
}

// CreateCustomEnvironmentDomains - List of domains associated with this environment
type CreateCustomEnvironmentDomains struct {
	Name                string   `json:"name"`
	ApexName            string   `json:"apexName"`
	ProjectID           string   `json:"projectId"`
	Redirect            *string  `json:"redirect,omitempty"`
	RedirectStatusCode  *float64 `json:"redirectStatusCode,omitempty"`
	GitBranch           *string  `json:"gitBranch,omitempty"`
	CustomEnvironmentID *string  `json:"customEnvironmentId,omitempty"`
	UpdatedAt           *float64 `json:"updatedAt,omitempty"`
	CreatedAt           *float64 `json:"createdAt,omitempty"`
	// `true` if the domain is verified for use with the project. If `false` it will not be used as an alias on this project until the challenge in `verification` is completed.
	Verified bool `json:"verified"`
	// A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
	Verification []CreateCustomEnvironmentVerification `json:"verification,omitempty"`
}

func (o *CreateCustomEnvironmentDomains) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateCustomEnvironmentDomains) GetApexName() string {
	if o == nil {
		return ""
	}
	return o.ApexName
}

func (o *CreateCustomEnvironmentDomains) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *CreateCustomEnvironmentDomains) GetRedirect() *string {
	if o == nil {
		return nil
	}
	return o.Redirect
}

func (o *CreateCustomEnvironmentDomains) GetRedirectStatusCode() *float64 {
	if o == nil {
		return nil
	}
	return o.RedirectStatusCode
}

func (o *CreateCustomEnvironmentDomains) GetGitBranch() *string {
	if o == nil {
		return nil
	}
	return o.GitBranch
}

func (o *CreateCustomEnvironmentDomains) GetCustomEnvironmentID() *string {
	if o == nil {
		return nil
	}
	return o.CustomEnvironmentID
}

func (o *CreateCustomEnvironmentDomains) GetUpdatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CreateCustomEnvironmentDomains) GetCreatedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CreateCustomEnvironmentDomains) GetVerified() bool {
	if o == nil {
		return false
	}
	return o.Verified
}

func (o *CreateCustomEnvironmentDomains) GetVerification() []CreateCustomEnvironmentVerification {
	if o == nil {
		return nil
	}
	return o.Verification
}

// CreateCustomEnvironmentResponseBody - Internal representation of a custom environment with all required properties
type CreateCustomEnvironmentResponseBody struct {
	// Unique identifier for the custom environment (format: env_*)
	ID string `json:"id"`
	// URL-friendly name of the environment
	Slug string `json:"slug"`
	// The type of environment (production, preview, or development)
	Type CreateCustomEnvironmentEnvironmentType `json:"type"`
	// Optional description of the environment's purpose
	Description *string `json:"description,omitempty"`
	// Configuration for matching git branches to this environment
	BranchMatcher *CreateCustomEnvironmentBranchMatcher `json:"branchMatcher,omitempty"`
	// List of domains associated with this environment
	Domains []CreateCustomEnvironmentDomains `json:"domains,omitempty"`
	// List of aliases for the current deployment
	CurrentDeploymentAliases []string `json:"currentDeploymentAliases,omitempty"`
	// Timestamp when the environment was created
	CreatedAt float64 `json:"createdAt"`
	// Timestamp when the environment was last updated
	UpdatedAt float64 `json:"updatedAt"`
}

func (o *CreateCustomEnvironmentResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateCustomEnvironmentResponseBody) GetSlug() string {
	if o == nil {
		return ""
	}
	return o.Slug
}

func (o *CreateCustomEnvironmentResponseBody) GetType() CreateCustomEnvironmentEnvironmentType {
	if o == nil {
		return CreateCustomEnvironmentEnvironmentType("")
	}
	return o.Type
}

func (o *CreateCustomEnvironmentResponseBody) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateCustomEnvironmentResponseBody) GetBranchMatcher() *CreateCustomEnvironmentBranchMatcher {
	if o == nil {
		return nil
	}
	return o.BranchMatcher
}

func (o *CreateCustomEnvironmentResponseBody) GetDomains() []CreateCustomEnvironmentDomains {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *CreateCustomEnvironmentResponseBody) GetCurrentDeploymentAliases() []string {
	if o == nil {
		return nil
	}
	return o.CurrentDeploymentAliases
}

func (o *CreateCustomEnvironmentResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *CreateCustomEnvironmentResponseBody) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

type CreateCustomEnvironmentResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *CreateCustomEnvironmentResponseBody
}

func (o *CreateCustomEnvironmentResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateCustomEnvironmentResponse) GetObject() *CreateCustomEnvironmentResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
