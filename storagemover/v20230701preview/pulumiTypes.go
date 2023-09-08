// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

var _ = utilities.GetEnvOrDefault

type AgentPropertiesResponseErrorDetails struct {
	// Error code reported by Agent
	Code *string `pulumi:"code"`
	// Expanded description of reported error code
	Message *string `pulumi:"message"`
}

type AgentPropertiesResponseErrorDetailsOutput struct{ *pulumi.OutputState }

func (AgentPropertiesResponseErrorDetailsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgentPropertiesResponseErrorDetails)(nil)).Elem()
}

func (o AgentPropertiesResponseErrorDetailsOutput) ToAgentPropertiesResponseErrorDetailsOutput() AgentPropertiesResponseErrorDetailsOutput {
	return o
}

func (o AgentPropertiesResponseErrorDetailsOutput) ToAgentPropertiesResponseErrorDetailsOutputWithContext(ctx context.Context) AgentPropertiesResponseErrorDetailsOutput {
	return o
}

// Error code reported by Agent
func (o AgentPropertiesResponseErrorDetailsOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPropertiesResponseErrorDetails) *string { return v.Code }).(pulumi.StringPtrOutput)
}

// Expanded description of reported error code
func (o AgentPropertiesResponseErrorDetailsOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgentPropertiesResponseErrorDetails) *string { return v.Message }).(pulumi.StringPtrOutput)
}

// The Azure Key Vault secret URIs which store the credentials.
type AzureKeyVaultSmbCredentials struct {
	// The Azure Key Vault secret URI which stores the password. Use empty string to clean-up existing value.
	PasswordUri *string `pulumi:"passwordUri"`
	// The Credentials type.
	// Expected value is 'AzureKeyVaultSmb'.
	Type string `pulumi:"type"`
	// The Azure Key Vault secret URI which stores the username. Use empty string to clean-up existing value.
	UsernameUri *string `pulumi:"usernameUri"`
}

// The Azure Key Vault secret URIs which store the credentials.
type AzureKeyVaultSmbCredentialsResponse struct {
	// The Azure Key Vault secret URI which stores the password. Use empty string to clean-up existing value.
	PasswordUri *string `pulumi:"passwordUri"`
	// The Credentials type.
	// Expected value is 'AzureKeyVaultSmb'.
	Type string `pulumi:"type"`
	// The Azure Key Vault secret URI which stores the username. Use empty string to clean-up existing value.
	UsernameUri *string `pulumi:"usernameUri"`
}

// The properties of Azure Storage blob container endpoint.
type AzureStorageBlobContainerEndpointProperties struct {
	// The name of the Storage blob container that is the target destination.
	BlobContainerName string `pulumi:"blobContainerName"`
	// A description for the Endpoint.
	Description *string `pulumi:"description"`
	// The Endpoint resource type.
	// Expected value is 'AzureStorageBlobContainer'.
	EndpointType string `pulumi:"endpointType"`
	// The Azure Resource ID of the storage account that is the target destination.
	StorageAccountResourceId string `pulumi:"storageAccountResourceId"`
}

// The properties of Azure Storage blob container endpoint.
type AzureStorageBlobContainerEndpointPropertiesResponse struct {
	// The name of the Storage blob container that is the target destination.
	BlobContainerName string `pulumi:"blobContainerName"`
	// A description for the Endpoint.
	Description *string `pulumi:"description"`
	// The Endpoint resource type.
	// Expected value is 'AzureStorageBlobContainer'.
	EndpointType string `pulumi:"endpointType"`
	// The provisioning state of this resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The Azure Resource ID of the storage account that is the target destination.
	StorageAccountResourceId string `pulumi:"storageAccountResourceId"`
}

// The properties of Azure Storage SMB file share endpoint.
type AzureStorageSmbFileShareEndpointProperties struct {
	// A description for the Endpoint.
	Description *string `pulumi:"description"`
	// The Endpoint resource type.
	// Expected value is 'AzureStorageSmbFileShare'.
	EndpointType string `pulumi:"endpointType"`
	// The name of the Azure Storage file share.
	FileShareName string `pulumi:"fileShareName"`
	// The Azure Resource ID of the storage account.
	StorageAccountResourceId string `pulumi:"storageAccountResourceId"`
}

// The properties of Azure Storage SMB file share endpoint.
type AzureStorageSmbFileShareEndpointPropertiesResponse struct {
	// A description for the Endpoint.
	Description *string `pulumi:"description"`
	// The Endpoint resource type.
	// Expected value is 'AzureStorageSmbFileShare'.
	EndpointType string `pulumi:"endpointType"`
	// The name of the Azure Storage file share.
	FileShareName string `pulumi:"fileShareName"`
	// The provisioning state of this resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The Azure Resource ID of the storage account.
	StorageAccountResourceId string `pulumi:"storageAccountResourceId"`
}

// The properties of NFS share endpoint.
type NfsMountEndpointProperties struct {
	// A description for the Endpoint.
	Description *string `pulumi:"description"`
	// The Endpoint resource type.
	// Expected value is 'NfsMount'.
	EndpointType string `pulumi:"endpointType"`
	// The directory being exported from the server.
	Export string `pulumi:"export"`
	// The host name or IP address of the server exporting the file system.
	Host string `pulumi:"host"`
	// The NFS protocol version.
	NfsVersion *string `pulumi:"nfsVersion"`
}

// The properties of NFS share endpoint.
type NfsMountEndpointPropertiesResponse struct {
	// A description for the Endpoint.
	Description *string `pulumi:"description"`
	// The Endpoint resource type.
	// Expected value is 'NfsMount'.
	EndpointType string `pulumi:"endpointType"`
	// The directory being exported from the server.
	Export string `pulumi:"export"`
	// The host name or IP address of the server exporting the file system.
	Host string `pulumi:"host"`
	// The NFS protocol version.
	NfsVersion *string `pulumi:"nfsVersion"`
	// The provisioning state of this resource.
	ProvisioningState string `pulumi:"provisioningState"`
}

// The properties of SMB share endpoint.
type SmbMountEndpointProperties struct {
	// The Azure Key Vault secret URIs which store the required credentials to access the SMB share.
	Credentials *AzureKeyVaultSmbCredentials `pulumi:"credentials"`
	// A description for the Endpoint.
	Description *string `pulumi:"description"`
	// The Endpoint resource type.
	// Expected value is 'SmbMount'.
	EndpointType string `pulumi:"endpointType"`
	// The host name or IP address of the server exporting the file system.
	Host string `pulumi:"host"`
	// The name of the SMB share being exported from the server.
	ShareName string `pulumi:"shareName"`
}

// The properties of SMB share endpoint.
type SmbMountEndpointPropertiesResponse struct {
	// The Azure Key Vault secret URIs which store the required credentials to access the SMB share.
	Credentials *AzureKeyVaultSmbCredentialsResponse `pulumi:"credentials"`
	// A description for the Endpoint.
	Description *string `pulumi:"description"`
	// The Endpoint resource type.
	// Expected value is 'SmbMount'.
	EndpointType string `pulumi:"endpointType"`
	// The host name or IP address of the server exporting the file system.
	Host string `pulumi:"host"`
	// The provisioning state of this resource.
	ProvisioningState string `pulumi:"provisioningState"`
	// The name of the SMB share being exported from the server.
	ShareName string `pulumi:"shareName"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponse struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *string `pulumi:"createdAt"`
	// The identity that created the resource.
	CreatedBy *string `pulumi:"createdBy"`
	// The type of identity that created the resource.
	CreatedByType *string `pulumi:"createdByType"`
	// The timestamp of resource last modification (UTC)
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// The identity that last modified the resource.
	LastModifiedBy *string `pulumi:"lastModifiedBy"`
	// The type of identity that last modified the resource.
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

// The timestamp of resource creation (UTC).
func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// The identity that created the resource.
func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that created the resource.
func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

// The timestamp of resource last modification (UTC)
func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

// The identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

// The type of identity that last modified the resource.
func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AgentPropertiesResponseErrorDetailsOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}