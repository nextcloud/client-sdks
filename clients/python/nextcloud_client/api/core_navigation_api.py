# coding: utf-8

"""
    nextcloud

    Nextcloud APIs

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import re  # noqa: F401
import io
import warnings

from pydantic import validate_arguments, ValidationError
from typing_extensions import Annotated

from pydantic import Field, StrictInt, StrictStr

from typing import Optional

from nextcloud_client.models.core_navigation_get_apps_navigation200_response import CoreNavigationGetAppsNavigation200Response

from nextcloud_client.api_client import ApiClient
from nextcloud_client.api_response import ApiResponse
from nextcloud_client.exceptions import (  # noqa: F401
    ApiTypeError,
    ApiValueError
)


class CoreNavigationApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient.get_default()
        self.api_client = api_client

    @validate_arguments
    def core_navigation_get_apps_navigation(self, ocs_api_request : StrictStr, absolute : Annotated[Optional[StrictInt], Field(description="Rewrite URLs to absolute ones")] = None, **kwargs) -> CoreNavigationGetAppsNavigation200Response:  # noqa: E501
        """Get the apps navigation  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_navigation_get_apps_navigation(ocs_api_request, absolute, async_req=True)
        >>> result = thread.get()

        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param absolute: Rewrite URLs to absolute ones
        :type absolute: int
        :param async_req: Whether to execute the request asynchronously.
        :type async_req: bool, optional
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: Returns the result object.
                 If the method is called asynchronously,
                 returns the request thread.
        :rtype: CoreNavigationGetAppsNavigation200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the core_navigation_get_apps_navigation_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.core_navigation_get_apps_navigation_with_http_info(ocs_api_request, absolute, **kwargs)  # noqa: E501

    @validate_arguments
    def core_navigation_get_apps_navigation_with_http_info(self, ocs_api_request : StrictStr, absolute : Annotated[Optional[StrictInt], Field(description="Rewrite URLs to absolute ones")] = None, **kwargs) -> ApiResponse:  # noqa: E501
        """Get the apps navigation  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_navigation_get_apps_navigation_with_http_info(ocs_api_request, absolute, async_req=True)
        >>> result = thread.get()

        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param absolute: Rewrite URLs to absolute ones
        :type absolute: int
        :param async_req: Whether to execute the request asynchronously.
        :type async_req: bool, optional
        :param _preload_content: if False, the ApiResponse.data will
                                 be set to none and raw_data will store the 
                                 HTTP response body without reading/decoding.
                                 Default is True.
        :type _preload_content: bool, optional
        :param _return_http_data_only: response data instead of ApiResponse
                                       object with status code, headers, etc
        :type _return_http_data_only: bool, optional
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :param _request_auth: set to override the auth_settings for an a single
                              request; this effectively ignores the authentication
                              in the spec for a single request.
        :type _request_auth: dict, optional
        :type _content_type: string, optional: force content-type for the request
        :return: Returns the result object.
                 If the method is called asynchronously,
                 returns the request thread.
        :rtype: tuple(CoreNavigationGetAppsNavigation200Response, status_code(int), headers(HTTPHeaderDict))
        """

        _params = locals()

        _all_params = [
            'ocs_api_request',
            'absolute'
        ]
        _all_params.extend(
            [
                'async_req',
                '_return_http_data_only',
                '_preload_content',
                '_request_timeout',
                '_request_auth',
                '_content_type',
                '_headers'
            ]
        )

        # validate the arguments
        for _key, _val in _params['kwargs'].items():
            if _key not in _all_params:
                raise ApiTypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method core_navigation_get_apps_navigation" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}

        # process the query parameters
        _query_params = []
        if _params.get('absolute') is not None:  # noqa: E501
            _query_params.append(('absolute', _params['absolute']))

        # process the header parameters
        _header_params = dict(_params.get('_headers', {}))
        if _params['ocs_api_request']:
            _header_params['OCS-APIRequest'] = _params['ocs_api_request']

        # process the form parameters
        _form_params = []
        _files = {}
        # process the body parameter
        _body_params = None
        # set the HTTP header `Accept`
        _header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # authentication setting
        _auth_settings = ['basic_auth', 'bearer_auth']  # noqa: E501

        _response_types_map = {
            '200': "CoreNavigationGetAppsNavigation200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/core/navigation/apps', 'GET',
            _path_params,
            _query_params,
            _header_params,
            body=_body_params,
            post_params=_form_params,
            files=_files,
            response_types_map=_response_types_map,
            auth_settings=_auth_settings,
            async_req=_params.get('async_req'),
            _return_http_data_only=_params.get('_return_http_data_only'),  # noqa: E501
            _preload_content=_params.get('_preload_content', True),
            _request_timeout=_params.get('_request_timeout'),
            collection_formats=_collection_formats,
            _request_auth=_params.get('_request_auth'))

    @validate_arguments
    def core_navigation_get_settings_navigation(self, ocs_api_request : StrictStr, absolute : Annotated[Optional[StrictInt], Field(description="Rewrite URLs to absolute ones")] = None, **kwargs) -> CoreNavigationGetAppsNavigation200Response:  # noqa: E501
        """Get the settings navigation  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_navigation_get_settings_navigation(ocs_api_request, absolute, async_req=True)
        >>> result = thread.get()

        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param absolute: Rewrite URLs to absolute ones
        :type absolute: int
        :param async_req: Whether to execute the request asynchronously.
        :type async_req: bool, optional
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: Returns the result object.
                 If the method is called asynchronously,
                 returns the request thread.
        :rtype: CoreNavigationGetAppsNavigation200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the core_navigation_get_settings_navigation_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.core_navigation_get_settings_navigation_with_http_info(ocs_api_request, absolute, **kwargs)  # noqa: E501

    @validate_arguments
    def core_navigation_get_settings_navigation_with_http_info(self, ocs_api_request : StrictStr, absolute : Annotated[Optional[StrictInt], Field(description="Rewrite URLs to absolute ones")] = None, **kwargs) -> ApiResponse:  # noqa: E501
        """Get the settings navigation  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_navigation_get_settings_navigation_with_http_info(ocs_api_request, absolute, async_req=True)
        >>> result = thread.get()

        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param absolute: Rewrite URLs to absolute ones
        :type absolute: int
        :param async_req: Whether to execute the request asynchronously.
        :type async_req: bool, optional
        :param _preload_content: if False, the ApiResponse.data will
                                 be set to none and raw_data will store the 
                                 HTTP response body without reading/decoding.
                                 Default is True.
        :type _preload_content: bool, optional
        :param _return_http_data_only: response data instead of ApiResponse
                                       object with status code, headers, etc
        :type _return_http_data_only: bool, optional
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :param _request_auth: set to override the auth_settings for an a single
                              request; this effectively ignores the authentication
                              in the spec for a single request.
        :type _request_auth: dict, optional
        :type _content_type: string, optional: force content-type for the request
        :return: Returns the result object.
                 If the method is called asynchronously,
                 returns the request thread.
        :rtype: tuple(CoreNavigationGetAppsNavigation200Response, status_code(int), headers(HTTPHeaderDict))
        """

        _params = locals()

        _all_params = [
            'ocs_api_request',
            'absolute'
        ]
        _all_params.extend(
            [
                'async_req',
                '_return_http_data_only',
                '_preload_content',
                '_request_timeout',
                '_request_auth',
                '_content_type',
                '_headers'
            ]
        )

        # validate the arguments
        for _key, _val in _params['kwargs'].items():
            if _key not in _all_params:
                raise ApiTypeError(
                    "Got an unexpected keyword argument '%s'"
                    " to method core_navigation_get_settings_navigation" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}

        # process the query parameters
        _query_params = []
        if _params.get('absolute') is not None:  # noqa: E501
            _query_params.append(('absolute', _params['absolute']))

        # process the header parameters
        _header_params = dict(_params.get('_headers', {}))
        if _params['ocs_api_request']:
            _header_params['OCS-APIRequest'] = _params['ocs_api_request']

        # process the form parameters
        _form_params = []
        _files = {}
        # process the body parameter
        _body_params = None
        # set the HTTP header `Accept`
        _header_params['Accept'] = self.api_client.select_header_accept(
            ['application/json'])  # noqa: E501

        # authentication setting
        _auth_settings = ['basic_auth', 'bearer_auth']  # noqa: E501

        _response_types_map = {
            '200': "CoreNavigationGetAppsNavigation200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/core/navigation/settings', 'GET',
            _path_params,
            _query_params,
            _header_params,
            body=_body_params,
            post_params=_form_params,
            files=_files,
            response_types_map=_response_types_map,
            auth_settings=_auth_settings,
            async_req=_params.get('async_req'),
            _return_http_data_only=_params.get('_return_http_data_only'),  # noqa: E501
            _preload_content=_params.get('_preload_content', True),
            _request_timeout=_params.get('_request_timeout'),
            collection_formats=_collection_formats,
            _request_auth=_params.get('_request_auth'))
