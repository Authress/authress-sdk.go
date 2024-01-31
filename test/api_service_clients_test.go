/*
Authress

Testing ServiceClientsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package authress

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "//"
)

func Test_authress_ServiceClientsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ServiceClientsAPIService CreateClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServiceClientsAPI.CreateClient(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsAPIService DeleteAccessKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clientId string
		var keyId string

		httpRes, err := apiClient.ServiceClientsAPI.DeleteAccessKey(context.Background(), clientId, keyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsAPIService DeleteClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clientId string

		httpRes, err := apiClient.ServiceClientsAPI.DeleteClient(context.Background(), clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsAPIService GetClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clientId string

		resp, httpRes, err := apiClient.ServiceClientsAPI.GetClient(context.Background(), clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsAPIService GetClients", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ServiceClientsAPI.GetClients(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsAPIService RequestAccessKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clientId string

		resp, httpRes, err := apiClient.ServiceClientsAPI.RequestAccessKey(context.Background(), clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsAPIService UpdateClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var clientId string

		resp, httpRes, err := apiClient.ServiceClientsAPI.UpdateClient(context.Background(), clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
