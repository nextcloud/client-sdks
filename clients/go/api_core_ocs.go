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


type CoreOcsAPI interface {

	/*
	CoreOcsGetCapabilities Get the capabilities

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CoreOcsAPICoreOcsGetCapabilitiesRequest
	*/
	CoreOcsGetCapabilities(ctx context.Context) CoreOcsAPICoreOcsGetCapabilitiesRequest

	// CoreOcsGetCapabilitiesExecute executes the request
	//  @return CoreOcsGetCapabilities200Response
	CoreOcsGetCapabilitiesExecute(r CoreOcsAPICoreOcsGetCapabilitiesRequest) (*CoreOcsGetCapabilities200Response, *http.Response, error)
}

// CoreOcsAPIService CoreOcsAPI service
type CoreOcsAPIService service

type CoreOcsAPICoreOcsGetCapabilitiesRequest struct {
	ctx context.Context
	ApiService CoreOcsAPI
	oCSAPIRequest *string
}

func (r CoreOcsAPICoreOcsGetCapabilitiesRequest) OCSAPIRequest(oCSAPIRequest string) CoreOcsAPICoreOcsGetCapabilitiesRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r CoreOcsAPICoreOcsGetCapabilitiesRequest) Execute() (*CoreOcsGetCapabilities200Response, *http.Response, error) {
	return r.ApiService.CoreOcsGetCapabilitiesExecute(r)
}

/*
CoreOcsGetCapabilities Get the capabilities

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CoreOcsAPICoreOcsGetCapabilitiesRequest
*/
func (a *CoreOcsAPIService) CoreOcsGetCapabilities(ctx context.Context) CoreOcsAPICoreOcsGetCapabilitiesRequest {
	return CoreOcsAPICoreOcsGetCapabilitiesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreOcsGetCapabilities200Response
func (a *CoreOcsAPIService) CoreOcsGetCapabilitiesExecute(r CoreOcsAPICoreOcsGetCapabilitiesRequest) (*CoreOcsGetCapabilities200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreOcsGetCapabilities200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreOcsAPIService.CoreOcsGetCapabilities")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/cloud/capabilities"

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