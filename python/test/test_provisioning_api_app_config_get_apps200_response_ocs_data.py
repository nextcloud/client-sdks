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
from openapi_client.models.provisioning_api_app_config_get_apps200_response_ocs_data import ProvisioningApiAppConfigGetApps200ResponseOcsData  # noqa: E501
from openapi_client.rest import ApiException

class TestProvisioningApiAppConfigGetApps200ResponseOcsData(unittest.TestCase):
    """ProvisioningApiAppConfigGetApps200ResponseOcsData unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test ProvisioningApiAppConfigGetApps200ResponseOcsData
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ProvisioningApiAppConfigGetApps200ResponseOcsData`
        """
        model = openapi_client.models.provisioning_api_app_config_get_apps200_response_ocs_data.ProvisioningApiAppConfigGetApps200ResponseOcsData()  # noqa: E501
        if include_optional :
            return ProvisioningApiAppConfigGetApps200ResponseOcsData(
                data = [
                    ''
                    ]
            )
        else :
            return ProvisioningApiAppConfigGetApps200ResponseOcsData(
                data = [
                    ''
                    ],
        )
        """

    def testProvisioningApiAppConfigGetApps200ResponseOcsData(self):
        """Test ProvisioningApiAppConfigGetApps200ResponseOcsData"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
