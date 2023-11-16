// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetMonitorDefaultKey(ctx *pulumi.Context, args *GetMonitorDefaultKeyArgs, opts ...pulumi.InvokeOption) (*GetMonitorDefaultKeyResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetMonitorDefaultKeyResult
	err := ctx.Invoke("azure-native:datadog/v20220601:getMonitorDefaultKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetMonitorDefaultKeyArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type GetMonitorDefaultKeyResult struct {
	// The time of creation of the API key.
	Created *string `pulumi:"created"`
	// The user that created the API key.
	CreatedBy *string `pulumi:"createdBy"`
	// The value of the API key.
	Key string `pulumi:"key"`
	// The name of the API key.
	Name *string `pulumi:"name"`
}

func GetMonitorDefaultKeyOutput(ctx *pulumi.Context, args GetMonitorDefaultKeyOutputArgs, opts ...pulumi.InvokeOption) GetMonitorDefaultKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetMonitorDefaultKeyResult, error) {
			args := v.(GetMonitorDefaultKeyArgs)
			r, err := GetMonitorDefaultKey(ctx, &args, opts...)
			var s GetMonitorDefaultKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetMonitorDefaultKeyResultOutput)
}

type GetMonitorDefaultKeyOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetMonitorDefaultKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitorDefaultKeyArgs)(nil)).Elem()
}

type GetMonitorDefaultKeyResultOutput struct{ *pulumi.OutputState }

func (GetMonitorDefaultKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetMonitorDefaultKeyResult)(nil)).Elem()
}

func (o GetMonitorDefaultKeyResultOutput) ToGetMonitorDefaultKeyResultOutput() GetMonitorDefaultKeyResultOutput {
	return o
}

func (o GetMonitorDefaultKeyResultOutput) ToGetMonitorDefaultKeyResultOutputWithContext(ctx context.Context) GetMonitorDefaultKeyResultOutput {
	return o
}

// The time of creation of the API key.
func (o GetMonitorDefaultKeyResultOutput) Created() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorDefaultKeyResult) *string { return v.Created }).(pulumi.StringPtrOutput)
}

// The user that created the API key.
func (o GetMonitorDefaultKeyResultOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorDefaultKeyResult) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

// The value of the API key.
func (o GetMonitorDefaultKeyResultOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v GetMonitorDefaultKeyResult) string { return v.Key }).(pulumi.StringOutput)
}

// The name of the API key.
func (o GetMonitorDefaultKeyResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetMonitorDefaultKeyResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetMonitorDefaultKeyResultOutput{})
}
