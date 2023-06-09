/*
nextcloud

Testing CoreAvatarApiService

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

func Test_openapi_CoreAvatarApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CoreAvatarApiService CoreAvatarGetAvatar", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string
		var size int64

		resp, httpRes, err := apiClient.CoreAvatarApi.CoreAvatarGetAvatar(context.Background(), userId, size).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreAvatarApiService CoreAvatarGetAvatarDark", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var userId string
		var size int64

		resp, httpRes, err := apiClient.CoreAvatarApi.CoreAvatarGetAvatarDark(context.Background(), userId, size).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
