package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resources/armresources"
)

type ResourceGroupsClientInterface interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceGroup armresources.ResourceGroup) error
	CheckExistence(ctx context.Context, resourceGroupName string) (armresources.ResourceGroupsClientCheckExistenceResponse, error)
}

type resourceGroupsClient struct {
	armresourcesGroupsClient *armresources.ResourceGroupsClient
}

func NewResourceGroupsClient(subscriptionID string, credential *azidentity.ClientSecretCredential, cloud cloud.Configuration) (*resourceGroupsClient, error) {
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
		armresourcesGroupsClient: clientFactory.NewResourceGroupsClient(),
	}, nil
}

func (cl *resourceGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceGroup armresources.ResourceGroup) error {
	_, err := cl.armresourcesGroupsClient.CreateOrUpdate(ctx, resourceGroupName, resourceGroup, nil)
	return err
}

func (cl *resourceGroupsClient) CheckExistence(ctx context.Context, resourceGroupName string) (armresources.ResourceGroupsClientCheckExistenceResponse, error) {
	return cl.armresourcesGroupsClient.CheckExistence(ctx, resourceGroupName, nil)
}
