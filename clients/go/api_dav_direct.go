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


type DavDirectAPI interface {

	/*
	DavDirectGetUrl Get a direct link to a file

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return DavDirectAPIDavDirectGetUrlRequest
	*/
	DavDirectGetUrl(ctx context.Context) DavDirectAPIDavDirectGetUrlRequest

	// DavDirectGetUrlExecute executes the request
	//  @return DavDirectGetUrl200Response
	DavDirectGetUrlExecute(r DavDirectAPIDavDirectGetUrlRequest) (*DavDirectGetUrl200Response, *http.Response, error)
}

// DavDirectAPIService DavDirectAPI service
type DavDirectAPIService service

type DavDirectAPIDavDirectGetUrlRequest struct {
	ctx context.Context
	ApiService DavDirectAPI
	fileId *int64
	oCSAPIRequest *string
	expirationTime *int64
}

// ID of the file
func (r DavDirectAPIDavDirectGetUrlRequest) FileId(fileId int64) DavDirectAPIDavDirectGetUrlRequest {
	r.fileId = &fileId
	return r
}

func (r DavDirectAPIDavDirectGetUrlRequest) OCSAPIRequest(oCSAPIRequest string) DavDirectAPIDavDirectGetUrlRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

// Duration until the link expires
func (r DavDirectAPIDavDirectGetUrlRequest) ExpirationTime(expirationTime int64) DavDirectAPIDavDirectGetUrlRequest {
	r.expirationTime = &expirationTime
	return r
}

func (r DavDirectAPIDavDirectGetUrlRequest) Execute() (*DavDirectGetUrl200Response, *http.Response, error) {
	return r.ApiService.DavDirectGetUrlExecute(r)
}

/*
DavDirectGetUrl Get a direct link to a file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return DavDirectAPIDavDirectGetUrlRequest
*/
func (a *DavDirectAPIService) DavDirectGetUrl(ctx context.Context) DavDirectAPIDavDirectGetUrlRequest {
	return DavDirectAPIDavDirectGetUrlRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DavDirectGetUrl200Response
func (a *DavDirectAPIService) DavDirectGetUrlExecute(r DavDirectAPIDavDirectGetUrlRequest) (*DavDirectGetUrl200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DavDirectGetUrl200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DavDirectAPIService.DavDirectGetUrl")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/dav/api/v1/direct"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fileId == nil {
		return localVarReturnValue, nil, reportError("fileId is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fileId", r.fileId, "")
	if r.expirationTime != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expirationTime", r.expirationTime, "")
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
