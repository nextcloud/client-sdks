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
from openapi_client.models.core_auto_complete_get200_response import CoreAutoCompleteGet200Response  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreAutoCompleteGet200Response(unittest.TestCase):
    """CoreAutoCompleteGet200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreAutoCompleteGet200Response
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreAutoCompleteGet200Response`
        """
        model = openapi_client.models.core_auto_complete_get200_response.CoreAutoCompleteGet200Response()  # noqa: E501
        if include_optional :
            return CoreAutoCompleteGet200Response(
                ocs = openapi_client.models.core_auto_complete_get_200_response_ocs.core_auto_complete_get_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = [
                        openapi_client.models.core_autocomplete_result.CoreAutocompleteResult(
                            id = '', 
                            label = '', 
                            icon = '', 
                            source = '', 
                            status = '', 
                            subline = '', 
                            share_with_display_name_unique = '', )
                        ], )
            )
        else :
            return CoreAutoCompleteGet200Response(
                ocs = openapi_client.models.core_auto_complete_get_200_response_ocs.core_auto_complete_get_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = [
                        openapi_client.models.core_autocomplete_result.CoreAutocompleteResult(
                            id = '', 
                            label = '', 
                            icon = '', 
                            source = '', 
                            status = '', 
                            subline = '', 
                            share_with_display_name_unique = '', )
                        ], ),
        )
        """

    def testCoreAutoCompleteGet200Response(self):
        """Test CoreAutoCompleteGet200Response"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()