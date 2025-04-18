// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"mockserver/internal/sdk/models/components"
)

type PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigRequestBody struct {
	Data map[string]any `json:"data"`
}

func (o *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigRequestBody) GetData() map[string]any {
	if o == nil {
		return map[string]any{}
	}
	return o.Data
}

type PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigRequest struct {
	IntegrationConfigurationID string                                                                                               `pathParam:"style=simple,explode=false,name=integrationConfigurationId"`
	ResourceID                 string                                                                                               `pathParam:"style=simple,explode=false,name=resourceId"`
	RequestBody                *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigRequestBody `request:"mediaType=application/json"`
}

func (o *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigRequest) GetIntegrationConfigurationID() string {
	if o == nil {
		return ""
	}
	return o.IntegrationConfigurationID
}

func (o *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigRequest) GetResourceID() string {
	if o == nil {
		return ""
	}
	return o.ResourceID
}

func (o *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigRequest) GetRequestBody() *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigRequestBody {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

// PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponseBody - The Edge Config was updated
type PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponseBody struct {
	Items     map[string]*components.EdgeConfigItemValue `json:"items"`
	UpdatedAt float64                                    `json:"updatedAt"`
	Digest    string                                     `json:"digest"`
}

func (o *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponseBody) GetItems() map[string]*components.EdgeConfigItemValue {
	if o == nil {
		return map[string]*components.EdgeConfigItemValue{}
	}
	return o.Items
}

func (o *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponseBody) GetUpdatedAt() float64 {
	if o == nil {
		return 0.0
	}
	return o.UpdatedAt
}

func (o *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponseBody) GetDigest() string {
	if o == nil {
		return ""
	}
	return o.Digest
}

type PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The Edge Config was updated
	Object *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponseBody
}

func (o *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponse) GetObject() *PutV1InstallationsIntegrationConfigurationIDResourcesResourceIDExperimentationEdgeConfigResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
