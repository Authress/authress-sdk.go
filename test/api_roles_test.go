/*
Authress

Testing RolesApi

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

func Test_authress_RolesApi(t *testing.T) {

	url, _ := url.Parse("https://authress.company.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})

	t.Run("Test RolesApi CreateRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := authressClient.Roles.CreateRole(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesApi DeleteRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleId string

		httpRes, err := authressClient.Roles.DeleteRole(context.Background(), roleId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesApi GetRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleId string

		resp, httpRes, err := authressClient.Roles.GetRole(context.Background(), roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesApi GetRoles", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := authressClient.Roles.GetRoles(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolesApi UpdateRole", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var roleId string

		resp, httpRes, err := authressClient.Roles.UpdateRole(context.Background(), roleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
