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
	"os"
)


// ThemingUserThemeApiService ThemingUserThemeApi service
type ThemingUserThemeApiService service

type ApiThemingUserThemeDeleteBackgroundRequest struct {
	ctx context.Context
	ApiService *ThemingUserThemeApiService
	oCSAPIRequest *string
}

func (r ApiThemingUserThemeDeleteBackgroundRequest) OCSAPIRequest(oCSAPIRequest string) ApiThemingUserThemeDeleteBackgroundRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiThemingUserThemeDeleteBackgroundRequest) Execute() (*ThemingBackground, *http.Response, error) {
	return r.ApiService.ThemingUserThemeDeleteBackgroundExecute(r)
}

/*
ThemingUserThemeDeleteBackground Delete the background

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiThemingUserThemeDeleteBackgroundRequest
*/
func (a *ThemingUserThemeApiService) ThemingUserThemeDeleteBackground(ctx context.Context) ApiThemingUserThemeDeleteBackgroundRequest {
	return ApiThemingUserThemeDeleteBackgroundRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ThemingBackground
func (a *ThemingUserThemeApiService) ThemingUserThemeDeleteBackgroundExecute(r ApiThemingUserThemeDeleteBackgroundRequest) (*ThemingBackground, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ThemingBackground
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThemingUserThemeApiService.ThemingUserThemeDeleteBackground")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/apps/theming/background/custom"

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

type ApiThemingUserThemeDisableThemeRequest struct {
	ctx context.Context
	ApiService *ThemingUserThemeApiService
	themeId string
	oCSAPIRequest *string
}

func (r ApiThemingUserThemeDisableThemeRequest) OCSAPIRequest(oCSAPIRequest string) ApiThemingUserThemeDisableThemeRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiThemingUserThemeDisableThemeRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.ThemingUserThemeDisableThemeExecute(r)
}

/*
ThemingUserThemeDisableTheme Disable theme

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param themeId the theme ID
 @return ApiThemingUserThemeDisableThemeRequest
*/
func (a *ThemingUserThemeApiService) ThemingUserThemeDisableTheme(ctx context.Context, themeId string) ApiThemingUserThemeDisableThemeRequest {
	return ApiThemingUserThemeDisableThemeRequest{
		ApiService: a,
		ctx: ctx,
		themeId: themeId,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *ThemingUserThemeApiService) ThemingUserThemeDisableThemeExecute(r ApiThemingUserThemeDisableThemeRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThemingUserThemeApiService.ThemingUserThemeDisableTheme")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/theming/api/v1/theme/{themeId}"
	localVarPath = strings.Replace(localVarPath, "{"+"themeId"+"}", url.PathEscape(parameterValueToString(r.themeId, "themeId")), -1)

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

type ApiThemingUserThemeEnableThemeRequest struct {
	ctx context.Context
	ApiService *ThemingUserThemeApiService
	themeId string
	oCSAPIRequest *string
}

func (r ApiThemingUserThemeEnableThemeRequest) OCSAPIRequest(oCSAPIRequest string) ApiThemingUserThemeEnableThemeRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiThemingUserThemeEnableThemeRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.ThemingUserThemeEnableThemeExecute(r)
}

/*
ThemingUserThemeEnableTheme Enable theme

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param themeId the theme ID
 @return ApiThemingUserThemeEnableThemeRequest
*/
func (a *ThemingUserThemeApiService) ThemingUserThemeEnableTheme(ctx context.Context, themeId string) ApiThemingUserThemeEnableThemeRequest {
	return ApiThemingUserThemeEnableThemeRequest{
		ApiService: a,
		ctx: ctx,
		themeId: themeId,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *ThemingUserThemeApiService) ThemingUserThemeEnableThemeExecute(r ApiThemingUserThemeEnableThemeRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThemingUserThemeApiService.ThemingUserThemeEnableTheme")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/theming/api/v1/theme/{themeId}/enable"
	localVarPath = strings.Replace(localVarPath, "{"+"themeId"+"}", url.PathEscape(parameterValueToString(r.themeId, "themeId")), -1)

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

type ApiThemingUserThemeGetBackgroundRequest struct {
	ctx context.Context
	ApiService *ThemingUserThemeApiService
	oCSAPIRequest *string
}

func (r ApiThemingUserThemeGetBackgroundRequest) OCSAPIRequest(oCSAPIRequest string) ApiThemingUserThemeGetBackgroundRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiThemingUserThemeGetBackgroundRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.ThemingUserThemeGetBackgroundExecute(r)
}

/*
ThemingUserThemeGetBackground Get the background image

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiThemingUserThemeGetBackgroundRequest
*/
func (a *ThemingUserThemeApiService) ThemingUserThemeGetBackground(ctx context.Context) ApiThemingUserThemeGetBackgroundRequest {
	return ApiThemingUserThemeGetBackgroundRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *ThemingUserThemeApiService) ThemingUserThemeGetBackgroundExecute(r ApiThemingUserThemeGetBackgroundRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThemingUserThemeApiService.ThemingUserThemeGetBackground")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/apps/theming/background"

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
	localVarHTTPHeaderAccepts := []string{"*/*"}

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

type ApiThemingUserThemeSetBackgroundRequest struct {
	ctx context.Context
	ApiService *ThemingUserThemeApiService
	type_ string
	oCSAPIRequest *string
	value *string
	color *string
}

func (r ApiThemingUserThemeSetBackgroundRequest) OCSAPIRequest(oCSAPIRequest string) ApiThemingUserThemeSetBackgroundRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

// Path of the background image
func (r ApiThemingUserThemeSetBackgroundRequest) Value(value string) ApiThemingUserThemeSetBackgroundRequest {
	r.value = &value
	return r
}

// Color for the background
func (r ApiThemingUserThemeSetBackgroundRequest) Color(color string) ApiThemingUserThemeSetBackgroundRequest {
	r.color = &color
	return r
}

func (r ApiThemingUserThemeSetBackgroundRequest) Execute() (*ThemingBackground, *http.Response, error) {
	return r.ApiService.ThemingUserThemeSetBackgroundExecute(r)
}

/*
ThemingUserThemeSetBackground Set the background

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param type_ Type of background
 @return ApiThemingUserThemeSetBackgroundRequest
*/
func (a *ThemingUserThemeApiService) ThemingUserThemeSetBackground(ctx context.Context, type_ string) ApiThemingUserThemeSetBackgroundRequest {
	return ApiThemingUserThemeSetBackgroundRequest{
		ApiService: a,
		ctx: ctx,
		type_: type_,
	}
}

// Execute executes the request
//  @return ThemingBackground
func (a *ThemingUserThemeApiService) ThemingUserThemeSetBackgroundExecute(r ApiThemingUserThemeSetBackgroundRequest) (*ThemingBackground, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ThemingBackground
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ThemingUserThemeApiService.ThemingUserThemeSetBackground")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/apps/theming/background/{type}"
	localVarPath = strings.Replace(localVarPath, "{"+"type"+"}", url.PathEscape(parameterValueToString(r.type_, "type_")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	if r.value != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "value", r.value, "")
	}
	if r.color != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "color", r.color, "")
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
