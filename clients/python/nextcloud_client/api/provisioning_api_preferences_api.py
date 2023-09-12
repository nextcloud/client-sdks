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

from pydantic import Field, StrictStr, conlist

from nextcloud_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response

from nextcloud_client.api_client import ApiClient
from nextcloud_client.api_response import ApiResponse
from nextcloud_client.exceptions import (  # noqa: F401
    ApiTypeError,
    ApiValueError
)


class ProvisioningApiPreferencesApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient.get_default()
        self.api_client = api_client

    @validate_arguments
    def provisioning_api_preferences_delete_multiple_preference(self, config_keys : Annotated[conlist(StrictStr), Field(..., description="Keys to delete")], app_id : Annotated[StrictStr, Field(..., description="ID of the app")], ocs_api_request : StrictStr, **kwargs) -> CoreWhatsNewDismiss200Response:  # noqa: E501
        """Delete multiple preferences for an app  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_preferences_delete_multiple_preference(config_keys, app_id, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param config_keys: Keys to delete (required)
        :type config_keys: List[str]
        :param app_id: ID of the app (required)
        :type app_id: str
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
            raise ValueError("Error! Please call the provisioning_api_preferences_delete_multiple_preference_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.provisioning_api_preferences_delete_multiple_preference_with_http_info(config_keys, app_id, ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def provisioning_api_preferences_delete_multiple_preference_with_http_info(self, config_keys : Annotated[conlist(StrictStr), Field(..., description="Keys to delete")], app_id : Annotated[StrictStr, Field(..., description="ID of the app")], ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """Delete multiple preferences for an app  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_preferences_delete_multiple_preference_with_http_info(config_keys, app_id, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param config_keys: Keys to delete (required)
        :type config_keys: List[str]
        :param app_id: ID of the app (required)
        :type app_id: str
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
            'config_keys',
            'app_id',
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
                    " to method provisioning_api_preferences_delete_multiple_preference" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}
        if _params['app_id']:
            _path_params['appId'] = _params['app_id']


        # process the query parameters
        _query_params = []
        if _params.get('config_keys') is not None:  # noqa: E501
            _query_params.append(('configKeys[]', _params['config_keys']))
            _collection_formats['configKeys[]'] = 'multi'

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
            '/ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}', 'DELETE',
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
    def provisioning_api_preferences_delete_preference(self, app_id : Annotated[StrictStr, Field(..., description="ID of the app")], config_key : Annotated[StrictStr, Field(..., description="Key to delete")], ocs_api_request : StrictStr, **kwargs) -> CoreWhatsNewDismiss200Response:  # noqa: E501
        """Delete a preference for an app  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_preferences_delete_preference(app_id, config_key, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param app_id: ID of the app (required)
        :type app_id: str
        :param config_key: Key to delete (required)
        :type config_key: str
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
            raise ValueError("Error! Please call the provisioning_api_preferences_delete_preference_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.provisioning_api_preferences_delete_preference_with_http_info(app_id, config_key, ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def provisioning_api_preferences_delete_preference_with_http_info(self, app_id : Annotated[StrictStr, Field(..., description="ID of the app")], config_key : Annotated[StrictStr, Field(..., description="Key to delete")], ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """Delete a preference for an app  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_preferences_delete_preference_with_http_info(app_id, config_key, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param app_id: ID of the app (required)
        :type app_id: str
        :param config_key: Key to delete (required)
        :type config_key: str
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
            'app_id',
            'config_key',
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
                    " to method provisioning_api_preferences_delete_preference" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}
        if _params['app_id']:
            _path_params['appId'] = _params['app_id']

        if _params['config_key']:
            _path_params['configKey'] = _params['config_key']


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
            '/ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}/{configKey}', 'DELETE',
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
    def provisioning_api_preferences_set_multiple_preferences(self, configs : Annotated[StrictStr, Field(..., description="Key-value pairs of the preferences")], app_id : Annotated[StrictStr, Field(..., description="ID of the app")], ocs_api_request : StrictStr, **kwargs) -> CoreWhatsNewDismiss200Response:  # noqa: E501
        """Update multiple preference values of an app  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_preferences_set_multiple_preferences(configs, app_id, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param configs: Key-value pairs of the preferences (required)
        :type configs: str
        :param app_id: ID of the app (required)
        :type app_id: str
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
            raise ValueError("Error! Please call the provisioning_api_preferences_set_multiple_preferences_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.provisioning_api_preferences_set_multiple_preferences_with_http_info(configs, app_id, ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def provisioning_api_preferences_set_multiple_preferences_with_http_info(self, configs : Annotated[StrictStr, Field(..., description="Key-value pairs of the preferences")], app_id : Annotated[StrictStr, Field(..., description="ID of the app")], ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """Update multiple preference values of an app  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_preferences_set_multiple_preferences_with_http_info(configs, app_id, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param configs: Key-value pairs of the preferences (required)
        :type configs: str
        :param app_id: ID of the app (required)
        :type app_id: str
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
            'configs',
            'app_id',
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
                    " to method provisioning_api_preferences_set_multiple_preferences" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}
        if _params['app_id']:
            _path_params['appId'] = _params['app_id']


        # process the query parameters
        _query_params = []
        if _params.get('configs') is not None:  # noqa: E501
            _query_params.append(('configs', _params['configs']))

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
            '/ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}', 'POST',
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
    def provisioning_api_preferences_set_preference(self, config_value : Annotated[StrictStr, Field(..., description="New value")], app_id : Annotated[StrictStr, Field(..., description="ID of the app")], config_key : Annotated[StrictStr, Field(..., description="Key of the preference")], ocs_api_request : StrictStr, **kwargs) -> CoreWhatsNewDismiss200Response:  # noqa: E501
        """Update a preference value of an app  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_preferences_set_preference(config_value, app_id, config_key, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param config_value: New value (required)
        :type config_value: str
        :param app_id: ID of the app (required)
        :type app_id: str
        :param config_key: Key of the preference (required)
        :type config_key: str
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
            raise ValueError("Error! Please call the provisioning_api_preferences_set_preference_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.provisioning_api_preferences_set_preference_with_http_info(config_value, app_id, config_key, ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def provisioning_api_preferences_set_preference_with_http_info(self, config_value : Annotated[StrictStr, Field(..., description="New value")], app_id : Annotated[StrictStr, Field(..., description="ID of the app")], config_key : Annotated[StrictStr, Field(..., description="Key of the preference")], ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """Update a preference value of an app  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.provisioning_api_preferences_set_preference_with_http_info(config_value, app_id, config_key, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param config_value: New value (required)
        :type config_value: str
        :param app_id: ID of the app (required)
        :type app_id: str
        :param config_key: Key of the preference (required)
        :type config_key: str
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
            'config_value',
            'app_id',
            'config_key',
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
                    " to method provisioning_api_preferences_set_preference" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}
        if _params['app_id']:
            _path_params['appId'] = _params['app_id']

        if _params['config_key']:
            _path_params['configKey'] = _params['config_key']


        # process the query parameters
        _query_params = []
        if _params.get('config_value') is not None:  # noqa: E501
            _query_params.append(('configValue', _params['config_value']))

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
            '/ocs/v2.php/apps/provisioning_api/api/v1/config/users/{appId}/{configKey}', 'POST',
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