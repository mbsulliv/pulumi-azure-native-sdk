// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets status telemetry of a live event.
// Azure REST API version: 2022-11-01.
func GetLiveEventStatus(ctx *pulumi.Context, args *GetLiveEventStatusArgs, opts ...pulumi.InvokeOption) (*GetLiveEventStatusResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetLiveEventStatusResult
	err := ctx.Invoke("azure-native:media:getLiveEventStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLiveEventStatusArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The name of the live event, maximum length is 32.
	LiveEventName string `pulumi:"liveEventName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Get live event status result.
type GetLiveEventStatusResult struct {
	// The result of the get live event status.
	Value []LiveEventStatusResponse `pulumi:"value"`
}

func GetLiveEventStatusOutput(ctx *pulumi.Context, args GetLiveEventStatusOutputArgs, opts ...pulumi.InvokeOption) GetLiveEventStatusResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLiveEventStatusResult, error) {
			args := v.(GetLiveEventStatusArgs)
			r, err := GetLiveEventStatus(ctx, &args, opts...)
			var s GetLiveEventStatusResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLiveEventStatusResultOutput)
}

type GetLiveEventStatusOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the live event, maximum length is 32.
	LiveEventName pulumi.StringInput `pulumi:"liveEventName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetLiveEventStatusOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLiveEventStatusArgs)(nil)).Elem()
}

// Get live event status result.
type GetLiveEventStatusResultOutput struct{ *pulumi.OutputState }

func (GetLiveEventStatusResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLiveEventStatusResult)(nil)).Elem()
}

func (o GetLiveEventStatusResultOutput) ToGetLiveEventStatusResultOutput() GetLiveEventStatusResultOutput {
	return o
}

func (o GetLiveEventStatusResultOutput) ToGetLiveEventStatusResultOutputWithContext(ctx context.Context) GetLiveEventStatusResultOutput {
	return o
}

// The result of the get live event status.
func (o GetLiveEventStatusResultOutput) Value() LiveEventStatusResponseArrayOutput {
	return o.ApplyT(func(v GetLiveEventStatusResult) []LiveEventStatusResponse { return v.Value }).(LiveEventStatusResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLiveEventStatusResultOutput{})
}