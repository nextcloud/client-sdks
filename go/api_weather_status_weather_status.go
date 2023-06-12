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


// WeatherStatusWeatherStatusApiService WeatherStatusWeatherStatusApi service
type WeatherStatusWeatherStatusApiService service

type ApiWeatherStatusWeatherStatusGetFavoritesRequest struct {
	ctx context.Context
	ApiService *WeatherStatusWeatherStatusApiService
	oCSAPIRequest *string
}

func (r ApiWeatherStatusWeatherStatusGetFavoritesRequest) OCSAPIRequest(oCSAPIRequest string) ApiWeatherStatusWeatherStatusGetFavoritesRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiWeatherStatusWeatherStatusGetFavoritesRequest) Execute() (*ProvisioningApiGroupsGetSubAdminsOfGroup200Response, *http.Response, error) {
	return r.ApiService.WeatherStatusWeatherStatusGetFavoritesExecute(r)
}

/*
WeatherStatusWeatherStatusGetFavorites Get favorites list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWeatherStatusWeatherStatusGetFavoritesRequest
*/
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusGetFavorites(ctx context.Context) ApiWeatherStatusWeatherStatusGetFavoritesRequest {
	return ApiWeatherStatusWeatherStatusGetFavoritesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ProvisioningApiGroupsGetSubAdminsOfGroup200Response
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusGetFavoritesExecute(r ApiWeatherStatusWeatherStatusGetFavoritesRequest) (*ProvisioningApiGroupsGetSubAdminsOfGroup200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ProvisioningApiGroupsGetSubAdminsOfGroup200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WeatherStatusWeatherStatusApiService.WeatherStatusWeatherStatusGetFavorites")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/weather_status/api/v1/favorites"

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

type ApiWeatherStatusWeatherStatusGetForecastRequest struct {
	ctx context.Context
	ApiService *WeatherStatusWeatherStatusApiService
	oCSAPIRequest *string
}

func (r ApiWeatherStatusWeatherStatusGetForecastRequest) OCSAPIRequest(oCSAPIRequest string) ApiWeatherStatusWeatherStatusGetForecastRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiWeatherStatusWeatherStatusGetForecastRequest) Execute() (*WeatherStatusWeatherStatusGetForecast200Response, *http.Response, error) {
	return r.ApiService.WeatherStatusWeatherStatusGetForecastExecute(r)
}

/*
WeatherStatusWeatherStatusGetForecast Get forecast for current location

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWeatherStatusWeatherStatusGetForecastRequest
*/
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusGetForecast(ctx context.Context) ApiWeatherStatusWeatherStatusGetForecastRequest {
	return ApiWeatherStatusWeatherStatusGetForecastRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WeatherStatusWeatherStatusGetForecast200Response
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusGetForecastExecute(r ApiWeatherStatusWeatherStatusGetForecastRequest) (*WeatherStatusWeatherStatusGetForecast200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WeatherStatusWeatherStatusGetForecast200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WeatherStatusWeatherStatusApiService.WeatherStatusWeatherStatusGetForecast")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/weather_status/api/v1/forecast"

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

type ApiWeatherStatusWeatherStatusGetLocationRequest struct {
	ctx context.Context
	ApiService *WeatherStatusWeatherStatusApiService
	oCSAPIRequest *string
}

func (r ApiWeatherStatusWeatherStatusGetLocationRequest) OCSAPIRequest(oCSAPIRequest string) ApiWeatherStatusWeatherStatusGetLocationRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiWeatherStatusWeatherStatusGetLocationRequest) Execute() (*WeatherStatusWeatherStatusGetLocation200Response, *http.Response, error) {
	return r.ApiService.WeatherStatusWeatherStatusGetLocationExecute(r)
}

/*
WeatherStatusWeatherStatusGetLocation Get stored user location

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWeatherStatusWeatherStatusGetLocationRequest
*/
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusGetLocation(ctx context.Context) ApiWeatherStatusWeatherStatusGetLocationRequest {
	return ApiWeatherStatusWeatherStatusGetLocationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WeatherStatusWeatherStatusGetLocation200Response
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusGetLocationExecute(r ApiWeatherStatusWeatherStatusGetLocationRequest) (*WeatherStatusWeatherStatusGetLocation200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WeatherStatusWeatherStatusGetLocation200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WeatherStatusWeatherStatusApiService.WeatherStatusWeatherStatusGetLocation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/weather_status/api/v1/location"

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

type ApiWeatherStatusWeatherStatusSetFavoritesRequest struct {
	ctx context.Context
	ApiService *WeatherStatusWeatherStatusApiService
	favorites *string
	oCSAPIRequest *string
}

// Favorite addresses
func (r ApiWeatherStatusWeatherStatusSetFavoritesRequest) Favorites(favorites string) ApiWeatherStatusWeatherStatusSetFavoritesRequest {
	r.favorites = &favorites
	return r
}

func (r ApiWeatherStatusWeatherStatusSetFavoritesRequest) OCSAPIRequest(oCSAPIRequest string) ApiWeatherStatusWeatherStatusSetFavoritesRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiWeatherStatusWeatherStatusSetFavoritesRequest) Execute() (*CoreReferenceApiTouchProvider200Response, *http.Response, error) {
	return r.ApiService.WeatherStatusWeatherStatusSetFavoritesExecute(r)
}

/*
WeatherStatusWeatherStatusSetFavorites Set favorites list

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWeatherStatusWeatherStatusSetFavoritesRequest
*/
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusSetFavorites(ctx context.Context) ApiWeatherStatusWeatherStatusSetFavoritesRequest {
	return ApiWeatherStatusWeatherStatusSetFavoritesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreReferenceApiTouchProvider200Response
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusSetFavoritesExecute(r ApiWeatherStatusWeatherStatusSetFavoritesRequest) (*CoreReferenceApiTouchProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreReferenceApiTouchProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WeatherStatusWeatherStatusApiService.WeatherStatusWeatherStatusSetFavorites")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/weather_status/api/v1/favorites"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.favorites == nil {
		return localVarReturnValue, nil, reportError("favorites is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "favorites", r.favorites, "")
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

type ApiWeatherStatusWeatherStatusSetLocationRequest struct {
	ctx context.Context
	ApiService *WeatherStatusWeatherStatusApiService
	oCSAPIRequest *string
	address *string
	lat *float32
	lon *float32
}

func (r ApiWeatherStatusWeatherStatusSetLocationRequest) OCSAPIRequest(oCSAPIRequest string) ApiWeatherStatusWeatherStatusSetLocationRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

// Any approximative or exact address
func (r ApiWeatherStatusWeatherStatusSetLocationRequest) Address(address string) ApiWeatherStatusWeatherStatusSetLocationRequest {
	r.address = &address
	return r
}

// Latitude in decimal degree format
func (r ApiWeatherStatusWeatherStatusSetLocationRequest) Lat(lat float32) ApiWeatherStatusWeatherStatusSetLocationRequest {
	r.lat = &lat
	return r
}

// Longitude in decimal degree format
func (r ApiWeatherStatusWeatherStatusSetLocationRequest) Lon(lon float32) ApiWeatherStatusWeatherStatusSetLocationRequest {
	r.lon = &lon
	return r
}

func (r ApiWeatherStatusWeatherStatusSetLocationRequest) Execute() (*WeatherStatusWeatherStatusUsePersonalAddress200Response, *http.Response, error) {
	return r.ApiService.WeatherStatusWeatherStatusSetLocationExecute(r)
}

/*
WeatherStatusWeatherStatusSetLocation Set address and resolve it to get coordinates or directly set coordinates and get address with reverse geocoding

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWeatherStatusWeatherStatusSetLocationRequest
*/
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusSetLocation(ctx context.Context) ApiWeatherStatusWeatherStatusSetLocationRequest {
	return ApiWeatherStatusWeatherStatusSetLocationRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WeatherStatusWeatherStatusUsePersonalAddress200Response
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusSetLocationExecute(r ApiWeatherStatusWeatherStatusSetLocationRequest) (*WeatherStatusWeatherStatusUsePersonalAddress200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WeatherStatusWeatherStatusUsePersonalAddress200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WeatherStatusWeatherStatusApiService.WeatherStatusWeatherStatusSetLocation")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/weather_status/api/v1/location"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	if r.address != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "address", r.address, "")
	}
	if r.lat != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lat", r.lat, "")
	}
	if r.lon != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lon", r.lon, "")
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

type ApiWeatherStatusWeatherStatusSetModeRequest struct {
	ctx context.Context
	ApiService *WeatherStatusWeatherStatusApiService
	mode *int64
	oCSAPIRequest *string
}

// New mode
func (r ApiWeatherStatusWeatherStatusSetModeRequest) Mode(mode int64) ApiWeatherStatusWeatherStatusSetModeRequest {
	r.mode = &mode
	return r
}

func (r ApiWeatherStatusWeatherStatusSetModeRequest) OCSAPIRequest(oCSAPIRequest string) ApiWeatherStatusWeatherStatusSetModeRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiWeatherStatusWeatherStatusSetModeRequest) Execute() (*CoreReferenceApiTouchProvider200Response, *http.Response, error) {
	return r.ApiService.WeatherStatusWeatherStatusSetModeExecute(r)
}

/*
WeatherStatusWeatherStatusSetMode Change the weather status mode. There are currently 2 modes: - ask the browser - use the user defined address

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWeatherStatusWeatherStatusSetModeRequest
*/
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusSetMode(ctx context.Context) ApiWeatherStatusWeatherStatusSetModeRequest {
	return ApiWeatherStatusWeatherStatusSetModeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CoreReferenceApiTouchProvider200Response
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusSetModeExecute(r ApiWeatherStatusWeatherStatusSetModeRequest) (*CoreReferenceApiTouchProvider200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CoreReferenceApiTouchProvider200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WeatherStatusWeatherStatusApiService.WeatherStatusWeatherStatusSetMode")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/weather_status/api/v1/mode"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.mode == nil {
		return localVarReturnValue, nil, reportError("mode is required and must be specified")
	}
	if r.oCSAPIRequest == nil {
		return localVarReturnValue, nil, reportError("oCSAPIRequest is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "mode", r.mode, "")
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

type ApiWeatherStatusWeatherStatusUsePersonalAddressRequest struct {
	ctx context.Context
	ApiService *WeatherStatusWeatherStatusApiService
	oCSAPIRequest *string
}

func (r ApiWeatherStatusWeatherStatusUsePersonalAddressRequest) OCSAPIRequest(oCSAPIRequest string) ApiWeatherStatusWeatherStatusUsePersonalAddressRequest {
	r.oCSAPIRequest = &oCSAPIRequest
	return r
}

func (r ApiWeatherStatusWeatherStatusUsePersonalAddressRequest) Execute() (*WeatherStatusWeatherStatusUsePersonalAddress200Response, *http.Response, error) {
	return r.ApiService.WeatherStatusWeatherStatusUsePersonalAddressExecute(r)
}

/*
WeatherStatusWeatherStatusUsePersonalAddress Try to use the address set in user personal settings as weather location

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiWeatherStatusWeatherStatusUsePersonalAddressRequest
*/
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusUsePersonalAddress(ctx context.Context) ApiWeatherStatusWeatherStatusUsePersonalAddressRequest {
	return ApiWeatherStatusWeatherStatusUsePersonalAddressRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WeatherStatusWeatherStatusUsePersonalAddress200Response
func (a *WeatherStatusWeatherStatusApiService) WeatherStatusWeatherStatusUsePersonalAddressExecute(r ApiWeatherStatusWeatherStatusUsePersonalAddressRequest) (*WeatherStatusWeatherStatusUsePersonalAddress200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WeatherStatusWeatherStatusUsePersonalAddress200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "WeatherStatusWeatherStatusApiService.WeatherStatusWeatherStatusUsePersonalAddress")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/ocs/v2.php/apps/weather_status/api/v1/use-personal"

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
