# coding: utf-8

"""
    nextcloud

    Nextcloud APIs  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


import re  # noqa: F401
import io
import warnings

from pydantic import validate_arguments, ValidationError
from typing_extensions import Annotated

from pydantic import Field, StrictStr

from openapi_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
from openapi_client.models.core_whats_new_get200_response import CoreWhatsNewGet200Response

from openapi_client.api_client import ApiClient
from openapi_client.api_response import ApiResponse
from openapi_client.exceptions import (  # noqa: F401
    ApiTypeError,
    ApiValueError
)


class CoreWhatsNewApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient.get_default()
        self.api_client = api_client

    @validate_arguments
    def core_whats_new_dismiss(self, version : Annotated[StrictStr, Field(..., description="Version to dismiss the changes for")], ocs_api_request : StrictStr, **kwargs) -> CoreWhatsNewDismiss200Response:  # noqa: E501
        """Dismiss the changes  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_whats_new_dismiss(version, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param version: Version to dismiss the changes for (required)
        :type version: str
        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param async_req: Whether to execute the request asynchronously.
        :type async_req: bool, optional
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: Returns the result object.
                 If the method is called asynchronously,
                 returns the request thread.
        :rtype: CoreWhatsNewDismiss200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the core_whats_new_dismiss_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.core_whats_new_dismiss_with_http_info(version, ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def core_whats_new_dismiss_with_http_info(self, version : Annotated[StrictStr, Field(..., description="Version to dismiss the changes for")], ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """Dismiss the changes  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_whats_new_dismiss_with_http_info(version, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param version: Version to dismiss the changes for (required)
        :type version: str
        :param ocs_api_request: (required)
        :type ocs_api_request: str
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
        :rtype: tuple(CoreWhatsNewDismiss200Response, status_code(int), headers(HTTPHeaderDict))
        """

        _params = locals()

        _all_params = [
            'version',
            'ocs_api_request'
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
                    " to method core_whats_new_dismiss" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}

        # process the query parameters
        _query_params = []
        if _params.get('version') is not None:  # noqa: E501
            _query_params.append(('version', _params['version']))

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
            '200': "CoreWhatsNewDismiss200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/core/whatsnew', 'POST',
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
    def core_whats_new_get(self, ocs_api_request : StrictStr, **kwargs) -> CoreWhatsNewGet200Response:  # noqa: E501
        """Get the changes  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_whats_new_get(ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param async_req: Whether to execute the request asynchronously.
        :type async_req: bool, optional
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: Returns the result object.
                 If the method is called asynchronously,
                 returns the request thread.
        :rtype: CoreWhatsNewGet200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the core_whats_new_get_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.core_whats_new_get_with_http_info(ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def core_whats_new_get_with_http_info(self, ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """Get the changes  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_whats_new_get_with_http_info(ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param ocs_api_request: (required)
        :type ocs_api_request: str
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
        :rtype: tuple(CoreWhatsNewGet200Response, status_code(int), headers(HTTPHeaderDict))
        """

        _params = locals()

        _all_params = [
            'ocs_api_request'
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
                    " to method core_whats_new_get" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}

        # process the query parameters
        _query_params = []
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
            '200': "CoreWhatsNewGet200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/core/whatsnew', 'GET',
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