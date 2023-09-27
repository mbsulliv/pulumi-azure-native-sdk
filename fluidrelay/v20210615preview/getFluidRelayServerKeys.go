// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210615preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// The set of available keys for this server.
func GetFluidRelayServerKeys(ctx *pulumi.Context, args *GetFluidRelayServerKeysArgs, opts ...pulumi.InvokeOption) (*GetFluidRelayServerKeysResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetFluidRelayServerKeysResult
	err := ctx.Invoke("azure-native:fluidrelay/v20210615preview:getFluidRelayServerKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetFluidRelayServerKeysArgs struct {
	// The resource name.
	Name string `pulumi:"name"`
	// The resource group containing the resource.
	ResourceGroup string `pulumi:"resourceGroup"`
}

// The set of available keys for this server.
type GetFluidRelayServerKeysResult struct {
	// The primary key for this server
	Key1 string `pulumi:"key1"`
	// The secondary key for this server
	Key2 string `pulumi:"key2"`
}

func GetFluidRelayServerKeysOutput(ctx *pulumi.Context, args GetFluidRelayServerKeysOutputArgs, opts ...pulumi.InvokeOption) GetFluidRelayServerKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFluidRelayServerKeysResult, error) {
			args := v.(GetFluidRelayServerKeysArgs)
			r, err := GetFluidRelayServerKeys(ctx, &args, opts...)
			var s GetFluidRelayServerKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFluidRelayServerKeysResultOutput)
}

type GetFluidRelayServerKeysOutputArgs struct {
	// The resource name.
	Name pulumi.StringInput `pulumi:"name"`
	// The resource group containing the resource.
	ResourceGroup pulumi.StringInput `pulumi:"resourceGroup"`
}

func (GetFluidRelayServerKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFluidRelayServerKeysArgs)(nil)).Elem()
}

// The set of available keys for this server.
type GetFluidRelayServerKeysResultOutput struct{ *pulumi.OutputState }

func (GetFluidRelayServerKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFluidRelayServerKeysResult)(nil)).Elem()
}

func (o GetFluidRelayServerKeysResultOutput) ToGetFluidRelayServerKeysResultOutput() GetFluidRelayServerKeysResultOutput {
	return o
}

func (o GetFluidRelayServerKeysResultOutput) ToGetFluidRelayServerKeysResultOutputWithContext(ctx context.Context) GetFluidRelayServerKeysResultOutput {
	return o
}

func (o GetFluidRelayServerKeysResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetFluidRelayServerKeysResult] {
	return pulumix.Output[GetFluidRelayServerKeysResult]{
		OutputState: o.OutputState,
	}
}

// The primary key for this server
func (o GetFluidRelayServerKeysResultOutput) Key1() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluidRelayServerKeysResult) string { return v.Key1 }).(pulumi.StringOutput)
}

// The secondary key for this server
func (o GetFluidRelayServerKeysResultOutput) Key2() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluidRelayServerKeysResult) string { return v.Key2 }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFluidRelayServerKeysResultOutput{})
}
