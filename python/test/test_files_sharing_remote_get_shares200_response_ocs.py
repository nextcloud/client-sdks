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
from openapi_client.models.files_sharing_remote_get_shares200_response_ocs import FilesSharingRemoteGetShares200ResponseOcs  # noqa: E501
from openapi_client.rest import ApiException

class TestFilesSharingRemoteGetShares200ResponseOcs(unittest.TestCase):
    """FilesSharingRemoteGetShares200ResponseOcs unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test FilesSharingRemoteGetShares200ResponseOcs
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `FilesSharingRemoteGetShares200ResponseOcs`
        """
        model = openapi_client.models.files_sharing_remote_get_shares200_response_ocs.FilesSharingRemoteGetShares200ResponseOcs()  # noqa: E501
        if include_optional :
            return FilesSharingRemoteGetShares200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ), 
                data = [
                    openapi_client.models.files_sharing_remote_share.FilesSharingRemoteShare(
                        accepted = True, 
                        file_id = 56, 
                        id = 56, 
                        mimetype = '', 
                        mountpoint = '', 
                        mtime = 56, 
                        name = '', 
                        owner = '', 
                        parent = 56, 
                        permissions = 56, 
                        remote = '', 
                        remote_id = '', 
                        share_token = '', 
                        share_type = 56, 
                        type = '', 
                        user = '', )
                    ]
            )
        else :
            return FilesSharingRemoteGetShares200ResponseOcs(
                meta = openapi_client.models.ocs_meta.OCSMeta(
                    status = '', 
                    statuscode = 56, 
                    message = '', 
                    totalitems = '', 
                    itemsperpage = '', ),
                data = [
                    openapi_client.models.files_sharing_remote_share.FilesSharingRemoteShare(
                        accepted = True, 
                        file_id = 56, 
                        id = 56, 
                        mimetype = '', 
                        mountpoint = '', 
                        mtime = 56, 
                        name = '', 
                        owner = '', 
                        parent = 56, 
                        permissions = 56, 
                        remote = '', 
                        remote_id = '', 
                        share_token = '', 
                        share_type = 56, 
                        type = '', 
                        user = '', )
                    ],
        )
        """

    def testFilesSharingRemoteGetShares200ResponseOcs(self):
        """Test FilesSharingRemoteGetShares200ResponseOcs"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
