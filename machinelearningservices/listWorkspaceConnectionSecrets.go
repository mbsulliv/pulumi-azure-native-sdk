// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2023-06-01-preview.
func ListWorkspaceConnectionSecrets(ctx *pulumi.Context, args *ListWorkspaceConnectionSecretsArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceConnectionSecretsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkspaceConnectionSecretsResult
	err := ctx.Invoke("azure-native:machinelearningservices:listWorkspaceConnectionSecrets", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceConnectionSecretsArgs struct {
	// Friendly name of the workspace connection
	ConnectionName string `pulumi:"connectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type ListWorkspaceConnectionSecretsResult struct {
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

func ListWorkspaceConnectionSecretsOutput(ctx *pulumi.Context, args ListWorkspaceConnectionSecretsOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceConnectionSecretsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceConnectionSecretsResult, error) {
			args := v.(ListWorkspaceConnectionSecretsArgs)
			r, err := ListWorkspaceConnectionSecrets(ctx, &args, opts...)
			var s ListWorkspaceConnectionSecretsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceConnectionSecretsResultOutput)
}

type ListWorkspaceConnectionSecretsOutputArgs struct {
	// Friendly name of the workspace connection
	ConnectionName pulumi.StringInput `pulumi:"connectionName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceConnectionSecretsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceConnectionSecretsArgs)(nil)).Elem()
}

type ListWorkspaceConnectionSecretsResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceConnectionSecretsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceConnectionSecretsResult)(nil)).Elem()
}

func (o ListWorkspaceConnectionSecretsResultOutput) ToListWorkspaceConnectionSecretsResultOutput() ListWorkspaceConnectionSecretsResultOutput {
	return o
}

func (o ListWorkspaceConnectionSecretsResultOutput) ToListWorkspaceConnectionSecretsResultOutputWithContext(ctx context.Context) ListWorkspaceConnectionSecretsResultOutput {
	return o
}

func (o ListWorkspaceConnectionSecretsResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWorkspaceConnectionSecretsResult] {
	return pulumix.Output[ListWorkspaceConnectionSecretsResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o ListWorkspaceConnectionSecretsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceConnectionSecretsResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o ListWorkspaceConnectionSecretsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceConnectionSecretsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWorkspaceConnectionSecretsResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v ListWorkspaceConnectionSecretsResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Azure Resource Manager metadata containing createdBy and modifiedBy information.
func (o ListWorkspaceConnectionSecretsResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v ListWorkspaceConnectionSecretsResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o ListWorkspaceConnectionSecretsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceConnectionSecretsResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceConnectionSecretsResultOutput{})
}
