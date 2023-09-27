// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List storage account keys of a workspace.
func ListWorkspaceStorageAccountKeys(ctx *pulumi.Context, args *ListWorkspaceStorageAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceStorageAccountKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkspaceStorageAccountKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20220101preview:listWorkspaceStorageAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceStorageAccountKeysArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

type ListWorkspaceStorageAccountKeysResult struct {
	UserStorageKey string `pulumi:"userStorageKey"`
}

func ListWorkspaceStorageAccountKeysOutput(ctx *pulumi.Context, args ListWorkspaceStorageAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceStorageAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceStorageAccountKeysResult, error) {
			args := v.(ListWorkspaceStorageAccountKeysArgs)
			r, err := ListWorkspaceStorageAccountKeys(ctx, &args, opts...)
			var s ListWorkspaceStorageAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceStorageAccountKeysResultOutput)
}

type ListWorkspaceStorageAccountKeysOutputArgs struct {
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceStorageAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceStorageAccountKeysArgs)(nil)).Elem()
}

type ListWorkspaceStorageAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceStorageAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceStorageAccountKeysResult)(nil)).Elem()
}

func (o ListWorkspaceStorageAccountKeysResultOutput) ToListWorkspaceStorageAccountKeysResultOutput() ListWorkspaceStorageAccountKeysResultOutput {
	return o
}

func (o ListWorkspaceStorageAccountKeysResultOutput) ToListWorkspaceStorageAccountKeysResultOutputWithContext(ctx context.Context) ListWorkspaceStorageAccountKeysResultOutput {
	return o
}

func (o ListWorkspaceStorageAccountKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWorkspaceStorageAccountKeysResult] {
	return pulumix.Output[ListWorkspaceStorageAccountKeysResult]{
		OutputState: o.OutputState,
	}
}

func (o ListWorkspaceStorageAccountKeysResultOutput) UserStorageKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceStorageAccountKeysResult) string { return v.UserStorageKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceStorageAccountKeysResultOutput{})
}
