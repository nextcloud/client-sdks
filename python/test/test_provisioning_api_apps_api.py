# coding: utf-8

"""
    nextcloud

    Nextcloud APIs  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


import unittest

import openapi_client
from openapi_client.api.provisioning_api_apps_api import ProvisioningApiAppsApi  # noqa: E501
from openapi_client.rest import ApiException


class TestProvisioningApiAppsApi(unittest.TestCase):
    """ProvisioningApiAppsApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.provisioning_api_apps_api.ProvisioningApiAppsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_provisioning_api_apps_disable(self):
        """Test case for provisioning_api_apps_disable

        Disable an app  # noqa: E501
        """
        pass

    def test_provisioning_api_apps_enable(self):
        """Test case for provisioning_api_apps_enable

        Enable an app  # noqa: E501
        """
        pass

    def test_provisioning_api_apps_get_app_info(self):
        """Test case for provisioning_api_apps_get_app_info

        Get the app info for an app  # noqa: E501
        """
        pass

    def test_provisioning_api_apps_get_apps(self):
        """Test case for provisioning_api_apps_get_apps

        Get a list of installed apps  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
