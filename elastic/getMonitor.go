// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package elastic

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Monitor resource.
// Azure REST API version: 2023-06-01.
//
// Other available API versions: 2023-06-15-preview, 2023-07-01-preview, 2023-10-01-preview.
func LookupMonitor(ctx *pulumi.Context, args *LookupMonitorArgs, opts ...pulumi.InvokeOption) (*LookupMonitorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupMonitorResult
	err := ctx.Invoke("azure-native:elastic:getMonitor", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitorArgs struct {
	// Monitor resource name
	MonitorName string `pulumi:"monitorName"`
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Monitor resource.
type LookupMonitorResult struct {
	// ARM id of the monitor resource.
	Id string `pulumi:"id"`
	// Identity properties of the monitor resource.
	Identity *IdentityPropertiesResponse `pulumi:"identity"`
	// The location of the monitor resource
	Location string `pulumi:"location"`
	// Name of the monitor resource.
	Name string `pulumi:"name"`
	// Properties of the monitor resource.
	Properties MonitorPropertiesResponse `pulumi:"properties"`
	// SKU of the monitor resource.
	Sku *ResourceSkuResponse `pulumi:"sku"`
	// The system metadata relating to this resource
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The tags of the monitor resource.
	Tags map[string]string `pulumi:"tags"`
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
	// The name of the resource group to which the Elastic resource belongs.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMonitorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitorArgs)(nil)).Elem()
}

// Monitor resource.
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

// ARM id of the monitor resource.
func (o LookupMonitorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Id }).(pulumi.StringOutput)
}

// Identity properties of the monitor resource.
func (o LookupMonitorResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupMonitorResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

// The location of the monitor resource
func (o LookupMonitorResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Location }).(pulumi.StringOutput)
}

// Name of the monitor resource.
func (o LookupMonitorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the monitor resource.
func (o LookupMonitorResultOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMonitorResult) MonitorPropertiesResponse { return v.Properties }).(MonitorPropertiesResponseOutput)
}

// SKU of the monitor resource.
func (o LookupMonitorResultOutput) Sku() ResourceSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupMonitorResult) *ResourceSkuResponse { return v.Sku }).(ResourceSkuResponsePtrOutput)
}

// The system metadata relating to this resource
func (o LookupMonitorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMonitorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The tags of the monitor resource.
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
