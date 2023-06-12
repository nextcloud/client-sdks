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
from openapi_client.models.core_status import CoreStatus  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreStatus(unittest.TestCase):
    """CoreStatus unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreStatus
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreStatus`
        """
        model = openapi_client.models.core_status.CoreStatus()  # noqa: E501
        if include_optional :
            return CoreStatus(
                installed = True, 
                maintenance = True, 
                needs_db_upgrade = True, 
                version = '', 
                versionstring = '', 
                edition = '', 
                productname = '', 
                extended_support = True
            )
        else :
            return CoreStatus(
                installed = True,
                maintenance = True,
                needs_db_upgrade = True,
                version = '',
                versionstring = '',
                edition = '',
                productname = '',
                extended_support = True,
        )
        """

    def testCoreStatus(self):
        """Test CoreStatus"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
