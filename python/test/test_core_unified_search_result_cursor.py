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
from openapi_client.models.core_unified_search_result_cursor import CoreUnifiedSearchResultCursor  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreUnifiedSearchResultCursor(unittest.TestCase):
    """CoreUnifiedSearchResultCursor unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreUnifiedSearchResultCursor
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreUnifiedSearchResultCursor`
        """
        model = openapi_client.models.core_unified_search_result_cursor.CoreUnifiedSearchResultCursor()  # noqa: E501
        if include_optional :
            return CoreUnifiedSearchResultCursor(
            )
        else :
            return CoreUnifiedSearchResultCursor(
        )
        """

    def testCoreUnifiedSearchResultCursor(self):
        """Test CoreUnifiedSearchResultCursor"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
