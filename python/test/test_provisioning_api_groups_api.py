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
from openapi_client.api.provisioning_api_groups_api import ProvisioningApiGroupsApi  # noqa: E501
from openapi_client.rest import ApiException


class TestProvisioningApiGroupsApi(unittest.TestCase):
    """ProvisioningApiGroupsApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.provisioning_api_groups_api.ProvisioningApiGroupsApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_provisioning_api_groups_add_group(self):
        """Test case for provisioning_api_groups_add_group

        Create a new group  # noqa: E501
        """
        pass

    def test_provisioning_api_groups_delete_group(self):
        """Test case for provisioning_api_groups_delete_group

        Delete a group  # noqa: E501
        """
        pass

    def test_provisioning_api_groups_get_group(self):
        """Test case for provisioning_api_groups_get_group

        Get a list of users in the specified group  # noqa: E501
        """
        pass

    def test_provisioning_api_groups_get_group_users(self):
        """Test case for provisioning_api_groups_get_group_users

        Get a list of users in the specified group  # noqa: E501
        """
        pass

    def test_provisioning_api_groups_get_group_users_details(self):
        """Test case for provisioning_api_groups_get_group_users_details

        Get a list of users details in the specified group  # noqa: E501
        """
        pass

    def test_provisioning_api_groups_get_groups(self):
        """Test case for provisioning_api_groups_get_groups

        Get a list of groups  # noqa: E501
        """
        pass

    def test_provisioning_api_groups_get_groups_details(self):
        """Test case for provisioning_api_groups_get_groups_details

        Get a list of groups details  # noqa: E501
        """
        pass

    def test_provisioning_api_groups_get_sub_admins_of_group(self):
        """Test case for provisioning_api_groups_get_sub_admins_of_group

        Get the list of user IDs that are a subadmin of the group  # noqa: E501
        """
        pass

    def test_provisioning_api_groups_update_group(self):
        """Test case for provisioning_api_groups_update_group

        Update a group  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
