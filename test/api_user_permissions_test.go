/*
Authress

Testing UserPermissionsApi

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

func Test_authress_UserPermissionsApi(t *testing.T) {

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})

	t.Run("Test UserPermissionsApi AuthorizeUser", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId UserId
		var resourceUri string
		var permission Action

		httpRes, err := authressClient.UserPermissions.AuthorizeUser(context.Background(), userId, resourceUri, permission).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserPermissionsApi GetUserPermissionsForResource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId UserId
		var resourceUri string

		resp, httpRes, err := authressClient.UserPermissions.GetUserPermissionsForResource(context.Background(), userId, resourceUri).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserPermissionsApi GetUserResources", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId UserId

		resp, httpRes, err := authressClient.UserPermissions.GetUserResources(context.Background(), userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test UserPermissionsApi GetUserRolesForResource", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var userId UserId
		var resourceUri string

		resp, httpRes, err := authressClient.UserPermissions.GetUserRolesForResource(context.Background(), userId, resourceUri).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
