// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package machinelearning

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// List the authorization keys associated with this workspace.
// Azure REST API version: 2019-10-01.
func ListWorkspaceKeys(ctx *pulumi.Context, args *ListWorkspaceKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListWorkspaceKeysResult
	err := ctx.Invoke("azure-native:machinelearning:listWorkspaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceKeysArgs struct {
	// The name of the resource group to which the machine learning workspace belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the machine learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Workspace authorization keys for a workspace.
type ListWorkspaceKeysResult struct {
	// Primary authorization key for this workspace.
	PrimaryToken *string `pulumi:"primaryToken"`
	// Secondary authorization key for this workspace.
	SecondaryToken *string `pulumi:"secondaryToken"`
}

func ListWorkspaceKeysOutput(ctx *pulumi.Context, args ListWorkspaceKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceKeysResult, error) {
			args := v.(ListWorkspaceKeysArgs)
			r, err := ListWorkspaceKeys(ctx, &args, opts...)
			var s ListWorkspaceKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceKeysResultOutput)
}

type ListWorkspaceKeysOutputArgs struct {
	// The name of the resource group to which the machine learning workspace belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the machine learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceKeysArgs)(nil)).Elem()
}

// Workspace authorization keys for a workspace.
type ListWorkspaceKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceKeysResult)(nil)).Elem()
}

func (o ListWorkspaceKeysResultOutput) ToListWorkspaceKeysResultOutput() ListWorkspaceKeysResultOutput {
	return o
}

func (o ListWorkspaceKeysResultOutput) ToListWorkspaceKeysResultOutputWithContext(ctx context.Context) ListWorkspaceKeysResultOutput {
	return o
}

func (o ListWorkspaceKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListWorkspaceKeysResult] {
	return pulumix.Output[ListWorkspaceKeysResult]{
		OutputState: o.OutputState,
	}
}

// Primary authorization key for this workspace.
func (o ListWorkspaceKeysResultOutput) PrimaryToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) *string { return v.PrimaryToken }).(pulumi.StringPtrOutput)
}

// Secondary authorization key for this workspace.
func (o ListWorkspaceKeysResultOutput) SecondaryToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWorkspaceKeysResult) *string { return v.SecondaryToken }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceKeysResultOutput{})
}
