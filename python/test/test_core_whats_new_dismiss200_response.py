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
from openapi_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreWhatsNewDismiss200Response(unittest.TestCase):
    """CoreWhatsNewDismiss200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreWhatsNewDismiss200Response
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreWhatsNewDismiss200Response`
        """
        model = openapi_client.models.core_whats_new_dismiss200_response.CoreWhatsNewDismiss200Response()  # noqa: E501
        if include_optional :
            return CoreWhatsNewDismiss200Response(
                ocs = openapi_client.models.core_whats_new_dismiss_200_response_ocs.core_whats_new_dismiss_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = { }, )
            )
        else :
            return CoreWhatsNewDismiss200Response(
                ocs = openapi_client.models.core_whats_new_dismiss_200_response_ocs.core_whats_new_dismiss_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = { }, ),
        )
        """

    def testCoreWhatsNewDismiss200Response(self):
        """Test CoreWhatsNewDismiss200Response"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
