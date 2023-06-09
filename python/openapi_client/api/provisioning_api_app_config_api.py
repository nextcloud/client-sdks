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

from typing import Optional

from openapi_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
from openapi_client.models.provisioning_api_app_config_get_apps200_response import ProvisioningApiAppConfigGetApps200Response
from openapi_client.models.provisioning_api_app_config_get_value200_response import ProvisioningApiAppConfigGetValue200Response

from openapi_client.api_client import ApiClient
from openapi_client.api_response import ApiResponse
from openapi_client.exceptions import (  # noqa: F401
    ApiTypeError,
    ApiValueError
)


class ProvisioningApiAppConfigApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient.get_default()
        self.api_client = api_client

    @validate_arguments
    def provisioning_api_app_config_delete_key(self, app : Annotated[StrictStr, Field(..., description="ID of the app")], key : Annotated[StrictStr, Field(..., description="Key to delete")], ocs_api_request : StrictStr, **kwargs) -> CoreWhatsNewDismiss200Response:  # noqa: E501
        """Delete a config key of an app  # noqa: E501

        This endpoint requires admin access  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_app_config_delete_key(app, key, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param app: ID of the app (required)
        :type app: str
        :param key: Key to delete (required)
        :type key: str
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
            raise ValueError("Error! Please call the provisioning_api_app_config_delete_key_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.provisioning_api_app_config_delete_key_with_http_info(app, key, ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def provisioning_api_app_config_delete_key_with_http_info(self, app : Annotated[StrictStr, Field(..., description="ID of the app")], key : Annotated[StrictStr, Field(..., description="Key to delete")], ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """Delete a config key of an app  # noqa: E501

        This endpoint requires admin access  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_app_config_delete_key_with_http_info(app, key, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param app: ID of the app (required)
        :type app: str
        :param key: Key to delete (required)
        :type key: str
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
            'app',
            'key',
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
                    " to method provisioning_api_app_config_delete_key" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}
        if _params['app']:
            _path_params['app'] = _params['app']

        if _params['key']:
            _path_params['key'] = _params['key']


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
            '200': "CoreWhatsNewDismiss200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key}', 'DELETE',
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
    def provisioning_api_app_config_get_apps(self, ocs_api_request : StrictStr, **kwargs) -> ProvisioningApiAppConfigGetApps200Response:  # noqa: E501
        """Get a list of apps  # noqa: E501

        This endpoint requires admin access  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_app_config_get_apps(ocs_api_request, async_req=True)
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
        :rtype: ProvisioningApiAppConfigGetApps200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the provisioning_api_app_config_get_apps_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.provisioning_api_app_config_get_apps_with_http_info(ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def provisioning_api_app_config_get_apps_with_http_info(self, ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """Get a list of apps  # noqa: E501

        This endpoint requires admin access  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_app_config_get_apps_with_http_info(ocs_api_request, async_req=True)
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
        :rtype: tuple(ProvisioningApiAppConfigGetApps200Response, status_code(int), headers(HTTPHeaderDict))
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
                    " to method provisioning_api_app_config_get_apps" % _key
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
            '200': "ProvisioningApiAppConfigGetApps200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/apps/provisioning_api/api/v1/config/apps', 'GET',
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
    def provisioning_api_app_config_get_keys(self, app : Annotated[StrictStr, Field(..., description="ID of the app")], ocs_api_request : StrictStr, **kwargs) -> ProvisioningApiAppConfigGetApps200Response:  # noqa: E501
        """Get the config keys of an app  # noqa: E501

        This endpoint requires admin access  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_app_config_get_keys(app, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param app: ID of the app (required)
        :type app: str
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
        :rtype: ProvisioningApiAppConfigGetApps200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the provisioning_api_app_config_get_keys_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.provisioning_api_app_config_get_keys_with_http_info(app, ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def provisioning_api_app_config_get_keys_with_http_info(self, app : Annotated[StrictStr, Field(..., description="ID of the app")], ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """Get the config keys of an app  # noqa: E501

        This endpoint requires admin access  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_app_config_get_keys_with_http_info(app, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param app: ID of the app (required)
        :type app: str
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
        :rtype: tuple(ProvisioningApiAppConfigGetApps200Response, status_code(int), headers(HTTPHeaderDict))
        """

        _params = locals()

        _all_params = [
            'app',
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
                    " to method provisioning_api_app_config_get_keys" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}
        if _params['app']:
            _path_params['app'] = _params['app']


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
            '200': "ProvisioningApiAppConfigGetApps200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}', 'GET',
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
    def provisioning_api_app_config_get_value(self, app : Annotated[StrictStr, Field(..., description="ID if the app")], key : Annotated[StrictStr, Field(..., description="Key")], ocs_api_request : StrictStr, default_value : Annotated[Optional[StrictStr], Field(description="Default returned value if the value is empty")] = None, **kwargs) -> ProvisioningApiAppConfigGetValue200Response:  # noqa: E501
        """Get a the config value of an app  # noqa: E501

        This endpoint requires admin access  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_app_config_get_value(app, key, ocs_api_request, default_value, async_req=True)
        >>> result = thread.get()

        :param app: ID if the app (required)
        :type app: str
        :param key: Key (required)
        :type key: str
        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param default_value: Default returned value if the value is empty
        :type default_value: str
        :param async_req: Whether to execute the request asynchronously.
        :type async_req: bool, optional
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: Returns the result object.
                 If the method is called asynchronously,
                 returns the request thread.
        :rtype: ProvisioningApiAppConfigGetValue200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the provisioning_api_app_config_get_value_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.provisioning_api_app_config_get_value_with_http_info(app, key, ocs_api_request, default_value, **kwargs)  # noqa: E501

    @validate_arguments
    def provisioning_api_app_config_get_value_with_http_info(self, app : Annotated[StrictStr, Field(..., description="ID if the app")], key : Annotated[StrictStr, Field(..., description="Key")], ocs_api_request : StrictStr, default_value : Annotated[Optional[StrictStr], Field(description="Default returned value if the value is empty")] = None, **kwargs) -> ApiResponse:  # noqa: E501
        """Get a the config value of an app  # noqa: E501

        This endpoint requires admin access  # noqa: E501
        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_app_config_get_value_with_http_info(app, key, ocs_api_request, default_value, async_req=True)
        >>> result = thread.get()

        :param app: ID if the app (required)
        :type app: str
        :param key: Key (required)
        :type key: str
        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param default_value: Default returned value if the value is empty
        :type default_value: str
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
        :rtype: tuple(ProvisioningApiAppConfigGetValue200Response, status_code(int), headers(HTTPHeaderDict))
        """

        _params = locals()

        _all_params = [
            'app',
            'key',
            'ocs_api_request',
            'default_value'
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
                    " to method provisioning_api_app_config_get_value" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}
        if _params['app']:
            _path_params['app'] = _params['app']

        if _params['key']:
            _path_params['key'] = _params['key']


        # process the query parameters
        _query_params = []
        if _params.get('default_value') is not None:  # noqa: E501
            _query_params.append(('defaultValue', _params['default_value']))

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
            '200': "ProvisioningApiAppConfigGetValue200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key}', 'GET',
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
    def provisioning_api_app_config_set_value(self, value : Annotated[StrictStr, Field(..., description="New value for the key")], app : Annotated[StrictStr, Field(..., description="ID of the app")], key : Annotated[StrictStr, Field(..., description="Key to update")], ocs_api_request : StrictStr, **kwargs) -> CoreWhatsNewDismiss200Response:  # noqa: E501
        """Update the config value of an app  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_app_config_set_value(value, app, key, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param value: New value for the key (required)
        :type value: str
        :param app: ID of the app (required)
        :type app: str
        :param key: Key to update (required)
        :type key: str
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
            raise ValueError("Error! Please call the provisioning_api_app_config_set_value_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.provisioning_api_app_config_set_value_with_http_info(value, app, key, ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def provisioning_api_app_config_set_value_with_http_info(self, value : Annotated[StrictStr, Field(..., description="New value for the key")], app : Annotated[StrictStr, Field(..., description="ID of the app")], key : Annotated[StrictStr, Field(..., description="Key to update")], ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """Update the config value of an app  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_app_config_set_value_with_http_info(value, app, key, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param value: New value for the key (required)
        :type value: str
        :param app: ID of the app (required)
        :type app: str
        :param key: Key to update (required)
        :type key: str
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
            'value',
            'app',
            'key',
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
                    " to method provisioning_api_app_config_set_value" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}
        if _params['app']:
            _path_params['app'] = _params['app']

        if _params['key']:
            _path_params['key'] = _params['key']


        # process the query parameters
        _query_params = []
        if _params.get('value') is not None:  # noqa: E501
            _query_params.append(('value', _params['value']))

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
            '/ocs/v2.php/apps/provisioning_api/api/v1/config/apps/{app}/{key}', 'POST',
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
