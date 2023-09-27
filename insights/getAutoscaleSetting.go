// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets an autoscale setting
// Azure REST API version: 2022-10-01.
func LookupAutoscaleSetting(ctx *pulumi.Context, args *LookupAutoscaleSettingArgs, opts ...pulumi.InvokeOption) (*LookupAutoscaleSettingResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAutoscaleSettingResult
	err := ctx.Invoke("azure-native:insights:getAutoscaleSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupAutoscaleSettingArgs struct {
	// The autoscale setting name.
	AutoscaleSettingName string `pulumi:"autoscaleSettingName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The autoscale setting resource.
type LookupAutoscaleSettingResult struct {
	// the enabled flag. Specifies whether automatic scaling is enabled for the resource. The default value is 'false'.
	Enabled *bool `pulumi:"enabled"`
	// Azure resource Id
	Id string `pulumi:"id"`
	// Resource location
	Location string `pulumi:"location"`
	// Azure resource name
	Name string `pulumi:"name"`
	// the collection of notifications.
	Notifications []AutoscaleNotificationResponse `pulumi:"notifications"`
	// the predictive autoscale policy mode.
	PredictiveAutoscalePolicy *PredictiveAutoscalePolicyResponse `pulumi:"predictiveAutoscalePolicy"`
	// the collection of automatic scaling profiles that specify different scaling parameters for different time periods. A maximum of 20 profiles can be specified.
	Profiles []AutoscaleProfileResponse `pulumi:"profiles"`
	// The system metadata related to the response.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater in length than 128 characters and a value no greater in length than 256 characters.
	Tags map[string]string `pulumi:"tags"`
	// the location of the resource that the autoscale setting should be added to.
	TargetResourceLocation *string `pulumi:"targetResourceLocation"`
	// the resource identifier of the resource that the autoscale setting should be added to.
	TargetResourceUri *string `pulumi:"targetResourceUri"`
	// Azure resource type
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupAutoscaleSettingResult
func (val *LookupAutoscaleSettingResult) Defaults() *LookupAutoscaleSettingResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if tmp.Enabled == nil {
		enabled_ := false
		tmp.Enabled = &enabled_
	}
	return &tmp
}

func LookupAutoscaleSettingOutput(ctx *pulumi.Context, args LookupAutoscaleSettingOutputArgs, opts ...pulumi.InvokeOption) LookupAutoscaleSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAutoscaleSettingResult, error) {
			args := v.(LookupAutoscaleSettingArgs)
			r, err := LookupAutoscaleSetting(ctx, &args, opts...)
			var s LookupAutoscaleSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAutoscaleSettingResultOutput)
}

type LookupAutoscaleSettingOutputArgs struct {
	// The autoscale setting name.
	AutoscaleSettingName pulumi.StringInput `pulumi:"autoscaleSettingName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupAutoscaleSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoscaleSettingArgs)(nil)).Elem()
}

// The autoscale setting resource.
type LookupAutoscaleSettingResultOutput struct{ *pulumi.OutputState }

func (LookupAutoscaleSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAutoscaleSettingResult)(nil)).Elem()
}

func (o LookupAutoscaleSettingResultOutput) ToLookupAutoscaleSettingResultOutput() LookupAutoscaleSettingResultOutput {
	return o
}

func (o LookupAutoscaleSettingResultOutput) ToLookupAutoscaleSettingResultOutputWithContext(ctx context.Context) LookupAutoscaleSettingResultOutput {
	return o
}

func (o LookupAutoscaleSettingResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAutoscaleSettingResult] {
	return pulumix.Output[LookupAutoscaleSettingResult]{
		OutputState: o.OutputState,
	}
}

// the enabled flag. Specifies whether automatic scaling is enabled for the resource. The default value is 'false'.
func (o LookupAutoscaleSettingResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// Azure resource Id
func (o LookupAutoscaleSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

// Resource location
func (o LookupAutoscaleSettingResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) string { return v.Location }).(pulumi.StringOutput)
}

// Azure resource name
func (o LookupAutoscaleSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

// the collection of notifications.
func (o LookupAutoscaleSettingResultOutput) Notifications() AutoscaleNotificationResponseArrayOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) []AutoscaleNotificationResponse { return v.Notifications }).(AutoscaleNotificationResponseArrayOutput)
}

// the predictive autoscale policy mode.
func (o LookupAutoscaleSettingResultOutput) PredictiveAutoscalePolicy() PredictiveAutoscalePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) *PredictiveAutoscalePolicyResponse {
		return v.PredictiveAutoscalePolicy
	}).(PredictiveAutoscalePolicyResponsePtrOutput)
}

// the collection of automatic scaling profiles that specify different scaling parameters for different time periods. A maximum of 20 profiles can be specified.
func (o LookupAutoscaleSettingResultOutput) Profiles() AutoscaleProfileResponseArrayOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) []AutoscaleProfileResponse { return v.Profiles }).(AutoscaleProfileResponseArrayOutput)
}

// The system metadata related to the response.
func (o LookupAutoscaleSettingResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no greater in length than 128 characters and a value no greater in length than 256 characters.
func (o LookupAutoscaleSettingResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// the location of the resource that the autoscale setting should be added to.
func (o LookupAutoscaleSettingResultOutput) TargetResourceLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) *string { return v.TargetResourceLocation }).(pulumi.StringPtrOutput)
}

// the resource identifier of the resource that the autoscale setting should be added to.
func (o LookupAutoscaleSettingResultOutput) TargetResourceUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) *string { return v.TargetResourceUri }).(pulumi.StringPtrOutput)
}

// Azure resource type
func (o LookupAutoscaleSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAutoscaleSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAutoscaleSettingResultOutput{})
}
