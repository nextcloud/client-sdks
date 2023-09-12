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


type CoreTranslationApiAPI interface {

	/*
	CoreTranslationApiLanguages Get the list of supported languages

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CoreTranslationApiAPICoreTranslationApiLanguagesRequest
	*/
	CoreTranslationApiLanguages(ctx context.Context) CoreTranslationApiAPICoreTranslationApiLanguagesRequest

	// CoreTranslationApiLanguagesExecute executes the request
	//  @return CoreTranslationApiLanguages200Response
	CoreTranslationApiLanguagesExecute(r CoreTranslationApiAPICoreTranslationApiLanguagesRequest) (*CoreTranslationApiLanguages200Response, *http.Response, error)

	/*
	CoreTranslationApiTranslate Translate a text

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return CoreTranslationApiAPICoreTranslationApiTranslateRequest
	*/
	CoreTranslationApiTranslate(ctx context.Context) CoreTranslationApiAPICoreTranslationApiTranslateRequest

	// CoreTranslationApiTranslateExecute executes the request
	//  @return CoreTranslationApiTranslate200Response
	CoreTranslationApiTranslateExecute(r CoreTranslationApiAPICoreTranslationApiTranslateRequest) (*CoreTranslationApiTranslate200Response, *http.Response, error)
}

// CoreTranslationApiAPIService CoreTranslationApiAPI service
type CoreTranslationApiAPIService service

type CoreTranslationApiAPICoreTranslationApiLanguagesRequest struct {
	ctx context.Context
	ApiService CoreTranslationApiAPI
	oCSAPIRequest *string
}

func (r CoreTranslationApiAPICoreTranslationApiLanguagesRequest) OCSAPIRequest(oCSAPIRequest string) CoreTranslationApiAPICoreTranslationApiLanguagesRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r CoreTranslationApiAPICoreTranslationApiLanguagesRequest) Execute() (*CoreTranslationApiLanguages200Response, *http.Response, error) {
	return r.ApiService.CoreTranslationApiLanguagesExecute(r)
}

/*
CoreTranslationApiLanguages Get the list of supported languages

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CoreTranslationApiAPICoreTranslationApiLanguagesRequest
*/
func (a *CoreTranslationApiAPIService) CoreTranslationApiLanguages(ctx context.Context) CoreTranslationApiAPICoreTranslationApiLanguagesRequest {
	return CoreTranslationApiAPICoreTranslationApiLanguagesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreTranslationApiLanguages200Response
func (a *CoreTranslationApiAPIService) CoreTranslationApiLanguagesExecute(r CoreTranslationApiAPICoreTranslationApiLanguagesRequest) (*CoreTranslationApiLanguages200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreTranslationApiLanguages200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreTranslationApiAPIService.CoreTranslationApiLanguages")
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

type CoreTranslationApiAPICoreTranslationApiTranslateRequest struct {
	ctx context.Context
	ApiService CoreTranslationApiAPI
	text *string
	toLanguage *string
	oCSAPIRequest *string
	fromLanguage *string
}

// Text to be translated
func (r CoreTranslationApiAPICoreTranslationApiTranslateRequest) Text(text string) CoreTranslationApiAPICoreTranslationApiTranslateRequest {
	r.text = &text
	return r
}

// Language to translate to
func (r CoreTranslationApiAPICoreTranslationApiTranslateRequest) ToLanguage(toLanguage string) CoreTranslationApiAPICoreTranslationApiTranslateRequest {
	r.toLanguage = &toLanguage
	return r
}

func (r CoreTranslationApiAPICoreTranslationApiTranslateRequest) OCSAPIRequest(oCSAPIRequest string) CoreTranslationApiAPICoreTranslationApiTranslateRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

// Language to translate from
func (r CoreTranslationApiAPICoreTranslationApiTranslateRequest) FromLanguage(fromLanguage string) CoreTranslationApiAPICoreTranslationApiTranslateRequest {
	r.fromLanguage = &fromLanguage
	return r
}

func (r CoreTranslationApiAPICoreTranslationApiTranslateRequest) Execute() (*CoreTranslationApiTranslate200Response, *http.Response, error) {
	return r.ApiService.CoreTranslationApiTranslateExecute(r)
}

/*
CoreTranslationApiTranslate Translate a text

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return CoreTranslationApiAPICoreTranslationApiTranslateRequest
*/
func (a *CoreTranslationApiAPIService) CoreTranslationApiTranslate(ctx context.Context) CoreTranslationApiAPICoreTranslationApiTranslateRequest {
	return CoreTranslationApiAPICoreTranslationApiTranslateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreTranslationApiTranslate200Response
func (a *CoreTranslationApiAPIService) CoreTranslationApiTranslateExecute(r CoreTranslationApiAPICoreTranslationApiTranslateRequest) (*CoreTranslationApiTranslate200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreTranslationApiTranslate200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CoreTranslationApiAPIService.CoreTranslationApiTranslate")
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