// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"fmt"
	"log"
	"mockserver/internal/handler/assert"
	"mockserver/internal/logging"
	"mockserver/internal/sdk/models/components"
	"mockserver/internal/sdk/models/operations"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"mockserver/internal/tracking"
	"net/http"
)

func pathGetV6Deployments(dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")
		instanceID := req.Header.Get("x-speakeasy-test-instance-id")

		count := rt.GetRequestCount(test, instanceID)

		switch fmt.Sprintf("%s[%d]", test, count) {
		case "getDeployments[0]":
			dir.HandlerFunc("getDeployments", testGetDeploymentsGetDeployments0)(w, req)
		default:
			http.Error(w, fmt.Sprintf("Unknown test: %s[%d]", test, count), http.StatusBadRequest)
		}
	}
}

func testGetDeploymentsGetDeployments0(w http.ResponseWriter, req *http.Request) {
	if err := assert.SecurityAuthorizationHeader(req, true, "Bearer"); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	if err := assert.AcceptHeader(req, []string{"application/json"}); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := assert.HeaderExists(req, "User-Agent"); err != nil {
		log.Printf("assertion error: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	respBody := &operations.GetDeploymentsResponseBody{
		Pagination: components.Pagination{
			Count: 20,
			Next:  types.Float64(1540095775951),
			Prev:  types.Float64(1540095775951),
		},
		Deployments: []operations.Deployments{
			operations.Deployments{
				UID:                    "dpl_2euZBFqxYdDMDG1jTrHFnNZ2eUVa",
				Name:                   "docs",
				URL:                    "docs-9jaeg38me.vercel.app",
				Created:                1609492210000,
				Deleted:                types.Float64(1609492210000),
				Undeleted:              types.Float64(1609492210000),
				SoftDeletedByRetention: types.Bool(true),
				Source:                 operations.GetDeploymentsSourceCli.ToPointer(),
				State:                  operations.GetDeploymentsStateReady.ToPointer(),
				ReadyState:             operations.GetDeploymentsReadyStateReady.ToPointer(),
				Type:                   operations.GetDeploymentsTypeLambdas,
				Creator: operations.GetDeploymentsCreator{
					UID:         "eLrCnEgbKhsHyfbiNR7E8496",
					Email:       types.String("example@example.com"),
					Username:    types.String("johndoe"),
					GithubLogin: types.String("johndoe"),
					GitlabLogin: types.String("johndoe"),
				},
				Target:       operations.GetDeploymentsTargetProduction.ToPointer(),
				CreatedAt:    types.Float64(1609492210000),
				BuildingAt:   types.Float64(1609492210000),
				Ready:        types.Float64(1609492210000),
				InspectorURL: types.String("https://vercel.com/acme/nextjs/J1hXN00qjUeoYfpEEf7dnDtpSiVq"),
			},
			operations.Deployments{
				UID:                    "dpl_2euZBFqxYdDMDG1jTrHFnNZ2eUVa",
				Name:                   "docs",
				URL:                    "docs-9jaeg38me.vercel.app",
				Created:                1609492210000,
				Deleted:                types.Float64(1609492210000),
				Undeleted:              types.Float64(1609492210000),
				SoftDeletedByRetention: types.Bool(true),
				Source:                 operations.GetDeploymentsSourceCli.ToPointer(),
				State:                  operations.GetDeploymentsStateReady.ToPointer(),
				ReadyState:             operations.GetDeploymentsReadyStateReady.ToPointer(),
				Type:                   operations.GetDeploymentsTypeLambdas,
				Creator: operations.GetDeploymentsCreator{
					UID:         "eLrCnEgbKhsHyfbiNR7E8496",
					Email:       types.String("example@example.com"),
					Username:    types.String("johndoe"),
					GithubLogin: types.String("johndoe"),
					GitlabLogin: types.String("johndoe"),
				},
				Target:       operations.GetDeploymentsTargetProduction.ToPointer(),
				CreatedAt:    types.Float64(1609492210000),
				BuildingAt:   types.Float64(1609492210000),
				Ready:        types.Float64(1609492210000),
				InspectorURL: types.String("https://vercel.com/acme/nextjs/J1hXN00qjUeoYfpEEf7dnDtpSiVq"),
			},
		},
	}
	respBodyBytes, err := utils.MarshalJSON(respBody, "", true)

	if err != nil {
		http.Error(
			w,
			"Unable to encode response body as JSON: "+err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(respBodyBytes)
}
