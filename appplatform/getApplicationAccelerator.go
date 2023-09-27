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

// Get the application accelerator.
// Azure REST API version: 2023-05-01-preview.
func LookupApplicationAccelerator(ctx *pulumi.Context, args *LookupApplicationAcceleratorArgs, opts ...pulumi.InvokeOption) (*LookupApplicationAcceleratorResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupApplicationAcceleratorResult
	err := ctx.Invoke("azure-native:appplatform:getApplicationAccelerator", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupApplicationAcceleratorArgs struct {
	// The name of the application accelerator.
	ApplicationAcceleratorName string `pulumi:"applicationAcceleratorName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName string `pulumi:"serviceName"`
}

// Application accelerator resource
type LookupApplicationAcceleratorResult struct {
	// Fully qualified resource Id for the resource.
	Id string `pulumi:"id"`
	// The name of the resource.
	Name string `pulumi:"name"`
	// Application accelerator properties payload
	Properties ApplicationAcceleratorPropertiesResponse `pulumi:"properties"`
	// Sku of the application accelerator resource
	Sku *SkuResponse `pulumi:"sku"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource.
	Type string `pulumi:"type"`
}

// Defaults sets the appropriate defaults for LookupApplicationAcceleratorResult
func (val *LookupApplicationAcceleratorResult) Defaults() *LookupApplicationAcceleratorResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}

func LookupApplicationAcceleratorOutput(ctx *pulumi.Context, args LookupApplicationAcceleratorOutputArgs, opts ...pulumi.InvokeOption) LookupApplicationAcceleratorResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApplicationAcceleratorResult, error) {
			args := v.(LookupApplicationAcceleratorArgs)
			r, err := LookupApplicationAccelerator(ctx, &args, opts...)
			var s LookupApplicationAcceleratorResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApplicationAcceleratorResultOutput)
}

type LookupApplicationAcceleratorOutputArgs struct {
	// The name of the application accelerator.
	ApplicationAcceleratorName pulumi.StringInput `pulumi:"applicationAcceleratorName"`
	// The name of the resource group that contains the resource. You can obtain this value from the Azure Resource Manager API or the portal.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the Service resource.
	ServiceName pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApplicationAcceleratorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationAcceleratorArgs)(nil)).Elem()
}

// Application accelerator resource
type LookupApplicationAcceleratorResultOutput struct{ *pulumi.OutputState }

func (LookupApplicationAcceleratorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApplicationAcceleratorResult)(nil)).Elem()
}

func (o LookupApplicationAcceleratorResultOutput) ToLookupApplicationAcceleratorResultOutput() LookupApplicationAcceleratorResultOutput {
	return o
}

func (o LookupApplicationAcceleratorResultOutput) ToLookupApplicationAcceleratorResultOutputWithContext(ctx context.Context) LookupApplicationAcceleratorResultOutput {
	return o
}

func (o LookupApplicationAcceleratorResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupApplicationAcceleratorResult] {
	return pulumix.Output[LookupApplicationAcceleratorResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource Id for the resource.
func (o LookupApplicationAcceleratorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource.
func (o LookupApplicationAcceleratorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) string { return v.Name }).(pulumi.StringOutput)
}

// Application accelerator properties payload
func (o LookupApplicationAcceleratorResultOutput) Properties() ApplicationAcceleratorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) ApplicationAcceleratorPropertiesResponse {
		return v.Properties
	}).(ApplicationAcceleratorPropertiesResponseOutput)
}

// Sku of the application accelerator resource
func (o LookupApplicationAcceleratorResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupApplicationAcceleratorResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource.
func (o LookupApplicationAcceleratorResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApplicationAcceleratorResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApplicationAcceleratorResultOutput{})
}
