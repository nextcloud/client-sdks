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
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing(unittest.TestCase):
    """CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing`
        """
        model = openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing.CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing()  # noqa: E501
        if include_optional :
            return CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing(
                api_enabled = True, 
                public = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public(
                    enabled = True, 
                    password = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_password.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_password(
                        enforced = True, 
                        ask_for_optional_password = True, ), 
                    multiple_links = True, 
                    expire_date = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_expire_date(
                        enabled = True, 
                        days = 56, 
                        enforced = True, ), 
                    expire_date_internal = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_expire_date(
                        enabled = True, 
                        days = 56, 
                        enforced = True, ), 
                    expire_date_remote = , 
                    send_mail = True, 
                    upload = True, 
                    upload_files_drop = True, ), 
                user = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user(
                    send_mail = True, 
                    expire_date = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date(
                        enabled = True, ), ), 
                resharing = True, 
                group_sharing = True, 
                group = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_group.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_group(
                    enabled = True, 
                    expire_date = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date(
                        enabled = True, ), ), 
                default_permissions = 56, 
                federation = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_federation.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_federation(
                    outgoing = True, 
                    incoming = True, 
                    expire_date = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date(
                        enabled = True, ), 
                    expire_date_supported = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date(
                        enabled = True, ), ), 
                sharee = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_sharee.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_sharee(
                    query_lookup_default = True, 
                    always_show_unique = True, )
            )
        else :
            return CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing(
                api_enabled = True,
                public = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public(
                    enabled = True, 
                    password = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_password.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_password(
                        enforced = True, 
                        ask_for_optional_password = True, ), 
                    multiple_links = True, 
                    expire_date = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_expire_date(
                        enabled = True, 
                        days = 56, 
                        enforced = True, ), 
                    expire_date_internal = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_public_expire_date(
                        enabled = True, 
                        days = 56, 
                        enforced = True, ), 
                    expire_date_remote = , 
                    send_mail = True, 
                    upload = True, 
                    upload_files_drop = True, ),
                user = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user(
                    send_mail = True, 
                    expire_date = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date(
                        enabled = True, ), ),
                resharing = True,
                federation = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_federation.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_federation(
                    outgoing = True, 
                    incoming = True, 
                    expire_date = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date(
                        enabled = True, ), 
                    expire_date_supported = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_user_expire_date(
                        enabled = True, ), ),
                sharee = openapi_client.models.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_sharee.core_ocs_get_capabilities_200_response_ocs_data_capabilities_files_sharing_sharee(
                    query_lookup_default = True, 
                    always_show_unique = True, ),
        )
        """

    def testCoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing(self):
        """Test CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()