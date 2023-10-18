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
from openapi_client.models.provisioning_api_apps_get_apps200_response_ocs import ProvisioningApiAppsGetApps200ResponseOcs  # noqa: E501
from openapi_client.rest import ApiException

class TestProvisioningApiAppsGetApps200ResponseOcs(unittest.TestCase):
    """ProvisioningApiAppsGetApps200ResponseOcs unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test ProvisioningApiAppsGetApps200ResponseOcs
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ProvisioningApiAppsGetApps200ResponseOcs`
        """
        model = openapi_client.models.provisioning_api_apps_get_apps200_response_ocs.ProvisioningApiAppsGetApps200ResponseOcs()  # noqa: E501
        if include_optional :
            return ProvisioningApiAppsGetApps200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ), 
                data = openapi_client.models.provisioning_api_apps_get_apps_200_response_ocs_data.provisioning_api_apps_get_apps_200_response_ocs_data(
                    apps = [
                        ''
                        ], )
            )
        else :
            return ProvisioningApiAppsGetApps200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ),
                data = openapi_client.models.provisioning_api_apps_get_apps_200_response_ocs_data.provisioning_api_apps_get_apps_200_response_ocs_data(
                    apps = [
                        ''
                        ], ),
        )
        """

    def testProvisioningApiAppsGetApps200ResponseOcs(self):
        """Test ProvisioningApiAppsGetApps200ResponseOcs"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()