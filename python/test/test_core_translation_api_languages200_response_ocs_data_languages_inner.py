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
from openapi_client.models.core_translation_api_languages200_response_ocs_data_languages_inner import CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner  # noqa: E501
from openapi_client.rest import ApiException

class TestCoreTranslationApiLanguages200ResponseOcsDataLanguagesInner(unittest.TestCase):
    """CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional):
        """Test CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner`
        """
        model = openapi_client.models.core_translation_api_languages200_response_ocs_data_languages_inner.CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner()  # noqa: E501
        if include_optional :
            return CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner(
                var_from = '', 
                from_label = '', 
                to = '', 
                to_label = ''
            )
        else :
            return CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner(
                var_from = '',
                from_label = '',
                to = '',
                to_label = '',
        )
        """

    def testCoreTranslationApiLanguages200ResponseOcsDataLanguagesInner(self):
        """Test CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()