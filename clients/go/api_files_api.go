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


type FilesApiAPI interface {

	/*
	FilesApiGetThumbnail Gets a thumbnail of the specified file

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param x Width of the thumbnail
	@param y Height of the thumbnail
	@param file URL-encoded filename
	@return FilesApiAPIFilesApiGetThumbnailRequest
	*/
	FilesApiGetThumbnail(ctx context.Context, x int64, y int64, file string) FilesApiAPIFilesApiGetThumbnailRequest

	// FilesApiGetThumbnailExecute executes the request
	//  @return *os.File
	FilesApiGetThumbnailExecute(r FilesApiAPIFilesApiGetThumbnailRequest) (*os.File, *http.Response, error)

	/*
	FilesApiServiceWorker Get the service-worker Javascript for previews

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FilesApiAPIFilesApiServiceWorkerRequest
	*/
	FilesApiServiceWorker(ctx context.Context) FilesApiAPIFilesApiServiceWorkerRequest

	// FilesApiServiceWorkerExecute executes the request
	//  @return *os.File
	FilesApiServiceWorkerExecute(r FilesApiAPIFilesApiServiceWorkerRequest) (*os.File, *http.Response, error)
}

// FilesApiAPIService FilesApiAPI service
type FilesApiAPIService service

type FilesApiAPIFilesApiGetThumbnailRequest struct {
	ctx context.Context
	ApiService FilesApiAPI
	x int64
	y int64
	file string
}

func (r FilesApiAPIFilesApiGetThumbnailRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.FilesApiGetThumbnailExecute(r)
}

/*
FilesApiGetThumbnail Gets a thumbnail of the specified file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param x Width of the thumbnail
 @param y Height of the thumbnail
 @param file URL-encoded filename
 @return FilesApiAPIFilesApiGetThumbnailRequest
*/
func (a *FilesApiAPIService) FilesApiGetThumbnail(ctx context.Context, x int64, y int64, file string) FilesApiAPIFilesApiGetThumbnailRequest {
	return FilesApiAPIFilesApiGetThumbnailRequest{
		ApiService: a,
		ctx: ctx,
		x: x,
		y: y,
		file: file,
	}
}

// Execute executes the request
//  @return *os.File
func (a *FilesApiAPIService) FilesApiGetThumbnailExecute(r FilesApiAPIFilesApiGetThumbnailRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiAPIService.FilesApiGetThumbnail")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/apps/files/api/v1/thumbnail/{x}/{y}/{file}"
	localVarPath = strings.Replace(localVarPath, "{"+"x"+"}", url.PathEscape(parameterValueToString(r.x, "x")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"y"+"}", url.PathEscape(parameterValueToString(r.y, "y")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"file"+"}", url.PathEscape(parameterValueToString(r.file, "file")), -1)

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

type FilesApiAPIFilesApiServiceWorkerRequest struct {
	ctx context.Context
	ApiService FilesApiAPI
}

func (r FilesApiAPIFilesApiServiceWorkerRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.FilesApiServiceWorkerExecute(r)
}

/*
FilesApiServiceWorker Get the service-worker Javascript for previews

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FilesApiAPIFilesApiServiceWorkerRequest
*/
func (a *FilesApiAPIService) FilesApiServiceWorker(ctx context.Context) FilesApiAPIFilesApiServiceWorkerRequest {
	return FilesApiAPIFilesApiServiceWorkerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *FilesApiAPIService) FilesApiServiceWorkerExecute(r FilesApiAPIFilesApiServiceWorkerRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesApiAPIService.FilesApiServiceWorker")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/apps/files/preview-service-worker.js"

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
	localVarHTTPHeaderAccepts := []string{"application/javascript"}

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
