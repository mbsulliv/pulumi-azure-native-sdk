// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20231101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// List globally enabled APMs for a Service.
func ListServiceGloballyEnabledApms(ctx *pulumi.Context, args *ListServiceGloballyEnabledApmsArgs, opts ...pulumi.InvokeOption) (*ListServiceGloballyEnabledApmsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv ListServiceGloballyEnabledApmsResult
	err := ctx.Invoke("azure-native:appplatform/v20231101preview:listServiceGloballyEnabledApms", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListServiceGloballyEnabledApmsArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Globally enabled APMs payload
type ListServiceGloballyEnabledApmsResult struct {
	// Collection of the globally enabled APMs
	Value []string `pulumi:"value"`
}

func ListServiceGloballyEnabledApmsOutput(ctx *pulumi.Context, args ListServiceGloballyEnabledApmsOutputArgs, opts ...pulumi.InvokeOption) ListServiceGloballyEnabledApmsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListServiceGloballyEnabledApmsResult, error) {
			args := v.(ListServiceGloballyEnabledApmsArgs)
			r, err := ListServiceGloballyEnabledApms(ctx, &args, opts...)
			var s ListServiceGloballyEnabledApmsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListServiceGloballyEnabledApmsResultOutput)
}

type ListServiceGloballyEnabledApmsOutputArgs struct {
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (ListServiceGloballyEnabledApmsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServiceGloballyEnabledApmsArgs)(nil)).Elem()
}

// Globally enabled APMs payload
type ListServiceGloballyEnabledApmsResultOutput struct{ *pulumi.OutputState }

func (ListServiceGloballyEnabledApmsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListServiceGloballyEnabledApmsResult)(nil)).Elem()
}

func (o ListServiceGloballyEnabledApmsResultOutput) ToListServiceGloballyEnabledApmsResultOutput() ListServiceGloballyEnabledApmsResultOutput {
	return o
}

func (o ListServiceGloballyEnabledApmsResultOutput) ToListServiceGloballyEnabledApmsResultOutputWithContext(ctx context.Context) ListServiceGloballyEnabledApmsResultOutput {
	return o
}

// Collection of the globally enabled APMs
func (o ListServiceGloballyEnabledApmsResultOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListServiceGloballyEnabledApmsResult) []string { return v.Value }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListServiceGloballyEnabledApmsResultOutput{})
}
