// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20191001

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
	err := ctx.Invoke("azure-native:machinelearning/v20191001:getWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceArgs struct {
	// The name of the resource group to which the machine learning workspace belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the machine learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// An object that represents a machine learning workspace.
type LookupWorkspaceResult struct {
	// The creation time for this workspace resource.
	CreationTime string `pulumi:"creationTime"`
	// The resource ID.
	Id string `pulumi:"id"`
	// The key vault identifier used for encrypted workspaces.
	KeyVaultIdentifierId *string `pulumi:"keyVaultIdentifierId"`
	// The location of the resource. This cannot be changed after the resource is created.
	Location string `pulumi:"location"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// The email id of the owner for this workspace.
	OwnerEmail string `pulumi:"ownerEmail"`
	// The sku of the workspace.
	Sku *SkuResponse `pulumi:"sku"`
	// The regional endpoint for the machine learning studio service which hosts this workspace.
	StudioEndpoint string `pulumi:"studioEndpoint"`
	// The tags of the resource.
	Tags map[string]string `pulumi:"tags"`
	// The type of the resource.
	Type string `pulumi:"type"`
	// The fully qualified arm id of the storage account associated with this workspace.
	UserStorageAccountId string `pulumi:"userStorageAccountId"`
	// The immutable id associated with this workspace.
	WorkspaceId string `pulumi:"workspaceId"`
	// The current state of workspace resource.
	WorkspaceState string `pulumi:"workspaceState"`
	// The type of this workspace.
	WorkspaceType string `pulumi:"workspaceType"`
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
	// The name of the resource group to which the machine learning workspace belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the machine learning workspace.
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

// The creation time for this workspace resource.
func (o LookupWorkspaceResultOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.CreationTime }).(pulumi.StringOutput)
}

// The resource ID.
func (o LookupWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The key vault identifier used for encrypted workspaces.
func (o LookupWorkspaceResultOutput) KeyVaultIdentifierId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *string { return v.KeyVaultIdentifierId }).(pulumi.StringPtrOutput)
}

// The location of the resource. This cannot be changed after the resource is created.
func (o LookupWorkspaceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Location }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The email id of the owner for this workspace.
func (o LookupWorkspaceResultOutput) OwnerEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.OwnerEmail }).(pulumi.StringOutput)
}

// The sku of the workspace.
func (o LookupWorkspaceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// The regional endpoint for the machine learning studio service which hosts this workspace.
func (o LookupWorkspaceResultOutput) StudioEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.StudioEndpoint }).(pulumi.StringOutput)
}

// The tags of the resource.
func (o LookupWorkspaceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the resource.
func (o LookupWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The fully qualified arm id of the storage account associated with this workspace.
func (o LookupWorkspaceResultOutput) UserStorageAccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.UserStorageAccountId }).(pulumi.StringOutput)
}

// The immutable id associated with this workspace.
func (o LookupWorkspaceResultOutput) WorkspaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceId }).(pulumi.StringOutput)
}

// The current state of workspace resource.
func (o LookupWorkspaceResultOutput) WorkspaceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceState }).(pulumi.StringOutput)
}

// The type of this workspace.
func (o LookupWorkspaceResultOutput) WorkspaceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceResult) string { return v.WorkspaceType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceResultOutput{})
}
