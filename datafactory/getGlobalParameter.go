// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package datafactory

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Global parameter
// Azure REST API version: 2018-06-01.
func LookupGlobalParameter(ctx *pulumi.Context, args *LookupGlobalParameterArgs, opts ...pulumi.InvokeOption) (*LookupGlobalParameterResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupGlobalParameterResult
	err := ctx.Invoke("azure-native:datafactory:getGlobalParameter", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGlobalParameterArgs struct {
	// The factory name.
	FactoryName string `pulumi:"factoryName"`
	// The global parameter name.
	GlobalParameterName string `pulumi:"globalParameterName"`
	// The resource group name.
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

// Global parameters resource type.
type LookupGlobalParameterResult struct {
	// Etag identifies change in the resource.
	Etag string `pulumi:"etag"`
	// The resource identifier.
	Id string `pulumi:"id"`
	// The resource name.
	Name string `pulumi:"name"`
	// Properties of the global parameter.
	Properties map[string]GlobalParameterSpecificationResponse `pulumi:"properties"`
	// The resource type.
	Type string `pulumi:"type"`
}

func LookupGlobalParameterOutput(ctx *pulumi.Context, args LookupGlobalParameterOutputArgs, opts ...pulumi.InvokeOption) LookupGlobalParameterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGlobalParameterResult, error) {
			args := v.(LookupGlobalParameterArgs)
			r, err := LookupGlobalParameter(ctx, &args, opts...)
			var s LookupGlobalParameterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGlobalParameterResultOutput)
}

type LookupGlobalParameterOutputArgs struct {
	// The factory name.
	FactoryName pulumi.StringInput `pulumi:"factoryName"`
	// The global parameter name.
	GlobalParameterName pulumi.StringInput `pulumi:"globalParameterName"`
	// The resource group name.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupGlobalParameterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalParameterArgs)(nil)).Elem()
}

// Global parameters resource type.
type LookupGlobalParameterResultOutput struct{ *pulumi.OutputState }

func (LookupGlobalParameterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGlobalParameterResult)(nil)).Elem()
}

func (o LookupGlobalParameterResultOutput) ToLookupGlobalParameterResultOutput() LookupGlobalParameterResultOutput {
	return o
}

func (o LookupGlobalParameterResultOutput) ToLookupGlobalParameterResultOutputWithContext(ctx context.Context) LookupGlobalParameterResultOutput {
	return o
}

// Etag identifies change in the resource.
func (o LookupGlobalParameterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalParameterResult) string { return v.Etag }).(pulumi.StringOutput)
}

// The resource identifier.
func (o LookupGlobalParameterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalParameterResult) string { return v.Id }).(pulumi.StringOutput)
}

// The resource name.
func (o LookupGlobalParameterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalParameterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Properties of the global parameter.
func (o LookupGlobalParameterResultOutput) Properties() GlobalParameterSpecificationResponseMapOutput {
	return o.ApplyT(func(v LookupGlobalParameterResult) map[string]GlobalParameterSpecificationResponse {
		return v.Properties
	}).(GlobalParameterSpecificationResponseMapOutput)
}

// The resource type.
func (o LookupGlobalParameterResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGlobalParameterResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGlobalParameterResultOutput{})
}
