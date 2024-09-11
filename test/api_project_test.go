/*
Ory APIs

Testing ProjectAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/ory/client-go"
)

func Test_client_ProjectAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectAPIService CreateOrganization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.CreateOrganization(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService CreateProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectAPI.CreateProject(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService CreateProjectApiKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string

		resp, httpRes, err := apiClient.ProjectAPI.CreateProjectApiKey(context.Background(), project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService DeleteOrganization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var organizationId string

		httpRes, err := apiClient.ProjectAPI.DeleteOrganization(context.Background(), projectId, organizationId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService DeleteProjectApiKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var tokenId string

		httpRes, err := apiClient.ProjectAPI.DeleteProjectApiKey(context.Background(), project, tokenId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService GetOrganization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var organizationId string

		resp, httpRes, err := apiClient.ProjectAPI.GetOrganization(context.Background(), projectId, organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService GetProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.GetProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService GetProjectMembers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string

		resp, httpRes, err := apiClient.ProjectAPI.GetProjectMembers(context.Background(), project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService ListOrganizations", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.ListOrganizations(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService ListProjectApiKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string

		resp, httpRes, err := apiClient.ProjectAPI.ListProjectApiKeys(context.Background(), project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService ListProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectAPI.ListProjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService PatchProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.PatchProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService PatchProjectWithRevision", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var revisionId string

		resp, httpRes, err := apiClient.ProjectAPI.PatchProjectWithRevision(context.Background(), projectId, revisionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService PurgeProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		httpRes, err := apiClient.ProjectAPI.PurgeProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService RemoveProjectMember", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var project string
		var member string

		httpRes, err := apiClient.ProjectAPI.RemoveProjectMember(context.Background(), project, member).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService SetProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string

		resp, httpRes, err := apiClient.ProjectAPI.SetProject(context.Background(), projectId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectAPIService UpdateOrganization", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var projectId string
		var organizationId string

		resp, httpRes, err := apiClient.ProjectAPI.UpdateOrganization(context.Background(), projectId, organizationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
