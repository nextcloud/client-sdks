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


// CoreTranslationApiApiService CoreTranslationApiApi service
type CoreTranslationApiApiService service

type ApiCoreTranslationApiLanguagesRequest struct {
	ctx context.Context
	ApiService *CoreTranslationApiApiService
	oCSAPIRequest *string
}

func (r ApiCoreTranslationApiLanguagesRequest) OCSAPIRequest(oCSAPIRequest string) ApiCoreTranslationApiLanguagesRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiCoreTranslationApiLanguagesRequest) Execute() (*CoreTranslationApiLanguages200Response, *http.Response, error) {
	return r.ApiService.CoreTranslationApiLanguagesExecute(r)
}

/*
CoreTranslationApiLanguages Get the list of supported languages

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCoreTranslationApiLanguagesRequest
*/
func (a *CoreTranslationApiApiService) CoreTranslationApiLanguages(ctx context.Context) ApiCoreTranslationApiLanguagesRequest {
	return ApiCoreTranslationApiLanguagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreTranslationApiLanguages200Response
func (a *CoreTranslationApiApiService) CoreTranslationApiLanguagesExecute(r ApiCoreTranslationApiLanguagesRequest) (*CoreTranslationApiLanguages200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreTranslationApiLanguages200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreTranslationApiApiService.CoreTranslationApiLanguages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/translation/languages"

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

type ApiCoreTranslationApiTranslateRequest struct {
	ctx context.Context
	ApiService *CoreTranslationApiApiService
	text *string
	toLanguage *string
	oCSAPIRequest *string
	fromLanguage *string
}

// Text to be translated
func (r ApiCoreTranslationApiTranslateRequest) Text(text string) ApiCoreTranslationApiTranslateRequest {
	r.text = &text
	return r
}

// Language to translate to
func (r ApiCoreTranslationApiTranslateRequest) ToLanguage(toLanguage string) ApiCoreTranslationApiTranslateRequest {
	r.toLanguage = &toLanguage
	return r
}

func (r ApiCoreTranslationApiTranslateRequest) OCSAPIRequest(oCSAPIRequest string) ApiCoreTranslationApiTranslateRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

// Language to translate from
func (r ApiCoreTranslationApiTranslateRequest) FromLanguage(fromLanguage string) ApiCoreTranslationApiTranslateRequest {
	r.fromLanguage = &fromLanguage
	return r
}

func (r ApiCoreTranslationApiTranslateRequest) Execute() (*CoreTranslationApiTranslate200Response, *http.Response, error) {
	return r.ApiService.CoreTranslationApiTranslateExecute(r)
}

/*
CoreTranslationApiTranslate Translate a text

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCoreTranslationApiTranslateRequest
*/
func (a *CoreTranslationApiApiService) CoreTranslationApiTranslate(ctx context.Context) ApiCoreTranslationApiTranslateRequest {
	return ApiCoreTranslationApiTranslateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreTranslationApiTranslate200Response
func (a *CoreTranslationApiApiService) CoreTranslationApiTranslateExecute(r ApiCoreTranslationApiTranslateRequest) (*CoreTranslationApiTranslate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreTranslationApiTranslate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreTranslationApiApiService.CoreTranslationApiTranslate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/translation/translate"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.text == nil {
		return localVarReturnValue, nil, reportError("text is required and must be specified")
	}
	if r.toLanguage == nil {
		return localVarReturnValue, nil, reportError("toLanguage is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "text", r.text, "")
	if r.fromLanguage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fromLanguage", r.fromLanguage, "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "toLanguage", r.toLanguage, "")
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
