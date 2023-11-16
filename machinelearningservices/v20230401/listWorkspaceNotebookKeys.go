// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List keys of a notebook.
func ListWorkspaceNotebookKeys(ctx *pulumi.Context, args *ListWorkspaceNotebookKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceNotebookKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkspaceNotebookKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401:listWorkspaceNotebookKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceNotebookKeysArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type ListWorkspaceNotebookKeysResult struct {
	PrimaryAccessKey   string `pulumi:"primaryAccessKey"`
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
}

func ListWorkspaceNotebookKeysOutput(ctx *pulumi.Context, args ListWorkspaceNotebookKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceNotebookKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceNotebookKeysResult, error) {
			args := v.(ListWorkspaceNotebookKeysArgs)
			r, err := ListWorkspaceNotebookKeys(ctx, &args, opts...)
			var s ListWorkspaceNotebookKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceNotebookKeysResultOutput)
}

type ListWorkspaceNotebookKeysOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceNotebookKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceNotebookKeysArgs)(nil)).Elem()
}

type ListWorkspaceNotebookKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceNotebookKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceNotebookKeysResult)(nil)).Elem()
}

func (o ListWorkspaceNotebookKeysResultOutput) ToListWorkspaceNotebookKeysResultOutput() ListWorkspaceNotebookKeysResultOutput {
	return o
}

func (o ListWorkspaceNotebookKeysResultOutput) ToListWorkspaceNotebookKeysResultOutputWithContext(ctx context.Context) ListWorkspaceNotebookKeysResultOutput {
	return o
}

func (o ListWorkspaceNotebookKeysResultOutput) PrimaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookKeysResult) string { return v.PrimaryAccessKey }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookKeysResultOutput) SecondaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookKeysResult) string { return v.SecondaryAccessKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceNotebookKeysResultOutput{})
}
