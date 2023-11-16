// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a workspace manager assignment
func LookupWorkspaceManagerAssignment(ctx *pulumi.Context, args *LookupWorkspaceManagerAssignmentArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceManagerAssignmentResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupWorkspaceManagerAssignmentResult
	err := ctx.Invoke("azure-native:securityinsights/v20230601preview:getWorkspaceManagerAssignment", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceManagerAssignmentArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the workspace manager assignment
	WorkspaceManagerAssignmentName string `pulumi:"workspaceManagerAssignmentName"`
	// The name of the workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The workspace manager assignment
type LookupWorkspaceManagerAssignmentResult struct {
	// Resource Etag.
	Etag string `pulumi:"etag"`
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// List of resources included in this workspace manager assignment
	Items []AssignmentItemResponse `pulumi:"items"`
	// The time the last job associated to this assignment ended at
	LastJobEndTime string `pulumi:"lastJobEndTime"`
	// State of the last job associated to this assignment
	LastJobProvisioningState string `pulumi:"lastJobProvisioningState"`
	// The name of the resource
	Name string `pulumi:"name"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource name of the workspace manager group targeted by the workspace manager assignment
	TargetResourceName string `pulumi:"targetResourceName"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceManagerAssignmentOutput(ctx *pulumi.Context, args LookupWorkspaceManagerAssignmentOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceManagerAssignmentResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceManagerAssignmentResult, error) {
			args := v.(LookupWorkspaceManagerAssignmentArgs)
			r, err := LookupWorkspaceManagerAssignment(ctx, &args, opts...)
			var s LookupWorkspaceManagerAssignmentResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceManagerAssignmentResultOutput)
}

type LookupWorkspaceManagerAssignmentOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the workspace manager assignment
	WorkspaceManagerAssignmentName pulumi.StringInput `pulumi:"workspaceManagerAssignmentName"`
	// The name of the workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceManagerAssignmentOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceManagerAssignmentArgs)(nil)).Elem()
}

// The workspace manager assignment
type LookupWorkspaceManagerAssignmentResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceManagerAssignmentResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceManagerAssignmentResult)(nil)).Elem()
}

func (o LookupWorkspaceManagerAssignmentResultOutput) ToLookupWorkspaceManagerAssignmentResultOutput() LookupWorkspaceManagerAssignmentResultOutput {
	return o
}

func (o LookupWorkspaceManagerAssignmentResultOutput) ToLookupWorkspaceManagerAssignmentResultOutputWithContext(ctx context.Context) LookupWorkspaceManagerAssignmentResultOutput {
	return o
}

// Resource Etag.
func (o LookupWorkspaceManagerAssignmentResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerAssignmentResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceManagerAssignmentResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerAssignmentResult) string { return v.Id }).(pulumi.StringOutput)
}

// List of resources included in this workspace manager assignment
func (o LookupWorkspaceManagerAssignmentResultOutput) Items() AssignmentItemResponseArrayOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerAssignmentResult) []AssignmentItemResponse { return v.Items }).(AssignmentItemResponseArrayOutput)
}

// The time the last job associated to this assignment ended at
func (o LookupWorkspaceManagerAssignmentResultOutput) LastJobEndTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerAssignmentResult) string { return v.LastJobEndTime }).(pulumi.StringOutput)
}

// State of the last job associated to this assignment
func (o LookupWorkspaceManagerAssignmentResultOutput) LastJobProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerAssignmentResult) string { return v.LastJobProvisioningState }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceManagerAssignmentResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerAssignmentResult) string { return v.Name }).(pulumi.StringOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWorkspaceManagerAssignmentResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerAssignmentResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource name of the workspace manager group targeted by the workspace manager assignment
func (o LookupWorkspaceManagerAssignmentResultOutput) TargetResourceName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerAssignmentResult) string { return v.TargetResourceName }).(pulumi.StringOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceManagerAssignmentResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceManagerAssignmentResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceManagerAssignmentResultOutput{})
}
