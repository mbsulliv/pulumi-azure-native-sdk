// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package databricks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the workspace.
// Azure REST API version: 2023-02-01.
//
// Other available API versions: 2023-09-15-preview.
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:databricks:getWorkspace", args, &rv, opts...)
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

// Information about workspace.
type LookupWorkspaceResult struct {
	// The workspace provider authorizations.
	Authorizations []WorkspaceProviderAuthorizationResponse `pulumi:"authorizations"`
	// Indicates the Object ID, PUID and Application ID of entity that created the workspace.
	CreatedBy *CreatedByResponse `pulumi:"createdBy"`
	// Specifies the date and time when the workspace is created.
	CreatedDateTime string `pulumi:"createdDateTime"`
	// The resource Id of the managed disk encryption set.
	DiskEncryptionSetId string `pulumi:"diskEncryptionSetId"`
	// Encryption properties for databricks workspace
	Encryption *WorkspacePropertiesResponseEncryption `pulumi:"encryption"`
	// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The details of Managed Identity of Disk Encryption Set used for Managed Disk Encryption
	ManagedDiskIdentity *ManagedIdentityConfigurationResponse `pulumi:"managedDiskIdentity"`
	// The managed resource group Id.
	ManagedResourceGroupId string `pulumi:"managedResourceGroupId"`
	// The name of the resource
	Name string `pulumi:"name"`
	// The workspace's custom parameters.
	Parameters *WorkspaceCustomParametersResponse `pulumi:"parameters"`
	// Private endpoint connections created on the workspace
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// The workspace provisioning state.
	ProvisioningState string `pulumi:"provisioningState"`
	// The network access type for accessing workspace. Set value to disabled to access workspace only via private link.
	PublicNetworkAccess *string `pulumi:"publicNetworkAccess"`
	// Gets or sets a value indicating whether data plane (clusters) to control plane communication happen over private endpoint. Supported values are 'AllRules' and 'NoAzureDatabricksRules'. 'NoAzureServiceRules' value is for internal use only.
	RequiredNsgRules *string `pulumi:"requiredNsgRules"`
	// The SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// The details of Managed Identity of Storage Account
	StorageAccountIdentity *ManagedIdentityConfigurationResponse `pulumi:"storageAccountIdentity"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
	Type string `pulumi:"type"`
	// The blob URI where the UI definition file is located.
	UiDefinitionUri *string `pulumi:"uiDefinitionUri"`
	// Indicates the Object ID, PUID and Application ID of entity that last updated the workspace.
	UpdatedBy *CreatedByResponse `pulumi:"updatedBy"`
	// The unique identifier of the databricks workspace in databricks control plane.
	WorkspaceId string `pulumi:"workspaceId"`
	// The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
	WorkspaceUrl string `pulumi:"workspaceUrl"`
}

// Defaults sets the appropriate defaults for LookupWorkspaceResult
func (val *LookupWorkspaceResult) Defaults() *LookupWorkspaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Parameters = tmp.Parameters.Defaults()

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

// Information about workspace.
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

// The workspace provider authorizations.
func (o LookupWorkspaceResultOutput) Authorizations() WorkspaceProviderAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []WorkspaceProviderAuthorizationResponse { return v.Authorizations }).(WorkspaceProviderAuthorizationResponseArrayOutput)
}

// Indicates the Object ID, PUID and Application ID of entity that created the workspace.
func (o LookupWorkspaceResultOutput) CreatedBy() CreatedByResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *CreatedByResponse { return v.CreatedBy }).(CreatedByResponsePtrOutput)
}

// Specifies the date and time when the workspace is created.
func (o LookupWorkspaceResultOutput) CreatedDateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CreatedDateTime }).(pulumi.StringOutput)
}

// The resource Id of the managed disk encryption set.
func (o LookupWorkspaceResultOutput) DiskEncryptionSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.DiskEncryptionSetId }).(pulumi.StringOutput)
}

// Encryption properties for databricks workspace
func (o LookupWorkspaceResultOutput) Encryption() WorkspacePropertiesResponseEncryptionPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *WorkspacePropertiesResponseEncryption { return v.Encryption }).(WorkspacePropertiesResponseEncryptionPtrOutput)
}

// Fully qualified resource Id for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The geo-location where the resource lives
func (o LookupWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The details of Managed Identity of Disk Encryption Set used for Managed Disk Encryption
func (o LookupWorkspaceResultOutput) ManagedDiskIdentity() ManagedIdentityConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ManagedIdentityConfigurationResponse { return v.ManagedDiskIdentity }).(ManagedIdentityConfigurationResponsePtrOutput)
}

// The managed resource group Id.
func (o LookupWorkspaceResultOutput) ManagedResourceGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ManagedResourceGroupId }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The workspace's custom parameters.
func (o LookupWorkspaceResultOutput) Parameters() WorkspaceCustomParametersResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *WorkspaceCustomParametersResponse { return v.Parameters }).(WorkspaceCustomParametersResponsePtrOutput)
}

// Private endpoint connections created on the workspace
func (o LookupWorkspaceResultOutput) PrivateEndpointConnections() PrivateEndpointConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []PrivateEndpointConnectionResponse { return v.PrivateEndpointConnections }).(PrivateEndpointConnectionResponseArrayOutput)
}

// The workspace provisioning state.
func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// The network access type for accessing workspace. Set value to disabled to access workspace only via private link.
func (o LookupWorkspaceResultOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

// Gets or sets a value indicating whether data plane (clusters) to control plane communication happen over private endpoint. Supported values are 'AllRules' and 'NoAzureDatabricksRules'. 'NoAzureServiceRules' value is for internal use only.
func (o LookupWorkspaceResultOutput) RequiredNsgRules() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.RequiredNsgRules }).(pulumi.StringPtrOutput)
}

// The SKU of the resource.
func (o LookupWorkspaceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// The details of Managed Identity of Storage Account
func (o LookupWorkspaceResultOutput) StorageAccountIdentity() ManagedIdentityConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *ManagedIdentityConfigurationResponse { return v.StorageAccountIdentity }).(ManagedIdentityConfigurationResponsePtrOutput)
}

// The system metadata relating to this resource
func (o LookupWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. Ex- Microsoft.Compute/virtualMachines or Microsoft.Storage/storageAccounts.
func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The blob URI where the UI definition file is located.
func (o LookupWorkspaceResultOutput) UiDefinitionUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.UiDefinitionUri }).(pulumi.StringPtrOutput)
}

// Indicates the Object ID, PUID and Application ID of entity that last updated the workspace.
func (o LookupWorkspaceResultOutput) UpdatedBy() CreatedByResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *CreatedByResponse { return v.UpdatedBy }).(CreatedByResponsePtrOutput)
}

// The unique identifier of the databricks workspace in databricks control plane.
func (o LookupWorkspaceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

// The workspace URL which is of the format 'adb-{workspaceId}.{random}.azuredatabricks.net'
func (o LookupWorkspaceResultOutput) WorkspaceUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceUrl }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
