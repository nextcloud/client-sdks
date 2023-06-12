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
from openapi_client.api.provisioning_api_users_api import ProvisioningApiUsersApi  # noqa: E501
from openapi_client.rest import ApiException


class TestProvisioningApiUsersApi(unittest.TestCase):
    """ProvisioningApiUsersApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.provisioning_api_users_api.ProvisioningApiUsersApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_provisioning_api_users_add_sub_admin(self):
        """Test case for provisioning_api_users_add_sub_admin

        Make a user a subadmin of a group  # noqa: E501
        """
        pass

    def test_provisioning_api_users_add_to_group(self):
        """Test case for provisioning_api_users_add_to_group

        Add a user to a group  # noqa: E501
        """
        pass

    def test_provisioning_api_users_add_user(self):
        """Test case for provisioning_api_users_add_user

        Create a new user  # noqa: E501
        """
        pass

    def test_provisioning_api_users_delete_user(self):
        """Test case for provisioning_api_users_delete_user

        Delete a user  # noqa: E501
        """
        pass

    def test_provisioning_api_users_disable_user(self):
        """Test case for provisioning_api_users_disable_user

        Disable a user  # noqa: E501
        """
        pass

    def test_provisioning_api_users_edit_user(self):
        """Test case for provisioning_api_users_edit_user

        Update a value of the user's details  # noqa: E501
        """
        pass

    def test_provisioning_api_users_edit_user_multi_value(self):
        """Test case for provisioning_api_users_edit_user_multi_value

        Update multiple values of the user's details  # noqa: E501
        """
        pass

    def test_provisioning_api_users_enable_user(self):
        """Test case for provisioning_api_users_enable_user

        Enable a user  # noqa: E501
        """
        pass

    def test_provisioning_api_users_get_current_user(self):
        """Test case for provisioning_api_users_get_current_user

        Get the details of the current user  # noqa: E501
        """
        pass

    def test_provisioning_api_users_get_editable_fields(self):
        """Test case for provisioning_api_users_get_editable_fields

        Get a list of fields that are editable for the current user  # noqa: E501
        """
        pass

    def test_provisioning_api_users_get_editable_fields_for_user(self):
        """Test case for provisioning_api_users_get_editable_fields_for_user

        Get a list of fields that are editable for a user  # noqa: E501
        """
        pass

    def test_provisioning_api_users_get_user(self):
        """Test case for provisioning_api_users_get_user

        Get the details of a user  # noqa: E501
        """
        pass

    def test_provisioning_api_users_get_user_sub_admin_groups(self):
        """Test case for provisioning_api_users_get_user_sub_admin_groups

        Get the groups a user is a subadmin of  # noqa: E501
        """
        pass

    def test_provisioning_api_users_get_users(self):
        """Test case for provisioning_api_users_get_users

        Get a list of users  # noqa: E501
        """
        pass

    def test_provisioning_api_users_get_users_details(self):
        """Test case for provisioning_api_users_get_users_details

        Get a list of users and their details  # noqa: E501
        """
        pass

    def test_provisioning_api_users_get_users_groups(self):
        """Test case for provisioning_api_users_get_users_groups

        Get a list of groups the user belongs to  # noqa: E501
        """
        pass

    def test_provisioning_api_users_remove_from_group(self):
        """Test case for provisioning_api_users_remove_from_group

        Remove a user from a group  # noqa: E501
        """
        pass

    def test_provisioning_api_users_remove_sub_admin(self):
        """Test case for provisioning_api_users_remove_sub_admin

        Remove a user from the subadmins of a group  # noqa: E501
        """
        pass

    def test_provisioning_api_users_resend_welcome_message(self):
        """Test case for provisioning_api_users_resend_welcome_message

        Resend the welcome message  # noqa: E501
        """
        pass

    def test_provisioning_api_users_search_by_phone_numbers(self):
        """Test case for provisioning_api_users_search_by_phone_numbers

        Search users by their phone numbers  # noqa: E501
        """
        pass

    def test_provisioning_api_users_wipe_user_devices(self):
        """Test case for provisioning_api_users_wipe_user_devices

        Wipe all devices of a user  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
