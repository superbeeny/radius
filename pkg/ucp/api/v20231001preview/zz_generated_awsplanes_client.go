//go:build go1.18
// +build go1.18

// Licensed under the Apache License, Version 2.0 . See LICENSE in the repository root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package v20231001preview

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// AwsPlanesClient contains the methods for the AwsPlanes group.
// Don't use this type directly, use NewAwsPlanesClient() instead.
type AwsPlanesClient struct {
	internal *arm.Client
}

// NewAwsPlanesClient creates a new instance of AwsPlanesClient with the specified values.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewAwsPlanesClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AwsPlanesClient, error) {
	cl, err := arm.NewClient(moduleName+".AwsPlanesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &AwsPlanesClient{
	internal: cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update a plane
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - planeName - The plane name.
//   - resource - Resource create parameters.
//   - options - AwsPlanesClientBeginCreateOrUpdateOptions contains the optional parameters for the AwsPlanesClient.BeginCreateOrUpdate
//     method.
func (client *AwsPlanesClient) BeginCreateOrUpdate(ctx context.Context, planeName string, resource AwsPlaneResource, options *AwsPlanesClientBeginCreateOrUpdateOptions) (*runtime.Poller[AwsPlanesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, planeName, resource, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AwsPlanesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AwsPlanesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Create or update a plane
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *AwsPlanesClient) createOrUpdate(ctx context.Context, planeName string, resource AwsPlaneResource, options *AwsPlanesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.createOrUpdateCreateRequest(ctx, planeName, resource, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *AwsPlanesClient) createOrUpdateCreateRequest(ctx context.Context, planeName string, resource AwsPlaneResource, options *AwsPlanesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/planes/aws/{planeName}"
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, resource); err != nil {
	return nil, err
}
	return req, nil
}

// BeginDelete - Delete a plane
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - planeName - The plane name.
//   - options - AwsPlanesClientBeginDeleteOptions contains the optional parameters for the AwsPlanesClient.BeginDelete method.
func (client *AwsPlanesClient) BeginDelete(ctx context.Context, planeName string, options *AwsPlanesClientBeginDeleteOptions) (*runtime.Poller[AwsPlanesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, planeName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AwsPlanesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AwsPlanesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Delete a plane
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *AwsPlanesClient) deleteOperation(ctx context.Context, planeName string, options *AwsPlanesClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	req, err := client.deleteCreateRequest(ctx, planeName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *AwsPlanesClient) deleteCreateRequest(ctx context.Context, planeName string, options *AwsPlanesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/planes/aws/{planeName}"
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get a plane by name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - planeName - The plane name.
//   - options - AwsPlanesClientGetOptions contains the optional parameters for the AwsPlanesClient.Get method.
func (client *AwsPlanesClient) Get(ctx context.Context, planeName string, options *AwsPlanesClientGetOptions) (AwsPlanesClientGetResponse, error) {
	var err error
	req, err := client.getCreateRequest(ctx, planeName, options)
	if err != nil {
		return AwsPlanesClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return AwsPlanesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return AwsPlanesClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *AwsPlanesClient) getCreateRequest(ctx context.Context, planeName string, options *AwsPlanesClientGetOptions) (*policy.Request, error) {
	urlPath := "/planes/aws/{planeName}"
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AwsPlanesClient) getHandleResponse(resp *http.Response) (AwsPlanesClientGetResponse, error) {
	result := AwsPlanesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AwsPlaneResource); err != nil {
		return AwsPlanesClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - List AWS planes
//
// Generated from API version 2023-10-01-preview
//   - options - AwsPlanesClientListOptions contains the optional parameters for the AwsPlanesClient.NewListPager method.
func (client *AwsPlanesClient) NewListPager(options *AwsPlanesClientListOptions) (*runtime.Pager[AwsPlanesClientListResponse]) {
	return runtime.NewPager(runtime.PagingHandler[AwsPlanesClientListResponse]{
		More: func(page AwsPlanesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AwsPlanesClientListResponse) (AwsPlanesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AwsPlanesClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return AwsPlanesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AwsPlanesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *AwsPlanesClient) listCreateRequest(ctx context.Context, options *AwsPlanesClientListOptions) (*policy.Request, error) {
	urlPath := "/planes/aws"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *AwsPlanesClient) listHandleResponse(resp *http.Response) (AwsPlanesClientListResponse, error) {
	result := AwsPlanesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AwsPlaneResourceListResult); err != nil {
		return AwsPlanesClientListResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Update a plane
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - planeName - The plane name.
//   - properties - The resource properties to be updated.
//   - options - AwsPlanesClientBeginUpdateOptions contains the optional parameters for the AwsPlanesClient.BeginUpdate method.
func (client *AwsPlanesClient) BeginUpdate(ctx context.Context, planeName string, properties AwsPlaneResourceTagsUpdate, options *AwsPlanesClientBeginUpdateOptions) (*runtime.Poller[AwsPlanesClientUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.update(ctx, planeName, properties, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[AwsPlanesClientUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[AwsPlanesClientUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Update - Update a plane
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
func (client *AwsPlanesClient) update(ctx context.Context, planeName string, properties AwsPlaneResourceTagsUpdate, options *AwsPlanesClientBeginUpdateOptions) (*http.Response, error) {
	var err error
	req, err := client.updateCreateRequest(ctx, planeName, properties, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// updateCreateRequest creates the Update request.
func (client *AwsPlanesClient) updateCreateRequest(ctx context.Context, planeName string, properties AwsPlaneResourceTagsUpdate, options *AwsPlanesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/planes/aws/{planeName}"
	if planeName == "" {
		return nil, errors.New("parameter planeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{planeName}", url.PathEscape(planeName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, properties); err != nil {
	return nil, err
}
	return req, nil
}

