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

// EnvironmentsClient contains the methods for the Environments group.
// Don't use this type directly, use NewEnvironmentsClient() instead.
type EnvironmentsClient struct {
	internal *arm.Client
	rootScope string
}

// NewEnvironmentsClient creates a new instance of EnvironmentsClient with the specified values.
//   - rootScope - The scope in which the resource is present. UCP Scope is /planes/{planeType}/{planeName}/resourceGroup/{resourcegroupID}
//     and Azure resource scope is
//     /subscriptions/{subscriptionID}/resourceGroup/{resourcegroupID}
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewEnvironmentsClient(rootScope string, credential azcore.TokenCredential, options *arm.ClientOptions) (*EnvironmentsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &EnvironmentsClient{
		rootScope: rootScope,
	internal: cl,
	}
	return client, nil
}

// CreateOrUpdate - Create a EnvironmentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - environmentName - environment name
//   - resource - Resource create parameters.
//   - options - EnvironmentsClientCreateOrUpdateOptions contains the optional parameters for the EnvironmentsClient.CreateOrUpdate
//     method.
func (client *EnvironmentsClient) CreateOrUpdate(ctx context.Context, environmentName string, resource EnvironmentResource, options *EnvironmentsClientCreateOrUpdateOptions) (EnvironmentsClientCreateOrUpdateResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "EnvironmentsClient.CreateOrUpdate", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, environmentName, resource, options)
	if err != nil {
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.createOrUpdateHandleResponse(httpResp)
	return resp, err
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *EnvironmentsClient) createOrUpdateCreateRequest(ctx context.Context, environmentName string, resource EnvironmentResource, _ *EnvironmentsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
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
;	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *EnvironmentsClient) createOrUpdateHandleResponse(resp *http.Response) (EnvironmentsClientCreateOrUpdateResponse, error) {
	result := EnvironmentsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentResource); err != nil {
		return EnvironmentsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete a EnvironmentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - environmentName - environment name
//   - options - EnvironmentsClientDeleteOptions contains the optional parameters for the EnvironmentsClient.Delete method.
func (client *EnvironmentsClient) Delete(ctx context.Context, environmentName string, options *EnvironmentsClientDeleteOptions) (EnvironmentsClientDeleteResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "EnvironmentsClient.Delete", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, environmentName, options)
	if err != nil {
		return EnvironmentsClientDeleteResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentsClientDeleteResponse{}, err
	}
	return EnvironmentsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *EnvironmentsClient) deleteCreateRequest(ctx context.Context, environmentName string, _ *EnvironmentsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
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

// Get - Get a EnvironmentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - environmentName - environment name
//   - options - EnvironmentsClientGetOptions contains the optional parameters for the EnvironmentsClient.Get method.
func (client *EnvironmentsClient) Get(ctx context.Context, environmentName string, options *EnvironmentsClientGetOptions) (EnvironmentsClientGetResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "EnvironmentsClient.Get", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, environmentName, options)
	if err != nil {
		return EnvironmentsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *EnvironmentsClient) getCreateRequest(ctx context.Context, environmentName string, _ *EnvironmentsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
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
func (client *EnvironmentsClient) getHandleResponse(resp *http.Response) (EnvironmentsClientGetResponse, error) {
	result := EnvironmentsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentResource); err != nil {
		return EnvironmentsClientGetResponse{}, err
	}
	return result, nil
}

// GetMetadata - Gets recipe metadata including parameters and any constraints on the parameters.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - environmentName - environment name
//   - body - The content of the action request
//   - options - EnvironmentsClientGetMetadataOptions contains the optional parameters for the EnvironmentsClient.GetMetadata
//     method.
func (client *EnvironmentsClient) GetMetadata(ctx context.Context, environmentName string, body RecipeGetMetadata, options *EnvironmentsClientGetMetadataOptions) (EnvironmentsClientGetMetadataResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "EnvironmentsClient.GetMetadata", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getMetadataCreateRequest(ctx, environmentName, body, options)
	if err != nil {
		return EnvironmentsClientGetMetadataResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentsClientGetMetadataResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentsClientGetMetadataResponse{}, err
	}
	resp, err := client.getMetadataHandleResponse(httpResp)
	return resp, err
}

// getMetadataCreateRequest creates the GetMetadata request.
func (client *EnvironmentsClient) getMetadataCreateRequest(ctx context.Context, environmentName string, body RecipeGetMetadata, _ *EnvironmentsClientGetMetadataOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments/{environmentName}/getMetadata"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2023-10-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, body); err != nil {
	return nil, err
}
;	return req, nil
}

// getMetadataHandleResponse handles the GetMetadata response.
func (client *EnvironmentsClient) getMetadataHandleResponse(resp *http.Response) (EnvironmentsClientGetMetadataResponse, error) {
	result := EnvironmentsClientGetMetadataResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RecipeGetMetadataResponse); err != nil {
		return EnvironmentsClientGetMetadataResponse{}, err
	}
	return result, nil
}

// NewListByScopePager - List EnvironmentResource resources by Scope
//
// Generated from API version 2023-10-01-preview
//   - options - EnvironmentsClientListByScopeOptions contains the optional parameters for the EnvironmentsClient.NewListByScopePager
//     method.
func (client *EnvironmentsClient) NewListByScopePager(options *EnvironmentsClientListByScopeOptions) (*runtime.Pager[EnvironmentsClientListByScopeResponse]) {
	return runtime.NewPager(runtime.PagingHandler[EnvironmentsClientListByScopeResponse]{
		More: func(page EnvironmentsClientListByScopeResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *EnvironmentsClientListByScopeResponse) (EnvironmentsClientListByScopeResponse, error) {
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByScopeCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return EnvironmentsClientListByScopeResponse{}, err
			}
			return client.listByScopeHandleResponse(resp)
			},
		Tracer: client.internal.Tracer(),
	})
}

// listByScopeCreateRequest creates the ListByScope request.
func (client *EnvironmentsClient) listByScopeCreateRequest(ctx context.Context, _ *EnvironmentsClientListByScopeOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
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

// listByScopeHandleResponse handles the ListByScope response.
func (client *EnvironmentsClient) listByScopeHandleResponse(resp *http.Response) (EnvironmentsClientListByScopeResponse, error) {
	result := EnvironmentsClientListByScopeResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentResourceListResult); err != nil {
		return EnvironmentsClientListByScopeResponse{}, err
	}
	return result, nil
}

// Update - Update a EnvironmentResource
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2023-10-01-preview
//   - environmentName - environment name
//   - properties - The resource properties to be updated.
//   - options - EnvironmentsClientUpdateOptions contains the optional parameters for the EnvironmentsClient.Update method.
func (client *EnvironmentsClient) Update(ctx context.Context, environmentName string, properties EnvironmentResourceUpdate, options *EnvironmentsClientUpdateOptions) (EnvironmentsClientUpdateResponse, error) {
	var err error
	ctx, endSpan := runtime.StartSpan(ctx, "EnvironmentsClient.Update", client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateCreateRequest(ctx, environmentName, properties, options)
	if err != nil {
		return EnvironmentsClientUpdateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return EnvironmentsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return EnvironmentsClientUpdateResponse{}, err
	}
	resp, err := client.updateHandleResponse(httpResp)
	return resp, err
}

// updateCreateRequest creates the Update request.
func (client *EnvironmentsClient) updateCreateRequest(ctx context.Context, environmentName string, properties EnvironmentResourceUpdate, _ *EnvironmentsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/{rootScope}/providers/Applications.Core/environments/{environmentName}"
	urlPath = strings.ReplaceAll(urlPath, "{rootScope}", client.rootScope)
	if environmentName == "" {
		return nil, errors.New("parameter environmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{environmentName}", url.PathEscape(environmentName))
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
;	return req, nil
}

// updateHandleResponse handles the Update response.
func (client *EnvironmentsClient) updateHandleResponse(resp *http.Response) (EnvironmentsClientUpdateResponse, error) {
	result := EnvironmentsClientUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.EnvironmentResource); err != nil {
		return EnvironmentsClientUpdateResponse{}, err
	}
	return result, nil
}

