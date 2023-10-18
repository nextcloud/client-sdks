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
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_public_expire_date import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate(unittest.TestCase):
    """CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate`
        """
        model = openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_public_expire_date.CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate()  # noqa: E501
        if include_optional :
            return CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate(
                enabled = True, 
                days = 56, 
                enforced = True
            )
        else :
            return CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate(
                enabled = True,
        )
        """

    def testCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate(self):
        """Test CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()