package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/operationalinsights/armoperationalinsights"
	"github.com/Azure/azure-sdk-for-go/services/operationalinsights/mgmt/2020-08-01/operationalinsights"
	"github.com/Azure/go-autorest/autorest"
)

type WorkplacesClientInterface interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters operationalinsights.Workspace) (operationalinsights.WorkspacesCreateOrUpdateFuture, error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (operationalinsights.Workspace, error)
	AsyncCreateUpdateResult(asyncRet operationalinsights.WorkspacesCreateOrUpdateFuture) (operationalinsights.Workspace, error)
}

type workplacesClient struct {
	workplacesClient                       operationalinsights.WorkspacesClient
	armoperationalinsightsWorkspacesClient *armoperationalinsights.WorkspacesClient
}

func NewWorkplacesClient(authorizer autorest.Authorizer, baseURL, subscriptionID string, credential *azidentity.ClientSecretCredential, cloud cloud.Configuration) (*workplacesClient, error) {
	client := operationalinsights.NewWorkspacesClientWithBaseURI(baseURL, subscriptionID)
	client.Authorizer = authorizer

	options := arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud,
		},
	}
	clientFactory, err := armoperationalinsights.NewClientFactory(subscriptionID, credential, &options)
	if err != nil {
		return nil, err
	}

	return &workplacesClient{
		workplacesClient:                       client,
		armoperationalinsightsWorkspacesClient: clientFactory.NewWorkspacesClient(),
	}, nil
}

func (c *workplacesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, parameters operationalinsights.Workspace) (operationalinsights.WorkspacesCreateOrUpdateFuture, error) {
	return c.workplacesClient.CreateOrUpdate(ctx, resourceGroupName, workspaceName, parameters)
}

func (c *workplacesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string) (operationalinsights.Workspace, error) {
	return c.workplacesClient.Get(ctx, resourceGroupName, workspaceName)
}

func (c *workplacesClient) AsyncCreateUpdateResult(asyncRet operationalinsights.WorkspacesCreateOrUpdateFuture) (operationalinsights.Workspace, error) {
	return asyncRet.Result(c.workplacesClient)
}
