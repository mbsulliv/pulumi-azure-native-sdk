// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get private endpoint connection properties for the given private endpoint.
func LookupPrivateEndpointConnection(ctx *pulumi.Context, args *LookupPrivateEndpointConnectionArgs, opts ...pulumi.InvokeOption) (*LookupPrivateEndpointConnectionResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupPrivateEndpointConnectionResult
	err := ctx.Invoke("azure-native:digitaltwins/v20201201:getPrivateEndpointConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrivateEndpointConnectionArgs struct {
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName string `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the DigitalTwinsInstance.
	ResourceName string `pulumi:"resourceName"`
}

// The private endpoint connection of a Digital Twin.
type LookupPrivateEndpointConnectionResult struct {
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name       string                                      `pulumi:"name"`
	Properties PrivateEndpointConnectionResponseProperties `pulumi:"properties"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupPrivateEndpointConnectionOutput(ctx *pulumi.Context, args LookupPrivateEndpointConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateEndpointConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateEndpointConnectionResult, error) {
			args := v.(LookupPrivateEndpointConnectionArgs)
			r, err := LookupPrivateEndpointConnection(ctx, &args, opts...)
			var s LookupPrivateEndpointConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateEndpointConnectionResultOutput)
}

type LookupPrivateEndpointConnectionOutputArgs struct {
	// The name of the private endpoint connection.
	PrivateEndpointConnectionName pulumi.StringInput `pulumi:"privateEndpointConnectionName"`
	// The name of the resource group that contains the DigitalTwinsInstance.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the DigitalTwinsInstance.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
}

func (LookupPrivateEndpointConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionArgs)(nil)).Elem()
}

// The private endpoint connection of a Digital Twin.
type LookupPrivateEndpointConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateEndpointConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateEndpointConnectionResult)(nil)).Elem()
}

func (o LookupPrivateEndpointConnectionResultOutput) ToLookupPrivateEndpointConnectionResultOutput() LookupPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionResultOutput) ToLookupPrivateEndpointConnectionResultOutputWithContext(ctx context.Context) LookupPrivateEndpointConnectionResultOutput {
	return o
}

func (o LookupPrivateEndpointConnectionResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupPrivateEndpointConnectionResult] {
	return pulumix.Output[LookupPrivateEndpointConnectionResult]{
		OutputState: o.OutputState,
	}
}

// The resource identifier.
func (o LookupPrivateEndpointConnectionResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupPrivateEndpointConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrivateEndpointConnectionResultOutput) Properties() PrivateEndpointConnectionResponsePropertiesOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) PrivateEndpointConnectionResponseProperties {
		return v.Properties
	}).(PrivateEndpointConnectionResponsePropertiesOutput)
}

// The resource type.
func (o LookupPrivateEndpointConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateEndpointConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateEndpointConnectionResultOutput{})
}
