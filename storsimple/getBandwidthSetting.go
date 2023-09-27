// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package storsimple

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Returns the properties of the specified bandwidth setting name.
// Azure REST API version: 2017-06-01.
func LookupBandwidthSetting(ctx *pulumi.Context, args *LookupBandwidthSettingArgs, opts ...pulumi.InvokeOption) (*LookupBandwidthSettingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBandwidthSettingResult
	err := ctx.Invoke("azure-native:storsimple:getBandwidthSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBandwidthSettingArgs struct {
	// The name of bandwidth setting to be fetched.
	BandwidthSettingName string `pulumi:"bandwidthSettingName"`
	// The manager name
	ManagerName string `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The bandwidth setting.
type LookupBandwidthSettingResult struct {
	// The path ID that uniquely identifies the object.
	Id string `pulumi:"id"`
	// The Kind of the object. Currently only Series8000 is supported
	Kind *string `pulumi:"kind"`
	// The name of the object.
	Name string `pulumi:"name"`
	// The schedules.
	Schedules []BandwidthScheduleResponse `pulumi:"schedules"`
	// The hierarchical type of the object.
	Type string `pulumi:"type"`
	// The number of volumes that uses the bandwidth setting.
	VolumeCount int `pulumi:"volumeCount"`
}

func LookupBandwidthSettingOutput(ctx *pulumi.Context, args LookupBandwidthSettingOutputArgs, opts ...pulumi.InvokeOption) LookupBandwidthSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBandwidthSettingResult, error) {
			args := v.(LookupBandwidthSettingArgs)
			r, err := LookupBandwidthSetting(ctx, &args, opts...)
			var s LookupBandwidthSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBandwidthSettingResultOutput)
}

type LookupBandwidthSettingOutputArgs struct {
	// The name of bandwidth setting to be fetched.
	BandwidthSettingName pulumi.StringInput `pulumi:"bandwidthSettingName"`
	// The manager name
	ManagerName pulumi.StringInput `pulumi:"managerName"`
	// The resource group name
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupBandwidthSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBandwidthSettingArgs)(nil)).Elem()
}

// The bandwidth setting.
type LookupBandwidthSettingResultOutput struct{ *pulumi.OutputState }

func (LookupBandwidthSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBandwidthSettingResult)(nil)).Elem()
}

func (o LookupBandwidthSettingResultOutput) ToLookupBandwidthSettingResultOutput() LookupBandwidthSettingResultOutput {
	return o
}

func (o LookupBandwidthSettingResultOutput) ToLookupBandwidthSettingResultOutputWithContext(ctx context.Context) LookupBandwidthSettingResultOutput {
	return o
}

func (o LookupBandwidthSettingResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBandwidthSettingResult] {
	return pulumix.Output[LookupBandwidthSettingResult]{
		OutputState: o.OutputState,
	}
}

// The path ID that uniquely identifies the object.
func (o LookupBandwidthSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// The Kind of the object. Currently only Series8000 is supported
func (o LookupBandwidthSettingResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// The name of the object.
func (o LookupBandwidthSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// The schedules.
func (o LookupBandwidthSettingResultOutput) Schedules() BandwidthScheduleResponseArrayOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) []BandwidthScheduleResponse { return v.Schedules }).(BandwidthScheduleResponseArrayOutput)
}

// The hierarchical type of the object.
func (o LookupBandwidthSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

// The number of volumes that uses the bandwidth setting.
func (o LookupBandwidthSettingResultOutput) VolumeCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBandwidthSettingResult) int { return v.VolumeCount }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBandwidthSettingResultOutput{})
}
