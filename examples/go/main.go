package main

import (
	"context"
	"encoding/base64"
	"fmt"

	openapi "github.com/nextcloud/api-sdk"
)

func main() {
	ctx := context.Background()
	cfg := &openapi.Configuration{
		Scheme: "http",
		Host:   "localhost",
		DefaultHeader: map[string]string{
			"Authorization": fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte("admin:admin"))),
		},
		Servers: openapi.ServerConfigurations{
			{
				URL: "http://localhost",
			},
		},
	}
	client := openapi.NewAPIClient(cfg)

	response, _, err := client.ProvisioningApiUsersApi.ProvisioningApiUsersGetCurrentUser(ctx).OCSAPIRequest("true").Execute()
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Ocs.Data.Id) // admin
}
