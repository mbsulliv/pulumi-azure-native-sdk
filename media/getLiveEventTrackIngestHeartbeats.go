// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package media

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get track ingest heartbeat events telemetry of a live event.
// Azure REST API version: 2022-11-01.
func GetLiveEventTrackIngestHeartbeats(ctx *pulumi.Context, args *GetLiveEventTrackIngestHeartbeatsArgs, opts ...pulumi.InvokeOption) (*GetLiveEventTrackIngestHeartbeatsResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv GetLiveEventTrackIngestHeartbeatsResult
	err := ctx.Invoke("azure-native:media:getLiveEventTrackIngestHeartbeats", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetLiveEventTrackIngestHeartbeatsArgs struct {
	// The Media Services account name.
	AccountName string `pulumi:"accountName"`
	// The name of the live event, maximum length is 32.
	LiveEventName string `pulumi:"liveEventName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Get live event track ingest heart beats result.
type GetLiveEventTrackIngestHeartbeatsResult struct {
	// The result of the get live event track events.
	Value []LiveEventTrackEventResponse `pulumi:"value"`
}

func GetLiveEventTrackIngestHeartbeatsOutput(ctx *pulumi.Context, args GetLiveEventTrackIngestHeartbeatsOutputArgs, opts ...pulumi.InvokeOption) GetLiveEventTrackIngestHeartbeatsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetLiveEventTrackIngestHeartbeatsResult, error) {
			args := v.(GetLiveEventTrackIngestHeartbeatsArgs)
			r, err := GetLiveEventTrackIngestHeartbeats(ctx, &args, opts...)
			var s GetLiveEventTrackIngestHeartbeatsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetLiveEventTrackIngestHeartbeatsResultOutput)
}

type GetLiveEventTrackIngestHeartbeatsOutputArgs struct {
	// The Media Services account name.
	AccountName pulumi.StringInput `pulumi:"accountName"`
	// The name of the live event, maximum length is 32.
	LiveEventName pulumi.StringInput `pulumi:"liveEventName"`
	// The name of the resource group within the Azure subscription.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (GetLiveEventTrackIngestHeartbeatsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLiveEventTrackIngestHeartbeatsArgs)(nil)).Elem()
}

// Get live event track ingest heart beats result.
type GetLiveEventTrackIngestHeartbeatsResultOutput struct{ *pulumi.OutputState }

func (GetLiveEventTrackIngestHeartbeatsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLiveEventTrackIngestHeartbeatsResult)(nil)).Elem()
}

func (o GetLiveEventTrackIngestHeartbeatsResultOutput) ToGetLiveEventTrackIngestHeartbeatsResultOutput() GetLiveEventTrackIngestHeartbeatsResultOutput {
	return o
}

func (o GetLiveEventTrackIngestHeartbeatsResultOutput) ToGetLiveEventTrackIngestHeartbeatsResultOutputWithContext(ctx context.Context) GetLiveEventTrackIngestHeartbeatsResultOutput {
	return o
}

// The result of the get live event track events.
func (o GetLiveEventTrackIngestHeartbeatsResultOutput) Value() LiveEventTrackEventResponseArrayOutput {
	return o.ApplyT(func(v GetLiveEventTrackIngestHeartbeatsResult) []LiveEventTrackEventResponse { return v.Value }).(LiveEventTrackEventResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLiveEventTrackIngestHeartbeatsResultOutput{})
}