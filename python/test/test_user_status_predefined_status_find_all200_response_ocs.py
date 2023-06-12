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
from openapi_client.models.user_status_predefined_status_find_all200_response_ocs import UserStatusPredefinedStatusFindAll200ResponseOcs  # noqa: E501
from openapi_client.rest import ApiException

class TestUserStatusPredefinedStatusFindAll200ResponseOcs(unittest.TestCase):
    """UserStatusPredefinedStatusFindAll200ResponseOcs unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test UserStatusPredefinedStatusFindAll200ResponseOcs
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `UserStatusPredefinedStatusFindAll200ResponseOcs`
        """
        model = openapi_client.models.user_status_predefined_status_find_all200_response_ocs.UserStatusPredefinedStatusFindAll200ResponseOcs()  # noqa: E501
        if include_optional :
            return UserStatusPredefinedStatusFindAll200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ), 
                data = [
                    openapi_client.models.user_status_predefined.UserStatusPredefined(
                        id = '', 
                        icon = '', 
                        message = '', 
                        clear_at = openapi_client.models.user_status_clear_at.UserStatusClearAt(
                            type = 'period', 
                            time = null, ), 
                        visible = True, )
                    ]
            )
        else :
            return UserStatusPredefinedStatusFindAll200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ),
                data = [
                    openapi_client.models.user_status_predefined.UserStatusPredefined(
                        id = '', 
                        icon = '', 
                        message = '', 
                        clear_at = openapi_client.models.user_status_clear_at.UserStatusClearAt(
                            type = 'period', 
                            time = null, ), 
                        visible = True, )
                    ],
        )
        """

    def testUserStatusPredefinedStatusFindAll200ResponseOcs(self):
        """Test UserStatusPredefinedStatusFindAll200ResponseOcs"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
