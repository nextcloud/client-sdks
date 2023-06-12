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


// UserStatusStatusesApiService UserStatusStatusesApi service
type UserStatusStatusesApiService service

type ApiUserStatusStatusesFindRequest struct {
	ctx context.Context
	ApiService *UserStatusStatusesApiService
	userId string
	oCSAPIRequest *string
}

func (r ApiUserStatusStatusesFindRequest) OCSAPIRequest(oCSAPIRequest string) ApiUserStatusStatusesFindRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiUserStatusStatusesFindRequest) Execute() (*UserStatusStatusesFind200Response, *http.Response, error) {
	return r.ApiService.UserStatusStatusesFindExecute(r)
}

/*
UserStatusStatusesFind Find the status of a user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of the user
 @return ApiUserStatusStatusesFindRequest
*/
func (a *UserStatusStatusesApiService) UserStatusStatusesFind(ctx context.Context, userId string) ApiUserStatusStatusesFindRequest {
	return ApiUserStatusStatusesFindRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
	}
}

// Execute executes the request
//  @return UserStatusStatusesFind200Response
func (a *UserStatusStatusesApiService) UserStatusStatusesFindExecute(r ApiUserStatusStatusesFindRequest) (*UserStatusStatusesFind200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserStatusStatusesFind200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserStatusStatusesApiService.UserStatusStatusesFind")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/user_status/api/v1/statuses/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type ApiUserStatusStatusesFindAllRequest struct {
	ctx context.Context
	ApiService *UserStatusStatusesApiService
	oCSAPIRequest *string
	limit *int64
	offset *int64
}

func (r ApiUserStatusStatusesFindAllRequest) OCSAPIRequest(oCSAPIRequest string) ApiUserStatusStatusesFindAllRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

// Maximum number of statuses to find
func (r ApiUserStatusStatusesFindAllRequest) Limit(limit int64) ApiUserStatusStatusesFindAllRequest {
	r.limit = &limit
	return r
}

// Offset for finding statuses
func (r ApiUserStatusStatusesFindAllRequest) Offset(offset int64) ApiUserStatusStatusesFindAllRequest {
	r.offset = &offset
	return r
}

func (r ApiUserStatusStatusesFindAllRequest) Execute() (*UserStatusStatusesFindAll200Response, *http.Response, error) {
	return r.ApiService.UserStatusStatusesFindAllExecute(r)
}

/*
UserStatusStatusesFindAll Find statuses of users

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiUserStatusStatusesFindAllRequest
*/
func (a *UserStatusStatusesApiService) UserStatusStatusesFindAll(ctx context.Context) ApiUserStatusStatusesFindAllRequest {
	return ApiUserStatusStatusesFindAllRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserStatusStatusesFindAll200Response
func (a *UserStatusStatusesApiService) UserStatusStatusesFindAllExecute(r ApiUserStatusStatusesFindAllRequest) (*UserStatusStatusesFindAll200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserStatusStatusesFindAll200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserStatusStatusesApiService.UserStatusStatusesFindAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/user_status/api/v1/statuses"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.offset != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offset", r.offset, "")
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
