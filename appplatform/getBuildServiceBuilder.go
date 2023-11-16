// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package appplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a KPack builder.
// Azure REST API version: 2023-05-01-preview.
//
// Other available API versions: 2023-07-01-preview, 2023-09-01-preview, 2023-11-01-preview.
func LookupBuildServiceBuilder(ctx *pulumi.Context, args *LookupBuildServiceBuilderArgs, opts ...pulumi.InvokeOption) (*LookupBuildServiceBuilderResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupBuildServiceBuilderResult
	err := ctx.Invoke("azure-native:appplatform:getBuildServiceBuilder", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBuildServiceBuilderArgs struct {
	// The name of the build service resource.
	BuildServiceName string `pulumi:"buildServiceName"`
	// The name of the builder resource.
	BuilderName string `pulumi:"builderName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// KPack Builder resource
type LookupBuildServiceBuilderResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Property of the Builder resource.
	Properties BuilderPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupBuildServiceBuilderOutput(ctx *pulumi.Context, args LookupBuildServiceBuilderOutputArgs, opts ...pulumi.InvokeOption) LookupBuildServiceBuilderResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBuildServiceBuilderResult, error) {
			args := v.(LookupBuildServiceBuilderArgs)
			r, err := LookupBuildServiceBuilder(ctx, &args, opts...)
			var s LookupBuildServiceBuilderResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBuildServiceBuilderResultOutput)
}

type LookupBuildServiceBuilderOutputArgs struct {
	// The name of the build service resource.
	BuildServiceName pulumi.StringInput `pulumi:"buildServiceName"`
	// The name of the builder resource.
	BuilderName pulumi.StringInput `pulumi:"builderName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupBuildServiceBuilderOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildServiceBuilderArgs)(nil)).Elem()
}

// KPack Builder resource
type LookupBuildServiceBuilderResultOutput struct{ *pulumi.OutputState }

func (LookupBuildServiceBuilderResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBuildServiceBuilderResult)(nil)).Elem()
}

func (o LookupBuildServiceBuilderResultOutput) ToLookupBuildServiceBuilderResultOutput() LookupBuildServiceBuilderResultOutput {
	return o
}

func (o LookupBuildServiceBuilderResultOutput) ToLookupBuildServiceBuilderResultOutputWithContext(ctx context.Context) LookupBuildServiceBuilderResultOutput {
	return o
}

// Fully qualified resource Id for the resource.
func (o LookupBuildServiceBuilderResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceBuilderResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupBuildServiceBuilderResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceBuilderResult) string { return v.Name }).(pulumi.StringOutput)
}

// Property of the Builder resource.
func (o LookupBuildServiceBuilderResultOutput) Properties() BuilderPropertiesResponseOutput {
	return o.ApplyT(func(v LookupBuildServiceBuilderResult) BuilderPropertiesResponse { return v.Properties }).(BuilderPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupBuildServiceBuilderResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupBuildServiceBuilderResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupBuildServiceBuilderResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBuildServiceBuilderResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBuildServiceBuilderResultOutput{})
}
