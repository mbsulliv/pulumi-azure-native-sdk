// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package securityinsights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets a workspace manager group
// Azure REST API version: 2023-06-01-preview.
func LookupWorkspaceManagerGroup(ctx *pulumi.Context, args *LookupWorkspaceManagerGroupArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceManagerGroupResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceManagerGroupResult
	err := ctx.Invoke("azure-native:securityinsights:getWorkspaceManagerGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceManagerGroupArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace manager group
	WorkspaceManagerGroupName string `pulumi:"workspaceManagerGroupName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The workspace manager group
type LookupWorkspaceManagerGroupResult struct {
	// The description of the workspace manager group
	Description *string `pulumi:"description"`
	// The display name of the workspace manager group
	DisplayName string `pulumi:"displayName"`
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The names of the workspace manager members participating in this group.
	MemberResourceNames []string `pulumi:"memberResourceNames"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceManagerGroupOutput(ctx *pulumi.Context, args LookupWorkspaceManagerGroupOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceManagerGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceManagerGroupResult, error) {
			args := v.(LookupWorkspaceManagerGroupArgs)
			r, err := LookupWorkspaceManagerGroup(ctx, &args, opts...)
			var s LookupWorkspaceManagerGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceManagerGroupResultOutput)
}

type LookupWorkspaceManagerGroupOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace manager group
	WorkspaceManagerGroupName pulumi.StringInput `pulumi:"workspaceManagerGroupName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceManagerGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceManagerGroupArgs)(nil)).Elem()
}

// The workspace manager group
type LookupWorkspaceManagerGroupResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceManagerGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceManagerGroupResult)(nil)).Elem()
}

func (o LookupWorkspaceManagerGroupResultOutput) ToLookupWorkspaceManagerGroupResultOutput() LookupWorkspaceManagerGroupResultOutput {
	return o
}

func (o LookupWorkspaceManagerGroupResultOutput) ToLookupWorkspaceManagerGroupResultOutputWithContext(ctx context.Context) LookupWorkspaceManagerGroupResultOutput {
	return o
}

func (o LookupWorkspaceManagerGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupWorkspaceManagerGroupResult] {
	return pulumix.Output[LookupWorkspaceManagerGroupResult]{
		OutputState: o.OutputState,
	}
}

// The description of the workspace manager group
func (o LookupWorkspaceManagerGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The display name of the workspace manager group
func (o LookupWorkspaceManagerGroupResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerGroupResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Resource Etag.
func (o LookupWorkspaceManagerGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceManagerGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

// The names of the workspace manager members participating in this group.
func (o LookupWorkspaceManagerGroupResultOutput) MemberResourceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerGroupResult) []string { return v.MemberResourceNames }).(pulumi.StringArrayOutput)
}

// The name of the resource
func (o LookupWorkspaceManagerGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWorkspaceManagerGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceManagerGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceManagerGroupResultOutput{})
}
