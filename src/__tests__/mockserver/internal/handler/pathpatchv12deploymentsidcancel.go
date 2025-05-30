// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package handler

import (
	"fmt"
	"log"
	"mockserver/internal/handler/assert"
	"mockserver/internal/logging"
	"mockserver/internal/sdk/models/operations"
	"mockserver/internal/sdk/types"
	"mockserver/internal/sdk/utils"
	"mockserver/internal/tracking"
	"net/http"
)

func pathPatchV12DeploymentsIDCancel(dir *logging.HTTPFileDirectory, rt *tracking.RequestTracker) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		test := req.Header.Get("x-speakeasy-test-name")
		instanceID := req.Header.Get("x-speakeasy-test-instance-id")

		count := rt.GetRequestCount(test, instanceID)

		switch fmt.Sprintf("%s[%d]", test, count) {
		case "cancelDeployment[0]":
			dir.HandlerFunc("cancelDeployment", testCancelDeploymentCancelDeployment0)(w, req)
		default:
			http.Error(w, fmt.Sprintf("Unknown test: %s[%d]", test, count), http.StatusBadRequest)
		}
	}
}

func testCancelDeploymentCancelDeployment0(w http.ResponseWriter, req *http.Request) {
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
	respBody := &operations.CancelDeploymentResponseBody{
		Build: operations.CancelDeploymentBuild{
			Env: []string{
				"<value>",
				"<value>",
				"<value>",
			},
		},
		Env: []string{
			"<value>",
			"<value>",
		},
		InspectorURL:              types.String("https://concrete-gerbil.org"),
		IsInConcurrentBuildsQueue: false,
		IsInSystemBuildsQueue:     true,
		ProjectSettings:           operations.CancelDeploymentProjectSettings{},
		AliasAssigned:             false,
		BootedAt:                  8528.74,
		BuildingAt:                7590.79,
		BuildSkipped:              false,
		Creator: operations.CancelDeploymentCreator{
			UID: "<id>",
		},
		Public:    true,
		Status:    operations.CancelDeploymentStatusBuilding,
		ID:        "<id>",
		CreatedAt: 3071.53,
		Name:      "<value>",
		Meta: map[string]string{
			"key":  "<value>",
			"key1": "<value>",
			"key2": "<value>",
		},
		ReadyState: operations.CancelDeploymentReadyStateQueued,
		Regions: []string{
			"<value>",
		},
		Type:      operations.CancelDeploymentTypeLambdas,
		URL:       "https://timely-reboot.name",
		Version:   6067.47,
		CreatedIn: "<value>",
		OwnerID:   "<id>",
		Plan:      operations.CancelDeploymentPlanPro,
		ProjectID: "<id>",
		Routes: []operations.CancelDeploymentRoutes{
			operations.CreateCancelDeploymentRoutesCancelDeploymentRoutes2(
				operations.CancelDeploymentRoutes2{
					Handle: operations.CancelDeploymentRoutesHandleResource,
				},
			),
			operations.CreateCancelDeploymentRoutesCancelDeploymentRoutes1(
				operations.CancelDeploymentRoutes1{
					Src:        "<value>",
					Continue:   types.Bool(true),
					Middleware: types.Float64(4831.29),
				},
			),
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
