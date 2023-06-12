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
from openapi_client.models.files_sharing_shareapi_create_share200_response import FilesSharingShareapiCreateShare200Response  # noqa: E501
from openapi_client.rest import ApiException

class TestFilesSharingShareapiCreateShare200Response(unittest.TestCase):
    """FilesSharingShareapiCreateShare200Response unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test FilesSharingShareapiCreateShare200Response
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `FilesSharingShareapiCreateShare200Response`
        """
        model = openapi_client.models.files_sharing_shareapi_create_share200_response.FilesSharingShareapiCreateShare200Response()  # noqa: E501
        if include_optional :
            return FilesSharingShareapiCreateShare200Response(
                ocs = openapi_client.models.files_sharing_shareapi_create_share_200_response_ocs.files_sharing_shareapi_create_share_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = openapi_client.models.files_sharing_share.FilesSharingShare(
                        attributes = '', 
                        can_delete = True, 
                        can_edit = True, 
                        displayname_file_owner = '', 
                        displayname_owner = '', 
                        expiration = '', 
                        file_parent = 56, 
                        file_source = 56, 
                        file_target = '', 
                        has_preview = True, 
                        id = '', 
                        item_source = 56, 
                        item_type = '', 
                        label = '', 
                        mail_send = 56, 
                        mimetype = '', 
                        note = '', 
                        password = '', 
                        password_expiration_time = '', 
                        path = '', 
                        permissions = 56, 
                        send_password_by_talk = True, 
                        share_type = 56, 
                        share_with = '', 
                        share_with_avatar = '', 
                        share_with_displayname = '', 
                        share_with_link = '', 
                        status = null, 
                        stime = 56, 
                        storage = 56, 
                        storage_id = '', 
                        token = '', 
                        uid_file_owner = '', 
                        uid_owner = '', 
                        url = '', ), )
            )
        else :
            return FilesSharingShareapiCreateShare200Response(
                ocs = openapi_client.models.files_sharing_shareapi_create_share_200_response_ocs.files_sharing_shareapi_create_share_200_response_ocs(
                    meta = openapi_client.models.ocs_meta.OCSMeta(
                        status = '', 
                        statuscode = 56, 
                        message = '', 
                        totalitems = '', 
                        itemsperpage = '', ), 
                    data = openapi_client.models.files_sharing_share.FilesSharingShare(
                        attributes = '', 
                        can_delete = True, 
                        can_edit = True, 
                        displayname_file_owner = '', 
                        displayname_owner = '', 
                        expiration = '', 
                        file_parent = 56, 
                        file_source = 56, 
                        file_target = '', 
                        has_preview = True, 
                        id = '', 
                        item_source = 56, 
                        item_type = '', 
                        label = '', 
                        mail_send = 56, 
                        mimetype = '', 
                        note = '', 
                        password = '', 
                        password_expiration_time = '', 
                        path = '', 
                        permissions = 56, 
                        send_password_by_talk = True, 
                        share_type = 56, 
                        share_with = '', 
                        share_with_avatar = '', 
                        share_with_displayname = '', 
                        share_with_link = '', 
                        status = null, 
                        stime = 56, 
                        storage = 56, 
                        storage_id = '', 
                        token = '', 
                        uid_file_owner = '', 
                        uid_owner = '', 
                        url = '', ), ),
        )
        """

    def testFilesSharingShareapiCreateShare200Response(self):
        """Test FilesSharingShareapiCreateShare200Response"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
