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

from nextcloud_client.models.core_text_processing_api_list_tasks_by_app200_response import CoreTextProcessingApiListTasksByApp200Response
from nextcloud_client.models.core_text_processing_api_schedule200_response import CoreTextProcessingApiSchedule200Response
from nextcloud_client.models.core_text_processing_api_task_types200_response import CoreTextProcessingApiTaskTypes200Response

from nextcloud_client.api_client import ApiClient
from nextcloud_client.api_response import ApiResponse
from nextcloud_client.exceptions import (  # noqa: F401
    ApiTypeError,
    ApiValueError
)


class CoreTextProcessingApiApi(object):
    """NOTE: This class is auto generated by OpenAPI Generator
    Ref: https://openapi-generator.tech

    Do not edit the class manually.
    """

    def __init__(self, api_client=None):
        if api_client is None:
            api_client = ApiClient.get_default()
        self.api_client = api_client

    @validate_arguments
    def core_text_processing_api_delete_task(self, id : Annotated[StrictInt, Field(..., description="The id of the task")], ocs_api_request : StrictStr, **kwargs) -> CoreTextProcessingApiSchedule200Response:  # noqa: E501
        """This endpoint allows to delete a scheduled task for a user  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_text_processing_api_delete_task(id, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param id: The id of the task (required)
        :type id: int
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
        :rtype: CoreTextProcessingApiSchedule200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the core_text_processing_api_delete_task_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.core_text_processing_api_delete_task_with_http_info(id, ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def core_text_processing_api_delete_task_with_http_info(self, id : Annotated[StrictInt, Field(..., description="The id of the task")], ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """This endpoint allows to delete a scheduled task for a user  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_text_processing_api_delete_task_with_http_info(id, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param id: The id of the task (required)
        :type id: int
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
        :rtype: tuple(CoreTextProcessingApiSchedule200Response, status_code(int), headers(HTTPHeaderDict))
        """

        _params = locals()

        _all_params = [
            'id',
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
                    " to method core_text_processing_api_delete_task" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}
        if _params['id']:
            _path_params['id'] = _params['id']


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
            '200': "CoreTextProcessingApiSchedule200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/textprocessing/task/{id}', 'DELETE',
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
    def core_text_processing_api_get_task(self, id : Annotated[StrictInt, Field(..., description="The id of the task")], ocs_api_request : StrictStr, **kwargs) -> CoreTextProcessingApiSchedule200Response:  # noqa: E501
        """This endpoint allows checking the status and results of a task. Tasks are removed 1 week after receiving their last update.  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_text_processing_api_get_task(id, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param id: The id of the task (required)
        :type id: int
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
        :rtype: CoreTextProcessingApiSchedule200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the core_text_processing_api_get_task_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.core_text_processing_api_get_task_with_http_info(id, ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def core_text_processing_api_get_task_with_http_info(self, id : Annotated[StrictInt, Field(..., description="The id of the task")], ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """This endpoint allows checking the status and results of a task. Tasks are removed 1 week after receiving their last update.  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_text_processing_api_get_task_with_http_info(id, ocs_api_request, async_req=True)
        >>> result = thread.get()

        :param id: The id of the task (required)
        :type id: int
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
        :rtype: tuple(CoreTextProcessingApiSchedule200Response, status_code(int), headers(HTTPHeaderDict))
        """

        _params = locals()

        _all_params = [
            'id',
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
                    " to method core_text_processing_api_get_task" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}
        if _params['id']:
            _path_params['id'] = _params['id']


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
            '200': "CoreTextProcessingApiSchedule200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/textprocessing/task/{id}', 'GET',
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
    def core_text_processing_api_list_tasks_by_app(self, app_id : Annotated[StrictStr, Field(..., description="ID of the app")], ocs_api_request : StrictStr, identifier : Annotated[Optional[StrictStr], Field(description="An arbitrary identifier for the task")] = None, **kwargs) -> CoreTextProcessingApiListTasksByApp200Response:  # noqa: E501
        """This endpoint returns a list of tasks of a user that are related with a specific appId and optionally with an identifier  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_text_processing_api_list_tasks_by_app(app_id, ocs_api_request, identifier, async_req=True)
        >>> result = thread.get()

        :param app_id: ID of the app (required)
        :type app_id: str
        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param identifier: An arbitrary identifier for the task
        :type identifier: str
        :param async_req: Whether to execute the request asynchronously.
        :type async_req: bool, optional
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: Returns the result object.
                 If the method is called asynchronously,
                 returns the request thread.
        :rtype: CoreTextProcessingApiListTasksByApp200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the core_text_processing_api_list_tasks_by_app_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.core_text_processing_api_list_tasks_by_app_with_http_info(app_id, ocs_api_request, identifier, **kwargs)  # noqa: E501

    @validate_arguments
    def core_text_processing_api_list_tasks_by_app_with_http_info(self, app_id : Annotated[StrictStr, Field(..., description="ID of the app")], ocs_api_request : StrictStr, identifier : Annotated[Optional[StrictStr], Field(description="An arbitrary identifier for the task")] = None, **kwargs) -> ApiResponse:  # noqa: E501
        """This endpoint returns a list of tasks of a user that are related with a specific appId and optionally with an identifier  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_text_processing_api_list_tasks_by_app_with_http_info(app_id, ocs_api_request, identifier, async_req=True)
        >>> result = thread.get()

        :param app_id: ID of the app (required)
        :type app_id: str
        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param identifier: An arbitrary identifier for the task
        :type identifier: str
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
        :rtype: tuple(CoreTextProcessingApiListTasksByApp200Response, status_code(int), headers(HTTPHeaderDict))
        """

        _params = locals()

        _all_params = [
            'app_id',
            'ocs_api_request',
            'identifier'
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
                    " to method core_text_processing_api_list_tasks_by_app" % _key
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
        if _params.get('identifier') is not None:  # noqa: E501
            _query_params.append(('identifier', _params['identifier']))

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
            '200': "CoreTextProcessingApiListTasksByApp200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/textprocessing/tasks/app/{appId}', 'GET',
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
    def core_text_processing_api_schedule(self, input : Annotated[StrictStr, Field(..., description="Input text")], type : Annotated[StrictStr, Field(..., description="Type of the task")], app_id : Annotated[StrictStr, Field(..., description="ID of the app that will execute the task")], ocs_api_request : StrictStr, identifier : Annotated[Optional[StrictStr], Field(description="An arbitrary identifier for the task")] = None, **kwargs) -> CoreTextProcessingApiSchedule200Response:  # noqa: E501
        """This endpoint allows scheduling a language model task  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_text_processing_api_schedule(input, type, app_id, ocs_api_request, identifier, async_req=True)
        >>> result = thread.get()

        :param input: Input text (required)
        :type input: str
        :param type: Type of the task (required)
        :type type: str
        :param app_id: ID of the app that will execute the task (required)
        :type app_id: str
        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param identifier: An arbitrary identifier for the task
        :type identifier: str
        :param async_req: Whether to execute the request asynchronously.
        :type async_req: bool, optional
        :param _request_timeout: timeout setting for this request. If one
                                 number provided, it will be total request
                                 timeout. It can also be a pair (tuple) of
                                 (connection, read) timeouts.
        :return: Returns the result object.
                 If the method is called asynchronously,
                 returns the request thread.
        :rtype: CoreTextProcessingApiSchedule200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the core_text_processing_api_schedule_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.core_text_processing_api_schedule_with_http_info(input, type, app_id, ocs_api_request, identifier, **kwargs)  # noqa: E501

    @validate_arguments
    def core_text_processing_api_schedule_with_http_info(self, input : Annotated[StrictStr, Field(..., description="Input text")], type : Annotated[StrictStr, Field(..., description="Type of the task")], app_id : Annotated[StrictStr, Field(..., description="ID of the app that will execute the task")], ocs_api_request : StrictStr, identifier : Annotated[Optional[StrictStr], Field(description="An arbitrary identifier for the task")] = None, **kwargs) -> ApiResponse:  # noqa: E501
        """This endpoint allows scheduling a language model task  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_text_processing_api_schedule_with_http_info(input, type, app_id, ocs_api_request, identifier, async_req=True)
        >>> result = thread.get()

        :param input: Input text (required)
        :type input: str
        :param type: Type of the task (required)
        :type type: str
        :param app_id: ID of the app that will execute the task (required)
        :type app_id: str
        :param ocs_api_request: (required)
        :type ocs_api_request: str
        :param identifier: An arbitrary identifier for the task
        :type identifier: str
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
        :rtype: tuple(CoreTextProcessingApiSchedule200Response, status_code(int), headers(HTTPHeaderDict))
        """

        _params = locals()

        _all_params = [
            'input',
            'type',
            'app_id',
            'ocs_api_request',
            'identifier'
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
                    " to method core_text_processing_api_schedule" % _key
                )
            _params[_key] = _val
        del _params['kwargs']

        _collection_formats = {}

        # process the path parameters
        _path_params = {}

        # process the query parameters
        _query_params = []
        if _params.get('input') is not None:  # noqa: E501
            _query_params.append(('input', _params['input']))

        if _params.get('type') is not None:  # noqa: E501
            _query_params.append(('type', _params['type']))

        if _params.get('app_id') is not None:  # noqa: E501
            _query_params.append(('appId', _params['app_id']))

        if _params.get('identifier') is not None:  # noqa: E501
            _query_params.append(('identifier', _params['identifier']))

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
            '200': "CoreTextProcessingApiSchedule200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/textprocessing/schedule', 'POST',
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
    def core_text_processing_api_task_types(self, ocs_api_request : StrictStr, **kwargs) -> CoreTextProcessingApiTaskTypes200Response:  # noqa: E501
        """This endpoint returns all available LanguageModel task types  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_text_processing_api_task_types(ocs_api_request, async_req=True)
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
        :rtype: CoreTextProcessingApiTaskTypes200Response
        """
        kwargs['_return_http_data_only'] = True
        if '_preload_content' in kwargs:
            raise ValueError("Error! Please call the core_text_processing_api_task_types_with_http_info method with `_preload_content` instead and obtain raw data from ApiResponse.raw_data")
        return self.core_text_processing_api_task_types_with_http_info(ocs_api_request, **kwargs)  # noqa: E501

    @validate_arguments
    def core_text_processing_api_task_types_with_http_info(self, ocs_api_request : StrictStr, **kwargs) -> ApiResponse:  # noqa: E501
        """This endpoint returns all available LanguageModel task types  # noqa: E501

        This method makes a synchronous HTTP request by default. To make an
        asynchronous HTTP request, please pass async_req=True

        >>> thread = api.core_text_processing_api_task_types_with_http_info(ocs_api_request, async_req=True)
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
        :rtype: tuple(CoreTextProcessingApiTaskTypes200Response, status_code(int), headers(HTTPHeaderDict))
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
                    " to method core_text_processing_api_task_types" % _key
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
            '200': "CoreTextProcessingApiTaskTypes200Response",
        }

        return self.api_client.call_api(
            '/ocs/v2.php/textprocessing/tasktypes', 'GET',
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
