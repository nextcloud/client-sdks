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
)


type FilesOpenLocalEditorAPI interface {

	/*
	FilesOpenLocalEditorCreate Create a local editor

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest
	*/
	FilesOpenLocalEditorCreate(ctx context.Context) FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest

	// FilesOpenLocalEditorCreateExecute executes the request
	//  @return FilesOpenLocalEditorCreate200Response
	FilesOpenLocalEditorCreateExecute(r FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest) (*FilesOpenLocalEditorCreate200Response, *http.Response, error)

	/*
	FilesOpenLocalEditorValidate Validate a local editor

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param token Token of the local editor
	@return FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest
	*/
	FilesOpenLocalEditorValidate(ctx context.Context, token string) FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest

	// FilesOpenLocalEditorValidateExecute executes the request
	//  @return FilesOpenLocalEditorValidate200Response
	FilesOpenLocalEditorValidateExecute(r FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest) (*FilesOpenLocalEditorValidate200Response, *http.Response, error)
}

// FilesOpenLocalEditorAPIService FilesOpenLocalEditorAPI service
type FilesOpenLocalEditorAPIService service

type FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest struct {
	ctx context.Context
	ApiService FilesOpenLocalEditorAPI
	path *string
	oCSAPIRequest *string
}

// Path of the file
func (r FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest) Path(path string) FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest {
	r.path = &path
	return r
}

func (r FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest) OCSAPIRequest(oCSAPIRequest string) FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest) Execute() (*FilesOpenLocalEditorCreate200Response, *http.Response, error) {
	return r.ApiService.FilesOpenLocalEditorCreateExecute(r)
}

/*
FilesOpenLocalEditorCreate Create a local editor

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest
*/
func (a *FilesOpenLocalEditorAPIService) FilesOpenLocalEditorCreate(ctx context.Context) FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest {
	return FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FilesOpenLocalEditorCreate200Response
func (a *FilesOpenLocalEditorAPIService) FilesOpenLocalEditorCreateExecute(r FilesOpenLocalEditorAPIFilesOpenLocalEditorCreateRequest) (*FilesOpenLocalEditorCreate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilesOpenLocalEditorCreate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOpenLocalEditorAPIService.FilesOpenLocalEditorCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files/api/v1/openlocaleditor"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.path == nil {
		return localVarReturnValue, nil, reportError("path is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "path", r.path, "")
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

type FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest struct {
	ctx context.Context
	ApiService FilesOpenLocalEditorAPI
	path *string
	token string
	oCSAPIRequest *string
}

// Path of the file
func (r FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest) Path(path string) FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest {
	r.path = &path
	return r
}

func (r FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest) OCSAPIRequest(oCSAPIRequest string) FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest) Execute() (*FilesOpenLocalEditorValidate200Response, *http.Response, error) {
	return r.ApiService.FilesOpenLocalEditorValidateExecute(r)
}

/*
FilesOpenLocalEditorValidate Validate a local editor

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param token Token of the local editor
 @return FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest
*/
func (a *FilesOpenLocalEditorAPIService) FilesOpenLocalEditorValidate(ctx context.Context, token string) FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest {
	return FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest{
		ApiService: a,
		ctx: ctx,
		token: token,
	}
}

// Execute executes the request
//  @return FilesOpenLocalEditorValidate200Response
func (a *FilesOpenLocalEditorAPIService) FilesOpenLocalEditorValidateExecute(r FilesOpenLocalEditorAPIFilesOpenLocalEditorValidateRequest) (*FilesOpenLocalEditorValidate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilesOpenLocalEditorValidate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesOpenLocalEditorAPIService.FilesOpenLocalEditorValidate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files/api/v1/openlocaleditor/{token}"
	localVarPath = strings.Replace(localVarPath, "{"+"token"+"}", url.PathEscape(parameterValueToString(r.token, "token")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.path == nil {
		return localVarReturnValue, nil, reportError("path is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "path", r.path, "")
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
