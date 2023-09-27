// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Get a Service by name.
func LookupEndpointVariant(ctx *pulumi.Context, args *LookupEndpointVariantArgs, opts ...pulumi.InvokeOption) (*LookupEndpointVariantResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupEndpointVariantResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210401:getEndpointVariant", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupEndpointVariantArgs struct {
	// Set to True to include Model details.
	Expand *bool `pulumi:"expand"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// Name of the Azure Machine Learning service.
	ServiceName string `pulumi:"serviceName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName string `pulumi:"workspaceName"`
}

// Machine Learning service object wrapped into ARM resource envelope.
type LookupEndpointVariantResult struct {
	// Specifies the resource ID.
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// Specifies the location of the resource.
	Location *string `pulumi:"location"`
	// Specifies the name of the resource.
	Name string `pulumi:"name"`
	// Service properties
	Properties interface{} `pulumi:"properties"`
	// The sku of the workspace.
	Sku *SkuResponse `pulumi:"sku"`
	// Read only system data
	SystemData SystemDataResponse `pulumi:"systemData"`
	// Contains resource tags defined as key/value pairs.
	Tags map[string]string `pulumi:"tags"`
	// Specifies the type of the resource.
	Type string `pulumi:"type"`
}

func LookupEndpointVariantOutput(ctx *pulumi.Context, args LookupEndpointVariantOutputArgs, opts ...pulumi.InvokeOption) LookupEndpointVariantResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupEndpointVariantResult, error) {
			args := v.(LookupEndpointVariantArgs)
			r, err := LookupEndpointVariant(ctx, &args, opts...)
			var s LookupEndpointVariantResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupEndpointVariantResultOutput)
}

type LookupEndpointVariantOutputArgs struct {
	// Set to True to include Model details.
	Expand pulumi.BoolPtrInput `pulumi:"expand"`
	// Name of the resource group in which workspace is located.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// Name of the Azure Machine Learning service.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
	// Name of Azure Machine Learning workspace.
	WorkspaceName pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupEndpointVariantOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointVariantArgs)(nil)).Elem()
}

// Machine Learning service object wrapped into ARM resource envelope.
type LookupEndpointVariantResultOutput struct{ *pulumi.OutputState }

func (LookupEndpointVariantResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEndpointVariantResult)(nil)).Elem()
}

func (o LookupEndpointVariantResultOutput) ToLookupEndpointVariantResultOutput() LookupEndpointVariantResultOutput {
	return o
}

func (o LookupEndpointVariantResultOutput) ToLookupEndpointVariantResultOutputWithContext(ctx context.Context) LookupEndpointVariantResultOutput {
	return o
}

func (o LookupEndpointVariantResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupEndpointVariantResult] {
	return pulumix.Output[LookupEndpointVariantResult]{
		OutputState: o.OutputState,
	}
}

// Specifies the resource ID.
func (o LookupEndpointVariantResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupEndpointVariantResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// Specifies the location of the resource.
func (o LookupEndpointVariantResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Specifies the name of the resource.
func (o LookupEndpointVariantResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) string { return v.Name }).(pulumi.StringOutput)
}

// Service properties
func (o LookupEndpointVariantResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// The sku of the workspace.
func (o LookupEndpointVariantResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Read only system data
func (o LookupEndpointVariantResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// Contains resource tags defined as key/value pairs.
func (o LookupEndpointVariantResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Specifies the type of the resource.
func (o LookupEndpointVariantResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEndpointVariantResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEndpointVariantResultOutput{})
}
