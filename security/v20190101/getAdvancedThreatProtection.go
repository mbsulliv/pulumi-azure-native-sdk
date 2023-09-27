// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20190101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the Advanced Threat Protection settings for the specified resource.
func LookupAdvancedThreatProtection(ctx *pulumi.Context, args *LookupAdvancedThreatProtectionArgs, opts ...pulumi.InvokeOption) (*LookupAdvancedThreatProtectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupAdvancedThreatProtectionResult
	err := ctx.Invoke("azure-native:security/v20190101:getAdvancedThreatProtection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAdvancedThreatProtectionArgs struct {
	// The identifier of the resource.
	ResourceId string `pulumi:"resourceId"`
	// Advanced Threat Protection setting name.
	SettingName string `pulumi:"settingName"`
}

// The Advanced Threat Protection resource.
type LookupAdvancedThreatProtectionResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// Indicates whether Advanced Threat Protection is enabled.
	IsEnabled *bool `pulumi:"isEnabled"`
	// Resource name
	Name string `pulumi:"name"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupAdvancedThreatProtectionOutput(ctx *pulumi.Context, args LookupAdvancedThreatProtectionOutputArgs, opts ...pulumi.InvokeOption) LookupAdvancedThreatProtectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAdvancedThreatProtectionResult, error) {
			args := v.(LookupAdvancedThreatProtectionArgs)
			r, err := LookupAdvancedThreatProtection(ctx, &args, opts...)
			var s LookupAdvancedThreatProtectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAdvancedThreatProtectionResultOutput)
}

type LookupAdvancedThreatProtectionOutputArgs struct {
	// The identifier of the resource.
	ResourceId pulumi.StringInput `pulumi:"resourceId"`
	// Advanced Threat Protection setting name.
	SettingName pulumi.StringInput `pulumi:"settingName"`
}

func (LookupAdvancedThreatProtectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdvancedThreatProtectionArgs)(nil)).Elem()
}

// The Advanced Threat Protection resource.
type LookupAdvancedThreatProtectionResultOutput struct{ *pulumi.OutputState }

func (LookupAdvancedThreatProtectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAdvancedThreatProtectionResult)(nil)).Elem()
}

func (o LookupAdvancedThreatProtectionResultOutput) ToLookupAdvancedThreatProtectionResultOutput() LookupAdvancedThreatProtectionResultOutput {
	return o
}

func (o LookupAdvancedThreatProtectionResultOutput) ToLookupAdvancedThreatProtectionResultOutputWithContext(ctx context.Context) LookupAdvancedThreatProtectionResultOutput {
	return o
}

func (o LookupAdvancedThreatProtectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAdvancedThreatProtectionResult] {
	return pulumix.Output[LookupAdvancedThreatProtectionResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id
func (o LookupAdvancedThreatProtectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdvancedThreatProtectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// Indicates whether Advanced Threat Protection is enabled.
func (o LookupAdvancedThreatProtectionResultOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAdvancedThreatProtectionResult) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// Resource name
func (o LookupAdvancedThreatProtectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdvancedThreatProtectionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource type
func (o LookupAdvancedThreatProtectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAdvancedThreatProtectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAdvancedThreatProtectionResultOutput{})
}
