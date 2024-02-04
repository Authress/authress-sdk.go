/*
Authress

Testing UsersApi

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

func Test_authress_UsersApi(t *testing.T) {

	authressClient := authress.AuthressClient.New(AuthressSettings{})

	t.Run("Test UsersApi DeleteUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId UserId

		httpRes, err := apiClient.Users.DeleteUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApi GetUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId UserId

		resp, httpRes, err := apiClient.Users.GetUser(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UsersApi GetUsers", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.Users.GetUsers(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
