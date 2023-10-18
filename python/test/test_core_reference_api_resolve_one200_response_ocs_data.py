# coding: utf-8

"""
    nextcloud

    Nextcloud APIs  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


import unittest
import datetime

import openapi_client
from openapi_client.models.core_reference_api_resolve_one200_response_ocs_data import CoreReferenceApiResolveOne200ResponseOcsData  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreReferenceApiResolveOne200ResponseOcsData(unittest.TestCase):
    """CoreReferenceApiResolveOne200ResponseOcsData unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreReferenceApiResolveOne200ResponseOcsData
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreReferenceApiResolveOne200ResponseOcsData`
        """
        model = openapi_client.models.core_reference_api_resolve_one200_response_ocs_data.CoreReferenceApiResolveOne200ResponseOcsData()  # noqa: E501
        if include_optional :
            return CoreReferenceApiResolveOne200ResponseOcsData(
                references = {
                    'key' : openapi_client.models.core_reference.CoreReference(
                        rich_object_type = '', 
                        rich_object = {
                            'key' : None
                            }, 
                        open_graph_object = openapi_client.models.core_reference_open_graph_object.CoreReference_openGraphObject(
                            id = '', 
                            name = '', 
                            description = '', 
                            thumb = '', 
                            link = '', ), 
                        accessible = True, )
                    }
            )
        else :
            return CoreReferenceApiResolveOne200ResponseOcsData(
                references = {
                    'key' : openapi_client.models.core_reference.CoreReference(
                        rich_object_type = '', 
                        rich_object = {
                            'key' : None
                            }, 
                        open_graph_object = openapi_client.models.core_reference_open_graph_object.CoreReference_openGraphObject(
                            id = '', 
                            name = '', 
                            description = '', 
                            thumb = '', 
                            link = '', ), 
                        accessible = True, )
                    },
        )
        """

    def testCoreReferenceApiResolveOne200ResponseOcsData(self):
        """Test CoreReferenceApiResolveOne200ResponseOcsData"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()