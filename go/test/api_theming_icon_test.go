/*
nextcloud

Testing ThemingIconApiService

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

func Test_openapi_ThemingIconApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ThemingIconApiService ThemingIconGetFavicon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string

		resp, httpRes, err := apiClient.ThemingIconApi.ThemingIconGetFavicon(context.Background(), app).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemingIconApiService ThemingIconGetThemedIcon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string
		var image string

		resp, httpRes, err := apiClient.ThemingIconApi.ThemingIconGetThemedIcon(context.Background(), app, image).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ThemingIconApiService ThemingIconGetTouchIcon", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var app string

		resp, httpRes, err := apiClient.ThemingIconApi.ThemingIconGetTouchIcon(context.Background(), app).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}