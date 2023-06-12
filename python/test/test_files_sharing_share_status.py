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
from openapi_client.models.files_sharing_share_status import FilesSharingShareStatus  # noqa: E501
from openapi_client.rest import ApiException

class TestFilesSharingShareStatus(unittest.TestCase):
    """FilesSharingShareStatus unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test FilesSharingShareStatus
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `FilesSharingShareStatus`
        """
        model = openapi_client.models.files_sharing_share_status.FilesSharingShareStatus()  # noqa: E501
        if include_optional :
            return FilesSharingShareStatus(
                status = '', 
                message = '', 
                icon = '', 
                clear_at = 56
            )
        else :
            return FilesSharingShareStatus(
                status = '',
                message = '',
                icon = '',
                clear_at = 56,
        )
        """

    def testFilesSharingShareStatus(self):
        """Test FilesSharingShareStatus"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
