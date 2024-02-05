// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20240101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get the APM by name.
func LookupApm(ctx *pulumi.Context, args *LookupApmArgs, opts ...pulumi.InvokeOption) (*LookupApmResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApmResult
	err := ctx.Invoke("azure-native:appplatform/v20240101preview:getApm", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApmArgs struct {
	// The name of the APM
	ApmName string `pulumi:"apmName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// APM Resource object
type LookupApmResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Properties of an APM
	Properties ApmPropertiesResponse `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

func LookupApmOutput(ctx *pulumi.Context, args LookupApmOutputArgs, opts ...pulumi.InvokeOption) LookupApmResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApmResult, error) {
			args := v.(LookupApmArgs)
			r, err := LookupApm(ctx, &args, opts...)
			var s LookupApmResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApmResultOutput)
}

type LookupApmOutputArgs struct {
	// The name of the APM
	ApmName pulumi.StringInput `pulumi:"apmName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApmOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApmArgs)(nil)).Elem()
}

// APM Resource object
type LookupApmResultOutput struct{ *pulumi.OutputState }

func (LookupApmResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApmResult)(nil)).Elem()
}

func (o LookupApmResultOutput) ToLookupApmResultOutput() LookupApmResultOutput {
	return o
}

func (o LookupApmResultOutput) ToLookupApmResultOutputWithContext(ctx context.Context) LookupApmResultOutput {
	return o
}

// Fully qualified resource Id for the resource.
func (o LookupApmResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApmResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupApmResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApmResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of an APM
func (o LookupApmResultOutput) Properties() ApmPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApmResult) ApmPropertiesResponse { return v.Properties }).(ApmPropertiesResponseOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupApmResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApmResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupApmResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApmResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApmResultOutput{})
}
