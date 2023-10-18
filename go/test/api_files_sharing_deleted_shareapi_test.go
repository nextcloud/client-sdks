/*
nextcloud

Testing FilesSharingDeletedShareapiApiService

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

func Test_openapi_FilesSharingDeletedShareapiApiService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FilesSharingDeletedShareapiApiService FilesSharingDeletedShareapiList", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FilesSharingDeletedShareapiApi.FilesSharingDeletedShareapiList(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FilesSharingDeletedShareapiApiService FilesSharingDeletedShareapiUndelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var id string

		resp, httpRes, err := apiClient.FilesSharingDeletedShareapiApi.FilesSharingDeletedShareapiUndelete(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}