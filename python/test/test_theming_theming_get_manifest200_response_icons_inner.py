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
from openapi_client.models.theming_theming_get_manifest200_response_icons_inner import ThemingThemingGetManifest200ResponseIconsInner  # noqa: E501
from openapi_client.rest import ApiException

class TestThemingThemingGetManifest200ResponseIconsInner(unittest.TestCase):
    """ThemingThemingGetManifest200ResponseIconsInner unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test ThemingThemingGetManifest200ResponseIconsInner
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ThemingThemingGetManifest200ResponseIconsInner`
        """
        model = openapi_client.models.theming_theming_get_manifest200_response_icons_inner.ThemingThemingGetManifest200ResponseIconsInner()  # noqa: E501
        if include_optional :
            return ThemingThemingGetManifest200ResponseIconsInner(
                src = '', 
                type = '', 
                sizes = ''
            )
        else :
            return ThemingThemingGetManifest200ResponseIconsInner(
                src = '',
                type = '',
                sizes = '',
        )
        """

    def testThemingThemingGetManifest200ResponseIconsInner(self):
        """Test ThemingThemingGetManifest200ResponseIconsInner"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
