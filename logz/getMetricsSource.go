// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package logz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Azure REST API version: 2022-01-01-preview.
func LookupMetricsSource(ctx *pulumi.Context, args *LookupMetricsSourceArgs, opts ...pulumi.InvokeOption) (*LookupMetricsSourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMetricsSourceResult
	err := ctx.Invoke("azure-native:logz:getMetricsSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetricsSourceArgs struct {
	// Metrics Account resource name
	MetricsSourceName string `pulumi:"metricsSourceName"`
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupMetricsSourceResult struct {
	// ARM id of the monitor resource.
	Id       string                      `pulumi:"id"`
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	Location string                      `pulumi:"location"`
	// Name of the monitor resource.
	Name string `pulumi:"name"`
	// Properties specific to the monitor resource.
	Properties MonitorPropertiesResponse `pulumi:"properties"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	Tags       map[string]string  `pulumi:"tags"`
	// The type of the monitor resource.
	Type string `pulumi:"type"`
}

func LookupMetricsSourceOutput(ctx *pulumi.Context, args LookupMetricsSourceOutputArgs, opts ...pulumi.InvokeOption) LookupMetricsSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetricsSourceResult, error) {
			args := v.(LookupMetricsSourceArgs)
			r, err := LookupMetricsSource(ctx, &args, opts...)
			var s LookupMetricsSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetricsSourceResultOutput)
}

type LookupMetricsSourceOutputArgs struct {
	// Metrics Account resource name
	MetricsSourceName pulumi.StringInput `pulumi:"metricsSourceName"`
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMetricsSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricsSourceArgs)(nil)).Elem()
}

type LookupMetricsSourceResultOutput struct{ *pulumi.OutputState }

func (LookupMetricsSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricsSourceResult)(nil)).Elem()
}

func (o LookupMetricsSourceResultOutput) ToLookupMetricsSourceResultOutput() LookupMetricsSourceResultOutput {
	return o
}

func (o LookupMetricsSourceResultOutput) ToLookupMetricsSourceResultOutputWithContext(ctx context.Context) LookupMetricsSourceResultOutput {
	return o
}

func (o LookupMetricsSourceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMetricsSourceResult] {
	return pulumix.Output[LookupMetricsSourceResult]{
		OutputState: o.OutputState,
	}
}

// ARM id of the monitor resource.
func (o LookupMetricsSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMetricsSourceResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupMetricsSourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the monitor resource.
func (o LookupMetricsSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties specific to the monitor resource.
func (o LookupMetricsSourceResultOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) MonitorPropertiesResponse { return v.Properties }).(MonitorPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o LookupMetricsSourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMetricsSourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the monitor resource.
func (o LookupMetricsSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetricsSourceResultOutput{})
}
