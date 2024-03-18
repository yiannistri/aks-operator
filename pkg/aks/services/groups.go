package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-10-01/resources"
	"github.com/Azure/go-autorest/autorest"
)

type ResourceGroupsClientInterface interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceGroup resources.Group) (resources.Group, error)
	CheckExistence(ctx context.Context, resourceGroupName string) (result autorest.Response, err error)
	Delete(ctx context.Context, resourceGroupName string) (result resources.GroupsDeleteFuture, err error)
}

type resourceGroupsClient struct {
	groupsClient             resources.GroupsClient
	armresourcesGroupsClient *armresources.ResourceGroupsClient
}

func NewResourceGroupsClient(authorizer autorest.Authorizer, baseURL, subscriptionID string, credential *azidentity.ClientSecretCredential, cloud cloud.Configuration) (*resourceGroupsClient, error) {
	client := resources.NewGroupsClientWithBaseURI(baseURL, subscriptionID)
	client.Authorizer = authorizer

	options := arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud,
		},
	}
	clientFactory, err := armresources.NewClientFactory(subscriptionID, credential, &options)
	if err != nil {
		return nil, err
	}

	return &resourceGroupsClient{
		groupsClient:             client,
		armresourcesGroupsClient: clientFactory.NewResourceGroupsClient(),
	}, nil
}

func (cl *resourceGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceGroup resources.Group) (resources.Group, error) {
	return cl.groupsClient.CreateOrUpdate(ctx, resourceGroupName, resourceGroup)
}

func (cl *resourceGroupsClient) CheckExistence(ctx context.Context, resourceGroupName string) (result autorest.Response, err error) {
	return cl.groupsClient.CheckExistence(ctx, resourceGroupName)
}

func (cl *resourceGroupsClient) Delete(ctx context.Context, resourceGroupName string) (result resources.GroupsDeleteFuture, err error) {
	return cl.groupsClient.Delete(ctx, resourceGroupName)
}
