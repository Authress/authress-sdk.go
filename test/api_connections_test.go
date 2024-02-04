/*
Authress

Testing ConnectionsApi

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package authress

import (
	"context"
	"testing"

	authress "github.com/authress/authress-sdk.go"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_authress_ConnectionsApi(t *testing.T) {

	authressClient := authress.AuthressClient.New(AuthressSettings{})

	t.Run("Test ConnectionsApi CreateConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Connections.CreateConnection(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectionsApi DeleteConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectionId string

		httpRes, err := apiClient.Connections.DeleteConnection(context.Background(), connectionId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectionsApi GetConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectionId string

		resp, httpRes, err := apiClient.Connections.GetConnection(context.Background(), connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectionsApi GetConnectionCredentials", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectionId string
		var userId UserId

		resp, httpRes, err := apiClient.Connections.GetConnectionCredentials(context.Background(), connectionId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectionsApi GetConnections", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Connections.GetConnections(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ConnectionsApi UpdateConnection", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var connectionId string

		resp, httpRes, err := apiClient.Connections.UpdateConnection(context.Background(), connectionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
