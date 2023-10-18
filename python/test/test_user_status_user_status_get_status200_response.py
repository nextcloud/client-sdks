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
from openapi_client.models.user_status_user_status_get_status200_response import UserStatusUserStatusGetStatus200Response  # noqa: E501
from openapi_client.rest import ApiException

class TestUserStatusUserStatusGetStatus200Response(unittest.TestCase):
    """UserStatusUserStatusGetStatus200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test UserStatusUserStatusGetStatus200Response
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `UserStatusUserStatusGetStatus200Response`
        """
        model = openapi_client.models.user_status_user_status_get_status200_response.UserStatusUserStatusGetStatus200Response()  # noqa: E501
        if include_optional :
            return UserStatusUserStatusGetStatus200Response(
                ocs = openapi_client.models.user_status_user_status_get_status_200_response_ocs.user_status_user_status_get_status_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = null, )
            )
        else :
            return UserStatusUserStatusGetStatus200Response(
                ocs = openapi_client.models.user_status_user_status_get_status_200_response_ocs.user_status_user_status_get_status_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = null, ),
        )
        """

    def testUserStatusUserStatusGetStatus200Response(self):
        """Test UserStatusUserStatusGetStatus200Response"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()