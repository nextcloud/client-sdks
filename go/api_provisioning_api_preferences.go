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


// ProvisioningApiPreferencesApiService ProvisioningApiPreferencesApi service
type ProvisioningApiPreferencesApiService service

type ApiProvisioningApiPreferencesDeleteMultiplePreferenceRequest struct {
	ctx context.Context
	ApiService *ProvisioningApiPreferencesApiService
	configKeys *string
	appId string
	oCSAPIRequest *string
}

// Keys to delete
func (r ApiProvisioningApiPreferencesDeleteMultiplePreferenceRequest) ConfigKeys(configKeys string) ApiProvisioningApiPreferencesDeleteMultiplePreferenceRequest {
	r.configKeys = &configKeys
	return r
}

func (r ApiProvisioningApiPreferencesDeleteMultiplePreferenceRequest) OCSAPIRequest(oCSAPIRequest string) ApiProvisioningApiPreferencesDeleteMultiplePreferenceRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiProvisioningApiPreferencesDeleteMultiplePreferenceRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.ProvisioningApiPreferencesDeleteMultiplePreferenceExecute(r)
}

/*
ProvisioningApiPreferencesDeleteMultiplePreference Delete multiple preferences for an app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId ID of the app
 @return ApiProvisioningApiPreferencesDeleteMultiplePreferenceRequest
*/
func (a *ProvisioningApiPreferencesApiService) ProvisioningApiPreferencesDeleteMultiplePreference(ctx context.Context, appId string) ApiProvisioningApiPreferencesDeleteMultiplePreferenceRequest {
	return ApiProvisioningApiPreferencesDeleteMultiplePreferenceRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *ProvisioningApiPreferencesApiService) ProvisioningApiPreferencesDeleteMultiplePreferenceExecute(r ApiProvisioningApiPreferencesDeleteMultiplePreferenceRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisioningApiPreferencesApiService.ProvisioningApiPreferencesDeleteMultiplePreference")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.configKeys == nil {
		return localVarReturnValue, nil, reportError("configKeys is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "configKeys", r.configKeys, "")
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

type ApiProvisioningApiPreferencesDeletePreferenceRequest struct {
	ctx context.Context
	ApiService *ProvisioningApiPreferencesApiService
	appId string
	configKey string
	oCSAPIRequest *string
}

func (r ApiProvisioningApiPreferencesDeletePreferenceRequest) OCSAPIRequest(oCSAPIRequest string) ApiProvisioningApiPreferencesDeletePreferenceRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiProvisioningApiPreferencesDeletePreferenceRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.ProvisioningApiPreferencesDeletePreferenceExecute(r)
}

/*
ProvisioningApiPreferencesDeletePreference Delete a preference for an app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId ID of the app
 @param configKey Key to delete
 @return ApiProvisioningApiPreferencesDeletePreferenceRequest
*/
func (a *ProvisioningApiPreferencesApiService) ProvisioningApiPreferencesDeletePreference(ctx context.Context, appId string, configKey string) ApiProvisioningApiPreferencesDeletePreferenceRequest {
	return ApiProvisioningApiPreferencesDeletePreferenceRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *ProvisioningApiPreferencesApiService) ProvisioningApiPreferencesDeletePreferenceExecute(r ApiProvisioningApiPreferencesDeletePreferenceRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisioningApiPreferencesApiService.ProvisioningApiPreferencesDeletePreference")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}/{configKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

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

type ApiProvisioningApiPreferencesSetMultiplePreferencesRequest struct {
	ctx context.Context
	ApiService *ProvisioningApiPreferencesApiService
	configs *string
	appId string
	oCSAPIRequest *string
}

// Key-value pairs of the preferences
func (r ApiProvisioningApiPreferencesSetMultiplePreferencesRequest) Configs(configs string) ApiProvisioningApiPreferencesSetMultiplePreferencesRequest {
	r.configs = &configs
	return r
}

func (r ApiProvisioningApiPreferencesSetMultiplePreferencesRequest) OCSAPIRequest(oCSAPIRequest string) ApiProvisioningApiPreferencesSetMultiplePreferencesRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiProvisioningApiPreferencesSetMultiplePreferencesRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.ProvisioningApiPreferencesSetMultiplePreferencesExecute(r)
}

/*
ProvisioningApiPreferencesSetMultiplePreferences Update multiple preference values of an app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId ID of the app
 @return ApiProvisioningApiPreferencesSetMultiplePreferencesRequest
*/
func (a *ProvisioningApiPreferencesApiService) ProvisioningApiPreferencesSetMultiplePreferences(ctx context.Context, appId string) ApiProvisioningApiPreferencesSetMultiplePreferencesRequest {
	return ApiProvisioningApiPreferencesSetMultiplePreferencesRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *ProvisioningApiPreferencesApiService) ProvisioningApiPreferencesSetMultiplePreferencesExecute(r ApiProvisioningApiPreferencesSetMultiplePreferencesRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisioningApiPreferencesApiService.ProvisioningApiPreferencesSetMultiplePreferences")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.configs == nil {
		return localVarReturnValue, nil, reportError("configs is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "configs", r.configs, "")
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

type ApiProvisioningApiPreferencesSetPreferenceRequest struct {
	ctx context.Context
	ApiService *ProvisioningApiPreferencesApiService
	configValue *string
	appId string
	configKey string
	oCSAPIRequest *string
}

// New value
func (r ApiProvisioningApiPreferencesSetPreferenceRequest) ConfigValue(configValue string) ApiProvisioningApiPreferencesSetPreferenceRequest {
	r.configValue = &configValue
	return r
}

func (r ApiProvisioningApiPreferencesSetPreferenceRequest) OCSAPIRequest(oCSAPIRequest string) ApiProvisioningApiPreferencesSetPreferenceRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiProvisioningApiPreferencesSetPreferenceRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.ProvisioningApiPreferencesSetPreferenceExecute(r)
}

/*
ProvisioningApiPreferencesSetPreference Update a preference value of an app

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param appId ID of the app
 @param configKey Key of the preference
 @return ApiProvisioningApiPreferencesSetPreferenceRequest
*/
func (a *ProvisioningApiPreferencesApiService) ProvisioningApiPreferencesSetPreference(ctx context.Context, appId string, configKey string) ApiProvisioningApiPreferencesSetPreferenceRequest {
	return ApiProvisioningApiPreferencesSetPreferenceRequest{
		ApiService: a,
		ctx: ctx,
		appId: appId,
		configKey: configKey,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *ProvisioningApiPreferencesApiService) ProvisioningApiPreferencesSetPreferenceExecute(r ApiProvisioningApiPreferencesSetPreferenceRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ProvisioningApiPreferencesApiService.ProvisioningApiPreferencesSetPreference")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}/{configKey}"
	localVarPath = strings.Replace(localVarPath, "{"+"appId"+"}", url.PathEscape(parameterValueToString(r.appId, "appId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"configKey"+"}", url.PathEscape(parameterValueToString(r.configKey, "configKey")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.configValue == nil {
		return localVarReturnValue, nil, reportError("configValue is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "configValue", r.configValue, "")
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
