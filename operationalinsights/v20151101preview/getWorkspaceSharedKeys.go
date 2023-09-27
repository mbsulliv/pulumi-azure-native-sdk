// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20151101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the shared keys for a workspace.
func GetWorkspaceSharedKeys(ctx *pulumi.Context, args *GetWorkspaceSharedKeysArgs, opts ...pulumi.InvokeOption) (*GetWorkspaceSharedKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetWorkspaceSharedKeysResult
	err := ctx.Invoke("azure-native:operationalinsights/v20151101preview:getWorkspaceSharedKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetWorkspaceSharedKeysArgs struct {
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Log Analytics Workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// The shared keys for a workspace.
type GetWorkspaceSharedKeysResult struct {
	// The primary shared key of a workspace.
	PrimarySharedKey *string `pulumi:"primarySharedKey"`
	// The secondary shared key of a workspace.
	SecondarySharedKey *string `pulumi:"secondarySharedKey"`
}

func GetWorkspaceSharedKeysOutput(ctx *pulumi.Context, args GetWorkspaceSharedKeysOutputArgs, opts ...pulumi.InvokeOption) GetWorkspaceSharedKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetWorkspaceSharedKeysResult, error) {
			args := v.(GetWorkspaceSharedKeysArgs)
			r, err := GetWorkspaceSharedKeys(ctx, &args, opts...)
			var s GetWorkspaceSharedKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetWorkspaceSharedKeysResultOutput)
}

type GetWorkspaceSharedKeysOutputArgs struct {
	// The name of the resource group to get. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the Log Analytics Workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (GetWorkspaceSharedKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkspaceSharedKeysArgs)(nil)).Elem()
}

// The shared keys for a workspace.
type GetWorkspaceSharedKeysResultOutput struct{ *pulumi.OutputState }

func (GetWorkspaceSharedKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetWorkspaceSharedKeysResult)(nil)).Elem()
}

func (o GetWorkspaceSharedKeysResultOutput) ToGetWorkspaceSharedKeysResultOutput() GetWorkspaceSharedKeysResultOutput {
	return o
}

func (o GetWorkspaceSharedKeysResultOutput) ToGetWorkspaceSharedKeysResultOutputWithContext(ctx context.Context) GetWorkspaceSharedKeysResultOutput {
	return o
}

func (o GetWorkspaceSharedKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetWorkspaceSharedKeysResult] {
	return pulumix.Output[GetWorkspaceSharedKeysResult]{
		OutputState: o.OutputState,
	}
}

// The primary shared key of a workspace.
func (o GetWorkspaceSharedKeysResultOutput) PrimarySharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWorkspaceSharedKeysResult) *string { return v.PrimarySharedKey }).(pulumi.StringPtrOutput)
}

// The secondary shared key of a workspace.
func (o GetWorkspaceSharedKeysResultOutput) SecondarySharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetWorkspaceSharedKeysResult) *string { return v.SecondarySharedKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetWorkspaceSharedKeysResultOutput{})
}
