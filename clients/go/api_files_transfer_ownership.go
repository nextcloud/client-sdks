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


type FilesTransferOwnershipAPI interface {

	/*
	FilesTransferOwnershipAccept Accept an ownership transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the ownership transfer
	@return FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest
	*/
	FilesTransferOwnershipAccept(ctx context.Context, id int64) FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest

	// FilesTransferOwnershipAcceptExecute executes the request
	//  @return CoreWhatsNewDismiss200Response
	FilesTransferOwnershipAcceptExecute(r FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error)

	/*
	FilesTransferOwnershipReject Reject an ownership transfer

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id ID of the ownership transfer
	@return FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest
	*/
	FilesTransferOwnershipReject(ctx context.Context, id int64) FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest

	// FilesTransferOwnershipRejectExecute executes the request
	//  @return CoreWhatsNewDismiss200Response
	FilesTransferOwnershipRejectExecute(r FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error)

	/*
	FilesTransferOwnershipTransfer Transfer the ownership to another user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest
	*/
	FilesTransferOwnershipTransfer(ctx context.Context) FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest

	// FilesTransferOwnershipTransferExecute executes the request
	//  @return CoreWhatsNewDismiss200Response
	FilesTransferOwnershipTransferExecute(r FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error)
}

// FilesTransferOwnershipAPIService FilesTransferOwnershipAPI service
type FilesTransferOwnershipAPIService service

type FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest struct {
	ctx context.Context
	ApiService FilesTransferOwnershipAPI
	id int64
	oCSAPIRequest *string
}

func (r FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest) OCSAPIRequest(oCSAPIRequest string) FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.FilesTransferOwnershipAcceptExecute(r)
}

/*
FilesTransferOwnershipAccept Accept an ownership transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the ownership transfer
 @return FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest
*/
func (a *FilesTransferOwnershipAPIService) FilesTransferOwnershipAccept(ctx context.Context, id int64) FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest {
	return FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *FilesTransferOwnershipAPIService) FilesTransferOwnershipAcceptExecute(r FilesTransferOwnershipAPIFilesTransferOwnershipAcceptRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesTransferOwnershipAPIService.FilesTransferOwnershipAccept")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files/api/v1/transferownership/{id}"
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

type FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest struct {
	ctx context.Context
	ApiService FilesTransferOwnershipAPI
	id int64
	oCSAPIRequest *string
}

func (r FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest) OCSAPIRequest(oCSAPIRequest string) FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.FilesTransferOwnershipRejectExecute(r)
}

/*
FilesTransferOwnershipReject Reject an ownership transfer

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the ownership transfer
 @return FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest
*/
func (a *FilesTransferOwnershipAPIService) FilesTransferOwnershipReject(ctx context.Context, id int64) FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest {
	return FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *FilesTransferOwnershipAPIService) FilesTransferOwnershipRejectExecute(r FilesTransferOwnershipAPIFilesTransferOwnershipRejectRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesTransferOwnershipAPIService.FilesTransferOwnershipReject")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files/api/v1/transferownership/{id}"
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

type FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest struct {
	ctx context.Context
	ApiService FilesTransferOwnershipAPI
	recipient *string
	path *string
	oCSAPIRequest *string
}

// Username of the recipient
func (r FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest) Recipient(recipient string) FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest {
	r.recipient = &recipient
	return r
}

// Path of the file
func (r FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest) Path(path string) FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest {
	r.path = &path
	return r
}

func (r FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest) OCSAPIRequest(oCSAPIRequest string) FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest) Execute() (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	return r.ApiService.FilesTransferOwnershipTransferExecute(r)
}

/*
FilesTransferOwnershipTransfer Transfer the ownership to another user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest
*/
func (a *FilesTransferOwnershipAPIService) FilesTransferOwnershipTransfer(ctx context.Context) FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest {
	return FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreWhatsNewDismiss200Response
func (a *FilesTransferOwnershipAPIService) FilesTransferOwnershipTransferExecute(r FilesTransferOwnershipAPIFilesTransferOwnershipTransferRequest) (*CoreWhatsNewDismiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreWhatsNewDismiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "FilesTransferOwnershipAPIService.FilesTransferOwnershipTransfer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/files/api/v1/transferownership"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.recipient == nil {
		return localVarReturnValue, nil, reportError("recipient is required and must be specified")
	}
	if r.path == nil {
		return localVarReturnValue, nil, reportError("path is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "recipient", r.recipient, "")
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
