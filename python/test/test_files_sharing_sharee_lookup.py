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
from openapi_client.models.files_sharing_sharee_lookup import FilesSharingShareeLookup  # noqa: E501
from openapi_client.rest import ApiException

class TestFilesSharingShareeLookup(unittest.TestCase):
    """FilesSharingShareeLookup unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test FilesSharingShareeLookup
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `FilesSharingShareeLookup`
        """
        model = openapi_client.models.files_sharing_sharee_lookup.FilesSharingShareeLookup()  # noqa: E501
        if include_optional :
            return FilesSharingShareeLookup(
                count = 56, 
                label = '', 
                extra = openapi_client.models.files_sharing_sharee_lookup_all_of_extra.FilesSharingShareeLookup_allOf_extra(
                    federation_id = '', 
                    name = openapi_client.models.files_sharing_lookup.FilesSharingLookup(
                        value = '', 
                        verified = 56, ), 
                    email = openapi_client.models.files_sharing_lookup.FilesSharingLookup(
                        value = '', 
                        verified = 56, ), 
                    address = , 
                    website = , 
                    twitter = , 
                    phone = , 
                    twitter_signature = , 
                    website_signature = , 
                    userid = , ), 
                value = None
            )
        else :
            return FilesSharingShareeLookup(
                count = 56,
                label = '',
                extra = openapi_client.models.files_sharing_sharee_lookup_all_of_extra.FilesSharingShareeLookup_allOf_extra(
                    federation_id = '', 
                    name = openapi_client.models.files_sharing_lookup.FilesSharingLookup(
                        value = '', 
                        verified = 56, ), 
                    email = openapi_client.models.files_sharing_lookup.FilesSharingLookup(
                        value = '', 
                        verified = 56, ), 
                    address = , 
                    website = , 
                    twitter = , 
                    phone = , 
                    twitter_signature = , 
                    website_signature = , 
                    userid = , ),
                value = None,
        )
        """

    def testFilesSharingShareeLookup(self):
        """Test FilesSharingShareeLookup"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
