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
	"os"
)


type FilesTrashbinPreviewAPI interface {

	/*
	FilesTrashbinPreviewGetPreview Get the preview for a file

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest
	*/
	FilesTrashbinPreviewGetPreview(ctx context.Context) FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest

	// FilesTrashbinPreviewGetPreviewExecute executes the request
	//  @return *os.File
	FilesTrashbinPreviewGetPreviewExecute(r FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest) (*os.File, *http.Response, error)
}

// FilesTrashbinPreviewAPIService FilesTrashbinPreviewAPI service
type FilesTrashbinPreviewAPIService service

type FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest struct {
	ctx context.Context
	ApiService FilesTrashbinPreviewAPI
	fileId *int64
	x *int64
	y *int64
	a *int32
}

// ID of the file
func (r FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest) FileId(fileId int64) FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest {
	r.fileId = &fileId
	return r
}

// Width of the preview
func (r FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest) X(x int64) FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest {
	r.x = &x
	return r
}

// Height of the preview
func (r FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest) Y(y int64) FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest {
	r.y = &y
	return r
}

// Whether to not crop the preview
func (r FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest) A(a int32) FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest {
	r.a = &a
	return r
}

func (r FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest) Execute() (*os.File, *http.Response, error) {
	return r.ApiService.FilesTrashbinPreviewGetPreviewExecute(r)
}

/*
FilesTrashbinPreviewGetPreview Get the preview for a file

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest
*/
func (a *FilesTrashbinPreviewAPIService) FilesTrashbinPreviewGetPreview(ctx context.Context) FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest {
	return FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return *os.File
func (a *FilesTrashbinPreviewAPIService) FilesTrashbinPreviewGetPreviewExecute(r FilesTrashbinPreviewAPIFilesTrashbinPreviewGetPreviewRequest) (*os.File, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *os.File
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesTrashbinPreviewAPIService.FilesTrashbinPreviewGetPreview")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/index.php/apps/files_trashbin/preview"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fileId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fileId", r.fileId, "")
	}
	if r.x != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "x", r.x, "")
	}
	if r.y != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "y", r.y, "")
	}
	if r.a != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "a", r.a, "")
	}
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
