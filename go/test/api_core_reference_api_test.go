/*
nextcloud

Testing CoreReferenceApiApiService

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

func Test_openapi_CoreReferenceApiApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test CoreReferenceApiApiService CoreReferenceApiExtract", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CoreReferenceApiApi.CoreReferenceApiExtract(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreReferenceApiApiService CoreReferenceApiGetProvidersInfo", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CoreReferenceApiApi.CoreReferenceApiGetProvidersInfo(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreReferenceApiApiService CoreReferenceApiResolve", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CoreReferenceApiApi.CoreReferenceApiResolve(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreReferenceApiApiService CoreReferenceApiResolveOne", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.CoreReferenceApiApi.CoreReferenceApiResolveOne(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test CoreReferenceApiApiService CoreReferenceApiTouchProvider", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var providerId string

		resp, httpRes, err := apiClient.CoreReferenceApiApi.CoreReferenceApiTouchProvider(context.Background(), providerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
