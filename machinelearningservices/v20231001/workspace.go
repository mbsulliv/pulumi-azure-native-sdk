// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231001

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// An object that represents a machine learning workspace.
type Workspace struct {
	pulumi.CustomResourceState

	// The flag to indicate whether to allow public access when behind VNet.
	AllowPublicAccessWhenBehindVnet pulumi.BoolPtrOutput `pulumi:"allowPublicAccessWhenBehindVnet"`
	// ARM id of the application insights associated with this workspace.
	ApplicationInsights pulumi.StringPtrOutput `pulumi:"applicationInsights"`
	// ARM id of the container registry associated with this workspace.
	ContainerRegistry pulumi.StringPtrOutput `pulumi:"containerRegistry"`
	// The description of this workspace.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Url for the discovery service to identify regional endpoints for machine learning experimentation services
	DiscoveryUrl pulumi.StringPtrOutput `pulumi:"discoveryUrl"`
	// The encryption settings of Azure ML workspace.
	Encryption EncryptionPropertyResponsePtrOutput `pulumi:"encryption"`
	// Settings for feature store type workspace.
	FeatureStoreSettings FeatureStoreSettingsResponsePtrOutput `pulumi:"featureStoreSettings"`
	// The friendly name for this workspace. This name in mutable
	FriendlyName pulumi.StringPtrOutput `pulumi:"friendlyName"`
	// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
	HbiWorkspace pulumi.BoolPtrOutput `pulumi:"hbiWorkspace"`
	// The identity of the resource.
	Identity ManagedServiceIdentityResponsePtrOutput `pulumi:"identity"`
	// The compute name for image build
	ImageBuildCompute pulumi.StringPtrOutput `pulumi:"imageBuildCompute"`
	// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
	KeyVault pulumi.StringPtrOutput `pulumi:"keyVault"`
	// Specifies the location of the resource.
	Location pulumi.StringPtrOutput `pulumi:"location"`
	// Managed Network settings for a machine learning workspace.
	ManagedNetwork ManagedNetworkSettingsResponsePtrOutput `pulumi:"managedNetwork"`
	// The URI associated with this workspace that machine learning flow must point at to set up tracking.
	MlFlowTrackingUri pulumi.StringOutput `pulumi:"mlFlowTrackingUri"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The notebook info of Azure ML workspace.
	NotebookInfo NotebookResourceInfoResponseOutput `pulumi:"notebookInfo"`
	// The user assigned identity resource id that represents the workspace identity.
	PrimaryUserAssignedIdentity pulumi.StringPtrOutput `pulumi:"primaryUserAssignedIdentity"`
	// The list of private endpoint connections in the workspace.
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Count of private connections in the workspace
	PrivateLinkCount pulumi.IntOutput `pulumi:"privateLinkCount"`
	// The current deployment state of workspace resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Whether requests from Public Network are allowed.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Settings for serverless compute created in the workspace
	ServerlessComputeSettings ServerlessComputeSettingsResponsePtrOutput `pulumi:"serverlessComputeSettings"`
	// The service managed resource settings.
	ServiceManagedResourcesSettings ServiceManagedResourcesSettingsResponsePtrOutput `pulumi:"serviceManagedResourcesSettings"`
	// The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace
	ServiceProvisionedResourceGroup pulumi.StringOutput `pulumi:"serviceProvisionedResourceGroup"`
	// The list of shared private link resources in this workspace.
	SharedPrivateLinkResources SharedPrivateLinkResourceResponseArrayOutput `pulumi:"sharedPrivateLinkResources"`
	// The sku of the workspace.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
	StorageAccount pulumi.StringPtrOutput `pulumi:"storageAccount"`
	// If the storage associated with the workspace has hierarchical namespace(HNS) enabled.
	StorageHnsEnabled pulumi.BoolOutput `pulumi:"storageHnsEnabled"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The tenant id associated with this workspace.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Enabling v1_legacy_mode may prevent you from using features provided by the v2 API.
	V1LegacyMode pulumi.BoolPtrOutput `pulumi:"v1LegacyMode"`
	// The immutable id associated with this workspace.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.AllowPublicAccessWhenBehindVnet == nil {
		args.AllowPublicAccessWhenBehindVnet = pulumi.BoolPtr(false)
	}
	if args.HbiWorkspace == nil {
		args.HbiWorkspace = pulumi.BoolPtr(false)
	}
	if args.V1LegacyMode == nil {
		args.V1LegacyMode = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20180301preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20181119:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190501:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20190601:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20191101:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200101:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200218preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200301:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200501preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200515preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200601:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200801:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210101:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210701:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220101preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220201preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220501:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221201preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230201preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230401preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230601preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20230801preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20231001:Workspace", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkspace gets an existing Workspace resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkspace(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkspaceState, opts ...pulumi.ResourceOption) (*Workspace, error) {
	var resource Workspace
	err := ctx.ReadResource("azure-native:machinelearningservices/v20231001:Workspace", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Workspace resources.
type workspaceState struct {
}

type WorkspaceState struct {
}

func (WorkspaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceState)(nil)).Elem()
}

type workspaceArgs struct {
	// The flag to indicate whether to allow public access when behind VNet.
	AllowPublicAccessWhenBehindVnet *bool `pulumi:"allowPublicAccessWhenBehindVnet"`
	// ARM id of the application insights associated with this workspace.
	ApplicationInsights *string `pulumi:"applicationInsights"`
	// ARM id of the container registry associated with this workspace.
	ContainerRegistry *string `pulumi:"containerRegistry"`
	// The description of this workspace.
	Description *string `pulumi:"description"`
	// Url for the discovery service to identify regional endpoints for machine learning experimentation services
	DiscoveryUrl *string `pulumi:"discoveryUrl"`
	// The encryption settings of Azure ML workspace.
	Encryption *EncryptionProperty `pulumi:"encryption"`
	// Settings for feature store type workspace.
	FeatureStoreSettings *FeatureStoreSettings `pulumi:"featureStoreSettings"`
	// The friendly name for this workspace. This name in mutable
	FriendlyName *string `pulumi:"friendlyName"`
	// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
	HbiWorkspace *bool `pulumi:"hbiWorkspace"`
	// The identity of the resource.
	Identity *ManagedServiceIdentity `pulumi:"identity"`
	// The compute name for image build
	ImageBuildCompute *string `pulumi:"imageBuildCompute"`
	// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
	KeyVault *string `pulumi:"keyVault"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Managed Network settings for a machine learning workspace.
	ManagedNetwork *ManagedNetworkSettings `pulumi:"managedNetwork"`
	// The user assigned identity resource id that represents the workspace identity.
	PrimaryUserAssignedIdentity *string `pulumi:"primaryUserAssignedIdentity"`
	// Whether requests from Public Network are allowed.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Settings for serverless compute created in the workspace
	ServerlessComputeSettings *ServerlessComputeSettings `pulumi:"serverlessComputeSettings"`
	// The service managed resource settings.
	ServiceManagedResourcesSettings *ServiceManagedResourcesSettings `pulumi:"serviceManagedResourcesSettings"`
	// The list of shared private link resources in this workspace.
	SharedPrivateLinkResources []SharedPrivateLinkResource `pulumi:"sharedPrivateLinkResources"`
	// The sku of the workspace.
	Sku *Sku `pulumi:"sku"`
	// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
	StorageAccount *string `pulumi:"storageAccount"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Enabling v1_legacy_mode may prevent you from using features provided by the v2 API.
	V1LegacyMode *bool `pulumi:"v1LegacyMode"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName *string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The flag to indicate whether to allow public access when behind VNet.
	AllowPublicAccessWhenBehindVnet pulumi.BoolPtrInput
	// ARM id of the application insights associated with this workspace.
	ApplicationInsights pulumi.StringPtrInput
	// ARM id of the container registry associated with this workspace.
	ContainerRegistry pulumi.StringPtrInput
	// The description of this workspace.
	Description pulumi.StringPtrInput
	// Url for the discovery service to identify regional endpoints for machine learning experimentation services
	DiscoveryUrl pulumi.StringPtrInput
	// The encryption settings of Azure ML workspace.
	Encryption EncryptionPropertyPtrInput
	// Settings for feature store type workspace.
	FeatureStoreSettings FeatureStoreSettingsPtrInput
	// The friendly name for this workspace. This name in mutable
	FriendlyName pulumi.StringPtrInput
	// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
	HbiWorkspace pulumi.BoolPtrInput
	// The identity of the resource.
	Identity ManagedServiceIdentityPtrInput
	// The compute name for image build
	ImageBuildCompute pulumi.StringPtrInput
	// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
	KeyVault pulumi.StringPtrInput
	// Specifies the location of the resource.
	Location pulumi.StringPtrInput
	// Managed Network settings for a machine learning workspace.
	ManagedNetwork ManagedNetworkSettingsPtrInput
	// The user assigned identity resource id that represents the workspace identity.
	PrimaryUserAssignedIdentity pulumi.StringPtrInput
	// Whether requests from Public Network are allowed.
	PublicNetworkAccess pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Settings for serverless compute created in the workspace
	ServerlessComputeSettings ServerlessComputeSettingsPtrInput
	// The service managed resource settings.
	ServiceManagedResourcesSettings ServiceManagedResourcesSettingsPtrInput
	// The list of shared private link resources in this workspace.
	SharedPrivateLinkResources SharedPrivateLinkResourceArrayInput
	// The sku of the workspace.
	Sku SkuPtrInput
	// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
	StorageAccount pulumi.StringPtrInput
	// Contains resource tags defined as key/value pairs.
	Tags pulumi.StringMapInput
	// Enabling v1_legacy_mode may prevent you from using features provided by the v2 API.
	V1LegacyMode pulumi.BoolPtrInput
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringPtrInput
}

func (WorkspaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workspaceArgs)(nil)).Elem()
}

type WorkspaceInput interface {
	pulumi.Input

	ToWorkspaceOutput() WorkspaceOutput
	ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput
}

func (*Workspace) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (i *Workspace) ToWorkspaceOutput() WorkspaceOutput {
	return i.ToWorkspaceOutputWithContext(context.Background())
}

func (i *Workspace) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkspaceOutput)
}

func (i *Workspace) ToOutput(ctx context.Context) pulumix.Output[*Workspace] {
	return pulumix.Output[*Workspace]{
		OutputState: i.ToWorkspaceOutputWithContext(ctx).OutputState,
	}
}

type WorkspaceOutput struct{ *pulumi.OutputState }

func (WorkspaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workspace)(nil)).Elem()
}

func (o WorkspaceOutput) ToWorkspaceOutput() WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToWorkspaceOutputWithContext(ctx context.Context) WorkspaceOutput {
	return o
}

func (o WorkspaceOutput) ToOutput(ctx context.Context) pulumix.Output[*Workspace] {
	return pulumix.Output[*Workspace]{
		OutputState: o.OutputState,
	}
}

// The flag to indicate whether to allow public access when behind VNet.
func (o WorkspaceOutput) AllowPublicAccessWhenBehindVnet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.AllowPublicAccessWhenBehindVnet }).(pulumi.BoolPtrOutput)
}

// ARM id of the application insights associated with this workspace.
func (o WorkspaceOutput) ApplicationInsights() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ApplicationInsights }).(pulumi.StringPtrOutput)
}

// ARM id of the container registry associated with this workspace.
func (o WorkspaceOutput) ContainerRegistry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ContainerRegistry }).(pulumi.StringPtrOutput)
}

// The description of this workspace.
func (o WorkspaceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Url for the discovery service to identify regional endpoints for machine learning experimentation services
func (o WorkspaceOutput) DiscoveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.DiscoveryUrl }).(pulumi.StringPtrOutput)
}

// The encryption settings of Azure ML workspace.
func (o WorkspaceOutput) Encryption() EncryptionPropertyResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) EncryptionPropertyResponsePtrOutput { return v.Encryption }).(EncryptionPropertyResponsePtrOutput)
}

// Settings for feature store type workspace.
func (o WorkspaceOutput) FeatureStoreSettings() FeatureStoreSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) FeatureStoreSettingsResponsePtrOutput { return v.FeatureStoreSettings }).(FeatureStoreSettingsResponsePtrOutput)
}

// The friendly name for this workspace. This name in mutable
func (o WorkspaceOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
func (o WorkspaceOutput) HbiWorkspace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.HbiWorkspace }).(pulumi.BoolPtrOutput)
}

// The identity of the resource.
func (o WorkspaceOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

// The compute name for image build
func (o WorkspaceOutput) ImageBuildCompute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ImageBuildCompute }).(pulumi.StringPtrOutput)
}

// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
func (o WorkspaceOutput) KeyVault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.KeyVault }).(pulumi.StringPtrOutput)
}

// Specifies the location of the resource.
func (o WorkspaceOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

// Managed Network settings for a machine learning workspace.
func (o WorkspaceOutput) ManagedNetwork() ManagedNetworkSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ManagedNetworkSettingsResponsePtrOutput { return v.ManagedNetwork }).(ManagedNetworkSettingsResponsePtrOutput)
}

// The URI associated with this workspace that machine learning flow must point at to set up tracking.
func (o WorkspaceOutput) MlFlowTrackingUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.MlFlowTrackingUri }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The notebook info of Azure ML workspace.
func (o WorkspaceOutput) NotebookInfo() NotebookResourceInfoResponseOutput {
	return o.ApplyT(func(v *Workspace) NotebookResourceInfoResponseOutput { return v.NotebookInfo }).(NotebookResourceInfoResponseOutput)
}

// The user assigned identity resource id that represents the workspace identity.
func (o WorkspaceOutput) PrimaryUserAssignedIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PrimaryUserAssignedIdentity }).(pulumi.StringPtrOutput)
}

// The list of private endpoint connections in the workspace.
func (o WorkspaceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Count of private connections in the workspace
func (o WorkspaceOutput) PrivateLinkCount() pulumi.IntOutput {
	return o.ApplyT(func(v *Workspace) pulumi.IntOutput { return v.PrivateLinkCount }).(pulumi.IntOutput)
}

// The current deployment state of workspace resource. The provisioningState is to indicate states for resource provisioning.
func (o WorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Whether requests from Public Network are allowed.
func (o WorkspaceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Settings for serverless compute created in the workspace
func (o WorkspaceOutput) ServerlessComputeSettings() ServerlessComputeSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ServerlessComputeSettingsResponsePtrOutput { return v.ServerlessComputeSettings }).(ServerlessComputeSettingsResponsePtrOutput)
}

// The service managed resource settings.
func (o WorkspaceOutput) ServiceManagedResourcesSettings() ServiceManagedResourcesSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ServiceManagedResourcesSettingsResponsePtrOutput {
		return v.ServiceManagedResourcesSettings
	}).(ServiceManagedResourcesSettingsResponsePtrOutput)
}

// The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace
func (o WorkspaceOutput) ServiceProvisionedResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ServiceProvisionedResourceGroup }).(pulumi.StringOutput)
}

// The list of shared private link resources in this workspace.
func (o WorkspaceOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) SharedPrivateLinkResourceResponseArrayOutput { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

// The sku of the workspace.
func (o WorkspaceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
func (o WorkspaceOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

// If the storage associated with the workspace has hierarchical namespace(HNS) enabled.
func (o WorkspaceOutput) StorageHnsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolOutput { return v.StorageHnsEnabled }).(pulumi.BoolOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o WorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Workspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Contains resource tags defined as key/value pairs.
func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The tenant id associated with this workspace.
func (o WorkspaceOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Enabling v1_legacy_mode may prevent you from using features provided by the v2 API.
func (o WorkspaceOutput) V1LegacyMode() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.V1LegacyMode }).(pulumi.BoolPtrOutput)
}

// The immutable id associated with this workspace.
func (o WorkspaceOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
