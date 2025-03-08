// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"mockserver/internal/sdk/models/components"
)

// DeliveryFormat - The delivery log format
type DeliveryFormat string

const (
	DeliveryFormatJSON   DeliveryFormat = "json"
	DeliveryFormatNdjson DeliveryFormat = "ndjson"
)

func (e DeliveryFormat) ToPointer() *DeliveryFormat {
	return &e
}
func (e *DeliveryFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "ndjson":
		*e = DeliveryFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for DeliveryFormat: %v", v)
	}
}

type Sources string

const (
	SourcesStatic   Sources = "static"
	SourcesLambda   Sources = "lambda"
	SourcesBuild    Sources = "build"
	SourcesEdge     Sources = "edge"
	SourcesExternal Sources = "external"
	SourcesFirewall Sources = "firewall"
)

func (e Sources) ToPointer() *Sources {
	return &e
}
func (e *Sources) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "static":
		fallthrough
	case "lambda":
		fallthrough
	case "build":
		fallthrough
	case "edge":
		fallthrough
	case "external":
		fallthrough
	case "firewall":
		*e = Sources(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Sources: %v", v)
	}
}

type Environments string

const (
	EnvironmentsPreview    Environments = "preview"
	EnvironmentsProduction Environments = "production"
)

func (e Environments) ToPointer() *Environments {
	return &e
}
func (e *Environments) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "preview":
		fallthrough
	case "production":
		*e = Environments(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Environments: %v", v)
	}
}

type CreateConfigurableLogDrainRequestBody struct {
	// The delivery log format
	DeliveryFormat DeliveryFormat `json:"deliveryFormat"`
	// The log drain url
	URL string `json:"url"`
	// Headers to be sent together with the request
	Headers      map[string]string `json:"headers,omitempty"`
	ProjectIds   []string          `json:"projectIds,omitempty"`
	Sources      []Sources         `json:"sources"`
	Environments []Environments    `json:"environments,omitempty"`
	// Custom secret of log drain
	Secret *string `json:"secret,omitempty"`
	// The sampling rate for this log drain. It should be a percentage rate between 0 and 100. With max 2 decimal points
	SamplingRate *float64 `json:"samplingRate,omitempty"`
	// The custom name of this log drain.
	Name *string `json:"name,omitempty"`
}

func (o *CreateConfigurableLogDrainRequestBody) GetDeliveryFormat() DeliveryFormat {
	if o == nil {
		return DeliveryFormat("")
	}
	return o.DeliveryFormat
}

func (o *CreateConfigurableLogDrainRequestBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *CreateConfigurableLogDrainRequestBody) GetHeaders() map[string]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateConfigurableLogDrainRequestBody) GetProjectIds() []string {
	if o == nil {
		return nil
	}
	return o.ProjectIds
}

func (o *CreateConfigurableLogDrainRequestBody) GetSources() []Sources {
	if o == nil {
		return []Sources{}
	}
	return o.Sources
}

func (o *CreateConfigurableLogDrainRequestBody) GetEnvironments() []Environments {
	if o == nil {
		return nil
	}
	return o.Environments
}

func (o *CreateConfigurableLogDrainRequestBody) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *CreateConfigurableLogDrainRequestBody) GetSamplingRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SamplingRate
}

func (o *CreateConfigurableLogDrainRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

type CreateConfigurableLogDrainRequest struct {
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug        *string                               `queryParam:"style=form,explode=true,name=slug"`
	RequestBody CreateConfigurableLogDrainRequestBody `request:"mediaType=application/json"`
}

func (o *CreateConfigurableLogDrainRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *CreateConfigurableLogDrainRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CreateConfigurableLogDrainRequest) GetRequestBody() CreateConfigurableLogDrainRequestBody {
	if o == nil {
		return CreateConfigurableLogDrainRequestBody{}
	}
	return o.RequestBody
}

type CreateConfigurableLogDrainSources string

const (
	CreateConfigurableLogDrainSourcesBuild    CreateConfigurableLogDrainSources = "build"
	CreateConfigurableLogDrainSourcesEdge     CreateConfigurableLogDrainSources = "edge"
	CreateConfigurableLogDrainSourcesLambda   CreateConfigurableLogDrainSources = "lambda"
	CreateConfigurableLogDrainSourcesStatic   CreateConfigurableLogDrainSources = "static"
	CreateConfigurableLogDrainSourcesExternal CreateConfigurableLogDrainSources = "external"
	CreateConfigurableLogDrainSourcesFirewall CreateConfigurableLogDrainSources = "firewall"
)

func (e CreateConfigurableLogDrainSources) ToPointer() *CreateConfigurableLogDrainSources {
	return &e
}
func (e *CreateConfigurableLogDrainSources) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "build":
		fallthrough
	case "edge":
		fallthrough
	case "lambda":
		fallthrough
	case "static":
		fallthrough
	case "external":
		fallthrough
	case "firewall":
		*e = CreateConfigurableLogDrainSources(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrainSources: %v", v)
	}
}

type CreateConfigurableLogDrainEnvironments string

const (
	CreateConfigurableLogDrainEnvironmentsProduction CreateConfigurableLogDrainEnvironments = "production"
	CreateConfigurableLogDrainEnvironmentsPreview    CreateConfigurableLogDrainEnvironments = "preview"
)

func (e CreateConfigurableLogDrainEnvironments) ToPointer() *CreateConfigurableLogDrainEnvironments {
	return &e
}
func (e *CreateConfigurableLogDrainEnvironments) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "production":
		fallthrough
	case "preview":
		*e = CreateConfigurableLogDrainEnvironments(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrainEnvironments: %v", v)
	}
}

type CreateConfigurableLogDrainStatus string

const (
	CreateConfigurableLogDrainStatusEnabled  CreateConfigurableLogDrainStatus = "enabled"
	CreateConfigurableLogDrainStatusDisabled CreateConfigurableLogDrainStatus = "disabled"
	CreateConfigurableLogDrainStatusErrored  CreateConfigurableLogDrainStatus = "errored"
)

func (e CreateConfigurableLogDrainStatus) ToPointer() *CreateConfigurableLogDrainStatus {
	return &e
}
func (e *CreateConfigurableLogDrainStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "enabled":
		fallthrough
	case "disabled":
		fallthrough
	case "errored":
		*e = CreateConfigurableLogDrainStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrainStatus: %v", v)
	}
}

type CreateConfigurableLogDrainDisabledReason string

const (
	CreateConfigurableLogDrainDisabledReasonDisabledByOwner      CreateConfigurableLogDrainDisabledReason = "disabled-by-owner"
	CreateConfigurableLogDrainDisabledReasonFeatureNotAvailable  CreateConfigurableLogDrainDisabledReason = "feature-not-available"
	CreateConfigurableLogDrainDisabledReasonAccountPlanDowngrade CreateConfigurableLogDrainDisabledReason = "account-plan-downgrade"
	CreateConfigurableLogDrainDisabledReasonDisabledByAdmin      CreateConfigurableLogDrainDisabledReason = "disabled-by-admin"
)

func (e CreateConfigurableLogDrainDisabledReason) ToPointer() *CreateConfigurableLogDrainDisabledReason {
	return &e
}
func (e *CreateConfigurableLogDrainDisabledReason) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "disabled-by-owner":
		fallthrough
	case "feature-not-available":
		fallthrough
	case "account-plan-downgrade":
		fallthrough
	case "disabled-by-admin":
		*e = CreateConfigurableLogDrainDisabledReason(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrainDisabledReason: %v", v)
	}
}

type CreateConfigurableLogDrainCreatedFrom string

const (
	CreateConfigurableLogDrainCreatedFromSelfServed  CreateConfigurableLogDrainCreatedFrom = "self-served"
	CreateConfigurableLogDrainCreatedFromIntegration CreateConfigurableLogDrainCreatedFrom = "integration"
)

func (e CreateConfigurableLogDrainCreatedFrom) ToPointer() *CreateConfigurableLogDrainCreatedFrom {
	return &e
}
func (e *CreateConfigurableLogDrainCreatedFrom) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "self-served":
		fallthrough
	case "integration":
		*e = CreateConfigurableLogDrainCreatedFrom(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrainCreatedFrom: %v", v)
	}
}

type CreateConfigurableLogDrainDeliveryFormat string

const (
	CreateConfigurableLogDrainDeliveryFormatJSON   CreateConfigurableLogDrainDeliveryFormat = "json"
	CreateConfigurableLogDrainDeliveryFormatNdjson CreateConfigurableLogDrainDeliveryFormat = "ndjson"
	CreateConfigurableLogDrainDeliveryFormatSyslog CreateConfigurableLogDrainDeliveryFormat = "syslog"
)

func (e CreateConfigurableLogDrainDeliveryFormat) ToPointer() *CreateConfigurableLogDrainDeliveryFormat {
	return &e
}
func (e *CreateConfigurableLogDrainDeliveryFormat) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		fallthrough
	case "ndjson":
		fallthrough
	case "syslog":
		*e = CreateConfigurableLogDrainDeliveryFormat(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateConfigurableLogDrainDeliveryFormat: %v", v)
	}
}

type CreateConfigurableLogDrainResponseBody struct {
	// The secret to validate the log-drain payload
	Secret              *string                                   `json:"secret,omitempty"`
	ClientID            *string                                   `json:"clientId,omitempty"`
	ConfigurationID     *string                                   `json:"configurationId,omitempty"`
	Sources             []CreateConfigurableLogDrainSources       `json:"sources,omitempty"`
	Environments        []CreateConfigurableLogDrainEnvironments  `json:"environments"`
	Status              *CreateConfigurableLogDrainStatus         `json:"status,omitempty"`
	DisabledAt          *float64                                  `json:"disabledAt,omitempty"`
	DisabledReason      *CreateConfigurableLogDrainDisabledReason `json:"disabledReason,omitempty"`
	DisabledBy          *string                                   `json:"disabledBy,omitempty"`
	FirstErrorTimestamp *float64                                  `json:"firstErrorTimestamp,omitempty"`
	SamplingRate        *float64                                  `json:"samplingRate,omitempty"`
	HideIPAddresses     *bool                                     `json:"hideIpAddresses,omitempty"`
	ID                  string                                    `json:"id"`
	CreatedAt           float64                                   `json:"createdAt"`
	DeletedAt           *float64                                  `json:"deletedAt"`
	UpdatedAt           float64                                   `json:"updatedAt"`
	URL                 string                                    `json:"url"`
	Headers             map[string]string                         `json:"headers,omitempty"`
	ProjectIds          []string                                  `json:"projectIds,omitempty"`
	Name                string                                    `json:"name"`
	TeamID              *string                                   `json:"teamId,omitempty"`
	OwnerID             string                                    `json:"ownerId"`
	CreatedFrom         *CreateConfigurableLogDrainCreatedFrom    `json:"createdFrom,omitempty"`
	DeliveryFormat      CreateConfigurableLogDrainDeliveryFormat  `json:"deliveryFormat"`
}

func (o *CreateConfigurableLogDrainResponseBody) GetSecret() *string {
	if o == nil {
		return nil
	}
	return o.Secret
}

func (o *CreateConfigurableLogDrainResponseBody) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *CreateConfigurableLogDrainResponseBody) GetConfigurationID() *string {
	if o == nil {
		return nil
	}
	return o.ConfigurationID
}

func (o *CreateConfigurableLogDrainResponseBody) GetSources() []CreateConfigurableLogDrainSources {
	if o == nil {
		return nil
	}
	return o.Sources
}

func (o *CreateConfigurableLogDrainResponseBody) GetEnvironments() []CreateConfigurableLogDrainEnvironments {
	if o == nil {
		return []CreateConfigurableLogDrainEnvironments{}
	}
	return o.Environments
}

func (o *CreateConfigurableLogDrainResponseBody) GetStatus() *CreateConfigurableLogDrainStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *CreateConfigurableLogDrainResponseBody) GetDisabledAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DisabledAt
}

func (o *CreateConfigurableLogDrainResponseBody) GetDisabledReason() *CreateConfigurableLogDrainDisabledReason {
	if o == nil {
		return nil
	}
	return o.DisabledReason
}

func (o *CreateConfigurableLogDrainResponseBody) GetDisabledBy() *string {
	if o == nil {
		return nil
	}
	return o.DisabledBy
}

func (o *CreateConfigurableLogDrainResponseBody) GetFirstErrorTimestamp() *float64 {
	if o == nil {
		return nil
	}
	return o.FirstErrorTimestamp
}

func (o *CreateConfigurableLogDrainResponseBody) GetSamplingRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SamplingRate
}

func (o *CreateConfigurableLogDrainResponseBody) GetHideIPAddresses() *bool {
	if o == nil {
		return nil
	}
	return o.HideIPAddresses
}

func (o *CreateConfigurableLogDrainResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *CreateConfigurableLogDrainResponseBody) GetCreatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.CreatedAt
}

func (o *CreateConfigurableLogDrainResponseBody) GetDeletedAt() *float64 {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *CreateConfigurableLogDrainResponseBody) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

func (o *CreateConfigurableLogDrainResponseBody) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

func (o *CreateConfigurableLogDrainResponseBody) GetHeaders() map[string]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *CreateConfigurableLogDrainResponseBody) GetProjectIds() []string {
	if o == nil {
		return nil
	}
	return o.ProjectIds
}

func (o *CreateConfigurableLogDrainResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateConfigurableLogDrainResponseBody) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *CreateConfigurableLogDrainResponseBody) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *CreateConfigurableLogDrainResponseBody) GetCreatedFrom() *CreateConfigurableLogDrainCreatedFrom {
	if o == nil {
		return nil
	}
	return o.CreatedFrom
}

func (o *CreateConfigurableLogDrainResponseBody) GetDeliveryFormat() CreateConfigurableLogDrainDeliveryFormat {
	if o == nil {
		return CreateConfigurableLogDrainDeliveryFormat("")
	}
	return o.DeliveryFormat
}

type CreateConfigurableLogDrainResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	Object   *CreateConfigurableLogDrainResponseBody
}

func (o *CreateConfigurableLogDrainResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *CreateConfigurableLogDrainResponse) GetObject() *CreateConfigurableLogDrainResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
