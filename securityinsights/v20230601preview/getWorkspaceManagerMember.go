// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a workspace manager member
func LookupWorkspaceManagerMember(ctx *pulumi.Context, args *LookupWorkspaceManagerMemberArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceManagerMemberResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceManagerMemberResult
	err := ctx.Invoke("azure-native:securityinsights/v20230601preview:getWorkspaceManagerMember", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceManagerMemberArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace manager member
	WorkspaceManagerMemberName string `pulumi:"workspaceManagerMemberName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The workspace manager member
type LookupWorkspaceManagerMemberResult struct {
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Fully qualified resource ID of the target Sentinel workspace joining the given Sentinel workspace manager
	TargetWorkspaceResourceId string `pulumi:"targetWorkspaceResourceId"`
	// Tenant id of the target Sentinel workspace joining the given Sentinel workspace manager
	TargetWorkspaceTenantId string `pulumi:"targetWorkspaceTenantId"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceManagerMemberOutput(ctx *pulumi.Context, args LookupWorkspaceManagerMemberOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceManagerMemberResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceManagerMemberResult, error) {
			args := v.(LookupWorkspaceManagerMemberArgs)
			r, err := LookupWorkspaceManagerMember(ctx, &args, opts...)
			var s LookupWorkspaceManagerMemberResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceManagerMemberResultOutput)
}

type LookupWorkspaceManagerMemberOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace manager member
	WorkspaceManagerMemberName pulumi.StringInput `pulumi:"workspaceManagerMemberName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceManagerMemberOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceManagerMemberArgs)(nil)).Elem()
}

// The workspace manager member
type LookupWorkspaceManagerMemberResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceManagerMemberResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceManagerMemberResult)(nil)).Elem()
}

func (o LookupWorkspaceManagerMemberResultOutput) ToLookupWorkspaceManagerMemberResultOutput() LookupWorkspaceManagerMemberResultOutput {
	return o
}

func (o LookupWorkspaceManagerMemberResultOutput) ToLookupWorkspaceManagerMemberResultOutputWithContext(ctx context.Context) LookupWorkspaceManagerMemberResultOutput {
	return o
}

func (o LookupWorkspaceManagerMemberResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceManagerMemberResult] {
	return pulumix.Output[LookupWorkspaceManagerMemberResult]{
		OutputState: o.OutputState,
	}
}

// Resource Etag.
func (o LookupWorkspaceManagerMemberResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerMemberResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceManagerMemberResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerMemberResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceManagerMemberResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerMemberResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWorkspaceManagerMemberResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerMemberResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Fully qualified resource ID of the target Sentinel workspace joining the given Sentinel workspace manager
func (o LookupWorkspaceManagerMemberResultOutput) TargetWorkspaceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerMemberResult) string { return v.TargetWorkspaceResourceId }).(pulumi.StringOutput)
}

// Tenant id of the target Sentinel workspace joining the given Sentinel workspace manager
func (o LookupWorkspaceManagerMemberResultOutput) TargetWorkspaceTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerMemberResult) string { return v.TargetWorkspaceTenantId }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceManagerMemberResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerMemberResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceManagerMemberResultOutput{})
}
