// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkspaceConnection(ctx *pulumi.Context, args *LookupWorkspaceConnectionArgs, opts ...pulumi.InvokeOption) (*LookupWorkspaceConnectionResult, error) {
	var rv LookupWorkspaceConnectionResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:getWorkspaceConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkspaceConnectionArgs struct {
	// Friendly name of the workspace connection
	ConnectionName string `pulumi:"connectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type LookupWorkspaceConnectionResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
	// Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupWorkspaceConnectionOutput(ctx *pulumi.Context, args LookupWorkspaceConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupWorkspaceConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkspaceConnectionResult, error) {
			args := v.(LookupWorkspaceConnectionArgs)
			r, err := LookupWorkspaceConnection(ctx, &args, opts...)
			var s LookupWorkspaceConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkspaceConnectionResultOutput)
}

type LookupWorkspaceConnectionOutputArgs struct {
	// Friendly name of the workspace connection
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupWorkspaceConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceConnectionArgs)(nil)).Elem()
}

type LookupWorkspaceConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupWorkspaceConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkspaceConnectionResult)(nil)).Elem()
}

func (o LookupWorkspaceConnectionResultOutput) ToLookupWorkspaceConnectionResultOutput() LookupWorkspaceConnectionResultOutput {
	return o
}

func (o LookupWorkspaceConnectionResultOutput) ToLookupWorkspaceConnectionResultOutputWithContext(ctx context.Context) LookupWorkspaceConnectionResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupWorkspaceConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupWorkspaceConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkspaceConnectionResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o LookupWorkspaceConnectionResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupWorkspaceConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkspaceConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkspaceConnectionResultOutput{})
}