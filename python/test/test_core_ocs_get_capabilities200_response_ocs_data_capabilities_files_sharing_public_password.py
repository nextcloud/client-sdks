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
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_public_password import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword(unittest.TestCase):
    """CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword`
        """
        model = openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_public_password.CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword()  # noqa: E501
        if include_optional :
            return CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword(
                enforced = True, 
                ask_for_optional_password = True
            )
        else :
            return CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword(
                enforced = True,
                ask_for_optional_password = True,
        )
        """

    def testCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword(self):
        """Test CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
