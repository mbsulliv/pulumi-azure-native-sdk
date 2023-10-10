// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Keys for endpoint authentication.
func ListServerlessEndpointKeys(ctx *pulumi.Context, args *ListServerlessEndpointKeysArgs, opts ...pulumi.InvokeOption) (*ListServerlessEndpointKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListServerlessEndpointKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230801preview:listServerlessEndpointKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServerlessEndpointKeysArgs struct {
	// Serverless Endpoint name.
	Name string `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Keys for endpoint authentication.
type ListServerlessEndpointKeysResult struct {
	// The primary key.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The secondary key.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListServerlessEndpointKeysOutput(ctx *pulumi.Context, args ListServerlessEndpointKeysOutputArgs, opts ...pulumi.InvokeOption) ListServerlessEndpointKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServerlessEndpointKeysResult, error) {
			args := v.(ListServerlessEndpointKeysArgs)
			r, err := ListServerlessEndpointKeys(ctx, &args, opts...)
			var s ListServerlessEndpointKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListServerlessEndpointKeysResultOutput)
}

type ListServerlessEndpointKeysOutputArgs struct {
	// Serverless Endpoint name.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListServerlessEndpointKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServerlessEndpointKeysArgs)(nil)).Elem()
}

// Keys for endpoint authentication.
type ListServerlessEndpointKeysResultOutput struct{ *pulumi.OutputState }

func (ListServerlessEndpointKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServerlessEndpointKeysResult)(nil)).Elem()
}

func (o ListServerlessEndpointKeysResultOutput) ToListServerlessEndpointKeysResultOutput() ListServerlessEndpointKeysResultOutput {
	return o
}

func (o ListServerlessEndpointKeysResultOutput) ToListServerlessEndpointKeysResultOutputWithContext(ctx context.Context) ListServerlessEndpointKeysResultOutput {
	return o
}

func (o ListServerlessEndpointKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListServerlessEndpointKeysResult] {
	return pulumix.Output[ListServerlessEndpointKeysResult]{
		OutputState: o.OutputState,
	}
}

// The primary key.
func (o ListServerlessEndpointKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServerlessEndpointKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// The secondary key.
func (o ListServerlessEndpointKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListServerlessEndpointKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServerlessEndpointKeysResultOutput{})
}
