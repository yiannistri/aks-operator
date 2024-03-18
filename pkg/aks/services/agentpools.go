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

type AgentPoolsClientInterface interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, agentPoolName string, parameters containerservice.AgentPool) (containerservice.AgentPoolsCreateOrUpdateFuture, error)
	Delete(ctx context.Context, resourceGroupName string, clusterName string, agentPoolName string) (containerservice.AgentPoolsDeleteFuture, error)
}

type agentPoolClient struct {
	agentPoolClient                     containerservice.AgentPoolsClient
	armcontainerserviceAgentPoolsClient *armcontainerservice.AgentPoolsClient
}

func NewAgentPoolClient(authorizer autorest.Authorizer, baseURL, subscriptionID string, credential *azidentity.ClientSecretCredential, cloud cloud.Configuration) (*agentPoolClient, error) {
	client := containerservice.NewAgentPoolsClientWithBaseURI(baseURL, subscriptionID)
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

	return &agentPoolClient{
		agentPoolClient:                     client,
		armcontainerserviceAgentPoolsClient: clientFactory.NewAgentPoolsClient(),
	}, nil
}

func (cl *agentPoolClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, agentPoolName string, parameters containerservice.AgentPool) (containerservice.AgentPoolsCreateOrUpdateFuture, error) {
	return cl.agentPoolClient.CreateOrUpdate(ctx, resourceGroupName, clusterName, agentPoolName, parameters)
}

func (cl *agentPoolClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, agentPoolName string) (containerservice.AgentPoolsDeleteFuture, error) {
	return cl.agentPoolClient.Delete(ctx, resourceGroupName, clusterName, agentPoolName)
}
