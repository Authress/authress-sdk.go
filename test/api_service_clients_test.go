/*
Authress

Testing ServiceClientsApi

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package test

import (
	"context"
	"net/url"
	"testing"

	authress "github.com/authress/authress-sdk.go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_authress_ServiceClientsApi(t *testing.T) {

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
	})

	t.Run("Test ServiceClientsApi CreateClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := authressClient.ServiceClients.CreateClient(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsApi DeleteAccessKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string
		var keyId string

		httpRes, err := authressClient.ServiceClients.DeleteAccessKey(context.Background(), clientId, keyId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsApi DeleteClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string

		httpRes, err := authressClient.ServiceClients.DeleteClient(context.Background(), clientId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsApi GetClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string

		resp, httpRes, err := authressClient.ServiceClients.GetClient(context.Background(), clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsApi GetClients", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := authressClient.ServiceClients.GetClients(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsApi RequestAccessKey", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string

		resp, httpRes, err := authressClient.ServiceClients.RequestAccessKey(context.Background(), clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ServiceClientsApi UpdateClient", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var clientId string

		resp, httpRes, err := authressClient.ServiceClients.UpdateClient(context.Background(), clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
