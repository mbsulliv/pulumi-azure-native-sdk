// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Information about workspace.
// Azure REST API version: 2023-02-01. Prior API version in Azure Native 1.x: 2018-04-01
type Workspace struct {
	pulumi.CustomResourceState

	// The workspace provider authorizations.
	Authorizations WorkspaceProviderAuthorizationResponseArrayOutput `pulumi:"authorizations"`
	// Indicates the Object ID, PUID and Application ID of entity that created the workspace.
	CreatedBy CreatedByResponsePtrOutput `pulumi:"createdBy"`
	// Specifies the date and time when the workspace is created.
	CreatedDateTime pulumi.StringOutput `pulumi:"createdDateTime"`
	// The resource Id of the managed disk encryption set.
	DiskEncryptionSetId pulumi.StringOutput `pulumi:"diskEncryptionSetId"`
	// Encryption properties for databricks workspace
	Encryption WorkspacePropertiesResponseEncryptionPtrOutput `pulumi:"encryption"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// The details of Managed Identity of Disk Encryption Set used for Managed Disk Encryption
	ManagedDiskIdentity ManagedIdentityConfigurationResponsePtrOutput `pulumi:"managedDiskIdentity"`
	// The managed resource group Id.
	ManagedResourceGroupId pulumi.StringOutput `pulumi:"managedResourceGroupId"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// The workspace's custom parameters.
	Parameters WorkspaceCustomParametersResponsePtrOutput `pulumi:"parameters"`
	// Private endpoint connections created on the workspace
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// The workspace provisioning state.
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// The network access type for accessing workspace. Set value to disabled to access workspace only via private link.
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Gets or sets a value indicating whether data plane (clusters) to control plane communication happen over private endpoint. Supported values are 'AllRules' and 'NoAzureDatabricksRules'. 'NoAzureServiceRules' value is for internal use only.
	RequiredNsgRules pulumi.StringPtrOutput `pulumi:"requiredNsgRules"`
	// The SKU of the resource.
	Sku SkuResponsePtrOutput `pulumi:"sku"`
	// The details of Managed Identity of Storage Account
	StorageAccountIdentity ManagedIdentityConfigurationResponsePtrOutput `pulumi:"storageAccountIdentity"`
	// The system metadata relating to this resource
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type pulumi.StringOutput `pulumi:"type"`
	// The blob URI where the UI definition file is located.
	UiDefinitionUri pulumi.StringPtrOutput `pulumi:"uiDefinitionUri"`
	// Indicates the Object ID, PUID and Application ID of entity that last updated the workspace.
	UpdatedBy CreatedByResponsePtrOutput `pulumi:"updatedBy"`
	// The unique identifier of the databricks workspace in databricks control plane.
	WorkspaceId pulumi.StringOutput `pulumi:"workspaceId"`
	// The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
	WorkspaceUrl pulumi.StringOutput `pulumi:"workspaceUrl"`
}

// NewWorkspace registers a new resource with the given unique name, arguments, and options.
func NewWorkspace(ctx *pulumi.Context,
	name string, args *WorkspaceArgs, opts ...pulumi.ResourceOption) (*Workspace, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ManagedResourceGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagedResourceGroupId'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Parameters != nil {
		args.Parameters = args.Parameters.ToWorkspaceCustomParametersPtrOutput().ApplyT(func(v *WorkspaceCustomParameters) *WorkspaceCustomParameters { return v.Defaults() }).(WorkspaceCustomParametersPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:databricks/v20180401:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:databricks/v20210401preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:databricks/v20220401preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:databricks/v20230201:Workspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:databricks:Workspace", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:databricks:Workspace", name, id, state, &resource, opts...)
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
	// The workspace provider authorizations.
	Authorizations []WorkspaceProviderAuthorization `pulumi:"authorizations"`
	// Encryption properties for databricks workspace
	Encryption *WorkspacePropertiesEncryption `pulumi:"encryption"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// The managed resource group Id.
	ManagedResourceGroupId string `pulumi:"managedResourceGroupId"`
	// The workspace's custom parameters.
	Parameters *WorkspaceCustomParameters `pulumi:"parameters"`
	// The network access type for accessing workspace. Set value to disabled to access workspace only via private link.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Gets or sets a value indicating whether data plane (clusters) to control plane communication happen over private endpoint. Supported values are 'AllRules' and 'NoAzureDatabricksRules'. 'NoAzureServiceRules' value is for internal use only.
	RequiredNsgRules *string `pulumi:"requiredNsgRules"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The SKU of the resource.
	Sku *Sku `pulumi:"sku"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The blob URI where the UI definition file is located.
	UiDefinitionUri *string `pulumi:"uiDefinitionUri"`
	// The name of the workspace.
	WorkspaceName *string `pulumi:"workspaceName"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// The workspace provider authorizations.
	Authorizations WorkspaceProviderAuthorizationArrayInput
	// Encryption properties for databricks workspace
	Encryption WorkspacePropertiesEncryptionPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// The managed resource group Id.
	ManagedResourceGroupId pulumi.StringInput
	// The workspace's custom parameters.
	Parameters WorkspaceCustomParametersPtrInput
	// The network access type for accessing workspace. Set value to disabled to access workspace only via private link.
	PublicNetworkAccess pulumi.StringPtrInput
	// Gets or sets a value indicating whether data plane (clusters) to control plane communication happen over private endpoint. Supported values are 'AllRules' and 'NoAzureDatabricksRules'. 'NoAzureServiceRules' value is for internal use only.
	RequiredNsgRules pulumi.StringPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// The SKU of the resource.
	Sku SkuPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// The blob URI where the UI definition file is located.
	UiDefinitionUri pulumi.StringPtrInput
	// The name of the workspace.
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

// The workspace provider authorizations.
func (o WorkspaceOutput) Authorizations() WorkspaceProviderAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceProviderAuthorizationResponseArrayOutput { return v.Authorizations }).(WorkspaceProviderAuthorizationResponseArrayOutput)
}

// Indicates the Object ID, PUID and Application ID of entity that created the workspace.
func (o WorkspaceOutput) CreatedBy() CreatedByResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) CreatedByResponsePtrOutput { return v.CreatedBy }).(CreatedByResponsePtrOutput)
}

// Specifies the date and time when the workspace is created.
func (o WorkspaceOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.CreatedDateTime }).(pulumi.StringOutput)
}

// The resource Id of the managed disk encryption set.
func (o WorkspaceOutput) DiskEncryptionSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.DiskEncryptionSetId }).(pulumi.StringOutput)
}

// Encryption properties for databricks workspace
func (o WorkspaceOutput) Encryption() WorkspacePropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspacePropertiesResponseEncryptionPtrOutput { return v.Encryption }).(WorkspacePropertiesResponseEncryptionPtrOutput)
}

// The geo-location where the resource lives
func (o WorkspaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The details of Managed Identity of Disk Encryption Set used for Managed Disk Encryption
func (o WorkspaceOutput) ManagedDiskIdentity() ManagedIdentityConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ManagedIdentityConfigurationResponsePtrOutput { return v.ManagedDiskIdentity }).(ManagedIdentityConfigurationResponsePtrOutput)
}

// The managed resource group Id.
func (o WorkspaceOutput) ManagedResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ManagedResourceGroupId }).(pulumi.StringOutput)
}

// The name of the resource
func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The workspace's custom parameters.
func (o WorkspaceOutput) Parameters() WorkspaceCustomParametersResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceCustomParametersResponsePtrOutput { return v.Parameters }).(WorkspaceCustomParametersResponsePtrOutput)
}

// Private endpoint connections created on the workspace
func (o WorkspaceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// The workspace provisioning state.
func (o WorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The network access type for accessing workspace. Set value to disabled to access workspace only via private link.
func (o WorkspaceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Gets or sets a value indicating whether data plane (clusters) to control plane communication happen over private endpoint. Supported values are 'AllRules' and 'NoAzureDatabricksRules'. 'NoAzureServiceRules' value is for internal use only.
func (o WorkspaceOutput) RequiredNsgRules() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.RequiredNsgRules }).(pulumi.StringPtrOutput)
}

// The SKU of the resource.
func (o WorkspaceOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

// The details of Managed Identity of Storage Account
func (o WorkspaceOutput) StorageAccountIdentity() ManagedIdentityConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ManagedIdentityConfigurationResponsePtrOutput { return v.StorageAccountIdentity }).(ManagedIdentityConfigurationResponsePtrOutput)
}

// The system metadata relating to this resource
func (o WorkspaceOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Workspace) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The blob URI where the UI definition file is located.
func (o WorkspaceOutput) UiDefinitionUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.UiDefinitionUri }).(pulumi.StringPtrOutput)
}

// Indicates the Object ID, PUID and Application ID of entity that last updated the workspace.
func (o WorkspaceOutput) UpdatedBy() CreatedByResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) CreatedByResponsePtrOutput { return v.UpdatedBy }).(CreatedByResponsePtrOutput)
}

// The unique identifier of the databricks workspace in databricks control plane.
func (o WorkspaceOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.WorkspaceId }).(pulumi.StringOutput)
}

// The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
func (o WorkspaceOutput) WorkspaceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.WorkspaceUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
