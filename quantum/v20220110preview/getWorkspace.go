// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220110preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the Workspace resource associated with the given name.
func LookupWorkspace(ctx *pulumi.Context, args *LookupWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceResult
	err := ctx.Invoke("azure-native:quantum/v20220110preview:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the quantum workspace resource.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The resource proxy definition object for quantum workspace.
type LookupWorkspaceResult struct {
	// The URI of the workspace endpoint.
	EndpointUri string `pulumi:"endpointUri"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// Managed Identity information.
	Identity *QuantumWorkspaceResponseIdentity `pulumi:"identity"`
	// The geo-location where the resource lives
	Location string `pulumi:"location"`
	// The name of the resource
	Name string `pulumi:"name"`
	// List of Providers selected for this Workspace
	Providers []ProviderResponse `pulumi:"providers"`
	// Provisioning status field
	ProvisioningState string `pulumi:"provisioningState"`
	// ARM Resource Id of the storage account associated with this workspace.
	StorageAccount *string `pulumi:"storageAccount"`
	// System metadata
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Resource tags.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
	// Whether the current workspace is ready to accept Jobs.
	Usable string `pulumi:"usable"`
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
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the quantum workspace resource.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceArgs)(nil)).Elem()
}

// The resource proxy definition object for quantum workspace.
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

// The URI of the workspace endpoint.
func (o LookupWorkspaceResultOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.EndpointUri }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// Managed Identity information.
func (o LookupWorkspaceResultOutput) Identity() QuantumWorkspaceResponseIdentityPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *QuantumWorkspaceResponseIdentity { return v.Identity }).(QuantumWorkspaceResponseIdentityPtrOutput)
}

// The geo-location where the resource lives
func (o LookupWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// List of Providers selected for this Workspace
func (o LookupWorkspaceResultOutput) Providers() ProviderResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) []ProviderResponse { return v.Providers }).(ProviderResponseArrayOutput)
}

// Provisioning status field
func (o LookupWorkspaceResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

// ARM Resource Id of the storage account associated with this workspace.
func (o LookupWorkspaceResultOutput) StorageAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.StorageAccount }).(pulumi.StringPtrOutput)
}

// System metadata
func (o LookupWorkspaceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Resource tags.
func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

// Whether the current workspace is ready to accept Jobs.
func (o LookupWorkspaceResultOutput) Usable() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Usable }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
