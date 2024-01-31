/*
Authress

Testing ApplicationsAPIService

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

func Test_authress_ApplicationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ApplicationsAPIService DelegateUserLogin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var applicationId string
		var userId UserId

		resp, httpRes, err := apiClient.ApplicationsAPI.DelegateUserLogin(context.Background(), applicationId, userId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
