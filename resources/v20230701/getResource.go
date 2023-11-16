// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20230701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a resource.
func LookupResource(ctx *pulumi.Context, args *LookupResourceArgs, opts ...pulumi.InvokeOption) (*LookupResourceResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceResult
	err := ctx.Invoke("azure-native:resources/v20230701:getResource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceArgs struct {
	// The parent resource identity.
	ParentResourcePath string `pulumi:"parentResourcePath"`
	// The name of the resource group containing the resource to get. The name is case insensitive.
	ResourceGroupName string `pulumi:"resourceGroupName"`
	// The name of the resource to get.
	ResourceName string `pulumi:"resourceName"`
	// The namespace of the resource provider.
	ResourceProviderNamespace string `pulumi:"resourceProviderNamespace"`
	// The resource type of the resource.
	ResourceType string `pulumi:"resourceType"`
}

// Resource information.
type LookupResourceResult struct {
	// Resource extended location.
	ExtendedLocation *ExtendedLocationResponse `pulumi:"extendedLocation"`
	// Resource ID
	Id string `pulumi:"id"`
	// The identity of the resource.
	Identity *IdentityResponse `pulumi:"identity"`
	// The kind of the resource.
	Kind *string `pulumi:"kind"`
	// Resource location
	Location *string `pulumi:"location"`
	// ID of the resource that manages this resource.
	ManagedBy *string `pulumi:"managedBy"`
	// Resource name
	Name string `pulumi:"name"`
	// The plan of the resource.
	Plan *PlanResponse `pulumi:"plan"`
	// The resource properties.
	Properties interface{} `pulumi:"properties"`
	// The SKU of the resource.
	Sku *SkuResponse `pulumi:"sku"`
	// Resource tags
	Tags map[string]string `pulumi:"tags"`
	// Resource type
	Type string `pulumi:"type"`
}

func LookupResourceOutput(ctx *pulumi.Context, args LookupResourceOutputArgs, opts ...pulumi.InvokeOption) LookupResourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceResult, error) {
			args := v.(LookupResourceArgs)
			r, err := LookupResource(ctx, &args, opts...)
			var s LookupResourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceResultOutput)
}

type LookupResourceOutputArgs struct {
	// The parent resource identity.
	ParentResourcePath pulumi.StringInput `pulumi:"parentResourcePath"`
	// The name of the resource group containing the resource to get. The name is case insensitive.
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	// The name of the resource to get.
	ResourceName pulumi.StringInput `pulumi:"resourceName"`
	// The namespace of the resource provider.
	ResourceProviderNamespace pulumi.StringInput `pulumi:"resourceProviderNamespace"`
	// The resource type of the resource.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
}

func (LookupResourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceArgs)(nil)).Elem()
}

// Resource information.
type LookupResourceResultOutput struct{ *pulumi.OutputState }

func (LookupResourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceResult)(nil)).Elem()
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutput() LookupResourceResultOutput {
	return o
}

func (o LookupResourceResultOutput) ToLookupResourceResultOutputWithContext(ctx context.Context) LookupResourceResultOutput {
	return o
}

// Resource extended location.
func (o LookupResourceResultOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *ExtendedLocationResponse { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

// Resource ID
func (o LookupResourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Id }).(pulumi.StringOutput)
}

// The identity of the resource.
func (o LookupResourceResultOutput) Identity() IdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *IdentityResponse { return v.Identity }).(IdentityResponsePtrOutput)
}

// The kind of the resource.
func (o LookupResourceResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

// Resource location
func (o LookupResourceResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// ID of the resource that manages this resource.
func (o LookupResourceResultOutput) ManagedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *string { return v.ManagedBy }).(pulumi.StringPtrOutput)
}

// Resource name
func (o LookupResourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Name }).(pulumi.StringOutput)
}

// The plan of the resource.
func (o LookupResourceResultOutput) Plan() PlanResponsePtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *PlanResponse { return v.Plan }).(PlanResponsePtrOutput)
}

// The resource properties.
func (o LookupResourceResultOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupResourceResult) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

// The SKU of the resource.
func (o LookupResourceResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupResourceResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

// Resource tags
func (o LookupResourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupResourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Resource type
func (o LookupResourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceResultOutput{})
}
