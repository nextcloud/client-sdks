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


type CoreWhatsNewAPI interface {

	/*
	CoreWhatsNewDismiss Dismiss the changes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CoreWhatsNewAPICoreWhatsNewDismissRequest
	*/
	CoreWhatsNewDismiss(ctx context.Context) CoreWhatsNewAPICoreWhatsNewDismissRequest

	// CoreWhatsNewDismissExecute executes the request
	//  @return CoreWhatsNewDismiss200Response
	CoreWhatsNewDismissExecute(r CoreWhatsNewAPICoreWhatsNewDismissRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error)

	/*
	CoreWhatsNewGet Get the changes

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CoreWhatsNewAPICoreWhatsNewGetRequest
	*/
	CoreWhatsNewGet(ctx context.Context) CoreWhatsNewAPICoreWhatsNewGetRequest

	// CoreWhatsNewGetExecute executes the request
	//  @return CoreWhatsNewGet200Response
	CoreWhatsNewGetExecute(r CoreWhatsNewAPICoreWhatsNewGetRequest) (*CoreWhatsNewGet200Response, *http.Response, error)
}

// CoreWhatsNewAPIService CoreWhatsNewAPI service
type CoreWhatsNewAPIService service

type CoreWhatsNewAPICoreWhatsNewDismissRequest struct {
	ctx context.Context
	ApiService CoreWhatsNewAPI
	version *string
	oCSAPIRequest *string
}

// Version to dismiss the changes for
func (r CoreWhatsNewAPICoreWhatsNewDismissRequest) Version(version string) CoreWhatsNewAPICoreWhatsNewDismissRequest {
	r.version = &version
	return r
}

func (r CoreWhatsNewAPICoreWhatsNewDismissRequest) OCSAPIRequest(oCSAPIRequest string) CoreWhatsNewAPICoreWhatsNewDismissRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r CoreWhatsNewAPICoreWhatsNewDismissRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.CoreWhatsNewDismissExecute(r)
}

/*
CoreWhatsNewDismiss Dismiss the changes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CoreWhatsNewAPICoreWhatsNewDismissRequest
*/
func (a *CoreWhatsNewAPIService) CoreWhatsNewDismiss(ctx context.Context) CoreWhatsNewAPICoreWhatsNewDismissRequest {
	return CoreWhatsNewAPICoreWhatsNewDismissRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *CoreWhatsNewAPIService) CoreWhatsNewDismissExecute(r CoreWhatsNewAPICoreWhatsNewDismissRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreWhatsNewAPIService.CoreWhatsNewDismiss")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/core/whatsnew"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.version == nil {
		return localVarReturnValue, nil, reportError("version is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "version", r.version, "")
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

type CoreWhatsNewAPICoreWhatsNewGetRequest struct {
	ctx context.Context
	ApiService CoreWhatsNewAPI
	oCSAPIRequest *string
}

func (r CoreWhatsNewAPICoreWhatsNewGetRequest) OCSAPIRequest(oCSAPIRequest string) CoreWhatsNewAPICoreWhatsNewGetRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r CoreWhatsNewAPICoreWhatsNewGetRequest) Execute() (*CoreWhatsNewGet200Response, *http.Response, error) {
	return r.ApiService.CoreWhatsNewGetExecute(r)
}

/*
CoreWhatsNewGet Get the changes

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CoreWhatsNewAPICoreWhatsNewGetRequest
*/
func (a *CoreWhatsNewAPIService) CoreWhatsNewGet(ctx context.Context) CoreWhatsNewAPICoreWhatsNewGetRequest {
	return CoreWhatsNewAPICoreWhatsNewGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreWhatsNewGet200Response
func (a *CoreWhatsNewAPIService) CoreWhatsNewGetExecute(r CoreWhatsNewAPICoreWhatsNewGetRequest) (*CoreWhatsNewGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreWhatsNewAPIService.CoreWhatsNewGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/core/whatsnew"

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