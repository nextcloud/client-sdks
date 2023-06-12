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
from openapi_client.api.files_sharing_shareapi_api import FilesSharingShareapiApi  # noqa: E501
from openapi_client.rest import ApiException


class TestFilesSharingShareapiApi(unittest.TestCase):
    """FilesSharingShareapiApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.files_sharing_shareapi_api.FilesSharingShareapiApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_files_sharing_shareapi_accept_share(self):
        """Test case for files_sharing_shareapi_accept_share

        Accept a share  # noqa: E501
        """
        pass

    def test_files_sharing_shareapi_create_share(self):
        """Test case for files_sharing_shareapi_create_share

        Create a share  # noqa: E501
        """
        pass

    def test_files_sharing_shareapi_delete_share(self):
        """Test case for files_sharing_shareapi_delete_share

        Delete a share  # noqa: E501
        """
        pass

    def test_files_sharing_shareapi_get_inherited_shares(self):
        """Test case for files_sharing_shareapi_get_inherited_shares

        Get all shares relative to a file, including parent folders shares rights  # noqa: E501
        """
        pass

    def test_files_sharing_shareapi_get_share(self):
        """Test case for files_sharing_shareapi_get_share

        Get a specific share by id  # noqa: E501
        """
        pass

    def test_files_sharing_shareapi_get_shares(self):
        """Test case for files_sharing_shareapi_get_shares

        Get shares of the current user  # noqa: E501
        """
        pass

    def test_files_sharing_shareapi_pending_shares(self):
        """Test case for files_sharing_shareapi_pending_shares

        Get all shares that are still pending  # noqa: E501
        """
        pass

    def test_files_sharing_shareapi_update_share(self):
        """Test case for files_sharing_shareapi_update_share

        Update a share  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
