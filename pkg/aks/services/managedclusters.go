package services

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/containerservice/armcontainerservice/v4"
	"github.com/Azure/azure-sdk-for-go/services/containerservice/mgmt/2020-11-01/containerservice"
	"github.com/Azure/go-autorest/autorest"
)

type ManagedClustersClientInterface interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters containerservice.ManagedCluster) (containerservice.ManagedClustersCreateOrUpdateFuture, error)
	Get(ctx context.Context, resourceGroupName string, clusterName string) (containerservice.ManagedCluster, error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string) (containerservice.ManagedClustersDeleteFuture, error)
	WaitForTaskCompletion(context.Context, containerservice.ManagedClustersDeleteFuture) error
	GetAccessProfile(ctx context.Context, resourceGroupName string, resourceName string, roleName string) (containerservice.ManagedClusterAccessProfile, error)
	UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.TagsObject) (containerservice.ManagedClustersUpdateTagsFuture, error)
	AsyncUpdateTagsResult(asyncRet containerservice.ManagedClustersUpdateTagsFuture) (containerservice.ManagedCluster, error)
}

type managedClustersClient struct {
	managedClustersClient                    containerservice.ManagedClustersClient
	armcontainerserviceManagedClustersClient *armcontainerservice.ManagedClustersClient
}

func NewManagedClustersClient(authorizer autorest.Authorizer, baseURL, subscriptionID string, credential *azidentity.ClientSecretCredential, cloud cloud.Configuration) (*managedClustersClient, error) {
	client := containerservice.NewManagedClustersClientWithBaseURI(baseURL, subscriptionID)
	client.Authorizer = authorizer

	options := arm.ClientOptions{
		ClientOptions: azcore.ClientOptions{
			Cloud: cloud,
		},
	}
	clientFactory, err := armcontainerservice.NewClientFactory(subscriptionID, credential, &options)
	if err != nil {
		return nil, err
	}

	return &managedClustersClient{
		managedClustersClient:                    client,
		armcontainerserviceManagedClustersClient: clientFactory.NewManagedClustersClient(),
	}, nil
}

func (cl *managedClustersClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, parameters containerservice.ManagedCluster) (containerservice.ManagedClustersCreateOrUpdateFuture, error) {
	return cl.managedClustersClient.CreateOrUpdate(ctx, resourceGroupName, clusterName, parameters)
}

func (cl *managedClustersClient) Get(ctx context.Context, resourceGroupName string, clusterName string) (containerservice.ManagedCluster, error) {
	return cl.managedClustersClient.Get(ctx, resourceGroupName, clusterName)
}

func (cl *managedClustersClient) Delete(ctx context.Context, resourceGroupName string, clusterName string) (containerservice.ManagedClustersDeleteFuture, error) {
	return cl.managedClustersClient.Delete(ctx, resourceGroupName, clusterName)
}

func (cl *managedClustersClient) WaitForTaskCompletion(ctx context.Context, future containerservice.ManagedClustersDeleteFuture) error {
	return future.WaitForCompletionRef(ctx, cl.managedClustersClient.Client)
}

func (cl *managedClustersClient) GetAccessProfile(ctx context.Context, resourceGroupName string, resourceName string, roleName string) (containerservice.ManagedClusterAccessProfile, error) {
	return cl.managedClustersClient.GetAccessProfile(ctx, resourceGroupName, resourceName, roleName)
}

func (cl *managedClustersClient) UpdateTags(ctx context.Context, resourceGroupName string, resourceName string, parameters containerservice.TagsObject) (containerservice.ManagedClustersUpdateTagsFuture, error) {
	return cl.managedClustersClient.UpdateTags(ctx, resourceGroupName, resourceName, parameters)
}

func (cl *managedClustersClient) AsyncUpdateTagsResult(asyncRet containerservice.ManagedClustersUpdateTagsFuture) (containerservice.ManagedCluster, error) {
	return asyncRet.Result(cl.managedClustersClient)
}
