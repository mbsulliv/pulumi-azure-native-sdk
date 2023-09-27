// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

func LookupMonitor(ctx *pulumi.Context, args *LookupMonitorArgs, opts ...pulumi.InvokeOption) (*LookupMonitorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMonitorResult
	err := ctx.Invoke("azure-native:logz/v20220101preview:getMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitorArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupMonitorResult struct {
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

func LookupMonitorOutput(ctx *pulumi.Context, args LookupMonitorOutputArgs, opts ...pulumi.InvokeOption) LookupMonitorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMonitorResult, error) {
			args := v.(LookupMonitorArgs)
			r, err := LookupMonitor(ctx, &args, opts...)
			var s LookupMonitorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMonitorResultOutput)
}

type LookupMonitorOutputArgs struct {
	// Monitor resource name
	MonitorName pulumi.StringInput `pulumi:"monitorName"`
	// The name of the resource group. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMonitorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitorArgs)(nil)).Elem()
}

type LookupMonitorResultOutput struct{ *pulumi.OutputState }

func (LookupMonitorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitorResult)(nil)).Elem()
}

func (o LookupMonitorResultOutput) ToLookupMonitorResultOutput() LookupMonitorResultOutput {
	return o
}

func (o LookupMonitorResultOutput) ToLookupMonitorResultOutputWithContext(ctx context.Context) LookupMonitorResultOutput {
	return o
}

func (o LookupMonitorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupMonitorResult] {
	return pulumix.Output[LookupMonitorResult]{
		OutputState: o.OutputState,
	}
}

// ARM id of the monitor resource.
func (o LookupMonitorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMonitorResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupMonitorResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupMonitorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the monitor resource.
func (o LookupMonitorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties specific to the monitor resource.
func (o LookupMonitorResultOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMonitorResult) MonitorPropertiesResponse { return v.Properties }).(MonitorPropertiesResponseOutput)
}

// The system metadata relating to this resource
func (o LookupMonitorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMonitorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMonitorResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMonitorResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// The type of the monitor resource.
func (o LookupMonitorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMonitorResultOutput{})
}
