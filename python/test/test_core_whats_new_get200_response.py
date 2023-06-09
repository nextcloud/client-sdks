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
from openapi_client.models.core_whats_new_get200_response import CoreWhatsNewGet200Response  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreWhatsNewGet200Response(unittest.TestCase):
    """CoreWhatsNewGet200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreWhatsNewGet200Response
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreWhatsNewGet200Response`
        """
        model = openapi_client.models.core_whats_new_get200_response.CoreWhatsNewGet200Response()  # noqa: E501
        if include_optional :
            return CoreWhatsNewGet200Response(
                ocs = openapi_client.models.core_whats_new_get_200_response_ocs.core_whats_new_get_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = openapi_client.models.core_whats_new_get_200_response_ocs_data.core_whats_new_get_200_response_ocs_data(
                        changelog_url = '', 
                        product = '', 
                        version = '', 
                        whats_new = openapi_client.models.core_whats_new_get_200_response_ocs_data_whats_new.core_whats_new_get_200_response_ocs_data_whatsNew(
                            regular = [
                                ''
                                ], 
                            admin = [
                                ''
                                ], ), ), )
            )
        else :
            return CoreWhatsNewGet200Response(
                ocs = openapi_client.models.core_whats_new_get_200_response_ocs.core_whats_new_get_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = openapi_client.models.core_whats_new_get_200_response_ocs_data.core_whats_new_get_200_response_ocs_data(
                        changelog_url = '', 
                        product = '', 
                        version = '', 
                        whats_new = openapi_client.models.core_whats_new_get_200_response_ocs_data_whats_new.core_whats_new_get_200_response_ocs_data_whatsNew(
                            regular = [
                                ''
                                ], 
                            admin = [
                                ''
                                ], ), ), ),
        )
        """

    def testCoreWhatsNewGet200Response(self):
        """Test CoreWhatsNewGet200Response"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
