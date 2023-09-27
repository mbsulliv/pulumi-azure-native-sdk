// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package synapse

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a workspace
// Azure REST API version: 2021-06-01.
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:synapse:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWorkspaceArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// A workspace
type LookupWorkspaceResult struct {
	// The ADLA resource ID.
	AdlaResourceId string `pulumi:"adlaResourceId"`
	// Connectivity endpoints
	ConnectivityEndpoints map[string]string `pulumi:"connectivityEndpoints"`
	// Initial workspace AAD admin properties for a CSP subscription
	CspWorkspaceAdminProperties *CspWorkspaceAdminPropertiesResponse `pulumi:"cspWorkspaceAdminProperties"`
	// Workspace default data lake storage account details
	DefaultDataLakeStorage *DataLakeStorageAccountDetailsResponse `pulumi:"defaultDataLakeStorage"`
	// The encryption details of the workspace
	Encryption *EncryptionDetailsResponse `pulumi:"encryption"`
	// Workspace level configs and feature flags
	ExtraProperties interface{} `pulumi:"extraProperties"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Identity of the workspace
	Identity *ManagedIdentityResponse `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
	ManagedResourceGroupName *string `pulumi:"managedResourceGroupName"`
	// Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
	ManagedVirtualNetwork *string `pulumi:"managedVirtualNetwork"`
	// Managed Virtual Network Settings
	ManagedVirtualNetworkSettings *ManagedVirtualNetworkSettingsResponse `pulumi:"managedVirtualNetworkSettings"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Private endpoint connections to the workspace
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Resource provisioning state
	ProvisioningState string `pulumi:"provisioningState"`
	// Enable or Disable public network access to workspace
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Purview Configuration
	PurviewConfiguration *PurviewConfigurationResponse `pulumi:"purviewConfiguration"`
	// Workspace settings
	Settings map[string]interface{} `pulumi:"settings"`
	// Login for workspace SQL active directory administrator
	SqlAdministratorLogin *string `pulumi:"sqlAdministratorLogin"`
	// SQL administrator login password
	SqlAdministratorLoginPassword *string `pulumi:"sqlAdministratorLoginPassword"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// Is trustedServiceBypassEnabled for the workspace
	TrustedServiceBypassEnabled *bool `pulumi:"trustedServiceBypassEnabled"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Virtual Network profile
	VirtualNetworkProfile *VirtualNetworkProfileResponse `pulumi:"virtualNetworkProfile"`
	// Git integration settings
	WorkspaceRepositoryConfiguration *WorkspaceRepositoryConfigurationResponse `pulumi:"workspaceRepositoryConfiguration"`
	// The workspace unique identifier
	WorkspaceUID string `pulumi:"workspaceUID"`
}

// Defaults sets the appropriate defaults for LookupWorkspaceResult
func (val *LookupWorkspaceResult) Defaults() *LookupWorkspaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.PublicNetworkAccess == nil {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	if tmp.TrustedServiceBypassEnabled == nil {
		trustedServiceBypassEnabled_ := false
		tmp.TrustedServiceBypassEnabled = &trustedServiceBypassEnabled_
	}
	return &tmp
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
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}

// A workspace
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

func (o LookupWorkspaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceResult] {
	return pulumix.Output[LookupWorkspaceResult]{
		OutputState: o.OutputState,
	}
}

// The ADLA resource ID.
func (o LookupWorkspaceResultOutput) AdlaResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.AdlaResourceId }).(pulumi.StringOutput)
}

// Connectivity endpoints
func (o LookupWorkspaceResultOutput) ConnectivityEndpoints() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.ConnectivityEndpoints }).(pulumi.StringMapOutput)
}

// Initial workspace AAD admin properties for a CSP subscription
func (o LookupWorkspaceResultOutput) CspWorkspaceAdminProperties() CspWorkspaceAdminPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *CspWorkspaceAdminPropertiesResponse {
		return v.CspWorkspaceAdminProperties
	}).(CspWorkspaceAdminPropertiesResponsePtrOutput)
}

// Workspace default data lake storage account details
func (o LookupWorkspaceResultOutput) DefaultDataLakeStorage() DataLakeStorageAccountDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *DataLakeStorageAccountDetailsResponse { return v.DefaultDataLakeStorage }).(DataLakeStorageAccountDetailsResponsePtrOutput)
}

// The encryption details of the workspace
func (o LookupWorkspaceResultOutput) Encryption() EncryptionDetailsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *EncryptionDetailsResponse { return v.Encryption }).(EncryptionDetailsResponsePtrOutput)
}

// Workspace level configs and feature flags
func (o LookupWorkspaceResultOutput) ExtraProperties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) interface{} { return v.ExtraProperties }).(pulumi.AnyOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity of the workspace
func (o LookupWorkspaceResultOutput) Identity() ManagedIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ManagedIdentityResponse { return v.Identity }).(ManagedIdentityResponsePtrOutput)
}

// The geo-location where the resource lives
func (o LookupWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Workspace managed resource group. The resource group name uniquely identifies the resource group within the user subscriptionId. The resource group name must be no longer than 90 characters long, and must be alphanumeric characters (Char.IsLetterOrDigit()) and '-', '_', '(', ')' and'.'. Note that the name cannot end with '.'
func (o LookupWorkspaceResultOutput) ManagedResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ManagedResourceGroupName }).(pulumi.StringPtrOutput)
}

// Setting this to 'default' will ensure that all compute for this workspace is in a virtual network managed on behalf of the user.
func (o LookupWorkspaceResultOutput) ManagedVirtualNetwork() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ManagedVirtualNetwork }).(pulumi.StringPtrOutput)
}

// Managed Virtual Network Settings
func (o LookupWorkspaceResultOutput) ManagedVirtualNetworkSettings() ManagedVirtualNetworkSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ManagedVirtualNetworkSettingsResponse {
		return v.ManagedVirtualNetworkSettings
	}).(ManagedVirtualNetworkSettingsResponsePtrOutput)
}

// The name of the resource
func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Private endpoint connections to the workspace
func (o LookupWorkspaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// Resource provisioning state
func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// Enable or Disable public network access to workspace
func (o LookupWorkspaceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Purview Configuration
func (o LookupWorkspaceResultOutput) PurviewConfiguration() PurviewConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *PurviewConfigurationResponse { return v.PurviewConfiguration }).(PurviewConfigurationResponsePtrOutput)
}

// Workspace settings
func (o LookupWorkspaceResultOutput) Settings() pulumi.MapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]interface{} { return v.Settings }).(pulumi.MapOutput)
}

// Login for workspace SQL active directory administrator
func (o LookupWorkspaceResultOutput) SqlAdministratorLogin() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.SqlAdministratorLogin }).(pulumi.StringPtrOutput)
}

// SQL administrator login password
func (o LookupWorkspaceResultOutput) SqlAdministratorLoginPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.SqlAdministratorLoginPassword }).(pulumi.StringPtrOutput)
}

// Resource tags.
func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Is trustedServiceBypassEnabled for the workspace
func (o LookupWorkspaceResultOutput) TrustedServiceBypassEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.TrustedServiceBypassEnabled }).(pulumi.BoolPtrOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

// Virtual Network profile
func (o LookupWorkspaceResultOutput) VirtualNetworkProfile() VirtualNetworkProfileResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *VirtualNetworkProfileResponse { return v.VirtualNetworkProfile }).(VirtualNetworkProfileResponsePtrOutput)
}

// Git integration settings
func (o LookupWorkspaceResultOutput) WorkspaceRepositoryConfiguration() WorkspaceRepositoryConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *WorkspaceRepositoryConfigurationResponse {
		return v.WorkspaceRepositoryConfiguration
	}).(WorkspaceRepositoryConfigurationResponsePtrOutput)
}

// The workspace unique identifier
func (o LookupWorkspaceResultOutput) WorkspaceUID() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceUID }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
