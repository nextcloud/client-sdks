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
from openapi_client.models.core_reference_provider import CoreReferenceProvider  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreReferenceProvider(unittest.TestCase):
    """CoreReferenceProvider unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreReferenceProvider
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreReferenceProvider`
        """
        model = openapi_client.models.core_reference_provider.CoreReferenceProvider()  # noqa: E501
        if include_optional :
            return CoreReferenceProvider(
                id = '', 
                title = '', 
                icon_url = '', 
                order = 56, 
                search_providers_ids = [
                    ''
                    ]
            )
        else :
            return CoreReferenceProvider(
                id = '',
                title = '',
                icon_url = '',
                order = 56,
                search_providers_ids = [
                    ''
                    ],
        )
        """

    def testCoreReferenceProvider(self):
        """Test CoreReferenceProvider"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
