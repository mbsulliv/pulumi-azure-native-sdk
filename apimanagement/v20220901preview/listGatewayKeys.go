// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Retrieves gateway keys.
func ListGatewayKeys(ctx *pulumi.Context, args *ListGatewayKeysArgs, opts ...pulumi.InvokeOption) (*ListGatewayKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListGatewayKeysResult
	err := ctx.Invoke("azure-native:apimanagement/v20220901preview:listGatewayKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGatewayKeysArgs struct {
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId string `pulumi:"gatewayId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName string `pulumi:"serviceName"`
}

// Gateway authentication keys.
type ListGatewayKeysResult struct {
	// Primary gateway key.
	Primary *string `pulumi:"primary"`
	// Secondary gateway key.
	Secondary *string `pulumi:"secondary"`
}

func ListGatewayKeysOutput(ctx *pulumi.Context, args ListGatewayKeysOutputArgs, opts ...pulumi.InvokeOption) ListGatewayKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListGatewayKeysResult, error) {
			args := v.(ListGatewayKeysArgs)
			r, err := ListGatewayKeys(ctx, &args, opts...)
			var s ListGatewayKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListGatewayKeysResultOutput)
}

type ListGatewayKeysOutputArgs struct {
	// Gateway entity identifier. Must be unique in the current API Management service instance. Must not have value 'managed'
	GatewayId pulumi.StringInput `pulumi:"gatewayId"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the API Management service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ListGatewayKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGatewayKeysArgs)(nil)).Elem()
}

// Gateway authentication keys.
type ListGatewayKeysResultOutput struct{ *pulumi.OutputState }

func (ListGatewayKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListGatewayKeysResult)(nil)).Elem()
}

func (o ListGatewayKeysResultOutput) ToListGatewayKeysResultOutput() ListGatewayKeysResultOutput {
	return o
}

func (o ListGatewayKeysResultOutput) ToListGatewayKeysResultOutputWithContext(ctx context.Context) ListGatewayKeysResultOutput {
	return o
}

func (o ListGatewayKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[ListGatewayKeysResult] {
	return pulumix.Output[ListGatewayKeysResult]{
		OutputState: o.OutputState,
	}
}

// Primary gateway key.
func (o ListGatewayKeysResultOutput) Primary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListGatewayKeysResult) *string { return v.Primary }).(pulumi.StringPtrOutput)
}

// Secondary gateway key.
func (o ListGatewayKeysResultOutput) Secondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListGatewayKeysResult) *string { return v.Secondary }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListGatewayKeysResultOutput{})
}
