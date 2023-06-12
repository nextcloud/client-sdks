/*
nextcloud

Nextcloud APIs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// ProvisioningApiAppsApiService ProvisioningApiAppsApi service
type ProvisioningApiAppsApiService service

type ApiProvisioningApiAppsDisableRequest struct {
	ctx context.Context
	ApiService *ProvisioningApiAppsApiService
	app string
	oCSAPIRequest *string
}

func (r ApiProvisioningApiAppsDisableRequest) OCSAPIRequest(oCSAPIRequest string) ApiProvisioningApiAppsDisableRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiProvisioningApiAppsDisableRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.ProvisioningApiAppsDisableExecute(r)
}

/*
ProvisioningApiAppsDisable Disable an app

This endpoint requires admin access

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param app ID of the app
 @return ApiProvisioningApiAppsDisableRequest
*/
func (a *ProvisioningApiAppsApiService) ProvisioningApiAppsDisable(ctx context.Context, app string) ApiProvisioningApiAppsDisableRequest {
	return ApiProvisioningApiAppsDisableRequest{
		ApiService: a,
		ctx: ctx,
		app: app,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *ProvisioningApiAppsApiService) ProvisioningApiAppsDisableExecute(r ApiProvisioningApiAppsDisableRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisioningApiAppsApiService.ProvisioningApiAppsDisable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/cloud/apps/{app}"
	localVarPath = strings.Replace(localVarPath, "{"+"app"+"}", url.PathEscape(parameterValueToString(r.app, "app")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "OCS-APIRequest", r.oCSAPIRequest, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiProvisioningApiAppsEnableRequest struct {
	ctx context.Context
	ApiService *ProvisioningApiAppsApiService
	app string
	oCSAPIRequest *string
}

func (r ApiProvisioningApiAppsEnableRequest) OCSAPIRequest(oCSAPIRequest string) ApiProvisioningApiAppsEnableRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiProvisioningApiAppsEnableRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.ProvisioningApiAppsEnableExecute(r)
}

/*
ProvisioningApiAppsEnable Enable an app

This endpoint requires admin access

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param app ID of the app
 @return ApiProvisioningApiAppsEnableRequest
*/
func (a *ProvisioningApiAppsApiService) ProvisioningApiAppsEnable(ctx context.Context, app string) ApiProvisioningApiAppsEnableRequest {
	return ApiProvisioningApiAppsEnableRequest{
		ApiService: a,
		ctx: ctx,
		app: app,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *ProvisioningApiAppsApiService) ProvisioningApiAppsEnableExecute(r ApiProvisioningApiAppsEnableRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisioningApiAppsApiService.ProvisioningApiAppsEnable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/cloud/apps/{app}"
	localVarPath = strings.Replace(localVarPath, "{"+"app"+"}", url.PathEscape(parameterValueToString(r.app, "app")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "OCS-APIRequest", r.oCSAPIRequest, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiProvisioningApiAppsGetAppInfoRequest struct {
	ctx context.Context
	ApiService *ProvisioningApiAppsApiService
	app string
	oCSAPIRequest *string
}

func (r ApiProvisioningApiAppsGetAppInfoRequest) OCSAPIRequest(oCSAPIRequest string) ApiProvisioningApiAppsGetAppInfoRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiProvisioningApiAppsGetAppInfoRequest) Execute() (*ProvisioningApiAppsGetAppInfo200Response, *http.Response, error) {
	return r.ApiService.ProvisioningApiAppsGetAppInfoExecute(r)
}

/*
ProvisioningApiAppsGetAppInfo Get the app info for an app

This endpoint requires admin access

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param app ID of the app
 @return ApiProvisioningApiAppsGetAppInfoRequest
*/
func (a *ProvisioningApiAppsApiService) ProvisioningApiAppsGetAppInfo(ctx context.Context, app string) ApiProvisioningApiAppsGetAppInfoRequest {
	return ApiProvisioningApiAppsGetAppInfoRequest{
		ApiService: a,
		ctx: ctx,
		app: app,
	}
}

// Execute executes the request
//  @return ProvisioningApiAppsGetAppInfo200Response
func (a *ProvisioningApiAppsApiService) ProvisioningApiAppsGetAppInfoExecute(r ApiProvisioningApiAppsGetAppInfoRequest) (*ProvisioningApiAppsGetAppInfo200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProvisioningApiAppsGetAppInfo200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisioningApiAppsApiService.ProvisioningApiAppsGetAppInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/cloud/apps/{app}"
	localVarPath = strings.Replace(localVarPath, "{"+"app"+"}", url.PathEscape(parameterValueToString(r.app, "app")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "OCS-APIRequest", r.oCSAPIRequest, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiProvisioningApiAppsGetAppsRequest struct {
	ctx context.Context
	ApiService *ProvisioningApiAppsApiService
	oCSAPIRequest *string
	filter *string
}

func (r ApiProvisioningApiAppsGetAppsRequest) OCSAPIRequest(oCSAPIRequest string) ApiProvisioningApiAppsGetAppsRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

// Filter for enabled or disabled apps
func (r ApiProvisioningApiAppsGetAppsRequest) Filter(filter string) ApiProvisioningApiAppsGetAppsRequest {
	r.filter = &filter
	return r
}

func (r ApiProvisioningApiAppsGetAppsRequest) Execute() (*ProvisioningApiAppsGetApps200Response, *http.Response, error) {
	return r.ApiService.ProvisioningApiAppsGetAppsExecute(r)
}

/*
ProvisioningApiAppsGetApps Get a list of installed apps

This endpoint requires admin access

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiProvisioningApiAppsGetAppsRequest
*/
func (a *ProvisioningApiAppsApiService) ProvisioningApiAppsGetApps(ctx context.Context) ApiProvisioningApiAppsGetAppsRequest {
	return ApiProvisioningApiAppsGetAppsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProvisioningApiAppsGetApps200Response
func (a *ProvisioningApiAppsApiService) ProvisioningApiAppsGetAppsExecute(r ApiProvisioningApiAppsGetAppsRequest) (*ProvisioningApiAppsGetApps200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProvisioningApiAppsGetApps200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisioningApiAppsApiService.ProvisioningApiAppsGetApps")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/cloud/apps"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	parameterAddToHeaderOrQuery(localVarHeaderParams, "OCS-APIRequest", r.oCSAPIRequest, "")
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
