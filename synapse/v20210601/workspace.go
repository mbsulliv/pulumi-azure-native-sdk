// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210601

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// A workspace
type Workspace struct {
	pulumi.CustomResourceState

	// The ADLA resource ID.
	AdlaResourceId pulumi.StringOutput `pulumi:"adlaResourceId"`
	// Connectivity endpoints
	ConnectivityEndpoints pulumi.StringMapOutput `pulumi:"connectivityEndpoints"`
	// Initial workspace AAD admin properties for a CSP subscription
	CspWorkspaceAdminProperties CspWorkspaceAdminPropertiesResponsePtrOutput `pulumi:"cspWorkspaceAdminProperties"`
	// Workspace default data lake storage account details
	DefaultDataLakeStorage DataLakeStorageAccountDetailsResponsePtrOutput `pulumi:"defaultDataLakeStorage"`
	// The encryption details of the workspace
	Encryption EncryptionDetailsResponsePtrOutput `pulumi:"encryption"`
	// Workspace level configs and feature flags
	ExtraProperties pulumi.AnyOutput `pulumi:"extraProperties"`
	// Identity of the workspace
	Identity ManagedIdentityResponsePtrOutput `pulumi:"identity"`
	// The geo-location where the resource lives
	Location pulumi.StringOutput `pulumi:"location"`
	// Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
	ManagedResourceGroupName pulumi.StringPtrOutput `pulumi:"managedResourceGroupName"`
	// Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
	ManagedVirtualNetwork pulumi.StringPtrOutput `pulumi:"managedVirtualNetwork"`
	// Managed Virtual Network Settings
	ManagedVirtualNetworkSettings ManagedVirtualNetworkSettingsResponsePtrOutput `pulumi:"managedVirtualNetworkSettings"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// Private endpoint connections to the workspace
	PrivateEndpointConnections PrivateEndpointConnectionResponseArrayOutput `pulumi:"privateEndpointConnections"`
	// Resource provisioning state
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	// Enable or Disable public network access to workspace
	PublicNetworkAccess pulumi.StringPtrOutput `pulumi:"publicNetworkAccess"`
	// Purview Configuration
	PurviewConfiguration PurviewConfigurationResponsePtrOutput `pulumi:"purviewConfiguration"`
	// Workspace settings
	Settings pulumi.MapOutput `pulumi:"settings"`
	// Login for workspace SQL active directory administrator
	SqlAdministratorLogin pulumi.StringPtrOutput `pulumi:"sqlAdministratorLogin"`
	// SQL administrator login password
	SqlAdministratorLoginPassword pulumi.StringPtrOutput `pulumi:"sqlAdministratorLoginPassword"`
	// Resource tags.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Is trustedServiceBypassEnabled for the workspace
	TrustedServiceBypassEnabled pulumi.BoolPtrOutput `pulumi:"trustedServiceBypassEnabled"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type pulumi.StringOutput `pulumi:"type"`
	// Virtual Network profile
	VirtualNetworkProfile VirtualNetworkProfileResponsePtrOutput `pulumi:"virtualNetworkProfile"`
	// Git integration settings
	WorkspaceRepositoryConfiguration WorkspaceRepositoryConfigurationResponsePtrOutput `pulumi:"workspaceRepositoryConfiguration"`
	// The workspace unique identifier
	WorkspaceUID pulumi.StringOutput `pulumi:"workspaceUID"`
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
	if args.PublicNetworkAccess == nil {
		args.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	if args.TrustedServiceBypassEnabled == nil {
		args.TrustedServiceBypassEnabled = pulumi.BoolPtr(false)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:synapse:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20190601preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20201201:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210301:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210401preview:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210501:Workspace"),
		},
		{
			Type: pulumi.String("azure-native:synapse/v20210601preview:Workspace"),
		},
	})
	opts = append(opts, aliases)
	opts = utilities.PkgResourceDefaultOpts(opts)
	var resource Workspace
	err := ctx.RegisterResource("azure-native:synapse/v20210601:Workspace", name, args, &resource, opts...)
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
	err := ctx.ReadResource("azure-native:synapse/v20210601:Workspace", name, id, state, &resource, opts...)
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
	// Enable or Disable AzureADOnlyAuthentication on All Workspace subresource
	AzureADOnlyAuthentication *bool `pulumi:"azureADOnlyAuthentication"`
	// Initial workspace AAD admin properties for a CSP subscription
	CspWorkspaceAdminProperties *CspWorkspaceAdminProperties `pulumi:"cspWorkspaceAdminProperties"`
	// Workspace default data lake storage account details
	DefaultDataLakeStorage *DataLakeStorageAccountDetails `pulumi:"defaultDataLakeStorage"`
	// The encryption details of the workspace
	Encryption *EncryptionDetails `pulumi:"encryption"`
	// Identity of the workspace
	Identity *ManagedIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location *string `pulumi:"location"`
	// Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
	ManagedVirtualNetwork *string `pulumi:"managedVirtualNetwork"`
	// Managed Virtual Network Settings
	ManagedVirtualNetworkSettings *ManagedVirtualNetworkSettings `pulumi:"managedVirtualNetworkSettings"`
	// Private endpoint connections to the workspace
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	PrivateEndpointConnections []PrivateEndpointConnectionType `pulumi:"privateEndpointConnections"`
	// Enable or Disable public network access to workspace
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Purview Configuration
	PurviewConfiguration *PurviewConfiguration `pulumi:"purviewConfiguration"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Login for workspace SQL active directory administrator
	SqlAdministratorLogin *string `pulumi:"sqlAdministratorLogin"`
	// SQL administrator login password
	SqlAdministratorLoginPassword *string `pulumi:"sqlAdministratorLoginPassword"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Is trustedServiceBypassEnabled for the workspace
	TrustedServiceBypassEnabled *bool `pulumi:"trustedServiceBypassEnabled"`
	// Virtual Network profile
	VirtualNetworkProfile *VirtualNetworkProfile `pulumi:"virtualNetworkProfile"`
	// The name of the workspace.
	WorkspaceName *string `pulumi:"workspaceName"`
	// Git integration settings
	WorkspaceRepositoryConfiguration *WorkspaceRepositoryConfiguration `pulumi:"workspaceRepositoryConfiguration"`
}

// The set of arguments for constructing a Workspace resource.
type WorkspaceArgs struct {
	// Enable or Disable AzureADOnlyAuthentication on All Workspace subresource
	AzureADOnlyAuthentication pulumi.BoolPtrInput
	// Initial workspace AAD admin properties for a CSP subscription
	CspWorkspaceAdminProperties CspWorkspaceAdminPropertiesPtrInput
	// Workspace default data lake storage account details
	DefaultDataLakeStorage DataLakeStorageAccountDetailsPtrInput
	// The encryption details of the workspace
	Encryption EncryptionDetailsPtrInput
	// Identity of the workspace
	Identity ManagedIdentityPtrInput
	// The geo-location where the resource lives
	Location pulumi.StringPtrInput
	// Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
	ManagedResourceGroupName pulumi.StringPtrInput
	// Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
	ManagedVirtualNetwork pulumi.StringPtrInput
	// Managed Virtual Network Settings
	ManagedVirtualNetworkSettings ManagedVirtualNetworkSettingsPtrInput
	// Private endpoint connections to the workspace
	// These are also available as standalone resources. Do not mix inline and standalone resource as they will conflict with each other, leading to resources deletion.
	PrivateEndpointConnections PrivateEndpointConnectionTypeArrayInput
	// Enable or Disable public network access to workspace
	PublicNetworkAccess pulumi.StringPtrInput
	// Purview Configuration
	PurviewConfiguration PurviewConfigurationPtrInput
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput
	// Login for workspace SQL active directory administrator
	SqlAdministratorLogin pulumi.StringPtrInput
	// SQL administrator login password
	SqlAdministratorLoginPassword pulumi.StringPtrInput
	// Resource tags.
	Tags pulumi.StringMapInput
	// Is trustedServiceBypassEnabled for the workspace
	TrustedServiceBypassEnabled pulumi.BoolPtrInput
	// Virtual Network profile
	VirtualNetworkProfile VirtualNetworkProfilePtrInput
	// The name of the workspace.
	WorkspaceName pulumi.StringPtrInput
	// Git integration settings
	WorkspaceRepositoryConfiguration WorkspaceRepositoryConfigurationPtrInput
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

// The ADLA resource ID.
func (o WorkspaceOutput) AdlaResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.AdlaResourceId }).(pulumi.StringOutput)
}

// Connectivity endpoints
func (o WorkspaceOutput) ConnectivityEndpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.ConnectivityEndpoints }).(pulumi.StringMapOutput)
}

// Initial workspace AAD admin properties for a CSP subscription
func (o WorkspaceOutput) CspWorkspaceAdminProperties() CspWorkspaceAdminPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) CspWorkspaceAdminPropertiesResponsePtrOutput { return v.CspWorkspaceAdminProperties }).(CspWorkspaceAdminPropertiesResponsePtrOutput)
}

// Workspace default data lake storage account details
func (o WorkspaceOutput) DefaultDataLakeStorage() DataLakeStorageAccountDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) DataLakeStorageAccountDetailsResponsePtrOutput { return v.DefaultDataLakeStorage }).(DataLakeStorageAccountDetailsResponsePtrOutput)
}

// The encryption details of the workspace
func (o WorkspaceOutput) Encryption() EncryptionDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) EncryptionDetailsResponsePtrOutput { return v.Encryption }).(EncryptionDetailsResponsePtrOutput)
}

// Workspace level configs and feature flags
func (o WorkspaceOutput) ExtraProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v *Workspace) pulumi.AnyOutput { return v.ExtraProperties }).(pulumi.AnyOutput)
}

// Identity of the workspace
func (o WorkspaceOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ManagedIdentityResponsePtrOutput { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o WorkspaceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
func (o WorkspaceOutput) ManagedResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ManagedResourceGroupName }).(pulumi.StringPtrOutput)
}

// Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
func (o WorkspaceOutput) ManagedVirtualNetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.ManagedVirtualNetwork }).(pulumi.StringPtrOutput)
}

// Managed Virtual Network Settings
func (o WorkspaceOutput) ManagedVirtualNetworkSettings() ManagedVirtualNetworkSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) ManagedVirtualNetworkSettingsResponsePtrOutput {
		return v.ManagedVirtualNetworkSettings
	}).(ManagedVirtualNetworkSettingsResponsePtrOutput)
}

// The name of the resource
func (o WorkspaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint connections to the workspace
func (o WorkspaceOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v *Workspace) PrivateEndpointConnectionResponseArrayOutput { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Resource provisioning state
func (o WorkspaceOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Enable or Disable public network access to workspace
func (o WorkspaceOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Purview Configuration
func (o WorkspaceOutput) PurviewConfiguration() PurviewConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) PurviewConfigurationResponsePtrOutput { return v.PurviewConfiguration }).(PurviewConfigurationResponsePtrOutput)
}

// Workspace settings
func (o WorkspaceOutput) Settings() pulumi.MapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.MapOutput { return v.Settings }).(pulumi.MapOutput)
}

// Login for workspace SQL active directory administrator
func (o WorkspaceOutput) SqlAdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.SqlAdministratorLogin }).(pulumi.StringPtrOutput)
}

// SQL administrator login password
func (o WorkspaceOutput) SqlAdministratorLoginPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringPtrOutput { return v.SqlAdministratorLoginPassword }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o WorkspaceOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

// Is trustedServiceBypassEnabled for the workspace
func (o WorkspaceOutput) TrustedServiceBypassEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Workspace) pulumi.BoolPtrOutput { return v.TrustedServiceBypassEnabled }).(pulumi.BoolPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o WorkspaceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Virtual Network profile
func (o WorkspaceOutput) VirtualNetworkProfile() VirtualNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) VirtualNetworkProfileResponsePtrOutput { return v.VirtualNetworkProfile }).(VirtualNetworkProfileResponsePtrOutput)
}

// Git integration settings
func (o WorkspaceOutput) WorkspaceRepositoryConfiguration() WorkspaceRepositoryConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Workspace) WorkspaceRepositoryConfigurationResponsePtrOutput {
		return v.WorkspaceRepositoryConfiguration
	}).(WorkspaceRepositoryConfigurationResponsePtrOutput)
}

// The workspace unique identifier
func (o WorkspaceOutput) WorkspaceUID() pulumi.StringOutput {
	return o.ApplyT(func(v *Workspace) pulumi.StringOutput { return v.WorkspaceUID }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkspaceOutput{})
}
