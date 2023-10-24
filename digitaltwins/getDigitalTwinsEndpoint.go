// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package digitaltwins

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get DigitalTwinsInstances Endpoint.
// Azure REST API version: 2023-01-31.
//
// Other available API versions: 2020-03-01-preview.
func LookupDigitalTwinsEndpoint(ctx *pulumi.Context, args *LookupDigitalTwinsEndpointArgs, opts ...pulumi.InvokeOption) (*LookupDigitalTwinsEndpointResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupDigitalTwinsEndpointResult
	err := ctx.Invoke("azure-native:digitaltwins:getDigitalTwinsEndpoint", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDigitalTwinsEndpointArgs struct {
	// Name of Endpoint Resource.
	EndpointName string `pulumi:"endpointName"`
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the DigitalTwinsInstance.
	ResourceName string `pulumi:"resourceName"`
}

// DigitalTwinsInstance endpoint resource.
type LookupDigitalTwinsEndpointResult struct {
	// The resource identifier.
	Id string `pulumi:"id"`
	// Extension resource name.
	Name string `pulumi:"name"`
	// DigitalTwinsInstance endpoint resource properties.
	Properties interface{} `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupDigitalTwinsEndpointOutput(ctx *pulumi.Context, args LookupDigitalTwinsEndpointOutputArgs, opts ...pulumi.InvokeOption) LookupDigitalTwinsEndpointResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDigitalTwinsEndpointResult, error) {
			args := v.(LookupDigitalTwinsEndpointArgs)
			r, err := LookupDigitalTwinsEndpoint(ctx, &args, opts...)
			var s LookupDigitalTwinsEndpointResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDigitalTwinsEndpointResultOutput)
}

type LookupDigitalTwinsEndpointOutputArgs struct {
	// Name of Endpoint Resource.
	EndpointName pulumi.StringInput `pulumi:"endpointName"`
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the DigitalTwinsInstance.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupDigitalTwinsEndpointOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDigitalTwinsEndpointArgs)(nil)).Elem()
}

// DigitalTwinsInstance endpoint resource.
type LookupDigitalTwinsEndpointResultOutput struct{ *pulumi.OutputState }

func (LookupDigitalTwinsEndpointResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDigitalTwinsEndpointResult)(nil)).Elem()
}

func (o LookupDigitalTwinsEndpointResultOutput) ToLookupDigitalTwinsEndpointResultOutput() LookupDigitalTwinsEndpointResultOutput {
	return o
}

func (o LookupDigitalTwinsEndpointResultOutput) ToLookupDigitalTwinsEndpointResultOutputWithContext(ctx context.Context) LookupDigitalTwinsEndpointResultOutput {
	return o
}

func (o LookupDigitalTwinsEndpointResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupDigitalTwinsEndpointResult] {
	return pulumix.Output[LookupDigitalTwinsEndpointResult]{
		OutputState: o.OutputState,
	}
}

// The resource identifier.
func (o LookupDigitalTwinsEndpointResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinsEndpointResult) string { return v.Id }).(pulumi.StringOutput)
}

// Extension resource name.
func (o LookupDigitalTwinsEndpointResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinsEndpointResult) string { return v.Name }).(pulumi.StringOutput)
}

// DigitalTwinsInstance endpoint resource properties.
func (o LookupDigitalTwinsEndpointResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupDigitalTwinsEndpointResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupDigitalTwinsEndpointResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupDigitalTwinsEndpointResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The resource type.
func (o LookupDigitalTwinsEndpointResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDigitalTwinsEndpointResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDigitalTwinsEndpointResultOutput{})
}
