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
	"reflect"
)


type CoreAutoCompleteAPI interface {

	/*
	CoreAutoCompleteGet Autocomplete a query

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CoreAutoCompleteAPICoreAutoCompleteGetRequest
	*/
	CoreAutoCompleteGet(ctx context.Context) CoreAutoCompleteAPICoreAutoCompleteGetRequest

	// CoreAutoCompleteGetExecute executes the request
	//  @return CoreAutoCompleteGet200Response
	CoreAutoCompleteGetExecute(r CoreAutoCompleteAPICoreAutoCompleteGetRequest) (*CoreAutoCompleteGet200Response, *http.Response, error)
}

// CoreAutoCompleteAPIService CoreAutoCompleteAPI service
type CoreAutoCompleteAPIService service

type CoreAutoCompleteAPICoreAutoCompleteGetRequest struct {
	ctx context.Context
	ApiService CoreAutoCompleteAPI
	search *string
	oCSAPIRequest *string
	itemType *string
	itemId *string
	sorter *string
	shareTypes *[]int64
	limit *int64
}

// Text to search for
func (r CoreAutoCompleteAPICoreAutoCompleteGetRequest) Search(search string) CoreAutoCompleteAPICoreAutoCompleteGetRequest {
	r.search = &search
	return r
}

func (r CoreAutoCompleteAPICoreAutoCompleteGetRequest) OCSAPIRequest(oCSAPIRequest string) CoreAutoCompleteAPICoreAutoCompleteGetRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

// Type of the items to search for
func (r CoreAutoCompleteAPICoreAutoCompleteGetRequest) ItemType(itemType string) CoreAutoCompleteAPICoreAutoCompleteGetRequest {
	r.itemType = &itemType
	return r
}

// ID of the items to search for
func (r CoreAutoCompleteAPICoreAutoCompleteGetRequest) ItemId(itemId string) CoreAutoCompleteAPICoreAutoCompleteGetRequest {
	r.itemId = &itemId
	return r
}

// can be piped, top prio first, e.g.: \&quot;commenters|share-recipients\&quot;
func (r CoreAutoCompleteAPICoreAutoCompleteGetRequest) Sorter(sorter string) CoreAutoCompleteAPICoreAutoCompleteGetRequest {
	r.sorter = &sorter
	return r
}

// Types of shares to search for
func (r CoreAutoCompleteAPICoreAutoCompleteGetRequest) ShareTypes(shareTypes []int64) CoreAutoCompleteAPICoreAutoCompleteGetRequest {
	r.shareTypes = &shareTypes
	return r
}

// Maximum number of results to return
func (r CoreAutoCompleteAPICoreAutoCompleteGetRequest) Limit(limit int64) CoreAutoCompleteAPICoreAutoCompleteGetRequest {
	r.limit = &limit
	return r
}

func (r CoreAutoCompleteAPICoreAutoCompleteGetRequest) Execute() (*CoreAutoCompleteGet200Response, *http.Response, error) {
	return r.ApiService.CoreAutoCompleteGetExecute(r)
}

/*
CoreAutoCompleteGet Autocomplete a query

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CoreAutoCompleteAPICoreAutoCompleteGetRequest
*/
func (a *CoreAutoCompleteAPIService) CoreAutoCompleteGet(ctx context.Context) CoreAutoCompleteAPICoreAutoCompleteGetRequest {
	return CoreAutoCompleteAPICoreAutoCompleteGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreAutoCompleteGet200Response
func (a *CoreAutoCompleteAPIService) CoreAutoCompleteGetExecute(r CoreAutoCompleteAPICoreAutoCompleteGetRequest) (*CoreAutoCompleteGet200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreAutoCompleteGet200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreAutoCompleteAPIService.CoreAutoCompleteGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/core/autocomplete/get"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.search == nil {
		return localVarReturnValue, nil, reportError("search is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "")
	if r.itemType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemType", r.itemType, "")
	}
	if r.itemId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemId", r.itemId, "")
	}
	if r.sorter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sorter", r.sorter, "")
	}
	if r.shareTypes != nil {
		t := *r.shareTypes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "shareTypes[]", s.Index(i).Interface(), "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "shareTypes[]", t, "multi")
		}
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
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