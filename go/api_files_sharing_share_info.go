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


// FilesSharingShareInfoApiService FilesSharingShareInfoApi service
type FilesSharingShareInfoApiService service

type ApiFilesSharingShareInfoInfoRequest struct {
	ctx context.Context
	ApiService *FilesSharingShareInfoApiService
	t *string
	password *string
	dir *string
	depth *int64
}

// Token of the share
func (r ApiFilesSharingShareInfoInfoRequest) T(t string) ApiFilesSharingShareInfoInfoRequest {
	r.t = &t
	return r
}

// Password of the share
func (r ApiFilesSharingShareInfoInfoRequest) Password(password string) ApiFilesSharingShareInfoInfoRequest {
	r.password = &password
	return r
}

// Subdirectory to get info about
func (r ApiFilesSharingShareInfoInfoRequest) Dir(dir string) ApiFilesSharingShareInfoInfoRequest {
	r.dir = &dir
	return r
}

// Maximum depth to get info about
func (r ApiFilesSharingShareInfoInfoRequest) Depth(depth int64) ApiFilesSharingShareInfoInfoRequest {
	r.depth = &depth
	return r
}

func (r ApiFilesSharingShareInfoInfoRequest) Execute() (*FilesSharingShareInfo, *http.Response, error) {
	return r.ApiService.FilesSharingShareInfoInfoExecute(r)
}

/*
FilesSharingShareInfoInfo Get the info about a share

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFilesSharingShareInfoInfoRequest
*/
func (a *FilesSharingShareInfoApiService) FilesSharingShareInfoInfo(ctx context.Context) ApiFilesSharingShareInfoInfoRequest {
	return ApiFilesSharingShareInfoInfoRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FilesSharingShareInfo
func (a *FilesSharingShareInfoApiService) FilesSharingShareInfoInfoExecute(r ApiFilesSharingShareInfoInfoRequest) (*FilesSharingShareInfo, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilesSharingShareInfo
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingShareInfoApiService.FilesSharingShareInfoInfo")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/apps/files_sharing/shareinfo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.t == nil {
		return localVarReturnValue, nil, reportError("t is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "t", r.t, "")
	if r.password != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "password", r.password, "")
	}
	if r.dir != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dir", r.dir, "")
	}
	if r.depth != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "depth", r.depth, "")
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
