// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a KPack build.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview.
func LookupBuildServiceBuild(ctx *pulumi.Context, args *LookupBuildServiceBuildArgs, opts ...pulumi.InvokeOption) (*LookupBuildServiceBuildResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBuildServiceBuildResult
	err := ctx.Invoke("azure-native:appplatform:getBuildServiceBuild", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupBuildServiceBuildArgs struct {
	// The name of the build resource.
	BuildName string `pulumi:"buildName"`
	// The name of the build service resource.
	BuildServiceName string `pulumi:"buildServiceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Build resource payload
type LookupBuildServiceBuildResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of the build resource
	Properties BuildPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupBuildServiceBuildResult
func (val *LookupBuildServiceBuildResult) Defaults() *LookupBuildServiceBuildResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = *tmp.Properties.Defaults()

	return &tmp
}

func LookupBuildServiceBuildOutput(ctx *pulumi.Context, args LookupBuildServiceBuildOutputArgs, opts ...pulumi.InvokeOption) LookupBuildServiceBuildResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBuildServiceBuildResult, error) {
			args := v.(LookupBuildServiceBuildArgs)
			r, err := LookupBuildServiceBuild(ctx, &args, opts...)
			var s LookupBuildServiceBuildResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBuildServiceBuildResultOutput)
}

type LookupBuildServiceBuildOutputArgs struct {
	// The name of the build resource.
	BuildName pulumi.StringInput `pulumi:"buildName"`
	// The name of the build service resource.
	BuildServiceName pulumi.StringInput `pulumi:"buildServiceName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupBuildServiceBuildOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildServiceBuildArgs)(nil)).Elem()
}

// Build resource payload
type LookupBuildServiceBuildResultOutput struct{ *pulumi.OutputState }

func (LookupBuildServiceBuildResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildServiceBuildResult)(nil)).Elem()
}

func (o LookupBuildServiceBuildResultOutput) ToLookupBuildServiceBuildResultOutput() LookupBuildServiceBuildResultOutput {
	return o
}

func (o LookupBuildServiceBuildResultOutput) ToLookupBuildServiceBuildResultOutputWithContext(ctx context.Context) LookupBuildServiceBuildResultOutput {
	return o
}

func (o LookupBuildServiceBuildResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupBuildServiceBuildResult] {
	return pulumix.Output[LookupBuildServiceBuildResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource.
func (o LookupBuildServiceBuildResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceBuildResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupBuildServiceBuildResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceBuildResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the build resource
func (o LookupBuildServiceBuildResultOutput) Properties() BuildPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBuildServiceBuildResult) BuildPropertiesResponse { return v.Properties }).(BuildPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupBuildServiceBuildResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBuildServiceBuildResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupBuildServiceBuildResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceBuildResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBuildServiceBuildResultOutput{})
}
