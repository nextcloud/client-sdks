# coding: utf-8

# flake8: noqa
"""
    nextcloud

    Nextcloud APIs  # noqa: E501

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""


# import models into model package
from openapi_client.models.cloud_federation_api_add_share import CloudFederationApiAddShare
from openapi_client.models.cloud_federation_api_capabilities import CloudFederationApiCapabilities
from openapi_client.models.cloud_federation_api_error import CloudFederationApiError
from openapi_client.models.cloud_federation_api_validation_error import CloudFederationApiValidationError
from openapi_client.models.cloud_federation_api_validation_error_all_of import CloudFederationApiValidationErrorAllOf
from openapi_client.models.cloud_federation_api_validation_error_all_of_validation_errors import CloudFederationApiValidationErrorAllOfValidationErrors
from openapi_client.models.core_app_password_get_app_password200_response import CoreAppPasswordGetAppPassword200Response
from openapi_client.models.core_app_password_get_app_password200_response_ocs import CoreAppPasswordGetAppPassword200ResponseOcs
from openapi_client.models.core_app_password_get_app_password200_response_ocs_data import CoreAppPasswordGetAppPassword200ResponseOcsData
from openapi_client.models.core_auto_complete_get200_response import CoreAutoCompleteGet200Response
from openapi_client.models.core_auto_complete_get200_response_ocs import CoreAutoCompleteGet200ResponseOcs
from openapi_client.models.core_autocomplete_result import CoreAutocompleteResult
from openapi_client.models.core_collaboration_resources_list_collection200_response import CoreCollaborationResourcesListCollection200Response
from openapi_client.models.core_collaboration_resources_list_collection200_response_ocs import CoreCollaborationResourcesListCollection200ResponseOcs
from openapi_client.models.core_collaboration_resources_search_collections200_response import CoreCollaborationResourcesSearchCollections200Response
from openapi_client.models.core_collaboration_resources_search_collections200_response_ocs import CoreCollaborationResourcesSearchCollections200ResponseOcs
from openapi_client.models.core_collection import CoreCollection
from openapi_client.models.core_contacts_action import CoreContactsAction
from openapi_client.models.core_hover_card_get_user200_response import CoreHoverCardGetUser200Response
from openapi_client.models.core_hover_card_get_user200_response_ocs import CoreHoverCardGetUser200ResponseOcs
from openapi_client.models.core_hover_card_get_user200_response_ocs_data import CoreHoverCardGetUser200ResponseOcsData
from openapi_client.models.core_login_flow_v2 import CoreLoginFlowV2
from openapi_client.models.core_login_flow_v2_credentials import CoreLoginFlowV2Credentials
from openapi_client.models.core_login_flow_v2_poll import CoreLoginFlowV2Poll
from openapi_client.models.core_navigation_entry import CoreNavigationEntry
from openapi_client.models.core_navigation_entry_order import CoreNavigationEntryOrder
from openapi_client.models.core_navigation_get_apps_navigation200_response import CoreNavigationGetAppsNavigation200Response
from openapi_client.models.core_navigation_get_apps_navigation200_response_ocs import CoreNavigationGetAppsNavigation200ResponseOcs
from openapi_client.models.core_ocs_get_capabilities200_response import CoreOcsGetCapabilities200Response
from openapi_client.models.core_ocs_get_capabilities200_response_ocs import CoreOcsGetCapabilities200ResponseOcs
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data import CoreOcsGetCapabilities200ResponseOcsData
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities import CoreOcsGetCapabilities200ResponseOcsDataCapabilities
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_dav import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesDav
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharing
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_federation import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingFederation
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_group import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingGroup
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_public import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublic
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_public_expire_date import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicExpireDate
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_public_password import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingPublicPassword
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_sharee import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingSharee
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_user import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUser
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_files_sharing_user_expire_date import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesFilesSharingUserExpireDate
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_ocm import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcm
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_ocm_resource_types_inner import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInner
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_ocm_resource_types_inner_protocols import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesOcmResourceTypesInnerProtocols
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_provisioning_api import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesProvisioningApi
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_theming import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesTheming
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_capabilities_user_status import CoreOcsGetCapabilities200ResponseOcsDataCapabilitiesUserStatus
from openapi_client.models.core_ocs_get_capabilities200_response_ocs_data_version import CoreOcsGetCapabilities200ResponseOcsDataVersion
from openapi_client.models.core_reference import CoreReference
from openapi_client.models.core_reference_api_get_providers_info200_response import CoreReferenceApiGetProvidersInfo200Response
from openapi_client.models.core_reference_api_get_providers_info200_response_ocs import CoreReferenceApiGetProvidersInfo200ResponseOcs
from openapi_client.models.core_reference_api_resolve_one200_response import CoreReferenceApiResolveOne200Response
from openapi_client.models.core_reference_api_resolve_one200_response_ocs import CoreReferenceApiResolveOne200ResponseOcs
from openapi_client.models.core_reference_api_resolve_one200_response_ocs_data import CoreReferenceApiResolveOne200ResponseOcsData
from openapi_client.models.core_reference_api_touch_provider200_response import CoreReferenceApiTouchProvider200Response
from openapi_client.models.core_reference_api_touch_provider200_response_ocs import CoreReferenceApiTouchProvider200ResponseOcs
from openapi_client.models.core_reference_api_touch_provider200_response_ocs_data import CoreReferenceApiTouchProvider200ResponseOcsData
from openapi_client.models.core_reference_open_graph_object import CoreReferenceOpenGraphObject
from openapi_client.models.core_reference_provider import CoreReferenceProvider
from openapi_client.models.core_status import CoreStatus
from openapi_client.models.core_translation_api_languages200_response import CoreTranslationApiLanguages200Response
from openapi_client.models.core_translation_api_languages200_response_ocs import CoreTranslationApiLanguages200ResponseOcs
from openapi_client.models.core_translation_api_languages200_response_ocs_data import CoreTranslationApiLanguages200ResponseOcsData
from openapi_client.models.core_translation_api_languages200_response_ocs_data_languages_inner import CoreTranslationApiLanguages200ResponseOcsDataLanguagesInner
from openapi_client.models.core_translation_api_translate200_response import CoreTranslationApiTranslate200Response
from openapi_client.models.core_translation_api_translate200_response_ocs import CoreTranslationApiTranslate200ResponseOcs
from openapi_client.models.core_translation_api_translate200_response_ocs_data import CoreTranslationApiTranslate200ResponseOcsData
from openapi_client.models.core_unified_search_get_providers200_response import CoreUnifiedSearchGetProviders200Response
from openapi_client.models.core_unified_search_get_providers200_response_ocs import CoreUnifiedSearchGetProviders200ResponseOcs
from openapi_client.models.core_unified_search_provider import CoreUnifiedSearchProvider
from openapi_client.models.core_unified_search_result import CoreUnifiedSearchResult
from openapi_client.models.core_unified_search_result_cursor import CoreUnifiedSearchResultCursor
from openapi_client.models.core_unified_search_result_entry import CoreUnifiedSearchResultEntry
from openapi_client.models.core_unified_search_search200_response import CoreUnifiedSearchSearch200Response
from openapi_client.models.core_unified_search_search200_response_ocs import CoreUnifiedSearchSearch200ResponseOcs
from openapi_client.models.core_whats_new_dismiss200_response import CoreWhatsNewDismiss200Response
from openapi_client.models.core_whats_new_dismiss200_response_ocs import CoreWhatsNewDismiss200ResponseOcs
from openapi_client.models.core_whats_new_get200_response import CoreWhatsNewGet200Response
from openapi_client.models.core_whats_new_get200_response_ocs import CoreWhatsNewGet200ResponseOcs
from openapi_client.models.core_whats_new_get200_response_ocs_data import CoreWhatsNewGet200ResponseOcsData
from openapi_client.models.core_whats_new_get200_response_ocs_data_whats_new import CoreWhatsNewGet200ResponseOcsDataWhatsNew
from openapi_client.models.core_wipe_check_wipe200_response import CoreWipeCheckWipe200Response
from openapi_client.models.dashboard_dashboard_api_get_widget_items200_response import DashboardDashboardApiGetWidgetItems200Response
from openapi_client.models.dashboard_dashboard_api_get_widget_items200_response_ocs import DashboardDashboardApiGetWidgetItems200ResponseOcs
from openapi_client.models.dashboard_dashboard_api_get_widgets200_response import DashboardDashboardApiGetWidgets200Response
from openapi_client.models.dashboard_dashboard_api_get_widgets200_response_ocs import DashboardDashboardApiGetWidgets200ResponseOcs
from openapi_client.models.dashboard_widget import DashboardWidget
from openapi_client.models.dashboard_widget_buttons_inner import DashboardWidgetButtonsInner
from openapi_client.models.dashboard_widget_item import DashboardWidgetItem
from openapi_client.models.dav_capabilities import DavCapabilities
from openapi_client.models.dav_direct_get_url200_response import DavDirectGetUrl200Response
from openapi_client.models.dav_direct_get_url200_response_ocs import DavDirectGetUrl200ResponseOcs
from openapi_client.models.dav_direct_get_url200_response_ocs_data import DavDirectGetUrl200ResponseOcsData
from openapi_client.models.files_sharing_capabilities import FilesSharingCapabilities
from openapi_client.models.files_sharing_deleted_share import FilesSharingDeletedShare
from openapi_client.models.files_sharing_deleted_shareapi_list200_response import FilesSharingDeletedShareapiList200Response
from openapi_client.models.files_sharing_deleted_shareapi_list200_response_ocs import FilesSharingDeletedShareapiList200ResponseOcs
from openapi_client.models.files_sharing_lookup import FilesSharingLookup
from openapi_client.models.files_sharing_remote_get_share200_response import FilesSharingRemoteGetShare200Response
from openapi_client.models.files_sharing_remote_get_share200_response_ocs import FilesSharingRemoteGetShare200ResponseOcs
from openapi_client.models.files_sharing_remote_get_shares200_response import FilesSharingRemoteGetShares200Response
from openapi_client.models.files_sharing_remote_get_shares200_response_ocs import FilesSharingRemoteGetShares200ResponseOcs
from openapi_client.models.files_sharing_remote_share import FilesSharingRemoteShare
from openapi_client.models.files_sharing_share import FilesSharingShare
from openapi_client.models.files_sharing_share_info import FilesSharingShareInfo
from openapi_client.models.files_sharing_share_status import FilesSharingShareStatus
from openapi_client.models.files_sharing_share_status_one_of import FilesSharingShareStatusOneOf
from openapi_client.models.files_sharing_shareapi_create_share200_response import FilesSharingShareapiCreateShare200Response
from openapi_client.models.files_sharing_shareapi_create_share200_response_ocs import FilesSharingShareapiCreateShare200ResponseOcs
from openapi_client.models.files_sharing_shareapi_get_shares200_response import FilesSharingShareapiGetShares200Response
from openapi_client.models.files_sharing_shareapi_get_shares200_response_ocs import FilesSharingShareapiGetShares200ResponseOcs
from openapi_client.models.files_sharing_sharee import FilesSharingSharee
from openapi_client.models.files_sharing_sharee_circle import FilesSharingShareeCircle
from openapi_client.models.files_sharing_sharee_circle_all_of import FilesSharingShareeCircleAllOf
from openapi_client.models.files_sharing_sharee_circle_all_of_value import FilesSharingShareeCircleAllOfValue
from openapi_client.models.files_sharing_sharee_circle_all_of_value_all_of import FilesSharingShareeCircleAllOfValueAllOf
from openapi_client.models.files_sharing_sharee_email import FilesSharingShareeEmail
from openapi_client.models.files_sharing_sharee_email_all_of import FilesSharingShareeEmailAllOf
from openapi_client.models.files_sharing_sharee_lookup import FilesSharingShareeLookup
from openapi_client.models.files_sharing_sharee_lookup_all_of import FilesSharingShareeLookupAllOf
from openapi_client.models.files_sharing_sharee_lookup_all_of_extra import FilesSharingShareeLookupAllOfExtra
from openapi_client.models.files_sharing_sharee_lookup_all_of_value import FilesSharingShareeLookupAllOfValue
from openapi_client.models.files_sharing_sharee_lookup_all_of_value_all_of import FilesSharingShareeLookupAllOfValueAllOf
from openapi_client.models.files_sharing_sharee_remote import FilesSharingShareeRemote
from openapi_client.models.files_sharing_sharee_remote_all_of import FilesSharingShareeRemoteAllOf
from openapi_client.models.files_sharing_sharee_remote_all_of_value import FilesSharingShareeRemoteAllOfValue
from openapi_client.models.files_sharing_sharee_remote_all_of_value_all_of import FilesSharingShareeRemoteAllOfValueAllOf
from openapi_client.models.files_sharing_sharee_remote_group import FilesSharingShareeRemoteGroup
from openapi_client.models.files_sharing_sharee_remote_group_all_of import FilesSharingShareeRemoteGroupAllOf
from openapi_client.models.files_sharing_sharee_user import FilesSharingShareeUser
from openapi_client.models.files_sharing_sharee_user_all_of import FilesSharingShareeUserAllOf
from openapi_client.models.files_sharing_sharee_user_all_of_status import FilesSharingShareeUserAllOfStatus
from openapi_client.models.files_sharing_sharee_value import FilesSharingShareeValue
from openapi_client.models.files_sharing_sharees_recommended_result import FilesSharingShareesRecommendedResult
from openapi_client.models.files_sharing_sharees_recommended_result_exact import FilesSharingShareesRecommendedResultExact
from openapi_client.models.files_sharing_sharees_search_result import FilesSharingShareesSearchResult
from openapi_client.models.files_sharing_sharees_search_result_exact import FilesSharingShareesSearchResultExact
from openapi_client.models.files_sharing_shareesapi_find_recommended200_response import FilesSharingShareesapiFindRecommended200Response
from openapi_client.models.files_sharing_shareesapi_find_recommended200_response_ocs import FilesSharingShareesapiFindRecommended200ResponseOcs
from openapi_client.models.files_sharing_shareesapi_search200_response import FilesSharingShareesapiSearch200Response
from openapi_client.models.files_sharing_shareesapi_search200_response_ocs import FilesSharingShareesapiSearch200ResponseOcs
from openapi_client.models.ocs_meta import OCSMeta
from openapi_client.models.oauth2_oauth_api_get_token200_response import Oauth2OauthApiGetToken200Response
from openapi_client.models.provisioning_api_app_config_get_apps200_response import ProvisioningApiAppConfigGetApps200Response
from openapi_client.models.provisioning_api_app_config_get_apps200_response_ocs import ProvisioningApiAppConfigGetApps200ResponseOcs
from openapi_client.models.provisioning_api_app_config_get_apps200_response_ocs_data import ProvisioningApiAppConfigGetApps200ResponseOcsData
from openapi_client.models.provisioning_api_app_config_get_value200_response import ProvisioningApiAppConfigGetValue200Response
from openapi_client.models.provisioning_api_app_config_get_value200_response_ocs import ProvisioningApiAppConfigGetValue200ResponseOcs
from openapi_client.models.provisioning_api_app_config_get_value200_response_ocs_data import ProvisioningApiAppConfigGetValue200ResponseOcsData
from openapi_client.models.provisioning_api_app_info import ProvisioningApiAppInfo
from openapi_client.models.provisioning_api_apps_get_app_info200_response import ProvisioningApiAppsGetAppInfo200Response
from openapi_client.models.provisioning_api_apps_get_app_info200_response_ocs import ProvisioningApiAppsGetAppInfo200ResponseOcs
from openapi_client.models.provisioning_api_apps_get_apps200_response import ProvisioningApiAppsGetApps200Response
from openapi_client.models.provisioning_api_apps_get_apps200_response_ocs import ProvisioningApiAppsGetApps200ResponseOcs
from openapi_client.models.provisioning_api_apps_get_apps200_response_ocs_data import ProvisioningApiAppsGetApps200ResponseOcsData
from openapi_client.models.provisioning_api_capabilities import ProvisioningApiCapabilities
from openapi_client.models.provisioning_api_group_details import ProvisioningApiGroupDetails
from openapi_client.models.provisioning_api_group_details_usercount import ProvisioningApiGroupDetailsUsercount
from openapi_client.models.provisioning_api_groups_get_group_users200_response import ProvisioningApiGroupsGetGroupUsers200Response
from openapi_client.models.provisioning_api_groups_get_group_users200_response_ocs import ProvisioningApiGroupsGetGroupUsers200ResponseOcs
from openapi_client.models.provisioning_api_groups_get_group_users200_response_ocs_data import ProvisioningApiGroupsGetGroupUsers200ResponseOcsData
from openapi_client.models.provisioning_api_groups_get_group_users_details200_response import ProvisioningApiGroupsGetGroupUsersDetails200Response
from openapi_client.models.provisioning_api_groups_get_group_users_details200_response_ocs import ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcs
from openapi_client.models.provisioning_api_groups_get_group_users_details200_response_ocs_data import ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsData
from openapi_client.models.provisioning_api_groups_get_group_users_details200_response_ocs_data_users_value import ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValue
from openapi_client.models.provisioning_api_groups_get_group_users_details200_response_ocs_data_users_value_one_of import ProvisioningApiGroupsGetGroupUsersDetails200ResponseOcsDataUsersValueOneOf
from openapi_client.models.provisioning_api_groups_get_groups200_response import ProvisioningApiGroupsGetGroups200Response
from openapi_client.models.provisioning_api_groups_get_groups200_response_ocs import ProvisioningApiGroupsGetGroups200ResponseOcs
from openapi_client.models.provisioning_api_groups_get_groups200_response_ocs_data import ProvisioningApiGroupsGetGroups200ResponseOcsData
from openapi_client.models.provisioning_api_groups_get_groups_details200_response import ProvisioningApiGroupsGetGroupsDetails200Response
from openapi_client.models.provisioning_api_groups_get_groups_details200_response_ocs import ProvisioningApiGroupsGetGroupsDetails200ResponseOcs
from openapi_client.models.provisioning_api_groups_get_groups_details200_response_ocs_data import ProvisioningApiGroupsGetGroupsDetails200ResponseOcsData
from openapi_client.models.provisioning_api_groups_get_sub_admins_of_group200_response import ProvisioningApiGroupsGetSubAdminsOfGroup200Response
from openapi_client.models.provisioning_api_groups_get_sub_admins_of_group200_response_ocs import ProvisioningApiGroupsGetSubAdminsOfGroup200ResponseOcs
from openapi_client.models.provisioning_api_user_details import ProvisioningApiUserDetails
from openapi_client.models.provisioning_api_user_details_backend_capabilities import ProvisioningApiUserDetailsBackendCapabilities
from openapi_client.models.provisioning_api_user_details_quota import ProvisioningApiUserDetailsQuota
from openapi_client.models.provisioning_api_user_details_quota_quota import ProvisioningApiUserDetailsQuotaQuota
from openapi_client.models.provisioning_api_users_add_user200_response import ProvisioningApiUsersAddUser200Response
from openapi_client.models.provisioning_api_users_add_user200_response_ocs import ProvisioningApiUsersAddUser200ResponseOcs
from openapi_client.models.provisioning_api_users_get_user200_response import ProvisioningApiUsersGetUser200Response
from openapi_client.models.provisioning_api_users_get_user200_response_ocs import ProvisioningApiUsersGetUser200ResponseOcs
from openapi_client.models.provisioning_api_users_search_by_phone_numbers200_response import ProvisioningApiUsersSearchByPhoneNumbers200Response
from openapi_client.models.provisioning_api_users_search_by_phone_numbers200_response_ocs import ProvisioningApiUsersSearchByPhoneNumbers200ResponseOcs
from openapi_client.models.theming_background import ThemingBackground
from openapi_client.models.theming_public_capabilities import ThemingPublicCapabilities
from openapi_client.models.theming_theming_get_manifest200_response import ThemingThemingGetManifest200Response
from openapi_client.models.theming_theming_get_manifest200_response_icons_inner import ThemingThemingGetManifest200ResponseIconsInner
from openapi_client.models.user_status_capabilities import UserStatusCapabilities
from openapi_client.models.user_status_clear_at import UserStatusClearAt
from openapi_client.models.user_status_clear_at_time import UserStatusClearAtTime
from openapi_client.models.user_status_clear_at_time_type import UserStatusClearAtTimeType
from openapi_client.models.user_status_predefined import UserStatusPredefined
from openapi_client.models.user_status_predefined_status_find_all200_response import UserStatusPredefinedStatusFindAll200Response
from openapi_client.models.user_status_predefined_status_find_all200_response_ocs import UserStatusPredefinedStatusFindAll200ResponseOcs
from openapi_client.models.user_status_private import UserStatusPrivate
from openapi_client.models.user_status_private_all_of import UserStatusPrivateAllOf
from openapi_client.models.user_status_public import UserStatusPublic
from openapi_client.models.user_status_statuses_find200_response import UserStatusStatusesFind200Response
from openapi_client.models.user_status_statuses_find200_response_ocs import UserStatusStatusesFind200ResponseOcs
from openapi_client.models.user_status_statuses_find_all200_response import UserStatusStatusesFindAll200Response
from openapi_client.models.user_status_statuses_find_all200_response_ocs import UserStatusStatusesFindAll200ResponseOcs
from openapi_client.models.user_status_user_status_get_status200_response import UserStatusUserStatusGetStatus200Response
from openapi_client.models.user_status_user_status_get_status200_response_ocs import UserStatusUserStatusGetStatus200ResponseOcs
from openapi_client.models.weather_status_capabilities import WeatherStatusCapabilities
from openapi_client.models.weather_status_forecast import WeatherStatusForecast
from openapi_client.models.weather_status_forecast_data import WeatherStatusForecastData
from openapi_client.models.weather_status_forecast_data_instant import WeatherStatusForecastDataInstant
from openapi_client.models.weather_status_forecast_data_instant_details import WeatherStatusForecastDataInstantDetails
from openapi_client.models.weather_status_forecast_data_next12_hours import WeatherStatusForecastDataNext12Hours
from openapi_client.models.weather_status_forecast_data_next12_hours_details import WeatherStatusForecastDataNext12HoursDetails
from openapi_client.models.weather_status_forecast_data_next12_hours_summary import WeatherStatusForecastDataNext12HoursSummary
from openapi_client.models.weather_status_forecast_data_next1_hours import WeatherStatusForecastDataNext1Hours
from openapi_client.models.weather_status_forecast_data_next1_hours_details import WeatherStatusForecastDataNext1HoursDetails
from openapi_client.models.weather_status_forecast_data_next6_hours import WeatherStatusForecastDataNext6Hours
from openapi_client.models.weather_status_forecast_data_next6_hours_details import WeatherStatusForecastDataNext6HoursDetails
from openapi_client.models.weather_status_weather_status_get_forecast200_response import WeatherStatusWeatherStatusGetForecast200Response
from openapi_client.models.weather_status_weather_status_get_forecast200_response_ocs import WeatherStatusWeatherStatusGetForecast200ResponseOcs
from openapi_client.models.weather_status_weather_status_get_location200_response import WeatherStatusWeatherStatusGetLocation200Response
from openapi_client.models.weather_status_weather_status_get_location200_response_ocs import WeatherStatusWeatherStatusGetLocation200ResponseOcs
from openapi_client.models.weather_status_weather_status_get_location200_response_ocs_data import WeatherStatusWeatherStatusGetLocation200ResponseOcsData
from openapi_client.models.weather_status_weather_status_use_personal_address200_response import WeatherStatusWeatherStatusUsePersonalAddress200Response
from openapi_client.models.weather_status_weather_status_use_personal_address200_response_ocs import WeatherStatusWeatherStatusUsePersonalAddress200ResponseOcs
from openapi_client.models.weather_status_weather_status_use_personal_address200_response_ocs_data import WeatherStatusWeatherStatusUsePersonalAddress200ResponseOcsData
