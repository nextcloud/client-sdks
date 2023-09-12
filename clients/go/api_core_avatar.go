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
	"strings"
	"os"
)


type CoreAvatarAPI interface {

	/*
	CoreAvatarGetAvatar Get the avatar

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of the user
	@param size Size of the avatar
	@return CoreAvatarAPICoreAvatarGetAvatarRequest
	*/
	CoreAvatarGetAvatar(ctx context.Context, userId string, size int64) CoreAvatarAPICoreAvatarGetAvatarRequest

	// CoreAvatarGetAvatarExecute executes the request
	//  @return *os.File
	CoreAvatarGetAvatarExecute(r CoreAvatarAPICoreAvatarGetAvatarRequest) (*os.File, *http.Response, error)

	/*
	CoreAvatarGetAvatarDark Get the dark avatar

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param userId ID of the user
	@param size Size of the avatar
	@return CoreAvatarAPICoreAvatarGetAvatarDarkRequest
	*/
	CoreAvatarGetAvatarDark(ctx context.Context, userId string, size int64) CoreAvatarAPICoreAvatarGetAvatarDarkRequest

	// CoreAvatarGetAvatarDarkExecute executes the request
	//  @return *os.File
	CoreAvatarGetAvatarDarkExecute(r CoreAvatarAPICoreAvatarGetAvatarDarkRequest) (*os.File, *http.Response, error)
}

// CoreAvatarAPIService CoreAvatarAPI service
type CoreAvatarAPIService service

type CoreAvatarAPICoreAvatarGetAvatarRequest struct {
	ctx context.Context
	ApiService CoreAvatarAPI
	userId string
	size int64
}

func (r CoreAvatarAPICoreAvatarGetAvatarRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.CoreAvatarGetAvatarExecute(r)
}

/*
CoreAvatarGetAvatar Get the avatar

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of the user
 @param size Size of the avatar
 @return CoreAvatarAPICoreAvatarGetAvatarRequest
*/
func (a *CoreAvatarAPIService) CoreAvatarGetAvatar(ctx context.Context, userId string, size int64) CoreAvatarAPICoreAvatarGetAvatarRequest {
	return CoreAvatarAPICoreAvatarGetAvatarRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		size: size,
	}
}

// Execute executes the request
//  @return *os.File
func (a *CoreAvatarAPIService) CoreAvatarGetAvatarExecute(r CoreAvatarAPICoreAvatarGetAvatarRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreAvatarAPIService.CoreAvatarGetAvatar")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/avatar/{userId}/{size}"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"size"+"}", url.PathEscape(parameterValueToString(r.size, "size")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type CoreAvatarAPICoreAvatarGetAvatarDarkRequest struct {
	ctx context.Context
	ApiService CoreAvatarAPI
	userId string
	size int64
}

func (r CoreAvatarAPICoreAvatarGetAvatarDarkRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.CoreAvatarGetAvatarDarkExecute(r)
}

/*
CoreAvatarGetAvatarDark Get the dark avatar

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param userId ID of the user
 @param size Size of the avatar
 @return CoreAvatarAPICoreAvatarGetAvatarDarkRequest
*/
func (a *CoreAvatarAPIService) CoreAvatarGetAvatarDark(ctx context.Context, userId string, size int64) CoreAvatarAPICoreAvatarGetAvatarDarkRequest {
	return CoreAvatarAPICoreAvatarGetAvatarDarkRequest{
		ApiService: a,
		ctx: ctx,
		userId: userId,
		size: size,
	}
}

// Execute executes the request
//  @return *os.File
func (a *CoreAvatarAPIService) CoreAvatarGetAvatarDarkExecute(r CoreAvatarAPICoreAvatarGetAvatarDarkRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreAvatarAPIService.CoreAvatarGetAvatarDark")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/avatar/{userId}/{size}/dark"
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"size"+"}", url.PathEscape(parameterValueToString(r.size, "size")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
