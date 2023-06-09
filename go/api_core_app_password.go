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
)


// CoreAppPasswordApiService CoreAppPasswordApi service
type CoreAppPasswordApiService service

type ApiCoreAppPasswordDeleteAppPasswordRequest struct {
	ctx context.Context
	ApiService *CoreAppPasswordApiService
	oCSAPIRequest *string
}

func (r ApiCoreAppPasswordDeleteAppPasswordRequest) OCSAPIRequest(oCSAPIRequest string) ApiCoreAppPasswordDeleteAppPasswordRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiCoreAppPasswordDeleteAppPasswordRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.CoreAppPasswordDeleteAppPasswordExecute(r)
}

/*
CoreAppPasswordDeleteAppPassword Delete app password

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCoreAppPasswordDeleteAppPasswordRequest
*/
func (a *CoreAppPasswordApiService) CoreAppPasswordDeleteAppPassword(ctx context.Context) ApiCoreAppPasswordDeleteAppPasswordRequest {
	return ApiCoreAppPasswordDeleteAppPasswordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *CoreAppPasswordApiService) CoreAppPasswordDeleteAppPasswordExecute(r ApiCoreAppPasswordDeleteAppPasswordRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreAppPasswordApiService.CoreAppPasswordDeleteAppPassword")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/core/apppassword"

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

type ApiCoreAppPasswordGetAppPasswordRequest struct {
	ctx context.Context
	ApiService *CoreAppPasswordApiService
	oCSAPIRequest *string
}

func (r ApiCoreAppPasswordGetAppPasswordRequest) OCSAPIRequest(oCSAPIRequest string) ApiCoreAppPasswordGetAppPasswordRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiCoreAppPasswordGetAppPasswordRequest) Execute() (*CoreAppPasswordGetAppPassword200Response, *http.Response, error) {
	return r.ApiService.CoreAppPasswordGetAppPasswordExecute(r)
}

/*
CoreAppPasswordGetAppPassword Create app password

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCoreAppPasswordGetAppPasswordRequest
*/
func (a *CoreAppPasswordApiService) CoreAppPasswordGetAppPassword(ctx context.Context) ApiCoreAppPasswordGetAppPasswordRequest {
	return ApiCoreAppPasswordGetAppPasswordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreAppPasswordGetAppPassword200Response
func (a *CoreAppPasswordApiService) CoreAppPasswordGetAppPasswordExecute(r ApiCoreAppPasswordGetAppPasswordRequest) (*CoreAppPasswordGetAppPassword200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreAppPasswordGetAppPassword200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreAppPasswordApiService.CoreAppPasswordGetAppPassword")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/core/getapppassword"

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

type ApiCoreAppPasswordRotateAppPasswordRequest struct {
	ctx context.Context
	ApiService *CoreAppPasswordApiService
	oCSAPIRequest *string
}

func (r ApiCoreAppPasswordRotateAppPasswordRequest) OCSAPIRequest(oCSAPIRequest string) ApiCoreAppPasswordRotateAppPasswordRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiCoreAppPasswordRotateAppPasswordRequest) Execute() (*CoreAppPasswordGetAppPassword200Response, *http.Response, error) {
	return r.ApiService.CoreAppPasswordRotateAppPasswordExecute(r)
}

/*
CoreAppPasswordRotateAppPassword Rotate app password

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCoreAppPasswordRotateAppPasswordRequest
*/
func (a *CoreAppPasswordApiService) CoreAppPasswordRotateAppPassword(ctx context.Context) ApiCoreAppPasswordRotateAppPasswordRequest {
	return ApiCoreAppPasswordRotateAppPasswordRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreAppPasswordGetAppPassword200Response
func (a *CoreAppPasswordApiService) CoreAppPasswordRotateAppPasswordExecute(r ApiCoreAppPasswordRotateAppPasswordRequest) (*CoreAppPasswordGetAppPassword200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreAppPasswordGetAppPassword200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreAppPasswordApiService.CoreAppPasswordRotateAppPassword")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/core/apppassword/rotate"

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
