// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package portal

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the Dashboard.
// Azure REST API version: 2020-09-01-preview.
func LookupDashboard(ctx *pulumi.Context, args *LookupDashboardArgs, opts ...pulumi.InvokeOption) (*LookupDashboardResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDashboardResult
	err := ctx.Invoke("azure-native:portal:getDashboard", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDashboardArgs struct {
	// The name of the dashboard.
	DashboardName string `pulumi:"dashboardName"`
	// The name of the resource group.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// The shared dashboard resource definition.
type LookupDashboardResult struct {
	// Resource Id
	Id string `pulumi:"id"`
	// The dashboard lenses.
	Lenses []DashboardLensResponse `pulumi:"lenses"`
	// Resource location
	Location string `pulumi:"location"`
	// The dashboard metadata.
	Metadata map[string]interface{} `pulumi:"metadata"`
	// Resource name
	Name string `pulumi:"name"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupDashboardOutput(ctx *pulumi.Context, args LookupDashboardOutputArgs, opts ...pulumi.InvokeOption) LookupDashboardResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDashboardResult, error) {
			args := v.(LookupDashboardArgs)
			r, err := LookupDashboard(ctx, &args, opts...)
			var s LookupDashboardResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDashboardResultOutput)
}

type LookupDashboardOutputArgs struct {
	// The name of the dashboard.
	DashboardName pulumi.StringInput `pulumi:"dashboardName"`
	// The name of the resource group.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupDashboardOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDashboardArgs)(nil)).Elem()
}

// The shared dashboard resource definition.
type LookupDashboardResultOutput struct{ *pulumi.OutputState }

func (LookupDashboardResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDashboardResult)(nil)).Elem()
}

func (o LookupDashboardResultOutput) ToLookupDashboardResultOutput() LookupDashboardResultOutput {
	return o
}

func (o LookupDashboardResultOutput) ToLookupDashboardResultOutputWithContext(ctx context.Context) LookupDashboardResultOutput {
	return o
}

func (o LookupDashboardResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDashboardResult] {
	return pulumix.Output[LookupDashboardResult]{
		OutputState: o.OutputState,
	}
}

// Resource Id
func (o LookupDashboardResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Id }).(pulumi.StringOutput)
}

// The dashboard lenses.
func (o LookupDashboardResultOutput) Lenses() DashboardLensResponseArrayOutput {
	return o.ApplyT(func(v LookupDashboardResult) []DashboardLensResponse { return v.Lenses }).(DashboardLensResponseArrayOutput)
}

// Resource location
func (o LookupDashboardResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Location }).(pulumi.StringOutput)
}

// The dashboard metadata.
func (o LookupDashboardResultOutput) Metadata() pulumi.MapOutput {
	return o.ApplyT(func(v LookupDashboardResult) map[string]interface{} { return v.Metadata }).(pulumi.MapOutput)
}

// Resource name
func (o LookupDashboardResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Name }).(pulumi.StringOutput)
}

// Resource tags
func (o LookupDashboardResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDashboardResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupDashboardResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDashboardResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDashboardResultOutput{})
}
