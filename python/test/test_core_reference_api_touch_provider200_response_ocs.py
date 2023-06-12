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
from openapi_client.models.core_reference_api_touch_provider200_response_ocs import CoreReferenceApiTouchProvider200ResponseOcs  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreReferenceApiTouchProvider200ResponseOcs(unittest.TestCase):
    """CoreReferenceApiTouchProvider200ResponseOcs unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreReferenceApiTouchProvider200ResponseOcs
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreReferenceApiTouchProvider200ResponseOcs`
        """
        model = openapi_client.models.core_reference_api_touch_provider200_response_ocs.CoreReferenceApiTouchProvider200ResponseOcs()  # noqa: E501
        if include_optional :
            return CoreReferenceApiTouchProvider200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ), 
                data = openapi_client.models.core_reference_api_touch_provider_200_response_ocs_data.core_reference_api_touch_provider_200_response_ocs_data(
                    success = True, )
            )
        else :
            return CoreReferenceApiTouchProvider200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ),
                data = openapi_client.models.core_reference_api_touch_provider_200_response_ocs_data.core_reference_api_touch_provider_200_response_ocs_data(
                    success = True, ),
        )
        """

    def testCoreReferenceApiTouchProvider200ResponseOcs(self):
        """Test CoreReferenceApiTouchProvider200ResponseOcs"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
