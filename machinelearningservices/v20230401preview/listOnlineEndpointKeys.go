// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Keys for endpoint authentication.
func ListOnlineEndpointKeys(ctx *pulumi.Context, args *ListOnlineEndpointKeysArgs, opts ...pulumi.InvokeOption) (*ListOnlineEndpointKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListOnlineEndpointKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20230401preview:listOnlineEndpointKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListOnlineEndpointKeysArgs struct {
	// Online Endpoint name.
	EndpointName string `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Keys for endpoint authentication.
type ListOnlineEndpointKeysResult struct {
	// The primary key.
	PrimaryKey *string `pulumi:"primaryKey"`
	// The secondary key.
	SecondaryKey *string `pulumi:"secondaryKey"`
}

func ListOnlineEndpointKeysOutput(ctx *pulumi.Context, args ListOnlineEndpointKeysOutputArgs, opts ...pulumi.InvokeOption) ListOnlineEndpointKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListOnlineEndpointKeysResult, error) {
			args := v.(ListOnlineEndpointKeysArgs)
			r, err := ListOnlineEndpointKeys(ctx, &args, opts...)
			var s ListOnlineEndpointKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListOnlineEndpointKeysResultOutput)
}

type ListOnlineEndpointKeysOutputArgs struct {
	// Online Endpoint name.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListOnlineEndpointKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOnlineEndpointKeysArgs)(nil)).Elem()
}

// Keys for endpoint authentication.
type ListOnlineEndpointKeysResultOutput struct{ *pulumi.OutputState }

func (ListOnlineEndpointKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListOnlineEndpointKeysResult)(nil)).Elem()
}

func (o ListOnlineEndpointKeysResultOutput) ToListOnlineEndpointKeysResultOutput() ListOnlineEndpointKeysResultOutput {
	return o
}

func (o ListOnlineEndpointKeysResultOutput) ToListOnlineEndpointKeysResultOutputWithContext(ctx context.Context) ListOnlineEndpointKeysResultOutput {
	return o
}

// The primary key.
func (o ListOnlineEndpointKeysResultOutput) PrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOnlineEndpointKeysResult) *string { return v.PrimaryKey }).(pulumi.StringPtrOutput)
}

// The secondary key.
func (o ListOnlineEndpointKeysResultOutput) SecondaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListOnlineEndpointKeysResult) *string { return v.SecondaryKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListOnlineEndpointKeysResultOutput{})
}
