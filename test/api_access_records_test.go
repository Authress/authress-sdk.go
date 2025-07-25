/*
Authress

Testing AccessRecordsApi

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

func Test_authress_AccessRecordsApi(t *testing.T) {

	url, _ := url.Parse("https://auth.yourdomain.com")
	authressClient := authress.NewAuthressClient(authress.AuthressSettings{
		AuthressApiUrl: url,
	})

	t.Run("Test AccessRecordsApi CreateClaim", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := authressClient.AccessRecords.CreateClaim(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRecordsApi CreateRecord", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := authressClient.AccessRecords.CreateRecord(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRecordsApi CreateRequest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := authressClient.AccessRecords.CreateRequest(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRecordsApi DeleteRecord", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var recordId string

		httpRes, err := authressClient.AccessRecords.DeleteRecord(context.Background(), recordId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRecordsApi DeleteRequest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var requestId string

		httpRes, err := authressClient.AccessRecords.DeleteRequest(context.Background(), requestId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRecordsApi GetRecord", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var recordId string

		resp, httpRes, err := authressClient.AccessRecords.GetRecord(context.Background(), recordId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRecordsApi GetRecords", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := authressClient.AccessRecords.GetRecords(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRecordsApi GetRequest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var requestId string

		resp, httpRes, err := authressClient.AccessRecords.GetRequest(context.Background(), requestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRecordsApi GetRequests", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := authressClient.AccessRecords.GetRequests(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRecordsApi RespondToAccessRequest", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var requestId string

		resp, httpRes, err := authressClient.AccessRecords.RespondToAccessRequest(context.Background(), requestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccessRecordsApi UpdateRecord", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var recordId string

		httpRes, err := authressClient.AccessRecords.UpdateRecord(context.Background(), recordId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
