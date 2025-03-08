// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"errors"
	"fmt"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/utils"
)

type GetBypassIPRequest struct {
	ProjectID string   `queryParam:"style=form,explode=true,name=projectId"`
	Limit     *float64 `queryParam:"style=form,explode=true,name=limit"`
	// Filter by source IP
	SourceIP *string `queryParam:"style=form,explode=true,name=sourceIp"`
	// Filter by domain
	Domain *string `queryParam:"style=form,explode=true,name=domain"`
	// Filter by project scoped rules
	ProjectScope *bool `queryParam:"style=form,explode=true,name=projectScope"`
	// Used for pagination. Retrieves results after the provided id
	Offset *string `queryParam:"style=form,explode=true,name=offset"`
	// The Team identifier to perform the request on behalf of.
	TeamID *string `queryParam:"style=form,explode=true,name=teamId"`
	// The Team slug to perform the request on behalf of.
	Slug *string `queryParam:"style=form,explode=true,name=slug"`
}

func (o *GetBypassIPRequest) GetProjectID() string {
	if o == nil {
		return ""
	}
	return o.ProjectID
}

func (o *GetBypassIPRequest) GetLimit() *float64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetBypassIPRequest) GetSourceIP() *string {
	if o == nil {
		return nil
	}
	return o.SourceIP
}

func (o *GetBypassIPRequest) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *GetBypassIPRequest) GetProjectScope() *bool {
	if o == nil {
		return nil
	}
	return o.ProjectScope
}

func (o *GetBypassIPRequest) GetOffset() *string {
	if o == nil {
		return nil
	}
	return o.Offset
}

func (o *GetBypassIPRequest) GetTeamID() *string {
	if o == nil {
		return nil
	}
	return o.TeamID
}

func (o *GetBypassIPRequest) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

type GetBypassIPResponseBodyAction string

const (
	GetBypassIPResponseBodyActionBlock  GetBypassIPResponseBodyAction = "block"
	GetBypassIPResponseBodyActionBypass GetBypassIPResponseBodyAction = "bypass"
)

func (e GetBypassIPResponseBodyAction) ToPointer() *GetBypassIPResponseBodyAction {
	return &e
}
func (e *GetBypassIPResponseBodyAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "bypass":
		*e = GetBypassIPResponseBodyAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetBypassIPResponseBodyAction: %v", v)
	}
}

type GetBypassIPResponseBodyResult struct {
	OwnerID       string                         `json:"OwnerId"`
	ID            string                         `json:"Id"`
	Domain        string                         `json:"Domain"`
	IP            string                         `json:"Ip"`
	Action        *GetBypassIPResponseBodyAction `json:"Action,omitempty"`
	ProjectID     *string                        `json:"ProjectId,omitempty"`
	IsProjectRule *bool                          `json:"IsProjectRule,omitempty"`
	Note          *string                        `json:"Note,omitempty"`
	CreatedAt     string                         `json:"CreatedAt"`
	ActorID       *string                        `json:"ActorId,omitempty"`
	UpdatedAt     string                         `json:"UpdatedAt"`
	UpdatedAtHour string                         `json:"UpdatedAtHour"`
	DeletedAt     *string                        `json:"DeletedAt,omitempty"`
	ExpiresAt     *float64                       `json:"ExpiresAt,omitempty"`
}

func (o *GetBypassIPResponseBodyResult) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *GetBypassIPResponseBodyResult) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetBypassIPResponseBodyResult) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *GetBypassIPResponseBodyResult) GetIP() string {
	if o == nil {
		return ""
	}
	return o.IP
}

func (o *GetBypassIPResponseBodyResult) GetAction() *GetBypassIPResponseBodyAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *GetBypassIPResponseBodyResult) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *GetBypassIPResponseBodyResult) GetIsProjectRule() *bool {
	if o == nil {
		return nil
	}
	return o.IsProjectRule
}

func (o *GetBypassIPResponseBodyResult) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *GetBypassIPResponseBodyResult) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *GetBypassIPResponseBodyResult) GetActorID() *string {
	if o == nil {
		return nil
	}
	return o.ActorID
}

func (o *GetBypassIPResponseBodyResult) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *GetBypassIPResponseBodyResult) GetUpdatedAtHour() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAtHour
}

func (o *GetBypassIPResponseBodyResult) GetDeletedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *GetBypassIPResponseBodyResult) GetExpiresAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

type GetBypassIPResponseBodyPagination struct {
	OwnerID string `json:"OwnerId"`
	ID      string `json:"Id"`
}

func (o *GetBypassIPResponseBodyPagination) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *GetBypassIPResponseBodyPagination) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type GetBypassIPResponseBody2 struct {
	Result     []GetBypassIPResponseBodyResult    `json:"result,omitempty"`
	Pagination *GetBypassIPResponseBodyPagination `json:"pagination,omitempty"`
}

func (o *GetBypassIPResponseBody2) GetResult() []GetBypassIPResponseBodyResult {
	if o == nil {
		return nil
	}
	return o.Result
}

func (o *GetBypassIPResponseBody2) GetPagination() *GetBypassIPResponseBodyPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

type ResponseBodyAction string

const (
	ResponseBodyActionBlock  ResponseBodyAction = "block"
	ResponseBodyActionBypass ResponseBodyAction = "bypass"
)

func (e ResponseBodyAction) ToPointer() *ResponseBodyAction {
	return &e
}
func (e *ResponseBodyAction) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "block":
		fallthrough
	case "bypass":
		*e = ResponseBodyAction(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ResponseBodyAction: %v", v)
	}
}

type ResponseBodyResult struct {
	OwnerID       string              `json:"OwnerId"`
	ID            string              `json:"Id"`
	Domain        string              `json:"Domain"`
	IP            string              `json:"Ip"`
	Action        *ResponseBodyAction `json:"Action,omitempty"`
	ProjectID     *string             `json:"ProjectId,omitempty"`
	IsProjectRule *bool               `json:"IsProjectRule,omitempty"`
	Note          *string             `json:"Note,omitempty"`
	CreatedAt     string              `json:"CreatedAt"`
	ActorID       *string             `json:"ActorId,omitempty"`
	UpdatedAt     string              `json:"UpdatedAt"`
	UpdatedAtHour string              `json:"UpdatedAtHour"`
	DeletedAt     *string             `json:"DeletedAt,omitempty"`
	ExpiresAt     *float64            `json:"ExpiresAt,omitempty"`
}

func (o *ResponseBodyResult) GetOwnerID() string {
	if o == nil {
		return ""
	}
	return o.OwnerID
}

func (o *ResponseBodyResult) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *ResponseBodyResult) GetDomain() string {
	if o == nil {
		return ""
	}
	return o.Domain
}

func (o *ResponseBodyResult) GetIP() string {
	if o == nil {
		return ""
	}
	return o.IP
}

func (o *ResponseBodyResult) GetAction() *ResponseBodyAction {
	if o == nil {
		return nil
	}
	return o.Action
}

func (o *ResponseBodyResult) GetProjectID() *string {
	if o == nil {
		return nil
	}
	return o.ProjectID
}

func (o *ResponseBodyResult) GetIsProjectRule() *bool {
	if o == nil {
		return nil
	}
	return o.IsProjectRule
}

func (o *ResponseBodyResult) GetNote() *string {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *ResponseBodyResult) GetCreatedAt() string {
	if o == nil {
		return ""
	}
	return o.CreatedAt
}

func (o *ResponseBodyResult) GetActorID() *string {
	if o == nil {
		return nil
	}
	return o.ActorID
}

func (o *ResponseBodyResult) GetUpdatedAt() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAt
}

func (o *ResponseBodyResult) GetUpdatedAtHour() string {
	if o == nil {
		return ""
	}
	return o.UpdatedAtHour
}

func (o *ResponseBodyResult) GetDeletedAt() *string {
	if o == nil {
		return nil
	}
	return o.DeletedAt
}

func (o *ResponseBodyResult) GetExpiresAt() *float64 {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

type GetBypassIPResponseBody1 struct {
	Result     []ResponseBodyResult `json:"result"`
	Pagination any                  `json:"pagination"`
}

func (o *GetBypassIPResponseBody1) GetResult() []ResponseBodyResult {
	if o == nil {
		return []ResponseBodyResult{}
	}
	return o.Result
}

func (o *GetBypassIPResponseBody1) GetPagination() any {
	if o == nil {
		return nil
	}
	return o.Pagination
}

type GetBypassIPResponseBodyType string

const (
	GetBypassIPResponseBodyTypeGetBypassIPResponseBody1 GetBypassIPResponseBodyType = "getBypassIp_responseBody_1"
	GetBypassIPResponseBodyTypeGetBypassIPResponseBody2 GetBypassIPResponseBodyType = "getBypassIp_responseBody_2"
)

type GetBypassIPResponseBody struct {
	GetBypassIPResponseBody1 *GetBypassIPResponseBody1
	GetBypassIPResponseBody2 *GetBypassIPResponseBody2

	Type GetBypassIPResponseBodyType
}

func CreateGetBypassIPResponseBodyGetBypassIPResponseBody1(getBypassIPResponseBody1 GetBypassIPResponseBody1) GetBypassIPResponseBody {
	typ := GetBypassIPResponseBodyTypeGetBypassIPResponseBody1

	return GetBypassIPResponseBody{
		GetBypassIPResponseBody1: &getBypassIPResponseBody1,
		Type:                     typ,
	}
}

func CreateGetBypassIPResponseBodyGetBypassIPResponseBody2(getBypassIPResponseBody2 GetBypassIPResponseBody2) GetBypassIPResponseBody {
	typ := GetBypassIPResponseBodyTypeGetBypassIPResponseBody2

	return GetBypassIPResponseBody{
		GetBypassIPResponseBody2: &getBypassIPResponseBody2,
		Type:                     typ,
	}
}

func (u *GetBypassIPResponseBody) UnmarshalJSON(data []byte) error {

	var getBypassIPResponseBody1 GetBypassIPResponseBody1 = GetBypassIPResponseBody1{}
	if err := utils.UnmarshalJSON(data, &getBypassIPResponseBody1, "", true, true); err == nil {
		u.GetBypassIPResponseBody1 = &getBypassIPResponseBody1
		u.Type = GetBypassIPResponseBodyTypeGetBypassIPResponseBody1
		return nil
	}

	var getBypassIPResponseBody2 GetBypassIPResponseBody2 = GetBypassIPResponseBody2{}
	if err := utils.UnmarshalJSON(data, &getBypassIPResponseBody2, "", true, true); err == nil {
		u.GetBypassIPResponseBody2 = &getBypassIPResponseBody2
		u.Type = GetBypassIPResponseBodyTypeGetBypassIPResponseBody2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for GetBypassIPResponseBody", string(data))
}

func (u GetBypassIPResponseBody) MarshalJSON() ([]byte, error) {
	if u.GetBypassIPResponseBody1 != nil {
		return utils.MarshalJSON(u.GetBypassIPResponseBody1, "", true)
	}

	if u.GetBypassIPResponseBody2 != nil {
		return utils.MarshalJSON(u.GetBypassIPResponseBody2, "", true)
	}

	return nil, errors.New("could not marshal union type GetBypassIPResponseBody: all fields are null")
}

type GetBypassIPResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	OneOf    *GetBypassIPResponseBody
}

func (o *GetBypassIPResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetBypassIPResponse) GetOneOf() *GetBypassIPResponseBody {
	if o == nil {
		return nil
	}
	return o.OneOf
}
