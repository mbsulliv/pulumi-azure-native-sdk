// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package providerhub

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Gets the sku details for the given resource type and sku name.
// Azure REST API version: 2021-09-01-preview.
func LookupSkusNestedResourceTypeSecond(ctx *pulumi.Context, args *LookupSkusNestedResourceTypeSecondArgs, opts ...pulumi.InvokeOption) (*LookupSkusNestedResourceTypeSecondResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSkusNestedResourceTypeSecondResult
	err := ctx.Invoke("azure-native:providerhub:getSkusNestedResourceTypeSecond", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSkusNestedResourceTypeSecondArgs struct {
	// The first child resource type.
	NestedResourceTypeFirst string `pulumi:"nestedResourceTypeFirst"`
	// The second child resource type.
	NestedResourceTypeSecond string `pulumi:"nestedResourceTypeSecond"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The resource type.
	ResourceType string `pulumi:"resourceType"`
	// The SKU.
	Sku string `pulumi:"sku"`
}

type LookupSkusNestedResourceTypeSecondResult struct {
	// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id string `pulumi:"id"`
	// The name of the resource
	Name       string                        `pulumi:"name"`
	Properties SkuResourceResponseProperties `pulumi:"properties"`
	// Metadata pertaining to creation and last modification of the resource.
	SystemData SystemDataResponse `pulumi:"systemData"`
	// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type string `pulumi:"type"`
}

func LookupSkusNestedResourceTypeSecondOutput(ctx *pulumi.Context, args LookupSkusNestedResourceTypeSecondOutputArgs, opts ...pulumi.InvokeOption) LookupSkusNestedResourceTypeSecondResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSkusNestedResourceTypeSecondResult, error) {
			args := v.(LookupSkusNestedResourceTypeSecondArgs)
			r, err := LookupSkusNestedResourceTypeSecond(ctx, &args, opts...)
			var s LookupSkusNestedResourceTypeSecondResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSkusNestedResourceTypeSecondResultOutput)
}

type LookupSkusNestedResourceTypeSecondOutputArgs struct {
	// The first child resource type.
	NestedResourceTypeFirst pulumi.StringInput `pulumi:"nestedResourceTypeFirst"`
	// The second child resource type.
	NestedResourceTypeSecond pulumi.StringInput `pulumi:"nestedResourceTypeSecond"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
	// The resource type.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// The SKU.
	Sku pulumi.StringInput `pulumi:"sku"`
}

func (LookupSkusNestedResourceTypeSecondOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusNestedResourceTypeSecondArgs)(nil)).Elem()
}

type LookupSkusNestedResourceTypeSecondResultOutput struct{ *pulumi.OutputState }

func (LookupSkusNestedResourceTypeSecondResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusNestedResourceTypeSecondResult)(nil)).Elem()
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) ToLookupSkusNestedResourceTypeSecondResultOutput() LookupSkusNestedResourceTypeSecondResultOutput {
	return o
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) ToLookupSkusNestedResourceTypeSecondResultOutputWithContext(ctx context.Context) LookupSkusNestedResourceTypeSecondResultOutput {
	return o
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupSkusNestedResourceTypeSecondResult] {
	return pulumix.Output[LookupSkusNestedResourceTypeSecondResult]{
		OutputState: o.OutputState,
	}
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSkusNestedResourceTypeSecondResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeSecondResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSkusNestedResourceTypeSecondResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeSecondResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSkusNestedResourceTypeSecondResultOutput) Properties() SkuResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeSecondResult) SkuResourceResponseProperties { return v.Properties }).(SkuResourceResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupSkusNestedResourceTypeSecondResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeSecondResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSkusNestedResourceTypeSecondResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeSecondResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSkusNestedResourceTypeSecondResultOutput{})
}
