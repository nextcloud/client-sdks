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


type FilesSharingRemoteAPI interface {

	/*
	FilesSharingRemoteAcceptShare Accept a remote share

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the share
	@return FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest
	*/
	FilesSharingRemoteAcceptShare(ctx context.Context, id int64) FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest

	// FilesSharingRemoteAcceptShareExecute executes the request
	//  @return CoreWhatsNewDismiss200Response
	FilesSharingRemoteAcceptShareExecute(r FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error)

	/*
	FilesSharingRemoteDeclineShare Decline a remote share

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the share
	@return FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest
	*/
	FilesSharingRemoteDeclineShare(ctx context.Context, id int64) FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest

	// FilesSharingRemoteDeclineShareExecute executes the request
	//  @return CoreWhatsNewDismiss200Response
	FilesSharingRemoteDeclineShareExecute(r FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error)

	/*
	FilesSharingRemoteGetOpenShares Get list of pending remote shares

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest
	*/
	FilesSharingRemoteGetOpenShares(ctx context.Context) FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest

	// FilesSharingRemoteGetOpenSharesExecute executes the request
	//  @return FilesSharingRemoteGetShares200Response
	FilesSharingRemoteGetOpenSharesExecute(r FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest) (*FilesSharingRemoteGetShares200Response, *http.Response, error)

	/*
	FilesSharingRemoteGetShare Get info of a remote share

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the share
	@return FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest
	*/
	FilesSharingRemoteGetShare(ctx context.Context, id int64) FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest

	// FilesSharingRemoteGetShareExecute executes the request
	//  @return FilesSharingRemoteGetShare200Response
	FilesSharingRemoteGetShareExecute(r FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest) (*FilesSharingRemoteGetShare200Response, *http.Response, error)

	/*
	FilesSharingRemoteGetShares Get a list of accepted remote shares

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest
	*/
	FilesSharingRemoteGetShares(ctx context.Context) FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest

	// FilesSharingRemoteGetSharesExecute executes the request
	//  @return FilesSharingRemoteGetShares200Response
	FilesSharingRemoteGetSharesExecute(r FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest) (*FilesSharingRemoteGetShares200Response, *http.Response, error)

	/*
	FilesSharingRemoteUnshare Unshare a remote share

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the share
	@return FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest
	*/
	FilesSharingRemoteUnshare(ctx context.Context, id int64) FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest

	// FilesSharingRemoteUnshareExecute executes the request
	//  @return CoreWhatsNewDismiss200Response
	FilesSharingRemoteUnshareExecute(r FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error)
}

// FilesSharingRemoteAPIService FilesSharingRemoteAPI service
type FilesSharingRemoteAPIService service

type FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest struct {
	ctx context.Context
	ApiService FilesSharingRemoteAPI
	id int64
	oCSAPIRequest *string
}

func (r FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest) OCSAPIRequest(oCSAPIRequest string) FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.FilesSharingRemoteAcceptShareExecute(r)
}

/*
FilesSharingRemoteAcceptShare Accept a remote share

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the share
 @return FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest
*/
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteAcceptShare(ctx context.Context, id int64) FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest {
	return FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteAcceptShareExecute(r FilesSharingRemoteAPIFilesSharingRemoteAcceptShareRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingRemoteAPIService.FilesSharingRemoteAcceptShare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest struct {
	ctx context.Context
	ApiService FilesSharingRemoteAPI
	id int64
	oCSAPIRequest *string
}

func (r FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest) OCSAPIRequest(oCSAPIRequest string) FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.FilesSharingRemoteDeclineShareExecute(r)
}

/*
FilesSharingRemoteDeclineShare Decline a remote share

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the share
 @return FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest
*/
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteDeclineShare(ctx context.Context, id int64) FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest {
	return FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteDeclineShareExecute(r FilesSharingRemoteAPIFilesSharingRemoteDeclineShareRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingRemoteAPIService.FilesSharingRemoteDeclineShare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest struct {
	ctx context.Context
	ApiService FilesSharingRemoteAPI
	oCSAPIRequest *string
}

func (r FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest) OCSAPIRequest(oCSAPIRequest string) FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest) Execute() (*FilesSharingRemoteGetShares200Response, *http.Response, error) {
	return r.ApiService.FilesSharingRemoteGetOpenSharesExecute(r)
}

/*
FilesSharingRemoteGetOpenShares Get list of pending remote shares

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest
*/
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteGetOpenShares(ctx context.Context) FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest {
	return FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FilesSharingRemoteGetShares200Response
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteGetOpenSharesExecute(r FilesSharingRemoteAPIFilesSharingRemoteGetOpenSharesRequest) (*FilesSharingRemoteGetShares200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilesSharingRemoteGetShares200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingRemoteAPIService.FilesSharingRemoteGetOpenShares")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files_sharing/api/v1/remote_shares/pending"

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

type FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest struct {
	ctx context.Context
	ApiService FilesSharingRemoteAPI
	id int64
	oCSAPIRequest *string
}

func (r FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest) OCSAPIRequest(oCSAPIRequest string) FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest) Execute() (*FilesSharingRemoteGetShare200Response, *http.Response, error) {
	return r.ApiService.FilesSharingRemoteGetShareExecute(r)
}

/*
FilesSharingRemoteGetShare Get info of a remote share

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the share
 @return FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest
*/
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteGetShare(ctx context.Context, id int64) FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest {
	return FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return FilesSharingRemoteGetShare200Response
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteGetShareExecute(r FilesSharingRemoteAPIFilesSharingRemoteGetShareRequest) (*FilesSharingRemoteGetShare200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilesSharingRemoteGetShare200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingRemoteAPIService.FilesSharingRemoteGetShare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files_sharing/api/v1/remote_shares/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest struct {
	ctx context.Context
	ApiService FilesSharingRemoteAPI
	oCSAPIRequest *string
}

func (r FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest) OCSAPIRequest(oCSAPIRequest string) FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest) Execute() (*FilesSharingRemoteGetShares200Response, *http.Response, error) {
	return r.ApiService.FilesSharingRemoteGetSharesExecute(r)
}

/*
FilesSharingRemoteGetShares Get a list of accepted remote shares

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest
*/
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteGetShares(ctx context.Context) FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest {
	return FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return FilesSharingRemoteGetShares200Response
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteGetSharesExecute(r FilesSharingRemoteAPIFilesSharingRemoteGetSharesRequest) (*FilesSharingRemoteGetShares200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *FilesSharingRemoteGetShares200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingRemoteAPIService.FilesSharingRemoteGetShares")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files_sharing/api/v1/remote_shares"

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

type FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest struct {
	ctx context.Context
	ApiService FilesSharingRemoteAPI
	id int64
	oCSAPIRequest *string
}

func (r FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest) OCSAPIRequest(oCSAPIRequest string) FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.FilesSharingRemoteUnshareExecute(r)
}

/*
FilesSharingRemoteUnshare Unshare a remote share

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the share
 @return FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest
*/
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteUnshare(ctx context.Context, id int64) FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest {
	return FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *FilesSharingRemoteAPIService) FilesSharingRemoteUnshareExecute(r FilesSharingRemoteAPIFilesSharingRemoteUnshareRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesSharingRemoteAPIService.FilesSharingRemoteUnshare")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files_sharing/api/v1/remote_shares/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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