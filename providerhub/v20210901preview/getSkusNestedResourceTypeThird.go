// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-azure-native-sdk/v2/utilities"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the sku details for the given resource type and sku name.
func LookupSkusNestedResourceTypeThird(ctx *pulumi.Context, args *LookupSkusNestedResourceTypeThirdArgs, opts ...pulumi.InvokeOption) (*LookupSkusNestedResourceTypeThirdResult, error) {
	opts = utilities.PkgInvokeDefaultOpts(opts)
	var rv LookupSkusNestedResourceTypeThirdResult
	err := ctx.Invoke("azure-native:providerhub/v20210901preview:getSkusNestedResourceTypeThird", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSkusNestedResourceTypeThirdArgs struct {
	// The first child resource type.
	NestedResourceTypeFirst string `pulumi:"nestedResourceTypeFirst"`
	// The second child resource type.
	NestedResourceTypeSecond string `pulumi:"nestedResourceTypeSecond"`
	// The third child resource type.
	NestedResourceTypeThird string `pulumi:"nestedResourceTypeThird"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace string `pulumi:"providerNamespace"`
	// The resource type.
	ResourceType string `pulumi:"resourceType"`
	// The SKU.
	Sku string `pulumi:"sku"`
}

type LookupSkusNestedResourceTypeThirdResult struct {
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

func LookupSkusNestedResourceTypeThirdOutput(ctx *pulumi.Context, args LookupSkusNestedResourceTypeThirdOutputArgs, opts ...pulumi.InvokeOption) LookupSkusNestedResourceTypeThirdResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSkusNestedResourceTypeThirdResult, error) {
			args := v.(LookupSkusNestedResourceTypeThirdArgs)
			r, err := LookupSkusNestedResourceTypeThird(ctx, &args, opts...)
			var s LookupSkusNestedResourceTypeThirdResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSkusNestedResourceTypeThirdResultOutput)
}

type LookupSkusNestedResourceTypeThirdOutputArgs struct {
	// The first child resource type.
	NestedResourceTypeFirst pulumi.StringInput `pulumi:"nestedResourceTypeFirst"`
	// The second child resource type.
	NestedResourceTypeSecond pulumi.StringInput `pulumi:"nestedResourceTypeSecond"`
	// The third child resource type.
	NestedResourceTypeThird pulumi.StringInput `pulumi:"nestedResourceTypeThird"`
	// The name of the resource provider hosted within ProviderHub.
	ProviderNamespace pulumi.StringInput `pulumi:"providerNamespace"`
	// The resource type.
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
	// The SKU.
	Sku pulumi.StringInput `pulumi:"sku"`
}

func (LookupSkusNestedResourceTypeThirdOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusNestedResourceTypeThirdArgs)(nil)).Elem()
}

type LookupSkusNestedResourceTypeThirdResultOutput struct{ *pulumi.OutputState }

func (LookupSkusNestedResourceTypeThirdResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSkusNestedResourceTypeThirdResult)(nil)).Elem()
}

func (o LookupSkusNestedResourceTypeThirdResultOutput) ToLookupSkusNestedResourceTypeThirdResultOutput() LookupSkusNestedResourceTypeThirdResultOutput {
	return o
}

func (o LookupSkusNestedResourceTypeThirdResultOutput) ToLookupSkusNestedResourceTypeThirdResultOutputWithContext(ctx context.Context) LookupSkusNestedResourceTypeThirdResultOutput {
	return o
}

// Fully qualified resource ID for the resource. Ex - /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
func (o LookupSkusNestedResourceTypeThirdResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeThirdResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the resource
func (o LookupSkusNestedResourceTypeThirdResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeThirdResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSkusNestedResourceTypeThirdResultOutput) Properties() SkuResourceResponsePropertiesOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeThirdResult) SkuResourceResponseProperties { return v.Properties }).(SkuResourceResponsePropertiesOutput)
}

// Metadata pertaining to creation and last modification of the resource.
func (o LookupSkusNestedResourceTypeThirdResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeThirdResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

// The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
func (o LookupSkusNestedResourceTypeThirdResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSkusNestedResourceTypeThirdResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSkusNestedResourceTypeThirdResultOutput{})
}
