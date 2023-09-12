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


type CoreClientFlowLoginV2API interface {

	/*
	CoreClientFlowLoginV2Init Init a login flow

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CoreClientFlowLoginV2APICoreClientFlowLoginV2InitRequest
	*/
	CoreClientFlowLoginV2Init(ctx context.Context) CoreClientFlowLoginV2APICoreClientFlowLoginV2InitRequest

	// CoreClientFlowLoginV2InitExecute executes the request
	//  @return CoreLoginFlowV2
	CoreClientFlowLoginV2InitExecute(r CoreClientFlowLoginV2APICoreClientFlowLoginV2InitRequest) (*CoreLoginFlowV2, *http.Response, error)

	/*
	CoreClientFlowLoginV2Poll Poll the login flow credentials

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest
	*/
	CoreClientFlowLoginV2Poll(ctx context.Context) CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest

	// CoreClientFlowLoginV2PollExecute executes the request
	//  @return CoreLoginFlowV2Credentials
	CoreClientFlowLoginV2PollExecute(r CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest) (*CoreLoginFlowV2Credentials, *http.Response, error)
}

// CoreClientFlowLoginV2APIService CoreClientFlowLoginV2API service
type CoreClientFlowLoginV2APIService service

type CoreClientFlowLoginV2APICoreClientFlowLoginV2InitRequest struct {
	ctx context.Context
	ApiService CoreClientFlowLoginV2API
}

func (r CoreClientFlowLoginV2APICoreClientFlowLoginV2InitRequest) Execute() (*CoreLoginFlowV2, *http.Response, error) {
	return r.ApiService.CoreClientFlowLoginV2InitExecute(r)
}

/*
CoreClientFlowLoginV2Init Init a login flow

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CoreClientFlowLoginV2APICoreClientFlowLoginV2InitRequest
*/
func (a *CoreClientFlowLoginV2APIService) CoreClientFlowLoginV2Init(ctx context.Context) CoreClientFlowLoginV2APICoreClientFlowLoginV2InitRequest {
	return CoreClientFlowLoginV2APICoreClientFlowLoginV2InitRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreLoginFlowV2
func (a *CoreClientFlowLoginV2APIService) CoreClientFlowLoginV2InitExecute(r CoreClientFlowLoginV2APICoreClientFlowLoginV2InitRequest) (*CoreLoginFlowV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreLoginFlowV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreClientFlowLoginV2APIService.CoreClientFlowLoginV2Init")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/login/v2"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest struct {
	ctx context.Context
	ApiService CoreClientFlowLoginV2API
	token *string
}

// Token of the flow
func (r CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest) Token(token string) CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest {
	r.token = &token
	return r
}

func (r CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest) Execute() (*CoreLoginFlowV2Credentials, *http.Response, error) {
	return r.ApiService.CoreClientFlowLoginV2PollExecute(r)
}

/*
CoreClientFlowLoginV2Poll Poll the login flow credentials

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest
*/
func (a *CoreClientFlowLoginV2APIService) CoreClientFlowLoginV2Poll(ctx context.Context) CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest {
	return CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreLoginFlowV2Credentials
func (a *CoreClientFlowLoginV2APIService) CoreClientFlowLoginV2PollExecute(r CoreClientFlowLoginV2APICoreClientFlowLoginV2PollRequest) (*CoreLoginFlowV2Credentials, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreLoginFlowV2Credentials
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreClientFlowLoginV2APIService.CoreClientFlowLoginV2Poll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/login/v2/poll"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token == nil {
		return localVarReturnValue, nil, reportError("token is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "")
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