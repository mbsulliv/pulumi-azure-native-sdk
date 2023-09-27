// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the properties of the specified machine learning workspace.
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200901preview:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWorkspaceArgs struct {
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// An object that represents a machine learning workspace.
type LookupWorkspaceResult struct {
	// The flag to indicate whether to allow public access when behind VNet.
	AllowPublicAccessWhenBehindVnet *bool `pulumi:"allowPublicAccessWhenBehindVnet"`
	// ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
	ApplicationInsights *string `pulumi:"applicationInsights"`
	// ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
	ContainerRegistry *string `pulumi:"containerRegistry"`
	// The creation time of the machine learning workspace in ISO8601 format.
	CreationTime string `pulumi:"creationTime"`
	// The description of this workspace.
	Description *string `pulumi:"description"`
	// Url for the discovery service to identify regional endpoints for machine learning experimentation services
	DiscoveryUrl *string `pulumi:"discoveryUrl"`
	// The encryption settings of Azure ML workspace.
	Encryption *EncryptionPropertyResponse `pulumi:"encryption"`
	// The friendly name for this workspace. This name in mutable
	FriendlyName *string `pulumi:"friendlyName"`
	// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
	HbiWorkspace *bool `pulumi:"hbiWorkspace"`
	// Specifies the resource ID.
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The compute name for image build
	ImageBuildCompute *string `pulumi:"imageBuildCompute"`
	// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
	KeyVault *string `pulumi:"keyVault"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// The list of private endpoint connections in the workspace.
	PrivateEndpointConnections []PrivateEndpointConnectionResponse `pulumi:"privateEndpointConnections"`
	// Count of private connections in the workspace
	PrivateLinkCount int `pulumi:"privateLinkCount"`
	// The current deployment state of workspace resource. The provisioningState is to indicate states for resource provisioning.
	ProvisioningState string `pulumi:"provisioningState"`
	// The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace
	ServiceProvisionedResourceGroup string `pulumi:"serviceProvisionedResourceGroup"`
	// The list of shared private link resources in this workspace.
	SharedPrivateLinkResources []SharedPrivateLinkResourceResponse `pulumi:"sharedPrivateLinkResources"`
	// The sku of the workspace.
	Sku *SkuResponse `pulumi:"sku"`
	// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
	StorageAccount *string `pulumi:"storageAccount"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
	// The immutable id associated with this workspace.
	WorkspaceId string `pulumi:"workspaceId"`
}

// Defaults sets the appropriate defaults for LookupWorkspaceResult
func (val *LookupWorkspaceResult) Defaults() *LookupWorkspaceResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.AllowPublicAccessWhenBehindVnet == nil {
		allowPublicAccessWhenBehindVnet_ := false
		tmp.AllowPublicAccessWhenBehindVnet = &allowPublicAccessWhenBehindVnet_
	}
	if tmp.HbiWorkspace == nil {
		hbiWorkspace_ := false
		tmp.HbiWorkspace = &hbiWorkspace_
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
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
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

func (o LookupWorkspaceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceResult] {
	return pulumix.Output[LookupWorkspaceResult]{
		OutputState: o.OutputState,
	}
}

// The flag to indicate whether to allow public access when behind VNet.
func (o LookupWorkspaceResultOutput) AllowPublicAccessWhenBehindVnet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.AllowPublicAccessWhenBehindVnet }).(pulumi.BoolPtrOutput)
}

// ARM id of the application insights associated with this workspace. This cannot be changed once the workspace has been created
func (o LookupWorkspaceResultOutput) ApplicationInsights() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ApplicationInsights }).(pulumi.StringPtrOutput)
}

// ARM id of the container registry associated with this workspace. This cannot be changed once the workspace has been created
func (o LookupWorkspaceResultOutput) ContainerRegistry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ContainerRegistry }).(pulumi.StringPtrOutput)
}

// The creation time of the machine learning workspace in ISO8601 format.
func (o LookupWorkspaceResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// The description of this workspace.
func (o LookupWorkspaceResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Url for the discovery service to identify regional endpoints for machine learning experimentation services
func (o LookupWorkspaceResultOutput) DiscoveryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.DiscoveryUrl }).(pulumi.StringPtrOutput)
}

// The encryption settings of Azure ML workspace.
func (o LookupWorkspaceResultOutput) Encryption() EncryptionPropertyResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *EncryptionPropertyResponse { return v.Encryption }).(EncryptionPropertyResponsePtrOutput)
}

// The friendly name for this workspace. This name in mutable
func (o LookupWorkspaceResultOutput) FriendlyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.FriendlyName }).(pulumi.StringPtrOutput)
}

// The flag to signal HBI data in the workspace and reduce diagnostic data collected by the service
func (o LookupWorkspaceResultOutput) HbiWorkspace() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *bool { return v.HbiWorkspace }).(pulumi.BoolPtrOutput)
}

// Specifies the resource ID.
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupWorkspaceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The compute name for image build
func (o LookupWorkspaceResultOutput) ImageBuildCompute() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.ImageBuildCompute }).(pulumi.StringPtrOutput)
}

// ARM id of the key vault associated with this workspace. This cannot be changed once the workspace has been created
func (o LookupWorkspaceResultOutput) KeyVault() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.KeyVault }).(pulumi.StringPtrOutput)
}

// Specifies the location of the resource.
func (o LookupWorkspaceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
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

// The name of the managed resource group created by workspace RP in customer subscription if the workspace is CMK workspace
func (o LookupWorkspaceResultOutput) ServiceProvisionedResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ServiceProvisionedResourceGroup }).(pulumi.StringOutput)
}

// The list of shared private link resources in this workspace.
func (o LookupWorkspaceResultOutput) SharedPrivateLinkResources() SharedPrivateLinkResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []SharedPrivateLinkResourceResponse { return v.SharedPrivateLinkResources }).(SharedPrivateLinkResourceResponseArrayOutput)
}

// The sku of the workspace.
func (o LookupWorkspaceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// ARM id of the storage account associated with this workspace. This cannot be changed once the workspace has been created
func (o LookupWorkspaceResultOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

// Contains resource tags defined as key/value pairs.
func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The immutable id associated with this workspace.
func (o LookupWorkspaceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
