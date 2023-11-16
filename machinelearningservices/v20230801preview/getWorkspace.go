// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// An object that represents a machine learning workspace.
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230801preview:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Azure Machine Learning Workspace Name
	WorkspaceName string `pulumi:"workspaceName"`
}

// An object that represents a machine learning workspace.
type LookupWorkspaceResult struct {
	// The flag to indicate whether to allow public access when behind VNet.
	AllowPublicAccessWhenBehindVnet *bool `pulumi:"allowPublicAccessWhenBehindVnet"`
	// ARM id of the application insights associated with this workspace.
	ApplicationInsights  *string  `pulumi:"applicationInsights"`
	AssociatedWorkspaces []string `pulumi:"associatedWorkspaces"`
	ContainerRegistries  []string `pulumi:"containerRegistries"`
	// ARM id of the container registry associated with this workspace.
	ContainerRegistry *string `pulumi:"containerRegistry"`
	// The description of this workspace.
	Description *string `pulumi:"description"`
	// Url for the discovery service to identify regional endpoints for machine learning experimentation services
	DiscoveryUrl        *string                     `pulumi:"discoveryUrl"`
	EnableDataIsolation *bool                       `pulumi:"enableDataIsolation"`
	Encryption          *EncryptionPropertyResponse `pulumi:"encryption"`
	ExistingWorkspaces  []string                    `pulumi:"existingWorkspaces"`
	// Settings for feature store type workspace.
	FeatureStoreSettings *FeatureStoreSettingsResponse `pulumi:"featureStoreSettings"`
	// The friendly name for this workspace. This name in mutable
	FriendlyName *string `pulumi:"friendlyName"`
	// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
	HbiWorkspace  *bool   `pulumi:"hbiWorkspace"`
	HubResourceId *string `pulumi:"hubResourceId"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Managed service identity (system assigned and/or user assigned identities)
	Identity *ManagedServiceIdentityResponse `pulumi:"identity"`
	// The compute name for image build
	ImageBuildCompute *string `pulumi:"imageBuildCompute"`
	// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
	KeyVault  *string  `pulumi:"keyVault"`
	KeyVaults []string `pulumi:"keyVaults"`
	Kind      *string  `pulumi:"kind"`
	Location  *string  `pulumi:"location"`
	// Managed Network settings for a machine learning workspace.
	ManagedNetwork *ManagedNetworkSettingsResponse `pulumi:"managedNetwork"`
	// The URI associated with this workspace that machine learning flow must point at to set up tracking.
	MlFlowTrackingUri string `pulumi:"mlFlowTrackingUri"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The notebook info of Azure ML workspace.
	NotebookInfo NotebookResourceInfoResponse `pulumi:"notebookInfo"`
	// The user assigned identity resource id that represents the workspace identity.
	PrimaryUserAssignedIdentity *string `pulumi:"primaryUserAssignedIdentity"`
	// The list of private endpoint connections in the workspace.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Count of private connections in the workspace
	PrivateLinkCount int `pulumi:"privateLinkCount"`
	// The current deployment state of workspace resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState string `pulumi:"provisioningState"`
	// Whether requests from Public Network are allowed.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Settings for serverless compute created in the workspace
	ServerlessComputeSettings *ServerlessComputeSettingsResponse `pulumi:"serverlessComputeSettings"`
	// The service managed resource settings.
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettingsResponse `pulumi:"serviceManagedResourcesSettings"`
	// The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace
	ServiceProvisionedResourceGroup string `pulumi:"serviceProvisionedResourceGroup"`
	// The list of shared private link resources in this workspace.
	SharedPrivateLinkResources []SharedPrivateLinkResourceResponse `pulumi:"sharedPrivateLinkResources"`
	// Optional. This field is required to be implemented by the RP because AML is supporting more than one tier
	Sku *SkuResponse `pulumi:"sku"`
	// Retention time in days after workspace get soft deleted.
	SoftDeleteRetentionInDays *int `pulumi:"softDeleteRetentionInDays"`
	// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
	StorageAccount  *string  `pulumi:"storageAccount"`
	StorageAccounts []string `pulumi:"storageAccounts"`
	// If the storage associated with the workspace has hierarchical namespace(HNS) enabled.
	StorageHnsEnabled bool `pulumi:"storageHnsEnabled"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The auth mode used for accessing the system datastores of the workspace.
	SystemDatastoresAuthMode *string           `pulumi:"systemDatastoresAuthMode"`
	Tags                     map[string]string `pulumi:"tags"`
	// The tenant id associated with this workspace.
	TenantId string `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Enabling v1_legacy_mode may prevent you from using features provided by the v2 API.
	V1LegacyMode *bool `pulumi:"v1LegacyMode"`
	// WorkspaceHub's configuration object.
	WorkspaceHubConfig *WorkspaceHubConfigResponse `pulumi:"workspaceHubConfig"`
	// The immutable id associated with this workspace.
	WorkspaceId string `pulumi:"workspaceId"`
}

func LookupWorkspaceOutput(ctx *pulumi.Context, args LookupWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceResult, error) {
			args := v.(LookupWorkspaceArgs)
			r, err := LookupWorkspace(ctx, &args, opts...)
			var s LookupWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceResultOutput)
}

type LookupWorkspaceOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Azure Machine Learning Workspace Name
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}

// An object that represents a machine learning workspace.
type LookupWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceResult)(nil)).Elem()
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutput() LookupWorkspaceResultOutput {
	return o
}

func (o LookupWorkspaceResultOutput) ToLookupWorkspaceResultOutputWithContext(ctx context.Context) LookupWorkspaceResultOutput {
	return o
}

// The flag to indicate whether to allow public access when behind VNet.
func (o LookupWorkspaceResultOutput) AllowPublicAccessWhenBehindVnet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.AllowPublicAccessWhenBehindVnet }).(pulumi.BoolPtrOutput)
}

// ARM id of the application insights associated with this workspace.
func (o LookupWorkspaceResultOutput) ApplicationInsights() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ApplicationInsights }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) AssociatedWorkspaces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []string { return v.AssociatedWorkspaces }).(pulumi.StringArrayOutput)
}

func (o LookupWorkspaceResultOutput) ContainerRegistries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []string { return v.ContainerRegistries }).(pulumi.StringArrayOutput)
}

// ARM id of the container registry associated with this workspace.
func (o LookupWorkspaceResultOutput) ContainerRegistry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ContainerRegistry }).(pulumi.StringPtrOutput)
}

// The description of this workspace.
func (o LookupWorkspaceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Url for the discovery service to identify regional endpoints for machine learning experimentation services
func (o LookupWorkspaceResultOutput) DiscoveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.DiscoveryUrl }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) EnableDataIsolation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.EnableDataIsolation }).(pulumi.BoolPtrOutput)
}

func (o LookupWorkspaceResultOutput) Encryption() EncryptionPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *EncryptionPropertyResponse { return v.Encryption }).(EncryptionPropertyResponsePtrOutput)
}

func (o LookupWorkspaceResultOutput) ExistingWorkspaces() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []string { return v.ExistingWorkspaces }).(pulumi.StringArrayOutput)
}

// Settings for feature store type workspace.
func (o LookupWorkspaceResultOutput) FeatureStoreSettings() FeatureStoreSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *FeatureStoreSettingsResponse { return v.FeatureStoreSettings }).(FeatureStoreSettingsResponsePtrOutput)
}

// The friendly name for this workspace. This name in mutable
func (o LookupWorkspaceResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
func (o LookupWorkspaceResultOutput) HbiWorkspace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.HbiWorkspace }).(pulumi.BoolPtrOutput)
}

func (o LookupWorkspaceResultOutput) HubResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.HubResourceId }).(pulumi.StringPtrOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed service identity (system assigned and/or user assigned identities)
func (o LookupWorkspaceResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The compute name for image build
func (o LookupWorkspaceResultOutput) ImageBuildCompute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ImageBuildCompute }).(pulumi.StringPtrOutput)
}

// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
func (o LookupWorkspaceResultOutput) KeyVault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.KeyVault }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) KeyVaults() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []string { return v.KeyVaults }).(pulumi.StringArrayOutput)
}

func (o LookupWorkspaceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Managed Network settings for a machine learning workspace.
func (o LookupWorkspaceResultOutput) ManagedNetwork() ManagedNetworkSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ManagedNetworkSettingsResponse { return v.ManagedNetwork }).(ManagedNetworkSettingsResponsePtrOutput)
}

// The URI associated with this workspace that machine learning flow must point at to set up tracking.
func (o LookupWorkspaceResultOutput) MlFlowTrackingUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.MlFlowTrackingUri }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The notebook info of Azure ML workspace.
func (o LookupWorkspaceResultOutput) NotebookInfo() NotebookResourceInfoResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) NotebookResourceInfoResponse { return v.NotebookInfo }).(NotebookResourceInfoResponseOutput)
}

// The user assigned identity resource id that represents the workspace identity.
func (o LookupWorkspaceResultOutput) PrimaryUserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.PrimaryUserAssignedIdentity }).(pulumi.StringPtrOutput)
}

// The list of private endpoint connections in the workspace.
func (o LookupWorkspaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Count of private connections in the workspace
func (o LookupWorkspaceResultOutput) PrivateLinkCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) int { return v.PrivateLinkCount }).(pulumi.IntOutput)
}

// The current deployment state of workspace resource. The provisioningState is to indicate states for resource provisioning.
func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether requests from Public Network are allowed.
func (o LookupWorkspaceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Settings for serverless compute created in the workspace
func (o LookupWorkspaceResultOutput) ServerlessComputeSettings() ServerlessComputeSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ServerlessComputeSettingsResponse { return v.ServerlessComputeSettings }).(ServerlessComputeSettingsResponsePtrOutput)
}

// The service managed resource settings.
func (o LookupWorkspaceResultOutput) ServiceManagedResourcesSettings() ServiceManagedResourcesSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ServiceManagedResourcesSettingsResponse {
		return v.ServiceManagedResourcesSettings
	}).(ServiceManagedResourcesSettingsResponsePtrOutput)
}

// The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace
func (o LookupWorkspaceResultOutput) ServiceProvisionedResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ServiceProvisionedResourceGroup }).(pulumi.StringOutput)
}

// The list of shared private link resources in this workspace.
func (o LookupWorkspaceResultOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []SharedPrivateLinkResourceResponse { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

// Optional. This field is required to be implemented by the RP because AML is supporting more than one tier
func (o LookupWorkspaceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Retention time in days after workspace get soft deleted.
func (o LookupWorkspaceResultOutput) SoftDeleteRetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *int { return v.SoftDeleteRetentionInDays }).(pulumi.IntPtrOutput)
}

// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
func (o LookupWorkspaceResultOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) StorageAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []string { return v.StorageAccounts }).(pulumi.StringArrayOutput)
}

// If the storage associated with the workspace has hierarchical namespace(HNS) enabled.
func (o LookupWorkspaceResultOutput) StorageHnsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) bool { return v.StorageHnsEnabled }).(pulumi.BoolOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The auth mode used for accessing the system datastores of the workspace.
func (o LookupWorkspaceResultOutput) SystemDatastoresAuthMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.SystemDatastoresAuthMode }).(pulumi.StringPtrOutput)
}

func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The tenant id associated with this workspace.
func (o LookupWorkspaceResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

// Enabling v1_legacy_mode may prevent you from using features provided by the v2 API.
func (o LookupWorkspaceResultOutput) V1LegacyMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.V1LegacyMode }).(pulumi.BoolPtrOutput)
}

// WorkspaceHub's configuration object.
func (o LookupWorkspaceResultOutput) WorkspaceHubConfig() WorkspaceHubConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *WorkspaceHubConfigResponse { return v.WorkspaceHubConfig }).(WorkspaceHubConfigResponsePtrOutput)
}

// The immutable id associated with this workspace.
func (o LookupWorkspaceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
