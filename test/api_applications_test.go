/*
Authress

Testing ApplicationsApi

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

func Test_authress_ApplicationsApi(t *testing.T) {

	authressClient := authress.AuthressClient.New(AuthressSettings{})

	t.Run("Test ApplicationsApi DelegateUserLogin", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string
		var userId UserId

		resp, httpRes, err := apiClient.Applications.DelegateUserLogin(context.Background(), applicationId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
