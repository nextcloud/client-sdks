/*
nextcloud

Testing ThemingUserThemeApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/nextcloud/api-sdk"
)

func Test_openapi_ThemingUserThemeApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ThemingUserThemeApiService ThemingUserThemeDeleteBackground", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ThemingUserThemeApi.ThemingUserThemeDeleteBackground(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemingUserThemeApiService ThemingUserThemeDisableTheme", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var themeId string

		resp, httpRes, err := apiClient.ThemingUserThemeApi.ThemingUserThemeDisableTheme(context.Background(), themeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemingUserThemeApiService ThemingUserThemeEnableTheme", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var themeId string

		resp, httpRes, err := apiClient.ThemingUserThemeApi.ThemingUserThemeEnableTheme(context.Background(), themeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemingUserThemeApiService ThemingUserThemeGetBackground", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ThemingUserThemeApi.ThemingUserThemeGetBackground(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemingUserThemeApiService ThemingUserThemeSetBackground", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var type_ string

		resp, httpRes, err := apiClient.ThemingUserThemeApi.ThemingUserThemeSetBackground(context.Background(), type_).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}