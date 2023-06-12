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
from openapi_client.models.files_sharing_sharee_circle import FilesSharingShareeCircle  # noqa: E501
from openapi_client.rest import ApiException

class TestFilesSharingShareeCircle(unittest.TestCase):
    """FilesSharingShareeCircle unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test FilesSharingShareeCircle
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `FilesSharingShareeCircle`
        """
        model = openapi_client.models.files_sharing_sharee_circle.FilesSharingShareeCircle()  # noqa: E501
        if include_optional :
            return FilesSharingShareeCircle(
                count = 56, 
                label = '', 
                share_with_description = '', 
                value = None
            )
        else :
            return FilesSharingShareeCircle(
                count = 56,
                label = '',
                share_with_description = '',
                value = None,
        )
        """

    def testFilesSharingShareeCircle(self):
        """Test FilesSharingShareeCircle"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
