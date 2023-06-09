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
from openapi_client.models.core_hover_card_get_user200_response import CoreHoverCardGetUser200Response  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreHoverCardGetUser200Response(unittest.TestCase):
    """CoreHoverCardGetUser200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreHoverCardGetUser200Response
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreHoverCardGetUser200Response`
        """
        model = openapi_client.models.core_hover_card_get_user200_response.CoreHoverCardGetUser200Response()  # noqa: E501
        if include_optional :
            return CoreHoverCardGetUser200Response(
                ocs = openapi_client.models.core_hover_card_get_user_200_response_ocs.core_hover_card_get_user_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = openapi_client.models.core_hover_card_get_user_200_response_ocs_data.core_hover_card_get_user_200_response_ocs_data(
                        user_id = '', 
                        display_name = '', 
                        actions = [
                            openapi_client.models.core_contacts_action.CoreContactsAction(
                                title = '', 
                                icon = '', 
                                hyperlink = '', 
                                app_id = '', )
                            ], ), )
            )
        else :
            return CoreHoverCardGetUser200Response(
                ocs = openapi_client.models.core_hover_card_get_user_200_response_ocs.core_hover_card_get_user_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = openapi_client.models.core_hover_card_get_user_200_response_ocs_data.core_hover_card_get_user_200_response_ocs_data(
                        user_id = '', 
                        display_name = '', 
                        actions = [
                            openapi_client.models.core_contacts_action.CoreContactsAction(
                                title = '', 
                                icon = '', 
                                hyperlink = '', 
                                app_id = '', )
                            ], ), ),
        )
        """

    def testCoreHoverCardGetUser200Response(self):
        """Test CoreHoverCardGetUser200Response"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
