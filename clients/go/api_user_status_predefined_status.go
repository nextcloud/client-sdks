/*
nextcloud

Nextcloud APIs

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client_sdk

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type UserStatusPredefinedStatusAPI interface {

	/*
	UserStatusPredefinedStatusFindAll Get all predefined messages

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest
	*/
	UserStatusPredefinedStatusFindAll(ctx context.Context) UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest

	// UserStatusPredefinedStatusFindAllExecute executes the request
	//  @return UserStatusPredefinedStatusFindAll200Response
	UserStatusPredefinedStatusFindAllExecute(r UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest) (*UserStatusPredefinedStatusFindAll200Response, *http.Response, error)
}

// UserStatusPredefinedStatusAPIService UserStatusPredefinedStatusAPI service
type UserStatusPredefinedStatusAPIService service

type UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest struct {
	ctx context.Context
	ApiService UserStatusPredefinedStatusAPI
	oCSAPIRequest *string
}

func (r UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest) OCSAPIRequest(oCSAPIRequest string) UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest) Execute() (*UserStatusPredefinedStatusFindAll200Response, *http.Response, error) {
	return r.ApiService.UserStatusPredefinedStatusFindAllExecute(r)
}

/*
UserStatusPredefinedStatusFindAll Get all predefined messages

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest
*/
func (a *UserStatusPredefinedStatusAPIService) UserStatusPredefinedStatusFindAll(ctx context.Context) UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest {
	return UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return UserStatusPredefinedStatusFindAll200Response
func (a *UserStatusPredefinedStatusAPIService) UserStatusPredefinedStatusFindAllExecute(r UserStatusPredefinedStatusAPIUserStatusPredefinedStatusFindAllRequest) (*UserStatusPredefinedStatusFindAll200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UserStatusPredefinedStatusFindAll200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserStatusPredefinedStatusAPIService.UserStatusPredefinedStatusFindAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/user_status/api/v1/predefined_statuses"

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
