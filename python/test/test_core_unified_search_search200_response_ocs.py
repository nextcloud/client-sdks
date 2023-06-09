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
from openapi_client.models.core_unified_search_search200_response_ocs import CoreUnifiedSearchSearch200ResponseOcs  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreUnifiedSearchSearch200ResponseOcs(unittest.TestCase):
    """CoreUnifiedSearchSearch200ResponseOcs unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreUnifiedSearchSearch200ResponseOcs
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreUnifiedSearchSearch200ResponseOcs`
        """
        model = openapi_client.models.core_unified_search_search200_response_ocs.CoreUnifiedSearchSearch200ResponseOcs()  # noqa: E501
        if include_optional :
            return CoreUnifiedSearchSearch200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ), 
                data = openapi_client.models.core_unified_search_result.CoreUnifiedSearchResult(
                    name = '', 
                    is_paginated = True, 
                    entries = [
                        openapi_client.models.core_unified_search_result_entry.CoreUnifiedSearchResultEntry(
                            thumbnail_url = '', 
                            title = '', 
                            subline = '', 
                            resource_url = '', 
                            icon = '', 
                            rounded = True, 
                            attributes = [
                                ''
                                ], )
                        ], 
                    cursor = null, )
            )
        else :
            return CoreUnifiedSearchSearch200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ),
                data = openapi_client.models.core_unified_search_result.CoreUnifiedSearchResult(
                    name = '', 
                    is_paginated = True, 
                    entries = [
                        openapi_client.models.core_unified_search_result_entry.CoreUnifiedSearchResultEntry(
                            thumbnail_url = '', 
                            title = '', 
                            subline = '', 
                            resource_url = '', 
                            icon = '', 
                            rounded = True, 
                            attributes = [
                                ''
                                ], )
                        ], 
                    cursor = null, ),
        )
        """

    def testCoreUnifiedSearchSearch200ResponseOcs(self):
        """Test CoreUnifiedSearchSearch200ResponseOcs"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
