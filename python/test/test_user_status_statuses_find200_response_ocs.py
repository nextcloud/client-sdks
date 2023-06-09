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
from openapi_client.models.user_status_statuses_find200_response_ocs import UserStatusStatusesFind200ResponseOcs  # noqa: E501
from openapi_client.rest import ApiException

class TestUserStatusStatusesFind200ResponseOcs(unittest.TestCase):
    """UserStatusStatusesFind200ResponseOcs unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test UserStatusStatusesFind200ResponseOcs
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `UserStatusStatusesFind200ResponseOcs`
        """
        model = openapi_client.models.user_status_statuses_find200_response_ocs.UserStatusStatusesFind200ResponseOcs()  # noqa: E501
        if include_optional :
            return UserStatusStatusesFind200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ), 
                data = openapi_client.models.user_status_public.UserStatusPublic(
                    user_id = '', 
                    message = '', 
                    icon = '', 
                    clear_at = 56, 
                    status = '', )
            )
        else :
            return UserStatusStatusesFind200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ),
                data = openapi_client.models.user_status_public.UserStatusPublic(
                    user_id = '', 
                    message = '', 
                    icon = '', 
                    clear_at = 56, 
                    status = '', ),
        )
        """

    def testUserStatusStatusesFind200ResponseOcs(self):
        """Test UserStatusStatusesFind200ResponseOcs"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
