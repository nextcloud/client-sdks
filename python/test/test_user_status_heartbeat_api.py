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
from openapi_client.api.user_status_heartbeat_api import UserStatusHeartbeatApi  # noqa: E501
from openapi_client.rest import ApiException


class TestUserStatusHeartbeatApi(unittest.TestCase):
    """UserStatusHeartbeatApi unit test stubs"""

    def setUp(self):
        self.api = openapi_client.api.user_status_heartbeat_api.UserStatusHeartbeatApi()  # noqa: E501

    def tearDown(self):
        pass

    def test_user_status_heartbeat_heartbeat(self):
        """Test case for user_status_heartbeat_heartbeat

        Keep the current status alive  # noqa: E501
        """
        pass


if __name__ == '__main__':
    unittest.main()
