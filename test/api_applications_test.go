/*
Authress

Testing ApplicationsApi

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

func Test_authress_ApplicationsApi(t *testing.T) {

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url, 
	})

	t.Run("Test ApplicationsApi DelegateUserLogin", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var applicationId string
		var userId string

		resp, httpRes, err := authressClient.Applications.DelegateUserLogin(context.Background(), applicationId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
